package common

import "fmt"

/**
 * Equation
 */

type Equation struct {
	in   map[string]float64
	conf *EquationConf
}

func NewEquation(in InParams, conf *EquationConf) Equationer {
	return &Equation{in: in, conf: conf}
}

func (e *Equation) String() string {
	return e.conf.Name
}

func (e *Equation) In(k string) (float64, bool) {
	v, ok := e.in[k]
	return v, ok
}

func (e *Equation) Validate() (bool, error) {
	for _, f := range e.conf.Validators {
		if r, err := f(e); err != nil {
			return r, err
		}
	}
	return true, nil
}

func (e *Equation) Calc() (float64, error) {
	if r, e := e.Validate(); !r {
		return 0.0, e
	}
	return e.conf.Calc(e), nil
}

type EquationConf struct {
	Name       string
	Extract    Extractor
	Validators []Validator
	Calc       Calculator
}

func NewEquationConf(name string, e Extractor, v []Validator, eq Calculator) *EquationConf {
	return &EquationConf{
		Name:       name,
		Extract:    e,
		Validators: v,
		Calc:       eq,
	}
}

type (
	InParams   map[string]float64
	Extractor  func(interface{}) InParams
	Validator  func(*Equation) (bool, error)
	Calculator func(*Equation) float64
)

type Equationer interface {
	In(string) (float64, bool)
	Validate() (bool, error)
	Calc() (float64, error)
}

/**
 * Validation
 */

func ValidateGender(expect int) Validator {
	return func(e *Equation) (bool, error) {
		if g, ok := e.In("gender"); !ok {
			return false, fmt.Errorf("Missing gender")
		} else if int(g) != expect {
			return false, fmt.Errorf("Valid for gender %d", expect)
		}
		return true, nil
	}
}

func ValidateAge(lower, upper float64) Validator {
	return func(e *Equation) (bool, error) {
		if age, ok := e.In("age"); !ok {
			return false, fmt.Errorf("Missing age measure")
		} else if age < lower || age > upper {
			return false, fmt.Errorf("Valid for ages between %.0f and %.0f", lower, upper)
		}
		return true, nil
	}
}

func ValidateMeasures(expect []string) Validator {
	return func(e *Equation) (bool, error) {
		for _, k := range expect {
			if _, ok := e.In(k); !ok {
				return false, fmt.Errorf("Missing %s measure", k)
			}
		}
		return true, nil
	}
}

/**
 * Classification
 */

func Classifier(value float64, classes map[int][2]float64, mapper map[int]string) string {
	cid := classifierId(value, classes)
	class, ok := mapper[cid]
	if !ok {
		return "No classification."
	}
	return class
}

func classifierId(value float64, classes map[int][2]float64) int {
	cid := -1
	for id_, limits := range classes {
		if value >= limits[0] && value < limits[1] {
			cid = id_
		}
	}
	return cid
}
