package entity

type Post struct {
	Id       int        `json:"id"`
	UserId   int        `json:"userId"`
	Title    string     `json:"title"`
	Body     string     `json:"body"`
	User     *User      `json:"user"`
	Comments []*Comment `json:"comments"`
}
