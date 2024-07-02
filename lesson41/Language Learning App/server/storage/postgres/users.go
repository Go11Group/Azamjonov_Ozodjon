package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"language_learning_app/pkg"
	"time"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// Create
func (u *UserRepo) CreateUser(newUser *model.User) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into users(name, email, birthday, password, created_at)
		values ($1, $2, $3, $4, $5)`
	_, err = tx.Exec(query, newUser.Name, newUser.Email, newUser.Birthday,
		newUser.Password, time.Now())
	return err
}

// Readall
func (u *UserRepo) GetUsers(filter *model.UserFilter) (*[]model.User, error) {
	query := `select * from users where deleted_at is null`
	params := []interface{}{}
	paramsCount := 1

	if filter.Name != nil {
		params = append(params, filter.Name)
		query += pkg.CreateNewParams(" and name = ", paramsCount)
		paramsCount++
	}
	if filter.Email != nil {
		params = append(params, filter.Email)
		query += pkg.CreateNewParams(" and email = ", paramsCount)
		paramsCount++
	}
	if filter.AgeFrom != nil {
		params = append(params, filter.AgeFrom)
		query += pkg.CreateNewParams(" and extract(year from age(current_timestamp, birthday)) >= ", paramsCount)
		paramsCount++
	}
	if filter.AgeTo != nil {
		params = append(params, filter.AgeTo)
		query += pkg.CreateNewParams(" and extract(year from age(current_timestamp, birthday)) <= ", paramsCount)
		paramsCount++
	}
	if filter.Limit != nil {
		params = append(params, filter.Limit)
		query += pkg.CreateNewParams(" limit ", paramsCount)
		paramsCount++
	}
	if filter.Offset != nil {
		params = append(params, filter.Offset)
		query += pkg.CreateNewParams(" offset ", paramsCount)
		paramsCount++
	}

	users := &[]model.User{}
	rows, err := u.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.UserId, &user.Name,
			&user.Email, &user.Birthday,
			&user.Password, &user.CreatedAt,
			&user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		*users = append(*users, user)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return users, nil
}

// Read by id
func (u *UserRepo) GetUserById(userId string) (*model.User, error) {
	query := `select * from users where user_id = $1`
	user := model.User{}
	err := u.DB.QueryRow(query, userId).Scan(&user.UserId,
		&user.Name, &user.Email, &user.Birthday,
		&user.Password, &user.CreatedAt,
		&user.UpdatedAt, &user.DeletedAt)

	return &user, err
}

// Update
func (u *UserRepo) UpdateUser(user *model.User) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update users set name=$1, email=$2, birthday=$3, 
	password=$4, updated_at=$5 
	where user_id=$6 and deleted_at is null`

	_, err = tx.Exec(query, user.Name, user.Email, user.Birthday,
		user.Password, time.Now(), user.UserId)
	return err
}

// Delete
func (u *UserRepo) DeleteUser(userId string) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update users set deleted_at=$1 where user_id=$2
	and deleted_at is null`

	_, err = tx.Exec(query, time.Now(), userId)
	return err
}

func (u *UserRepo) GetUsersCourses(userId string) (*[]model.Course, error) {
	query := `  select
					c.course_id,
					c.title,
					c.description
				from
					enrollments as e 
				inner join
					courses as c
						on c.course_id = e.course_id
				where
					user_id = $1;`

	courses := []model.Course{}
	rows, err := u.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &courses, nil
}
