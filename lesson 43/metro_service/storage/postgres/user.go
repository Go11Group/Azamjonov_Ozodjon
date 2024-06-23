package postgres

import (
	"atto/models"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *models.CreateUser) error {
	_, err := u.Db.Exec("INSERT INTO users (id, name, phone, age) VALUES ($1, $2, $3, $4)",
		uuid.NewString(), user.Name, user.Phone, user.Age)
	return err
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	var user models.User
	err := u.Db.QueryRow("SELECT id, name, phone, age FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Update(user *models.User) error {
	_, err := u.Db.Exec("UPDATE users SET name = $2, phone = $3, age = $4 WHERE id = $1",
		user.Id, user.Name, user.Phone, user.Age)
	return err
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.Db.Exec("UPDATE users SET deleted_at = $2 WHERE id = $1", id, time.Now())
	return err
}

func (u *UserRepo) GetAll() ([]models.User, error) {
	rows, err := u.Db.Query("SELECT id, name, phone, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
