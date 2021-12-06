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
// 	GET /api/provider/{providerID}
func (c Client) GetProvider(providerID uint) (response.Provider, error) {
	provider := response.Provider{}
	path := fmt.Sprintf("/api/provider/%d", providerID)
	err := c.GetDecoded(&provider, "PROVIDER", path, nil)
	return provider, err
}

// GetProviderList filters providers based on the parameters by invoking the HTTP
// request:
// 	GET /api/provider
func (c Client) GetProviderList(params ProviderSearch) ([]response.Provider, error) {
	providers := response.PaginatedProviders{}

	q, err := query.Values(params)
	if err != nil {
		return providers.List, err
	}

	path := "/api/provider"
	err = c.GetDecoded(&providers, "ARTIFACT", path, q)
	return providers.List, err
}

// UpdateProvider updates the provider with the specified ID by invoking the
// HTTP request:
// 	PUT /api/provider/{providerID}
func (c Client) UpdateProvider(providerID uint, provider request.ProviderUpdate) (response.Provider, error) {
	updatedProvider := response.Provider{}
	body, err := json.Marshal(provider)
	if err != nil {
		return updatedProvider, err
	}

	path := fmt.Sprintf("/api/provider/%d", providerID)
	err = c.PutDecoded(&updatedProvider, "PROVIDER", path, nil, body)
	return updatedProvider, err
}

// CreateProvider creates a new provider by invoking the HTTP request:
// 	POST /api/provider
func (c Client) PostProvider(provider request.Provider) (response.Provider, error) {
	newProvider := response.Provider{}
	body, err := json.Marshal(provider)
	if err != nil {
		return newProvider, err
	}

	path := "/api/provider"
	err = c.PostDecoded(&newProvider, "PROVIDER", path, nil, body)
	return newProvider, err
}
