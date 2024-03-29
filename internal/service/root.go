package service

import (
	"sync"

	"github.com/tedkimdev/go-web-skeleton/internal/repository"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository

	PostService *PostService
}

func NewService(repository *repository.Repository) *Service {
	if repository == nil {
		panic("repository is nil")
	}

	serviceInit.Do(func() {
		postService := newPostService(repository.PostsRepository)

		serviceInstance = &Service{
			repository:  repository,
			PostService: postService,
		}
	})

	return serviceInstance
}
