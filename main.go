package main

import (
	"caturilham05/user/app"
	"caturilham05/user/controller"
	"caturilham05/user/helper"
	"caturilham05/user/middleware"
	"caturilham05/user/repository"
	"caturilham05/user/service"
	"net/http"

	_ "caturilham05/user/docs"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// @title User API
// @version 1.0
// @description This is a sample API with JWT authentication
// @host localhost:3000

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @Security BearerAuth

func main() {
	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
		// Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
