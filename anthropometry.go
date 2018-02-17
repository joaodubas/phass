package phass

import (
	"fmt"
	"math"
	"strings"
)

/**
 * BMI Prime
 */

// NewBMIPrime creates a new anthropometric index able to calculate and
// classify a BMI Prime based on weight and height.
var NewBMIPrime = newAnthropoRatio(
	limitsForBMIPrime,
	func(w, h float64) Measurer {
		return NewBMI(w, h)
	},
	bmiPrimeConf,
	func(a Measurer) []string {
		i := a.(*AnthropometricRatio)
		v, _ := i.Calc()
		c, _ := i.Classify()
		return []string{
			fmt.Sprintf("BMI Prime: %.4f", v),
			fmt.Sprintf("BMI Prime classification: %s", c),
		}
	},
)

/**
 * BMI
 */

// NewBMI creates a new anthropometric index able to calculate and classify
// a BMI based on weight and height.
var NewBMI = newAnthropoRatio(
	limitsForBMI,
	func(w, h float64) Measurer {
		return NewAnthropometry(w, h)
	},
	bmiConf,
	func(a Measurer) []string {
		i := a.(*AnthropometricRatio)
		v, _ := i.Calc()
		c, _ := i.Classify()
		return []string{
			fmt.Sprintf("BMI: %2.f (kg/m^2)", v),
			fmt.Sprintf("BMI classification: %s", c),
		}
	},
)

/**
 * Anthropometry
 */

// Anthropometry represents the basic anthropometric measures (weight, height)
type Anthropometry struct {
	Weight float64
	Height float64
}

// NewAnthropometry returns a new Anthropometry instance.
func NewAnthropometry(weight float64, height float64) *Anthropometry {
	return &Anthropometry{Weight: weight, Height: height}
}

func (a *Anthropometry) String() string {
	return fmt.Sprintf("Weight: %.2f kg\nHeight: %.2f cm", a.Weight, a.Height)
}

// GetName return this measurement name.
func (a *Anthropometry) GetName() string {
	return "Anthropometry"
}

// Result get common representation for this measurement.
func (a *Anthropometry) Result() ([]string, error) {
	return []string{
		fmt.Sprintf("Weight: %.2f kg.", a.Weight),
		fmt.Sprintf("Height: %.2f cm.", a.Height),
	}, nil
}

/**
 * Anthropometric index
 */

// AnthropometricRatio represents a given anthropometric ratio. It's main
// responsibility is to implement Measurer interface for any ratio.
type AnthropometricRatio struct {
	*Anthropometry
	// lim represents the limits for a given classification value
	lim map[int][2]float64
	// prt is the parent measurement
	prt Measurer
	// conf defines the equation configuration for the given ratio
	conf *EquationConf
	// result knows how to represent the results for the given measurement
	result func(Measurer) []string
}

// newAnthropometricRatio create a anthropometric ratio, that is comprised of a
// limit mapper for classification, a parent measumerement function, an
// enquation configuration, and a result function.
// Returns a function that create a new AnthropoRatio instance.
func newAnthropoRatio(
	lim map[int][2]float64,
	prt func(float64, float64) Measurer,
	conf *EquationConf,
	result func(Measurer) []string,
) func(float64, float64) *AnthropometricRatio {
	ai := new(AnthropometricRatio)
	ai.lim = lim
	ai.conf = conf
	ai.result = result
	return func(weight, height float64) *AnthropometricRatio {
		ai.Anthropometry = NewAnthropometry(weight, height)
		ai.prt = prt(weight, height)
		return ai
	}
}

func (i *AnthropometricRatio) String() string {
	prs, _ := i.prt.Result()
	rs, _ := i.Result()
	return fmt.Sprintf("%s\n%s", strings.Join(prs, "\n"), strings.Join(rs, "\n"))
}

// GetName retrieves the common name for any anthropometric index
func (i *AnthropometricRatio) GetName() string {
	return "Anthropometry"
}

