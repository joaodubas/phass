package main

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	bf "github.com/joaodubas/phass/bodyfat"
	ccf "github.com/joaodubas/phass/circumference"
	skf "github.com/joaodubas/phass/skinfold"
	"os"
	"strings"
)

func main() {
	p, err := assess.NewPerson("João Paulo Dubas", "1978-Dec-15", assess.Male)
	handleError("Ops person was not born:", err)

	a, err := assess.NewAssessment("2015-May-15")
	handleError("Ops assessment not done:", err)

	// add anthropometric
	bmi := anthropo.NewBMIPrime(98.0, 168.0)
	a.AddMeasure(bmi)

	// add skinfold
	skfs := skf.NewSkinfolds(map[int]float64{
		skf.SKFChest:       5.0,
		skf.SKFAbdominal:   10.0,
		skf.SKFThigh:       15.0,
		skf.SKFTriceps:     20.0,
		skf.SKFMidaxillary: 25.0,
		skf.SKFSubscapular: 30.0,
		skf.SKFSuprailiac:  35.0,
	})
	a.AddMeasure(skfs)

	// add bodyfat
	a.AddMeasure(bf.NewMenSevenSKF(p, a, skfs))

	// add circunferences
	ccfs := ccf.NewCircumferences(map[int]float64{
		ccf.CCFWaist: 98.2,
		ccf.CCFHip: 104.1,
	})
	a.AddMeasure(ccfs)

	// add waist-to-hip
	a.AddMeasure(ccf.NewWaistToHipRatio(p, a, ccfs.Measures))

	// add conicity index
	a.AddMeasure(ccf.NewConicityIndex(bmi.Anthropometry, ccfs.Measures))

	// show result
	rs, err := a.Result()
	handleError("Ops an assessment failed:", err)
	fmt.Println(strings.Join(rs, "\n"))
}

func handleError(tmpl string, err error) {
	if err != nil {
		fmt.Printf("%s\n", fmt.Sprintf(tmpl, err))
		os.Exit(1)
	}
}
