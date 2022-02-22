package request

import "time"

// Log is a new log message to be added to a build.
type Log struct {
	BuildID   uint
	Message   string
	Timestamp time.Time
}
