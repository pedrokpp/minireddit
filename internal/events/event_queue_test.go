package events

import (
	"sync"
	"testing"

	"kpp.dev/minireddit/internal/infra/memory"
	"kpp.dev/minireddit/internal/usecase"
)

func TestNewCreate(t *testing.T) {
	eventQueue := NewEventQueue()
	go eventQueue.Loop()

	postRepository := memory.NewPostRepositoryMemory()
	input := usecase.PostCreateInput{
		Title:   "title",
		Content: "content",
		Author:  "author",
	}
	workers := 100
	semaphore := make(chan struct{}, int(workers/2))
	wg := sync.WaitGroup{}
	for range workers {
		wg.Go(func() {
			semaphore <- struct{}{}
			t.Log("creating post event")
			eventQueue.Enqueue(NewCreate(input, postRepository))
			t.Log("create post event queued")
			<-semaphore
		})
	}
	wg.Wait()
	allPosts, err := postRepository.GetAll()
	if err != nil {
		t.Errorf("failed to get all posts: %v", err)
	}
	t.Log("done")
	t.Logf("created %d posts", len(allPosts))
}
