package presenters

import (
	"github.com/labstack/echo/v4"
	"kpp.dev/minireddit/internal/presenters/http"
	"kpp.dev/minireddit/internal/repository"
)

func HTTP(postRepository repository.PostRepository) {
	e := echo.New()
	http.RegisterHandlers(e, postRepository)
	e.Logger.Fatal(e.Start(":1323"))
}
