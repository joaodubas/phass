package assessment

import "time"

const TimeLayout = "2006-Jan-02"

type Measurer interface {
	Result() ([]string, error)
	Classify() (string, error)
	Calc() (float64, error)
}

func NewPerson(fullName string, birth string, gender int) (*Person, error) {
	p := &Person{}
	b, err := time.Parse(TimeLayout, birth)
	if err != nil {
		return p, err
	}
	p.FullName = fullName
	p.Birthday = b
	p.Gender = gender
	return p, nil
}

func NewAssessment(date string) (*Assessment, error) {
	a := &Assessment{}
	d, err := time.Parse(TimeLayout, date)
	if err != nil {
		return a, err
	}
	a.Date = d
	return a, nil
}
