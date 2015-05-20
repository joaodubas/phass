package anthropometry

import (
	"fmt"
	"math"
)

type BMIPrime struct {
	*BMI
}

func (b *BMIPrime) String() string {
	return fmt.Sprintf("%s\nBMI prime: %.2f (%s)",  b.BMI.String(), b.Calc(), b.Classify())
}

func (b *BMIPrime) Classify() string {
	classes := map[int][]float64{
		VerySeverelyUnderweight: []float64{math.Inf(-1), 0.60},
		SeverelyUnderweight:     []float64{0.60, 0.64},
		Underweight:             []float64{0.64, 0.74},
		Normal:                  []float64{0.74, 1},
		Overweight:              []float64{1, 1.2},
		ObeseClassOne:           []float64{1.2, 1.4},
		ObeseClassTwo:           []float64{1.4, 1.6},
		ObeseClassThree:         []float64{1.6, math.Inf(+1)},
	}
	return classifier(b.Calc(), classes)
}

func (b *BMIPrime) Calc() float64 {
	return b.BMI.Calc() / 25.0
}
