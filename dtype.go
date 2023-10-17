// Package dtype provides a definition of a Dtype, which is a part of the type system that Gorgonia uses.
package dtype // import "gorgonia.org/dtype"

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
)

type readerInto interface {
	Read(ptrToSliceElem any)
	Err() error
}

// Datatype represents a data type of a Tensor. Concretely it's implemented as an embedded reflect.Type
// which allows for easy reflection operations. It also implements hm.Type, for type inference in Gorgonia
type Datatype[DT any] struct{}

func (dt Datatype[DT]) Name() string                                  { var v DT; return reflect.TypeOf(v).Name() }
func (dt Datatype[DT]) String() string                                { return dt.Name() }
func (dt Datatype[DT]) Size() uintptr                                 { var v DT; return unsafe.Sizeof(v) }
func (dt Datatype[DT]) Kind() reflect.Kind                            { var v DT; return reflect.TypeOf(v).Kind() }
func (dt Datatype[DT]) Apply(hm.Subs) hm.Substitutable                { return dt }
func (dt Datatype[DT]) FreeTypeVar() hm.TypeVarSet                    { return nil }
func (dt Datatype[DT]) Normalize(k, v hm.TypeVarSet) (hm.Type, error) { return dt, nil }
func (dt Datatype[DT]) Types() hm.Types                               { return nil }
func (dt Datatype[DT]) Format(s fmt.State, c rune)                    { fmt.Fprintf(s, "%s", dt.Name()) }
func (dt Datatype[DT]) Eq(other hm.Type) bool                         { return other == dt }

// SliceOf creates a slice of the given datatype with n elements.
// This method when working using only interfaces (such as when serializing and unserializing data).
func (dt Datatype[DT]) SliceOf(n int) any { return make([]DT, n) }

// ReadIntoSlice is useful for reading values into `data`, which has to be a []DT. This method is mainly used for serialization and deserialization.
func (dt Datatype[DT]) ReadIntoSlice(slice any, reader readerInto) {

	var v DT
	switch any(v).(type) {
	case int:
		s := slice.([]int)
		// variable sized int. So we gotta fake it
		var i64 int64
		for i := range s {
			reader.Read(&i64)
			s[i] = int(i64)
		}

	case uint:
		s := slice.([]uint)
		// variable sized uint. So we gotta fake it
		var u64 uint64
		for i := range s {
			reader.Read(&u64)
			s[i] = uint(u64)
		}
	default:
		s := slice.([]DT)
		for i := range s {
			reader.Read(&s[i])
		}
	}

}

// Dtype is the type-erased version of a Dtype. One may also think of it as a .... type variable!
type Dtype interface {
	hm.Type
	Kind() reflect.Kind
	Size() uintptr

	NumpyDtype() (string, error)

	// SliceOf creates a slice of the given datatype with n elements.
	// This method when working using only interfaces.
	SliceOf(n int) any

	// ReadIntoSlice is useful for reading data into slice
	ReadIntoSlice(data any, reader readerInto)
}

var numpyDtypes map[Dtype]string
var reverseNumpyDtypes map[string]Dtype

func init() {
	numpyDtypes = map[Dtype]string{
		Bool:       "b1",
		Int:        fmt.Sprintf("i%d", Int.Size()),
		Int8:       "i1",
		Int16:      "i2",
		Int32:      "i4",
		Int64:      "i8",
		Uint:       fmt.Sprintf("u%d", Uint.Size()),
		Uint8:      "u1",
		Uint16:     "u2",
		Uint32:     "u4",
		Uint64:     "u8",
		Float32:    "f4",
		Float64:    "f8",
		Complex64:  "c8",
		Complex128: "c16",
	}

	reverseNumpyDtypes = map[string]Dtype{
		"b1":  Bool,
		"i1":  Int8,
		"i2":  Int16,
		"i4":  Int32,
		"i8":  Int64,
		"u1":  Uint8,
		"u2":  Uint16,
		"u4":  Uint32,
		"u8":  Uint64,
		"f4":  Float32,
		"f8":  Float64,
		"c8":  Complex64,
		"c16": Complex128,
	}
}

// NumpyDtype returns the Numpy's Dtype equivalent. This is predominantly used in converting a Tensor to a Numpy ndarray,
// however, not all Dtypes are supported
func (dt Datatype[DT]) NumpyDtype() (string, error) {
	retVal, ok := numpyDtypes[dt]
	if !ok {
		return "v", errors.Errorf("Unsupported Dtype conversion to Numpy Dtype: %v", dt)
	}
	return retVal, nil
}

