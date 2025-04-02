package controller

import (
	"caturilham05/user/helper"
	"caturilham05/user/model/web"
	"caturilham05/user/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

// @Summary Login
// @Description Login with username and password to generate a token
// @Tags User
// @Accept json
// @Produce json
// @Param login body web.UserLoginRequest true "Login Information"
// @Success 200 {object} web.UserLoginResponse
// @Failure 400 {object} web.WebResponse "BAD REQUEST"
// @Router /api/login [post]
func (u *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(r, &userLoginRequest)

	userResponse := u.UserService.Login(r.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param create body web.UserCreateRequest true "User Create Information"
// @Success 200 {object} web.UserResponse
// @Failure 400 {object} web.WebResponse "BAD REQUEST"
// @Failure 409 {object} web.WebResponse "CONFLICT"
// @Router /api/users [post]
func (u *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := u.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Delete User
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} web.WebResponse "Success"
// @Failure 401 {object} web.WebResponse "UNAUTHORIZED"
// @Failure 404 {object} web.WebResponse "NOT FOUND"
// @Router /api/users/{id} [delete]
func (u *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	u.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Get All Users
// @Description Retrieve all users (Requires JWT Token)
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} web.UserResponse
// @Failure 401 {object} web.WebResponse "UNAUTHORIZED"
// @Router /api/users [get]
func (u *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRepsponses := u.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userRepsponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Get User by ID
// @Description Retrieve user details by ID
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} web.UserResponse
// @Failure 400 {object} web.WebResponse "BAD REQUEST"
// @Failure 401 {object} web.WebResponse "UNAUTHORIZED"
// @Failure 404 {object} web.WebResponse "NOT FOUND"
// @Router /api/users/{id} [get]
func (u *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := u.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Update User
// @Description Update an existing user
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param update body web.UserUpdateRequest true "User Update Information"
// @Success 200 {object} web.UserResponse
// @Failure 400 {object} web.WebResponse "BAD REQUEST"
// @Failure 401 {object} web.WebResponse "UNAUTHORIZED"
// @Failure 404 {object} web.WebResponse "NOT FOUND"
// @Failure 409 {object} web.WebResponse "CONFLICT"
// @Router /api/users/{id} [put]
func (u *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := u.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}
