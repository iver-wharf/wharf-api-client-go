package wharfapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
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
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildArtifactList(params ArtifactSearch, buildID uint) (response.PaginatedArtifacts, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedArtifacts{}, err
	}
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
//
// Added in wharf-api v0.7.1.
func (c *Client) GetBuildArtifact(buildID, artifactID uint) (response.Artifact, error) {
	// TODO: Receive multipart file, not response.Artifact
	if err := c.validateEndpointVersion(0, 7, 1); err != nil {
		return response.Artifact{}, err
	}
	var artifact response.Artifact
	path := fmt.Sprintf("/api/build/%d/artifact/%d", buildID, artifactID)
	err := c.getUnmarshal(path, nil, &artifact)
	return artifact, err
}

// CreateBuildArtifact uploads an artifact by invoking the HTTP request:
//  POST /api/build/{buildId}/artifact
//
// Added in wharf-api v0.4.9.
func (c Client) CreateBuildArtifact(buildID uint, fileName string, artifact io.Reader) error {
	if err := c.validateEndpointVersion(0, 4, 9); err != nil {
		return err
	}
	path := fmt.Sprintf("/api/build/%d/artifact", buildID)
	resp, err := c.uploadMultipart(http.MethodPost, path, map[string]file{
		"files": {
			fileName: fileName,
			reader:   artifact,
		},
	})
	if err != nil {
		return err
	}
	return resp.Close()
}
