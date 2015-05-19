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

func classifier(value float64, classes map[int][]float64) string {
	class := ""
	for id_, limits := range classes {
		if value >= limits[0] && value < limits[1] {
			class = BMIClassification[id_]
		}
	}
	return class
}
