package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course model.Course) error {
	_, err := c.DB.Exec("insert into course(name, field) values($1, $2)", course.Name, course.Field)
	if err != nil{
		return err
	}
	return nil
}

func (c *CourseRepo) Delete(id string) error {
	_, err := c.DB.Exec("delete from course where id = $1", id)
	if err != nil{
		return err
	}
	return nil
}

func (c *CourseRepo) Update(course model.Course) error {
	_, err := c.DB.Exec("update table course set name = $1, field = $2 where id = $3", course.Name, course.Field, course.Id)
	if err != nil{
		return err
	}
	return nil
}

func (c *CourseRepo) ShowAllCourse() ([]model.Course, error){
	rows, err := c.DB.Query("select * from course")
	if err != nil{
		return []model.Course{}, err
	}

	courses := []model.Course{}
	for rows.Next(){
		course := model.Course{}
		err = rows.Scan(&course.Id, &course.Name, &course.Field)
		if err != nil{
			return []model.Course{}, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *CourseRepo) GetById(id string) (model.Course, error){
	course := model.Course{}
	
	err := c.DB.QueryRow("select * from course where id = $1", id).
	Scan(&course.Id, &course.Name, &course.Field)
	if err != nil{
		return model.Course{}, err
	}
	return course, nil
}