package entity_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/camopy/posts-fetcher/app/domain/entity"
)

func TestNewComment(t *testing.T) {
	type input struct {
		id     int
		postId int
		name   string
		email  string
		body   string
	}
	tests := []struct {
		name  string
		input input
		want  *entity.Comment
	}{
		{
			name: "Comment 1",
			input: input{
				id:     1,
				postId: 1,
				name:   "Title",
				email:  "email@emai.com",
				body:   "Body",
			},
			want: &entity.Comment{
				Id:     1,
				PostId: 1,
				Name:   "Title",
				Email:  "email@emai.com",
				Body:   "Body",
			},
		},
		{
			name: "Comment 2",
			input: input{
				id:     2,
				postId: 2,
				name:   "Title2",
				email:  "email@emai2.com",
				body:   "Body2",
			},
			want: &entity.Comment{
				Id:     2,
				PostId: 2,
				Name:   "Title2",
				Email:  "email@emai2.com",
				Body:   "Body2",
			},
		},
	}

	for _, tt := range tests {
		got := entity.NewComment(tt.input.id, tt.input.postId, tt.input.name, tt.input.email, tt.input.body)
		if !reflect.DeepEqual(got, tt.want) {
			log.Fatalf("Test %s - Got %v - Want %v", tt.name, got, tt.want)
		}
	}
}
