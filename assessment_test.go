package phass

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
)

func TestPersonMeasurerInterface(t *testing.T) {
	cases := []struct {
		name   string
		dob    string
		gender int
		err    error
	}{
		{
			name:   "Person 1",
			dob:    "1978-Dec-15",
			gender: Male,
			err:    nil,
		},
		{
			name:   "Person 2",
			dob:    "1988-Mar-08",
			gender: Female,
			err:    nil,
		},
	}

	for _, data := range cases {
		p, err := NewPerson(data.name, data.dob, data.gender)
		if err != data.err {
			t.Error("This is a valid person")
		}

		if p.GetName() != "Person" {
			t.Error("Wrong measurement name")
		}

		rs, err := p.Result()
		if err != nil {
			t.Error("This is a valid result")
		} else if len(rs) != 4 {
			t.Error("There should be 4 information about person")
		}
		for index, item := range rs {
			switch index {
			case 0:
				if !strings.Contains(item, p.FullName) {
					t.Error("Full name should be present")
				}
			case 1:
				if !strings.Contains(item, p.genderRepr()) {
					t.Error("Gender representation should be present")
				}
			case 2:
				if !strings.Contains(item, fmt.Sprintf("%.0f", p.Age())) {
					t.Error("Age in years should be present")
				}
			case 3:
				if !strings.Contains(item, fmt.Sprintf("%.1f", p.AgeInMonths())) {
					t.Error("Age in months should be present")
				}
			}
		}
	}
}

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
		{
			person:      older,
			age:         34.5,
			ageInMonths: 414.0,
		},
		{
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
