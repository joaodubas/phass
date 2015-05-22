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
	return fmt.Sprintf("%s\nBMI: %.2f kg/m^2 (%s)", b.Anthropometry.String(), b.Calc(), b.Classify())
}

func (b *BMI) Classify() string {
	return common.Classifier(b.Calc(), limitsForBMI, BMIClassification)
}

func (b *BMI) Calc() float64 {
	return b.Weight / math.Pow(b.Height/100, 2)
}

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
