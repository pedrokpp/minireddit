package usecase

import (
	"github.com/google/uuid"
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/validation"
)

type PostDislike struct {
	postRepository repository.PostRepository
}

func NewPostDislike(postRepository repository.PostRepository) *PostDislike {
	return &PostDislike{
		postRepository: postRepository,
	}
}

type PostDislikeInput struct {
	ID string `json:"id"`
}

func (i *PostDislikeInput) Validate() error {
	if err := validation.UUID(i.ID); err != nil {
		return err
	}
	return nil
}

type PostDislikeOutput struct {
	ID string `json:"id"`
}

func (p *PostDislike) Execute(input PostDislikeInput) (*PostDislikeOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}

	if err := p.postRepository.Dislike(id); err != nil {
		return nil, err
	}

	return &PostDislikeOutput{ID: input.ID}, nil
}
