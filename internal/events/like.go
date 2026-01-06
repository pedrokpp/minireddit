package events

import (
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/usecase"
)

type Like struct {
	postRepository repository.PostRepository
	Input          usecase.PostLikeInput
}

func NewLike(input usecase.PostLikeInput, postRepository repository.PostRepository) *Like {
	return &Like{
		postRepository: postRepository,
		Input:          input,
	}
}

func (c *Like) Type() string {
	return "create"
}

func (c *Like) Execute() error {
	like := usecase.NewPostLike(c.postRepository)
	_, err := like.Execute(c.Input)
	if err != nil {
		return err
	}

	return nil
}
