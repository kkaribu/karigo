package karigo

import (
	"errors"
	"fmt"

	"github.com/kkaribu/jsonapi"
)

// Kernel ...
type Kernel func(ctx *Ctx) error

// executeKernel ...
func (a *App) executeKernel(ctx *Ctx) {
	ctx.AddToLog(fmt.Sprintf("Looking for %s.", ctx.Method+" "+ctx.URL.Route))

	if kernel, ok := a.Kernels[ctx.Method+" "+ctx.URL.Route]; ok {
		ctx.AddToLog("Kernel found.")
		err := kernel(ctx)
		if err != nil {
			panic(jsonapi.NewErrInternal())
		}

		// Pagination
		if ctx.URL.IsCol && ctx.Method == "GET" {
			var size int
			size, err = ctx.App.Store.CountCollectionSize(ctx.Tx, ctx.URL.ResType, ctx.URL.FromFilter, ctx.URL.Params)
			if err != nil {
				panic(jsonapi.NewErrInternal())
			}

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
		}

		// Inclusions
		inclusions, err := ctx.App.Store.SelectInclusions(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter, ctx.URL.Params)
		if err != nil {
			panic(jsonapi.NewErrInternal())
		}
		for _, inc := range inclusions {
			// id, typ := inc.IDAndType()
			// fmt.Printf("About to include %s %s\n\n\n", typ, id)
			ctx.Out.Include(inc)
		}
	} else {
		ctx.AddToLog("Kernel not found.")
		panic(jsonapi.NewErrNotFound())
	}
}

/*
 * DEFAULT KERNELS
 */

// KernelGetCollection ...
func KernelGetCollection(ctx *Ctx) error {
	col := ctx.App.Collection(ctx.URL.ResType)

	// Collection
	err := ctx.App.Store.SelectCollection(ctx.Tx, ctx.URL.ResType, ctx.URL.FromFilter, ctx.URL.Params, col)
	if err != nil {
		return err
	}

	ctx.Out.Data = col

	// body, err := jsonapi.Marshal(res, ctx.URL, ctx.Options)

	return nil
}

// KernelGetResource ...
func KernelGetResource(ctx *Ctx) error {
	res := ctx.App.Resource(ctx.URL.ResType)

	// Resource
	err := ctx.App.Store.SelectResource(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter, ctx.URL.Params, res)
	if err != nil {
		return err
	}

	ctx.Out.Data = res

	// body, err := jsonapi.Marshal(res, ctx.URL, ctx.Options)

	return nil
}

// KernelInsertResource ...
func KernelInsertResource(ctx *Ctx) error {
	res := ctx.In.Data.(jsonapi.Resource)

	errs := res.Validate()
	if len(errs) > 0 {
		return jsonapi.NewErrBadRequest()
	}

	err := ctx.App.Store.InsertResource(ctx.Tx, res)
	if err != nil {
		panic(err)
	}

	return nil
}

// KernelUpdateResource ...
func KernelUpdateResource(ctx *Ctx) error {
	res := ctx.In.Data.(jsonapi.Resource)

	errs := res.Validate()
	if len(errs) > 0 {
		return jsonapi.NewErrBadRequest()
	}

	err := ctx.App.Store.InsertResource(ctx.Tx, res) // TODO
	if err != nil {
		panic(err)
	}

	return nil
}

// KernelDeleteResource ...
func KernelDeleteResource(ctx *Ctx) error {
	err := ctx.App.Store.DeleteResource(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID)
	if err != nil {
		return errors.New("karigo: resource could not be deleted")
	}

	return nil
}

// KernelGetRelationship ...
func KernelGetRelationship(ctx *Ctx) error {
	rel, err := ctx.App.Store.SelectRelationship(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter.Type)
	if err != nil {
		return err
	}

	ctx.Out.Data = jsonapi.Identifier{
		Type: ctx.URL.FromFilter.Type,
		ID:   rel,
	}

	// body, err := jsonapi.Marshal(jsonapi.NewIdentifiers(ctx.URL.Rel.Type, []string{rel}), ctx.URL, ctx.Options)

	return err
}

// KernelGetRelationships ...
func KernelGetRelationships(ctx *Ctx) error {
	// fmt.Printf("REL: %+v\n", ctx.URL.Rel)
	rels, err := ctx.App.Store.SelectRelationships(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter.Type)
	if err != nil {
		return err
	}
	// fmt.Printf("LEN(RELS): %d\n", len(rels))

	ctx.Out.Data = jsonapi.NewIdentifiers(ctx.URL.FromFilter.Type, rels)

	// body, err := jsonapi.Marshal(jsonapi.NewIdentifiers(ctx.URL.Rel.Type, rels), ctx.URL, ctx.Options)

	return err
}

// KernelInsertRelationships ...
func KernelInsertRelationships(ctx *Ctx) error {
	relIDs := ctx.In.Data.(jsonapi.Identifiers)

	err := ctx.App.Store.InsertRelationships(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter.Type, relIDs.IDs())
	if err != nil {
		return err
	}

	return nil
}

// KernelUpdateRelationship ...
func KernelUpdateRelationship(ctx *Ctx) error {
	relID := ctx.In.Data.(jsonapi.Identifier)

	err := ctx.App.Store.UpdateRelationship(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter.Type, relID.ID)
	if err != nil {
		return err
	}

	return nil
}

// KernelUpdateRelationships ...
func KernelUpdateRelationships(ctx *Ctx) error {
	relIDs := ctx.In.Data.(jsonapi.Identifiers)

	err := ctx.App.Store.UpdateRelationships(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter.Type, relIDs.IDs())
	if err != nil {
		return err
	}

	return nil
}

// KernelDeleteRelationships ...
func KernelDeleteRelationships(ctx *Ctx) error {
	relIDs := ctx.In.Data.(jsonapi.Identifiers)

	err := ctx.App.Store.DeleteRelationships(ctx.Tx, ctx.URL.ResType, ctx.URL.ResID, ctx.URL.FromFilter.Name, relIDs.IDs())
	if err != nil {
		return err
	}

	return nil
}
