package api

import (
	"net/http"
)

type Api interface {
	Serve() error
	GetPosts(w http.ResponseWriter, r *http.Request)
}
