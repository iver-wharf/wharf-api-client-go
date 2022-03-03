package wharfapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstRuneLower(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Leaves as-is on empty",
			input: "",
			want:  "",
		},
		{
			name:  "Leaves as-is on lower-cased",
			input: "foo bar",
			want:  "foo bar",
		},
		{
			name:  "Leaves as-is on all-caps except first rune",
			input: "fOO BAR",
			want:  "fOO BAR",
		},
		{
			name:  "Lower-cases single rune",
			input: "F",
			want:  "f",
		},
		{
			name:  "Lower-cases only first rune",
			input: "FOO BAR",
			want:  "fOO BAR",
		},
		{
			name:  "Lower-cases non-ASCII rune",
			input: "ÄÅ gott",
			want:  "äÅ gott",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := firstRuneLower(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCutString_ok(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 string
		want2 string
	}{
		{
			name:  "Cuts at middle",
			input: "foo bar",
			want1: "foo",
			want2: "bar",
		},
		{
			name:  "Cuts at start",
			input: " bar",
			want1: "",
			want2: "bar",
		},
		{
			name:  "Cuts at end",
			input: "foo ",
			want1: "foo",
			want2: "",
		},
		{
			name:  "Cuts only first",
			input: "foo bar moo",
			want1: "foo",
			want2: "bar moo",
		},
		{
			name:  "Empty results",
			input: " ",
			want1: "",
			want2: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got1, got2, ok := cutString(test.input, ' ')
			assert.True(t, ok, "cut string OK?")
			assert.Equal(t, test.want1, got1, "return value 1")
			assert.Equal(t, test.want2, got2, "return value 2")
		})
	}
}

func TestCutString_notOk(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Empty string",
			input: "",
		},
		{
			name:  "No delimiter",
			input: "foobar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, _, ok := cutString(test.input, ' ')
			assert.False(t, ok, "cut string not OK?")
		})
	}
}
