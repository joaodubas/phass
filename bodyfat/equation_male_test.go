package bodyfat

import (
	skf "github.com/joaodubas/phass/skinfold"
	"testing"
)

func TestMaleSevenSkinfoldEquation(t *testing.T) {
	cases := []case_{
		newCase_(male, "2014-Dec-15", map[int]float64{skf.SKFSubscapular: 18.4, skf.SKFTriceps: 15.2, skf.SKFChest: 11.3, skf.SKFMidaxillary: 15.7, skf.SKFSuprailiac: 39.8, skf.SKFAbdominal: 29.1, skf.SKFThigh: 24.4}, "for-age-36", 22.46, ""),
		newCase_(male, "2030-Dec-15", map[int]float64{skf.SKFSubscapular: 25.6, skf.SKFTriceps: 16.2, skf.SKFChest: 10.8, skf.SKFMidaxillary: 7.6, skf.SKFSuprailiac: 29.1, skf.SKFAbdominal: 23.2, skf.SKFThigh: 15.2}, "for-age-52", 21.234, ""),
		newCase_(male, "2031-Dec-15", map[int]float64{skf.SKFSubscapular: 24.4, skf.SKFTriceps: 11.4, skf.SKFChest: 11, skf.SKFMidaxillary: 13.9, skf.SKFSuprailiac: 34.2, skf.SKFAbdominal: 44.5, skf.SKFThigh: 10.9}, "for-age-53", 24.242, ""),
		newCase_(male, "2005-Dec-15", map[int]float64{skf.SKFSubscapular: 27.8, skf.SKFTriceps: 14.2, skf.SKFChest: 10.7, skf.SKFMidaxillary: 15.1, skf.SKFSuprailiac: 22.6, skf.SKFAbdominal: 35.8, skf.SKFThigh: 27.8}, "for-age-27", 21.306, ""),
		newCase_(male, "2014-Dec-15", map[int]float64{skf.SKFSubscapular: 22.8, skf.SKFTriceps: 14.6, skf.SKFChest: 6.7, skf.SKFMidaxillary: 13.1, skf.SKFSuprailiac: 15.5, skf.SKFAbdominal: 27.1, skf.SKFThigh: 33.8}, "for-age-36", 19.94, ""),
		newCase_(male, "2009-Dec-15", map[int]float64{skf.SKFSubscapular: 16.4, skf.SKFTriceps: 13.3, skf.SKFChest: 7.5, skf.SKFMidaxillary: 15.2, skf.SKFSuprailiac: 44.9, skf.SKFAbdominal: 29.9, skf.SKFThigh: 15}, "for-age-31", 20.384, ""),
		newCase_(male, "2014-Dec-15", map[int]float64{skf.SKFSubscapular: 21.2, skf.SKFTriceps: 10.5, skf.SKFChest: 11.4, skf.SKFMidaxillary: 9.8, skf.SKFSuprailiac: 45, skf.SKFAbdominal: 40.7, skf.SKFThigh: 31.2}, "for-age-36", 24.31, ""),
		newCase_(male, "2031-Dec-15", map[int]float64{skf.SKFSubscapular: 29.8, skf.SKFTriceps: 15.4, skf.SKFChest: 6.8, skf.SKFMidaxillary: 16.2, skf.SKFSuprailiac: 43.8, skf.SKFAbdominal: 31.4, skf.SKFThigh: 25.2}, "for-age-53", 26.41, ""),
		newCase_(male, "2005-Dec-15", map[int]float64{skf.SKFSubscapular: 26.5, skf.SKFTriceps: 12.8, skf.SKFChest: 10.7, skf.SKFMidaxillary: 13.4, skf.SKFSuprailiac: 38.1, skf.SKFAbdominal: 41.6, skf.SKFThigh: 24.1}, "for-age-27", 22.841, ""),
		newCase_(male, "2028-Dec-15", map[int]float64{skf.SKFSubscapular: 32.4, skf.SKFTriceps: 12.3, skf.SKFChest: 10.6, skf.SKFMidaxillary: 16.1, skf.SKFSuprailiac: 37.2, skf.SKFAbdominal: 28.8, skf.SKFThigh: 32.3}, "for-age-50", 26.14, ""),
		newCase_(male, "2003-Dec-15", map[int]float64{skf.SKFSubscapular: 22.5, skf.SKFTriceps: 11.6, skf.SKFChest: 7.3, skf.SKFMidaxillary: 7.1, skf.SKFSuprailiac: 34.1, skf.SKFAbdominal: 44.1, skf.SKFThigh: 25}, "for-age-25", 20.772, ""),
		newCase_(male, "2022-Dec-15", map[int]float64{skf.SKFSubscapular: 27.9, skf.SKFTriceps: 14.4, skf.SKFChest: 11.2, skf.SKFMidaxillary: 8.7, skf.SKFSuprailiac: 19.1, skf.SKFAbdominal: 23.8, skf.SKFThigh: 12.5}, "for-age-44", 18.852, ""),
		newCase_(male, "2028-Dec-15", map[int]float64{skf.SKFSubscapular: 26.2, skf.SKFTriceps: 14.8, skf.SKFChest: 8.5, skf.SKFMidaxillary: 6.8, skf.SKFSuprailiac: 45.2, skf.SKFAbdominal: 19.6, skf.SKFThigh: 25}, "for-age-50", 23.332, ""),
		newCase_(male, "2023-Dec-15", map[int]float64{skf.SKFSubscapular: 27.3, skf.SKFTriceps: 11.7, skf.SKFChest: 7.8, skf.SKFMidaxillary: 15.7, skf.SKFSuprailiac: 44.1, skf.SKFAbdominal: 29.2, skf.SKFThigh: 29.6}, "for-age-45", 24.989, ""),
		newCase_(male, "2021-Dec-15", map[int]float64{skf.SKFSubscapular: 17.5, skf.SKFTriceps: 14.7, skf.SKFChest: 7.3, skf.SKFMidaxillary: 16.5, skf.SKFSuprailiac: 19.2, skf.SKFAbdominal: 45.4, skf.SKFThigh: 31.1}, "for-age-43", 23.106, ""),
		newCase_(male, "2027-Dec-15", map[int]float64{skf.SKFSubscapular: 16, skf.SKFTriceps: 15.1, skf.SKFChest: 7.2, skf.SKFMidaxillary: 14, skf.SKFSuprailiac: 22.4, skf.SKFAbdominal: 23.4, skf.SKFThigh: 12.7}, "for-age-49", 18.558, ""),
		newCase_(male, "2020-Dec-15", map[int]float64{skf.SKFSubscapular: 15.6, skf.SKFTriceps: 10.8, skf.SKFChest: 12.1, skf.SKFMidaxillary: 14.5, skf.SKFSuprailiac: 20.1, skf.SKFAbdominal: 31.8, skf.SKFThigh: 18.1}, "for-age-42", 19.322, ""),
		newCase_(male, "2018-Dec-15", map[int]float64{skf.SKFSubscapular: 29.6, skf.SKFTriceps: 13.7, skf.SKFChest: 12.8, skf.SKFMidaxillary: 12.7, skf.SKFSuprailiac: 31.6, skf.SKFAbdominal: 21.4, skf.SKFThigh: 34.9}, "for-age-40", 23.315, ""),
		newCase_(male, "2000-Dec-15", map[int]float64{skf.SKFSubscapular: 29.9, skf.SKFTriceps: 8.8, skf.SKFChest: 10, skf.SKFMidaxillary: 18.7, skf.SKFSuprailiac: 37.5, skf.SKFAbdominal: 31.2, skf.SKFThigh: 11.7}, "for-age-22", 19.915, ""),
		newCase_(male, "2020-Dec-15", map[int]float64{skf.SKFSubscapular: 27.5, skf.SKFTriceps: 13.4, skf.SKFChest: 8, skf.SKFMidaxillary: 8, skf.SKFSuprailiac: 45.8, skf.SKFAbdominal: 20.1, skf.SKFThigh: 18.9}, "for-age-42", 21.743, ""),
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
	cases := []case_{
		newCase_(male, "2019-Dec-15", map[int]float64{skf.SKFSubscapular: 23.8, skf.SKFTriceps: 12, skf.SKFChest: 6, skf.SKFMidaxillary: 11.3, skf.SKFSuprailiac: 37.7, skf.SKFAbdominal: 43.6, skf.SKFThigh: 25.9}, "for-age-41", 23.444, ""),
		newCase_(male, "2002-Dec-15", map[int]float64{skf.SKFSubscapular: 32.7, skf.SKFTriceps: 14.2, skf.SKFChest: 10.8, skf.SKFMidaxillary: 16.8, skf.SKFSuprailiac: 15.1, skf.SKFAbdominal: 20.6, skf.SKFThigh: 15}, "for-age-24", 13.358, ""),
		newCase_(male, "2029-Dec-15", map[int]float64{skf.SKFSubscapular: 23.5, skf.SKFTriceps: 13.2, skf.SKFChest: 10.3, skf.SKFMidaxillary: 16.6, skf.SKFSuprailiac: 20.2, skf.SKFAbdominal: 24.8, skf.SKFThigh: 30.9}, "for-age-51", 22.031, ""),
		newCase_(male, "2016-Dec-15", map[int]float64{skf.SKFSubscapular: 32.6, skf.SKFTriceps: 9.7, skf.SKFChest: 9.8, skf.SKFMidaxillary: 15.3, skf.SKFSuprailiac: 25.6, skf.SKFAbdominal: 26.5, skf.SKFThigh: 19.4}, "for-age-38", 17.636, ""),
		newCase_(male, "2009-Dec-15", map[int]float64{skf.SKFSubscapular: 23, skf.SKFTriceps: 13.9, skf.SKFChest: 10.4, skf.SKFMidaxillary: 14.7, skf.SKFSuprailiac: 31.3, skf.SKFAbdominal: 44.4, skf.SKFThigh: 34.6}, "for-age-31", 25.833, ""),
		newCase_(male, "2019-Dec-15", map[int]float64{skf.SKFSubscapular: 27, skf.SKFTriceps: 13.6, skf.SKFChest: 6.6, skf.SKFMidaxillary: 15, skf.SKFSuprailiac: 31.7, skf.SKFAbdominal: 17.1, skf.SKFThigh: 32.4}, "for-age-41", 18.092, ""),
		newCase_(male, "2012-Dec-15", map[int]float64{skf.SKFSubscapular: 17, skf.SKFTriceps: 9, skf.SKFChest: 6.3, skf.SKFMidaxillary: 6.1, skf.SKFSuprailiac: 19.1, skf.SKFAbdominal: 45.3, skf.SKFThigh: 12.8}, "for-age-34", 19.628, ""),
		newCase_(male, "2017-Dec-15", map[int]float64{skf.SKFSubscapular: 25.5, skf.SKFTriceps: 13.1, skf.SKFChest: 8.7, skf.SKFMidaxillary: 17.1, skf.SKFSuprailiac: 21.6, skf.SKFAbdominal: 33.8, skf.SKFThigh: 22.7}, "for-age-39", 20.424, ""),
		newCase_(male, "2021-Dec-15", map[int]float64{skf.SKFSubscapular: 28, skf.SKFTriceps: 15.9, skf.SKFChest: 10.4, skf.SKFMidaxillary: 15.8, skf.SKFSuprailiac: 45.8, skf.SKFAbdominal: 39.8, skf.SKFThigh: 12.9}, "for-age-43", 20.301, ""),
		newCase_(male, "2020-Dec-15", map[int]float64{skf.SKFSubscapular: 33.1, skf.SKFTriceps: 9.5, skf.SKFChest: 12.1, skf.SKFMidaxillary: 18.4, skf.SKFSuprailiac: 23, skf.SKFAbdominal: 30.4, skf.SKFThigh: 29.5}, "for-age-42", 22.625, ""),
		newCase_(male, "2026-Dec-15", map[int]float64{skf.SKFSubscapular: 16.9, skf.SKFTriceps: 15.7, skf.SKFChest: 8.1, skf.SKFMidaxillary: 16.9, skf.SKFSuprailiac: 28, skf.SKFAbdominal: 24.3, skf.SKFThigh: 13.7}, "for-age-48", 15.964, ""),
		newCase_(male, "2025-Dec-15", map[int]float64{skf.SKFSubscapular: 18.3, skf.SKFTriceps: 13.4, skf.SKFChest: 10.9, skf.SKFMidaxillary: 10.9, skf.SKFSuprailiac: 19.3, skf.SKFAbdominal: 16, skf.SKFThigh: 25.2}, "for-age-47", 17.619, ""),
		newCase_(male, "2015-Dec-15", map[int]float64{skf.SKFSubscapular: 27.6, skf.SKFTriceps: 11.1, skf.SKFChest: 7.8, skf.SKFMidaxillary: 9.1, skf.SKFSuprailiac: 33.3, skf.SKFAbdominal: 38, skf.SKFThigh: 29.9}, "for-age-37", 23.031, ""),
		newCase_(male, "2017-Dec-15", map[int]float64{skf.SKFSubscapular: 18, skf.SKFTriceps: 9.9, skf.SKFChest: 7.7, skf.SKFMidaxillary: 14.6, skf.SKFSuprailiac: 39.4, skf.SKFAbdominal: 45.1, skf.SKFThigh: 32.2}, "for-age-39", 25.673, ""),
		newCase_(male, "2029-Dec-15", map[int]float64{skf.SKFSubscapular: 35.3, skf.SKFTriceps: 12.8, skf.SKFChest: 6.8, skf.SKFMidaxillary: 11.6, skf.SKFSuprailiac: 41.2, skf.SKFAbdominal: 39.7, skf.SKFThigh: 25.4}, "for-age-51", 23.646, ""),
		newCase_(male, "1997-Dec-15", map[int]float64{skf.SKFSubscapular: 29.3, skf.SKFTriceps: 10.5, skf.SKFChest: 7.2, skf.SKFMidaxillary: 14.8, skf.SKFSuprailiac: 19.2, skf.SKFAbdominal: 28.7, skf.SKFThigh: 10.7}, "for-age-19", 12.859, ""),
		newCase_(male, "2013-Dec-15", map[int]float64{skf.SKFSubscapular: 19.8, skf.SKFTriceps: 11, skf.SKFChest: 7.5, skf.SKFMidaxillary: 6.7, skf.SKFSuprailiac: 39.1, skf.SKFAbdominal: 44.3, skf.SKFThigh: 29.9}, "for-age-35", 24.361, ""),
		newCase_(male, "2015-Dec-15", map[int]float64{skf.SKFSubscapular: 28.9, skf.SKFTriceps: 12.5, skf.SKFChest: 11.8, skf.SKFMidaxillary: 10.3, skf.SKFSuprailiac: 41.9, skf.SKFAbdominal: 17.5, skf.SKFThigh: 22.5}, "for-age-37", 16.398, ""),
		newCase_(male, "2016-Dec-15", map[int]float64{skf.SKFSubscapular: 30.1, skf.SKFTriceps: 12.8, skf.SKFChest: 6.6, skf.SKFMidaxillary: 17.6, skf.SKFSuprailiac: 28.5, skf.SKFAbdominal: 39.9, skf.SKFThigh: 24}, "for-age-38", 21.757, ""),
		newCase_(male, "2005-Dec-15", map[int]float64{skf.SKFSubscapular: 30, skf.SKFTriceps: 16.1, skf.SKFChest: 6.3, skf.SKFMidaxillary: 14.4, skf.SKFSuprailiac: 29.4, skf.SKFAbdominal: 42, skf.SKFThigh: 19.5}, "for-age-27", 19.758, ""),
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
	cases := []case_{
		newCase_(male, "1994-Dec-15", map[int]float64{skf.SKFSubscapular: 17.6, skf.SKFTriceps: 8.5, skf.SKFChest: 9.8, skf.SKFMidaxillary: 10.6, skf.SKFSuprailiac: 40.3, skf.SKFAbdominal: 28.4, skf.SKFThigh: 34.6, skf.SKFCalf: 18.4}, "for-age-16", 21.509, ""),
		newCase_(male, "1984-Dec-15", map[int]float64{skf.SKFSubscapular: 26.6, skf.SKFTriceps: 12, skf.SKFChest: 9.7, skf.SKFMidaxillary: 16.7, skf.SKFSuprailiac: 43.9, skf.SKFAbdominal: 42.8, skf.SKFThigh: 10.3, skf.SKFCalf: 24.1}, "for-age-6", 27.121, ""),
		newCase_(male, "1994-Dec-15", map[int]float64{skf.SKFSubscapular: 19.7, skf.SKFTriceps: 16.1, skf.SKFChest: 8.8, skf.SKFMidaxillary: 15.2, skf.SKFSuprailiac: 17, skf.SKFAbdominal: 35.3, skf.SKFThigh: 28.7, skf.SKFCalf: 24.4}, "for-age-16", 29.805, ""),
		newCase_(male, "1989-Dec-15", map[int]float64{skf.SKFSubscapular: 20.5, skf.SKFTriceps: 10.3, skf.SKFChest: 11.2, skf.SKFMidaxillary: 16.3, skf.SKFSuprailiac: 25, skf.SKFAbdominal: 39.4, skf.SKFThigh: 11.9, skf.SKFCalf: 23.7}, "for-age-11", 25.84, ""),
		newCase_(male, "1989-Dec-15", map[int]float64{skf.SKFSubscapular: 21.2, skf.SKFTriceps: 11.3, skf.SKFChest: 10.8, skf.SKFMidaxillary: 18.7, skf.SKFSuprailiac: 27.4, skf.SKFAbdominal: 18.1, skf.SKFThigh: 25.8, skf.SKFCalf: 22.8}, "for-age-11", 25.901, ""),
		newCase_(male, "1985-Dec-15", map[int]float64{skf.SKFSubscapular: 31.7, skf.SKFTriceps: 9.4, skf.SKFChest: 12.3, skf.SKFMidaxillary: 10.4, skf.SKFSuprailiac: 19.9, skf.SKFAbdominal: 38.2, skf.SKFThigh: 26.9, skf.SKFCalf: 19.9}, "for-age-7", 22.973, ""),
		newCase_(male, "1990-Dec-15", map[int]float64{skf.SKFSubscapular: 18.9, skf.SKFTriceps: 15.9, skf.SKFChest: 8, skf.SKFMidaxillary: 7.6, skf.SKFSuprailiac: 33.2, skf.SKFAbdominal: 26.4, skf.SKFThigh: 23.5, skf.SKFCalf: 16.7}, "for-age-12", 24.986, ""),
		newCase_(male, "1985-Dec-15", map[int]float64{skf.SKFSubscapular: 20.3, skf.SKFTriceps: 11.4, skf.SKFChest: 6.5, skf.SKFMidaxillary: 15.9, skf.SKFSuprailiac: 33.4, skf.SKFAbdominal: 15.7, skf.SKFThigh: 18.3, skf.SKFCalf: 23.1}, "for-age-7", 26.145, ""),
		newCase_(male, "1992-Dec-15", map[int]float64{skf.SKFSubscapular: 31.5, skf.SKFTriceps: 16.9, skf.SKFChest: 7, skf.SKFMidaxillary: 13.2, skf.SKFSuprailiac: 31.4, skf.SKFAbdominal: 41, skf.SKFThigh: 18.3, skf.SKFCalf: 17.2}, "for-age-14", 25.901, ""),
		newCase_(male, "1990-Dec-15", map[int]float64{skf.SKFSubscapular: 29.4, skf.SKFTriceps: 13.7, skf.SKFChest: 12.8, skf.SKFMidaxillary: 8.3, skf.SKFSuprailiac: 45.9, skf.SKFAbdominal: 23.5, skf.SKFThigh: 21.4, skf.SKFCalf: 14.9}, "for-age-12", 22.546, ""),
		newCase_(male, "1989-Dec-15", map[int]float64{skf.SKFSubscapular: 25.5, skf.SKFTriceps: 13.7, skf.SKFChest: 6.9, skf.SKFMidaxillary: 12.1, skf.SKFSuprailiac: 31.3, skf.SKFAbdominal: 40, skf.SKFThigh: 29.6, skf.SKFCalf: 23.3}, "for-age-11", 27.67, ""),
		newCase_(male, "1994-Dec-15", map[int]float64{skf.SKFSubscapular: 19.4, skf.SKFTriceps: 13.7, skf.SKFChest: 11.4, skf.SKFMidaxillary: 13.4, skf.SKFSuprailiac: 42, skf.SKFAbdominal: 34.5, skf.SKFThigh: 30.3, skf.SKFCalf: 25.5}, "for-age-16", 29.012, ""),
		newCase_(male, "1987-Dec-15", map[int]float64{skf.SKFSubscapular: 29.7, skf.SKFTriceps: 9.8, skf.SKFChest: 7.4, skf.SKFMidaxillary: 9.4, skf.SKFSuprailiac: 36.3, skf.SKFAbdominal: 38.6, skf.SKFThigh: 16, skf.SKFCalf: 24.1}, "for-age-9", 25.779, ""),
		newCase_(male, "1986-Dec-15", map[int]float64{skf.SKFSubscapular: 32.5, skf.SKFTriceps: 9.7, skf.SKFChest: 12.2, skf.SKFMidaxillary: 14.4, skf.SKFSuprailiac: 37, skf.SKFAbdominal: 44.3, skf.SKFThigh: 19.5, skf.SKFCalf: 11.5}, "for-age-8", 18.032, ""),
		newCase_(male, "1992-Dec-15", map[int]float64{skf.SKFSubscapular: 19.8, skf.SKFTriceps: 16, skf.SKFChest: 12, skf.SKFMidaxillary: 16.1, skf.SKFSuprailiac: 17, skf.SKFAbdominal: 25.2, skf.SKFThigh: 23.2, skf.SKFCalf: 6.4}, "for-age-14", 18.764, ""),
		newCase_(male, "1989-Dec-15", map[int]float64{skf.SKFSubscapular: 27.2, skf.SKFTriceps: 11.7, skf.SKFChest: 7.5, skf.SKFMidaxillary: 15.9, skf.SKFSuprailiac: 29, skf.SKFAbdominal: 17.5, skf.SKFThigh: 10, skf.SKFCalf: 24.4}, "for-age-11", 27.121, ""),
		newCase_(male, "1993-Dec-15", map[int]float64{skf.SKFSubscapular: 24.4, skf.SKFTriceps: 10.3, skf.SKFChest: 12.5, skf.SKFMidaxillary: 9.3, skf.SKFSuprailiac: 19.6, skf.SKFAbdominal: 29.5, skf.SKFThigh: 20.1, skf.SKFCalf: 23.4}, "for-age-15", 25.657, ""),
		newCase_(male, "1985-Dec-15", map[int]float64{skf.SKFSubscapular: 29.1, skf.SKFTriceps: 11.2, skf.SKFChest: 12.4, skf.SKFMidaxillary: 12, skf.SKFSuprailiac: 34.8, skf.SKFAbdominal: 18.1, skf.SKFThigh: 11.1, skf.SKFCalf: 14.8}, "for-age-7", 20.96, ""),
		newCase_(male, "1985-Dec-15", map[int]float64{skf.SKFSubscapular: 16.8, skf.SKFTriceps: 16.6, skf.SKFChest: 10.8, skf.SKFMidaxillary: 7.6, skf.SKFSuprailiac: 21.3, skf.SKFAbdominal: 45.4, skf.SKFThigh: 31.3, skf.SKFCalf: 14.5}, "for-age-7", 24.071, ""),
		newCase_(male, "1985-Dec-15", map[int]float64{skf.SKFSubscapular: 22.1, skf.SKFTriceps: 16.4, skf.SKFChest: 10.2, skf.SKFMidaxillary: 9.8, skf.SKFSuprailiac: 35.9, skf.SKFAbdominal: 15.9, skf.SKFThigh: 17.8, skf.SKFCalf: 17.7}, "for-age-7", 25.901, ""),
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
