package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Branch holds metadata about a Git branch for a given Wharf project.
type Branch struct {
	BranchID  uint   `json:"branchId"`
	ProjectID uint   `json:"projectId"`
	Name      string `json:"name"`
	TokenID   uint   `json:"tokenId"`
	Default   bool   `json:"default"`
}

// PutBranch tries to match an existing branch by ID or by name,
// or adds a new a provider if none matched, by invoking the HTTP request:
// 	PUT /api/branch
func (c Client) PutBranch(branch Branch) (Branch, error) {
	newBranch := Branch{}

	body, err := json.Marshal(branch)
	if err != nil {
		return newBranch, err
	}

	url := fmt.Sprintf("%s/api/branch", c.APIURL)
	ioBody, err := doRequest("POST | BRANCH |", http.MethodPost, url, body, c.AuthHeader)
	if err != nil {
		return newBranch, err
	}

	defer (*ioBody).Close()

	err = json.NewDecoder(*ioBody).Decode(&newBranch)
	if err != nil {
		return newBranch, err
	}

	return newBranch, nil
}

// PutBranches resets the default branch and list of branches for a project
// using the project ID from the first branch in the provided list by invoking
// the HTTP request:
// 	PUT /api/branches
func (c Client) PutBranches(branches []Branch) ([]Branch, error) {
	var newBranches []Branch
	body, err := json.Marshal(branches)
	if err != nil {
		return newBranches, err
	}

	url := fmt.Sprintf("%s/api/branches", c.APIURL)
	ioBody, err := doRequest("PUT | BRANCHES |", http.MethodPut, url, body, c.AuthHeader)
	if err != nil {
		return newBranches, err
	}

	defer (*ioBody).Close()

	err = json.NewDecoder(*ioBody).Decode(&newBranches)
	if err != nil {
		return newBranches, err
	}

	return newBranches, nil
}
