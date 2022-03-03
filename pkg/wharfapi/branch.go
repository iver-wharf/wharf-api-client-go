package wharfapi

import (
	"fmt"

	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
)

// CreateProjectBranch adds a branch to the project with the matching
// project ID by invoking the HTTP request:
//  POST /api/project/{projectId}/branch
//
// Added in wharf-api v5.0.0.
func (c *Client) CreateProjectBranch(projectID uint, branch request.Branch) (response.Branch, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.Branch{}, err
	}
	var newBranch response.Branch
	path := fmt.Sprintf("/api/project/%d/branch", projectID)
	err := c.postJSONUnmarshal(path, nil, branch, &newBranch)
	return newBranch, err
}

// UpdateProjectBranchList resets the default branch and list of branches for a project
// using the project ID from the first branch in the provided list by invoking
// the HTTP request:
//  PUT /api/project/{projectId}/branch
//
// Added in wharf-api v5.0.0.
func (c *Client) UpdateProjectBranchList(projectID uint, branches []request.Branch) ([]response.Branch, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return nil, err
	}
	var newBranches []response.Branch
	path := fmt.Sprintf("/api/project/%d/branch", projectID)
	err := c.putJSONUnmarshal(path, nil, branches, &newBranches)
	return newBranches, err
}

// GetProjectBranchList gets the branches for a project by invoking the HTTP
// request:
//  GET /api/project/{projectId}/branch
//
// Added in wharf-api v5.0.0.
func (c *Client) GetProjectBranchList(projectID uint) ([]response.Branch, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/api/project/%d/branch", projectID)
	var branches []response.Branch
	err := c.getUnmarshal(path, nil, &branches)
	return branches, err
}
