package query_test

import (
	"errors"
	"testing"

	"github.com/iver-wharf/wharf-api-client-go/pkg/wharfapi/query"

	"github.com/stretchr/testify/assert"
)

type ProjectSearch struct {
	OrderBy          []string `query:"name:orderBy"`
	Limit            *int     `query:"name:limit"`
	Offset           *int     `query:"name:offset;requires:Limit"`
	Name             *string  `query:"name:name"`
	GroupName        *string  `query:"name:groupName"`
	Description      *string  `query:"name:description"`
	TokenID          *uint    `query:"name:tokenId"`
	ProviderID       *uint    `query:"name:providerId"`
	GitURL           *string  `query:"name:gitUrl"`
	NameMatch        *string  `query:"name:nameMatch;excluded_with:Name"`
	GroupNameMatch   *string  `query:"name:groupNameMatch;excluded_with:GroupName"`
	DescriptionMatch *string  `query:"name:descriptionMatch;excluded_with:Description"`
	GitURLMatch      *string  `query:"name:gitUrlMatch;excluded_with:GitURL"`
	Match            *string  `query:"name:match"`
}

func intPtr(v int) *int {
	return &v
}

func strPtr(v string) *string {
	return &v
}

func TestQueryFromObj(t *testing.T) {
	tests := []struct {
		name    string
		input   ProjectSearch
		want    string
		wantErr error
	}{
		{
			name:    "empty struct",
			input:   ProjectSearch{},
			want:    "",
			wantErr: nil,
		},
		{
			name: "Slice one element should work",
			input: ProjectSearch{
				OrderBy: []string{"projectId desc"},
			},
			want:    "orderBy=projectId+desc",
			wantErr: nil,
		},
		{
			name: "Slice multiple elements should work",
			input: ProjectSearch{
				OrderBy: []string{"projectId desc", "tokenId asc"},
			},
			want:    "orderBy=projectId+desc&orderBy=tokenId+asc",
			wantErr: nil,
		},
		{
			name: "Limit with offset should work",
			input: ProjectSearch{
				Limit:  intPtr(10),
				Offset: intPtr(10),
			},
			want:    "limit=10&offset=10",
			wantErr: nil,
		},
		{
			name: "Limit without offset should work",
			input: ProjectSearch{
				Limit: intPtr(10),
			},
			want:    "limit=10",
			wantErr: nil,
		},
		{
			name: "Offset without limit should fail",
			input: ProjectSearch{
				Offset: intPtr(10),
			},
			want:    "",
			wantErr: errors.New("field \"Limit\" not set when using \"Offset\""),
		},
		{
			name: "Excluded with is used should fail",
			input: ProjectSearch{
				Name:      strPtr("test_proj"),
				NameMatch: strPtr("test_proj_match"),
			},
			want:    "",
			wantErr: errors.New("field \"Name\" set when using \"NameMatch\""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q, err := query.FromObj(test.input)
			assert.Equal(t, test.wantErr, err)

			got := q.Encode()
			assert.Equal(t, test.want, got)
		})
	}
}
