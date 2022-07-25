package api

import "github.com/camopy/posts-fetcher/app/domain/entity"

type Api interface {
	FetchPosts() ([]*entity.Post, error)
	FetchComments()
	FetchUsers()
}
