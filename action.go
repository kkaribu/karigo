package karigo

import (
	"strings"
)

// ModifyTitleAndAddTag ...
func ModifyTitleAndAddTag(articleID, tagID string) func(*Access) {
	return func(access *Access) {
		title := access.WillReadString("articles." + articleID + ".title")
		_ = access.WillReadString("tags." + tagID + ".id")
		access.WillWrite("articles." + articleID + ".title")
		access.WillWrite("articles." + articleID + ".tags")
		access.Ready()

		// Do some stuff...
		title = strings.ToUpper(title)
		access.SetString("articles."+articleID+".title", title)

		// More stuff
		access.AddToManyRel("articles."+articleID+".tags", "tags."+tagID+".id")
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
		ids := access.WillReadStrings("players.*.id", nil, []string{"score"}, 10, 1)
		access.WillWrite("players.*.rank")
		access.Ready()

		// Do stuff...
		access.Release("players.*.score")

		for i := 0; i < len(ids); i++ {
			access.SetInt("players."+".rank", i)
		}
	}
}
