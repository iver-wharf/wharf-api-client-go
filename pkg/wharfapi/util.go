package wharfapi

import (
	"fmt"
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
