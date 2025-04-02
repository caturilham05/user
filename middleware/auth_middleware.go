package middleware

import (
	"caturilham05/user/helper"
	"caturilham05/user/model/web"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	accessToken := request.Header.Get("Authorization")

	if strings.HasPrefix(request.URL.Path, "/api/login") {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	if accessToken == "" {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	_, err := helper.ValidateToken(accessToken)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	middleware.Handler.ServeHTTP(writer, request)
}
