package karigo

import (
	"strings"
)

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
		acc.SetToManyRel("articles.abc123.tags", []string{})
	}
}

// SetRankToTopPlayers ...
func SetRankToTopPlayers(limit int) func(*Access) {
	return func(acc *Access) {
		ids := acc.GetStrings("players.*.id[score]", nil, 10, 1)
		acc.WillSet("players.*.rank")
		acc.Ready()

		// Do stuff...
		acc.Release("players.*.score")

		for i := 0; i < len(ids); i++ {
			acc.SetInt("players."+".rank", i)
		}
	}
}
