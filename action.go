package karigo

import (
	"strings"

	"github.com/kkaribu/jsonapi"
)

// Action ...
type Action interface {
	Execute(a *Access)
}

/*
 * GENERIC ACTIONS
 */

// ActionGetRelatedIDs ...
func ActionGetRelatedIDs(typ, id, rel string, ids *[]string) func(*Access) {
	return func(acc *Access) {
		ids2 := acc.GetToManyRel(typ + "." + id + "." + rel)
		*ids = ids2
	}
}

// ActionGetCollection ...
func ActionGetCollection(typ string,
	btf jsonapi.BelongsToFilter,
	fields []string,
	filter *jsonapi.Condition,
	sort []string,
	pageSize,
	pageNumber int,
	v interface{},
) func(*Access) {
	return func(acc *Access) {
		ids := acc.GetString("typ.{reltyp.relid.relname}.{field1,fields2}[sort,sort2,-sort3]:10:1")
		v = ids
	}
}

// ActionInsertResource ...
func ActionInsertResource(res jsonapi.Resource) func(*Access) {
	return func(acc *Access) {
		id, typ := res.IDAndType()
		acc.WillSet(typ)
		acc.Ready()

		acc.SetString(typ+"."+id+".id", id)
		for _, attr := range res.Attrs() {
			acc.Set(typ+"."+attr.Name, res.Get(attr.Name))
		}
		for _, rel := range res.Rels() {
			if rel.ToOne {
				acc.SetToOneRel(typ+"."+rel.Name, res.GetToOne(rel.Name))
			} else {
				acc.SetToManyRel(typ+"."+rel.Name, res.GetToMany(rel.Name)...)
			}
		}
	}
}

/*
 * EXAMPLES
 */

// ModifyTitleAndAddTags ...
func ModifyTitleAndAddTags(articleID string, tagIDs ...string) func(*Access) {
	return func(acc *Access) {
		title := acc.GetString("articles." + articleID + ".title")
		for _, id := range tagIDs {
			_ = acc.GetString("tags." + id + ".id")
		}
		acc.WillSet("articles." + articleID + ".title")
		acc.WillSet("articles." + articleID + ".tags")
		acc.Ready()

		// Do some stuff...
		title = strings.ToUpper(title)
		acc.SetString("articles."+articleID+".title", title)

		// More stuff
		acc.AddToManyRel("articles."+articleID+".tags", tagIDs...)
	}
}

// RemoveAllTagsFromArticle ...
func RemoveAllTagsFromArticle(articleID string) func(*Access) {
	return func(acc *Access) {
		acc.WillSet("articles." + articleID + ".tags")
		acc.Ready()

		// Do stuff...
		acc.SetToManyRel("articles.abc123.tags", []string{}...)
	}
}

// SetRankToTopPlayers ...
func SetRankToTopPlayers(limit int) func(*Access) {
	return func(acc *Access) {
		ids := acc.GetStrings("players.*.id", nil, nil, 10, 1)
		acc.WillSet("players.*.rank")
		acc.Ready()

		// Do stuff...
		acc.Release("players.*.score")

		for i := 0; i < len(ids); i++ {
			acc.SetInt("players."+".rank", i)
		}
	}
}
