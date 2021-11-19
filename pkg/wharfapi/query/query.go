package query

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func FromObj(obj interface{}) (url.Values, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Struct {
		return nil, errors.New("object is not a struct")
	}

	fields := make(map[string]map[string]string)
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv := val.FieldByName(field.Name)
		if fv.IsNil() {
			continue
		}
		tag := field.Tag.Get("query")
		m, ok := stringToMap(tag)
		if !ok {
			continue
		}
		fields[field.Name] = m
	}

	q := url.Values{}
	for fieldName, tags := range fields {
		if v, ok := tags["requires"]; ok {
			f := val.FieldByName(v)
			if f.IsNil() || f.IsZero() {
				return nil, fmt.Errorf("field %q not set when using %q", v, fieldName)
			}
		}
		if v, ok := tags["excluded_with"]; ok {
			f := val.FieldByName(v)
			if !f.IsNil() || !f.IsZero() {
				return nil, fmt.Errorf("field %q set when using %q", v, fieldName)
			}
		}
		name := fieldName
		if v, ok := tags["name"]; ok {
			name = v
		}
		f := val.FieldByName(fieldName)
		switch reflect.TypeOf(f.Interface()).Kind() {
		case reflect.Slice, reflect.Array:
			times := f.Len()
			for i := 0; i < times; i++ {
				v := f.Index(i)
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				q.Add(name, fmt.Sprintf("%v", v))
			}
		default:
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			q.Add(name, fmt.Sprintf("%v", f))
		}
	}

	return q, nil
}

func stringToMap(s string) (map[string]string, bool) {
	if s == "" {
		return nil, false
	}
	pairs := strings.Split(s, ";")
	m := make(map[string]string, len(pairs))
	for _, pair := range pairs {
		split := strings.SplitN(pair, ":", 2)
		k, v := split[0], split[1]
		m[k] = v
	}
	return m, true
}

// OrderBy          []string `query:"name:orderBy"`
// Limit            *int     `query:"name:limit"`
// Offset           *int     `query:"name:offset;requires:Limit"`
// Name             *string  `query:"name:name"`
// GroupName        *string  `query:"name:groupName"`
// Description      *string  `query:"name:description"`
// TokenID          *uint    `query:"name:tokenId"`
// ProviderID       *uint    `query:"name:providerId"`
// GitURL           *string  `query:"name:gitUrl"`
// NameMatch        *string  `query:"name:nameMatch;excluded_with:Name"`
// GroupNameMatch   *string  `query:"name:groupNameMatch;excluded_with:GroupName"`
// DescriptionMatch *string  `query:"name:descriptionMatch;excluded_with:Description"`
// GitURLMatch      *string  `query:"name:gitUrlMatch;excluded_with:GitURL"`
// Match            *string  `query:"name:match"`
