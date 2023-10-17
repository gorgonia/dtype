package dtype

import (
	"fmt"

	"github.com/pkg/errors"
)

// ConsFromInt represents a constructor of a value of the given type.
type ConsFromInt func(int) interface{}

// FromInt is a function to create a value of the given dtype from a given int.
func FromInt(a Dtype, v int) (interface{}, error) {
	fn, ok := fromInt[a]
	if !ok {
		return nil, errors.Errorf("No ConsFromInt constructor found to be registered for %v", a)
	}
	return fn(v), nil
}

var fromInt = map[Dtype]ConsFromInt{
	Int:        intFromInt,
	Int8:       int8FromInt,
	Int16:      int16FromInt,
	Int32:      int32FromInt,
	Int64:      int64FromInt,
	Uint:       uintFromInt,
	Uint8:      uint8FromInt,
	Uint16:     uint16FromInt,
	Uint32:     uint32FromInt,
	Uint64:     uint64FromInt,
	Float32:    float32FromInt,
	Float64:    float64FromInt,
	Complex64:  complex64FromInt,
	Complex128: complex128FromInt,
	String:     stringFromInt,
}

func intFromInt(a int) interface{}        { return a }
func int8FromInt(a int) interface{}       { return int8(a) }
func int16FromInt(a int) interface{}      { return int16(a) }
func int32FromInt(a int) interface{}      { return int32(a) }
func int64FromInt(a int) interface{}      { return int64(a) }
func uintFromInt(a int) interface{}       { return uint(a) }
func uint8FromInt(a int) interface{}      { return uint8(a) }
func uint16FromInt(a int) interface{}     { return uint16(a) }
func uint32FromInt(a int) interface{}     { return uint32(a) }
func uint64FromInt(a int) interface{}     { return uint64(a) }
func float32FromInt(a int) interface{}    { return float32(a) }
func float64FromInt(a int) interface{}    { return float64(a) }
func complex64FromInt(a int) interface{}  { return complex64(complex(float64(a), 0)) }
func complex128FromInt(a int) interface{} { return complex128(complex(float64(a), 0)) }
func stringFromInt(a int) interface{}     { return fmt.Sprintf("%d", a) }
