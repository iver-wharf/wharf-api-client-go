package wharfapi

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api/pkg/model/request"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

type TokenSearch struct {
	Limit         *int     `url:"limit,omitempty"`
	Offset        *int     `url:"offset,omitempty"`
	OrderBy       []string `url:"orderby,omitempty"`
	UserName      *string  `url:"userName,omitempty"`
	UserNameMatch *string  `url:"userNameMatch,omitempty"`
}

// GetToken fetches a token by ID by invoking the HTTP request:
// 	GET /api/token/{tokenID}
func (c Client) GetToken(tokenID uint) (response.Token, error) {
	token := response.Token{}
	path := fmt.Sprintf("/api/token/%d", tokenID)
	err := c.GetDecoded(&token, "TOKEN", path, nil)
	return token, err
}

// GetTokenList filters tokens based on the parameters by invoking the HTTP
// request:
// 	GET /api/token
func (c Client) GetTokenList(params TokenSearch) ([]response.Token, error) {
	tokens := response.PaginatedTokens{}

	q, err := query.Values(params)
	if err != nil {
		return tokens.List, err
	}

	path := "/api/token"
	err = c.GetDecoded(&tokens, "TOKEN", path, q)
	return tokens.List, err
}

// UpdateToken updates the token with the specified ID by invoking the HTTP request:
// 	PUT /api/token
func (c Client) UpdateToken(tokenID uint, token request.TokenUpdate) (response.Token, error) {
	updatedToken := response.Token{}
	body, err := json.Marshal(token)
	if err != nil {
		return updatedToken, err
	}

	path := fmt.Sprintf("/api/token/%d", tokenID)
	err = c.PutDecoded(&updatedToken, "TOKEN", path, nil, body)
	return updatedToken, err
}

// CreateToken adds a new a token by invoking the HTTP request:
// 	POST /api/token
func (c Client) CreateToken(token request.Token) (response.Token, error) {
	newToken := response.Token{}
	body, err := json.Marshal(token)
	if err != nil {
		return newToken, err
	}

	path := "/api/token"
	err = c.PostDecoded(&newToken, "TOKEN", path, nil, body)
	return newToken, err
}
