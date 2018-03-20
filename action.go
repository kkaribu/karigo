package karigo

// Action ...
type Action interface {
	Execute(ctx Ctx, sw Switch)
}
