package wharfapi

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

func firstRuneLower(value string) string {
	if len(value) == 0 {
		return value
	}
	first, size := utf8.DecodeRuneInString(value)
	if first == utf8.RuneError {
		return value
	}
	return fmt.Sprintf("%c%s", unicode.ToLower(first), value[size:])
}

func cutString(str string, delimiter byte) (string, string, bool) {
	idx := strings.IndexByte(str, delimiter)
	if idx == -1 {
		return str, "", false
	}
	return str[:idx], str[idx+1:], true
}

func trimProtocol(v string) string {
	switch {
	case strings.HasPrefix(v, "http://"):
		return strings.TrimPrefix(v, "http://")
	case strings.HasPrefix(v, "https://"):
		return strings.TrimPrefix(v, "https://")
	default:
		return v
	}
}

func decodeJSONAndClose(r io.ReadCloser, obj interface{}) (finalErr error) {
	defer closeAndSetError(r, &finalErr)
	finalErr = json.NewDecoder(r).Decode(obj)
	return
}

// closeAndSetError may be used to set the named return variable inside a
// deferred call if that deferred call failed.
//
// NOTE: The errPtr argument must be a pointer to a named result parameters,
// otherwise it will not affect the calling function's returned value.
func closeAndSetError(closer io.Closer, errPtr *error) {
	closeErr := closer.Close()
	if errPtr != nil && *errPtr == nil {
		*errPtr = closeErr
	}
}
