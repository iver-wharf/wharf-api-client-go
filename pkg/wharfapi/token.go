package wharfapi

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/request"
	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
)

// TokenSearch is used when getting tokens without using a token ID
// through the HTTP request:
//  GET /api/token
type TokenSearch struct {
	Limit         *int     `url:"limit,omitempty"`
	Offset        *int     `url:"offset,omitempty"`
	OrderBy       []string `url:"orderby,omitempty"`
	UserName      *string  `url:"userName,omitempty"`
	UserNameMatch *string  `url:"userNameMatch,omitempty"`
}

// GetToken fetches a token by ID by invoking the HTTP request:
//  GET /api/token/{tokenID}
//
// Added in wharf-api v0.2.2.
func (c *Client) GetToken(tokenID uint) (response.Token, error) {
	if err := c.validateEndpointVersion(0, 2, 2); err != nil {
		return response.Token{}, err
	}
	var token response.Token
	path := fmt.Sprintf("/api/token/%d", tokenID)
	err := c.getUnmarshal(path, nil, &token)
	return token, err
}

// GetTokenList filters tokens based on the parameters by invoking the HTTP
// request:
//  GET /api/token
//
// Added in wharf-api v5.0.0.
func (c *Client) GetTokenList(params TokenSearch) (response.PaginatedTokens, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedTokens{}, err
	}
	var tokens response.PaginatedTokens
	q, err := query.Values(params)
	if err != nil {
		return tokens, err
	}
	path := "/api/token"
	err = c.getUnmarshal(path, q, &tokens)
	return tokens, err
}

// UpdateToken updates the token with the specified ID by invoking the HTTP request:
//  PUT /api/token/{tokenID}
//
// Added in wharf-api v5.0.0.
func (c *Client) UpdateToken(tokenID uint, token request.TokenUpdate) (response.Token, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.Token{}, err
	}
	var updatedToken response.Token
	path := fmt.Sprintf("/api/token/%d", tokenID)
	err := c.putJSONUnmarshal(path, nil, token, &updatedToken)
	return updatedToken, err
}

// CreateToken adds a new a token by invoking the HTTP request:
//  POST /api/token
//
// Added in wharf-api v0.2.0.
func (c *Client) CreateToken(token request.Token) (response.Token, error) {
	if err := c.validateEndpointVersion(0, 2, 0); err != nil {
		return response.Token{}, err
	}
	var newToken response.Token
	path := "/api/token"
	err := c.postJSONUnmarshal(path, nil, token, &newToken)
	return newToken, err
}
