package postgres

import (
	"database/sql"
	"fmt"
	"github.com/imtihon/model"
	"time"
)

// UserRepo represents the repository for managing users in PostgreSQL.
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo creates a new instance of UserRepo with the provided database connection.
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Get executes a query to retrieve users based on a dynamic query string and arguments.
func (u *UserRepo) Get(query string, args []interface{}) ([]model.Users, error) {
	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	users := []model.Users{}
	for rows.Next() {
		user := model.Users{}
		err = rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday,
			&user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetById retrieves a user by their ID from the database.
func (u *UserRepo) GetById(id string) (model.Users, error) {
	var user model.Users
	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1", id).Scan(
		&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return user, err
}

// UserCreate creates a new user record in the database.
func (u *UserRepo) UserCreate(user model.Users) error {
	query := `INSERT INTO users (name, email, birthday, password, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`
	err := u.db.QueryRow(query, user.Name, user.Email, user.Birthday, user.Password, time.Now(), time.Now()).Scan(&user.UserId)
	return err
}

// UserUpdate updates an existing user record in the database.
func (u *UserRepo) UserUpdate(user model.Users) error {
	query := `UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = $5, deleted_at = $6 WHERE user_id = $7`
	_, err := u.db.Exec(query, user.Name, user.Email, user.Birthday, user.Password, time.Now(), user.DeletedAt, user.UserId)
	return err
}

// UserDelete marks a user as deleted by setting the deleted_at timestamp.
func (u *UserRepo) UserDelete(id string) error {
	query := `UPDATE users SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE user_id = $1 AND deleted_at = 0`
	_, err := u.db.Exec(query, id)
	return err
}

// GetByNameOrEmail retrieves a user by their ID from the database.
func (u *UserRepo) GetByNameOrEmail(id string) (model.Users, error) {
	var user model.Users
	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1", id).Scan(
		&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return user, err
}

// SearchUsers retrieves users based on optional name and email filters.
func (u *UserRepo) SearchUsers(name string, email string) ([]model.Users, error) {
	query := `SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at
              FROM users
              WHERE deleted_at = 0`

	args := []interface{}{}
	argIdx := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIdx) // Use ILIKE for case-insensitive search
		args = append(args, "%"+name+"%")
		argIdx++
	}
	if email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", argIdx)
		args = append(args, "%"+email+"%")
		argIdx++
	}

	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
