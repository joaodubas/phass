package bodyfat

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

type womenSevenSKF struct {
	*assess.Person
	*anthropo.Anthropometry
	*skf.Skinfolds
}

func (m womenSevenSKF) New(p *assess.Person, a *anthropo.Anthropometry, s *skf.Skinfolds) BodyCompositionCalculator {
	return &womenSevenSKF{p, a, s}
}

func (w *womenSevenSKF) PercentBodyFat() (float64, error) {
	db, err := w.density()
	if err != nil {
		return 0, err
	}
	return (5.01/db - 4.57) * 100.0, nil
}

func (w *womenSevenSKF) density() (float64, error) {
	if use, err := w.CanUse(); !use {
		return 0.0, err
	}

	age := w.Age()
	sskf := w.SumSpecific(womenSevenSSKF)
	return 1.097 - 0.00046971*sskf + 0.00000056*math.Pow(sskf, 2) - 0.00012828*age, nil
}

func (w *womenSevenSKF) CanUse() (bool, error) {
	if w.Gender != assess.Female {
		return false, fmt.Errorf("Equation appropriate for women.")
	}

	age := w.Age()
	if age < 18 || age > 55 {
		return false, fmt.Errorf("Equation appropriate for ages between 18 up to 55.")
	}

	return true, nil
}

var womenSevenSSKF = []int{
	skf.SKFChest,
	skf.SKFAbdominal,
	skf.SKFThigh,
	skf.SKFAbdominal,
	skf.SKFSuprailiac,
	skf.SKFTriceps,
	skf.SKFMidaxillary,
}
