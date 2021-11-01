package user

import (
	"Fin-BEReview/constants"
	model "Fin-BEReview/model"
	"Fin-BEReview/service"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddFundHandler(s *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := new(model.FinRequest)
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

		// err = s.UserService.AddFund(user, int64(req.Amount))
		// if err != nil {
		// 	return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
		// 		Success:  false,
		// 		Response: "Add fund failed",
		// 		Error:    []string{err.Error()},
		// 	})
		// }
		request, err := json.Marshal(model.FinRequest{
			UserID: user.UserID.String(),
			Amount: req.Amount,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "Marshal request fail",
				Error:    []string{err.Error()},
			})
		}
		job := &model.Job{
			Key:   constants.FinAddFundJobKey,
			Value: string(request),
		}
		err = s.JobService.AddJob(job)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "add job failed",
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

		req := new(model.FinRequest)
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
		request, err := json.Marshal(model.FinRequest{
			UserID: user.UserID.String(),
			Amount: req.Amount,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "Marshal request fail",
				Error:    []string{err.Error()},
			})
		}
		job := &model.Job{
			Key:   constants.FinWithdrawJobKey,
			Value: string(request),
		}
		err = s.JobService.AddJob(job)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, model.StandardResponse{
				Success:  false,
				Response: "add job failed",
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