// Result returns the measurement representation, and an optional error if any
// violation was made for the measurement.
func (i *AnthropometricRatio) Result() ([]string, error) {
	rs := []string{}

	if _, err := i.Calc(); err != nil {
		return rs, err
	}

	if _, err := i.Classify(); err != nil {
		return rs, err
	}

	prs, err := i.prt.Result()
	if err != nil {
		return rs, err
	}

	rs = append(rs, prs...)
	rs = append(rs, i.result(i)...)
	return rs, nil
}

// Classify returns the classification for the given measurement calc, and an
// optional error.
func (i *AnthropometricRatio) Classify() (string, error) {
	v, err := i.Calc()
	if err != nil {
		return "", err
	}
	return Classifier(v, i.lim, BMIClassification), nil
}

// Calc returns the measurement value, and an optional error.
func (i *AnthropometricRatio) Calc() (float64, error) {
	return i.equation().Calc()
}

// equation returns the Equation instance used to validate and calculate the
// measurement value.
func (i *AnthropometricRatio) equation() Equationer {
	return NewEquation(i.conf.Extract(i.Anthropometry), i.conf)
}

/**
 * Equations
 */

// Equations for calculation of body mass index and body mass index prime.
var (
	bmiConf = NewEquationConf(
		"BMI",
		inParams,
		validators,
		func(e *Equation) float64 {
			w, _ := e.In("weight")
			h, _ := e.In("height")
			return w / math.Pow(h/100, 2)
		},
	)
	bmiPrimeConf = NewEquationConf(
		"BMIPrime",
		inParams,
		validators,
		func(e *Equation) float64 {
			return bmiConf.Calc(e) / 25.0
		},
	)
)

// List of validators for weight and height measures.
var validators = []Validator{ValidateMeasures([]string{"weight", "height"})}

// inParams method define base parameters for anthropometry. It receives an
// interface, that must comply with Anthropometry struct and returns a map
// with weight and height values.
func inParams(i interface{}) InParams {
	ci := i.(*Anthropometry)
	return map[string]float64{
		"weight": ci.Weight,
		"height": ci.Height,
	}
}

/**
 * Classification
 */

// Classification constants for BMI and BMIPrime
const (
	VerySeverelyUnderweight = iota
	SeverelyUnderweight
	Underweight
	Normal
	Overweight
	ObeseClassOne
	ObeseClassTwo
	ObeseClassThree
)

// BMIClassification map classification constants to their string representation.
var BMIClassification = map[int]string{
	VerySeverelyUnderweight: "Very severely underweight",
	SeverelyUnderweight:     "Severely underweight",
	Underweight:             "Underweight",
	Normal:                  "Normal",
	Overweight:              "Overweight",
	ObeseClassOne:           "Obese class one",
	ObeseClassTwo:           "Obese class two",
	ObeseClassThree:         "Obese class three",
}

// Mappers defining limits for each classification constant.
var (
	limitsForBMI = map[int][2]float64{
		VerySeverelyUnderweight: {math.Inf(-1), 15},
		SeverelyUnderweight:     {15, 16},
		Underweight:             {16, 18.5},
		Normal:                  {18.5, 25},
		Overweight:              {25, 30},
		ObeseClassOne:           {30, 35},
		ObeseClassTwo:           {35, 40},
		ObeseClassThree:         {40, math.Inf(+1)},
	}
	limitsForBMIPrime = map[int][2]float64{
		VerySeverelyUnderweight: {math.Inf(-1), 0.60},
		SeverelyUnderweight:     {0.60, 0.64},
		Underweight:             {0.64, 0.74},
		Normal:                  {0.74, 1},
		Overweight:              {1, 1.2},
		ObeseClassOne:           {1.2, 1.4},
		ObeseClassTwo:           {1.4, 1.6},
		ObeseClassThree:         {1.6, math.Inf(+1)},
	}
)
