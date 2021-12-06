package wharfapi

import (
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

type ArtifactSearch struct {
	BuildID       *uint    `url:"buildId,omitempty"`
	Limit         *int     `url:"limit,omitempty"`
	Offset        *int     `url:"offset,omitempty"`
	OrderBy       []string `url:"orderby,omitempty"`
	Name          *string  `url:"name,omitempty"`
	FileName      *string  `url:"fileName,omitempty"`
	NameMatch     *string  `url:"nameMatch,omitempty"`
	FileNameMatch *string  `url:"fileNameMatch,omitempty"`
	Match         *string  `url:"match,omitempty"`
}

// GetBuildArtifactList filters artifacts based on the parameters by invoking the HTTP
// request:
// 	GET /api/build/{buildId}/artifact
func (c Client) GetBuildArtifactList(params ArtifactSearch, buildID uint) ([]response.Artifact, error) {
	artifacts := response.PaginatedArtifacts{}

	q, err := query.Values(params)
	if err != nil {
		return artifacts.List, err
	}

	path := fmt.Sprintf("/api/build/%d/artifact", buildID)
	err = c.GetDecoded(&artifacts, "ARTIFACT", path, q)
	return artifacts.List, err
}

// GetBuildArtifact gets an artifact by invoking the HTTP request:
//  GET /api/build/{buildId}/artifact/{artifactId}
func (c Client) GetBuildArtifact(buildID, artifactID uint) (response.Artifact, error) {
	artifact := response.Artifact{}
	path := fmt.Sprintf("/api/build/%d/artifact/%d", buildID, artifactID)
	err := c.GetDecoded(&artifact, "ARTIFACT", path, nil)
	return artifact, err
}

// CreateBuildArtifact is not implemented yet.
// Should handle invoking the HTTP request:
//  POST /api/build/{buildId}/artifact
func (c Client) CreateBuildArtifact(buildID uint) error {
	return errors.New("not implemented yet")
}
