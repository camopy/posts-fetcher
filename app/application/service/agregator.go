package service

import "github.com/camopy/posts-fetcher/app/domain/entity"

type Agregator struct {
	Posts    map[int]*entity.Post
	Users    map[int]*entity.User
	Comments map[int][]*entity.Comment
}

func NewAgregator() *Agregator {
	return &Agregator{
		Posts: make(map[int]*entity.Post),
	}
}

func (a *Agregator) AgregateData(posts []*entity.Post, c []*entity.Comment, u []*entity.User) {
	a.createPostsMap(posts)
	a.createUsersMap(u)
	a.createCommentsMap(c)
	a.appendUsers()
	a.app
}

func (a *Agregator) createPostsMap(posts []*entity.Post) {
	for _, p := range posts {
		a.Posts[p.Id] = p
	}
}

func (a *Agregator) createUsersMap(users []*entity.User) {
	for _, u := range users {
		a.Users[u.Id] = u
	}
}

func (a *Agregator) createCommentsMap(comments []*entity.Comment) {
	for _, c := range comments {
		a.Posts[c.PostId].Comments = append(a.Posts[c.PostId].Comments, c)
	}
}

func (a *Agregator) appendUsers() {
	for _, p := range a.Posts {
		p.User = a.Users[p.UserId]
	}
}
