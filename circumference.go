package circumference

import (
	"fmt"
	"math"

	"github.com/joaodubas/phass"
)

/**
 * Constants
 */

// Circuferences constants.
// CCFNeck: neck circumference
// CCFShoulder: sholder circumference
// CCFChest: chest circumference
// CCFWaist: wait circumference
// CCFAbdominal: abdominal circumference
// CCFHip: hip circumference
// CCFRightArm: right arm circumference
// CCFRightForeArm: right forearm circumference
// CCFRightThigh: right thigh circumference
// CCFRightCalf: right calf circumference
// CCFLeftArm: left arm circumference
// CCFLeftForeArm: left forearm circumference
// CCFLeftThigh: left thigh circumference
// CCFLeftCalf: left calf circumference
const (
	CCFNeck int = iota
	CCFShoulder
	CCFChest
	CCFWaist
	CCFAbdominal
	CCFHip
	CCFRightArm
	CCFRightForeArm
	CCFRightThigh
	CCFRightCalf
	CCFLeftArm
	CCFLeftForeArm
	CCFLeftThigh
	CCFLeftCalf
)

/**
 * Conicity index
 */

// ConicityIndex is an abdominal adiposity proxy, that adjusts waist
// circumference by height and weight.
type ConicityIndex struct {
	*phass.Anthropometry
	*Circumferences
}

// NewConicityIndex creates a new conicity index, based in anthropometry and
// circumferences measures.
func NewConicityIndex(anthropometry *phass.Anthropometry, measures map[int]float64) *ConicityIndex {
	return &ConicityIndex{anthropometry, NewCircumferences(measures)}
}

func (c *ConicityIndex) String() string {
	v, _ := c.Calc()
	return fmt.Sprintf("%s\nConicity index: %.4f", c.Anthropometry.String(), v)
}

// GetName returns this measurement name.
func (c *ConicityIndex) GetName() string {
	return "Conicity index"
}

// Result returns relevant information about this measurement.
func (c *ConicityIndex) Result() ([]string, error) {
	v, err := c.Calc()
	if err != nil {
		return []string{}, err
	}
	return []string{fmt.Sprintf("Conicity index: %.4f.", v)}, nil
}

// Classify returns the classification for this measurement.
func (c *ConicityIndex) Classify() (string, error) {
	if _, err := c.Calc(); err != nil {
		return "", err
	}
	return "No classification available yet", nil
}

// Calc returns the value for this measurement.
func (c *ConicityIndex) Calc() (float64, error) {
	return c.equation().Calc()
}

// equation returns an equation, used to calculate conicity index.
func (c *ConicityIndex) equation() phass.Equationer {
	return phass.NewEquation(cidConf.Extract(c), cidConf)
}

/**
 * Waist-to-hip ratio
 */

// WaistToHip represents the waits-to-hip ratio, a proxy for abdominal
// adiposity.
type WaistToHip struct {
	*phass.Person
	*phass.Assessment
	*Circumferences
}

// NewWaistToHipRatio creates a new pointer to waist-to-hip, based in person,
// assessment and circumference measures.
func NewWaistToHipRatio(person *phass.Person, assessment *phass.Assessment, measures map[int]float64) *WaistToHip {
	return &WaistToHip{person, assessment, NewCircumferences(measures)}
}

func (w *WaistToHip) String() string {
	v, _ := w.Calc()
	c, _ := w.Classify()
	return fmt.Sprintf("%s\nWTH: %.2f (%s)", w.Person.String(), v, c)
}

// GetName returns this measurement name.
func (w *WaistToHip) GetName() string {
	return "Waist-to-Hip ratio"
}

// Result returns relevant information about waist-to-hip assessment.
func (w *WaistToHip) Result() ([]string, error) {
	rs := []string{}

	v, err := w.Calc()
	if err != nil {
		return rs, err
	}

	c, err := w.Classify()
	if err != nil {
		return rs, err
	}

	rs = append(
		rs,
		fmt.Sprintf("Waist-to-hip ratio: %.2f.", v),
		fmt.Sprintf("Waist-to-hip ratio classification: %s.", c),
	)
	return rs, nil
}

// Classify returns the classification for this assessment.
func (w *WaistToHip) Classify() (string, error) {
	v, err := w.Calc()
	if err != nil {
		return "", err
	}

	genderClass, ok := wthLimits[w.Person.Gender]
	if !ok {
		return "", fmt.Errorf("No classification for gender %d", w.Person.Gender)
	}

	age := w.Person.AgeFromDate(w.Assessment.Date)
	for limits, classes := range genderClass {
		if age < limits[0] || age >= limits[1] {
			continue
		}
		return phass.Classifier(v, classes, WTHClassification), nil
	}

	return "", fmt.Errorf("No classification for age %.0f", age)
}

