package postgres

import (
	"database/sql"
	"github.com/imtihon/model"
	_ "github.com/lib/pq"
	"time"
)

// CourseRepo represents the repository for managing courses in PostgreSQL.
type CourseRepo struct {
	db *sql.DB
}

// NewCourseRepo creates a new instance of CourseRepo with the provided database connection.
func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

// CourseGet retrieves a course by its ID from the database.
func (c *CourseRepo) CourseGet(id string) (model.Courses, error) {
	var course model.Courses
	err := c.db.QueryRow("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1", id).Scan(
		&course.Id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
	return course, err
}

// CourseCreate creates a new course record in the database.
func (c *CourseRepo) CourseCreate(course model.Courses) error {
	query := `INSERT INTO courses (title, description, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := c.db.Exec(query, course.Title, course.Description, time.Now(), time.Now(), 0)
	return err
}

// CourseUpdate updates an existing course record in the database.
func (c *CourseRepo) CourseUpdate(course model.Courses) error {
	query := `UPDATE courses SET title = $1, description = $2, updated_at = $3, deleted_at = $4 WHERE course_id = $5`
	_, err := c.db.Exec(query, course.Title, course.Description, time.Now(), course.DeletedAt, course.Id)
	return err
}

// CourseDelete marks a course as deleted by setting the deleted_at timestamp.
func (c *CourseRepo) CourseDelete(id string) error {
	query := `update courses set deleted_at = date_part('epoch', current_timestamp)::INT where course_id = $1 and deleted_at = 0`
	_, err := c.db.Exec(query, id)
	return err
}

// Get executes a query to retrieve courses based on a dynamic query string and arguments.
func (c *CourseRepo) Get(query string, args []interface{}) ([]model.Courses, error) {
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	courses := []model.Courses{}
	for rows.Next() {
		course := model.Courses{}
		err = rows.Scan(&course.Id, &course.Title, &course.Description,
			&course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// GetCoursesByUserID retrieves courses enrolled by a specific user from the database.
func (c *CourseRepo) GetCoursesByUserID(userID string) ([]model.Courses, error) {
	query := `
		SELECT c.course_id, c.title, c.description, c.created_at, c.updated_at, c.deleted_at
		FROM courses c
		JOIN enrollments e ON c.course_id = e.course_id
		WHERE e.user_id = $1 AND c.deleted_at = 0`
	rows, err := c.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Courses
	for rows.Next() {
		var course model.Courses
		err := rows.Scan(&course.Id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// GetPopularCourses retrieves courses with the most enrollments within a specified date range.
func (c *CourseRepo) GetPopularCourses(startDate, endDate string) ([]model.PopularCourse, error) {
	query := `
		SELECT c.course_id, c.title AS course_title, COUNT(e.enrollment_id) AS enrollments_count
		FROM courses c
		JOIN enrollments e ON c.course_id = e.course_id
		WHERE e.enrollment_date BETWEEN $1 AND $2
		GROUP BY c.course_id, c.title
		ORDER BY enrollments_count DESC
	`
	rows, err := c.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.PopularCourse
	for rows.Next() {
		var course model.PopularCourse
		err := rows.Scan(&course.CourseID, &course.CourseTitle, &course.EnrollmentsCount)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
