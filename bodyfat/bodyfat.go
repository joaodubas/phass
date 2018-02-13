package bodyfat

import (
	"fmt"
	"math"

	"github.com/joaodubas/phass"
	skf "github.com/joaodubas/phass/skinfold"
)

/**
 * Skinfold equations
 */

// Skinfold percentage of fat estimation for differents genders.
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

// BodyCompositionSKF contains data needed to estimate body composition of a
// given person, based in skilfold assessment. This is composition of different
// structs, such as: a person, assessment details, skinfolds and an equation.
type BodyCompositionSKF struct {
	*phass.Person
	*phass.Assessment
	*skf.Skinfolds
	*phass.EquationConf
}

// FactoryBodyCompositionSKF factory to create new body composition assesment by
// skilfolds methods. It returns a function to create new BodyCompositionSKF
// structs.
func FactoryBodyCompositionSKF(conf SKFEquationConf) func(*phass.Person, *phass.Assessment, *skf.Skinfolds) *BodyCompositionSKF {
	c := NewEquationConfForSKF(conf)
	return func(p *phass.Person, a *phass.Assessment, s *skf.Skinfolds) *BodyCompositionSKF {
		return NewBodyCompositionSKF(p, a, s, c)
	}
}

// NewBodyCompositionSKF create a new body composition assessment. It receives
// person, an assessment, skinfolds, and the equation to estimate body fat
// percentange. Returns a pointer to BodyCompostionSKF.
func NewBodyCompositionSKF(p *phass.Person, a *phass.Assessment, s *skf.Skinfolds, e *phass.EquationConf) *BodyCompositionSKF {
	return &BodyCompositionSKF{p, a, s, e}
}

func (b *BodyCompositionSKF) String() string {
	v, _ := b.Calc()
	return fmt.Sprintf("Body fat: %.2f %%", v)
}

// GetName returns this measurement name.
func (b *BodyCompositionSKF) GetName() string {
	return "Body composition"
}

// Result returns information about body composition assessment.
func (b *BodyCompositionSKF) Result() ([]string, error) {
	rs := []string{}

	v, err := b.Calc()
	if err != nil {
		return rs, err
	}

	rs = append(rs, fmt.Sprintf("Body fat: %.2f %%", v))
	return rs, nil
}

// Classify returns classification related to body fat percentage.
func (b *BodyCompositionSKF) Classify() (string, error) {
	return "", nil
}

// Calc returns value for estimate body fat percentage.
func (b *BodyCompositionSKF) Calc() (float64, error) {
	return b.equation().Calc()
}

// equation returns a equation, used to estimate body fat percentage.
func (b *BodyCompositionSKF) equation() phass.Equationer {
	return phass.NewEquation(b.EquationConf.Extract(b), b.EquationConf)
}

/**
 * Skfinfold conf definition
 */

// Popular skinfold equations to estimate body fat.
var (
	confWomenSevenSKF = SKFEquationConf{
		name:     "Women seven skinfold equation from Jackson, Pollock, Ward",
		gender:   phass.Female,
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
		equation: func(e *phass.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.097 - 0.00046971*sskf + 0.00000056*math.Pow(sskf, 2) - 0.00012828*age
			return (5.01/d - 4.57) * 100.0
		},
	}
	confWomenThreeSKF = SKFEquationConf{
		name:     "Women three skinfold equation from Jackson, Pollock, Ward",
		gender:   phass.Female,
		lowerAge: 18,
		upperAge: 55,
		skinfolds: []int{
			skf.SKFTriceps,
			skf.SKFSuprailiac,
			skf.SKFThigh,
		},
		equation: func(e *phass.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.0994921 - 0.0009929*sskf + 0.0000023*math.Pow(sskf, 2) - 0.0001392*age
			return (5.01/d - 4.57) * 100
		},
	}
	confWomenTwoSKF = SKFEquationConf{
		name:     "Women two skinfold equation from Slaughter et al.",
		gender:   phass.Female,
		lowerAge: 6,
		upperAge: 17,
		skinfolds: []int{
			skf.SKFTriceps,
			skf.SKFCalf,
		},
		equation: func(e *phass.Equation) float64 {
			sskf, _ := e.In("sskf")
			return 0.735*sskf + 1.0
		},
	}
	confMenSevenSKF = SKFEquationConf{
		name:     "Men seven skinfold equation from Jackson, Pollock",
		gender:   phass.Male,
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
		equation: func(e *phass.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.112 - 0.00043499*sskf + 0.00000055*math.Pow(sskf, 2) - 0.0002882*age
			return (4.95/d - 4.5) * 100.0
		},
	}
	confMenThreeSKF = SKFEquationConf{
		name:     "Men three skinfold equation from Jackson, Pollock",
		gender:   phass.Male,
		lowerAge: 18,
		upperAge: 61,
		skinfolds: []int{
			skf.SKFChest,
			skf.SKFAbdominal,
			skf.SKFThigh,
		},
		equation: func(e *phass.Equation) float64 {
			age, _ := e.In("age")
			sskf, _ := e.In("sskf")
			d := 1.109380 - 0.0008267*sskf + 0.0000016*math.Pow(sskf, 2) - 0.0002574*age
			return (4.95/d - 4.5) * 100.0
		},
	}
	confMenTwoSKF = SKFEquationConf{
		name:     "Men two skinfold equation from Slaughter et al.",
		gender:   phass.Male,
		lowerAge: 6,
		upperAge: 17,
		skinfolds: []int{
			skf.SKFTriceps,
			skf.SKFCalf,
		},
		equation: func(e *phass.Equation) float64 {
			sskf, _ := e.In("sskf")
			return 0.610*sskf + 5.1
		},
	}
)

/**
 * SKF equation conf
 */

// NewEquationConfForSKF returns an equation configuration based in provided
// configuration.
func NewEquationConfForSKF(conf SKFEquationConf) *phass.EquationConf {
	extractor := func(i interface{}) phass.InParams {
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
	validators := []phass.Validator{
		phass.ValidateMeasures([]string{"gender", "age", "sskf"}),
		phass.ValidateGender(conf.gender),
		phass.ValidateAge(conf.lowerAge, conf.upperAge),
		validateSkinfolds(conf.skinfolds),
	}
	return phass.NewEquationConf(conf.name, extractor, validators, conf.equation)
}

// SKFEquationConf common configuration for skinfold equations.
type SKFEquationConf struct {
	name      string
	gender    int
	lowerAge  float64
	upperAge  float64
	skinfolds []int
	equation  phass.Calculator
}

func validateSkinfolds(skfs []int) phass.Validator {
	return func(e *phass.Equation) (bool, error) {
		for _, k := range skfs {
			if _, ok := e.In(skf.NamedSkinfold(k)); !ok {
				return false, fmt.Errorf("Missing skinfold %s", skf.NamedSkinfold(k))
			}
		}
		return true, nil
	}
}