// Calc returns value for this waist-to-hip assessment.
func (w *WaistToHip) Calc() (float64, error) {
	return w.equation().Calc()
}

// equation returns an equation, used to calculate waist-to-hip.
func (w *WaistToHip) equation() phass.Equationer {
	return phass.NewEquation(wthConf.Extract(w), wthConf)
}

/**
 * Circumferences
 */

// Circumferences represent a collection of circumference measures.
type Circumferences struct {
	Measures map[int]float64
}

// NewCircumferences returns a new Circumferences instance.
func NewCircumferences(measures map[int]float64) *Circumferences {
	return &Circumferences{Measures: measures}
}

// GetName retrieves circumferences measurement name.
func (c *Circumferences) GetName() string {
	return "Cirumferences"
}

// Result show the Circumferences representation.
func (c *Circumferences) Result() ([]string, error) {
	rs := []string{}
	for k, v := range c.Measures {
		rs = append(rs, fmt.Sprintf("Circumference %s: %.2f cm.", NamedCircumference(k), v))
	}
	return rs, nil
}

// NamedCircumference returns the name for a given circumference constant.
func NamedCircumference(name int) string {
	named := map[int]string{
		CCFNeck:         "neck",
		CCFShoulder:     "shoulder",
		CCFChest:        "chest",
		CCFWaist:        "waist",
		CCFAbdominal:    "abdominal",
		CCFHip:          "hip",
		CCFRightArm:     "right arm",
		CCFRightForeArm: "right forearm",
		CCFRightThigh:   "right thigh",
		CCFRightCalf:    "right calf",
		CCFLeftArm:      "left arm",
		CCFLeftForeArm:  "left forearm",
		CCFLeftThigh:    "left thigh",
		CCFLeftCalf:     "left calf",
	}
	return named[name]
}

/**
 * Equation
 */

var (
	wthConf = phass.NewEquationConf(
		"Waist to Hip ratio",
		func(i interface{}) phass.InParams {
			ci := i.(*WaistToHip)
			return map[string]float64{
				"age":                        ci.Person.AgeFromDate(ci.Assessment.Date),
				"gender":                     float64(ci.Person.Gender),
				NamedCircumference(CCFWaist): ci.Circumferences.Measures[CCFWaist],
				NamedCircumference(CCFHip):   ci.Circumferences.Measures[CCFHip],
			}
		},
		[]phass.Validator{
			phass.ValidateAge(20, 69),
			phass.ValidateMeasures([]string{"age", "gender", NamedCircumference(CCFWaist), NamedCircumference(CCFHip)}),
		},
		func(e *phass.Equation) float64 {
			w, _ := e.In(NamedCircumference(CCFWaist))
			h, _ := e.In(NamedCircumference(CCFHip))
			return w / h
		},
	)
	cidConf = phass.NewEquationConf(
		"Conicity index",
		func(i interface{}) phass.InParams {
			ci := i.(*ConicityIndex)
			rs := map[string]float64{
				"weight": ci.Anthropometry.Weight,
				"height": ci.Anthropometry.Height,
			}
			if v, ok := ci.Circumferences.Measures[CCFWaist]; ok {
				rs[NamedCircumference(CCFWaist)] = v
			}
			return rs
		},
		[]phass.Validator{
			phass.ValidateMeasures([]string{"weight", "height", NamedCircumference(CCFWaist)}),
		},
		func(e *phass.Equation) float64 {
			w, _ := e.In("weight")
			h, _ := e.In("height")
			c, _ := e.In(NamedCircumference(CCFWaist))
			return c / 100 / (0.109 * math.Sqrt(w/h/100))
		},
	)
)

/**
 * Classification
 */

// Waist-to-hip classification constants.
const (
	WTHLow = iota
	WTHModerate
	WTHHigh
	WTHVeryHigh
)

// WTHClassification map between constant and string.
var WTHClassification = map[int]string{
	WTHLow:      "Low",
	WTHModerate: "Moderate",
	WTHHigh:     "High",
	WTHVeryHigh: "Very high",
}

// wthLimits represent the classification limits for any given gender, age, and
// waist-to-hip ratio vlaue.
var wthLimits = map[int]map[[2]float64]map[int][2]float64{
	phass.Male: map[[2]float64]map[int][2]float64{
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
	},
	phass.Female: map[[2]float64]map[int][2]float64{
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
	},
}