package user

import "math/rand"

type User struct {
	Id int
	Name string
	Email string
}

func (u *User) UpdateUserEmail(newEmail string) {
	u.Email = newEmail
}

func CreateUser(id int, name string, email string) User {
	user := User{
		Id: rand.Intn(100),
		Name: name,
		Email: email,
	}

	return user
}
