package wharfapi

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api/v5/pkg/model/response"
)

// ArtifactSearch is used when getting artifacts without using an artifact ID
// through the HTTP request:
//  GET /api/build/{buildId}/artifact
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
//  GET /api/build/{buildId}/artifact
func (c Client) GetBuildArtifactList(params ArtifactSearch, buildID uint) (response.PaginatedArtifacts, error) {
	var artifacts response.PaginatedArtifacts
	q, err := query.Values(params)
	if err != nil {
		return artifacts, err
	}
	path := fmt.Sprintf("/api/build/%d/artifact", buildID)
	err = c.getUnmarshal(path, q, &artifacts)
	return artifacts, err
}

// GetBuildArtifact gets an artifact by invoking the HTTP request:
//  GET /api/build/{buildId}/artifact/{artifactId}
func (c Client) GetBuildArtifact(buildID, artifactID uint) (response.Artifact, error) {
	var artifact response.Artifact
	path := fmt.Sprintf("/api/build/%d/artifact/%d", buildID, artifactID)
	err := c.getUnmarshal(path, nil, &artifact)
	return artifact, err
}
