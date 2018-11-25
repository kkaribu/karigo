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

// ActionGetCollection ...
func ActionGetCollection(key Key, col jsonapi.Collection) func(*Access) {
	return func(acc *Access) {
		data := acc.GetColFields(key)

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
func ActionGetResource(key Key, res jsonapi.Resource) func(*Access) {
	return func(acc *Access) {
		data := acc.GetResFields(key)

		for f, val := range data {
			res.Set(f, val)
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
