package wharfapi

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

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

type BuildParam struct {
	BuildID      uint   `json:"buildId"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}
