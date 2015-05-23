package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

var womenSevenSSKF = []int{
	skf.SKFChest,
	skf.SKFAbdominal,
	skf.SKFThigh,
	skf.SKFAbdominal,
	skf.SKFSuprailiac,
	skf.SKFTriceps,
	skf.SKFMidaxillary,
}

var confWomenSevenSKF = common.NewEquationConf(
	"Seven skinfold equation from Pollock",
	[]common.Validator{
		common.ValidateMeasures([]string{"gender", "age", "sskf"}),
		common.ValidateGender(assess.Female),
		common.ValidateAge(18, 55),
	},
	func(e *common.Equation) float64 {
		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.097 - 0.00046971*sskf + 0.00000056*math.Pow(sskf, 2) - 0.00012828*age

		return (5.01/d - 4.57) * 100.0
	},
)
