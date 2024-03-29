package repository

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	db *sqlx.DB

	PostsRepository *PostsRepository
}

func NewRepoistory(db *sqlx.DB) *Repository {
	if db == nil {
		panic("db is nil")
	}

	repositoryInit.Do(func() {
		postsRepository := newPostsRepository(db)
		repositoryInstance = &Repository{
			db:              db,
			PostsRepository: postsRepository,
		}
	})

	return repositoryInstance
}
