package repository

import (
	"caturilham05/user/helper"
	"caturilham05/user/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserRepositoryImpl struct {
}

// FindByName implements UserRepository.
func (u *UserRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "SELECT `id`, `name`, `username`, `password`, `created_at` FROM `user` WHERE `username` = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer row.Close()

	user := domain.User{}
	var createdAt []uint8
	if row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &createdAt)
		helper.PanicIfError(err)
	}

	if user.Id != 0 {
		createdAtStr := string(createdAt)
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		helper.PanicIfError(err)
		user.CreatedAt = parsedTime
	}

	return user, nil
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM `user` WHERE `id` = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

// FIndAll implements UserRepository.
func (u *UserRepositoryImpl) FIndAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT `id`, `name`, `username`, `created_at` FROM `user` WHERE 1 ORDER BY `id` DESC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	var createdAt []uint8
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &createdAt)
		helper.PanicIfError(err)

		// Konversi dari []uint8 ke string, lalu parsing ke time.Time
		createdAtStr := string(createdAt)
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		helper.PanicIfError(err)

		user.CreatedAt = parsedTime

		users = append(users, user)
	}

	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT `id`, `name`, `username`, `password`, `created_at` FROM `user` WHERE `id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	var createdAt []uint8
	if !rows.Next() {
		return user, errors.New("pengguna tidak ditemukan")
	}

	err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &createdAt)

	helper.PanicIfError(err)

	createdAtStr := string(createdAt)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	helper.PanicIfError(err)
	user.CreatedAt = parsedTime

	return user, nil

}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user(`name`, `username`, `password`) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	user.CreatedAt = time.Now()
	return user
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE `user` SET `name` = ?, `username` = ?, `password` = ? WHERE `id` = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Password, user.Id)
	helper.PanicIfError(err)
	return user
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
