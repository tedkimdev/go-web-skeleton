package posts

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/tedkimdev/go-web-skeleton/internal/model"
)

func Route(router chi.Router, postService PostService) {
	router.Get("/posts", newGetPostsHandler(postService))
	router.Post("/posts", newPostPostsHandler(postService))
	router.Put("/posts/{id}", newPutPostsHandler(postService))
	router.Delete("/posts/{id}", newDeletePostsHandler(postService))
}

type PostService interface {
	CreatePost(ctx context.Context, title, body string) (*model.Post, error)
	GetPosts(ctx context.Context) ([]model.Post, error)
	DeletePost(ctx context.Context, postID string) error
}
