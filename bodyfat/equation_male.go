package bodyfat

import (
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	skf "github.com/joaodubas/phass/skinfold"
	"math"
)

var menSevenSSKF = []int{
	skf.SKFChest,
	skf.SKFAbdominal,
	skf.SKFThigh,
	skf.SKFAbdominal,
	skf.SKFSuprailiac,
	skf.SKFTriceps,
	skf.SKFMidaxillary,
}

var confMenSevenSKF = common.NewEquationConf(
	"Seven skinfold equation from Pollock",
	[]func(*common.Equation) (bool, error){
		common.ValidateMeasures([]string{"gender", "age", "sskf"}),
		common.ValidateGender(assess.Male),
		common.ValidateAge(18, 61),
	},
	func(e *common.Equation) (float64, error) {
		if ok, err := e.Validator(); !ok {
			return 0.0, err
		}

		age, _ := e.In("age")
		sskf, _ := e.In("sskf")
		d := 1.112 - 0.00043499*sskf + 0.00000055*math.Pow(sskf, 2) - 0.0002882*age

		return (4.95/d - 4.5) * 100.0, nil
	},
)
