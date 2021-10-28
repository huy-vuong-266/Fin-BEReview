package internal

import (
	userapi "Fin-BEReview/internal/user"
	"Fin-BEReview/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, s *service.Service) http.Handler {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	})

	e.POST("/add-fund", userapi.AddFundHandler(s))
	e.POST("/withdraw", userapi.WithdrawHandler(s))
	return e.Server.Handler
}
