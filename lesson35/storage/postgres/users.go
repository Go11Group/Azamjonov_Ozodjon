package postgres

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson35/model"
)

type UsersRepo struct {
	Db *sql.DB
}

func (u *UsersRepo) Create(user model.Users) error {
	_, err := u.Db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)",
		user.Id, user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersRepo) GetByID(id string) (model.Users, error) {
	row := u.Db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id)
	var user model.Users
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UsersRepo) Update(user model.Users) error {
	_, err := u.Db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3",
		user.Name, user.Age, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersRepo) Delete(id string) error {
	_, err := u.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
