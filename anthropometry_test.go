package phass

import (
	"testing"
)

func TestBMI(t *testing.T) {
	cases := []caseAnthropometry{
		{
			bmi:      anthropo{height: 189.2, weight: 51.1},
			calc:     14.2751,
			classify: BMIClassification[VerySeverelyUnderweight],
		},
		{
			bmi:      anthropo{height: 193.6, weight: 58.9},
			calc:     15.7146,
			classify: BMIClassification[SeverelyUnderweight],
		},
		{
			bmi:      anthropo{height: 171.5, weight: 54.2},
			calc:     18.4277,
			classify: BMIClassification[Underweight],
		},
		{
			bmi:      anthropo{height: 172.6, weight: 71.3},
			calc:     23.9336,
			classify: BMIClassification[Normal],
		},
		{
			bmi:      anthropo{height: 173.5, weight: 88.3},
			calc:     29.3334,
			classify: BMIClassification[Overweight],
		},
		{
			bmi:      anthropo{height: 164.3, weight: 83.8},
			calc:     31.0434,
			classify: BMIClassification[ObeseClassOne],
		},
		{
			bmi:      anthropo{height: 171.1, weight: 106.9},
			calc:     36.5155,
			classify: BMIClassification[ObeseClassTwo],
		},
		{
			bmi:      anthropo{height: 168.1, weight: 118.1},
			calc:     41.7941,
			classify: BMIClassification[ObeseClassThree],
		},
	}

	for _, data := range cases {
		bmi := NewBMI(data.bmi.weight, data.bmi.height)
		if calc, _ := bmi.Calc(); !floatEqual(calc, data.calc, FloatLimit) {
			t.Errorf("BMI calculated is %.4f and expected is %.4f\n", calc, data.calc)
		}
		if classify, _ := bmi.Classify(); classify != data.classify {
			t.Errorf("Classification defined is %s and expected is %s\n", classify, data.classify)
		}
	}
}

func TestBMIPrime(t *testing.T) {
	cases := []caseAnthropometry{
		{
			bmi:      anthropo{height: 189.2, weight: 51.1},
			calc:     0.571,
			classify: BMIClassification[VerySeverelyUnderweight],
		},
		{
			bmi:      anthropo{height: 193.6, weight: 58.9},
			calc:     0.6286,
			classify: BMIClassification[SeverelyUnderweight],
		},
		{
			bmi:      anthropo{height: 171.5, weight: 54.2},
			calc:     0.7371,
			classify: BMIClassification[Underweight],
		},
		{
			bmi:      anthropo{height: 172.6, weight: 71.3},
			calc:     0.9573,
			classify: BMIClassification[Normal],
		},
		{
			bmi:      anthropo{height: 173.5, weight: 88.3},
			calc:     1.1733,
			classify: BMIClassification[Overweight],
		},
		{
			bmi:      anthropo{height: 164.3, weight: 83.8},
			calc:     1.2417,
			classify: BMIClassification[ObeseClassOne],
		},
		{
			bmi:      anthropo{height: 171.1, weight: 106.9},
			calc:     1.4606,
			classify: BMIClassification[ObeseClassTwo],
		},
		{
			bmi:      anthropo{height: 168.1, weight: 118.1},
			calc:     1.6718,
			classify: BMIClassification[ObeseClassThree],
		},
	}

	for _, data := range cases {
		bmi := NewBMIPrime(data.bmi.weight, data.bmi.height)
		if calc, _ := bmi.Calc(); !floatEqual(calc, data.calc, FloatLimit) {
			t.Errorf("BMI calculated is %.4f and expected is %.4f\n", calc, data.calc)
		}
		if classify, _ := bmi.Classify(); classify != data.classify {
			t.Errorf("Classification defined is %s and expected is %s\n", classify, data.classify)
		}
	}
}

type anthropo struct {
	weight float64
	height float64
}

type caseAnthropometry struct {
	bmi      anthropo
	calc     float64
	classify string
}

var FloatLimit = 0.0001
