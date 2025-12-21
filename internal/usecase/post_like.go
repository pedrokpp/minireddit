package usecase

import (
	"github.com/google/uuid"
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/validation"
)

type PostLike struct {
	postRepository repository.PostRepository
}

func NewPostLike(postRepository repository.PostRepository) *PostLike {
	return &PostLike{
		postRepository: postRepository,
	}
}

type PostLikeInput struct {
	ID string `json:"id"`
}

func (i *PostLikeInput) Validate() error {
	if err := validation.UUID(i.ID); err != nil {
		return err
	}

	return nil
}

type PostLikeOutput struct {
	ID string `json:"id"`
}

func (p *PostLike) Execute(input PostLikeInput) (*PostLikeOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}

	if err := p.postRepository.Like(id); err != nil {
		return nil, err
	}

	return &PostLikeOutput{ID: input.ID}, nil
}
