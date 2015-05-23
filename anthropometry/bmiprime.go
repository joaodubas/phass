package anthropometry

import (
	"fmt"
	"github.com/joaodubas/phass/common"
	"math"
)

type BMIPrime struct {
	*BMI
}

func (b *BMIPrime) String() string {
	v, _ := b.Calc()
	c, _ := b.Classify()
	return fmt.Sprintf("%s\nBMI prime: %.2f (%s)", b.BMI.String(), v, c)
}

func (b *BMIPrime) Result() ([]string, error) {
	return []string{}, nil
}

func (b *BMIPrime) Classify() (string, error) {
	v, err := b.Calc()
	if err != nil {
		return "", err
	}
	return common.Classifier(v, limitsForBMIPrime, BMIClassification), nil
}

func (b *BMIPrime) Calc() (float64, error) {
	v, err := b.BMI.Calc()
	if err != nil {
		return 0.0, err
	}
	return v / 25.0, nil
}

var limitsForBMIPrime = map[int][2]float64{
	VerySeverelyUnderweight: [2]float64{math.Inf(-1), 0.60},
	SeverelyUnderweight:     [2]float64{0.60, 0.64},
	Underweight:             [2]float64{0.64, 0.74},
	Normal:                  [2]float64{0.74, 1},
	Overweight:              [2]float64{1, 1.2},
	ObeseClassOne:           [2]float64{1.2, 1.4},
	ObeseClassTwo:           [2]float64{1.4, 1.6},
	ObeseClassThree:         [2]float64{1.6, math.Inf(+1)},
}
