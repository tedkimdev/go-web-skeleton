package service

import (
	"context"

	"github.com/tedkimdev/go-web-skeleton/internal/model"
)

type PostsRepository interface {
	GetPosts(ctx context.Context) ([]model.Post, error)
	DeletePost(ctx context.Context, postID string) error
	CreatePost(ctx context.Context, title, body string) (*model.Post, error)
}

type PostService struct {
	postsRepository PostsRepository
}

func newPostService(postsRepository PostsRepository) *PostService {
	return &PostService{
		postsRepository: postsRepository,
	}
}

func (s *PostService) GetPosts(ctx context.Context) ([]model.Post, error) {
	return s.postsRepository.GetPosts(ctx)
}

func (s *PostService) DeletePost(ctx context.Context, postID string) error {
	return s.postsRepository.DeletePost(ctx, postID)
}

func (s *PostService) CreatePost(ctx context.Context, title, body string) (*model.Post, error) {
	return s.postsRepository.CreatePost(ctx, title, body)
}
