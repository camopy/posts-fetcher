package fetchposts_test

import (
	"testing"

	fetchposts "github.com/camopy/posts-fetcher/app/application/usecase/fetchPosts"
	blogfetcher "github.com/camopy/posts-fetcher/app/infra/blogFetcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchPosts(t *testing.T) {
	blogFetcher, err := blogfetcher.New()
	require.NoError(t, err)

	fetchPosts := fetchposts.New(blogFetcher)

	input := fetchposts.Input{Start: 10, Size: 10}
	posts, err := fetchPosts.Execute(input)

	require.NoError(t, err)
	assert.Len(t, posts, 10)
	for _, post := range posts {
		assert.Greater(t, post.Id, input.Start)
		assert.LessOrEqual(t, post.Id, input.Start+input.Size)
	}
}
