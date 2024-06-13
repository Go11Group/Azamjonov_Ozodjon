package postgres

import (
	"database/sql"
	"github.com/imtihon/model"
	"time"
)

type EnrollmentRepo struct {
	db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}

// Get enrollment by ID
func (e *EnrollmentRepo) Enrollment_Get(id string) (model.Enrollments, error) {
	var enrollment model.Enrollments
	err := e.db.QueryRow("SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at FROM enrollments WHERE enrollment_id = $1", id).Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	return enrollment, err
}

// Create new enrollment
func (e *EnrollmentRepo) Enrollment_Create(enrollment model.Enrollments) error {
	query := `INSERT INTO enrollments (enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := e.db.Exec(query, enrollment.EnrollmentId, enrollment.UserId, enrollment.CourseId, time.Now(), time.Now(), enrollment.DeletedAt)
	return err
}

// Update existing enrollment
func (e *EnrollmentRepo) Enrollment_Update(enrollment model.Enrollments) error {
	query := `UPDATE enrollments SET user_id = $1, course_id = $2, enrollment_date = $3, updated_at = $4, deleted_at = $5 WHERE enrollment_id = $6`
	_, err := e.db.Exec(query, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate, time.Now(), enrollment.DeletedAt, enrollment.EnrollmentId)
	return err
}

// Delete enrollment by setting deleted_at timestamp
func (e *EnrollmentRepo) Enrollment_Delete(id string) error {
	query := `UPDATE enrollments SET deleted_at = $1 WHERE enrollment_id = $2`
	_, err := e.db.Exec(query, time.Now(), id)
	return err
}

//                             enrollment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//                             user_id UUID REFERENCES Users(user_id) ON DELETE CASCADE,
//                             course_id UUID REFERENCES Courses(course_id) ON DELETE CASCADE,
//                             enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                             deleted_at TIMESTAMP
