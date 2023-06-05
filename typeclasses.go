package dtype

type TypeClass int

const (
	All TypeClass = iota
	Specialized
	Addable
	Number
	Ord
	Eq
	Unsigned
	Signed
	SignedNonComplex
	Floats
	Complexes
	FloatComplex
	NonComplexNumber
	Generatable

	maxtypeclass
)

var typeclasses = [...]*typeclass{
	allTypes,
	specializedTypes,
	addableTypes,
	numberTypes,
	ordTypes,
	eqTypes,
	unsignedTypes,
	signedTypes,
	signedNonComplexTypes,
	floatTypes,
	complexTypes,
	floatcmplxTypes,
	nonComplexNumberTypes,
	generatableTypes,
}

// allTypes for indexing
var allTypes = &typeclass{
	name: "τ",
	set: []Datatype{
		Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Complex64, Complex128, String, Uintptr, UnsafePointer,
	},
}

// specialized types indicate that there are specialized code generated for these types
var specializedTypes = &typeclass{
	name: "Specialized",
	set: []Datatype{
		Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Complex64, Complex128, String,
	},
}

var addableTypes = &typeclass{
	name: "Addable",
	set: []Datatype{
		Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Complex64, Complex128, String,
	},
}

var numberTypes = &typeclass{
	name: "Number",
	set: []Datatype{
		Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Complex64, Complex128,
	},
}

var ordTypes = &typeclass{
	name: "Ord",
	set: []Datatype{
		Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, String,
	},
}

var eqTypes = &typeclass{
	name: "Eq",
	set: []Datatype{
		Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Complex64, Complex128, String, Uintptr, UnsafePointer,
	},
}

var unsignedTypes = &typeclass{
	name: "Unsigned",
	set:  []Datatype{Uint, Uint8, Uint16, Uint32, Uint64},
}

var signedTypes = &typeclass{
	name: "Signed",
	set: []Datatype{
		Int, Int8, Int16, Int32, Int64, Float32, Float64, Complex64, Complex128,
	},
}

// this typeclass is ever only used by Sub tests
var signedNonComplexTypes = &typeclass{
	name: "Signed NonComplex",
	set: []Datatype{
		Int, Int8, Int16, Int32, Int64, Float32, Float64,
	},
}

var floatTypes = &typeclass{
	name: "Float",
	set: []Datatype{
		Float32, Float64,
	},
}

var complexTypes = &typeclass{
	name: "Complex Numbers",
	set:  []Datatype{Complex64, Complex128},
}

var floatcmplxTypes = &typeclass{
	name: "Real",
	set: []Datatype{
		Float32, Float64, Complex64, Complex128,
	},
}

var nonComplexNumberTypes = &typeclass{
	name: "Non complex numbers",
	set: []Datatype{
		Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64,
	},
}

// this typeclass is ever only used by Pow tests
var generatableTypes = &typeclass{
	name: "Generatable types",
	set: []Datatype{
		Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, String,
	},
}
