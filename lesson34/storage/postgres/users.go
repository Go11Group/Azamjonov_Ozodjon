package postgres

import (
	"database/sql"
	"strconv"
	"time"
	"users_and_produncts/model"
)

type Users struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) *Users {
	return &Users{DB: db}
}

func (u *Users) CreateUser(user model.User) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}

	defer tr.Commit()

	query := `insert into users(username, email, password) values ($1, $2, $3)`
	_, err = tr.Exec(query, user.Username, user.Email, user.Password)
	return err
}

func (u *Users) GetUser(filter model.Filter) (*[]model.User, error) {
	query := `select * from users`

	isEmptyParams := true
	sch := 1
	params := []interface{}{}
	if filter.Id != nil {
		query += ` where`
		query += ` id=$` + strconv.Itoa(sch)
		params = append(params, filter.Id)
		isEmptyParams = false
		sch++
	}
	if filter.Username != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` username=$` + strconv.Itoa(sch)
		params = append(params, filter.Username)
		isEmptyParams = false
		sch++
	}
	if isEmptyParams {
		query += ` where deletedat is null`
	} else {
		query += ` and deletedat is null`
	}

	users := []model.User{}
	rows, err := u.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &users, err
}

func (u *Users) UpdateUser(user *model.User) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update users set username=$1, email=$2, password=$3, updatedAt=$4 where id=$5`
	_, err = tr.Exec(query, user.Username, user.Email, user.Password, time.Now(), user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) DeleteUser(userID int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update users set deletedat=$1 where id=$2`
	_, err = tr.Exec(query, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
