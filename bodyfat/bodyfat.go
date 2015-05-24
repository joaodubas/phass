package bodyfat

import (
	"fmt"
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
)

type bodyCompBySKF func(*assess.Person, *assess.Assessment, *skf.Skinfolds) *BodyCompositionSKF

func FactoryBodyCompositionSKF(conf SKFEquationConf) bodyCompBySKF {
	c := NewEquationConfForSKF(conf)
	return func(p *assess.Person, a *assess.Assessment, s *skf.Skinfolds) *BodyCompositionSKF {
		return NewBodyCompositionSKF(p, a, s, c)
	}
}

type BodyCompositionSKF struct {
	*assess.Person
	*assess.Assessment
	*skf.Skinfolds
	*common.EquationConf
}

func NewBodyCompositionSKF(p *assess.Person, a *assess.Assessment, s *skf.Skinfolds, e *common.EquationConf) *BodyCompositionSKF {
	return &BodyCompositionSKF{p, a, s, e}
}

func (b *BodyCompositionSKF) String() string {
	return ""
}

func (b *BodyCompositionSKF) Result() ([]string, error) {
	return []string{}, nil
}

func (b *BodyCompositionSKF) Classify() (string, error) {
	return "", nil
}

func (b *BodyCompositionSKF) Calc() (float64, error) {
	return b.equation().Calc()
}

func (b *BodyCompositionSKF) equation() common.Equationer {
	return common.NewEquation(b.EquationConf.Extract(b), b.EquationConf)
}

func NewEquationConfForSKF(conf SKFEquationConf) *common.EquationConf {
	extractor := func(i interface{}) common.InParams {
		c := i.(*BodyCompositionSKF)
		r := map[string]float64{
			"gender": float64(c.Gender),
			"age":    c.AgeFromDate(c.Date),
			"sskf":   c.SumSpecific(conf.skinfolds),
		}
		for _, s := range conf.skinfolds {
			if v, ok := c.Skinfolds.Measures[s]; ok {
				r[skf.NamedSkinfold(s)] = v
			}
		}
		return r
	}
	validators := []common.Validator{
		common.ValidateMeasures([]string{"gender", "age", "sskf"}),
		common.ValidateGender(conf.gender),
		common.ValidateAge(conf.lowerAge, conf.upperAge),
		validateSkinfolds(conf.skinfolds),
	}
	return common.NewEquationConf(conf.name, extractor, validators, conf.equation)
}

type SKFEquationConf struct {
	name      string
	gender    int
	lowerAge  float64
	upperAge  float64
	skinfolds []int
	equation  common.Calculator
}

func validateSkinfolds(skfs []int) common.Validator {
	return func(e *common.Equation) (bool, error) {
		for _, k := range skfs {
			if _, ok := e.In(skf.NamedSkinfold(k)); !ok {
				return false, fmt.Errorf("Missing skinfold %s", skf.NamedSkinfold(k))
			}
		}
		return true, nil
	}
}
