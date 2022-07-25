package api

import "github.com/camopy/posts-fetcher/app/domain/entity"

type Api interface {
	GetPosts() ([]*entity.Post, error)
}
