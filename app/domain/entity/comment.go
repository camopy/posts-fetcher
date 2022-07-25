package entity

type Comment struct {
	Id     int
	PostId int
	Name   string
	Email  string
	Body   string
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
