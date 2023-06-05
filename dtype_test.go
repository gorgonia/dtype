package dtype

import (
	"testing"
)

type Float16 uint16

func TestRegisterType(t *testing.T) {
	dt := Dtype[Float16]{}
	RegisterFloat(dt)

	if err := typeclassCheck(dt, floatTypes); err != nil {
		t.Errorf("Expected %v to be in floatTypes: %v", dt, err)
	}
	if err := typeclassCheck(dt, numberTypes); err != nil {
		t.Errorf("Expected %v to be in numberTypes: %v", dt, err)
	}
	if err := typeclassCheck(dt, ordTypes); err != nil {
		t.Errorf("Expected %v to be in ordTypes: %v", dt, err)
	}
	if err := typeclassCheck(dt, eqTypes); err != nil {
		t.Errorf("Expected %v to be in eqTypes: %v", dt, err)
	}
}

func TestDtypeConversions(t *testing.T) {
	for k, v := range reverseNumpyDtypes {
		if npdt, err := v.NumpyDtype(); npdt != k {
			t.Errorf("Expected %v to return numpy dtype of %q. Got %q instead", v, k, npdt)
		} else if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
	dt := Dtype[Float16]{}
	if _, err := dt.NumpyDtype(); err == nil {
		t.Errorf("Expected an error when passing in type unknown to np")
	}

	for k, v := range numpyDtypes {
		if dt, err := FromNumpyDtype(v); dt != k {
			// special cases
			if Int.Size() == 4 && v == "i4" && dt == Int {
				continue
			}
			if Int.Size() == 8 && v == "i8" && dt == Int {
				continue
			}

			if Uint.Size() == 4 && v == "u4" && dt == Uint {
				continue
			}
			if Uint.Size() == 8 && v == "u8" && dt == Uint {
				continue
			}
			t.Errorf("Expected %q to return %v. Got %v instead", v, k, dt)
		} else if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
	if _, err := FromNumpyDtype("EDIUH"); err == nil {
		t.Error("Expected error when nonsense is passed into fromNumpyDtype")
	}
}

func TestAllTypes(t *testing.T) {
	for _, tc := range []*typeclass{
		specializedTypes,
		addableTypes,
		numberTypes,
		ordTypes,
		eqTypes,
		signedTypes,
		unsignedTypes,
		signedNonComplexTypes,
		floatTypes,
		complexTypes,
		floatcmplxTypes,
		nonComplexNumberTypes,
		generatableTypes,
	} {
		tc.Lock()
		for _, typ := range tc.set {
			if ID(typ) == -1 {
				t.Errorf("Dtype %v has no ID in allTypes. It is not properly registered", typ)
			}
		}
		tc.Unlock()
	}
}

func TestFindByName(t *testing.T) {
	dt, err := FindByName("float64")
	if err != nil {
		t.Errorf("Expected \"float64\" to be found")
	}

	if dt != Float64 {
		t.Errorf("Got a different dtype than expected")
	}

	_, err = FindByName("f00b4rb4z")
	if err == nil {
		t.Errorf("Expected the Dtype named \"f00b4rb4z\" to not be found ")
	}
}

func TestRegisterFloat(t *testing.T) {
	// this is a repeat test, to test repeated additions of a given dtype
	dt := Dtype[Float16]{}
	RegisterFloat(dt)
	RegisterFloat(dt)

	var count int
	floatTypes.Lock()
	for _, ft := range floatTypes.set {
		if ft == dt {
			count++
		}
	}
	floatTypes.Unlock()
	if count != 1 {
		t.Errorf("Expected Float16 to only exist once in the float types set")
	}
}

func TestTypeClassCheck(t *testing.T) {
	dt := Float64

	cases := []struct {
		TypeClass
		willerr bool
	}{
		{Number, false},
		{maxtypeclass, true},
		{Unsigned, true},
		{-1, false},
	}

	for _, tc := range cases {
		err := TypeClassCheck(dt, tc.TypeClass)
		switch {
		case tc.willerr && err == nil:
			t.Errorf("Expected Float64 in %v to error.", tc.TypeClass)
		case !tc.willerr && err != nil:
			t.Errorf("Expected Float64 in %v to not error", tc.TypeClass)
		}
	}
}
