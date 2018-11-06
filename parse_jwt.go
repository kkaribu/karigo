package karigo

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var (
	publicKey []byte
)

func init() {
	publicKey = []byte("PUBLICKEY")
}

// parseJWT parses a JWT if one is found in a request's header and stores it in
// a context.
func (a *App) parseJWT(ctx *Ctx, r *http.Request) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err == nil {
		ctx.JWT = token
	}

	// Claims
	idInter := claim(ctx.JWT, "id")
	id := ""
	if i, ok := idInter.(string); ok {
		id = i
	}
	ctx.ID = id

	groupsInter := claim(ctx.JWT, "groups")
	groups := []string{}
	if iS, ok := groupsInter.([]interface{}); ok {
		for _, gI := range iS {
			if g, ok := gI.(string); ok {
				groups = append(groups, g)
			}
		}
	}
	ctx.Groups = groups
}

// Claim claims a value from a JWT stored in a context and returns it.
func claim(token *jwt.Token, claim string) interface{} {
	if token != nil {
		if _, found := token.Claims[claim]; found {
			return token.Claims[claim]
		}
	}

	return nil
}
