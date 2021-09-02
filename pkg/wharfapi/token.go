package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Token is a value used by the provider plugins to authenticate with the remote
// providers.
type Token struct {
	TokenID  uint   `json:"tokenId"`
	Token    string `json:"token"`
	UserName string `json:"userName"`
}

// GetTokenByID fetches a token by ID by invoking the HTTP request:
// 	GET /api/token/{tokenID}
func (c Client) GetTokenByID(tokenID uint) (Token, error) {
	newToken := Token{}

	apiURL := fmt.Sprintf("%s/api/token/%v", c.APIURL, tokenID)
	ioBody, err := doRequest("GET | TOKEN |", http.MethodGet, apiURL, []byte{}, c.AuthHeader)
	if err != nil {
		return newToken, err
	}

	defer (*ioBody).Close()

	err = json.NewDecoder(*ioBody).Decode(&newToken)
	if err != nil {
		return newToken, err
	}
	return newToken, nil
}

// GetToken tries to search for a token using the username+token pair by
// invoking the HTTP request:
// 	POST /api/token/search
func (c Client) GetToken(token string, userName string) (Token, error) {
	newToken := Token{}

	path := "/api/tokens/search"

	data := url.Values{}
	data.Set("Token", token)
	if userName != "" {
		data.Add("UserName", userName)
	}

	u, _ := url.ParseRequestURI(c.APIURL)
	u.Path = path
	u.RawQuery = data.Encode()

	ioBody, err := doRequest("GET | TOKEN |", http.MethodPost, fmt.Sprintf("%v", u), []byte{}, c.AuthHeader)
	if err != nil {
		return newToken, err
	}

	defer (*ioBody).Close()

	var tokens []Token
	err = json.NewDecoder(*ioBody).Decode(&tokens)
	if err != nil {
		return newToken, err
	}

	if len(tokens) == 0 {
		return newToken, nil
	}

	return tokens[0], nil
}

// SearchToken tries to match the given token on the token field and username
// field, by invoking the HTTP request:
// 	POST /api/tokens/search
func (c Client) SearchToken(token Token) ([]Token, error) {
	body, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/tokens/search", c.APIURL)
	ioBody, err := doRequest("SEARCH | TOKEN |", http.MethodPost, url, body, c.AuthHeader)
	if err != nil {
		return nil, err
	}

	defer (*ioBody).Close()

	var foundTokens []Token
	err = json.NewDecoder(*ioBody).Decode(&foundTokens)
	if err != nil {
		return nil, err
	}

	return foundTokens, nil
}

// PutToken tries to match an existing token by ID or username+token and updates
// it, or adds a new a token if none matched, by invoking the HTTP request:
// 	PUT /api/token
func (c Client) PutToken(token Token) (Token, error) {
	body, err := json.Marshal(token)
	if err != nil {
		return Token{}, err
	}

	apiURL := fmt.Sprintf("%s/api/token", c.APIURL)
	ioBody, err := doRequest("PUT | TOKEN |", http.MethodPut, apiURL, body, c.AuthHeader)
	if err != nil {
		return Token{}, err
	}

	defer (*ioBody).Close()

	var newToken Token
	if err := json.NewDecoder(*ioBody).Decode(&newToken); err != nil {
		return Token{}, err
	}

	return newToken, nil
}

// PostToken adds a new a token by invoking the HTTP request:
// 	POST /api/token
func (c Client) PostToken(token Token) (Token, error) {
	newToken := Token{}
	body, err := json.Marshal(token)
	if err != nil {
		return newToken, err
	}

	apiURL := fmt.Sprintf("%s/api/token", c.APIURL)
	ioBody, err := doRequest("POST | TOKEN", http.MethodPost, apiURL, body, c.AuthHeader)
	if err != nil {
		return newToken, err
	}

	defer (*ioBody).Close()

	err = json.NewDecoder(*ioBody).Decode(&newToken)
	if err != nil {
		return newToken, err
	}

	return newToken, nil
}
