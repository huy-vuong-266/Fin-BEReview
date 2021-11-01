package main

import (
	"Fin-BEReview/internal"
	"Fin-BEReview/internal/job"
	"Fin-BEReview/internal/user"
	"Fin-BEReview/model"
	"Fin-BEReview/service"
	"Fin-BEReview/storage"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	service := &service.Service{
		UserService: user.NewServiceUser(),
		JobService:  job.NewServiceJob(),
	}
	e.Validator = &model.CustomValidator{
		Validator: validator.New(),
	}
	handler := internal.NewRouter(e, service)

	err := storage.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected")

	server := http.Server{
		Addr:    "127.0.0.1:8011",
		Handler: handler,
	}

	fmt.Println("Server listen and server at port 8011")
	go func() {
		for {
			service.JobService.ProcessJob(*service)
		}
	}()
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
