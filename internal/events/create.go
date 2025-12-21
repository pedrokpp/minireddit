package events

import (
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/usecase"
)

type Create struct {
	postRepository repository.PostRepository
	Input          usecase.PostCreateInput
}

func NewCreate(input usecase.PostCreateInput, postRepository repository.PostRepository) *Create {
	return &Create{
		postRepository: postRepository,
		Input:          input,
	}
}

func (c *Create) Type() string {
	return "create"
}

func (c *Create) Execute() error {
	create := usecase.NewPostCreate(c.postRepository)
	_, err := create.Execute(c.Input)
	if err != nil {
		return err
	}

	return nil
}
