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
	Type            string
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
		Type:            url.ResType,
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

// String ...
// TODO Implement Query.String
func (k *Query) String() string {
	return ""
}
