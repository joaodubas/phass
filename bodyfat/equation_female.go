package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

var NewWomenSevenSKF = FactoryBodyCompositionSKF(confWomenSevenSKF)

var confWomenSevenSKF = SKFEquationConf{
	name:     "Seven skinfold equation from Pollock",
	gender:   assess.Female,
	lowerAge: 18,
	upperAge: 55,
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
		d := 1.097 - 0.00046971*sskf + 0.00000056*math.Pow(sskf, 2) - 0.00012828*age

		return (5.01/d - 4.57) * 100.0
	},
}
