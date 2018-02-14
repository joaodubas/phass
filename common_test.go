package phass

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestEquationValidation(t *testing.T) {
	cases := []caseCommon{
		{in: map[string]float64{}, ok: false, err: "Missing measure"},
		{in: map[string]float64{"age": 18.0}, ok: false, err: "Missing measure"},
		{in: map[string]float64{"sskf": 112.1}, ok: false, err: "Missing measure"},
		{in: map[string]float64{"age": 17.9, "sskf": 112.1}, ok: false, err: "Equation valid for age"},
	}

	for _, data := range cases {
		eq := NewEquation(data.in, conf)
		if v, err := eq.Validate(); v != data.ok {
			t.Error("Should get an error.")
		} else if !strings.Contains(err.Error(), data.err) {
			t.Error("Should get proper error message.")
		}

		if v, err := eq.Calc(); v > 0.0 {
			t.Error("Value should be 0")
		} else if !strings.Contains(err.Error(), data.err) {
			t.Error("Should get proper error message.")
		}
	}

	eq := NewEquation(map[string]float64{"age": 20, "sskf": 109.2}, conf)
	if v, err := eq.Validate(); !v {
		t.Error("Should be valid, instead get: %s", err)
	}
	if v, err := eq.Calc(); v <= 0.0 && err != nil {
		t.Error("Should get value greater that 0, instead get error %s", err)
	}

}

func TestEquationRetrieveIn(t *testing.T) {
	in := map[string]float64{
		"age":  18.8,
		"sskf": 210.1,
	}
	eq := NewEquation(in, conf)

	for k, v := range in {
		sv, ok := eq.In(k)
		if !ok {
			t.Errorf("Key %s should be defined.", k)
		}
		if v != sv {
			t.Errorf("Key %s should have value %.2f.", k, v)
		}
	}

	_, ok := eq.In("iDoNotExist")
	if ok {
		t.Error("Key should no be available")
	}
}

func TestAgeValidator(t *testing.T) {
	cases := []caseCommon{
		{
			in:  map[string]float64{},
			ok:  false,
			err: "Missing age",
		},
		{
			in:  map[string]float64{"age": 9},
			ok:  false,
			err: "Valid for ages",
		},
		{
			in:  map[string]float64{"age": 21},
			ok:  false,
			err: "Valid for ages",
		},
	}

	validator := ValidateAge(10, 20)
	for _, data := range cases {
		eq := NewEquation(data.in, conf).(*Equation)
		if ok, err := validator(eq); ok != data.ok {
			t.Error("Should receive a proper boolean")
		} else if !strings.Contains(err.Error(), data.err) {
			t.Error("Should show proper error message")
		}
	}

	eq := NewEquation(map[string]float64{"age": 15}, conf).(*Equation)
	if ok, err := validator(eq); !ok {
		t.Error("Should receive proper boolean")
	} else if err != nil {
		t.Error("Should not get an error")
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
		func(e *Equation) (bool, error) {
			keys := []string{"age", "sskf"}
			for _, k := range keys {
				if _, ok := e.In(k); !ok {
					return false, fmt.Errorf("Missing measure %s", k)
				}
			}
			return true, nil
		},
		func(e *Equation) (bool, error) {
			v, ok := e.In("age")
			if !ok {
				return false, fmt.Errorf("Missing measure age.")
			}

			lower, upper := 18.0, 64.0
			if v < lower || v >= upper {
				return false, fmt.Errorf("Equation valid for age %.0f up to %.0f", lower, upper)
			}

			return true, nil
		},
	},
	func(e *Equation) float64 {
		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.09 - 0.0002*sskf + 0.0000001*math.Pow(sskf, 2) - 0.00005*age
		return 495.0/d - 450.0
	},
)
