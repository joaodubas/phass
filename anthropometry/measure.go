package anthropometry

import "fmt"

type Anthropometry struct {
	Weight float64
	Height float64
}

func (a *Anthropometry) String() string {
	return fmt.Sprintf("Weight: %.2f kg\nHeight: %.2f cm", a.Weight, a.Height)
}
