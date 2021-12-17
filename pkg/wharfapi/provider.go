package wharfapi

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api/pkg/model/request"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

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
func (c Client) GetProvider(providerID uint) (response.Provider, error) {
	var provider response.Provider
	path := fmt.Sprintf("/api/provider/%d", providerID)
	err := c.GetUnmarshal(path, nil, &provider)
	return provider, err
}

// GetProviderList filters providers based on the parameters by invoking the HTTP
// request:
//  GET /api/provider
func (c Client) GetProviderList(params ProviderSearch) (response.PaginatedProviders, error) {
	var providers response.PaginatedProviders

	q, err := query.Values(params)
	if err != nil {
		return providers, err
	}

	path := "/api/provider"
	err = c.GetUnmarshal(path, q, &providers)
	return providers, err
}

// UpdateProvider updates the provider with the specified ID by invoking the
// HTTP request:
//  PUT /api/provider/{providerID}
func (c Client) UpdateProvider(providerID uint, provider request.ProviderUpdate) (response.Provider, error) {
	var updatedProvider response.Provider
	path := fmt.Sprintf("/api/provider/%d", providerID)
	err := c.PutJSONUnmarshal(path, nil, provider, &updatedProvider)
	return updatedProvider, err
}

// CreateProvider creates a new provider by invoking the HTTP request:
//  POST /api/provider
func (c Client) CreateProvider(provider request.Provider) (response.Provider, error) {
	var newProvider response.Provider
	path := "/api/provider"
	err := c.PostJSONUnmarshal(path, nil, provider, &newProvider)
	return newProvider, err
}
