package wharfapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/iver-wharf/wharf-core/pkg/problem"
)

const problemContentType = "application/problem+json"

func newProblemError(r problem.Response) error {
	trimmedTitle := strings.TrimRight(firstRuneLower(r.Title), ",.!; ")
	return ProblemError{
		Message: fmt.Sprintf("%s: %s", trimmedTitle, strings.Join(r.Errors, "; ")),
		Problem: r,
	}
}

func isProblemResponse(r *http.Response) bool {
	return r.Header.Get("Content-Type") == problemContentType
}

func parseProblemResponse(response *http.Response) (problem.Response, error) {
	resp, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return problem.Response{}, fmt.Errorf(
			"failed to read problem response body: %w", readErr)
	}
	if closeErr := response.Body.Close(); closeErr != nil {
		return problem.Response{}, fmt.Errorf(
			"failed to close problem response body reading: %w", closeErr)
	}
	var prob problem.Response
	if jsonErr := json.Unmarshal(resp, &prob); jsonErr != nil {
		return problem.Response{}, fmt.Errorf(
			`failed to parse "%s" problem response: %w`, problemContentType, jsonErr)
	}
	return prob, nil
}

// ProblemError is a class that conforms with the "error" interface and is
// returned by any HTTP request for which a problem response was sent.
//
// Since v4.0.0 of wharf-api and v1.3.0 of wharf-api-client-go all endpoints
// will return IETF RFC-7807 compatible problem errors whenever there is a
// non-2xx response.
type ProblemError struct {
	Problem problem.Response
	Message string
}

func (e ProblemError) Error() string {
	return e.Message
}
