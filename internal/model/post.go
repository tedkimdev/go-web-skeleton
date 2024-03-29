package model

type Post struct {
	ID    string  `db:"id"`
	Title *string `db:"title"`
	Body  *string `db:"body"`
}

type CreatePostParam struct {
	Title string `db:"title"`
	Body  string `db:"body"`
}
