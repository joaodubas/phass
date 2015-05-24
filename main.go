package main

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	bf "github.com/joaodubas/phass/bodyfat"
	skf "github.com/joaodubas/phass/skinfold"
	"os"
)

func main() {
	bmiPrime := anthropo.NewBMIPrime(98.0, 168.0)
	fmt.Println(bmiPrime)

	p, err := assess.NewPerson("João Paulo Dubas", "1978-Dec-15", assess.Male)
	handleError("Ops person was not born:", err)
	fmt.Printf("Age in years: %.2f\nAge in months %.2f\n", p.Age(), p.AgeInMonths())

	a, err := assess.NewAssessment("2015-Dec-15")
	handleError("Ops assessment not done:", err)
	fmt.Println(a)

	skfs := skf.NewSkinfolds(map[int]float64{
		skf.SKFChest:       5.0,
		skf.SKFAbdominal:   10.0,
		skf.SKFThigh:       15.0,
		skf.SKFTriceps:     20.0,
		skf.SKFMidaxillary: 25.0,
		skf.SKFSubscapular: 30.0,
		skf.SKFSuprailiac:  35.0,
	})
	fmt.Println(skfs)

	comp := bf.NewWomenSevenSKF(p, a, skfs)
	fmt.Println(comp)
	comp = bf.NewMenSevenSKF(p, a, skfs)
	fmt.Println(comp)
}

func handleError(tmpl string, err error) {
	if err != nil {
		fmt.Printf("%s\n", fmt.Sprintf(tmpl, err))
		os.Exit(1)
	}
}
