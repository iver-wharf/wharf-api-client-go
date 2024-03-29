package wharfapi

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
)

// BuildSearch is used when getting builds without using a build ID through the
// HTTP request:
//  GET /api/build
type BuildSearch struct {
	Limit     *int     `url:"limit,omitempty"`
	Offset    *int     `url:"offset,omitempty"`
	OrderBy   []string `url:"orderby,omitempty"`
	ProjectID *uint    `url:"projectId,omitempty"`

	ScheduledAfter  *time.Time `url:"scheduledAfter,omitempty"`
	ScheduledBefore *time.Time `url:"scheduledBefore,omitempty"`
	FinishedAfter   *time.Time `url:"finishedAfter,omitempty"`
	FinishedBefore  *time.Time `url:"finishedBefore,omitempty"`

	IsInvalid *bool    `url:"isInvalid,omitempty"`
	Status    []string `url:"status,omitempty"`
	StatusID  []int    `url:"statusId,omitempty"`

	Environment *string `url:"environment,omitempty"`
	GitBranch   *string `url:"gitBranch,omitempty"`
	Stage       *string `url:"stage,omitempty"`

	EnvironmentMatch *string `url:"environmentMatch,omitempty"`
	GitBranchMatch   *string `url:"gitBranchMatch,omitempty"`
	StageMatch       *string `url:"stageMatch,omitempty"`
	Match            *string `url:"match,omitempty"`
}

// ProjectStartBuild is a range of options you start a build with. The ProjectID and
// Stage fields are required when starting a build.
type ProjectStartBuild struct {
	Stage       string `url:"stage"`
	Branch      string `url:"branch,omitempty"`
	Environment string `url:"environment,omitempty"`
	Engine      string `url:"engine,omitempty"`
}

// GetBuildList filters builds based on the parameters by invoking the HTTP
// request:
//  GET /api/build
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildList(params BuildSearch) (response.PaginatedBuilds, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedBuilds{}, err
	}
	var builds response.PaginatedBuilds
	q, err := query.Values(&params)
	if err != nil {
		return builds, err
	}
	path := "/api/build"
	err = c.getUnmarshal(path, q, &builds)
	return builds, err
}

// GetBuild gets a build by invoking the HTTP request:
//  GET /api/build/{buildId}
//
// Added in wharf-api v0.3.5.
func (c *Client) GetBuild(buildID uint) (response.Build, error) {
	if err := c.validateEndpointVersion(0, 3, 5); err != nil {
		return response.Build{}, err
	}
	path := fmt.Sprintf("/api/build/%d", buildID)
	var build response.Build
	err := c.getUnmarshal(path, nil, &build)
	return build, err
}

// UpdateBuildStatus updates a build by invoking the HTTP request:
//  PUT /api/build/{buildId}/status
//
// Added in wharf-api v5.0.0.
func (c *Client) UpdateBuildStatus(buildID uint, status request.LogOrStatusUpdate) (response.Build, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.Build{}, err
	}
	var updatedBuild response.Build
	path := fmt.Sprintf("/api/build/%d/status", buildID)
	err := c.putJSONUnmarshal(path, nil, status, &updatedBuild)
	return updatedBuild, err
}

// CreateBuildLog adds a new log to a build by invoking the HTTP request:
//  POST /api/build/{buildId}/log
//
// Added in wharf-api v0.1.0.
func (c *Client) CreateBuildLog(buildID uint, buildLog request.LogOrStatusUpdate) error {
	if err := c.validateEndpointVersion(0, 1, 0); err != nil {
		return err
	}
	path := fmt.Sprintf("/api/build/%d/log", buildID)
	ioBody, err := c.postJSON(path, nil, buildLog)
	if err != nil {
		return err
	}
	return ioBody.Close()
}

// GetBuildLogList gets the logs for a build by invoking the HTTP request:
//  GET /api/build/{buildId}/log
//
// Added in wharf-api v0.3.8.
func (c *Client) GetBuildLogList(buildID uint) ([]response.Log, error) {
	if err := c.validateEndpointVersion(0, 3, 8); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/api/build/%d/log", buildID)
	var logs []response.Log
	err := c.getUnmarshal(path, nil, &logs)
	return logs, err
}

// StartProjectBuild starts a new build by invoking the HTTP request:
//  POST /api/project/{projectID}/build
//
// Added in wharf-api v5.0.0.
func (c *Client) StartProjectBuild(projectID uint, params ProjectStartBuild, inputs request.BuildInputs) (response.BuildReferenceWrapper, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.BuildReferenceWrapper{}, err
	}
	var newBuildRef response.BuildReferenceWrapper
	q, err := query.Values(params)
	if err != nil {
		return newBuildRef, err
	}

	path := fmt.Sprintf("/api/project/%d/build", projectID)
	err = c.postJSONUnmarshal(path, q, inputs, &newBuildRef)
	if err == nil {
		log.Debug().WithString("buildRef", newBuildRef.BuildReference).Message("Started build.")
	}
	return newBuildRef, err
}
