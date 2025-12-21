package memory

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"kpp.dev/minireddit/internal/entity"
)

type PostRepositoryMemory struct {
	posts map[uuid.UUID]*entity.Post
	mutex sync.RWMutex
}

func NewPostRepositoryMemory() *PostRepositoryMemory {
	return &PostRepositoryMemory{
		posts: make(map[uuid.UUID]*entity.Post),
	}
}

func (r *PostRepositoryMemory) GetAll() ([]entity.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	posts := make([]entity.Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, *post)
	}

	return posts, nil
}

func (r *PostRepositoryMemory) Save(post *entity.Post) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if post.ID == uuid.Nil {
		post.ID = uuid.New()
	}

	r.posts[post.ID] = post
	return nil
}

func (r *PostRepositoryMemory) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.posts[id]; !exists {
		return errors.New("post not found")
	}

	delete(r.posts, id)
	return nil
}

func (r *PostRepositoryMemory) Like(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	post, exists := r.posts[id]
	if !exists {
		return errors.New("post not found")
	}

	post.Likes++
	return nil
}

func (r *PostRepositoryMemory) Dislike(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	post, exists := r.posts[id]
	if !exists {
		return errors.New("post not found")
	}

	post.Dislikes++
	return nil
}
