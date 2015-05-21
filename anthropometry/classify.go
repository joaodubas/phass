package anthropometry

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
