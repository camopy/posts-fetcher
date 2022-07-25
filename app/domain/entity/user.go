package entity

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(id int, name, userName, email string) *User {
	return &User{
		Id:       id,
		Name:     name,
		UserName: userName,
		Email:    email,
	}
}
