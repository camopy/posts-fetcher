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
	users map[int]*entity.User
}

func New() (*BlogFetcher, error) {
	userMap, err := createUsersMap()
	if err != nil {
		return nil, err
	}
	return &BlogFetcher{
		users: userMap,
	}, nil
}

func createUsersMap() (map[int]*entity.User, error) {
	m := map[int]*entity.User{}
	res, err := http.Get(usersEndpoint)
	if err != nil {
		return nil, err
	}
	users := []*entity.User{}
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, err
	}
	for _, u := range users {
		m[u.Id] = u
	}
	return m, nil
}

func (f *BlogFetcher) FetchPosts(start, size int) ([]*entity.Post, error) {
	p, err := fetchPosts()
	if err != nil {
		return nil, err
	}
	return f.filterPosts(p, start, size)
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

func (f *BlogFetcher) filterPosts(posts []*entity.Post, start, size int) ([]*entity.Post, error) {
	filtered := []*entity.Post{}
	for _, p := range posts {
		if isFromQueryRange(start, size, p.Id) {
			c, err := f.fetchComments(p.Id)
			if err != nil {
				return nil, err
			}
			p.Comments = c
			p.User = f.users[p.UserId]
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}

func isFromQueryRange(start, size, value int) bool {
	return value > start && value <= (start+size)
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
