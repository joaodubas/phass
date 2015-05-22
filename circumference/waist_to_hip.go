package circumference

import (
	"fmt"
	assess "github.com/joaodubas/phass/assessment"
	"github.com/joaodubas/phass/common"
	"math"
)

type WaistToHip struct {
	*assess.Person
	*Circumferences
}

func (w *WaistToHip) String() string {
	return fmt.Sprintf("%s\nWTH: %.2f (%s)", w.Person.String(), w.Calc(), w.Classify())
}

func (w *WaistToHip) Classify() string {
	r := w.Calc()
	if r == 0.0 {
		return "No classification due to missing measure"
	}

	genderClass, ok := wthLimits[w.Gender]
	if !ok {
		return fmt.Sprintf("No classification for gender %d", w.Gender)
	}

	age := w.Age()
	for limits, classes := range genderClass {
		if age < limits[0] || age > limits[1] {
			continue
		}
		return common.Classifier(w.Calc(), classes, WTHClassification)
	}

	return fmt.Sprintf("No classification for age %.0f", age)
}

func (w *WaistToHip) Calc() float64 {
	waist, ok := w.Measures[CCFWaist]
	if !ok {
		return 0.0
	}
	hip, ok := w.Measures[CCFHip]
	if !ok {
		return 0.0
	}
	return waist / hip
}

var wthLimits = map[int]map[[2]float64]map[int][2]float64{
	assess.Male: wthLimitsForMales,
	assess.Female: wthLimitsForFemale,
}

var wthLimitsForMales = map[[2]float64]map[int][2]float64{
	[2]float64{20, 30}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.83},
		WTHModerate: [2]float64{0.83, 0.89},
		WTHHigh:     [2]float64{0.89, 0.94},
		WTHVeryHigh: [2]float64{0.94, math.Inf(+1)},
	},
	[2]float64{30, 40}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.84},
		WTHModerate: [2]float64{0.84, 0.92},
		WTHHigh:     [2]float64{0.92, 0.96},
		WTHVeryHigh: [2]float64{0.96, math.Inf(+1)},
	},
	[2]float64{40, 50}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.88},
		WTHModerate: [2]float64{0.88, 0.96},
		WTHHigh:     [2]float64{0.96, 1.00},
		WTHVeryHigh: [2]float64{1.00, math.Inf(+1)},
	},
	[2]float64{50, 60}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.90},
		WTHModerate: [2]float64{0.90, 0.97},
		WTHHigh:     [2]float64{0.97, 1.02},
		WTHVeryHigh: [2]float64{1.02, math.Inf(+1)},
	},
	[2]float64{60, 70}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.91},
		WTHModerate: [2]float64{0.91, 0.99},
		WTHHigh:     [2]float64{0.99, 1.03},
		WTHVeryHigh: [2]float64{1.03, math.Inf(+1)},
	},
}

var wthLimitsForFemale = map[[2]float64]map[int][2]float64{
	[2]float64{20, 30}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.71},
		WTHModerate: [2]float64{0.71, 0.78},
		WTHHigh:     [2]float64{0.78, 0.82},
		WTHVeryHigh: [2]float64{0.82, math.Inf(+1)},
	},
	[2]float64{30, 40}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.72},
		WTHModerate: [2]float64{0.72, 0.79},
		WTHHigh:     [2]float64{0.79, 0.84},
		WTHVeryHigh: [2]float64{0.84, math.Inf(+1)},
	},
	[2]float64{40, 50}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.73},
		WTHModerate: [2]float64{0.73, 0.80},
		WTHHigh:     [2]float64{0.80, 0.87},
		WTHVeryHigh: [2]float64{0.87, math.Inf(+1)},
	},
	[2]float64{50, 60}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.74},
		WTHModerate: [2]float64{0.74, 0.82},
		WTHHigh:     [2]float64{0.82, 0.88},
		WTHVeryHigh: [2]float64{0.88, math.Inf(+1)},
	},
	[2]float64{60, 70}: map[int][2]float64{
		WTHLow:      [2]float64{math.Inf(-1), 0.76},
		WTHModerate: [2]float64{0.76, 0.84},
		WTHHigh:     [2]float64{0.84, 0.90},
		WTHVeryHigh: [2]float64{0.90, math.Inf(+1)},
	},
}
