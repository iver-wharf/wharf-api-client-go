package wharfapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var redacted = "*REDACTED*"
var tokenPatternJSON = regexp.MustCompile(`("token"\s*:\s*"([a-zA-Z\d\s]+)")\s*`)
var tokenReplacementJSON = fmt.Sprintf(`"token":"%s"`, redacted)

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
		log.WithError(err).Warningln("Unable to redact token from URL: parse URL")
		return ""
	}

	params, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		log.WithError(err).Warningln("Unable to redact token from URL: parse query")
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
		log.WithError(err).WithField("new URL string", newURLStr).Warningln("Unable to redact token from URL: unescape path")
		return newURLStr
	}

	return sanitized
}

func doRequest(from string, method string, URLStr string, body []byte, authHeader string) (*io.ReadCloser, error) {
	log.WithFields(log.Fields{
		"method": method,
		"body":   redactTokenInJSON(string(body)),
		"url":    redactTokenInURL(URLStr),
	}).Debugln(from)

	req, err := http.NewRequest(method, URLStr, bytes.NewReader(body))
	if err != nil {
		log.WithError(err).Errorln("Unable to prepare http request")
		return nil, err
	}

	if authHeader != "" {
		req.Header.Add("Authorization", authHeader)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.WithError(err).Errorln("Unable to send http request")
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		response.Body.Close()
		log.WithField("response", response).Errorln("Unauthorized")
		realm := response.Header.Get("WWW-Authenticate")
		return nil, &AuthError{realm}
	}

	if response.StatusCode < 200 && response.StatusCode >= 300 {
		resp, err2 := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err2 == nil {
			log.WithFields(log.Fields{
				"status":        response.Status,
				"response body": string(resp)}).Debug("Got invalid status code")
		}
		return nil, fmt.Errorf("unexpected status code returned %v", response.StatusCode)
	}

	return &response.Body, nil
}
