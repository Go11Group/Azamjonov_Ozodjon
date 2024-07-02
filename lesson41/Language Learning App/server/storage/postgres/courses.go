package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"language_learning_app/pkg"
	"time"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

// Create
func (c *CourseRepo) CreateCourse(newCourse *model.Course) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into courses(title, description, created_at)
		values ($1, $2, $3)`
	_, err = tx.Exec(query, newCourse.Title, newCourse.Description, time.Now())
	return err
}

// Readall
func (c *CourseRepo) GetCourses(filter *model.CourseFilter) (*[]model.Course, error) {
	query := `select * from courses where deleted_at is null`
	params := []interface{}{}
	paramsCount := 1

	if filter.Title != nil {
		params = append(params, filter.Title)
		query += pkg.CreateNewParams(" and title = ", paramsCount)
		paramsCount++
	}
	if filter.Description != nil {
		params = append(params, filter.Description)
		query += pkg.CreateNewParams(" and description = ", paramsCount)
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

	courses := &[]model.Course{}
	rows, err := c.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.CourseId, &course.Title,
			&course.Description, &course.CreatedAt,
			&course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		*courses = append(*courses, course)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return courses, nil
}

// Read by id
func (c *CourseRepo) GetCourseById(courseId string) (*model.Course, error) {
	query := `select * from courses where course_id = $1`
	course := model.Course{}
	err := c.DB.QueryRow(query, courseId).Scan(&course.CourseId,
		&course.Title, &course.Description, &course.CreatedAt,
		&course.UpdatedAt, &course.DeletedAt)

	return &course, err
}

func (c *CourseRepo) GetLessonsbyCourse(courseId string) (*[]model.Lesson, error) {
	query := `select
					l.lesson_id,
					l.title,
					l.content
				from
					lessons as l 
				where
					course_id = $1`

	lessons := []model.Lesson{}
	rows, err := c.DB.Query(query, courseId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		lesson := model.Lesson{}
		err := rows.Scan(&lesson.LessonId, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &lessons, nil
}

func (c *CourseRepo) GetEnrolledUsersByCourse(courseId string) (*[]model.User, error) {
	query := `select
					u.user_id,
					u.name,
					u.email
				from
					enrollments as e
				inner join
					users as u
						on u.user_id = e.user_id
				where
					course_id = $1;`

	users := []model.User{}

	rows, err := c.DB.Query(query, courseId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (c *CourseRepo) GetMostPopularCourses(startDate, endDate, limit string) (*[]model.PopularCourse, error) {
	query := `with enrollmentsDate as (
					select 
						*
					from
						enrollments
					where
						enrollment_date >= $1::date and 
						enrollment_date <= $2::date
				)

				select
					c.course_id,
					c.title,
					count(*) as enrollments_count
				from
					enrollments as e
				inner join
					courses as c
						on c.course_id = e.course_id
				group by
					c.course_id
				order by 
					count(*) desc
				limit $3;`

	courses := []model.PopularCourse{}
	rows, err := c.DB.Query(query, startDate, endDate, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		course := model.PopularCourse{}
		err := rows.Scan(&course.CourseId, &course.Title, &course.Enrollments_count)
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

// Update
func (c *CourseRepo) UpdateCourse(course *model.Course) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update courses set title=$1, description=$2, updated_at=$3 
	where course_id=$4 and deleted_at is null`

	_, err = tx.Exec(query, course.Title, course.Description, time.Now(),
		course.CourseId)
	return err
}

// Delete
func (c *CourseRepo) DeleteCourse(courseId string) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update courses set deleted_at=$1 where course_id=$2
	and deleted_at is null`

	_, err = tx.Exec(query, time.Now(), courseId)
	return err
}
