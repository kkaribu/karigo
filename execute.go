package karigo

import (
	"fmt"

	"github.com/kkaribu/jsonapi"
)

// executeTx ...
func (a *App) executeTx(ctx *Ctx) {
	ctx.AddToLog(fmt.Sprintf("Looking for %s.", ctx.Method+" "+ctx.URL.Route))

	var tx func(acc *Access) error

	if ctx.Method == "GET" {
		tx = func(acc *Access) error {
			if ctx.URL.IsCol {
				// Get collection
				var size int
				err := ActionGetCollectionSize(ctx.Query, &size)(acc)
				if err != nil {
					return err
				}

				col := ctx.App.Collection(ctx.URL.ResType)
				err = ActionGetCollection(ctx.Query, col)(acc)
				if err != nil {
					return err
				}

				ctx.Out.Data = col

				// TODO just to make sure, but is it necessary?
				if ctx.URL.Params.PageSize <= 0 {
					ctx.URL.Params.PageSize = 1000
				}

				totalPages := size / ctx.URL.Params.PageSize
				if size%ctx.URL.Params.PageSize != 0 {
					totalPages++
				}

				// ctx.Out.Options.Meta["total-pages"] = (size / ctx.URL.Params.PageSize) + 1
				if totalPages == 0 {
					totalPages = 1
				}
				ctx.Out.Meta["total-pages"] = totalPages

				pageNumber := ctx.URL.Params.PageNumber

				ctx.Out.Links["self"] = jsonapi.Link{HRef: ctx.URL.NormalizePath()}

				ctx.URL.Params.PageNumber = 1
				ctx.Out.Links["first"] = jsonapi.Link{HRef: ctx.URL.NormalizePath()}

				ctx.URL.Params.PageNumber = pageNumber - 1
				if ctx.URL.Params.PageNumber == 0 {
					ctx.URL.Params.PageNumber = 1
				}
				ctx.Out.Links["prev"] = jsonapi.Link{HRef: ctx.URL.NormalizePath()}

				ctx.URL.Params.PageNumber = pageNumber + 1
				if ctx.URL.Params.PageNumber > totalPages {
					ctx.URL.Params.PageNumber = totalPages
				}
				ctx.Out.Links["next"] = jsonapi.Link{HRef: ctx.URL.NormalizePath()}

				ctx.URL.Params.PageNumber = totalPages
				ctx.Out.Links["last"] = jsonapi.Link{HRef: ctx.URL.NormalizePath()}

				ctx.URL.Params.PageNumber = pageNumber
			} else {
				// Get resource
				res := ctx.App.Resource(ctx.URL.ResType)
				err := ActionGetResource(ctx.Query, res)(acc)
				if err != nil {
					return err
				}

				ctx.Out.Data = res
			}

			// Get inclusions
			for _, inc := range ctx.URL.Params.Include {
				for i := 0; i < len(inc); i++ {
					path := inc[:i+1]
					rels := make([]string, 0, len(path))
					var typ string
					for _, p := range path {
						rels = append(rels, p.Name)
						typ = p.Type
					}
					inclusions := ctx.App.Registry.Collection(typ)
					err := ActionGetInclusions(ctx.Query, rels, ctx.URL.Params.Fields[typ], inclusions)(acc)
					if err != nil {
						return err
					}
				}
			}

			return nil
		}
	} else if ctx.Method == "POST" {
		tx = func(acc *Access) error {
			res := ctx.In.Data.(jsonapi.Resource)
			err := ActionInsertResource(res)(acc)
			if err != nil {
				return err
			}

			return nil
		}
	} else if ctx.Method == "PATCH" {
		tx = func(acc *Access) error {
			vals := map[string]interface{}{}
			err := ActionUpdateResource(ctx.URL.ResType, ctx.URL.ResID, vals)(acc)
			if err != nil {
				return err
			}

			return nil
		}
	} else if ctx.Method == "DELETE" {
		tx = func(acc *Access) error {
			err := ActionDeleteResource(ctx.URL.ResType, ctx.URL.ResID)(acc)
			if err != nil {
				return err
			}

			return nil
		}
	}

	ctx.Tx = func(acc *Access) error {
		tx(acc)
		return nil
	}
}
