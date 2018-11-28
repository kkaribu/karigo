package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// Action ...
type Action interface {
	Execute(a *Access)
}

/*
 * GENERIC ACTIONS
 */

// ActionGetCollectionSize ...
func ActionGetCollectionSize(query Query, size *int) func(*Access) {
	return func(acc *Access) {
		*size = acc.Count(query)
	}
}

// ActionGetCollection ...
func ActionGetCollection(query Query, col jsonapi.Collection) func(*Access) {
	return func(acc *Access) {
		data := acc.GetColFields(query)

		baseRes := col.Sample()
		for _, m := range data {
			newRes := baseRes.Copy()
			for f, val := range m {
				newRes.Set(f, val)
			}
			col.Add(newRes)
		}
	}
}

// ActionGetResource ...
func ActionGetResource(query Query, res jsonapi.Resource) func(*Access) {
	return func(acc *Access) {
		data := acc.GetResFields(query)

		for f, val := range data {
			res.Set(f, val)
		}
	}
}

// ActionGetInclusions ...
func ActionGetInclusions(query Query, rels, fields []string, col jsonapi.Collection) func(*Access) {
	return func(acc *Access) {
		data := acc.GetInclusions(query, rels, fields)

		var res jsonapi.Resource
		for id, fields := range data {
			res = col.Sample().Copy()
			res.Set("id", id)
			for field, val := range fields {
				res.Set(field, val)
			}
			col.Add(res)
		}
	}
}

// ActionInsertResource ...
func ActionInsertResource(res jsonapi.Resource) func(*Access) {
	return func(acc *Access) {
		id, typ := res.IDAndType()
		for _, attr := range res.Attrs() {
			acc.Set(typ, id, attr.Name, res.Get(attr.Name))
		}
		for _, rel := range res.Rels() {
			if rel.ToOne {
				acc.SetToOneRel(typ, id, rel.Name, res.GetToOne(rel.Name))
			} else {
				acc.SetToManyRel(typ, id, rel.Name, res.GetToMany(rel.Name)...)
			}
		}
	}
}

// ActionUpdateResource ...
func ActionUpdateResource(typ, id string, vals map[string]interface{}) func(*Access) {
	return func(acc *Access) {
		for field, val := range vals {
			acc.Set(typ, id, field, val)
		}
	}
}

// ActionDeleteResource ...
func ActionDeleteResource(typ, id string) func(*Access) {
	return func(acc *Access) {
		acc.Set(typ, id, "id", "")
	}
}
