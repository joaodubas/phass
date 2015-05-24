package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
	"strings"
	"testing"
)

func TestBodyFatCompositionValidation(t *testing.T) {
	newEquation := FactoryBodyCompositionSKF(SKFEquationConf{
		name:     "Dummy Dubas Two SKF",
		gender:   assess.Female,
		lowerAge: 18,
		upperAge: 55,
		skinfolds: []int{
			skf.SKFSuprailiac,
			skf.SKFThigh,
		},
		equation: func(e *common.Equation) float64 {
			age, _ := e.In("age")
			sum, _ := e.In("sskf")
			d := 1.01 - 0.0001*sum + 0.0000004*sum*sum - 0.000001*age
			return 495/d - 450
		},
	})
	cases := []case_{
		newCase_(
			female,
			"1988-Mar-15",
			map[int]float64{skf.SKFTriceps: 10.5, skf.SKFSuprailiac: 42.1, skf.SKFThigh: 21.2},
			"Too young",
			0.0,
			"Valid for age",
		),
		newCase_(
			female,
			"2048-Mar-15",
			map[int]float64{skf.SKFTriceps: 10.5, skf.SKFSuprailiac: 42.1, skf.SKFThigh: 21.2},
			"Too old",
			0.0,
			"Valid for age",
		),
		newCase_(
			male,
			"1998-Dec-15",
			map[int]float64{skf.SKFTriceps: 10.5, skf.SKFSuprailiac: 42.1, skf.SKFThigh: 21.2},
			"Wrong gender",
			0.0,
			"Valid for gender",
		),
		newCase_(
			female,
			"2008-Mar-15",
			map[int]float64{skf.SKFTriceps: 10.5},
			"Without skinfolds",
			0.0,
			"Missing skinfold",
		),
		newCase_(
			female,
			"2008-Mar-15",
			map[int]float64{skf.SKFTriceps: 10.5, skf.SKFSuprailiac: 42.1, skf.SKFThigh: 21.2},
			"Every little thing is gone be allright",
			42.41,
			"",
		),
	}

	for _, data := range cases {
		bc := newEquation(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); data.err != "" && err == nil {
			t.Errorf("Case _%s_ failed, should show a validation error", data.name)
		} else if data.err != "" && !strings.Contains(err.Error(), data.err) {
			t.Errorf("Case _%s_ failed, should show proper error message", data.name)
		} else if data.calc > 0.0 && !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

type case_ struct {
	person     *assess.Person
	assessment *assess.Assessment
	skinfold   *skf.Skinfolds
	name       string
	calc       float64
	err        string
}

func newCase_(p *assess.Person, d string, s map[int]float64, name string, calc float64, err string) case_ {
	assessment, _ := assess.NewAssessment(d)
	return case_{
		person:     p,
		assessment: assessment,
		skinfold:   skf.NewSkinfolds(s),
		name:       name,
		calc:       calc,
		err:        err,
	}
}

var (
	male, _   = assess.NewPerson("Joao Paulo Dubas", "1978-Dec-15", assess.Male)
	female, _ = assess.NewPerson("Ana Paula Dubas", "1988-Mar-15", assess.Female)
)

func floatEqual(original, expected, limit float64) bool {
	diff := math.Abs(original-expected)
	return diff <= limit
}
