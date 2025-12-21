package usecase

import (
	"kpp.dev/minireddit/internal/entity"
	"kpp.dev/minireddit/internal/repository"
)

type PostGetAll struct {
	postRepository repository.PostRepository
}

func NewPostGetAll(postRepository repository.PostRepository) *PostGetAll {
	return &PostGetAll{
		postRepository: postRepository,
	}
}

// type PostGetAllInput struct{}

type PostGetAllOutput struct {
	Posts []entity.Post
}

func (p *PostGetAll) Execute() (*PostGetAllOutput, error) {
	posts, err := p.postRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return &PostGetAllOutput{Posts: posts}, nil
}
