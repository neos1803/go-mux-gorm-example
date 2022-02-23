package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

var users []User

func MockUsers() []User {
	users = append(users, User{
		Id:       1,
		Name:     "John",
		Username: "john",
	})

	users = append(users, User{
		Id:       2,
		Name:     "Doe",
		Username: "doe",
	})

	return users
}
