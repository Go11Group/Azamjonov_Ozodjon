package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
	"my_project/strorage"
	"time"
)

type CourseRepository struct {
	Db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{Db: db}
}

// create course bu yerda course larni yaratamiz shu yerda  course kelgan larni table ga yozib qoyadi
func (repo *CourseRepository) CreateCourse(course model.Course) error {
	_, err := repo.Db.Exec("insert into courses(title,description) values ($1,$2)", course.Title, course.Description)
	return err

}

// update course bu yerda course yaratamiz birinchi course larni oqib olamiz agar fild larni bosh kelsa ularni ozgartirmaydi
func (repo *CourseRepository) UpdateCourse(id string, course model.Course) error {
	course1 := model.Course{}
	err := repo.Db.QueryRow("select title ,description  from courses where id=$1", id).Scan(&course1.Title, &course1.Description)
	if err != nil {
		return err
	}
	if len(course.Title) > 0 {
		course1.Title = course.Title
	}
	if len(course.Description) > 0 {
		course1.Description = course.Description
	}
	_, err = repo.Db.Exec("update courses  set title=$1,description=$2 where id =$3", course1.Title, course1.Description, id)
	_, err = repo.Db.Exec("update courses  set updated_at=$1  where id =$2", time.Now(), id)
	return err

}

// delete course bu delete course ni qiliadi

func (repo *CourseRepository) DeleteCourse(id string) error {
	fmt.Println("+++++++++++++++++", id)
	_, err := repo.Db.Exec("delete from courses where id=$1", id)
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err = repo.Db.Exec("update  courses set deleted_at=$1 where id=$2", DeletedAt, id)
	return err
}

// get course yoki filtr bu filter fildlari bilan yoki limit offset
func (repo *CourseRepository) GetCourses(f model.Filter) ([]model.Course, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, title, description, created_at, updated_at, deleted_at
	 	from courses where true`
	filter := ""

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.Description) > 0 {
		params["description"] = f.Description
		filter += " and description = :description "
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
		return nil, err
	}
	courses := []model.Course{}
	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.Id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil

}

// get course lessons bu yerda course ga tegisli bolgan course larni olib keladi
func (repo *CourseRepository) GetLessonByCourseId(courseId string) (string, []model.Lesson, error) { // 2-task bajarildi
	rows, err := repo.Db.Query("SELECT course_id, title, content FROM lessons WHERE course_id = $1", courseId)
	if err != nil {
		return "", nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var lessons []model.Lesson
	for rows.Next() {
		lesson := model.Lesson{}
		err := rows.Scan(&lesson.CourseId, &lesson.Title, &lesson.Content)
		if err != nil {
			return "", nil, err
		}
		lessons = append(lessons, lesson)
	}

	return courseId, lessons, nil
}

// get course enrollment id - bu yerda course id boyicha userlarni olib keladi
func (repo *CourseRepository) GetUserByCourseIdWithEnrollment(course_id string) (string, []model.User, error) { //3-task bajarildi

	rows, err := repo.Db.Query("select c.id,u.id,u.name,u.email  from courses c inner join  enrollments e on c.id=e.course_id inner join users u on e.user_id=u.id  where c.id=$1", course_id)
	if err != nil {
		return "", nil, err
	}
	users := []model.User{}
	var courseId string
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&courseId, &user.Id, &user.Name, &user.Email)
		if err != nil {
			return "", nil, err
		}
		users = append(users, user)
	}
	return courseId, users, nil
}

// id si boyicha qidirib topib keladi
func (repo *CourseRepository) GetById(id string) (*model.Course, error) {
	course := model.Course{}
	err := repo.Db.QueryRow("select id,title,description from courses where id=$1", id).Scan(&course.Id, &course.Title, &course.Description)
	if err != nil {
		fmt.Println("+++++++", err)
		return nil, err
	}
	return &course, nil
}

// get poplularni course olib keladi start time va end time lar kirib keladi shu time oraliqda enrolment id course idsi title descr enrolment lar soni chiqaradi
func (repo *CourseRepository) GetPopularyCourse(startTime, endTime time.Time) ([]map[string]interface{}, error) {
	// query qismi
	query := `select c.id, c.title, COUNT(e.id) as enrollments_count from courses c join enrollments e on c.id = e.course_id where e.enrollment_date between $1 and $2 group by  c.id, c.title order by  enrollments_count desc `
	rows, err := repo.Db.Query(query, startTime, endTime)
	if err != nil {
		fmt.Println("++++++++++++", err)
		return nil, err
	}
	// map ochilgan
	var popularCourses []map[string]interface{}
	fmt.Println("))))))))))", rows)
	for rows.Next() {
		fmt.Println("-------")
		var courseID string
		var courseTitle string
		var enrollmentsCount string
		if err := rows.Scan(&courseID, &courseTitle, &enrollmentsCount); err != nil {

			return nil, err
		}
		fmt.Println("++++++++++", courseID, courseTitle, enrollmentsCount)
		popularCourses = append(popularCourses, map[string]interface{}{
			"course_id":         courseID,
			"course_title":      courseTitle,
			"enrollments_count": enrollmentsCount,
		})
	}
	return popularCourses, nil

}
