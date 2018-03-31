package karigo

// Action ...
type Action func(ctx Ctx, sw Switch) error

// ActionDemo ...
func ActionDemo(ctx Ctx, sw Switch) error {
	sw.ReserveR("users")
	sw.ReserveR("articles.abc123.title")
	sw.ReserveR("comments.*.author")
	sw.ReserveW("users")
	sw.ReserveW("comments.*.author")
	sw.Ready()

	// Users
	col, _ := sw.GetCol("users", nil, [2]int{0, 0}, nil, nil)

	// Narrow down to the users that will be updated
	usersLocks := []string{}
	for i := 0; i < col.Len(); i++ {
		id, _ := col.Elem(i).IDAndType()
		usersLocks = append(usersLocks, "users."+id)
	}
	sw.NarrowDown("users", usersLocks)

	return nil
}
