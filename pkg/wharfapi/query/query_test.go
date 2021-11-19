package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func stringToMap(s string) (map[string]string, bool) {
// 	if s == "" {
// 		return nil, false
// 	}

// 	pairs := strings.Split(s, ";")
// 	m := make(map[string]string, len(pairs))
// 	for _, pair := range pairs {
// 		split := strings.SplitN(pair, ":", 2)
// 		k, v := split[0], split[1]
// 		m[k] = v
// 	}

// 	return m, true
// }

func TestStringToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantMap  map[string]string
		wantBool bool
	}{
		{
			name:  "one entry",
			input: "name:testName",
			wantMap: map[string]string{
				"name": "testName",
			},
			wantBool: true,
		},
		{
			name:  "two entries",
			input: "name:testName;requires:OtherField",
			wantMap: map[string]string{
				"name":     "testName",
				"requires": "OtherField",
			},
			wantBool: true,
		},
		{
			name:  "empty entry",
			input: "name:testName;requires:;excluded_with:OtherField",
			wantMap: map[string]string{
				"name":          "testName",
				"excluded_with": "OtherField",
				"requires":      "",
			},
			wantBool: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotMap, gotBool := stringToMap(test.input)
			assert.EqualValues(t, test.wantMap, gotMap)
			assert.Equal(t, test.wantBool, gotBool)
		})
	}
}
