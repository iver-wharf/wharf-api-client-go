package wharfapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/problem"
)

var redacted = "*REDACTED*"
var tokenPatternJSON = regexp.MustCompile(`("token"\s*:\s*"([a-zA-Z\d\s]+)")\s*`)
var tokenReplacementJSON = fmt.Sprintf(`"token":"%s"`, redacted)

func newRequest(method, authHeader, baseURL, path string, q url.Values, body io.Reader) (*http.Request, error) {
	u, err := newURL(baseURL, path, q)
	if err != nil {
		return nil, err
	}
	return newRequestFromURL(method, authHeader, u, body)
}

func newURL(baseURL, path string, q url.Values) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path
	u.RawQuery = q.Encode()
	return u, nil
}

func newRequestFromURL(method, authHeader string, u *url.URL, body io.Reader) (*http.Request, error) {
	urlStr := u.String()
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		log.Error().WithError(err).Message("Failed preparing HTTP request.")
		return nil, err
	}

	if authHeader != "" {
		req.Header.Add("Authorization", authHeader)
	}

	return req, nil
}

func doRequest(req *http.Request) (io.ReadCloser, error) {
	client := &http.Client{}
	response, err := client.Do(req)

	var redactedURL = redactTokenInURL(req.URL.String())
	var withRequestMeta = func(ev logger.Event) logger.Event {
		return ev.
			WithString("method", req.Method).
			WithString("url", redactedURL)
	}
	log.Debug().WithFunc(withRequestMeta).Message("")

	if err != nil {
		log.Error().WithFunc(withRequestMeta).
			WithError(err).
			Message("Failed sending HTTP request.")
		return nil, err
	}

	if isNonSuccessful(response.StatusCode) {
		defer response.Body.Close()
		if response.StatusCode == http.StatusUnauthorized {
			log.Error().WithFunc(withRequestMeta).
				WithInt("status", response.StatusCode).
				Message("Unauthorized.")
			realm := response.Header.Get("WWW-Authenticate")
			return nil, &AuthError{realm}
		}

		if problem.IsHTTPResponse(response) {
			prob, err := problem.ParseHTTPResponse(response)
			if err != nil {
				return nil, fmt.Errorf("unexpected status code returned: %s: %w", response.Status, err)
			}
			return nil, prob
		}

		log.Warn().WithFunc(withRequestMeta).
			WithInt("status", response.StatusCode).
			WithString("Content-Type", response.Header.Get("Content-Type")).
			Messagef("Non-2xx should have responded with a Content-Type of %q.", problem.HTTPContentType)
		return nil, fmt.Errorf("unexpected status code returned: %s", response.Status)
	}

	return response.Body, nil
}

func redactTokenInJSON(src string) string {
	if !tokenPatternJSON.MatchString(src) {
		return src
	}

	return tokenPatternJSON.ReplaceAllString(src, tokenReplacementJSON)
}

func redactTokenInURL(urlStr string) string {
	if urlStr == "" {
		return ""
	}

	uri, err := url.Parse(urlStr)
	if err != nil {
		log.Warn().
			WithError(err).
			WithString("action", "parse URL").
			Message("Unable to redact token from URL.")
		return ""
	}

	params, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.Warn().
			WithError(err).
			WithString("action", "parse query").
			Message("Unable to redact token from URL.")
		return ""
	}

	token := params.Get("Token")
	if token != "" {
		params.Set("Token", redacted)
	} else {
		token = params.Get("token")
		if token != "" {
			params.Set("token", redacted)
		}
	}

	uri.RawQuery = params.Encode()
	newURLStr := uri.String()

	sanitized, err := url.PathUnescape(newURLStr)
	if err != nil {
		log.Warn().
			WithString("action", "unescape path").
			WithError(err).
			WithString("newURL", newURLStr).
			Message("Unable to redact token from URL.")
		return newURLStr
	}

	return sanitized
}

func isNonSuccessful(statusCode int) bool {
	return statusCode < 200 || statusCode >= 300
}
