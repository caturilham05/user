package service

import (
	"caturilham05/user/exception"
	"caturilham05/user/helper"
	"caturilham05/user/model/domain"
	"caturilham05/user/model/web"
	"caturilham05/user/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

// Login implements UserService.
func (u *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := u.UserRepository.FindByName(ctx, tx, request.Username)
	if err != nil {
		helper.PanicIfError(err)
	}

	if err := helper.PasswordVerify(user.Password, request.Password); err != nil {
		panic(exception.NewBadRequestError("password tidak valid"))
	}

	token, err := helper.GenerateToken(user.Id, request.Username)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	userResponse := domain.User{
		Id:        user.Id,
		Name:      user.Name,
		Username:  request.Username,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

	return helper.ToLoginResponse(userResponse)
}

// Create implements UserService.
func (u *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userCheck, err := u.UserRepository.FindByName(ctx, tx, request.Username)
	if err != nil {
		helper.PanicIfError(err)
	}

	if userCheck.Id != 0 {
		panic(exception.NewConflictError("pengguna dengan " + request.Username + " sudah terdaftar"))
	}

	passwordHash, err := helper.HashPassword(request.Password)
	if err != nil {
		panic(exception.NewConflictError(err.Error()))
	}

	if request.Password != request.PasswordConfirm {
		panic(exception.NewConflictError("password tidak sesuai"))
	}

	user := domain.User{
		Name:            request.Name,
		Username:        request.Username,
		Password:        passwordHash,
		PasswordConfirm: request.PasswordConfirm,
	}

	user = u.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := u.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	u.UserRepository.Delete(ctx, tx, user)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := u.UserRepository.FIndAll(ctx, tx)
	return helper.ToUserReponses(users)
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := u.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

// Update implements UserService.
func (u *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// check user exist
	user, err := u.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// check username exist in another user
	userCheck, err := u.UserRepository.FindByName(ctx, tx, request.Username)
	if err != nil {
		helper.PanicIfError(err)
	}

	if userCheck.Id != user.Id {
		panic(exception.NewConflictError("pengguna dengan " + request.Username + " sudah terdaftar"))
	}

	// hash password
	var passwordHash string
	if request.Password != "" {
		if request.Password != request.PasswordConfirm {
			panic(exception.NewConflictError("password tidak sesuai"))
		}
		passwordHash, err = helper.HashPassword(request.Password)
		if err != nil {
			panic(exception.NewConflictError(err.Error()))
		}
	} else {
		passwordHash = user.Password // Gunakan password lama
	}

	user.Name = helper.DefaultIfEmpty(request.Name, user.Name)
	user.Username = helper.DefaultIfEmpty(request.Username, user.Username)
	user.Password = helper.DefaultIfEmpty(passwordHash, user.Password)

	user = u.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}
