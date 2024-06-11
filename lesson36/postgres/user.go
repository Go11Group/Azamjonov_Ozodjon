package postgres

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson36/model"
	_ "github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) User_Get(id int) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(
		&user.ID, &user.Username, &user.Email)
	return user, err
}

func (u *UserRepo) User_Create(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	return err
}

func (u *UserRepo) User_Update(id int, user model.User) error {
	_, err := u.db.Exec("UPDATE users SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, id)
	return err
}

func (u *UserRepo) User_Delete(id int) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
