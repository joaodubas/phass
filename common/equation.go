package common

func NewEquation(in InParams, conf *EquationConf) Equationer {
	return &Equation{in: in, conf: conf}
}

func NewEquationConf(name string, v []Validate, eq Calc) *EquationConf {
	return &EquationConf{
		Name:       name,
		Validators: v,
		Calc:       eq,
	}
}

type Equation struct {
	in   map[string]float64
	conf *EquationConf
}

func (e *Equation) String() string {
	return e.conf.Name
}

func (e *Equation) In(k string) (float64, bool) {
	v, ok := e.in[k]
	return v, ok
}

func (e *Equation) Validator() (bool, error) {
	for _, f := range e.conf.Validators {
		if r, err := f(e); err != nil {
			return r, err
		}
	}
	return true, nil
}

func (e *Equation) Calc() (float64, error) {
	if r, e := e.Validator(); !r {
		return 0.0, e
	}
	return e.conf.Calc(e), nil
}

type EquationConf struct {
	Name       string
	Validators []Validate
	Calc       Calc
}

type InParams map[string]float64

type Validate func(*Equation) (bool, error)

type Calc func(*Equation) float64

type Equationer interface {
	In(string) (float64, bool)
	Validator() (bool, error)
	Calc() (float64, error)
}
