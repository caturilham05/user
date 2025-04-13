package main

import (
	"caturilham05/user/app"
	"caturilham05/user/controller"
	_ "caturilham05/user/docs"
	"caturilham05/user/helper"
	"caturilham05/user/middleware"
	"caturilham05/user/repository"
	"caturilham05/user/service"
	"fmt"
	"net/http"

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
	// fmt.Println(helper.GenerateSecretKey())
	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: middleware.NewAuthMiddleware(router),
		// Handler: router,
	}

	fmt.Println("Server started on 0.0.0.0:3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
