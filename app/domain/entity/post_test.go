package entity_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/camopy/posts-fetcher/app/domain/entity"
)

func TestNewPost(t *testing.T) {
	type input struct {
		id     int
		userId int
		title  string
		body   string
	}
	tests := []struct {
		name  string
		input input
		want  *entity.Post
	}{
		{
			name: "Post 1",
			input: input{
				id:     1,
				userId: 1,
				title:  "Title",
				body:   "Body",
			},
			want: &entity.Post{
				Id:     1,
				PostId: 1,
				Title:  "Title",
				Body:   "Body",
			},
		},
		{
			name: "Post 2",
			input: input{
				id:     2,
				userId: 2,
				title:  "Title2",
				body:   "Body2",
			},
			want: &entity.Post{
				Id:     2,
				PostId: 2,
				Title:  "Title2",
				Body:   "Body2",
			},
		},
	}

	for _, tt := range tests {
		got := entity.NewPost(tt.input.id, tt.input.userId, tt.input.title, tt.input.body)
		if !reflect.DeepEqual(got, tt.want) {
			log.Fatalf("Test %s - Got %v - Want %v", tt.name, got, tt.want)
		}
	}
}
