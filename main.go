package main

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	assess "github.com/joaodubas/phass/assessment"
	"os"
)

func main() {
	bmiPrime := anthropo.NewBMIPrime(98.0, 168.0)
	fmt.Println(bmiPrime)

	p, err := assess.NewPerson("João Paulo Dubas", "1978-Dec-15", assess.Male)
	handleError("Ops person was not born:", err)
	fmt.Printf("Age in years: %.2f\nAge in months %.2f\n", p.Age(), p.AgeInMonths())
}

func handleError(tmpl string, err error) {
	if err != nil {
		fmt.Printf("%s\n", fmt.Sprintf(tmpl, err))
		os.Exit(1)
	}
}
