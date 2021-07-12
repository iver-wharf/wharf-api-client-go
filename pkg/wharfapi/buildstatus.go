package wharfapi

import "strconv"

// BuildStatus is the state of a build.
//
// The flow of the build status goes like this:
//
//            Scheduling
//                |
//                V
//             Running
//               / \
//  Completed <-/   \-> Failed
type BuildStatus int

const (
	// BuildScheduling means that the build has not started execution yet.
	BuildScheduling = BuildStatus(iota)
	// BuildRunning means that the build is currently executing code.
	BuildRunning
	// BuildCompleted means that the build ran successfully to completetion.
	BuildCompleted
	// BuildFailed means that the build ran unsuccessfully.
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
