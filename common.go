package phass

import "fmt"

/**
 * Equation
 */

// Equation represents the configuration for an equation used in a given
// assessment. This is represented by input parameters used in equation and a
// configuration.
type Equation struct {
	in   map[string]float64
	conf *EquationConf
}

// NewEquation returnas a Equationer interface. It receives input parameters
// and configuration for this equation.
func NewEquation(in InParams, conf *EquationConf) Equationer {
	return &Equation{in: in, conf: conf}
}

func (e *Equation) String() string {
	return e.conf.Name
}

// In verifies if a given input parameter was provided and it's value.
func (e *Equation) In(k string) (float64, bool) {
	v, ok := e.in[k]
	return v, ok
}

// Validate execute provided validators and returns boolean indicating if it's
// valid or not, and any error associated.
func (e *Equation) Validate() (bool, error) {
	for _, f := range e.conf.Validators {
		if r, err := f(e); err != nil {
			return r, err
		}
	}
	return true, nil
}

// Calc returns this equation value, and errors if the equation can't be
// calculated.
func (e *Equation) Calc() (float64, error) {
	if r, e := e.Validate(); !r {
		return 0.0, e
	}
	return e.conf.Calc(e), nil
}

// EquationConf represents a configuration for a given equation. It's
// represented by a name, a function that extract input parameters, a slice of
// validators methods, and a calculatator function that represents an equation.
type EquationConf struct {
	Name       string
	Extract    Extractor
	Validators []Validator
	Calc       Calculator
}

// NewEquationConf returns an EquationConf pointer, that receives a name, an
// extractor function, a slice of validators, and a calcutator function.
func NewEquationConf(name string, e Extractor, v []Validator, eq Calculator) *EquationConf {
	return &EquationConf{
		Name:       name,
		Extract:    e,
		Validators: v,
		Calc:       eq,
	}
}

type (
	// InParams is a map of string to float values, represent input params for
	// an equation.
	InParams map[string]float64
	// Extractor is a function to extract input parameters, receives an
	// interface and returns input params for an equation.
	Extractor func(interface{}) InParams
	// Validator is a function to validate input parameters provided in an
	// equation, receives an equation pointer and returns a boolean and error.
	Validator func(*Equation) (bool, error)
	// Calculator is a function calculate the value based in a equation.
	Calculator func(*Equation) float64
)

// Equationer is an interface that wraps an equation.
// In function is used to verify a given input parameter.
// Validate function is used to ensure input parameters are valid.
// Calc function is used to return this equation value.
type Equationer interface {
	In(string) (float64, bool)
	Validate() (bool, error)
	Calc() (float64, error)
}

/**
 * Validation
 */

// ValidateGender returns a Validator function, that ensure gender is equal to
// the one expected.
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

// ValidateAge returns a Validator function, that ensure a given age is
// between lower and upper limits.
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

// ValidateMeasures returns a Validator function, that ensure a list of
// expected measures are available.
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

// Classifier is used to classifiy a given value, with base in classes and a
// string mapper.
// classes is used to verify in which classification bin this value is
// contained, and mapper is used to convert the classification bin to a string.
func Classifier(value float64, classes map[int][2]float64, mapper map[int]string) string {
	cid := classifierIndex(value, classes)
	class, ok := mapper[cid]
	if !ok {
		return "No classification."
	}
	return class
}

// classifierIndex returns the classification bin index containing this value.
// Returns a positive integer representing the index where this value is
// classified, or -1 when no classification bin contains the provided value.
func classifierIndex(value float64, classes map[int][2]float64) int {
	cid := -1
	for index, limits := range classes {
		if value >= limits[0] && value < limits[1] {
			cid = index
		}
	}
	return cid
}
