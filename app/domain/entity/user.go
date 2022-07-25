package entity

type User struct {
	Id       int
	Name     string
	UserName string
	Email    string
}

func NewUser(id int, name, userName, email string) *User {
	return &User{
		Id:       id,
		Name:     name,
		UserName: userName,
		Email:    email,
	}
}
