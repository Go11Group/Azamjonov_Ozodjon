package postgres

import (
	"database/sql"
	"github.com/imtihon/model"
	_ "github.com/lib/pq"
	"time"
)

type CourseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

// Get course by ID
func (c *CourseRepo) Course_Get(id string) (model.Courses, error) {
	var course model.Courses
	err := c.db.QueryRow("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1", id).Scan(
		&course.Id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
	return course, err
}

// Create new course
func (c *CourseRepo) Course_Create(course model.Courses) error {
	query := `INSERT INTO courses (course_id, title, description, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := c.db.Exec(query, course.Id, course.Title, course.Description, time.Now(), time.Now(), nil)
	return err
}

// Update existing course
func (c *CourseRepo) Course_Update(course model.Courses) error {
	query := `UPDATE courses SET title = $1, description = $2, updated_at = $3, deleted_at = $4 WHERE course_id = $5`
	_, err := c.db.Exec(query, course.Title, course.Description, time.Now(), course.DeletedAt, course.Id)
	return err
}

// Delete course by setting deleted_at timestamp
func (c *CourseRepo) Course_Delete(id string) error {
	query := `UPDATE courses SET deleted_at = $1 WHERE course_id = $2`
	_, err := c.db.Exec(query, time.Now(), id)
	return err
}

//                         course_id UUID PRIMARY KEY,
//                         title VARCHAR(100) NOT NULL,
//                         description TEXT,
//                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                         deleted_at TIMESTAMP
