package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
	"my_project/strorage"
	"time"
)

type EnrollmentRepository struct {
	Db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) *EnrollmentRepository {
	return &EnrollmentRepository{Db: db}
}

// create enrollment enrollments larni yaratamiz
func (repo *EnrollmentRepository) CreateEnrollment(enrollment model.Enrollment) error {
	// parse qilib beradi string dan  time.Time
	enrollmentDate, err := time.Parse("2006-01-02", enrollment.EnrollmentDate)
	if err != nil {
		fmt.Println("++++++", err)
		return err
	}
	_, err = repo.Db.Exec("INSERT INTO enrollments(user_id, course_id, enrollment_date) VALUES ($1, $2, $3)", enrollment.UserId, enrollment.CourseId, enrollmentDate)
	return err
}

// course update qiladi id si boyicha enrollment larni update qiladi
func (repo *EnrollmentRepository) UpdateEnrollment(id string, enrollment model.Enrollment) error {
	enrollment1 := model.Enrollment{}
	err := repo.Db.QueryRow("select user_id,course_id ,enrollment_date from enrollments where id=$1 ", id).Scan(&enrollment1.UserId, &enrollment1.CourseId, &enrollment1.EnrollmentDate)
	if err != nil {
		return err
	}
	if len(enrollment.UserId) > 0 {
		enrollment1.UserId = enrollment.UserId
	}
	if len(enrollment.CourseId) > 0 {
		enrollment1.CourseId = enrollment.CourseId
	}
	var time1 time.Time
	if len(enrollment.EnrollmentDate) > 0 {
		enrollmentDate, err := time.Parse("2006-01-02", enrollment.EnrollmentDate)
		if err != nil {
			fmt.Println("++++++", err)
			return err
		}
		time1 = enrollmentDate

	}

	_, err = repo.Db.Exec("update enrollments  set user_id=$1,course_id=$2,enrollment_date=$3 where id =$4", enrollment1.UserId, enrollment1.CourseId, time1, id)
	_, err = repo.Db.Exec("update enrollments  set updated_at=$1 where id =$2", time.Now(), id)
	return err

}

// delete enrollment  - enrollment larni delete qilamiz
func (repo *EnrollmentRepository) DeleteEnrollment(id string) error {
	_, err := repo.Db.Exec("delete from enrollments where id=$1", id)
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err = repo.Db.Exec("update  enrollments set deleted_at=$1 where id=$2", DeletedAt, id)
	return err
}

// get all get va filter limit offset yoki filter qiladi fild lari boyicha

func (repo *EnrollmentRepository) GetEnrollments(f model.Filter) ([]model.Enrollment, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, user_id, course_id,enrollment_date, created_at, updated_at, deleted_at
	 	from enrollments where true`
	filter := ""

	if len(f.UserId) > 0 {
		params["user_id"] = f.UserId
		filter += " and user_id = :user_id "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :course_id "
	}
	if len(f.EnrollmentDate) > 0 {
		params["enrollment_date"] = f.EnrollmentDate
		filter += " and enrollment_date = :enrollment_date "
	}
	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if f.Offset > 0 {
		params["offset"] = f.Offset
		offset = ` OFFSET :offset`
	}

	query = query + filter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)

	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		fmt.Println("____", err)
		return nil, err
	}
	enrollments := []model.Enrollment{}
	for rows.Next() {
		enrollment := model.Enrollment{}
		err := rows.Scan(&enrollment.Id, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil

}

// get by id si boyicha yani idsi boyicha search qiladi
func (repo *EnrollmentRepository) GetById(id string) (*model.Enrollment, error) {
	enrollment := model.Enrollment{}
	err := repo.Db.QueryRow("select id,user_id,course_id,enrollment_date from enrollments where id=$1", id).Scan(&enrollment.Id, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}
