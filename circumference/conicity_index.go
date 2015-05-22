package circumference

import (
	"fmt"
	anthropo "github.com/joaodubas/phass/anthropometry"
	"math"
)

type ConicityIndex struct {
	*anthropo.Anthropometry
	*Circumferences
}

func (c *ConicityIndex) String() string {
	return fmt.Sprintf("%s\nConicity index: %.4f", c.Anthropometry.String(), c.Calc())
}

func (c *ConicityIndex) Classify() string {
	return "No classification available yet"
}

func (c *ConicityIndex) Calc() float64 {
	circ, ok := c.Measures[CCFWaist]
	if !ok {
		return 0.0
	}
	return circ / 100 / (0.109 * math.Sqrt(c.Weight/c.Height/100))
}
