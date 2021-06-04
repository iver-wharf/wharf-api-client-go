package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Token struct {
	TokenID    uint   `json:"tokenId"`
	Token      string `json:"token"`
	UserName   string `json:"userName"`
}

func (c Client) GetTokenByID(tokenID uint) (Token, error) {
	newToken := Token{}

	url := fmt.Sprintf("%s/api/token/%v", c.ApiUrl, tokenID)
	ioBody, err := doRequest("GET | TOKEN |", http.MethodGet, url, []byte{}, c.AuthHeader)
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

func (c Client) GetToken(token string, userName string) (Token, error) {
	newToken := Token{}

	path := "/api/tokens/search"

	data := url.Values{}
	data.Set("Token", token)
	if userName != "" {
		data.Add("UserName", userName)
	}

	u, _ := url.ParseRequestURI(c.ApiUrl)
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

func (c Client) PostToken(token Token) (Token, error) {
	newToken := Token{}
	body, err := json.Marshal(token)
	if err != nil {
		return newToken, err
	}

	url := fmt.Sprintf("%s/api/token", c.ApiUrl)
	ioBody, err := doRequest("POST | TOKEN", http.MethodPost, url, body, c.AuthHeader)
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
