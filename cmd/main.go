package main

import (
	"log"

	blogfetcher "github.com/camopy/posts-fetcher/app/infra/blogFetcher"
	rest "github.com/camopy/posts-fetcher/app/infra/rest"
)

func main() {
	blogFetcher, err := blogfetcher.New()
	if err != nil {
		log.Fatalf("failed to create blog fetcher: %v", err)
	}
	server := rest.New(blogFetcher)
	server.Serve()
}
