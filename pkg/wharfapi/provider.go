package wharfapi

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
)

// ProviderSearch is used when getting providers without using a provider ID
// through the HTTP request:
//  GET /api/provider
type ProviderSearch struct {
	Limit     *int     `url:"limit,omitempty"`
	Offset    *int     `url:"offset,omitempty"`
	OrderBy   []string `url:"orderby,omitempty"`
	Name      *string  `url:"name,omitempty"`
	URL       *string  `url:"url,omitempty"`
	NameMatch *string  `url:"nameMatch,omitempty"`
	URLMatch  *string  `url:"urlMatch,omitempty"`
	Match     *string  `url:"match,omitempty"`
}

// GetProvider fetches a provider by ID by invoking the HTTP request:
//  GET /api/provider/{providerID}
//
// Added in wharf-api v0.3.9.
func (c *Client) GetProvider(providerID uint) (response.Provider, error) {
	if err := c.validateEndpointVersion(0, 3, 9); err != nil {
		return response.Provider{}, err
	}
	var provider response.Provider
	path := fmt.Sprintf("/api/provider/%d", providerID)
	err := c.getUnmarshal(path, nil, &provider)
	return provider, err
}

// GetProviderList filters providers based on the parameters by invoking the HTTP
// request:
//  GET /api/provider
//
// Added in wharf-api v5.0.0.
func (c *Client) GetProviderList(params ProviderSearch) (response.PaginatedProviders, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedProviders{}, err
	}
	var providers response.PaginatedProviders

	q, err := query.Values(params)
	if err != nil {
		return providers, err
	}

	path := "/api/provider"
	err = c.getUnmarshal(path, q, &providers)
	return providers, err
}

// UpdateProvider updates the provider with the specified ID by invoking the
// HTTP request:
//  PUT /api/provider/{providerID}
//
// Added in wharf-api v5.0.0.
func (c *Client) UpdateProvider(providerID uint, provider request.ProviderUpdate) (response.Provider, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.Provider{}, err
	}
	var updatedProvider response.Provider
	path := fmt.Sprintf("/api/provider/%d", providerID)
	err := c.putJSONUnmarshal(path, nil, provider, &updatedProvider)
	return updatedProvider, err
}

// CreateProvider creates a new provider by invoking the HTTP request:
//  POST /api/provider
//
// Added in wharf-api v0.3.9.
func (c *Client) CreateProvider(provider request.Provider) (response.Provider, error) {
	if err := c.validateEndpointVersion(0, 3, 9); err != nil {
		return response.Provider{}, err
	}
	var newProvider response.Provider
	path := "/api/provider"
	err := c.postJSONUnmarshal(path, nil, provider, &newProvider)
	return newProvider, err
}
