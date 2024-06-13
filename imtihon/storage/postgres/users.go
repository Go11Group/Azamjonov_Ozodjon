package postgres

import (
	"database/sql"
	"fmt"
	"github.com/imtihon/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) User_Get(id string) (model.Users, error) {
	var user model.Users
	var deletedAt sql.NullString // sql.NullString to handle NULL values

	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1", id).Scan(
		&user.Id, // assuming user.Id is directly scanned
		&user.Name,
		&user.Email,
		&user.Birthday,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&deletedAt,
	)
	if err != nil {
		return user, err
	}

	// Assign deleted_at field
	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.String
	} else {
		user.DeletedAt = nil // or "" depending on your logic
	}

	return user, nil
}

func (u *UserRepo) User_Create(user model.Users) error {
	query := `INSERT INTO users (id, name, email, birthday, password, created_at, updated_at, deleted_at)
              VALUES ($1, $2, $3, $4, $5, DEFAULT, DEFAULT, DEFAULT)`
	_, err := u.db.Exec(query, user.Id, user.Name, user.Email, user.Birthday, user.Password)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

// User_Update updates an existing user in the database
func (u *UserRepo) User_Update(user model.Users) error {
	query := `UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = DEFAULT, deleted_at = $5 WHERE id = $6`
	_, err := u.db.Exec(query, user.Name, user.Email, user.Birthday, user.Password, user.DeletedAt, user.Id)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (u *UserRepo) User_Delete(id string) error {
	_, err := u.db.Exec("DELETE FROM users WHERE user_id = $1", id)
	return err
}
