package usecase

import (
	"kpp.dev/minireddit/internal/entity"
	"kpp.dev/minireddit/internal/repository"
)

type PostCreate struct {
	postRepository repository.PostRepository
}

func NewPostCreate(postRepository repository.PostRepository) *PostCreate {
	return &PostCreate{
		postRepository: postRepository,
	}
}

type PostCreateInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (i *PostCreateInput) Validate() error {
	return nil
}

type PostCreateOutput struct {
	ID string `json:"id"`
}

func (p *PostCreate) Execute(input PostCreateInput) (*PostCreateOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	post := entity.NewPost(input.Title, input.Content, input.Author)
	if err := p.postRepository.Save(post); err != nil {
		return nil, err
	}

	return &PostCreateOutput{ID: post.ID.String()}, nil
}
