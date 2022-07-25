package service

import "github.com/camopy/posts-fetcher/app/domain/entity"

type BlogFetcher interface {
	FetchPosts(int, int) (map[int]*entity.Post, error)
}
