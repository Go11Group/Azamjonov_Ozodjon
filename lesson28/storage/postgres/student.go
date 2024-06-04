package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/at_lesson/lesson28/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.Student, error) {
	rows, err := u.Db.Query(`select s.id, s.name, age, gender, c.name from students s
					left join course c on c.id = s.course_id `)
	if err != nil {
		return nil, err
	}

	var users []model.Student
	var user model.Student
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.CourseId)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.Student, error) {
	var user model.Student

	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from students s
					left join course c on c.id = s.course_id where s.id = $1`, id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.CourseId)
	if err != nil {
		return model.Student{}, err
	}

	return user, nil
}

func (u *StudentRepo) Create(user model.Student) error {

	_, err :=u.Db.Exec("insert into students(name, age, gender, course_id) values($1, $2, $3, $4)",
		user.Name, user.Age, user.Gender, user.CourseId)
	if err != nil{
		return err
	}
	return nil
}

func (u *StudentRepo) Update(user model.Student) error {
	q := fmt.Sprintf("update table students set name=%s, age=%d, gender=%s, course_id=%s", user.Name, user.Age, user.Gender, user.CourseId)
	_, err := u.Db.Exec(q)
	if err != nil{
		return err
	}
	return nil
}

func (u *StudentRepo) Delete(id string) error {
	_, err := u.Db.Exec("delete from students where id = $1", id)
	if err != nil{
		return err
	}
	return nil
}