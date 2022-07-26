package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/camopy/posts-fetcher/app/application/service"
	"github.com/go-chi/chi"
)

const START = "start"
const SIZE = "size"
const PORT = "8085"

type Rest struct {
	blogFetcher service.BlogFetcher
	server      *http.Server
}

func New(blogFetcher service.BlogFetcher) *Rest {
	rest := &Rest{
		blogFetcher: blogFetcher,
	}
	rest.server = &http.Server{
		Addr:         fmt.Sprintf(":%s", PORT),
		Handler:      rest.NewRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return rest
}

func (rest *Rest) NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/posts", rest.GetPosts)
	return r
}

func (rest *Rest) Serve() {
	fmt.Printf("listening on port %s", PORT)
	if err := rest.server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func (rest *Rest) GetPosts(w http.ResponseWriter, r *http.Request) {
	start, err := parseStartFromURL(r)
	if err != nil {
		fmt.Printf("can not parse start: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	size, err := parseSizeFromURL(r)
	if err != nil {
		fmt.Printf("can not parse size: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p, err := rest.blogFetcher.FetchPosts(int(start), int(size))
	if err != nil {
		fmt.Printf("can not fetch posts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(p) == 0 {
		fmt.Printf("no posts found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		fmt.Printf("can not encode posts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func parseStartFromURL(r *http.Request) (uint64, error) {
	return strconv.ParseUint(r.URL.Query().Get(START), 0, 64)
}

func parseSizeFromURL(r *http.Request) (uint64, error) {
	return strconv.ParseUint(r.URL.Query().Get(SIZE), 0, 64)
}
