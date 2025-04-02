package app

import (
	"caturilham05/user/controller"
	_ "caturilham05/user/docs" // Import dokumentasi Swagger
	"caturilham05/user/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpSwagger.WrapHandler(w, r)
	})

	router.POST("/api/login", userController.Login)

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId", userController.FindById)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
