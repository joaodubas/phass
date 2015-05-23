package common

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestEquationValidation(t *testing.T) {
	cases := []case_{
		case_{in: map[string]float64{}, rsVal: false, errVal: "Missing measure"},
		case_{in: map[string]float64{"age": 18.0}, rsVal: false, errVal: "Missing measure"},
		case_{in: map[string]float64{"sskf": 112.1}, rsVal: false, errVal: "Missing measure"},
		case_{in: map[string]float64{"age": 17.9, "sskf": 112.1}, rsVal: false, errVal: "Equation valid for age"},
	}

	for _, data := range cases {
		eq := NewEquation(data.in, conf)
		if v, err := eq.Validator(); v != data.rsVal {
			t.Error("Should get an error.")
		} else if !strings.Contains(err.Error(), data.errVal) {
			t.Error("Should get proper error message.")
		}

		if v, err := eq.Calc(); v > 0.0 {
			t.Error("Value should be 0")
		} else if !strings.Contains(err.Error(), data.errVal) {
			t.Error("Should get proper error message.")
		}
	}

	eq := NewEquation(map[string]float64{"age": 20, "sskf": 109.2}, conf)
	if v, err := eq.Validator(); !v {
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

type case_ struct {
	in     map[string]float64
	rsVal  bool
	errVal string
}

var conf = NewEquationConf(
	"Testing",
	[]Validate{
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
