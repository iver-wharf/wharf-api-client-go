package wharfapi

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api/pkg/model/request"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

// type Project response.Project
// type ProjectUpdate request.ProjectUpdate
type ProjectSearch struct {
	OrderBy          []string `url:"orderby,omitempty"`
	Limit            *int     `url:"limit,omitempty"`
	Offset           *int     `url:"offset,omitempty"`
	Name             *string  `url:"name,omitempty"`
	GroupName        *string  `url:"groupName,omitempty"`
	Description      *string  `url:"description,omitempty"`
	TokenID          *uint    `url:"tokenId,omitempty"`
	ProviderID       *uint    `url:"providerId,omitempty"`
	GitURL           *string  `url:"gitUrl,omitempty"`
	NameMatch        *string  `url:"nameMatch,omitempty"`
	GroupNameMatch   *string  `url:"groupNameMatch,omitempty"`
	DescriptionMatch *string  `url:"descriptionMatch,omitempty"`
	GitURLMatch      *string  `url:"gitUrlMatch,omitempty"`
	Match            *string  `url:"match,omitempty"`
}

// ProjectRun is a range of options you start a build with. The ProjectID and
// Stage fields are required when starting a build.
type ProjectRun struct {
	ProjectID   uint   `json:"projectId" url:"-"`
	Stage       string `json:"stage" url:"-"`
	Branch      string `json:"branch" url:"branch"`
	Environment string `json:"environment" url:"environment"`
}

// // ProjectRunResponse contains metadata about the newly started build.
// type ProjectRunResponse = response.BuildReferenceWrapper

// GetProjectByID fetches a project by ID by invoking the HTTP request:
// 	GET /api/project/{projectID}
func (c Client) GetProject(projectID uint) (response.Project, error) {
	path := fmt.Sprintf("/api/project/%v", projectID)
	project := response.Project{}
	err := c.GetDecoded(&project, "PROJECT", path, nil)
	return project, err
}

// GetProjectList filters projects based on the parameters by invoking the HTTP
// request:
// 	GET /api/project
func (c Client) GetProjectList(params ProjectSearch) ([]response.Project, error) {
	q, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	path := "/api/project"
	projects := response.PaginatedProjects{}
	err = c.GetDecoded(&projects, "PROJECT", path, q)
	return projects.List, err
}

// UpdateProject updates a project by ID by invoking the HTTP request:
// 	PUT /api/project/{projectID}
func (c Client) UpdateProject(projectID uint, project request.ProjectUpdate) (response.Project, error) {
	updatedProject := response.Project{}
	body, err := json.Marshal(project)
	if err != nil {
		return updatedProject, err
	}

	path := fmt.Sprintf("/api/project/%d", projectID)
	err = c.PutDecoded(&updatedProject, "PROJECT", path, nil, body)
	return updatedProject, err
}

// PostProjectRun starts a new build by invoking the HTTP request:
// 	POST /api/project/{projectID}/{stage}/run
func (c Client) StartProjectBuild(projectRun ProjectRun, inputs request.BuildInputs) (response.BuildReferenceWrapper, error) {
	newBuildRef := response.BuildReferenceWrapper{}

	body, err := json.Marshal(inputs)
	if err != nil {
		return newBuildRef, err
	}

	q, err := query.Values(projectRun)
	if err != nil {
		return newBuildRef, err
	}

	path := fmt.Sprintf("/api/project/%d/%s/run", projectRun.ProjectID, projectRun.Stage)
	err = c.PostDecoded(&newBuildRef, "PROJECT", path, q, body)
	if err == nil {
		log.Debug().WithString("buildRef", newBuildRef.BuildReference).Message("Started build.")
	}
	return newBuildRef, err
}
