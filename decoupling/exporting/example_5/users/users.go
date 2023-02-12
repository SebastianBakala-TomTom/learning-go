package users

type user struct {
	Name string
	ID   string
}

type Manager struct {
	Title string
	user
}
