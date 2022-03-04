package wharfapi

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
)

// ProjectSearch is used when getting projects without using a project ID through
// the HTTP request:
//  GET /api/project
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

// CreateProject adds a new project to the database by invoking the
// HTTP request:
//  POST /api/project
//
// Added in wharf-api v0.1.10.
func (c *Client) CreateProject(project request.Project) (response.Project, error) {
	if err := c.validateEndpointVersion(0, 1, 10); err != nil {
		return response.Project{}, err
	}
	var newProject response.Project
	path := "/api/project"
	err := c.postJSONUnmarshal(path, nil, project, &newProject)
	return newProject, err
}

// GetProject fetches a project by ID by invoking the HTTP request:
//  GET /api/project/{projectID}
//
// Added in wharf-api v0.1.10.
func (c *Client) GetProject(projectID uint) (response.Project, error) {
	if err := c.validateEndpointVersion(0, 1, 8); err != nil {
		return response.Project{}, err
	}
	path := fmt.Sprintf("/api/project/%v", projectID)
	var project response.Project
	err := c.getUnmarshal(path, nil, &project)
	return project, err
}

// GetProjectList filters projects based on the parameters by invoking the HTTP
// request:
//  GET /api/project
//
// Added in wharf-api v5.0.0.
func (c *Client) GetProjectList(params ProjectSearch) (response.PaginatedProjects, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedProjects{}, err
	}
	var projects response.PaginatedProjects
	q, err := query.Values(params)
	if err != nil {
		return projects, err
	}
	path := "/api/project"
	err = c.getUnmarshal(path, q, &projects)
	return projects, err
}

// UpdateProject updates a project by ID by invoking the HTTP request:
//  PUT /api/project/{projectID}
//
// Added in wharf-api v5.0.0.
func (c *Client) UpdateProject(projectID uint, project request.ProjectUpdate) (response.Project, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.Project{}, err
	}
	var updatedProject response.Project
	path := fmt.Sprintf("/api/project/%d", projectID)
	err := c.putJSONUnmarshal(path, nil, project, &updatedProject)
	return updatedProject, err
}

// GetProjectOverrides fetches a project's overrides by project ID by invoking the
// HTTP request:
//  GET /api/project/{projectID}/override
//
// Added in wharf-api v5.0.0.
func (c *Client) GetProjectOverrides(projectID uint) (response.ProjectOverrides, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.ProjectOverrides{}, err
	}
	path := fmt.Sprintf("/api/project/%v/override", projectID)
	var overrides response.ProjectOverrides
	err := c.getUnmarshal(path, nil, &overrides)
	return overrides, err
}

// UpdateProjectOverrides updates a project's overrides by project ID by
// invoking the HTTP request:
//  PUT /api/project/{projectID}/override
//
// Added in wharf-api v5.0.0.
func (c *Client) UpdateProjectOverrides(projectID uint, overrides request.ProjectOverridesUpdate) (response.ProjectOverrides, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.ProjectOverrides{}, err
	}
	var updatedOverrides response.ProjectOverrides
	path := fmt.Sprintf("/api/project/%d/override", projectID)
	err := c.putJSONUnmarshal(path, nil, overrides, &updatedOverrides)
	return updatedOverrides, err
}

// DeleteProjectOverrides clears a project's overrides by project ID by
// invoking the HTTP request:
//  DELETE /api/project/{projectID}/override
//
// Added in wharf-api v5.0.0.
func (c *Client) DeleteProjectOverrides(projectID uint) error {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return err
	}
	path := fmt.Sprintf("/api/project/%d/override", projectID)
	resp, err := c.delete(path, nil, nil)
	if err != nil {
		return err
	}
	return resp.Close()
}
