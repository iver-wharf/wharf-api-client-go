package wharfapi

import (
	"fmt"

	"github.com/iver-wharf/wharf-api/pkg/model/request"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

// CreateProjectBranch adds a branch to the project with the matching
// project ID by invoking the HTTP request:
//  POST /api/project/{projectId}/branch
func (c Client) CreateProjectBranch(projectID uint, branch request.Branch) (response.Branch, error) {
	var newBranch response.Branch
	path := fmt.Sprintf("/api/project/%d/branch", projectID)
	err := c.postJSONUnmarshal(path, nil, branch, &newBranch)
	return newBranch, err
}

// UpdateProjectBranchList resets the default branch and list of branches for a project
// using the project ID from the first branch in the provided list by invoking
// the HTTP request:
//  PUT /api/project/{projectId}/branch
func (c Client) UpdateProjectBranchList(projectID uint, branches []request.Branch) ([]response.Branch, error) {
	var newBranches []response.Branch
	path := fmt.Sprintf("/api/project/%d/branch", projectID)
	err := c.putJSONUnmarshal(path, nil, branches, &newBranches)
	return newBranches, err
}

// GetProjectBranchList gets the branches for a project by invoking the HTTP
// request:
//  GET /api/project/{projectId}/branch
func (c Client) GetProjectBranchList(projectID uint) ([]response.Branch, error) {
	path := fmt.Sprintf("/api/project/%d/branch", projectID)
	var branches []response.Branch
	err := c.getUnmarshal(path, nil, &branches)
	return branches, err
}
