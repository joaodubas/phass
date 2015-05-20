package bodyfat

import (
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	skf "github.com/joaodubas/phass/skinfold"
)

type BodyCompositionCalculator interface {
	PercentBodyFat() (float64, error)
	CanUse() (bool, error)
}

type SKFNewer interface {
	New(*assess.Person, *anthropo.Anthropometry, *skf.Skinfolds) BodyCompositionCalculator
}
