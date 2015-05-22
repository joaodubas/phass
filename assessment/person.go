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

func (p *Person) AgeFromDate(t time.Time) float64 {
	return elapsedFromDateIn(p.Birthday, t, secondsInYear)
}

func (p *Person) AgeInMonthsFromDate(t time.Time) float64 {
	return elapsedFromDateIn(p.Birthday, t, secondsInMonth)
}

func (p *Person) genderRepr() string {
	choices := map[int]string{
		Male:   "Male",
		Female: "Female",
	}
	return choices[p.Gender]
}

type Assessment struct {
	Date time.Time
}

func (a *Assessment) String() string {
	return fmt.Sprintf("Assessment made in %s", a.Date.Format(TimeLayout))
}

func elapsedFromNowIn(t time.Time, in float64) float64 {
	return elapsedFromDateIn(t, time.Now(), in)
}

func elapsedFromDateIn(from time.Time, to time.Time, in float64) float64 {
	return to.Sub(from).Seconds() / in
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
