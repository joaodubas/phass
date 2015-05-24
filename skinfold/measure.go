package skinfold

import "fmt"

const (
	SKFSubscapular = iota
	SKFTriceps
	SKFBiceps
	SKFChest
	SKFMidaxillary
	SKFSuprailiac
	SKFAbdominal
	SKFThigh
	SKFCalf
)

type Skinfolds struct {
	Measures map[int]float64
}

func (s *Skinfolds) String() string {
	return fmt.Sprintf("Sum %d skinfolds: %.2f mm", len(s.Measures), s.Sum())
}

func (s *Skinfolds) Sum() float64 {
	accum := 0.0
	for _, v := range s.Measures {
		accum += v
	}
	return accum
}

func (s *Skinfolds) SumSpecific(skinfolds []int) float64 {
	accum := 0.0
	for _, skf := range skinfolds {
		accum += s.Measures[skf]
	}
	return accum
}

func NamedSkinfold(name int) string {
	named := map[int]string{
		SKFSubscapular: "subscapular",
		SKFTriceps:     "triceps",
		SKFBiceps:      "biceps",
		SKFChest:       "chest",
		SKFMidaxillary: "mid-axillary",
		SKFSuprailiac:  "suprailiac",
		SKFAbdominal:   "abdominal",
		SKFThigh:       "thigh",
		SKFCalf:        "calf",
	}

	return named[name]
}
