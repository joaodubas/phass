package phass

import (
	"math"
	"testing"
	"time"
)

func TestPersonBirth(t *testing.T) {
	_, err := NewPerson("Someone", "1900-Dec-40", Male)
	if err == nil {
		t.Errorf("Date used should be invalid")
	}
	_, err = NewPerson("Otherone", "1900-Dec-18", Female)
	if err != nil {
		t.Errorf("Everything is all right with this person")
	}
}

func TestPersonAge(t *testing.T) {
	older, err := NewPerson("João Paulo Dubas", "1978-Dec-15", Male)
	if err != nil {
		t.Errorf("Received err creating person %s", err)
	}

	newer, err := NewPerson("Ana Paula Dubas", "1988-Mar-08", Female)
	if err != nil {
		t.Errorf("Received err creating person %s", err)
	}

	cases := []caseAssessment{
		caseAssessment{
			person:      older,
			age:         34.5,
			ageInMonths: 414.0,
		},
		caseAssessment{
			person:      newer,
			age:         25.27,
			ageInMonths: 303.24,
		},
	}

	for _, data := range cases {
		if age := data.person.AgeFromDate(refDate); !floatEqual(age, data.age, 0.9) {
			t.Errorf("Age calculated is %.3f, expected is %.3f", age, data.age)
		}
		if age := data.person.AgeInMonthsFromDate(refDate); !floatEqual(age, data.ageInMonths, 0.1) {
			t.Errorf("Age calculated is %.4f, expected is %.4f", age, data.ageInMonths)
		}
	}
}

func TestAssessmentDate(t *testing.T) {
	_, err := NewAssessment("1900-Dec-40")
	if err == nil {
		t.Errorf("Date used should be invalid")
	}
	_, err = NewAssessment("1900-Dec-18")
	if err != nil {
		t.Errorf("Everything is all right with this assessment")
	}
}

var refDate, _ = time.Parse(TimeLayout, "2013-Jun-15")

type caseAssessment struct {
	person      *Person
	age         float64
	ageInMonths float64
}

func floatEqual(data, expected, limit float64) bool {
	d := math.Abs(data - expected)
	return d < limit
}
