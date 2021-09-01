package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Provider holds metadata about a provider registered in Wharf.
type Provider struct {
	ProviderID uint   `json:"providerId"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	UploadURL  string `json:"uploadUrl"`
	TokenID    uint   `json:"tokenId"`
}

// GetProviderByID fetches a provider by ID by invoking the HTTP request:
// 	GET /api/provider/{providerID}
func (c Client) GetProviderByID(providerID uint) (Provider, error) {
	newProvider := Provider{}

	apiURL := fmt.Sprintf("%s/api/provider/%v", c.APIURL, providerID)
	ioBody, err := doRequest("GET | PROVIDER |", http.MethodGet, apiURL, []byte{}, c.AuthHeader)
	if err != nil {
		return newProvider, err
	}

	defer (*ioBody).Close()

	err = json.NewDecoder(*ioBody).Decode(&newProvider)
	if err != nil {
		return newProvider, err
	}
	return newProvider, nil
}

// GetProvider tries to find a provider based on its name, URL, etc. by invoking
// the HTTP request:
// 	POST /api/providers/search
func (c Client) GetProvider(providerName string, urlStr string, uploadURLStr string, tokenID uint) (Provider, error) {
	newProvider := Provider{}

	path := "/api/providers/search"
	data := url.Values{}
	data.Set("Name", providerName)
	data.Add("URL", urlStr)
	data.Add("UploadURL", uploadURLStr)

	if tokenID > 0 {
		data.Add("TokenID", fmt.Sprint(tokenID))
	}

	u, _ := url.ParseRequestURI(c.APIURL)
	u.Path = path
	u.RawQuery = data.Encode()
	apiURL := fmt.Sprintf("%v", u)

	ioBody, err := doRequest("GET | PROVIDER |", http.MethodPost, apiURL, []byte{}, c.AuthHeader)
	if err != nil {
		return newProvider, err
	}

	defer (*ioBody).Close()

	var providers []Provider
	err = json.NewDecoder(*ioBody).Decode(&providers)
	if err != nil {
		return newProvider, err
	}

	if len(providers) == 0 {
		return newProvider, nil
	}

	return providers[0], nil
}

// SearchProvider tries to match the given provider on the provider name, URL,
// upload URL, and token ID, by invoking the HTTP request:
// 	POST /api/providers/search
//
// The token ID is not queried if the argument's tokenID field is set to zero.
func (c Client) SearchProvider(provider Provider) ([]Provider, error) {
	body, err := json.Marshal(provider)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/projects/search", c.APIURL)
	ioBody, err := doRequest("SEARCH | PROVIDER |", http.MethodPost, url, body, c.AuthHeader)
	if err != nil {
		return nil, err
	}

	defer (*ioBody).Close()

	var foundProviders []Provider
	err = json.NewDecoder(*ioBody).Decode(&foundProviders)
	if err != nil {
		return nil, err
	}

	return foundProviders, nil
}

// PutProvider tries to match an existing provider by ID or combination of name,
// URL, etc. and updates it, or adds a new a provider if none matched,
// by invoking the HTTP request:
// 	PUT /api/provider
func (c Client) PutProvider(provider Provider) (Provider, error) {
	body, err := json.Marshal(provider)
	if err != nil {
		return Provider{}, err
	}

	apiURL := fmt.Sprintf("%s/api/provider", c.APIURL)
	ioBody, err := doRequest("PUT | PROVIDER |", http.MethodPut, apiURL, body, c.AuthHeader)
	if err != nil {
		return Provider{}, err
	}

	defer (*ioBody).Close()

	var newProvider Provider
	if err := json.NewDecoder(*ioBody).Decode(&newProvider); err != nil {
		return Provider{}, err
	}

	return newProvider, nil
}

// PostProvider adds a new a provider by invoking the HTTP request:
// 	POST /api/provider
func (c Client) PostProvider(provider Provider) (Provider, error) {
	newProvider := Provider{}
	body, err := json.Marshal(provider)
	if err != nil {
		return newProvider, err
	}

	apiURL := fmt.Sprintf("%s/api/provider", c.APIURL)
	ioBody, err := doRequest("POST | PROVIDER |", http.MethodPost, apiURL, body, c.AuthHeader)
	if err != nil {
		return newProvider, err
	}

	defer (*ioBody).Close()

	err = json.NewDecoder(*ioBody).Decode(&newProvider)
	if err != nil {
		return newProvider, err
	}

	return newProvider, nil
}
