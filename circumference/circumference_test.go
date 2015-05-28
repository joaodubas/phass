package circumference

import (
	assess "github.com/joaodubas/phass/assessment"
	"math"
	"testing"
)

func TestWaistToHipCalcAndClassification(t *testing.T) {
	male, err := assess.NewPerson("João Paulo Dubas", "1978-Dec-15", assess.Male)
	if err != nil {
		t.Errorf("Could not create a person: .", err)
	}

	female, err := assess.NewPerson("Ana Paula Dubas", "1988-Mar-08", assess.Female)
	if err != nil {
		t.Errorf("Could not create a person: .", err)
	}

	type wthSpec struct {
		person         *assess.Person
		assessmentDate string
		waist          float64
		hip            float64
		calc           float64
		classify       string
	}

	specs := []wthSpec{
		wthSpec{person: male, assessmentDate: "1998-Dec-15", waist: 82.0, hip: 100.6, calc: 0.8151, classify: WTHClassification[WTHLow]},
		wthSpec{person: male, assessmentDate: "1998-Dec-15", waist: 84.4, hip: 98.1, calc: 0.8603, classify: WTHClassification[WTHModerate]},
		wthSpec{person: male, assessmentDate: "1998-Dec-15", waist: 96.6, hip: 105.6, calc: 0.9148, classify: WTHClassification[WTHHigh]},
		wthSpec{person: male, assessmentDate: "1998-Dec-15", waist: 127.1, hip: 104.2, calc: 1.2198, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: male, assessmentDate: "2008-Dec-15", waist: 84.3, hip: 102.8, calc: 0.8200, classify: WTHClassification[WTHLow]},
		wthSpec{person: male, assessmentDate: "2008-Dec-15", waist: 92.7, hip: 105.3, calc: 0.8803, classify: WTHClassification[WTHModerate]},
		wthSpec{person: male, assessmentDate: "2008-Dec-15", waist: 97.5, hip: 103.7, calc: 0.9402, classify: WTHClassification[WTHHigh]},
		wthSpec{person: male, assessmentDate: "2008-Dec-15", waist: 115.5, hip: 93.9, calc: 1.2300, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: male, assessmentDate: "2018-Dec-15", waist: 80.5, hip: 95.8, calc: 0.8403, classify: WTHClassification[WTHLow]},
		wthSpec{person: male, assessmentDate: "2018-Dec-15", waist: 87.0, hip: 94.6, calc: 0.9197, classify: WTHClassification[WTHModerate]},
		wthSpec{person: male, assessmentDate: "2018-Dec-15", waist: 89.8, hip: 91.6, calc: 0.9803, classify: WTHClassification[WTHHigh]},
		wthSpec{person: male, assessmentDate: "2018-Dec-15", waist: 137.9, hip: 110.3, calc: 1.2502, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: male, assessmentDate: "2028-Dec-15", waist: 78.5, hip: 92.4, calc: 0.8496, classify: WTHClassification[WTHLow]},
		wthSpec{person: male, assessmentDate: "2028-Dec-15", waist: 101.2, hip: 108.2, calc: 0.9353, classify: WTHClassification[WTHModerate]},
		wthSpec{person: male, assessmentDate: "2028-Dec-15", waist: 93.1, hip: 93.6, calc: 0.9947, classify: WTHClassification[WTHHigh]},
		wthSpec{person: male, assessmentDate: "2028-Dec-15", waist: 113.5, hip: 90.1, calc: 1.2597, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: male, assessmentDate: "2038-Dec-15", waist: 83.2, hip: 97.3, calc: 0.8551, classify: WTHClassification[WTHLow]},
		wthSpec{person: male, assessmentDate: "2038-Dec-15", waist: 100.1, hip: 105.4, calc: 0.9497, classify: WTHClassification[WTHModerate]},
		wthSpec{person: male, assessmentDate: "2038-Dec-15", waist: 107.0, hip: 105.9, calc: 1.0104, classify: WTHClassification[WTHHigh]},
		wthSpec{person: male, assessmentDate: "2038-Dec-15", waist: 118.4, hip: 93.6, calc: 1.2650, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: female, assessmentDate: "2008-Mar-15", waist: 64.1, hip: 97.8, calc: 0.6554, classify: WTHClassification[WTHLow]},
		wthSpec{person: female, assessmentDate: "2008-Mar-15", waist: 69.7, hip: 93.6, calc: 0.7447, classify: WTHClassification[WTHModerate]},
		wthSpec{person: female, assessmentDate: "2008-Mar-15", waist: 79.5, hip: 99.4, calc: 0.7998, classify: WTHClassification[WTHHigh]},
		wthSpec{person: female, assessmentDate: "2008-Mar-15", waist: 84.3, hip: 92.6, calc: 0.9104, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: female, assessmentDate: "2018-Mar-15", waist: 65.3, hip: 99.0, calc: 0.6596, classify: WTHClassification[WTHLow]},
		wthSpec{person: female, assessmentDate: "2018-Mar-15", waist: 76.6, hip: 101.5, calc: 0.7547, classify: WTHClassification[WTHModerate]},
		wthSpec{person: female, assessmentDate: "2018-Mar-15", waist: 82.0, hip: 100.6, calc: 0.8151, classify: WTHClassification[WTHHigh]},
		wthSpec{person: female, assessmentDate: "2018-Mar-15", waist: 88.9, hip: 96.6, calc: 0.9203, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: female, assessmentDate: "2028-Mar-15", waist: 64.0, hip: 96.3, calc: 0.6646, classify: WTHClassification[WTHLow]},
		wthSpec{person: female, assessmentDate: "2028-Mar-15", waist: 80.4, hip: 105.1, calc: 0.7650, classify: WTHClassification[WTHModerate]},
		wthSpec{person: female, assessmentDate: "2028-Mar-15", waist: 81.7, hip: 97.9, calc: 0.8345, classify: WTHClassification[WTHHigh]},
		wthSpec{person: female, assessmentDate: "2028-Mar-15", waist: 89.9, hip: 96.1, calc: 0.9355, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: female, assessmentDate: "2038-Mar-15", waist: 66.7, hip: 99.5, calc: 0.6704, classify: WTHClassification[WTHLow]},
		wthSpec{person: female, assessmentDate: "2038-Mar-15", waist: 71.6, hip: 91.8, calc: 0.7800, classify: WTHClassification[WTHModerate]},
		wthSpec{person: female, assessmentDate: "2038-Mar-15", waist: 91.1, hip: 107.2, calc: 0.8498, classify: WTHClassification[WTHHigh]},
		wthSpec{person: female, assessmentDate: "2038-Mar-15", waist: 100.6, hip: 107.0, calc: 0.9402, classify: WTHClassification[WTHVeryHigh]},
		wthSpec{person: female, assessmentDate: "2048-Mar-15", waist: 62.1, hip: 91.3, calc: 0.6802, classify: WTHClassification[WTHLow]},
		wthSpec{person: female, assessmentDate: "2048-Mar-15", waist: 74.6, hip: 93.3, calc: 0.7996, classify: WTHClassification[WTHModerate]},
		wthSpec{person: female, assessmentDate: "2048-Mar-15", waist: 88.6, hip: 101.8, calc: 0.8703, classify: WTHClassification[WTHHigh]},
		wthSpec{person: female, assessmentDate: "2048-Mar-15", waist: 93.2, hip: 98.1, calc: 0.9501, classify: WTHClassification[WTHVeryHigh]},
	}

	type wthCase struct {
		wth      *WaistToHip
		calc     float64
		classify string
	}

	cases := []wthCase{}

	for _, spec := range specs {
		assessment, err := assess.NewAssessment(spec.assessmentDate)
		if err != nil {
			t.Errorf("Could not create assessment: %s", err)
		}
		cases = append(cases, wthCase{
			wth:      NewWaistToHipRatio(spec.person, assessment, map[int]float64{CCFWaist: spec.waist, CCFHip: spec.hip}),
			calc:     spec.calc,
			classify: spec.classify,
		})
	}

	for _, data := range cases {
		if calc, _ := data.wth.Calc(); !floatEqual(calc, data.calc, wthLimit) {
			t.Errorf("Calc is %.3f, expected is %.3f", calc, data.calc)
		}

		if classify, _ := data.wth.Classify(); classify != data.classify {
			t.Errorf("Classify is %s, expected is %s", classify, data.classify)
		}
	}
}

