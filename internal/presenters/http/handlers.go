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

func RegisterHandlers(e *echo.Echo, postRepository repository.PostRepository) {
	eventQueue := events.NewEventQueue()
	e.GET("/health", Health())
	e.POST("/posts", CreatePost(eventQueue, postRepository))
	e.GET("/posts", GetAll(postRepository))
	go eventQueue.Loop()
}
