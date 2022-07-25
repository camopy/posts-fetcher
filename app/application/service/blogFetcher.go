package service

import "github.com/camopy/posts-fetcher/app/domain/entity"

type BlogFetcher interface {
	FetchPosts(int, int) ([]*entity.Post, error)
}
