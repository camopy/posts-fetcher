package usecase_test

import (
	"testing"

	"github.com/camopy/posts-fetcher/app/infra/rest"
	"github.com/stretchr/testify/assert"
)

func TestFetchPosts(t *testing.T) {
	rest := rest.New()
	p, err := rest.FetchPosts()
	assert.NoError(t, err)

	c, err := rest.FetchComments()
	assert.NoError(t, err)

	u, err := rest.FetchUsers()
	assert.NoError(t, err)

	got := agregateData(p, c, u)
}
