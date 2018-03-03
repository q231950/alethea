package ci

// CI represents the type of CI
type CI int

const (
	Unknown CI = 0
	Circle  CI = 1
	Jenkins CI = 2
)

func (ci CI) String() string {
	switch ci {
	case Circle:
		return "Circle CI"
	case Jenkins:
		return "Jenkins CI"
	default:
		return "Unknown CI"
	}
}
