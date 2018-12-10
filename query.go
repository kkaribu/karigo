package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// Query represents a resource or a collection.
//
// To get a resource, only Type and ID are used.
// To get a collection, all fields except ID are considered.
//
// Filter is not necessary.
// Fields will default to a slice containing the string "id"
// if nil or empty.
type Query struct {
	Set             string
	ID              string
	Fields          []string
	BelongsToFilter jsonapi.BelongsToFilter
	Filter          *jsonapi.Condition
	Sort            []string
	PageSize        int
	PageNumber      int
}

// NewQuery creates a new *Query object from a *jsonapi.URL object.
func NewQuery(url *jsonapi.URL) *Query {
	var fields []string
	if f, ok := url.Params.Fields[url.ResType]; ok {
		fields = make([]string, len(f))
		copy(fields, f)
	} else {
		fields = []string{"id"}
	}

	query := &Query{
		Set:             url.ResType,
		ID:              url.ResID,
		Fields:          fields,
		BelongsToFilter: url.BelongsToFilter,
		Filter:          url.Params.Filter,
		Sort:            url.Params.SortingRules,
		PageSize:        url.Params.PageSize,
		PageNumber:      url.Params.PageNumber,
	}

	return query
}

// NewQueryFromKey ...
func NewQueryFromKey(key Key) *Query {
	var fields []string
	if key.Field != "" {
		fields = make([]string, 1)
		fields[0] = key.Field
	} else {
		fields = []string{}
	}

	return &Query{
		Set:    key.Set,
		ID:     key.ID,
		Fields: fields,
	}
}

// String ...
// TODO Implement Query.String
func (k *Query) String() string {
	return ""
}
