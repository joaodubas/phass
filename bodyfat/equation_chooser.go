package bodyfat

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	skf "github.com/joaodubas/phass/skinfold"
)

func NewSumSKF(person *assess.Person, anthropometry *anthropo.Anthropometry, skinfolds *skf.Skinfolds) *SumSKF {
	return &SumSKF{person, anthropometry, skinfolds}
}

type SumSKF struct {
	*assess.Person
	*anthropo.Anthropometry
	*skf.Skinfolds
}

func (s *SumSKF) String() string {
	p, _ := s.PercentBodyFat()
	return fmt.Sprintf("%s\n%s\n%s\nBody fat: %.4f %%", s.Person, s.Anthropometry, s.Skinfolds, p)
}

func (s *SumSKF) PercentBodyFat() (float64, error) {
	bc := s.chooser()
	if use, err := bc.CanUse(); !use {
		return 0.0, err
	}
	return bc.PercentBodyFat()
}

func (s *SumSKF) CanUse() (bool, error) {
	return s.chooser().CanUse()
}

func (s *SumSKF) chooser() BodyCompositionCalculator {
	var bc BodyCompositionCalculator
	for _, fn := range skfEquations {
		tbc := fn(s.Person, s.Anthropometry, s.Skinfolds)
		if use, _ := tbc.CanUse(); !use {
			continue
		}
		bc = tbc
	}
	return bc
}

var skfEquations = []func(*assess.Person, *anthropo.Anthropometry, *skf.Skinfolds) BodyCompositionCalculator{}

func registerSKF(equation interface{}) func(*assess.Person, *anthropo.Anthropometry, *skf.Skinfolds) BodyCompositionCalculator {
	return func(p *assess.Person, a *anthropo.Anthropometry, s *skf.Skinfolds) BodyCompositionCalculator {
		eq := equation.(SKFNewer)
		return eq.New(p, a, s)
	}
}
