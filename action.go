package karigo

import (
	"strings"
)

// ModifyTitleAndAddTag ...
func ModifyTitleAndAddTag(articleID, tagID string) func(*Access) NTx {
	return func(access *Access) NTx {
		access.WillRead("articles." + articleID + ".title")
		access.WillRead("tags." + tagID + ".id")
		access.WillWrite("articles." + articleID + ".tags")
		access.Ready()

		// Do some stuff...
		title := access.GetString("articles.abc123.title")
		title = strings.ToUpper(title)
		access.SetString("articles.abc123.title", title)
		access.Release("articles.abc123.title")

		// More stuff
		access.AddToManyRel("articles.abc123.tags", "tags.mytag.id")

		// End the action
		return access.End()
	}
}

// RemoveAllTagsFromArticle ...
func RemoveAllTagsFromArticle(articleID string) func(*Access) NTx {
	return func(access *Access) NTx {
		access.WillRead("articles." + articleID + ".tags")
		access.Ready()

		// Do stuff...
		access.SetToManyRel("articles.abc123.tags", []string{})

		// End the action
		return access.End()
	}
}

// SetRankToTopPlayers ...
func SetRankToTopPlayers(limit int) func(*Access) NTx {
	return func(access *Access) NTx {
		access.WillRead("players.*.score")
		access.Ready()
		access.WillWrite("players.*.rank")

		// Do stuff...
		topPlayers := []map[string]uint{}
		access.GetManyStructs("players", []string{"score"}, nil, []string{"score"}, 10, 1, topPlayers)
		access.Release("players.*.score")

		for i := 0; i < len(topPlayers); i++ {
			access.SetInt("players."+".rank", i)
		}

		// End the action
		return access.End()
	}
}
