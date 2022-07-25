package rest

import (
	"encoding/json"
	"net/http"

	"github.com/camopy/posts-fetcher/app/domain/entity"
)

const postsEndpoint = "https://jsonplaceholder.typicode.com/posts"
const commentsEndpoint = "https://jsonplaceholder.typicode.com/posts/1/comments"
const usersEndpoint = "https://jsonplaceholder.typicode.com/users"

type Rest struct {
	client http.Client
}

func New() *Rest {
	return &Rest{
		client: *http.DefaultClient,
	}
}

func (r *Rest) FetchPosts() ([]*entity.Post, error) {
	res, err := r.client.Get(postsEndpoint)
	if err != nil {
		return nil, err
	}
	p := []*entity.Post{}
	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Rest) FetchComments() ([]*entity.Comment, error) {
	res, err := r.client.Get(postsEndpoint)
	if err != nil {
		return nil, err
	}
	c := []*entity.Comment{}
	if err := json.NewDecoder(res.Body).Decode(&c); err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Rest) FetchUsers() ([]*entity.User, error) {
	res, err := r.client.Get(postsEndpoint)
	if err != nil {
		return nil, err
	}
	u := []*entity.User{}
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, err
	}
	return u, nil
}
