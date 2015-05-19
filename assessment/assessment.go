package assessment

import "time"

type Classifier interface {
	Classify() string
}

func NewPerson(fullName string, birth string, gender int) (*Person, error) {
	p := &Person{}
	b, err := time.Parse("2006-Jan-02", birth)
	if err != nil {
		return p, err
	}
	p.FullName = fullName
	p.Birthday = b
	p.Gender = gender
	return p, nil
}
