package karigo

// Action ...
type Action interface {
	Execute(ctx Ctx, sw Switch)
}

// ActionFunc ...
type ActionFunc func(Switch) error

// ActionDemo ...
func ActionDemo() ([]string, ActionFunc) {
	deps := []string{
		"users",
		"articles.abc123.title",
		"comments.*.author",
	}

	f := func(sw Switch) error {
		col, _ := sw.GetCol("users", nil, [2]int{0, 0}, nil, nil)

		usersLocks := []string{}
		for i := 0; i < col.Len(); i++ {
			id, _ := col.Elem(i).IDAndType()
			usersLocks = append(usersLocks, "users."+id)
		}

		sw.NarrowDown("users", usersLocks)

		return nil
	}

	return deps, f
}
