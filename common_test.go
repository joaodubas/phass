package phass

import (
	"math"
	"testing"
)

func TestEquationValidation(t *testing.T) {
	cases := []struct {
		in  InParams
		ok  bool
		err error
	}{
		{
			in:  map[string]float64{},
			ok:  false,
			err: ErrMissingMeasure,
		},
		{
			in:  map[string]float64{"age": 18.0},
			ok:  false,
			err: ErrMissingMeasure,
		},
		{
			in:  map[string]float64{"sskf": 112.1},
			ok:  false,
			err: ErrMissingMeasure,
		},
		{
			in:  map[string]float64{"age": 17.9, "sskf": 112.1},
			ok:  false,
			err: ErrInvalidAge,
		},
		{
			in:  map[string]float64{"age": 20, "sskf": 109.2},
			ok:  true,
			err: nil,
		},
	}
	for _, data := range cases {
		eq := NewEquation(data.in, conf)
		if v, err := eq.Validate(); v != data.ok {
			t.Error("Should get an error")
		} else if err != data.err {
			t.Error("Should get proper error message")
		}

		// NOTE: this checks if calc returns 0 when an error occurs.
		if !data.ok {
			if v, err := eq.Calc(); v > 0.0 {
				t.Error("Value should be 0")
			} else if err != data.err {
				t.Error("Should get proper error message.")
			}
		}
	}
}

func TestEquationRetrieveInputParameters(t *testing.T) {
	in := map[string]float64{
		"age":  18.8,
		"sskf": 210.1,
	}
	eq := NewEquation(in, conf)
	for k, v := range in {
		if sv, ok := eq.In(k); !ok {
			t.Errorf("Key %s should be defined", k)
		} else if v != sv {
			t.Errorf("Key %s should have value %.2f", k, v)
		}
	}
	if _, ok := eq.In("iDoNotExist"); ok {
		t.Error("Key should no be available")
	}
}

func TestAgeValidator(t *testing.T) {
	cases := []struct {
		in  InParams
		ok  bool
		err error
	}{
		{
			in:  map[string]float64{},
			ok:  false,
			err: ErrMissingAge,
		},
		{
			in:  map[string]float64{"age": 9},
			ok:  false,
			err: ErrInvalidAge,
		},
		{
			in:  map[string]float64{"age": 21},
			ok:  false,
			err: ErrInvalidAge,
		},
		{
			in:  map[string]float64{"age": 15},
			ok:  true,
			err: nil,
		},
	}
	validator := ValidateAge(10, 20)
	for _, data := range cases {
		eq := NewEquation(data.in, conf).(*Equation)
		if ok, err := validator(eq); ok != data.ok {
			t.Error("Should receive a proper boolean")
		} else if err != data.err {
			t.Error("Should show proper error message")
		}
	}
}

func TestMeasureValidator(t *testing.T) {
	cases := []caseCommon{}
	validator := ValidateMeasures([]string{"age", "weight", "height"})
	for _, data := range cases {
		eq := NewEquation(data.in, conf).(*Equation)
		if ok, err := validator(eq); ok != data.ok {
			t.Error("Should receive proper boolean")
		} else if strings.Contains(err.Error(), data.err) {
			t.Error("Should show proper error message.")
		}
	}
}

type caseCommon struct {
	in  InParams
	ok  bool
	err string
}

var conf = NewEquationConf(
	"Testing",
	func(i interface{}) InParams {
		return map[string]float64{}
	},
	[]Validator{
		// ensure age and sskf measures are set
		func(e *Equation) (bool, error) {
			keys := []string{"age", "sskf"}
			for _, k := range keys {
				if _, ok := e.In(k); !ok {
					return false, ErrMissingMeasure
				}
			}
			return true, nil
		},
		// ensure age is between 18 and 64 years
		func(e *Equation) (bool, error) {
			v, ok := e.In("age")
			if !ok {
				return false, ErrMissingAge
			}

			lower, upper := 18.0, 64.0
			if v < lower || v >= upper {
				return false, ErrInvalidAge
			}

			return true, nil
		},
	},
	// equation that calculates body density and convert into body fat.
	func(e *Equation) float64 {
		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.09 - 0.0002*sskf + 0.0000001*math.Pow(sskf, 2) - 0.00005*age
		return 495.0/d - 450.0
	},
)
