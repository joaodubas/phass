package common

func Classifier(value float64, classes map[int][]float64, mapper map[int]string) string {
	cid := classifierId(value, classes)
	class, ok := mapper[cid]
	if !ok {
		return "No classification."
	}
	return class
}

func classifierId(value float64, classes map[int][]float64) int {
	cid := -1
	for id_, limits := range classes {
		if value >= limits[0] && value < limits[1] {
			cid = id_
		}
	}
	return cid
}
