package phass

import (
	"math"
	"testing"
)

/**
 * Test common work for equations
 */

func TestBodyFatCompositionValidation(t *testing.T) {
	newEquation := FactoryBodyCompositionSKF(SKFEquationConf{
		name:     "Dummy Dubas Two SKF",
		gender:   Female,
		lowerAge: 18,
		upperAge: 55,
		skinfolds: []int{
			SKFSuprailiac,
			SKFThigh,
		},
		equation: func(e *Equation) float64 {
			age, _ := e.In("age")
			sum, _ := e.In("sskf")
			d := 1.01 - 0.0001*sum + 0.0000004*sum*sum - 0.000001*age
			return 495/d - 450
		},
	})
	cases := []caseBodyFat{
		newCaseBodyFat(
			female,
			"1988-Mar-15",
			map[int]float64{SKFTriceps: 10.5, SKFSuprailiac: 26.9, SKFThigh: 21.2},
			"Too young",
			0.0,
			ErrInvalidAge.Error(),
		),
		newCaseBodyFat(
			female,
			"2048-Mar-15",
			map[int]float64{SKFTriceps: 10.5, SKFSuprailiac: 26.9, SKFThigh: 21.2},
			"Too old",
			0.0,
			ErrInvalidAge.Error(),
		),
		newCaseBodyFat(
			male,
			"1998-Dec-15",
			map[int]float64{SKFTriceps: 10.5, SKFSuprailiac: 26.9, SKFThigh: 21.2},
			"Wrong gender",
			0.0,
			ErrInvalidGender.Error(),
		),
		newCaseBodyFat(
			female,
			"2008-Mar-15",
			map[int]float64{SKFTriceps: 10.5},
			"Without skinfolds",
			0.0,
			"Missing skinfold suprailiac",
		),
		newCaseBodyFat(
			female,
			"2008-Mar-15",
			map[int]float64{SKFTriceps: 10.5, SKFSuprailiac: 26.9, SKFThigh: 21.2},
			"The answer to life the universe and everything",
			42.00,
			"",
		),
	}

	for _, data := range cases {
		bc := newEquation(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); data.err != "" && err == nil {
			t.Errorf("Case _%s_ failed, should show a validation error", data.name)
		} else if data.err != "" && err.Error() != data.err {
			t.Errorf("Case _%s_ failed, should show proper error message", data.name)
		} else if data.calc > 0.0 && !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

/**
 * Test women equations
 */

func TestFemaleSevenSkinfoldEquation(t *testing.T) {
	cases := []caseBodyFat{
		newCaseBodyFat(female, "2039-Mar-15", map[int]float64{SKFSubscapular: 21.1, SKFTriceps: 17, SKFChest: 11.6, SKFMidaxillary: 12, SKFSuprailiac: 43.9, SKFAbdominal: 40.8, SKFThigh: 34.8}, "for-age-51", 32.385, ""),
		newCaseBodyFat(female, "2011-Mar-15", map[int]float64{SKFSubscapular: 17.1, SKFTriceps: 10.4, SKFChest: 8.3, SKFMidaxillary: 11, SKFSuprailiac: 41.4, SKFAbdominal: 19.9, SKFThigh: 27.2}, "for-age-23", 24.384, ""),
		newCaseBodyFat(female, "2011-Mar-15", map[int]float64{SKFSubscapular: 21.3, SKFTriceps: 14.2, SKFChest: 7.3, SKFMidaxillary: 10.1, SKFSuprailiac: 45.5, SKFAbdominal: 40.1, SKFThigh: 19.6}, "for-age-23", 27.626, ""),
		newCaseBodyFat(female, "2016-Mar-15", map[int]float64{SKFSubscapular: 30.3, SKFTriceps: 16.9, SKFChest: 9.6, SKFMidaxillary: 10.4, SKFSuprailiac: 35, SKFAbdominal: 15.6, SKFThigh: 33.2}, "for-age-28", 26.941, ""),
		newCaseBodyFat(female, "2036-Mar-15", map[int]float64{SKFSubscapular: 31.6, SKFTriceps: 9.1, SKFChest: 12, SKFMidaxillary: 15.3, SKFSuprailiac: 15.7, SKFAbdominal: 24.7, SKFThigh: 19.4}, "for-age-48", 24.749, ""),
		newCaseBodyFat(female, "2036-Mar-15", map[int]float64{SKFSubscapular: 21, SKFTriceps: 11, SKFChest: 7.6, SKFMidaxillary: 18.5, SKFSuprailiac: 34.7, SKFAbdominal: 37.2, SKFThigh: 29.9}, "for-age-48", 29.382, ""),
		newCaseBodyFat(female, "2029-Mar-15", map[int]float64{SKFSubscapular: 32.9, SKFTriceps: 13.8, SKFChest: 9.7, SKFMidaxillary: 18.8, SKFSuprailiac: 41.7, SKFAbdominal: 26.7, SKFThigh: 18}, "for-age-41", 29.191, ""),
		newCaseBodyFat(female, "2010-Mar-15", map[int]float64{SKFSubscapular: 29.9, SKFTriceps: 12.7, SKFChest: 10.6, SKFMidaxillary: 16.1, SKFSuprailiac: 34.8, SKFAbdominal: 39.8, SKFThigh: 25.8}, "for-age-22", 29.127, ""),
		newCaseBodyFat(female, "2019-Mar-15", map[int]float64{SKFSubscapular: 19.5, SKFTriceps: 9.6, SKFChest: 11.4, SKFMidaxillary: 8.7, SKFSuprailiac: 32.3, SKFAbdominal: 28.1, SKFThigh: 20.2}, "for-age-31", 24.042, ""),
		newCaseBodyFat(female, "2037-Mar-15", map[int]float64{SKFSubscapular: 30.7, SKFTriceps: 11.5, SKFChest: 10.5, SKFMidaxillary: 9, SKFSuprailiac: 16.1, SKFAbdominal: 43, SKFThigh: 28.2}, "for-age-49", 27.92, ""),
		newCaseBodyFat(female, "2039-Mar-15", map[int]float64{SKFSubscapular: 22, SKFTriceps: 12.2, SKFChest: 11.3, SKFMidaxillary: 18, SKFSuprailiac: 26, SKFAbdominal: 24.5, SKFThigh: 16.6}, "for-age-51", 25.35, ""),
		newCaseBodyFat(female, "2030-Mar-15", map[int]float64{SKFSubscapular: 15.7, SKFTriceps: 8.8, SKFChest: 9.1, SKFMidaxillary: 7.5, SKFSuprailiac: 41.5, SKFAbdominal: 25.9, SKFThigh: 30.2}, "for-age-42", 26.014, ""),
		newCaseBodyFat(female, "2006-Mar-15", map[int]float64{SKFSubscapular: 35.9, SKFTriceps: 11.8, SKFChest: 12.2, SKFMidaxillary: 6.9, SKFSuprailiac: 22.3, SKFAbdominal: 41, SKFThigh: 17.6}, "for-age-18", 25.877, ""),
		newCaseBodyFat(female, "2019-Mar-15", map[int]float64{SKFSubscapular: 32.1, SKFTriceps: 12.9, SKFChest: 9.6, SKFMidaxillary: 16.5, SKFSuprailiac: 44.6, SKFAbdominal: 22.3, SKFThigh: 29.1}, "for-age-31", 29.327, ""),
		newCaseBodyFat(female, "2012-Mar-15", map[int]float64{SKFSubscapular: 31.1, SKFTriceps: 10.9, SKFChest: 6.7, SKFMidaxillary: 18.1, SKFSuprailiac: 44.6, SKFAbdominal: 19.2, SKFThigh: 22.9}, "for-age-24", 27.051, ""),
		newCaseBodyFat(female, "2043-Mar-15", map[int]float64{SKFSubscapular: 31.9, SKFTriceps: 8.9, SKFChest: 12.8, SKFMidaxillary: 18.8, SKFSuprailiac: 28.5, SKFAbdominal: 29.6, SKFThigh: 29.3}, "for-age-55", 29.793, ""),
		newCaseBodyFat(female, "2037-Mar-15", map[int]float64{SKFSubscapular: 35.2, SKFTriceps: 13.9, SKFChest: 10.8, SKFMidaxillary: 12.4, SKFSuprailiac: 45.1, SKFAbdominal: 21.8, SKFThigh: 34.2}, "for-age-49", 31.254, ""),
		newCaseBodyFat(female, "2029-Mar-15", map[int]float64{SKFSubscapular: 35.4, SKFTriceps: 16.8, SKFChest: 10.3, SKFMidaxillary: 14.7, SKFSuprailiac: 41.6, SKFAbdominal: 17.7, SKFThigh: 23.6}, "for-age-41", 28.986, ""),
		newCaseBodyFat(female, "2031-Mar-15", map[int]float64{SKFSubscapular: 20.4, SKFTriceps: 14.8, SKFChest: 9.9, SKFMidaxillary: 14.8, SKFSuprailiac: 27.6, SKFAbdominal: 34.2, SKFThigh: 17.6}, "for-age-43", 26.162, ""),
		newCaseBodyFat(female, "2015-Mar-15", map[int]float64{SKFSubscapular: 25.4, SKFTriceps: 10.1, SKFChest: 10, SKFMidaxillary: 7.8, SKFSuprailiac: 26.3, SKFAbdominal: 37.1, SKFThigh: 11.6}, "for-age-27", 23.58, ""),
	}

	for _, data := range cases {
		bc := NewWomenSevenSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); err != nil {
			age := data.person.AgeFromDate(data.assessment.Date)
			t.Logf("Got age %f", data.person.AgeFromDate(data.assessment.Date))
			t.Logf("Golt age %f %f", age, math.Ceil(age))
			t.Errorf("Case _%s_ failed, should not show a validation error, instead got %s", data.name, err)
		} else if !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

func TestFemaleThreeSkinfoldEquation(t *testing.T) {
	cases := []caseBodyFat{
		newCaseBodyFat(female, "2014-Mar-15", map[int]float64{SKFSubscapular: 33.3, SKFTriceps: 16.3, SKFChest: 10.9, SKFMidaxillary: 8.9, SKFSuprailiac: 26.6, SKFAbdominal: 46, SKFThigh: 16.1}, "for-age-26", 22.289, ""),
		newCaseBodyFat(female, "2017-Mar-15", map[int]float64{SKFSubscapular: 22.5, SKFTriceps: 16.8, SKFChest: 6.4, SKFMidaxillary: 7.6, SKFSuprailiac: 17.3, SKFAbdominal: 40.1, SKFThigh: 29.8}, "for-age-29", 24.083, ""),
		newCaseBodyFat(female, "2025-Mar-15", map[int]float64{SKFSubscapular: 23.9, SKFTriceps: 16.1, SKFChest: 8.5, SKFMidaxillary: 14.6, SKFSuprailiac: 45.9, SKFAbdominal: 44.8, SKFThigh: 21.1}, "for-age-37", 30.489, ""),
		newCaseBodyFat(female, "2022-Mar-15", map[int]float64{SKFSubscapular: 25.8, SKFTriceps: 10.6, SKFChest: 7.4, SKFMidaxillary: 14.6, SKFSuprailiac: 28.1, SKFAbdominal: 39.4, SKFThigh: 24.9}, "for-age-34", 24.308, ""),
		newCaseBodyFat(female, "2027-Mar-15", map[int]float64{SKFSubscapular: 20.6, SKFTriceps: 9.9, SKFChest: 8.6, SKFMidaxillary: 18.1, SKFSuprailiac: 31.6, SKFAbdominal: 45.8, SKFThigh: 28.5}, "for-age-39", 26.67, ""),
		newCaseBodyFat(female, "2010-Mar-15", map[int]float64{SKFSubscapular: 35.6, SKFTriceps: 13.9, SKFChest: 12.1, SKFMidaxillary: 9.5, SKFSuprailiac: 27.8, SKFAbdominal: 33.1, SKFThigh: 34}, "for-age-22", 27.317, ""),
		newCaseBodyFat(female, "2037-Mar-15", map[int]float64{SKFSubscapular: 24.1, SKFTriceps: 14.3, SKFChest: 12.3, SKFMidaxillary: 11.7, SKFSuprailiac: 31.5, SKFAbdominal: 16.7, SKFThigh: 28}, "for-age-49", 28.502, ""),
		newCaseBodyFat(female, "2027-Mar-15", map[int]float64{SKFSubscapular: 15.4, SKFTriceps: 8.9, SKFChest: 10.3, SKFMidaxillary: 17.4, SKFSuprailiac: 40.9, SKFAbdominal: 26.9, SKFThigh: 19.8}, "for-age-39", 26.545, ""),
		newCaseBodyFat(female, "2037-Mar-15", map[int]float64{SKFSubscapular: 26.7, SKFTriceps: 12.4, SKFChest: 8.8, SKFMidaxillary: 11, SKFSuprailiac: 21.6, SKFAbdominal: 15.9, SKFThigh: 14.8}, "for-age-49", 20.281, ""),
		newCaseBodyFat(female, "2021-Mar-15", map[int]float64{SKFSubscapular: 32.9, SKFTriceps: 10.9, SKFChest: 8.9, SKFMidaxillary: 6.3, SKFSuprailiac: 24.2, SKFAbdominal: 25.5, SKFThigh: 22.5}, "for-age-33", 22.271, ""),
		newCaseBodyFat(female, "2030-Mar-15", map[int]float64{SKFSubscapular: 16.1, SKFTriceps: 15.4, SKFChest: 7.6, SKFMidaxillary: 10.8, SKFSuprailiac: 35.6, SKFAbdominal: 31.9, SKFThigh: 10.5}, "for-age-42", 24.138, ""),
		newCaseBodyFat(female, "2008-Mar-15", map[int]float64{SKFSubscapular: 21.4, SKFTriceps: 8.3, SKFChest: 6.6, SKFMidaxillary: 7, SKFSuprailiac: 37.3, SKFAbdominal: 42.2, SKFThigh: 22}, "for-age-20", 24.685, ""),
		newCaseBodyFat(female, "2013-Mar-15", map[int]float64{SKFSubscapular: 32, SKFTriceps: 12.2, SKFChest: 11.4, SKFMidaxillary: 15.3, SKFSuprailiac: 31.3, SKFAbdominal: 22.2, SKFThigh: 28.4}, "for-age-25", 26.352, ""),
		newCaseBodyFat(female, "2041-Mar-15", map[int]float64{SKFSubscapular: 29.8, SKFTriceps: 14, SKFChest: 8.7, SKFMidaxillary: 8.7, SKFSuprailiac: 16.7, SKFAbdominal: 34, SKFThigh: 29.3}, "for-age-53", 24.351, ""),
		newCaseBodyFat(female, "2023-Mar-15", map[int]float64{SKFSubscapular: 28.7, SKFTriceps: 14.7, SKFChest: 12.8, SKFMidaxillary: 15.6, SKFSuprailiac: 24.4, SKFAbdominal: 25.5, SKFThigh: 17.1}, "for-age-35", 21.929, ""),
		newCaseBodyFat(female, "2010-Mar-15", map[int]float64{SKFSubscapular: 19.3, SKFTriceps: 16.3, SKFChest: 9.2, SKFMidaxillary: 13.4, SKFSuprailiac: 17.8, SKFAbdominal: 34.2, SKFThigh: 26.5}, "for-age-22", 22.561, ""),
		newCaseBodyFat(female, "2042-Mar-15", map[int]float64{SKFSubscapular: 17.3, SKFTriceps: 16.1, SKFChest: 8, SKFMidaxillary: 12.8, SKFSuprailiac: 43.9, SKFAbdominal: 23, SKFThigh: 33.4}, "for-age-54", 34.513, ""),
		newCaseBodyFat(female, "2043-Mar-15", map[int]float64{SKFSubscapular: 18.9, SKFTriceps: 12.6, SKFChest: 6.4, SKFMidaxillary: 9.7, SKFSuprailiac: 21.4, SKFAbdominal: 37.6, SKFThigh: 34.6}, "for-age-55", 27.27, ""),
		newCaseBodyFat(female, "2014-Mar-15", map[int]float64{SKFSubscapular: 15.8, SKFTriceps: 11.8, SKFChest: 10.9, SKFMidaxillary: 13.6, SKFSuprailiac: 25.2, SKFAbdominal: 35.4, SKFThigh: 28}, "for-age-26", 24.244, ""),
		newCaseBodyFat(female, "2017-Mar-15", map[int]float64{SKFSubscapular: 31, SKFTriceps: 16.4, SKFChest: 9.8, SKFMidaxillary: 8.7, SKFSuprailiac: 33.3, SKFAbdominal: 28.7, SKFThigh: 14}, "for-age-29", 24.018, ""),
	}

	for _, data := range cases {
		bc := NewWomenThreeSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); err != nil {
			t.Logf("Got age %f", data.person.AgeFromDate(data.assessment.Date))
			t.Errorf("Case _%s_ failed, should not show a validation error, instead got %s", data.name, err)
		} else if !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

func TestFemaleTwoSkinfoldEquation(t *testing.T) {
	cases := []caseBodyFat{
		newCaseBodyFat(female, "2004-Mar-15", map[int]float64{SKFSubscapular: 28, SKFTriceps: 13.1, SKFChest: 7.5, SKFMidaxillary: 7.5, SKFSuprailiac: 30, SKFAbdominal: 18.5, SKFThigh: 31.1, SKFCalf: 14.5}, "for-age-16", 21.286, ""),
		newCaseBodyFat(female, "2004-Mar-15", map[int]float64{SKFSubscapular: 18.5, SKFTriceps: 9.8, SKFChest: 9.6, SKFMidaxillary: 7.7, SKFSuprailiac: 21.1, SKFAbdominal: 34.1, SKFThigh: 20.8, SKFCalf: 20.6}, "for-age-16", 23.344, ""),
		newCaseBodyFat(female, "1995-Mar-15", map[int]float64{SKFSubscapular: 27.2, SKFTriceps: 16.1, SKFChest: 11.9, SKFMidaxillary: 6.8, SKFSuprailiac: 36.5, SKFAbdominal: 45.4, SKFThigh: 34.6, SKFCalf: 10.4}, "for-age-7", 20.478, ""),
		newCaseBodyFat(female, "2001-Mar-15", map[int]float64{SKFSubscapular: 21.7, SKFTriceps: 12.3, SKFChest: 12.8, SKFMidaxillary: 11.9, SKFSuprailiac: 36.9, SKFAbdominal: 32.7, SKFThigh: 27.9, SKFCalf: 17.4}, "for-age-13", 22.83, ""),
		newCaseBodyFat(female, "1995-Mar-15", map[int]float64{SKFSubscapular: 32.5, SKFTriceps: 13.9, SKFChest: 8.2, SKFMidaxillary: 8.5, SKFSuprailiac: 30.8, SKFAbdominal: 25.1, SKFThigh: 13.7, SKFCalf: 20.1}, "for-age-7", 25.99, ""),
		newCaseBodyFat(female, "2003-Mar-15", map[int]float64{SKFSubscapular: 21.3, SKFTriceps: 11.5, SKFChest: 8.2, SKFMidaxillary: 12.8, SKFSuprailiac: 31.9, SKFAbdominal: 20.1, SKFThigh: 31.2, SKFCalf: 16.8}, "for-age-15", 21.801, ""),
		newCaseBodyFat(female, "1995-Mar-15", map[int]float64{SKFSubscapular: 35.4, SKFTriceps: 13.7, SKFChest: 6.5, SKFMidaxillary: 11.4, SKFSuprailiac: 38.6, SKFAbdominal: 29.3, SKFThigh: 27.6, SKFCalf: 20.7}, "for-age-7", 26.284, ""),
		newCaseBodyFat(female, "2004-Mar-15", map[int]float64{SKFSubscapular: 30.5, SKFTriceps: 10.1, SKFChest: 6.2, SKFMidaxillary: 11.1, SKFSuprailiac: 42, SKFAbdominal: 26.3, SKFThigh: 26.4, SKFCalf: 16.7}, "for-age-16", 20.698, ""),
		newCaseBodyFat(female, "2000-Mar-15", map[int]float64{SKFSubscapular: 29.7, SKFTriceps: 9.3, SKFChest: 9.7, SKFMidaxillary: 12.6, SKFSuprailiac: 36.7, SKFAbdominal: 17.5, SKFThigh: 28.4, SKFCalf: 9.6}, "for-age-12", 14.892, ""),
		newCaseBodyFat(female, "2003-Mar-15", map[int]float64{SKFSubscapular: 29.7, SKFTriceps: 12.8, SKFChest: 8.2, SKFMidaxillary: 16.4, SKFSuprailiac: 24.5, SKFAbdominal: 26.7, SKFThigh: 30.1, SKFCalf: 25.6}, "for-age-15", 29.224, ""),
		newCaseBodyFat(female, "1999-Mar-15", map[int]float64{SKFSubscapular: 16.1, SKFTriceps: 16.1, SKFChest: 11.6, SKFMidaxillary: 11.8, SKFSuprailiac: 42.6, SKFAbdominal: 32.2, SKFThigh: 14.1, SKFCalf: 14.9}, "for-age-11", 23.785, ""),
		newCaseBodyFat(female, "2002-Mar-15", map[int]float64{SKFSubscapular: 24.1, SKFTriceps: 11.4, SKFChest: 6.3, SKFMidaxillary: 6.7, SKFSuprailiac: 38.9, SKFAbdominal: 18, SKFThigh: 34.6, SKFCalf: 18.4}, "for-age-14", 22.903, ""),
		newCaseBodyFat(female, "1995-Mar-15", map[int]float64{SKFSubscapular: 34.5, SKFTriceps: 16.1, SKFChest: 8.7, SKFMidaxillary: 11.6, SKFSuprailiac: 45.3, SKFAbdominal: 15.3, SKFThigh: 27.5, SKFCalf: 19.1}, "for-age-7", 26.872, ""),
		newCaseBodyFat(female, "2001-Mar-15", map[int]float64{SKFSubscapular: 28.1, SKFTriceps: 14.9, SKFChest: 12.6, SKFMidaxillary: 11.1, SKFSuprailiac: 18.5, SKFAbdominal: 44.1, SKFThigh: 26.4, SKFCalf: 21.4}, "for-age-13", 27.681, ""),
		newCaseBodyFat(female, "2002-Mar-15", map[int]float64{SKFSubscapular: 23.1, SKFTriceps: 8.8, SKFChest: 11, SKFMidaxillary: 13.6, SKFSuprailiac: 20, SKFAbdominal: 32.9, SKFThigh: 20.9, SKFCalf: 8.2}, "for-age-14", 13.495, ""),
		newCaseBodyFat(female, "2005-Mar-15", map[int]float64{SKFSubscapular: 18.9, SKFTriceps: 15.8, SKFChest: 12.7, SKFMidaxillary: 11.4, SKFSuprailiac: 25.3, SKFAbdominal: 29, SKFThigh: 20.2, SKFCalf: 16.8}, "for-age-17", 24.961, ""),
		newCaseBodyFat(female, "1999-Mar-15", map[int]float64{SKFSubscapular: 20.9, SKFTriceps: 8, SKFChest: 6.5, SKFMidaxillary: 8.4, SKFSuprailiac: 45.4, SKFAbdominal: 15.6, SKFThigh: 33.5, SKFCalf: 23.1}, "for-age-11", 23.859, ""),
		newCaseBodyFat(female, "1995-Mar-15", map[int]float64{SKFSubscapular: 30.8, SKFTriceps: 14.6, SKFChest: 7.2, SKFMidaxillary: 15, SKFSuprailiac: 22.8, SKFAbdominal: 43.2, SKFThigh: 22.5, SKFCalf: 8.4}, "for-age-7", 17.905, ""),
		newCaseBodyFat(female, "1999-Mar-15", map[int]float64{SKFSubscapular: 27.4, SKFTriceps: 8.1, SKFChest: 7.8, SKFMidaxillary: 9.1, SKFSuprailiac: 44.1, SKFAbdominal: 24.7, SKFThigh: 34.1, SKFCalf: 23.9}, "for-age-11", 24.52, ""),
		newCaseBodyFat(female, "1998-Mar-15", map[int]float64{SKFSubscapular: 27.5, SKFTriceps: 16.5, SKFChest: 7.8, SKFMidaxillary: 14.4, SKFSuprailiac: 33.9, SKFAbdominal: 38.2, SKFThigh: 16.5, SKFCalf: 19.9}, "for-age-10", 27.754, ""),
	}

	for _, data := range cases {
		bc := NewWomenTwoSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); err != nil {
			t.Logf("Got age %f", data.person.AgeFromDate(data.assessment.Date))
			t.Errorf("Case _%s_ failed, should not show a validation error, instead got %s", data.name, err)
		} else if !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

/**
 * Test men equations
 */

func TestMaleSevenSkinfoldEquation(t *testing.T) {
	cases := []caseBodyFat{
		newCaseBodyFat(male, "2014-Dec-15", map[int]float64{SKFSubscapular: 18.4, SKFTriceps: 15.2, SKFChest: 11.3, SKFMidaxillary: 15.7, SKFSuprailiac: 39.8, SKFAbdominal: 29.1, SKFThigh: 24.4}, "for-age-36", 22.46, ""),
		newCaseBodyFat(male, "2030-Dec-15", map[int]float64{SKFSubscapular: 25.6, SKFTriceps: 16.2, SKFChest: 10.8, SKFMidaxillary: 7.6, SKFSuprailiac: 29.1, SKFAbdominal: 23.2, SKFThigh: 15.2}, "for-age-52", 21.234, ""),
		newCaseBodyFat(male, "2031-Dec-15", map[int]float64{SKFSubscapular: 24.4, SKFTriceps: 11.4, SKFChest: 11, SKFMidaxillary: 13.9, SKFSuprailiac: 34.2, SKFAbdominal: 44.5, SKFThigh: 10.9}, "for-age-53", 24.242, ""),
		newCaseBodyFat(male, "2005-Dec-15", map[int]float64{SKFSubscapular: 27.8, SKFTriceps: 14.2, SKFChest: 10.7, SKFMidaxillary: 15.1, SKFSuprailiac: 22.6, SKFAbdominal: 35.8, SKFThigh: 27.8}, "for-age-27", 21.306, ""),
		newCaseBodyFat(male, "2014-Dec-15", map[int]float64{SKFSubscapular: 22.8, SKFTriceps: 14.6, SKFChest: 6.7, SKFMidaxillary: 13.1, SKFSuprailiac: 15.5, SKFAbdominal: 27.1, SKFThigh: 33.8}, "for-age-36", 19.94, ""),
		newCaseBodyFat(male, "2009-Dec-15", map[int]float64{SKFSubscapular: 16.4, SKFTriceps: 13.3, SKFChest: 7.5, SKFMidaxillary: 15.2, SKFSuprailiac: 44.9, SKFAbdominal: 29.9, SKFThigh: 15}, "for-age-31", 20.384, ""),
		newCaseBodyFat(male, "2014-Dec-15", map[int]float64{SKFSubscapular: 21.2, SKFTriceps: 10.5, SKFChest: 11.4, SKFMidaxillary: 9.8, SKFSuprailiac: 45, SKFAbdominal: 40.7, SKFThigh: 31.2}, "for-age-36", 24.31, ""),
		newCaseBodyFat(male, "2031-Dec-15", map[int]float64{SKFSubscapular: 29.8, SKFTriceps: 15.4, SKFChest: 6.8, SKFMidaxillary: 16.2, SKFSuprailiac: 43.8, SKFAbdominal: 31.4, SKFThigh: 25.2}, "for-age-53", 26.41, ""),
		newCaseBodyFat(male, "2005-Dec-15", map[int]float64{SKFSubscapular: 26.5, SKFTriceps: 12.8, SKFChest: 10.7, SKFMidaxillary: 13.4, SKFSuprailiac: 38.1, SKFAbdominal: 41.6, SKFThigh: 24.1}, "for-age-27", 22.841, ""),
		newCaseBodyFat(male, "2028-Dec-15", map[int]float64{SKFSubscapular: 32.4, SKFTriceps: 12.3, SKFChest: 10.6, SKFMidaxillary: 16.1, SKFSuprailiac: 37.2, SKFAbdominal: 28.8, SKFThigh: 32.3}, "for-age-50", 26.14, ""),
		newCaseBodyFat(male, "2003-Dec-15", map[int]float64{SKFSubscapular: 22.5, SKFTriceps: 11.6, SKFChest: 7.3, SKFMidaxillary: 7.1, SKFSuprailiac: 34.1, SKFAbdominal: 44.1, SKFThigh: 25}, "for-age-25", 20.772, ""),
		newCaseBodyFat(male, "2022-Dec-15", map[int]float64{SKFSubscapular: 27.9, SKFTriceps: 14.4, SKFChest: 11.2, SKFMidaxillary: 8.7, SKFSuprailiac: 19.1, SKFAbdominal: 23.8, SKFThigh: 12.5}, "for-age-44", 18.852, ""),
		newCaseBodyFat(male, "2028-Dec-15", map[int]float64{SKFSubscapular: 26.2, SKFTriceps: 14.8, SKFChest: 8.5, SKFMidaxillary: 6.8, SKFSuprailiac: 45.2, SKFAbdominal: 19.6, SKFThigh: 25}, "for-age-50", 23.332, ""),
		newCaseBodyFat(male, "2023-Dec-15", map[int]float64{SKFSubscapular: 27.3, SKFTriceps: 11.7, SKFChest: 7.8, SKFMidaxillary: 15.7, SKFSuprailiac: 44.1, SKFAbdominal: 29.2, SKFThigh: 29.6}, "for-age-45", 24.989, ""),
		newCaseBodyFat(male, "2021-Dec-15", map[int]float64{SKFSubscapular: 17.5, SKFTriceps: 14.7, SKFChest: 7.3, SKFMidaxillary: 16.5, SKFSuprailiac: 19.2, SKFAbdominal: 45.4, SKFThigh: 31.1}, "for-age-43", 23.106, ""),
		newCaseBodyFat(male, "2027-Dec-15", map[int]float64{SKFSubscapular: 16, SKFTriceps: 15.1, SKFChest: 7.2, SKFMidaxillary: 14, SKFSuprailiac: 22.4, SKFAbdominal: 23.4, SKFThigh: 12.7}, "for-age-49", 18.558, ""),
		newCaseBodyFat(male, "2020-Dec-15", map[int]float64{SKFSubscapular: 15.6, SKFTriceps: 10.8, SKFChest: 12.1, SKFMidaxillary: 14.5, SKFSuprailiac: 20.1, SKFAbdominal: 31.8, SKFThigh: 18.1}, "for-age-42", 19.322, ""),
		newCaseBodyFat(male, "2018-Dec-15", map[int]float64{SKFSubscapular: 29.6, SKFTriceps: 13.7, SKFChest: 12.8, SKFMidaxillary: 12.7, SKFSuprailiac: 31.6, SKFAbdominal: 21.4, SKFThigh: 34.9}, "for-age-40", 23.315, ""),
		newCaseBodyFat(male, "2000-Dec-15", map[int]float64{SKFSubscapular: 29.9, SKFTriceps: 8.8, SKFChest: 10, SKFMidaxillary: 18.7, SKFSuprailiac: 37.5, SKFAbdominal: 31.2, SKFThigh: 11.7}, "for-age-22", 19.915, ""),
		newCaseBodyFat(male, "2020-Dec-15", map[int]float64{SKFSubscapular: 27.5, SKFTriceps: 13.4, SKFChest: 8, SKFMidaxillary: 8, SKFSuprailiac: 45.8, SKFAbdominal: 20.1, SKFThigh: 18.9}, "for-age-42", 21.743, ""),
	}

	for _, data := range cases {
		bc := NewMenSevenSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); err != nil {
			t.Errorf("Case _%s_ failed, should not show a validation error, instead got %s", data.name, err)
		} else if !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

func TestMaleThreeSkinfoldEquation(t *testing.T) {
	cases := []caseBodyFat{
		newCaseBodyFat(male, "2019-Dec-15", map[int]float64{SKFSubscapular: 23.8, SKFTriceps: 12, SKFChest: 6, SKFMidaxillary: 11.3, SKFSuprailiac: 37.7, SKFAbdominal: 43.6, SKFThigh: 25.9}, "for-age-41", 23.444, ""),
		newCaseBodyFat(male, "2002-Dec-15", map[int]float64{SKFSubscapular: 32.7, SKFTriceps: 14.2, SKFChest: 10.8, SKFMidaxillary: 16.8, SKFSuprailiac: 15.1, SKFAbdominal: 20.6, SKFThigh: 15}, "for-age-24", 13.358, ""),
		newCaseBodyFat(male, "2029-Dec-15", map[int]float64{SKFSubscapular: 23.5, SKFTriceps: 13.2, SKFChest: 10.3, SKFMidaxillary: 16.6, SKFSuprailiac: 20.2, SKFAbdominal: 24.8, SKFThigh: 30.9}, "for-age-51", 22.031, ""),
		newCaseBodyFat(male, "2016-Dec-15", map[int]float64{SKFSubscapular: 32.6, SKFTriceps: 9.7, SKFChest: 9.8, SKFMidaxillary: 15.3, SKFSuprailiac: 25.6, SKFAbdominal: 26.5, SKFThigh: 19.4}, "for-age-38", 17.636, ""),
		newCaseBodyFat(male, "2009-Dec-15", map[int]float64{SKFSubscapular: 23, SKFTriceps: 13.9, SKFChest: 10.4, SKFMidaxillary: 14.7, SKFSuprailiac: 31.3, SKFAbdominal: 44.4, SKFThigh: 34.6}, "for-age-31", 25.833, ""),
		newCaseBodyFat(male, "2019-Dec-15", map[int]float64{SKFSubscapular: 27, SKFTriceps: 13.6, SKFChest: 6.6, SKFMidaxillary: 15, SKFSuprailiac: 31.7, SKFAbdominal: 17.1, SKFThigh: 32.4}, "for-age-41", 18.092, ""),
		newCaseBodyFat(male, "2012-Dec-15", map[int]float64{SKFSubscapular: 17, SKFTriceps: 9, SKFChest: 6.3, SKFMidaxillary: 6.1, SKFSuprailiac: 19.1, SKFAbdominal: 45.3, SKFThigh: 12.8}, "for-age-34", 19.628, ""),
		newCaseBodyFat(male, "2017-Dec-15", map[int]float64{SKFSubscapular: 25.5, SKFTriceps: 13.1, SKFChest: 8.7, SKFMidaxillary: 17.1, SKFSuprailiac: 21.6, SKFAbdominal: 33.8, SKFThigh: 22.7}, "for-age-39", 20.424, ""),
		newCaseBodyFat(male, "2021-Dec-15", map[int]float64{SKFSubscapular: 28, SKFTriceps: 15.9, SKFChest: 10.4, SKFMidaxillary: 15.8, SKFSuprailiac: 45.8, SKFAbdominal: 39.8, SKFThigh: 12.9}, "for-age-43", 20.301, ""),
		newCaseBodyFat(male, "2020-Dec-15", map[int]float64{SKFSubscapular: 33.1, SKFTriceps: 9.5, SKFChest: 12.1, SKFMidaxillary: 18.4, SKFSuprailiac: 23, SKFAbdominal: 30.4, SKFThigh: 29.5}, "for-age-42", 22.625, ""),
		newCaseBodyFat(male, "2026-Dec-15", map[int]float64{SKFSubscapular: 16.9, SKFTriceps: 15.7, SKFChest: 8.1, SKFMidaxillary: 16.9, SKFSuprailiac: 28, SKFAbdominal: 24.3, SKFThigh: 13.7}, "for-age-48", 15.964, ""),
		newCaseBodyFat(male, "2025-Dec-15", map[int]float64{SKFSubscapular: 18.3, SKFTriceps: 13.4, SKFChest: 10.9, SKFMidaxillary: 10.9, SKFSuprailiac: 19.3, SKFAbdominal: 16, SKFThigh: 25.2}, "for-age-47", 17.619, ""),
		newCaseBodyFat(male, "2015-Dec-15", map[int]float64{SKFSubscapular: 27.6, SKFTriceps: 11.1, SKFChest: 7.8, SKFMidaxillary: 9.1, SKFSuprailiac: 33.3, SKFAbdominal: 38, SKFThigh: 29.9}, "for-age-37", 23.031, ""),
		newCaseBodyFat(male, "2017-Dec-15", map[int]float64{SKFSubscapular: 18, SKFTriceps: 9.9, SKFChest: 7.7, SKFMidaxillary: 14.6, SKFSuprailiac: 39.4, SKFAbdominal: 45.1, SKFThigh: 32.2}, "for-age-39", 25.673, ""),
		newCaseBodyFat(male, "2029-Dec-15", map[int]float64{SKFSubscapular: 35.3, SKFTriceps: 12.8, SKFChest: 6.8, SKFMidaxillary: 11.6, SKFSuprailiac: 41.2, SKFAbdominal: 39.7, SKFThigh: 25.4}, "for-age-51", 23.646, ""),
		newCaseBodyFat(male, "1997-Dec-15", map[int]float64{SKFSubscapular: 29.3, SKFTriceps: 10.5, SKFChest: 7.2, SKFMidaxillary: 14.8, SKFSuprailiac: 19.2, SKFAbdominal: 28.7, SKFThigh: 10.7}, "for-age-19", 12.859, ""),
		newCaseBodyFat(male, "2013-Dec-15", map[int]float64{SKFSubscapular: 19.8, SKFTriceps: 11, SKFChest: 7.5, SKFMidaxillary: 6.7, SKFSuprailiac: 39.1, SKFAbdominal: 44.3, SKFThigh: 29.9}, "for-age-35", 24.361, ""),
		newCaseBodyFat(male, "2015-Dec-15", map[int]float64{SKFSubscapular: 28.9, SKFTriceps: 12.5, SKFChest: 11.8, SKFMidaxillary: 10.3, SKFSuprailiac: 41.9, SKFAbdominal: 17.5, SKFThigh: 22.5}, "for-age-37", 16.398, ""),
		newCaseBodyFat(male, "2016-Dec-15", map[int]float64{SKFSubscapular: 30.1, SKFTriceps: 12.8, SKFChest: 6.6, SKFMidaxillary: 17.6, SKFSuprailiac: 28.5, SKFAbdominal: 39.9, SKFThigh: 24}, "for-age-38", 21.757, ""),
		newCaseBodyFat(male, "2005-Dec-15", map[int]float64{SKFSubscapular: 30, SKFTriceps: 16.1, SKFChest: 6.3, SKFMidaxillary: 14.4, SKFSuprailiac: 29.4, SKFAbdominal: 42, SKFThigh: 19.5}, "for-age-27", 19.758, ""),
	}

	for _, data := range cases {
		bc := NewMenThreeSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); err != nil {
			t.Errorf("Case _%s_ failed, should not show a validation error, instead got %s", data.name, err)
		} else if !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

func TestMaleTwoSkinfoldEquation(t *testing.T) {
	cases := []caseBodyFat{
		newCaseBodyFat(male, "1994-Dec-15", map[int]float64{SKFSubscapular: 17.6, SKFTriceps: 8.5, SKFChest: 9.8, SKFMidaxillary: 10.6, SKFSuprailiac: 40.3, SKFAbdominal: 28.4, SKFThigh: 34.6, SKFCalf: 18.4}, "for-age-16", 21.509, ""),
		newCaseBodyFat(male, "1984-Dec-15", map[int]float64{SKFSubscapular: 26.6, SKFTriceps: 12, SKFChest: 9.7, SKFMidaxillary: 16.7, SKFSuprailiac: 43.9, SKFAbdominal: 42.8, SKFThigh: 10.3, SKFCalf: 24.1}, "for-age-6", 27.121, ""),
		newCaseBodyFat(male, "1994-Dec-15", map[int]float64{SKFSubscapular: 19.7, SKFTriceps: 16.1, SKFChest: 8.8, SKFMidaxillary: 15.2, SKFSuprailiac: 17, SKFAbdominal: 35.3, SKFThigh: 28.7, SKFCalf: 24.4}, "for-age-16", 29.805, ""),
		newCaseBodyFat(male, "1989-Dec-15", map[int]float64{SKFSubscapular: 20.5, SKFTriceps: 10.3, SKFChest: 11.2, SKFMidaxillary: 16.3, SKFSuprailiac: 25, SKFAbdominal: 39.4, SKFThigh: 11.9, SKFCalf: 23.7}, "for-age-11", 25.84, ""),
		newCaseBodyFat(male, "1989-Dec-15", map[int]float64{SKFSubscapular: 21.2, SKFTriceps: 11.3, SKFChest: 10.8, SKFMidaxillary: 18.7, SKFSuprailiac: 27.4, SKFAbdominal: 18.1, SKFThigh: 25.8, SKFCalf: 22.8}, "for-age-11", 25.901, ""),
		newCaseBodyFat(male, "1985-Dec-15", map[int]float64{SKFSubscapular: 31.7, SKFTriceps: 9.4, SKFChest: 12.3, SKFMidaxillary: 10.4, SKFSuprailiac: 19.9, SKFAbdominal: 38.2, SKFThigh: 26.9, SKFCalf: 19.9}, "for-age-7", 22.973, ""),
		newCaseBodyFat(male, "1990-Dec-15", map[int]float64{SKFSubscapular: 18.9, SKFTriceps: 15.9, SKFChest: 8, SKFMidaxillary: 7.6, SKFSuprailiac: 33.2, SKFAbdominal: 26.4, SKFThigh: 23.5, SKFCalf: 16.7}, "for-age-12", 24.986, ""),
		newCaseBodyFat(male, "1985-Dec-15", map[int]float64{SKFSubscapular: 20.3, SKFTriceps: 11.4, SKFChest: 6.5, SKFMidaxillary: 15.9, SKFSuprailiac: 33.4, SKFAbdominal: 15.7, SKFThigh: 18.3, SKFCalf: 23.1}, "for-age-7", 26.145, ""),
		newCaseBodyFat(male, "1992-Dec-15", map[int]float64{SKFSubscapular: 31.5, SKFTriceps: 16.9, SKFChest: 7, SKFMidaxillary: 13.2, SKFSuprailiac: 31.4, SKFAbdominal: 41, SKFThigh: 18.3, SKFCalf: 17.2}, "for-age-14", 25.901, ""),
		newCaseBodyFat(male, "1990-Dec-15", map[int]float64{SKFSubscapular: 29.4, SKFTriceps: 13.7, SKFChest: 12.8, SKFMidaxillary: 8.3, SKFSuprailiac: 45.9, SKFAbdominal: 23.5, SKFThigh: 21.4, SKFCalf: 14.9}, "for-age-12", 22.546, ""),
		newCaseBodyFat(male, "1989-Dec-15", map[int]float64{SKFSubscapular: 25.5, SKFTriceps: 13.7, SKFChest: 6.9, SKFMidaxillary: 12.1, SKFSuprailiac: 31.3, SKFAbdominal: 40, SKFThigh: 29.6, SKFCalf: 23.3}, "for-age-11", 27.67, ""),
		newCaseBodyFat(male, "1994-Dec-15", map[int]float64{SKFSubscapular: 19.4, SKFTriceps: 13.7, SKFChest: 11.4, SKFMidaxillary: 13.4, SKFSuprailiac: 42, SKFAbdominal: 34.5, SKFThigh: 30.3, SKFCalf: 25.5}, "for-age-16", 29.012, ""),
		newCaseBodyFat(male, "1987-Dec-15", map[int]float64{SKFSubscapular: 29.7, SKFTriceps: 9.8, SKFChest: 7.4, SKFMidaxillary: 9.4, SKFSuprailiac: 36.3, SKFAbdominal: 38.6, SKFThigh: 16, SKFCalf: 24.1}, "for-age-9", 25.779, ""),
		newCaseBodyFat(male, "1986-Dec-15", map[int]float64{SKFSubscapular: 32.5, SKFTriceps: 9.7, SKFChest: 12.2, SKFMidaxillary: 14.4, SKFSuprailiac: 37, SKFAbdominal: 44.3, SKFThigh: 19.5, SKFCalf: 11.5}, "for-age-8", 18.032, ""),
		newCaseBodyFat(male, "1992-Dec-15", map[int]float64{SKFSubscapular: 19.8, SKFTriceps: 16, SKFChest: 12, SKFMidaxillary: 16.1, SKFSuprailiac: 17, SKFAbdominal: 25.2, SKFThigh: 23.2, SKFCalf: 6.4}, "for-age-14", 18.764, ""),
		newCaseBodyFat(male, "1989-Dec-15", map[int]float64{SKFSubscapular: 27.2, SKFTriceps: 11.7, SKFChest: 7.5, SKFMidaxillary: 15.9, SKFSuprailiac: 29, SKFAbdominal: 17.5, SKFThigh: 10, SKFCalf: 24.4}, "for-age-11", 27.121, ""),
		newCaseBodyFat(male, "1993-Dec-15", map[int]float64{SKFSubscapular: 24.4, SKFTriceps: 10.3, SKFChest: 12.5, SKFMidaxillary: 9.3, SKFSuprailiac: 19.6, SKFAbdominal: 29.5, SKFThigh: 20.1, SKFCalf: 23.4}, "for-age-15", 25.657, ""),
		newCaseBodyFat(male, "1985-Dec-15", map[int]float64{SKFSubscapular: 29.1, SKFTriceps: 11.2, SKFChest: 12.4, SKFMidaxillary: 12, SKFSuprailiac: 34.8, SKFAbdominal: 18.1, SKFThigh: 11.1, SKFCalf: 14.8}, "for-age-7", 20.96, ""),
		newCaseBodyFat(male, "1985-Dec-15", map[int]float64{SKFSubscapular: 16.8, SKFTriceps: 16.6, SKFChest: 10.8, SKFMidaxillary: 7.6, SKFSuprailiac: 21.3, SKFAbdominal: 45.4, SKFThigh: 31.3, SKFCalf: 14.5}, "for-age-7", 24.071, ""),
		newCaseBodyFat(male, "1985-Dec-15", map[int]float64{SKFSubscapular: 22.1, SKFTriceps: 16.4, SKFChest: 10.2, SKFMidaxillary: 9.8, SKFSuprailiac: 35.9, SKFAbdominal: 15.9, SKFThigh: 17.8, SKFCalf: 17.7}, "for-age-7", 25.901, ""),
	}

	for _, data := range cases {
		bc := NewMenTwoSKF(data.person, data.assessment, data.skinfold)
		if calc, err := bc.Calc(); err != nil {
			t.Errorf("Case _%s_ failed, should not show a validation error, instead got %s", data.name, err)
		} else if !floatEqual(calc, data.calc, 0.009) {
			t.Errorf("Case _%s_ failed, should have value %.4f, instead got %.4f", data.name, data.calc, calc)
		}
	}
}

/**
 * Common data for testing
 */

type caseBodyFat struct {
	person     *Person
	assessment *Assessment
	skinfold   *Skinfolds
	name       string
	calc       float64
	err        string
}

func newCaseBodyFat(p *Person, d string, s map[int]float64, name string, calc float64, err string) caseBodyFat {
	assessment, _ := NewAssessment(d)
	return caseBodyFat{
		person:     p,
		assessment: assessment,
		skinfold:   NewSkinfolds(s),
		name:       name,
		calc:       calc,
		err:        err,
	}
}

var (
	male, _   = NewPerson("Joao Paulo Dubas", "1978-Dec-15", Male)
	female, _ = NewPerson("Ana Paula Dubas", "1988-Mar-15", Female)
)
