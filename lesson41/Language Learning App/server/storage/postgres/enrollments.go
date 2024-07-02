package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"language_learning_app/pkg"
	"time"
)

type EnrollmentRepo struct {
	DB *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{DB: db}
}

// Create
func (e *EnrollmentRepo) CreateEnrollment(newEnrollment *model.Enrollment) error {
	tx, err := e.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into enrollments(user_id, course_id, enrollment_date, 
	created_at) values ($1, $2, $3, $4)`
	_, err = tx.Exec(query, newEnrollment.UserId, newEnrollment.CourseId,
		time.Now(), time.Now())
	return err
}

// Readall
func (e *EnrollmentRepo) GetEnrollments(filter *model.EnrollmentFilter) (*[]model.Enrollment, error) {
	query := `select * from enrollments where deleted_at is null`
	params := []interface{}{}
	paramsCount := 1

	if filter.UserId != nil {
		params = append(params, filter.UserId)
		query += pkg.CreateNewParams(" and user_id = ", paramsCount)
		paramsCount++
	}
	if filter.CourseId != nil {
		params = append(params, filter.CourseId)
		query += pkg.CreateNewParams(" and course_id = ", paramsCount)
		paramsCount++
	}
	if filter.EnrollmentDate != nil {
		params = append(params, filter.EnrollmentDate)
		query += pkg.CreateNewParams(" and enrollment_date = ", paramsCount)
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

	enrollments := &[]model.Enrollment{}
	rows, err := e.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		enrollment := model.Enrollment{}
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId,
			&enrollment.CourseId, &enrollment.EnrollmentDate,
			&enrollment.CreatedAt, &enrollment.UpdatedAt,
			&enrollment.DeletedAt)
		if err != nil {
			return nil, err
		}
		*enrollments = append(*enrollments, enrollment)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return enrollments, nil
}

// Read by id
func (e *EnrollmentRepo) GetEnrollmentById(enrollmentId string) (*model.Enrollment, error) {
	query := `select * from enrollments where enrollment_id = $1`
	enrollment := model.Enrollment{}
	err := e.DB.QueryRow(query, enrollmentId).Scan(&enrollment.EnrollmentId,
		&enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate,
		&enrollment.CreatedAt, &enrollment.UpdatedAt,
		&enrollment.DeletedAt)

	return &enrollment, err
}

// Update
func (e *EnrollmentRepo) UpdateEnrollment(enrollment *model.Enrollment) error {
	tx, err := e.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update enrollments set user_id=$1, course_id=$2, 
	enrollment_date=$3, updated_at=$3 where enrollment_id=$4 and 
	deleted_at is null`

	_, err = tx.Exec(query, enrollment.UserId, enrollment.CourseId,
		enrollment.EnrollmentDate, time.Now(), enrollment.EnrollmentId)
	return err
}

// Delete
func (e *EnrollmentRepo) DeleteEnrollment(enrollmentId string) error {
	tx, err := e.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update enrollments set deleted_at=$1 where enrollment_id=$2
	and deleted_at is null`

	_, err = tx.Exec(query, time.Now(), enrollmentId)
	return err
}
