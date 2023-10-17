package dtype

import (
	"unsafe"
)

// oh how nice it'd be if I could make them immutable
var (
	Bool       = Datatype[bool]{}
	Int        = Datatype[int]{}
	Int8       = Datatype[int8]{}
	Int16      = Datatype[int16]{}
	Int32      = Datatype[int32]{}
	Int64      = Datatype[int64]{}
	Uint       = Datatype[uint]{}
	Uint8      = Datatype[uint8]{}
	Uint16     = Datatype[uint16]{}
	Uint32     = Datatype[uint32]{}
	Uint64     = Datatype[uint64]{}
	Float32    = Datatype[float32]{}
	Float64    = Datatype[float64]{}
	Complex64  = Datatype[complex64]{}
	Complex128 = Datatype[complex128]{}
	String     = Datatype[string]{}

	// aliases
	Byte = Uint8

	// extras
	Uintptr       = Datatype[uintptr]{}
	UnsafePointer = Datatype[unsafe.Pointer]{}
)
