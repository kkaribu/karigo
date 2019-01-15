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
func ActionGetCollectionSize(query Query, size *int) ActionFunc {
	return ActionFunc(func(acc *Access) error {
		*size = acc.Count(query)

		return nil
	})
}

// ActionGetCollection ...
func ActionGetCollection(query Query, col jsonapi.Collection) ActionFunc {
	return ActionFunc(func(acc *Access) error {
		data := acc.GetCol(query)

		baseRes := col.Sample()
		for _, m := range data {
			newRes := baseRes.Copy()
			for f, val := range m {
				newRes.Set(f, val)
			}
			col.Add(newRes)
		}

		return nil
	})
}

// ActionGetResource ...
func ActionGetResource(query Query, res jsonapi.Resource) ActionFunc {
	return ActionFunc(func(acc *Access) error {
		data := acc.GetRes(query)

		for f, val := range data {
			res.Set(f, val)
		}

		return nil
	})
}

// ActionGetInclusions ...
func ActionGetInclusions(query Query, rels, fields []string, col jsonapi.Collection) ActionFunc {
	return ActionFunc(func(acc *Access) error {
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
	})
}

// ActionInsertResource ...
func ActionInsertResource(res jsonapi.Resource) ActionFunc {
	return ActionFunc(func(acc *Access) error {
		id := res.GetID()
		typ := res.GetType()
		for _, attr := range res.Attrs() {
			acc.Do(Op{Set: typ, ID: id, Field: attr.Name, Op: "set", Val: res.Get(attr.Name)})
		}
		for _, rel := range res.Rels() {
			if rel.ToOne {
				acc.Do(Op{Set: typ, ID: id, Field: rel.Name, Op: "set", Val: res.GetToOne(rel.Name)})
			} else {
				acc.Do(Op{Set: typ, ID: id, Field: rel.Name, Op: "set", Val: res.GetToMany(rel.Name)})
			}
		}

		return nil
	})
}

// ActionUpdateResource ...
func ActionUpdateResource(typ, id string, vals map[string]interface{}) ActionFunc {
	return ActionFunc(func(acc *Access) error {
		for field, val := range vals {
			acc.Do(Op{Set: typ, ID: id, Field: field, Op: "set", Val: val})
		}

		return nil
	})
}

// ActionDeleteResource ...
func ActionDeleteResource(typ, id string) ActionFunc {
	return ActionFunc(func(acc *Access) error {
		acc.Do(Op{Set: typ, ID: id, Field: "id", Op: "set", Val: ""})

		return nil
	})
}

// ActionUnimplemented ...
func ActionUnimplemented() ActionFunc {
	return ActionFunc(func(acc *Access) error {
		return jsonapi.NewErrNotImplemented()
	})
}
