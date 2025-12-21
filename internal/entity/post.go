package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID uuid.UUID `json:"id"`

	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`

	Likes    uint `json:"likes"`
	Dislikes uint `json:"dislikes"`

	CreatedAt time.Time `json:"created_at"`
}

func NewPost(title, content, author string) *Post {
	return &Post{
		ID:        uuid.New(),
		Title:     title,
		Content:   content,
		Author:    author,
		Likes:     0,
		Dislikes:  0,
		CreatedAt: time.Now(),
	}
}
