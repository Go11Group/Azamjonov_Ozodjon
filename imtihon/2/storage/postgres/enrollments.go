package postgres

import (
	"database/sql"
	"github.com/imtihon/model"
	"time"
)

// EnrollmentRepo represents the repository for managing enrollments in PostgreSQL.
type EnrollmentRepo struct {
	db *sql.DB
}

// NewEnrollmentRepo creates a new instance of EnrollmentRepo with the provided database connection.
func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}

// EnrollmentGet retrieves an enrollment by its ID from the database.
func (e *EnrollmentRepo) EnrollmentGet(id string) (model.Enrollments, error) {
	var enrollment model.Enrollments
	err := e.db.QueryRow(`
		SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at 
		FROM enrollments 
		WHERE enrollment_id = $1`, id).Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	return enrollment, err
}

// EnrollmentCreate creates a new enrollment record in the database.
func (e *EnrollmentRepo) EnrollmentCreate(enrollment model.Enrollments) error {
	query := `
		INSERT INTO enrollments (user_id, course_id, enrollment_date, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := e.db.Exec(query, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate, time.Now(), time.Now(), 0)
	return err
}

// EnrollmentUpdate updates an existing enrollment record in the database.
func (e *EnrollmentRepo) EnrollmentUpdate(enrollment model.Enrollments) error {
	query := `
		UPDATE enrollments 
		SET user_id = $1, course_id = $2, enrollment_date = $3, updated_at = $4, deleted_at = $5 
		WHERE enrollment_id = $6`
	_, err := e.db.Exec(query, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate, time.Now(), enrollment.DeletedAt, enrollment.EnrollmentId)
	return err
}

// EnrollmentDelete marks an enrollment as deleted by setting the deleted_at timestamp.
func (e *EnrollmentRepo) EnrollmentDelete(id string) error {
	query := `
		UPDATE enrollments 
		SET deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE enrollment_id = $1 AND deleted_at = 0`
	_, err := e.db.Exec(query, id)
	return err
}

// Get executes a general query to retrieve multiple enrollments based on the provided query and arguments.
func (e *EnrollmentRepo) Get(query string, args []interface{}) ([]model.Enrollments, error) {
	rows, err := e.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollments
	for rows.Next() {
		var enrollment model.Enrollments
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate,
			&enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return enrollments, nil
}

// EnrolledByCourseId retrieves all enrollments associated with a specific course ID.
func (e *EnrollmentRepo) EnrolledByCourseId(courseID string) ([]model.Enrollments, error) {
	query := `
		SELECT e.enrollment_id, e.user_id, e.course_id, e.enrollment_date, e.created_at, e.updated_at, e.deleted_at
		FROM enrollments e
		JOIN courses c ON e.course_id = c.course_id
		WHERE c.course_id = $1 AND e.deleted_at = 0`
	rows, err := e.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollments
	for rows.Next() {
		var enrollment model.Enrollments
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
