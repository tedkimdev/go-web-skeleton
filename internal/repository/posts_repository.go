package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/tedkimdev/go-web-skeleton/internal/model"
)

type PostsRepository struct {
	db *sqlx.DB
}

func newPostsRepository(db *sqlx.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

func (r *PostsRepository) GetPosts(ctx context.Context) ([]model.Post, error) {
	query := "select * from test_posts"
	posts := make([]model.Post, 0)

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post model.Post
		err := rows.StructScan(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostsRepository) DeletePost(ctx context.Context, postID string) error {
	query := "delete from test_posts where id = $1"

	_, err := r.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostsRepository) CreatePost(ctx context.Context, title, body string) (*model.Post, error) {
	query := `
	INSERT INTO test_posts (title, body)
		VALUES ($1, $2)
	RETURNING *
	`

	var post model.Post
	params := model.CreatePostParam{
		Title: title,
		Body:  body,
	}
	err := r.db.GetContext(ctx, &post, query, params)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
