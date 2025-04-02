package app

import (
	"caturilham05/user/controller"
	"caturilham05/user/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/login", userController.Login)

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId", userController.FindById)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
