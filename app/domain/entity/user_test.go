package entity_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/camopy/posts-fetcher/app/domain/entity"
)

func TestNewUser(t *testing.T) {
	type input struct {
		id       int
		name     string
		email    string
		userName string
	}
	tests := []struct {
		name  string
		input input
		want  *entity.User
	}{
		{
			name: "User 1",
			input: input{
				id:       1,
				name:     "Title",
				email:    "email.com",
				userName: "Body",
			},
			want: &entity.User{
				Id:       1,
				Name:     "Title",
				Email:    "email.com",
				UserName: "Body",
			},
		},
		{
			name: "User 2",
			input: input{
				id:       2,
				name:     "Title2",
				email:    "email.com",
				userName: "Body2",
			},
			want: &entity.User{
				Id:       2,
				Name:     "Title2",
				Email:    "email.com",
				UserName: "Body2",
			},
		},
	}

	for _, tt := range tests {
		got := entity.NewUser(tt.input.id, tt.input.name, tt.input.userName, tt.input.email)
		if !reflect.DeepEqual(got, tt.want) {
			log.Fatalf("Test %s - Got %v - Want %v", tt.name, got, tt.want)
		}
	}
}
