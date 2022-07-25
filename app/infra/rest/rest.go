package rest

import (
	"net/http"
)

type Rest struct {
	client http.Client
}

func New() *Rest {
	return &Rest{
		client: *http.DefaultClient,
	}
}
