package repository

import (
	"caturilham05/user/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FIndAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByName(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
