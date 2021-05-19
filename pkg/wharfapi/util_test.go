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
