package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Branch struct {
	BranchID  uint   `json:"branchId"`
	ProjectID uint   `json:"projectId"`
	Name      string `json:"name"`
	TokenID   uint   `json:"tokenId"`
	Default   bool   `json:"default"`
}

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
