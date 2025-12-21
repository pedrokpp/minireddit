package repository

import (
	"github.com/google/uuid"
	"kpp.dev/minireddit/internal/entity"
)

type PostRepository interface {
	GetAll() ([]entity.Post, error)

	Save(post *entity.Post) error

	Delete(id uuid.UUID) error
	Like(id uuid.UUID) error
	Dislike(id uuid.UUID) error
}
