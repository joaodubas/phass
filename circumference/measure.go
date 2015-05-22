package circumference

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

type Circumferences struct {
	Measures map[int]float64
}
