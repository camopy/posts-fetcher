package entity

type Post struct {
	Id       int `json:"id"`
	UserId   int
	User     *User
	Comments []*Comment
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func NewPost(id, userId int, title, body string) *Post {
	return &Post{
		Id:    id,
		Title: title,
		Body:  body,
	}
}