// FromNumpyDtype returns a Dtype given a string that matches Numpy's Dtype.
func FromNumpyDtype(t string) (Dtype, error) {
	retVal, ok := reverseNumpyDtypes[t]
	if !ok {
		return nil, errors.Errorf("Unsupported Dtype conversion from %q to Dtype", t)
	}
	if t == "i4" && Int.Size() == 4 {
		return Int, nil
	}
	if t == "i8" && Int.Size() == 8 {
		return Int, nil
	}
	if t == "u4" && Uint.Size() == 4 {
		return Uint, nil
	}
	if t == "u8" && Uint.Size() == 8 {
		return Uint, nil
	}
	return retVal, nil
}

type typeclass struct {
	name string
	set  []Dtype

	sync.Mutex
}

// FindByName finds a given type by its name.
func FindByName(name string) (Dtype, error) {
	for _, dt := range allTypes.set {
		if dt.String() == name {
			return dt, nil
		}
	}
	return nil, errors.Errorf("Cannot find a Dtype named %q. Perhaps it hasn't been registered? Use dtype.Register() to register custom Dtypes.", name)
}

// TypeClassCheck checks if a given Dtype is in the given type class.
// It returns nil if it is in the given type class.
func TypeClassCheck(a Dtype, in TypeClass) error {
	if in >= maxtypeclass {
		return errors.Errorf("Unknown/Unsupported typeclass to check")
	}
	var tc *typeclass
	if in >= All {
		tc = typeclasses[in]
	}
	return typeclassCheck(a, tc)
}

func typeclassCheck(a Dtype, tc *typeclass) error {
	if tc == nil {
		return nil
	}
	tc.Lock()
	defer tc.Unlock()
	for _, s := range tc.set {
		if s == a {
			return nil
		}
	}
	return errors.Errorf("Type %v is not a member of %v", a, tc.name)
}

// RegisterNumber is a function required to register a new numerical Dtype.
// This package provides the following Dtype:
//
//	Int
//	Int8
//	Int16
//	Int32
//	Int64
//	Uint
//	Uint8
//	Uint16
//	Uint32
//	Uint64
//	Float32
//	Float64
//	Complex64
//	Complex128
//
// If a Dtype that is registered already exists on the list, it will not be added to the list.
func RegisterNumber(a Dtype, constructor ConsFromInt) {
	numberTypes.Lock()
	defer numberTypes.Unlock()
	for _, dt := range numberTypes.set {
		if dt == a {
			return
		}
	}
	if constructor != nil {
		fromInt[a] = constructor
	}
	numberTypes.set = append(numberTypes.set, a)
	RegisterEq(a)
}

// RegisterFloat registers a dtype as a type whose values are floating points.
// This implies that NaN, +Inf and -Inf are also well as values in this type.
func RegisterFloat(a Dtype) {
	floatTypes.Lock()
	defer floatTypes.Unlock()
	for _, dt := range floatTypes.set {
		if dt == a {
			return
		}
	}
	floatTypes.set = append(floatTypes.set, a)
	RegisterNumber(a, nil)
	RegisterOrd(a)
}

// RegisterOrd registers a dtype as a type whose values can be ordered.
func RegisterOrd(a Dtype) {
	ordTypes.Lock()
	defer ordTypes.Unlock()
	for _, dt := range ordTypes.set {
		if dt == a {
			return
		}
	}
	ordTypes.set = append(ordTypes.set, a)
	RegisterEq(a)
}

// RegisterEq registers a dtype as a type whose values can be compared for equality.
func RegisterEq(a Dtype) {
	eqTypes.Lock()
	defer eqTypes.Unlock()
	for _, dt := range eqTypes.set {
		if dt == a {
			return
		}
	}
	eqTypes.set = append(eqTypes.set, a)
	Register(a)
}

// Register registers a new Dtype into the registry.
func Register(a Dtype) {
	allTypes.Lock()
	defer allTypes.Unlock()
	for _, dt := range allTypes.set {
		if a == dt {
			return
		}
	}
	allTypes.set = append(allTypes.set, a)
}

// ID returns the ID of the Dtype in the registry.
func ID(a Dtype) int {
	allTypes.Lock()
	defer allTypes.Unlock()
	for i, v := range allTypes.set {
		if a == v {
			return i
		}
	}
	return -1
}
