package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// Action ...
type Action interface {
	Execute(a *Access) error
}

// ActionFunc ...
type ActionFunc func(*Access) error

// Execute ...
func (f ActionFunc) Execute(acc *Access) error {
	return f(acc)
}

/*
 * GENERIC ACTIONS
 */

// ActionGetCollectionSize ...
func ActionGetCollectionSize(query Query, size *int) func(*Access) error {
	return func(acc *Access) error {
		*size = acc.Count(query)

		return nil
	}
}

// ActionGetCollection ...
func ActionGetCollection(query Query, col jsonapi.Collection) func(*Access) error {
	return func(acc *Access) error {
		data := acc.GetColFields(query)

		baseRes := col.Sample()
		for _, m := range data {
			newRes := baseRes.Copy()
			for f, val := range m {
				newRes.Set(f, val)
			}
			col.Add(newRes)
		}

		return nil
	}
}

// ActionGetResource ...
func ActionGetResource(query Query, res jsonapi.Resource) func(*Access) error {
	return func(acc *Access) error {
		data := acc.GetResFields(query)

		for f, val := range data {
			res.Set(f, val)
		}

		return nil
	}
}

// ActionGetInclusions ...
func ActionGetInclusions(query Query, rels, fields []string, col jsonapi.Collection) func(*Access) error {
	return func(acc *Access) error {
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

		return nil
	}
}

// ActionInsertResource ...
func ActionInsertResource(res jsonapi.Resource) func(*Access) error {
	return func(acc *Access) error {
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

		return nil
	}
}

// ActionUpdateResource ...
func ActionUpdateResource(typ, id string, vals map[string]interface{}) func(*Access) error {
	return func(acc *Access) error {
		for field, val := range vals {
			acc.Set(typ, id, field, val)
		}

		return nil
	}
}

// ActionDeleteResource ...
func ActionDeleteResource(typ, id string) func(*Access) error {
	return func(acc *Access) error {
		acc.Set(typ, id, "id", "")

		return nil
	}
}
