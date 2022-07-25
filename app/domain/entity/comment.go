package entity

type Comment struct {
	Id     int    `json:"id"`
	PostId int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func NewComment(id, postId int, name, email, body string) *Comment {
	return &Comment{
		Id:     id,
		PostId: postId,
		Name:   name,
		Email:  email,
		Body:   body,
	}
}
