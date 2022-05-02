package wharfapi

import "github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"

// GetEngineList filters builds based on the parameters by invoking the HTTP
// request:
//  GET /api/build
//
// Added in wharf-api v5.1.0.
func (c *Client) GetEngineList() (response.EngineList, error) {
	if err := c.validateEndpointVersion(5, 1, 0); err != nil {
		return response.EngineList{}, err
	}
	var list response.EngineList
	err := c.getUnmarshal("/api/engine", nil, &list)
	return list, err
}
