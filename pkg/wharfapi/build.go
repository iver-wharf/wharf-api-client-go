package wharfapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/iver-wharf/wharf-api/pkg/model/request"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
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

// GetBuildList gets all builds by invoking the HTTP request:
//  GET /api/build
func (c Client) GetBuildList() ([]response.Build, error) {
	path := "/api/build"
	list := []response.Build{}
	err := c.PutDecoded(&list, "BUILD", path, nil, nil)
	if err != nil {
		return []response.Build{}, err
	}
	return list, nil
}

// GetBuild gets a build by invoking the HTTP request:
//  GET /api/build/{buildId}
func (c Client) GetBuild(buildID uint) (response.Build, error) {
	path := fmt.Sprintf("/api/build/%d", buildID)
	build := response.Build{}
	err := c.GetDecoded(&build, "BUILD", path, nil)
	return build, err
}

// UpdateBuildStatus updates a build by invoking the HTTP request:
// 	PUT /api/build/{buildId}/status
func (c Client) UpdateBuildStatus(buildID uint, status request.LogOrStatusUpdate) (response.Build, error) {
	updatedBuild := response.Build{}
	body, err := json.Marshal(&status)
	if err != nil {
		return updatedBuild, err
	}

	path := fmt.Sprintf("/api/build/%d/status", buildID)
	err = c.PutDecoded(&updatedBuild, "BUILD", path, nil, body)
	return updatedBuild, err
}

// CreateBuildLog adds a new log to a build by invoking the HTTP request:
// 	POST /api/build/{buildId}/log
func (c Client) CreateBuildLog(buildID uint, buildLog request.LogOrStatusUpdate) error {
	body, err := json.Marshal(buildLog)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/api/build/%d/log", buildID)
	_, err = c.Post("BUILD", path, nil, body)
	return err
}

// GetBuildLogList gets the logs for a build by invoking the HTTP request:
//  GET /api/build/{buildId}/log
func (c Client) GetBuildLogList(buildID uint) ([]response.Log, error) {
	path := fmt.Sprintf("/api/build/%d/log", buildID)
	list := []response.Log{}
	err := c.GetDecoded(&list, "BUILD", path, nil)
	return list, err
}

// StreamBuildLog is not implemented yet.
// Should handle invoking the HTTP request:
//  GET /api/build/{buildId}/stream
func (c Client) StreamBuildLog(buildID uint) error {
	return errors.New("not implemented yet")
}
