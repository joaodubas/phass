package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	skf "github.com/joaodubas/phass/skinfold"
	"strings"
	"testing"
)

func TestWomenSevenSKF(t *testing.T) {
	cases := []case_{
		newCase_(
			female,
			"1988-Mar-15",
			map[int]float64{},
			"Too young",
			0.0,
			"Valid for age",
		),
		newCase_(
			female,
			"2048-Mar-15",
			map[int]float64{},
			"Too old",
			0.0,
			"Valid for age",
		),
		newCase_(
			male,
			"1998-Dec-15",
			map[int]float64{},
			"Wrong gender",
			0.0,
			"Valid for gender",
		),
		newCase_(
			female,
			"2008-Mar-15",
			map[int]float64{},
			"Without skinfolds",
			0.0,
			"Missing skinfold",
		),
	}

	for _, data := range cases {
		bc := NewWomenSevenSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); data.err != "" && err == nil {
			t.Errorf("Case _%s_ failed, should show a validation error", data.name)
		} else if data.err != "" && !strings.Contains(err.Error(), data.err) {
			t.Logf("Error seen %s", err)
			t.Errorf("Case _%s_ failed, should show proper error message", data.name)
		} else if data.calc > 0.0 && calc != data.calc {
			t.Errorf("Case _%s_ failed, should have value %.2f, instead got %.2f")
		}
	}
}

func TestMenSevenSKF(t *testing.T) {

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
