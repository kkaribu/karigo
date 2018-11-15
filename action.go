package karigo

import (
	"strings"
)

// ModifyTitleAndAddTags ...
func ModifyTitleAndAddTags(articleID string, tagIDs ...string) func(*Access) {
	return func(access *Access) {
		title := access.GetString("articles." + articleID + ".title")
		for _, id := range tagIDs {
			_ = access.GetString("tags." + id + ".id")
		}
		access.WillWrite("articles." + articleID + ".title")
		access.WillWrite("articles." + articleID + ".tags")
		access.Ready()

		// Do some stuff...
		title = strings.ToUpper(title)
		access.SetString("articles."+articleID+".title", title)

		// More stuff
		access.AddToManyRel("articles."+articleID+".tags", tagIDs...)
	}
}

// RemoveAllTagsFromArticle ...
func RemoveAllTagsFromArticle(articleID string) func(*Access) {
	return func(access *Access) {
		access.WillWrite("articles." + articleID + ".tags")
		access.Ready()

		// Do stuff...
		access.SetToManyRel("articles.abc123.tags", []string{})
	}
}

// SetRankToTopPlayers ...
func SetRankToTopPlayers(limit int) func(*Access) {
	return func(access *Access) {
		ids := access.GetStrings("players.*.id[score]", nil, []string{"score"}, 10, 1)
		access.WillWrite("players.*.rank")
		access.Ready()

		// Do stuff...
		access.Release("players.*.score")

		for i := 0; i < len(ids); i++ {
			access.SetInt("players."+".rank", i)
		}
	}
}
