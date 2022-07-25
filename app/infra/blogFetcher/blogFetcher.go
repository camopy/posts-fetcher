package blogfetcher

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/camopy/posts-fetcher/app/domain/entity"
)

const postsEndpoint = "https://jsonplaceholder.typicode.com/posts"
const usersEndpoint = "https://jsonplaceholder.typicode.com/users"

var commentsEndpoint = "https://jsonplaceholder.typicode.com/posts/%v/comments"

type BlogFetcher struct {
	posts map[int]*entity.Post
	users map[int]*entity.User
}

func New() *BlogFetcher {
	return &BlogFetcher{
		posts: make(map[int]*entity.Post),
		users: make(map[int]*entity.User),
	}
}

func (f *BlogFetcher) FetchPosts() (map[int]*entity.Post, error) {
	p, err := fetchPosts()
	if err != nil {
		return nil, err
	}
	u, err := fetchUsers()
	if err != nil {
		return nil, err
	}
	f.createUsersMap(u)
	f.createPostsMap(p)

	return f.posts, nil
}

func fetchPosts() ([]*entity.Post, error) {
	res, err := http.Get(postsEndpoint)
	if err != nil {
		return nil, err
	}
	p := []*entity.Post{}
	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
		return nil, err
	}
	return p, nil
}

func fetchUsers() ([]*entity.User, error) {
	res, err := http.Get(usersEndpoint)
	if err != nil {
		return nil, err
	}
	u := []*entity.User{}
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, err
	}
	return u, nil
}

func (f *BlogFetcher) createUsersMap(users []*entity.User) {
	for _, u := range users {
		f.users[u.Id] = u
	}
}

func (f *BlogFetcher) createPostsMap(posts []*entity.Post) error {
	for _, p := range posts {
		f.posts[p.Id] = p
		f.posts[p.Id].User = f.users[p.UserId]
		c, err := f.fetchComments(p.Id)
		if err != nil {
			return err
		}
		f.posts[p.Id].Comments = c
	}
	return nil
}

func (r *BlogFetcher) fetchComments(postId int) ([]*entity.Comment, error) {
	url := fmt.Sprintf(commentsEndpoint, postId)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	c := []*entity.Comment{}
	if err := json.NewDecoder(res.Body).Decode(&c); err != nil {
		return nil, err
	}
	return c, nil
}
