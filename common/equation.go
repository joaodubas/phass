package common

func NewEquation(in InParams, conf *EquationConf) Equationer {
	return &Equation{in: in, conf: conf}
}

func NewEquationConf(name string, e Extractor, v []Validator, eq Calculator) *EquationConf {
	return &EquationConf{
		Name:       name,
		Extract:    e,
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

type(
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
