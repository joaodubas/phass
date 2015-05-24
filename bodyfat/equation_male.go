package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

var NewMenSevenSKF = FactoryBodyCompositionSKF(confMenSevenSKF)

var confMenSevenSKF = SKFEquationConf{
	name:     "Seven skinfold equation from Pollock",
	gender:   assess.Male,
	lowerAge: 18,
	upperAge: 61,
	skinfolds: []int{
		skf.SKFChest,
		skf.SKFAbdominal,
		skf.SKFThigh,
		skf.SKFAbdominal,
		skf.SKFSuprailiac,
		skf.SKFTriceps,
		skf.SKFMidaxillary,
	},
	equation: func(e *common.Equation) float64 {
		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.112 - 0.00043499*sskf + 0.00000055*math.Pow(sskf, 2) - 0.0002882*age

		return (4.95/d - 4.5) * 100.0
	},
}
