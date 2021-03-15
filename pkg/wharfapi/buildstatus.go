package wharfapi

import "strconv"

type BuildStatus int

const (
	BuildScheduling = BuildStatus(iota)
	BuildRunning
	BuildCompleted
	BuildFailed
)

func (bs BuildStatus) String() string {
	switch bs {
	case BuildScheduling:
		return "Scheduling"
	case BuildRunning:
		return "Running"
	case BuildCompleted:
		return "Completed"
	case BuildFailed:
		return "Failed"
	default:
		return strconv.Itoa(int(bs))
	}
}
