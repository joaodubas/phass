package circumference

import (
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
)

func NewCircumferences(measures map[int]float64) *Circumferences {
	return &Circumferences{Measures: measures}
}

func NewWaistToHipRation(p *assess.Person, measures map[int]float64) *WaistToHip {
	return &WaistToHip{p, NewCircumferences(measures)}
}

func NewConicityIndex(a *anthropo.Anthropometry, measures map[int]float64) *ConicityIndex {
	return &ConicityIndex{a, NewCircumferences(measures)}
}
