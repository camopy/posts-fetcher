package blogfetcher_test

import (
	"fmt"
	"testing"

	blogfetcher "github.com/camopy/posts-fetcher/app/infra/blogFetcher"
	"github.com/stretchr/testify/require"
)

func TestFetchPosts(t *testing.T) {
	fetcher := blogfetcher.New()
	posts, err := fetcher.FetchPosts()
	require.NoError(t, err)
	fmt.Printf("%+v", posts[1].Comments)
}
