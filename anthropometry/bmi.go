package anthropometry

import (
	"fmt"
	"math"
)

type BMI struct {
	*Anthropometry
}

func (b *BMI) String() string {
	return fmt.Sprintf("%s\nBMI: %.2f kg/m^2 (%s)", b.Anthropometry.String(), b.Calc(), b.Classify())
}

func (b *BMI) Classify() string {
	classes := map[int][]float64{
		VerySeverelyUnderweight: []float64{math.Inf(-1), 15},
		SeverelyUnderweight:     []float64{15, 16},
		Underweight:             []float64{16, 18.5},
		Normal:                  []float64{18.5, 25},
		Overweight:              []float64{25, 30},
		ObeseClassOne:           []float64{30, 35},
		ObeseClassTwo:           []float64{35, 40},
		ObeseClassThree:         []float64{40, math.Inf(+1)},
	}
	return classifier(b.Calc(), classes)
}

func (b *BMI) Calc() float64 {
	return b.Weight / math.Pow(b.Height/100, 2)
}
