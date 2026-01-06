package http

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"kpp.dev/minireddit/internal/events"
	"kpp.dev/minireddit/internal/repository"
	"kpp.dev/minireddit/internal/usecase"
)

func Health() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	}
}

func CreatePost(eventQueue *events.EventQueue, postRepository repository.PostRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := usecase.PostCreateInput{}
		if err := json.NewDecoder(c.Request().Body).Decode(&input); err != nil {
			return err
		}

		createEvent := events.NewCreate(input, postRepository)
		eventQueue.Enqueue(createEvent)

		return c.JSON(http.StatusOK, createEvent.Input)
	}
}

func GetAll(postRepository repository.PostRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		getAll := usecase.NewPostGetAll(postRepository)
		posts, err := getAll.Execute()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, posts)
	}
}

func LikePost(eventQueue *events.EventQueue, postRepository repository.PostRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := usecase.PostLikeInput{
			ID: c.Param("id"),
		}

		likeEvent := events.NewLike(input, postRepository)
		eventQueue.Enqueue(likeEvent)

		return c.JSON(http.StatusOK, likeEvent.Input)
	}
}

func DislikePost(eventQueue *events.EventQueue, postRepository repository.PostRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := usecase.PostDislikeInput{
			ID: c.Param("id"),
		}

		dislikeEvent := events.NewDislike(input, postRepository)
		eventQueue.Enqueue(dislikeEvent)

		return c.JSON(http.StatusOK, dislikeEvent.Input)
	}
}

func RegisterHandlers(e *echo.Echo, postRepository repository.PostRepository) {
	eventQueue := events.NewEventQueue()
	e.GET("/health", Health())
	e.GET("/posts", GetAll(postRepository))
	e.POST("/posts", CreatePost(eventQueue, postRepository))
	e.POST("/posts/:id/like", LikePost(eventQueue, postRepository))
	e.POST("/posts/:id/dislike", DislikePost(eventQueue, postRepository))
	go eventQueue.Loop()
}
