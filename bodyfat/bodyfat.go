package bodyfat

import (
	"fmt"
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

/**
 * Skinfold equations
 */

var (
	NewWomenSevenSKF = FactoryBodyCompositionSKF(confWomenSevenSKF)
	NewWomenThreeSKF = FactoryBodyCompositionSKF(confWomenThreeSKF)
	NewWomenTwoSKF   = FactoryBodyCompositionSKF(confWomenTwoSKF)
	NewMenSevenSKF   = FactoryBodyCompositionSKF(confMenSevenSKF)
	NewMenThreeSKF   = FactoryBodyCompositionSKF(confMenThreeSKF)
	NewMenTwoSKF     = FactoryBodyCompositionSKF(confMenTwoSKF)
)

/**
 * SKF equation definition
 */

type BodyCompositionSKF struct {
	*assess.Person
	*assess.Assessment
	*skf.Skinfolds
	*common.EquationConf
}

func FactoryBodyCompositionSKF(conf SKFEquationConf) func(*assess.Person, *assess.Assessment, *skf.Skinfolds) *BodyCompositionSKF {
	c := NewEquationConfForSKF(conf)
	return func(p *assess.Person, a *assess.Assessment, s *skf.Skinfolds) *BodyCompositionSKF {
		return NewBodyCompositionSKF(p, a, s, c)
	}
}

func NewBodyCompositionSKF(p *assess.Person, a *assess.Assessment, s *skf.Skinfolds, e *common.EquationConf) *BodyCompositionSKF {
	return &BodyCompositionSKF{p, a, s, e}
}

func (b *BodyCompositionSKF) String() string {
	v, _ := b.Calc()
	return fmt.Sprintf("Body fat: %.2f %%", v)
}

func (b *BodyCompositionSKF) GetName() string {
	return "Body composition"
}

func (b *BodyCompositionSKF) Result() ([]string, error) {
	rs := []string{}

	v, err := b.Calc()
	if err != nil {
		return rs, err
	}

	rs = append(rs, fmt.Sprintf("Body fat: %.2f %%", v))
	return rs, nil
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

/**
 * Skfinfold conf definition
 */

var (
	confWomenSevenSKF = SKFEquationConf{
		name:     "Women seven skinfold equation from Jackson, Pollock, Ward",
		gender:   assess.Female,
		lowerAge: 18,
		upperAge: 55,
		skinfolds: []int{
			skf.SKFSubscapular,
			skf.SKFTriceps,
			skf.SKFChest,
			skf.SKFMidaxillary,
			skf.SKFSuprailiac,
			skf.SKFAbdominal,
			skf.SKFThigh,
		},
		equation: func(e *common.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.097 - 0.00046971*sskf + 0.00000056*math.Pow(sskf, 2) - 0.00012828*age
			return (5.01/d - 4.57) * 100.0
		},
	}
	confWomenThreeSKF = SKFEquationConf{
		name:     "Women three skinfold equation from Jackson, Pollock, Ward",
		gender:   assess.Female,
		lowerAge: 18,
		upperAge: 55,
		skinfolds: []int{
			skf.SKFTriceps,
			skf.SKFSuprailiac,
			skf.SKFThigh,
		},
		equation: func(e *common.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.0994921 - 0.0009929*sskf + 0.0000023*math.Pow(sskf, 2) - 0.0001392*age
			return (5.01/d - 4.57) * 100
		},
	}
	confWomenTwoSKF = SKFEquationConf{
		name:     "Women two skinfold equation from Slaughter et al.",
		gender:   assess.Female,
		lowerAge: 6,
		upperAge: 17,
		skinfolds: []int{
			skf.SKFTriceps,
			skf.SKFCalf,
		},
		equation: func(e *common.Equation) float64 {
			sskf, _ := e.In("sskf")
			return 0.735*sskf + 1.0
		},
	}
	confMenSevenSKF = SKFEquationConf{
		name:     "Men seven skinfold equation from Jackson, Pollock",
		gender:   assess.Male,
		lowerAge: 18,
		upperAge: 61,
		skinfolds: []int{
			skf.SKFSubscapular,
			skf.SKFTriceps,
			skf.SKFChest,
			skf.SKFMidaxillary,
			skf.SKFSuprailiac,
			skf.SKFAbdominal,
			skf.SKFThigh,
		},
		equation: func(e *common.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.112 - 0.00043499*sskf + 0.00000055*math.Pow(sskf, 2) - 0.0002882*age
			return (4.95/d - 4.5) * 100.0
		},
	}
	confMenThreeSKF = SKFEquationConf{
		name:     "Men three skinfold equation from Jackson, Pollock",
		gender:   assess.Male,
		lowerAge: 18,
		upperAge: 61,
		skinfolds: []int{
			skf.SKFChest,
			skf.SKFAbdominal,
			skf.SKFThigh,
		},
		equation: func(e *common.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.109380 - 0.0008267*sskf + 0.0000016*math.Pow(sskf, 2) - 0.0002574*age
			return (4.95/d - 4.5) * 100.0
		},
	}
	confMenTwoSKF = SKFEquationConf{
		name:     "Men two skinfold equation from Slaughter et al.",
		gender:   assess.Male,
		lowerAge: 6,
		upperAge: 17,
		skinfolds: []int{
			skf.SKFTriceps,
			skf.SKFCalf,
		},
		equation: func(e *common.Equation) float64 {
			sskf, _ := e.In("sskf")
			return 0.610*sskf + 5.1
		},
	}
)

/**
 * SKF equation conf
 */

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
