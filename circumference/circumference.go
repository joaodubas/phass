package circumference

import (
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
)

func NewCircumferences(measures map[int]float64) *Circumferences {
	return &Circumferences{Measures: measures}
}

func NewWaistToHipRatio(person *assess.Person, assessment *assess.Assessment, measures map[int]float64) *WaistToHip {
	return &WaistToHip{person, assessment, NewCircumferences(measures)}
}

func NewConicityIndex(a *anthropo.Anthropometry, measures map[int]float64) *ConicityIndex {
	return &ConicityIndex{a, NewCircumferences(measures)}
}
