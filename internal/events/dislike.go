package events

import (
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/usecase"
)

type Dislike struct {
	postRepository repository.PostRepository
	Input          usecase.PostDislikeInput
}

func NewDislike(input usecase.PostDislikeInput, postRepository repository.PostRepository) *Dislike {
	return &Dislike{
		postRepository: postRepository,
		Input:          input,
	}
}

func (c *Dislike) Type() string {
	return "create"
}

func (c *Dislike) Execute() error {
	dislike := usecase.NewPostDislike(c.postRepository)
	_, err := dislike.Execute(c.Input)
	if err != nil {
		return err
	}

	return nil
}
