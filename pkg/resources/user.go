package resources

type User struct {
	Email string `schema:"email"`
	Pass string	`schema:"pass"`
}

type Settings struct {
	Users []User
}