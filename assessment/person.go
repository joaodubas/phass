package assessment

import (
	"fmt"
	"time"
)

const (
	Male int = iota
	Female
)

type Person struct {
	FullName string
	Birthday time.Time
	Gender   int
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %s\nGender: %s\nAge: %.0f", p.FullName, p.genderRepr(), p.Age())
}

func (p *Person) Age() float64 {
	return elapsedFromNowIn(p.Birthday, secondsInYear)
}

func (p *Person) AgeInMonths() float64 {
	return elapsedFromNowIn(p.Birthday, secondsInMonth)
}

func (p *Person) genderRepr() string {
	choices := map[int]string{
		Male:   "Male",
		Female: "Female",
	}
	return choices[p.Gender]
}

func elapsedFromNowIn(t time.Time, in float64) float64 {
	return time.Now().Sub(t).Seconds() / in
}

var (
	daysInYear      = 365.25
	daysInMonth     = daysInYear / 12
	secondsInMinute = 60.0
	minutesInHour   = 60.0
	hoursInDay      = 24.0
	secondsInDay    = hoursInDay * minutesInHour * secondsInMinute
	secondsInMonth  = daysInMonth * secondsInDay
	secondsInYear   = daysInYear * secondsInDay
)
