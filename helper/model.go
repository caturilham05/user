package helper

import (
	"caturilham05/user/model/domain"
	"caturilham05/user/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserReponses(user []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range user {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToLoginResponse(user domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Token:     user.Token,
		ExpiredAt: user.ExpiredAt,
	}
}