func TestWaistToHipMissingMeasure(t *testing.T) {
	p, err := assess.NewPerson("João Paulo Dubas", "1978-Dec-15", assess.Male)
	if err != nil {
		t.Errorf("Could not create person: %s", err)
	}

	a, err := assess.NewAssessment("2015-May-22")
	if err != nil {
		t.Errorf("Could not create assessment: %s", err)
	}

	cases := []*WaistToHip{
		// missing waist measure
		NewWaistToHipRatio(p, a, map[int]float64{CCFHip: 101.2}),
		// missing hip measure
		NewWaistToHipRatio(p, a, map[int]float64{CCFWaist: 80.3}),
		// wrong measures
		NewWaistToHipRatio(p, a, map[int]float64{CCFAbdominal: 100.1, CCFChest: 108.1}),
	}

	for _, data := range cases {
		if _, err := data.Calc(); err != nil {
			t.Error("Should not get a waist-to-hip value")
		}
		if _, err := data.Classify(); err != nil {
			t.Errorf("Should not get a waist-to-hip classification")
		}
	}
}

func TestWaistToHipOutsideAgeRange(t *testing.T) {
	m, err := assess.NewPerson("João Paulo Dubas", "1978-Dec-15", assess.Male)
	if err != nil {
		t.Errorf("Could not create person: %s", err)
	}

	f, err := assess.NewPerson("Ana Paula Dubas", "1988-Mar-15", assess.Female)
	if err != nil {
		t.Errorf("Could not create person: %s", err)
	}

	type wthSpec struct {
		person     *assess.Person
		assessment string
		measures   map[int]float64
	}
	specs := []wthSpec{
		// lower than 20 years
		wthSpec{person: m, assessment: "1998-Dec-14", measures: map[int]float64{CCFWaist: 100.1, CCFHip: 100.1}},
		wthSpec{person: f, assessment: "2008-Mar-14", measures: map[int]float64{CCFWaist: 100.1, CCFHip: 100.1}},
		// greater than 69 years
		wthSpec{person: m, assessment: "2048-Dec-15", measures: map[int]float64{CCFWaist: 100.1, CCFHip: 100.1}},
		wthSpec{person: m, assessment: "2058-Mar-15", measures: map[int]float64{CCFWaist: 100.1, CCFHip: 100.1}},
	}

	cases := []*WaistToHip{}
	for _, spec := range specs {
		a, err := assess.NewAssessment(spec.assessment)
		if err != nil {
			t.Errorf("Could not create assessment: %s", err)
		}
		cases = append(cases, NewWaistToHipRatio(spec.person, a, spec.measures))
	}

	for _, data := range cases {
		if r, err := data.Calc(); err == nil {
			t.Errorf("WTH should not have a value due to age outside valid range, %.4f", r)
		}
		if _, err := data.Classify(); err == nil {
			t.Error("WTH do not have classification due to age outside valid range")
		}
	}
}

func floatEqual(original, expected, limit float64) bool {
	diff := math.Abs(original - expected)
	return diff <= limit
}

var wthLimit = 0.001
