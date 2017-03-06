package karigo

import (
	"fmt"

	"github.com/dchest/uniuri"
	"github.com/kkaribu/jsonapi"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		panic(fmt.Sprintf("karigo: could not hash password: %s", err))
	}

	return string(hash)
}

// RandomID ...
func RandomID(length int) string {
	return uniuri.NewLen(length)
}

// TableName ...
func TableName(rel jsonapi.Rel) string {
	resType := rel.InverseType
	name := ""

	if rel.InverseName == "" {
		name += fmt.Sprintf("%s_%s", resType, rel.Name)
	} else {
		if resType+rel.Name < rel.Type+rel.InverseName {
			name += fmt.Sprintf("%s_%s_and_%s_%s", resType, rel.Name, rel.Type, rel.InverseName)
		} else if resType+rel.Name > rel.Type+rel.InverseName {
			name += fmt.Sprintf("%s_%s_and_%s_%s", rel.Type, rel.InverseName, resType, rel.Name)
		} else {
			// TODO
			name += fmt.Sprintf("%s_%s_twoway", resType, rel.Name)
		}
	}

	return name
}
