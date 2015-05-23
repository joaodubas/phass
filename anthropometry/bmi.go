package anthropometry

import (
	"fmt"
	"github.com/joaodubas/phass/common"
	"math"
)

type BMI struct {
	*Anthropometry
}

func (b *BMI) String() string {
	v, _ := b.Calc()
	c, _ := b.Classify()
	return fmt.Sprintf("%s\nBMI: %.2f kg/m^2 (%s)", b.Anthropometry.String(), v, c)
}

func (b *BMI) Result() ([]string, error) {
	return []string{}, nil
}

func (b *BMI) Classify() (string, error) {
	v, err := b.Calc()
	if err != nil {
		return "", err
	}
	return common.Classifier(v, limitsForBMI, BMIClassification), nil
}

func (b *BMI) Calc() (float64, error) {
	return b.equation().Calc()
}

func (b *BMI) equation() common.Equationer {
	return common.NewEquation(bmiConf.Extract(b), bmiConf)
}

var bmiConf = common.NewEquationConf(
	"BMI",
	func(i interface{}) common.InParams {
		c := i.(*BMI)
		return map[string]float64{
			"weight": c.Weight,
			"height": c.Height,
		}
	},
	[]common.Validator{
		common.ValidateMeasures([]string{"weight", "height"}),
	},
	func(e *common.Equation) float64 {
		w, _ := e.In("weight")
		h, _ := e.In("height")
		return w / math.Pow(h/100, 2)
	},
)

var limitsForBMI = map[int][2]float64{
	VerySeverelyUnderweight: [2]float64{math.Inf(-1), 15},
	SeverelyUnderweight:     [2]float64{15, 16},
	Underweight:             [2]float64{16, 18.5},
	Normal:                  [2]float64{18.5, 25},
	Overweight:              [2]float64{25, 30},
	ObeseClassOne:           [2]float64{30, 35},
	ObeseClassTwo:           [2]float64{35, 40},
	ObeseClassThree:         [2]float64{40, math.Inf(+1)},
}
