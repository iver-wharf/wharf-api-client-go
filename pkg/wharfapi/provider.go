package wharfapi

import (
	"encoding/json"
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
	provider := response.Provider{}
	path := fmt.Sprintf("/api/provider/%d", providerID)
	err := c.GetUnmarshal(path, nil, &provider)
	return provider, err
}

// GetProviderList filters providers based on the parameters by invoking the HTTP
// request:
//  GET /api/provider
func (c Client) GetProviderList(params ProviderSearch) (response.PaginatedProviders, error) {
	providers := response.PaginatedProviders{}

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
	updatedProvider := response.Provider{}
	body, err := json.Marshal(provider)
	if err != nil {
		return updatedProvider, err
	}

	path := fmt.Sprintf("/api/provider/%d", providerID)
	err = c.PutJSONUnmarshal(path, nil, body, &updatedProvider)
	return updatedProvider, err
}

// CreateProvider creates a new provider by invoking the HTTP request:
//  POST /api/provider
func (c Client) CreateProvider(provider request.Provider) (response.Provider, error) {
	newProvider := response.Provider{}
	body, err := json.Marshal(provider)
	if err != nil {
		return newProvider, err
	}

	path := "/api/provider"
	err = c.PostJSONUnmarshal(path, nil, body, &newProvider)
	return newProvider, err
}
