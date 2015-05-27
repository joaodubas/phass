package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

var (
	NewMenSevenSKF = FactoryBodyCompositionSKF(confMenSevenSKF)
	NewMenThreeSKF = FactoryBodyCompositionSKF(confMenThreeSKF)
	NewMenTwoSKF   = FactoryBodyCompositionSKF(confMenTwoSKF)
)

var confMenSevenSKF = SKFEquationConf{
	name:     "Men seven skinfold equation from Jackson, Pollock",
	gender:   assess.Male,
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
	equation: func(e *common.Equation) float64 {
		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.112 - 0.00043499*sskf + 0.00000055*math.Pow(sskf, 2) - 0.0002882*age
		return (4.95/d - 4.5) * 100.0
	},
}

var confMenThreeSKF = SKFEquationConf{
	name:     "Men three skinfold equation from Jackson, Pollock",
	gender:   assess.Male,
	lowerAge: 18,
	upperAge: 61,
	skinfolds: []int{
		skf.SKFChest,
		skf.SKFAbdominal,
		skf.SKFThigh,
	},
	equation: func(e *common.Equation) float64 {
		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.109380 - 0.0008267*sskf + 0.0000016*math.Pow(sskf, 2) - 0.0002574*age
		return (4.95/d - 4.5) * 100.0
	},
}

var confMenTwoSKF = SKFEquationConf{
	name:     "Men two skinfold equation from Slaughter et al.",
	gender:   assess.Male,
	lowerAge: 6,
	upperAge: 17,
	skinfolds: []int{
		skf.SKFTriceps,
		skf.SKFCalf,
	},
	equation: func(e *common.Equation) float64 {
		sskf, _ := e.In("sskf")
		return 0.610*sskf + 5.1
	},
}
