package wharfapi

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// Build is the metadata about a code execution for a certain project inside
// Wharf.
type Build struct {
	BuildID     uint         `json:"buildId"`
	StatusID    BuildStatus  `json:"statusId"`
	ProjectID   uint         `json:"projectId"`
	ScheduledOn *time.Time   `json:"scheduledOn"`
	StartedOn   *time.Time   `json:"startedOn"`
	CompletedOn *time.Time   `json:"finishedOn"`
	GitBranch   string       `json:"gitBranch"`
	Environment null.String  `json:"environment"`
	Stage       string       `json:"stage"`
	Params      []BuildParam `json:"params"`
	IsInvalid   bool         `json:"isInvalid"`
}

// BuildParam is an input parameter provided by the user or service that started
// the build.
type BuildParam struct {
	BuildID uint   `json:"buildId"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}
