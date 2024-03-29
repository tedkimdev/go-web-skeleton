package openapi

import "github.com/tedkimdev/go-web-skeleton/internal/model"

func ModelPostToDTO(post *model.Post) *Post {
	dto := Post{
		Id:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}

	return &dto
}

func ModelPostsToDTO(posts []model.Post) []*Post {
	dto := make([]*Post, 0)
	for _, post := range posts {
		p := Post{
			Id:    post.ID,
			Title: post.Title,
			Body:  post.Body,
		}
		dto = append(dto, &p)
	}

	return dto
}
