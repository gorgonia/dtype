package dtype

import (
	"unsafe"
)

// oh how nice it'd be if I could make them immutable
var (
	Bool       = Dtype[bool]{}
	Int        = Dtype[int]{}
	Int8       = Dtype[int8]{}
	Int16      = Dtype[int16]{}
	Int32      = Dtype[int32]{}
	Int64      = Dtype[int64]{}
	Uint       = Dtype[uint]{}
	Uint8      = Dtype[uint8]{}
	Uint16     = Dtype[uint16]{}
	Uint32     = Dtype[uint32]{}
	Uint64     = Dtype[uint64]{}
	Float32    = Dtype[float32]{}
	Float64    = Dtype[float64]{}
	Complex64  = Dtype[complex64]{}
	Complex128 = Dtype[complex128]{}
	String     = Dtype[string]{}

	// aliases
	Byte = Uint8

	// extras
	Uintptr       = Dtype[uintptr]{}
	UnsafePointer = Dtype[unsafe.Pointer]{}
)
