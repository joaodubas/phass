package circumference

const (
	WTHLow = iota
	WTHModerate
	WTHHigh
	WTHVeryHigh
)

var WTHClassification = map[int]string{
	WTHLow:      "Low",
	WTHModerate: "Moderate",
	WTHHigh:     "High",
	WTHVeryHigh: "Very high",
}
