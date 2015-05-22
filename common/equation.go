package common

func NewEquation(name string, in map[string]float64, v []func(e *Equation) (bool, error), eq func(e *Equation) float64) *Equation {
	return &Equation{Name: name, in: in, validators: v, equation: eq}
}

type Equation struct {
	Name       string
	in         map[string]float64
	validators []func(e *Equation) (bool, error)
	equation   func(e *Equation) float64
}

func (e *Equation) String() string {
	return e.Name
}

func (e *Equation) In(k string) (float64, bool) {
	v, ok := e.in[k]
	return v, ok
}

func (e *Equation) Validator() (bool, error) {
	for _, f := range e.validators {
		if r, err := f(e); err != nil {
			return r, err
		}
	}
	return true, nil
}

func (e *Equation) Equation() (float64, error) {
	if r, e := e.Validator(); !r {
		return 0.0, e
	}
	return e.equation(e), nil
}

type Equationer interface {
	In(string) (float64, bool)
	Validator() (bool, error)
	Equation() (float64, error)
}
