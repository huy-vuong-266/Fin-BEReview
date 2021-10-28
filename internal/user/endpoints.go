package user

import (
	model "Fin-BEReview/model"
	"Fin-BEReview/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddFundHandler(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		type AddFundRequest struct {
			UserID string `json:"user_id" validate:"required"`
			Amount int    `json:"amount" validate:"required"`
		}

		req := new(AddFundRequest)
		c.Bind(req)
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.StandardResponse{
				Success:  false,
				Response: "",
				Error:    []string{err.Error()},
			})
		}

		userRes, err := s.UserService.GetUserByID(req.UserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.StandardResponse{
				Success:  false,
				Response: "Get user failed",
				Error:    []string{err.Error()},
			})
		}

		user, ok := userRes.(*model.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "Parse user failed",
				Error:    []string{err.Error()},
			})
		}

		err = s.UserService.AddFund(user, int64(req.Amount))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "Add fund failed",
				Error:    []string{err.Error()},
			})
		}

		return c.JSON(http.StatusOK, model.StandardResponse{
			Success:  true,
			Response: "Add fund success",
			Error:    []string{},
		})

	}
}

func WithdrawHandler(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		type AddFundRequest struct {
			UserID string `json:"user_id" validate:"required"`
			Amount int    `json:"amount" validate:"required"`
		}

		req := new(AddFundRequest)
		c.Bind(req)
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.StandardResponse{
				Success:  false,
				Response: "",
				Error:    []string{err.Error()},
			})
		}

		userRes, err := s.UserService.GetUserByID(req.UserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.StandardResponse{
				Success:  false,
				Response: "Get user failed",
				Error:    []string{err.Error()},
			})
		}

		user, ok := userRes.(*model.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "Parse user failed",
				Error:    []string{err.Error()},
			})
		}

		if user.Budget < int64(req.Amount) {
			return echo.NewHTTPError(http.StatusBadRequest, model.StandardResponse{
				Success:  false,
				Response: "budget < request amount",
				Error:    []string{"budget < request amount"},
			})
		}
		err = s.UserService.Withdraw(user, int64(req.Amount))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "Add fund failed",
				Error:    []string{err.Error()},
			})
		}

		return c.JSON(http.StatusOK, model.StandardResponse{
			Success:  true,
			Response: "Withdraw success",
			Error:    []string{},
		})

	}
}
