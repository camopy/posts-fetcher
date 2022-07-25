package fetchposts

import (
	"github.com/camopy/posts-fetcher/app/application/service"
	"github.com/camopy/posts-fetcher/app/domain/entity"
)

type FetchPosts struct {
	blogFetcher service.BlogFetcher
}

func New(blogFetcher service.BlogFetcher) *FetchPosts {
	return &FetchPosts{
		blogFetcher: blogFetcher,
	}
}

type Input struct {
	Start int
	Size  int
}

func (f *FetchPosts) Execute(input Input) ([]*entity.Post, error) {
	return f.blogFetcher.FetchPosts(input.Start, input.Size)
}
