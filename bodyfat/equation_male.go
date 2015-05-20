package bodyfat

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

func init() {
	skfEquations = append(skfEquations, registerSKF(menSevenSKF{}))
}

type menSevenSKF struct {
	*assess.Person
	*anthropo.Anthropometry
	*skf.Skinfolds
}

func (m menSevenSKF) New(p *assess.Person, a *anthropo.Anthropometry, s *skf.Skinfolds) BodyCompositionCalculator {
	return &menSevenSKF{p, a, s}
}

func (m *menSevenSKF) PercentBodyFat() (float64, error) {
	db, err := m.density()
	if err != nil {
		return 0.0, err
	}
	return (4.95/db - 4.5) * 100.0, nil
}

func (m *menSevenSKF) density() (float64, error) {
	if use, err := m.CanUse(); !use {
		return 0.0, err
	}

	age := m.Age()
	sskf := m.SumSpecific(menSevenSSKF)
	return 1.112 - 0.00043499*sskf + 0.00000055*math.Pow(sskf, 2) - 0.0002882*age, nil
}

func (m *menSevenSKF) CanUse() (bool, error) {
	if m.Gender != assess.Male {
		return false, fmt.Errorf("Equation appropriate for mem.")
	}

	age := m.Age()
	if age < 18 || age > 61 {
		return false, fmt.Errorf("Equation appropriate for ages between 18 up to 61.")
	}

	return true, nil
}

var menSevenSSKF = []int{
	skf.SKFChest,
	skf.SKFAbdominal,
	skf.SKFThigh,
	skf.SKFAbdominal,
	skf.SKFSuprailiac,
	skf.SKFTriceps,
	skf.SKFMidaxillary,
}
