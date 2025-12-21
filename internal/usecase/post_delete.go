package usecase

import (
	"github.com/google/uuid"
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/validation"
)

type PostDelete struct {
	postRepository repository.PostRepository
}

func NewPostDelete(postRepository repository.PostRepository) *PostDelete {
	return &PostDelete{
		postRepository: postRepository,
	}
}

type PostDeleteInput struct {
	ID string `json:"id"`
}

func (i *PostDeleteInput) Validate() error {
	if err := validation.UUID(i.ID); err != nil {
		return err
	}

	return nil
}

type PostDeleteOutput struct {
	ID string `json:"id"`
}

func (p *PostDelete) Execute(input PostDeleteInput) (*PostDeleteOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}

	if err := p.postRepository.Delete(id); err != nil {
		return nil, err
	}

	return &PostDeleteOutput{ID: input.ID}, nil
}
