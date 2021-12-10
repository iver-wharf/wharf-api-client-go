package wharfapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api/pkg/model/request"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

type BuildSearch struct {
	Limit     *int     `url:"limit,omitempty"`
	Offset    *int     `url:"offset,omitempty"`
	OrderBy   []string `url:"orderby,omitempty"`
	ProjectID *uint    `url:"projectId,omitempty"`

	ScheduledAfter  *time.Time `url:"scheduledAfter,omitempty"`
	ScheduledBefore *time.Time `url:"scheduledBefore,omitempty"`
	FinishedAfter   *time.Time `url:"finishedAfter,omitempty"`
	FinishedBefore  *time.Time `url:"finishedBefore,omitempty"`

	IsInvalid *bool   `url:"isInvalid,omitempty"`
	Status    *string `url:"status,omitempty"`
	StatusID  *int    `url:"statusId,omitempty"`

	Environment *string `url:"environment,omitempty"`
	GitBranch   *string `url:"gitBranch,omitempty"`
	Stage       *string `url:"stage,omitempty"`

	EnvironmentMatch *string `url:"environmentMatch,omitempty"`
	GitBranchMatch   *string `url:"gitBranchMatch,omitempty"`
	StageMatch       *string `url:"stageMatch,omitempty"`
	Match            *string `url:"match,omitempty"`
}

// BuildParam is an input parameter provided by the user or service that started
// the build.
type BuildParam struct {
	BuildID uint   `json:"buildId"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

// ProjectStartBuild is a range of options you start a build with. The ProjectID and
// Stage fields are required when starting a build.
type ProjectStartBuild struct {
	Stage       string `url:"stage"`
	Branch      string `url:"branch,omitempty"`
	Environment string `url:"environment,omitempty"`
}

// GetBuildList gets all builds by invoking the HTTP request:
//  GET /api/build
func (c Client) GetBuildList(params BuildSearch) (response.PaginatedBuilds, error) {
	builds := response.PaginatedBuilds{}
	q, err := query.Values(&params)
	if err != nil {
		return builds, err
	}
	path := "/api/build"
	err = c.GetUnmarshal(path, q, &builds)
	return builds, err
}

// GetBuild gets a build by invoking the HTTP request:
//  GET /api/build/{buildId}
func (c Client) GetBuild(buildID uint) (response.Build, error) {
	path := fmt.Sprintf("/api/build/%d", buildID)
	build := response.Build{}
	err := c.GetUnmarshal(path, nil, &build)
	return build, err
}

// UpdateBuildStatus updates a build by invoking the HTTP request:
//  PUT /api/build/{buildId}/status
func (c Client) UpdateBuildStatus(buildID uint, status request.LogOrStatusUpdate) (response.Build, error) {
	updatedBuild := response.Build{}
	path := fmt.Sprintf("/api/build/%d/status", buildID)
	err := c.PutJSONUnmarshal(path, nil, &status, &updatedBuild)
	return updatedBuild, err
}

// CreateBuildLog adds a new log to a build by invoking the HTTP request:
//  POST /api/build/{buildId}/log
func (c Client) CreateBuildLog(buildID uint, buildLog request.LogOrStatusUpdate) error {
	body, err := json.Marshal(buildLog)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/api/build/%d/log", buildID)
	_, err = c.Post(path, nil, body)
	return err
}

// GetBuildLogList gets the logs for a build by invoking the HTTP request:
//  GET /api/build/{buildId}/log
func (c Client) GetBuildLogList(buildID uint) ([]response.Log, error) {
	path := fmt.Sprintf("/api/build/%d/log", buildID)
	logs := []response.Log{}
	err := c.GetUnmarshal(path, nil, &logs)
	return logs, err
}

// StartProjectBuild starts a new build by invoking the HTTP request:
//  POST /api/project/{projectID}/build
func (c Client) StartProjectBuild(projectID uint, params ProjectStartBuild, inputs request.BuildInputs) (response.BuildReferenceWrapper, error) {
	newBuildRef := response.BuildReferenceWrapper{}
	q, err := query.Values(params)
	if err != nil {
		return newBuildRef, err
	}

	path := fmt.Sprintf("/api/project/%d/build", projectID)
	err = c.PostJSONUnmarshal(path, q, &inputs, &newBuildRef)
	if err == nil {
		log.Debug().WithString("buildRef", newBuildRef.BuildReference).Message("Started build.")
	}
	return newBuildRef, err
}
