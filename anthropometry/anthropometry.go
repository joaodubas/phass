package anthropometry

func NewAnthropometry(weight float64, height float64) *Anthropometry {
	return &Anthropometry{Weight: weight, Height: height}
}

func NewBMI(weight float64, height float64) *BMI {
	return &BMI{NewAnthropometry(weight, height)}
}

func NewBMIPrime(weight float64, height float64) *BMIPrime {
	return &BMIPrime{NewBMI(weight, height)}
}
