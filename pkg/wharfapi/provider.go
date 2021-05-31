package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Provider struct {
	ProviderID uint   `json:"providerId"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	UploadURL  string `json:"uploadUrl"`
	TokenID    uint   `json:"tokenId"`
}

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
