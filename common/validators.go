package common

import "fmt"

func ValidateGender(expect int) Validator {
	return func(e *Equation) (bool, error) {
		if g, ok := e.In("gender"); !ok {
			return false, fmt.Errorf("Missing gender")
		} else if int(g) != expect {
			return false, fmt.Errorf("Valid for gender %d", expect)
		}
		return true, nil
	}
}

func ValidateAge(lower, upper float64) Validator {
	return func(e *Equation) (bool, error) {
		if age, ok := e.In("age"); !ok {
			return false, fmt.Errorf("Missing age measure")
		} else if age < lower || age > upper {
			return false, fmt.Errorf("Valid for ages between %.0f and %.0f", lower, upper)
		}
		return true, nil
	}
}

func ValidateMeasures(expect []string) Validator {
	return func(e *Equation) (bool, error) {
		for _, k := range expect {
			if _, ok := e.In(k); !ok {
				return false, fmt.Errorf("Missing %s measure", k)
			}
		}
		return true, nil
	}
}
