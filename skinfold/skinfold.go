package skinfold

func NewSkinfolds(measures map[int]float64) *Skinfolds {
	return &Skinfolds{Measures: measures}
}