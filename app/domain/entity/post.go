package entity

type Post struct {
	Id       int `json:"id"`
	UserId   int
	User     *User
	Comments []*Comment
	PostId   int    `json:"post_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func NewPost(id, userId int, title, body string) *Post {
	return &Post{
		Id:     id,
		PostId: userId,
		Title:  title,
		Body:   body,
	}
}
