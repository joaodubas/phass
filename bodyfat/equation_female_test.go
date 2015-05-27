package bodyfat

import (
	skf "github.com/joaodubas/phass/skinfold"
	"math"
	"testing"
)

func TestFemaleSevenSkinfoldEquation(t *testing.T) {
	cases := []case_{
		newCase_(female, "2039-Mar-15", map[int]float64{skf.SKFSubscapular: 21.1, skf.SKFTriceps: 17, skf.SKFChest: 11.6, skf.SKFMidaxillary: 12, skf.SKFSuprailiac: 43.9, skf.SKFAbdominal: 40.8, skf.SKFThigh: 34.8}, "for-age-51", 32.385, ""),
		newCase_(female, "2011-Mar-15", map[int]float64{skf.SKFSubscapular: 17.1, skf.SKFTriceps: 10.4, skf.SKFChest: 8.3, skf.SKFMidaxillary: 11, skf.SKFSuprailiac: 41.4, skf.SKFAbdominal: 19.9, skf.SKFThigh: 27.2}, "for-age-23", 24.384, ""),
		newCase_(female, "2011-Mar-15", map[int]float64{skf.SKFSubscapular: 21.3, skf.SKFTriceps: 14.2, skf.SKFChest: 7.3, skf.SKFMidaxillary: 10.1, skf.SKFSuprailiac: 45.5, skf.SKFAbdominal: 40.1, skf.SKFThigh: 19.6}, "for-age-23", 27.626, ""),
		newCase_(female, "2016-Mar-15", map[int]float64{skf.SKFSubscapular: 30.3, skf.SKFTriceps: 16.9, skf.SKFChest: 9.6, skf.SKFMidaxillary: 10.4, skf.SKFSuprailiac: 35, skf.SKFAbdominal: 15.6, skf.SKFThigh: 33.2}, "for-age-28", 26.941, ""),
		newCase_(female, "2036-Mar-15", map[int]float64{skf.SKFSubscapular: 31.6, skf.SKFTriceps: 9.1, skf.SKFChest: 12, skf.SKFMidaxillary: 15.3, skf.SKFSuprailiac: 15.7, skf.SKFAbdominal: 24.7, skf.SKFThigh: 19.4}, "for-age-48", 24.749, ""),
		newCase_(female, "2036-Mar-15", map[int]float64{skf.SKFSubscapular: 21, skf.SKFTriceps: 11, skf.SKFChest: 7.6, skf.SKFMidaxillary: 18.5, skf.SKFSuprailiac: 34.7, skf.SKFAbdominal: 37.2, skf.SKFThigh: 29.9}, "for-age-48", 29.382, ""),
		newCase_(female, "2029-Mar-15", map[int]float64{skf.SKFSubscapular: 32.9, skf.SKFTriceps: 13.8, skf.SKFChest: 9.7, skf.SKFMidaxillary: 18.8, skf.SKFSuprailiac: 41.7, skf.SKFAbdominal: 26.7, skf.SKFThigh: 18}, "for-age-41", 29.191, ""),
		newCase_(female, "2010-Mar-15", map[int]float64{skf.SKFSubscapular: 29.9, skf.SKFTriceps: 12.7, skf.SKFChest: 10.6, skf.SKFMidaxillary: 16.1, skf.SKFSuprailiac: 34.8, skf.SKFAbdominal: 39.8, skf.SKFThigh: 25.8}, "for-age-22", 29.127, ""),
		newCase_(female, "2019-Mar-15", map[int]float64{skf.SKFSubscapular: 19.5, skf.SKFTriceps: 9.6, skf.SKFChest: 11.4, skf.SKFMidaxillary: 8.7, skf.SKFSuprailiac: 32.3, skf.SKFAbdominal: 28.1, skf.SKFThigh: 20.2}, "for-age-31", 24.042, ""),
		newCase_(female, "2037-Mar-15", map[int]float64{skf.SKFSubscapular: 30.7, skf.SKFTriceps: 11.5, skf.SKFChest: 10.5, skf.SKFMidaxillary: 9, skf.SKFSuprailiac: 16.1, skf.SKFAbdominal: 43, skf.SKFThigh: 28.2}, "for-age-49", 27.92, ""),
		newCase_(female, "2039-Mar-15", map[int]float64{skf.SKFSubscapular: 22, skf.SKFTriceps: 12.2, skf.SKFChest: 11.3, skf.SKFMidaxillary: 18, skf.SKFSuprailiac: 26, skf.SKFAbdominal: 24.5, skf.SKFThigh: 16.6}, "for-age-51", 25.35, ""),
		newCase_(female, "2030-Mar-15", map[int]float64{skf.SKFSubscapular: 15.7, skf.SKFTriceps: 8.8, skf.SKFChest: 9.1, skf.SKFMidaxillary: 7.5, skf.SKFSuprailiac: 41.5, skf.SKFAbdominal: 25.9, skf.SKFThigh: 30.2}, "for-age-42", 26.014, ""),
		newCase_(female, "2006-Mar-15", map[int]float64{skf.SKFSubscapular: 35.9, skf.SKFTriceps: 11.8, skf.SKFChest: 12.2, skf.SKFMidaxillary: 6.9, skf.SKFSuprailiac: 22.3, skf.SKFAbdominal: 41, skf.SKFThigh: 17.6}, "for-age-18", 25.877, ""),
		newCase_(female, "2019-Mar-15", map[int]float64{skf.SKFSubscapular: 32.1, skf.SKFTriceps: 12.9, skf.SKFChest: 9.6, skf.SKFMidaxillary: 16.5, skf.SKFSuprailiac: 44.6, skf.SKFAbdominal: 22.3, skf.SKFThigh: 29.1}, "for-age-31", 29.327, ""),
		newCase_(female, "2012-Mar-15", map[int]float64{skf.SKFSubscapular: 31.1, skf.SKFTriceps: 10.9, skf.SKFChest: 6.7, skf.SKFMidaxillary: 18.1, skf.SKFSuprailiac: 44.6, skf.SKFAbdominal: 19.2, skf.SKFThigh: 22.9}, "for-age-24", 27.051, ""),
		newCase_(female, "2043-Mar-15", map[int]float64{skf.SKFSubscapular: 31.9, skf.SKFTriceps: 8.9, skf.SKFChest: 12.8, skf.SKFMidaxillary: 18.8, skf.SKFSuprailiac: 28.5, skf.SKFAbdominal: 29.6, skf.SKFThigh: 29.3}, "for-age-55", 29.793, ""),
		newCase_(female, "2037-Mar-15", map[int]float64{skf.SKFSubscapular: 35.2, skf.SKFTriceps: 13.9, skf.SKFChest: 10.8, skf.SKFMidaxillary: 12.4, skf.SKFSuprailiac: 45.1, skf.SKFAbdominal: 21.8, skf.SKFThigh: 34.2}, "for-age-49", 31.254, ""),
		newCase_(female, "2029-Mar-15", map[int]float64{skf.SKFSubscapular: 35.4, skf.SKFTriceps: 16.8, skf.SKFChest: 10.3, skf.SKFMidaxillary: 14.7, skf.SKFSuprailiac: 41.6, skf.SKFAbdominal: 17.7, skf.SKFThigh: 23.6}, "for-age-41", 28.986, ""),
		newCase_(female, "2031-Mar-15", map[int]float64{skf.SKFSubscapular: 20.4, skf.SKFTriceps: 14.8, skf.SKFChest: 9.9, skf.SKFMidaxillary: 14.8, skf.SKFSuprailiac: 27.6, skf.SKFAbdominal: 34.2, skf.SKFThigh: 17.6}, "for-age-43", 26.162, ""),
		newCase_(female, "2015-Mar-15", map[int]float64{skf.SKFSubscapular: 25.4, skf.SKFTriceps: 10.1, skf.SKFChest: 10, skf.SKFMidaxillary: 7.8, skf.SKFSuprailiac: 26.3, skf.SKFAbdominal: 37.1, skf.SKFThigh: 11.6}, "for-age-27", 23.58, ""),
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
	cases := []case_{
		newCase_(female, "2014-Mar-15", map[int]float64{skf.SKFSubscapular: 33.3, skf.SKFTriceps: 16.3, skf.SKFChest: 10.9, skf.SKFMidaxillary: 8.9, skf.SKFSuprailiac: 26.6, skf.SKFAbdominal: 46, skf.SKFThigh: 16.1}, "for-age-26", 22.289, ""),
		newCase_(female, "2017-Mar-15", map[int]float64{skf.SKFSubscapular: 22.5, skf.SKFTriceps: 16.8, skf.SKFChest: 6.4, skf.SKFMidaxillary: 7.6, skf.SKFSuprailiac: 17.3, skf.SKFAbdominal: 40.1, skf.SKFThigh: 29.8}, "for-age-29", 24.083, ""),
		newCase_(female, "2025-Mar-15", map[int]float64{skf.SKFSubscapular: 23.9, skf.SKFTriceps: 16.1, skf.SKFChest: 8.5, skf.SKFMidaxillary: 14.6, skf.SKFSuprailiac: 45.9, skf.SKFAbdominal: 44.8, skf.SKFThigh: 21.1}, "for-age-37", 30.489, ""),
		newCase_(female, "2022-Mar-15", map[int]float64{skf.SKFSubscapular: 25.8, skf.SKFTriceps: 10.6, skf.SKFChest: 7.4, skf.SKFMidaxillary: 14.6, skf.SKFSuprailiac: 28.1, skf.SKFAbdominal: 39.4, skf.SKFThigh: 24.9}, "for-age-34", 24.308, ""),
		newCase_(female, "2027-Mar-15", map[int]float64{skf.SKFSubscapular: 20.6, skf.SKFTriceps: 9.9, skf.SKFChest: 8.6, skf.SKFMidaxillary: 18.1, skf.SKFSuprailiac: 31.6, skf.SKFAbdominal: 45.8, skf.SKFThigh: 28.5}, "for-age-39", 26.67, ""),
		newCase_(female, "2010-Mar-15", map[int]float64{skf.SKFSubscapular: 35.6, skf.SKFTriceps: 13.9, skf.SKFChest: 12.1, skf.SKFMidaxillary: 9.5, skf.SKFSuprailiac: 27.8, skf.SKFAbdominal: 33.1, skf.SKFThigh: 34}, "for-age-22", 27.317, ""),
		newCase_(female, "2037-Mar-15", map[int]float64{skf.SKFSubscapular: 24.1, skf.SKFTriceps: 14.3, skf.SKFChest: 12.3, skf.SKFMidaxillary: 11.7, skf.SKFSuprailiac: 31.5, skf.SKFAbdominal: 16.7, skf.SKFThigh: 28}, "for-age-49", 28.502, ""),
		newCase_(female, "2027-Mar-15", map[int]float64{skf.SKFSubscapular: 15.4, skf.SKFTriceps: 8.9, skf.SKFChest: 10.3, skf.SKFMidaxillary: 17.4, skf.SKFSuprailiac: 40.9, skf.SKFAbdominal: 26.9, skf.SKFThigh: 19.8}, "for-age-39", 26.545, ""),
		newCase_(female, "2037-Mar-15", map[int]float64{skf.SKFSubscapular: 26.7, skf.SKFTriceps: 12.4, skf.SKFChest: 8.8, skf.SKFMidaxillary: 11, skf.SKFSuprailiac: 21.6, skf.SKFAbdominal: 15.9, skf.SKFThigh: 14.8}, "for-age-49", 20.281, ""),
		newCase_(female, "2021-Mar-15", map[int]float64{skf.SKFSubscapular: 32.9, skf.SKFTriceps: 10.9, skf.SKFChest: 8.9, skf.SKFMidaxillary: 6.3, skf.SKFSuprailiac: 24.2, skf.SKFAbdominal: 25.5, skf.SKFThigh: 22.5}, "for-age-33", 22.271, ""),
		newCase_(female, "2030-Mar-15", map[int]float64{skf.SKFSubscapular: 16.1, skf.SKFTriceps: 15.4, skf.SKFChest: 7.6, skf.SKFMidaxillary: 10.8, skf.SKFSuprailiac: 35.6, skf.SKFAbdominal: 31.9, skf.SKFThigh: 10.5}, "for-age-42", 24.138, ""),
		newCase_(female, "2008-Mar-15", map[int]float64{skf.SKFSubscapular: 21.4, skf.SKFTriceps: 8.3, skf.SKFChest: 6.6, skf.SKFMidaxillary: 7, skf.SKFSuprailiac: 37.3, skf.SKFAbdominal: 42.2, skf.SKFThigh: 22}, "for-age-20", 24.685, ""),
		newCase_(female, "2013-Mar-15", map[int]float64{skf.SKFSubscapular: 32, skf.SKFTriceps: 12.2, skf.SKFChest: 11.4, skf.SKFMidaxillary: 15.3, skf.SKFSuprailiac: 31.3, skf.SKFAbdominal: 22.2, skf.SKFThigh: 28.4}, "for-age-25", 26.352, ""),
		newCase_(female, "2041-Mar-15", map[int]float64{skf.SKFSubscapular: 29.8, skf.SKFTriceps: 14, skf.SKFChest: 8.7, skf.SKFMidaxillary: 8.7, skf.SKFSuprailiac: 16.7, skf.SKFAbdominal: 34, skf.SKFThigh: 29.3}, "for-age-53", 24.351, ""),
		newCase_(female, "2023-Mar-15", map[int]float64{skf.SKFSubscapular: 28.7, skf.SKFTriceps: 14.7, skf.SKFChest: 12.8, skf.SKFMidaxillary: 15.6, skf.SKFSuprailiac: 24.4, skf.SKFAbdominal: 25.5, skf.SKFThigh: 17.1}, "for-age-35", 21.929, ""),
		newCase_(female, "2010-Mar-15", map[int]float64{skf.SKFSubscapular: 19.3, skf.SKFTriceps: 16.3, skf.SKFChest: 9.2, skf.SKFMidaxillary: 13.4, skf.SKFSuprailiac: 17.8, skf.SKFAbdominal: 34.2, skf.SKFThigh: 26.5}, "for-age-22", 22.561, ""),
		newCase_(female, "2042-Mar-15", map[int]float64{skf.SKFSubscapular: 17.3, skf.SKFTriceps: 16.1, skf.SKFChest: 8, skf.SKFMidaxillary: 12.8, skf.SKFSuprailiac: 43.9, skf.SKFAbdominal: 23, skf.SKFThigh: 33.4}, "for-age-54", 34.513, ""),
		newCase_(female, "2043-Mar-15", map[int]float64{skf.SKFSubscapular: 18.9, skf.SKFTriceps: 12.6, skf.SKFChest: 6.4, skf.SKFMidaxillary: 9.7, skf.SKFSuprailiac: 21.4, skf.SKFAbdominal: 37.6, skf.SKFThigh: 34.6}, "for-age-55", 27.27, ""),
		newCase_(female, "2014-Mar-15", map[int]float64{skf.SKFSubscapular: 15.8, skf.SKFTriceps: 11.8, skf.SKFChest: 10.9, skf.SKFMidaxillary: 13.6, skf.SKFSuprailiac: 25.2, skf.SKFAbdominal: 35.4, skf.SKFThigh: 28}, "for-age-26", 24.244, ""),
		newCase_(female, "2017-Mar-15", map[int]float64{skf.SKFSubscapular: 31, skf.SKFTriceps: 16.4, skf.SKFChest: 9.8, skf.SKFMidaxillary: 8.7, skf.SKFSuprailiac: 33.3, skf.SKFAbdominal: 28.7, skf.SKFThigh: 14}, "for-age-29", 24.018, ""),
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
	cases := []case_{
		newCase_(female, "2004-Mar-15", map[int]float64{skf.SKFSubscapular: 28, skf.SKFTriceps: 13.1, skf.SKFChest: 7.5, skf.SKFMidaxillary: 7.5, skf.SKFSuprailiac: 30, skf.SKFAbdominal: 18.5, skf.SKFThigh: 31.1, skf.SKFCalf: 14.5}, "for-age-16", 21.286, ""),
		newCase_(female, "2004-Mar-15", map[int]float64{skf.SKFSubscapular: 18.5, skf.SKFTriceps: 9.8, skf.SKFChest: 9.6, skf.SKFMidaxillary: 7.7, skf.SKFSuprailiac: 21.1, skf.SKFAbdominal: 34.1, skf.SKFThigh: 20.8, skf.SKFCalf: 20.6}, "for-age-16", 23.344, ""),
		newCase_(female, "1995-Mar-15", map[int]float64{skf.SKFSubscapular: 27.2, skf.SKFTriceps: 16.1, skf.SKFChest: 11.9, skf.SKFMidaxillary: 6.8, skf.SKFSuprailiac: 36.5, skf.SKFAbdominal: 45.4, skf.SKFThigh: 34.6, skf.SKFCalf: 10.4}, "for-age-7", 20.478, ""),
		newCase_(female, "2001-Mar-15", map[int]float64{skf.SKFSubscapular: 21.7, skf.SKFTriceps: 12.3, skf.SKFChest: 12.8, skf.SKFMidaxillary: 11.9, skf.SKFSuprailiac: 36.9, skf.SKFAbdominal: 32.7, skf.SKFThigh: 27.9, skf.SKFCalf: 17.4}, "for-age-13", 22.83, ""),
		newCase_(female, "1995-Mar-15", map[int]float64{skf.SKFSubscapular: 32.5, skf.SKFTriceps: 13.9, skf.SKFChest: 8.2, skf.SKFMidaxillary: 8.5, skf.SKFSuprailiac: 30.8, skf.SKFAbdominal: 25.1, skf.SKFThigh: 13.7, skf.SKFCalf: 20.1}, "for-age-7", 25.99, ""),
		newCase_(female, "2003-Mar-15", map[int]float64{skf.SKFSubscapular: 21.3, skf.SKFTriceps: 11.5, skf.SKFChest: 8.2, skf.SKFMidaxillary: 12.8, skf.SKFSuprailiac: 31.9, skf.SKFAbdominal: 20.1, skf.SKFThigh: 31.2, skf.SKFCalf: 16.8}, "for-age-15", 21.801, ""),
		newCase_(female, "1995-Mar-15", map[int]float64{skf.SKFSubscapular: 35.4, skf.SKFTriceps: 13.7, skf.SKFChest: 6.5, skf.SKFMidaxillary: 11.4, skf.SKFSuprailiac: 38.6, skf.SKFAbdominal: 29.3, skf.SKFThigh: 27.6, skf.SKFCalf: 20.7}, "for-age-7", 26.284, ""),
		newCase_(female, "2004-Mar-15", map[int]float64{skf.SKFSubscapular: 30.5, skf.SKFTriceps: 10.1, skf.SKFChest: 6.2, skf.SKFMidaxillary: 11.1, skf.SKFSuprailiac: 42, skf.SKFAbdominal: 26.3, skf.SKFThigh: 26.4, skf.SKFCalf: 16.7}, "for-age-16", 20.698, ""),
		newCase_(female, "2000-Mar-15", map[int]float64{skf.SKFSubscapular: 29.7, skf.SKFTriceps: 9.3, skf.SKFChest: 9.7, skf.SKFMidaxillary: 12.6, skf.SKFSuprailiac: 36.7, skf.SKFAbdominal: 17.5, skf.SKFThigh: 28.4, skf.SKFCalf: 9.6}, "for-age-12", 14.892, ""),
		newCase_(female, "2003-Mar-15", map[int]float64{skf.SKFSubscapular: 29.7, skf.SKFTriceps: 12.8, skf.SKFChest: 8.2, skf.SKFMidaxillary: 16.4, skf.SKFSuprailiac: 24.5, skf.SKFAbdominal: 26.7, skf.SKFThigh: 30.1, skf.SKFCalf: 25.6}, "for-age-15", 29.224, ""),
		newCase_(female, "1999-Mar-15", map[int]float64{skf.SKFSubscapular: 16.1, skf.SKFTriceps: 16.1, skf.SKFChest: 11.6, skf.SKFMidaxillary: 11.8, skf.SKFSuprailiac: 42.6, skf.SKFAbdominal: 32.2, skf.SKFThigh: 14.1, skf.SKFCalf: 14.9}, "for-age-11", 23.785, ""),
		newCase_(female, "2002-Mar-15", map[int]float64{skf.SKFSubscapular: 24.1, skf.SKFTriceps: 11.4, skf.SKFChest: 6.3, skf.SKFMidaxillary: 6.7, skf.SKFSuprailiac: 38.9, skf.SKFAbdominal: 18, skf.SKFThigh: 34.6, skf.SKFCalf: 18.4}, "for-age-14", 22.903, ""),
		newCase_(female, "1995-Mar-15", map[int]float64{skf.SKFSubscapular: 34.5, skf.SKFTriceps: 16.1, skf.SKFChest: 8.7, skf.SKFMidaxillary: 11.6, skf.SKFSuprailiac: 45.3, skf.SKFAbdominal: 15.3, skf.SKFThigh: 27.5, skf.SKFCalf: 19.1}, "for-age-7", 26.872, ""),
		newCase_(female, "2001-Mar-15", map[int]float64{skf.SKFSubscapular: 28.1, skf.SKFTriceps: 14.9, skf.SKFChest: 12.6, skf.SKFMidaxillary: 11.1, skf.SKFSuprailiac: 18.5, skf.SKFAbdominal: 44.1, skf.SKFThigh: 26.4, skf.SKFCalf: 21.4}, "for-age-13", 27.681, ""),
		newCase_(female, "2002-Mar-15", map[int]float64{skf.SKFSubscapular: 23.1, skf.SKFTriceps: 8.8, skf.SKFChest: 11, skf.SKFMidaxillary: 13.6, skf.SKFSuprailiac: 20, skf.SKFAbdominal: 32.9, skf.SKFThigh: 20.9, skf.SKFCalf: 8.2}, "for-age-14", 13.495, ""),
		newCase_(female, "2005-Mar-15", map[int]float64{skf.SKFSubscapular: 18.9, skf.SKFTriceps: 15.8, skf.SKFChest: 12.7, skf.SKFMidaxillary: 11.4, skf.SKFSuprailiac: 25.3, skf.SKFAbdominal: 29, skf.SKFThigh: 20.2, skf.SKFCalf: 16.8}, "for-age-17", 24.961, ""),
		newCase_(female, "1999-Mar-15", map[int]float64{skf.SKFSubscapular: 20.9, skf.SKFTriceps: 8, skf.SKFChest: 6.5, skf.SKFMidaxillary: 8.4, skf.SKFSuprailiac: 45.4, skf.SKFAbdominal: 15.6, skf.SKFThigh: 33.5, skf.SKFCalf: 23.1}, "for-age-11", 23.859, ""),
		newCase_(female, "1995-Mar-15", map[int]float64{skf.SKFSubscapular: 30.8, skf.SKFTriceps: 14.6, skf.SKFChest: 7.2, skf.SKFMidaxillary: 15, skf.SKFSuprailiac: 22.8, skf.SKFAbdominal: 43.2, skf.SKFThigh: 22.5, skf.SKFCalf: 8.4}, "for-age-7", 17.905, ""),
		newCase_(female, "1999-Mar-15", map[int]float64{skf.SKFSubscapular: 27.4, skf.SKFTriceps: 8.1, skf.SKFChest: 7.8, skf.SKFMidaxillary: 9.1, skf.SKFSuprailiac: 44.1, skf.SKFAbdominal: 24.7, skf.SKFThigh: 34.1, skf.SKFCalf: 23.9}, "for-age-11", 24.52, ""),
		newCase_(female, "1998-Mar-15", map[int]float64{skf.SKFSubscapular: 27.5, skf.SKFTriceps: 16.5, skf.SKFChest: 7.8, skf.SKFMidaxillary: 14.4, skf.SKFSuprailiac: 33.9, skf.SKFAbdominal: 38.2, skf.SKFThigh: 16.5, skf.SKFCalf: 19.9}, "for-age-10", 27.754, ""),
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
