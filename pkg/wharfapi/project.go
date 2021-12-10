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

func (c Client) CreateProject(project request.Project) (response.Project, error) {
	newProject := response.Project{}
	body, err := json.Marshal(&project)
	if err != nil {
		return newProject, err
	}

	path := "/api/project"
	err = c.PostJSONUnmarshal(path, nil, body, &newProject)
	return newProject, err
}

// // ProjectRunResponse contains metadata about the newly started build.
// type ProjectRunResponse = response.BuildReferenceWrapper

// GetProjectByID fetches a project by ID by invoking the HTTP request:
//  GET /api/project/{projectID}
func (c Client) GetProject(projectID uint) (response.Project, error) {
	path := fmt.Sprintf("/api/project/%v", projectID)
	project := response.Project{}
	err := c.GetUnmarshal(path, nil, &project)
	return project, err
}

// GetProjectList filters projects based on the parameters by invoking the HTTP
// request:
//  GET /api/project
func (c Client) GetProjectList(params ProjectSearch) (response.PaginatedProjects, error) {
	projects := response.PaginatedProjects{}
	q, err := query.Values(params)
	if err != nil {
		return projects, err
	}
	path := "/api/project"
	err = c.GetUnmarshal(path, q, &projects)
	return projects, err
}

// UpdateProject updates a project by ID by invoking the HTTP request:
//  PUT /api/project/{projectID}
func (c Client) UpdateProject(projectID uint, project request.ProjectUpdate) (response.Project, error) {
	updatedProject := response.Project{}
	body, err := json.Marshal(project)
	if err != nil {
		return updatedProject, err
	}

	path := fmt.Sprintf("/api/project/%d", projectID)
	err = c.PutJSONUnmarshal(path, nil, body, &updatedProject)
	return updatedProject, err
}
