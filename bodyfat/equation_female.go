package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

var (
	NewWomenSevenSKF = FactoryBodyCompositionSKF(confWomenSevenSKF)
	NewWomenThreeSKF = FactoryBodyCompositionSKF(confWomenThreeSKF)
	NewWomenTwoSKF   = FactoryBodyCompositionSKF(confWomenTwoSKF)
)

var confWomenSevenSKF = SKFEquationConf{
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

var confWomenThreeSKF = SKFEquationConf{
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

var confWomenTwoSKF = SKFEquationConf{
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
