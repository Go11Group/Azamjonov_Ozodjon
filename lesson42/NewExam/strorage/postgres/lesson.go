package postgres

import (
	"database/sql"
	"my_project/model"
	"my_project/strorage"
	"time"
)

type LessonRepository struct {
	Db *sql.DB
}

func NewLessonRepository(db *sql.DB) *LessonRepository {
	return &LessonRepository{Db: db}
}

// create lesson bu lesson create qiladi yaratadi
func (repo *LessonRepository) CreateLesson(lesson model.Lesson) error {
	_, err := repo.Db.Exec("insert into lessons(course_id,title,content) values ($1,$2,$3)", lesson.CourseId, lesson.Title, lesson.Content)
	return err

}

// update lesson lessonlarni update qiladi  agar lessons fildlari bosh kelib qolsa bosh qilib qoymaydi
func (repo *LessonRepository) UpdateLesson(id string, lesson model.Lesson) error {
	lesson1 := model.Lesson{}
	err := repo.Db.QueryRow("select  course_id,title,content from lessons where id=$1", id).Scan(&lesson1.CourseId, &lesson1.Title, &lesson1.Content)
	if err != nil {
		return err
	}
	if len(lesson.CourseId) > 0 {
		lesson1.CourseId = lesson.CourseId
	}
	if len(lesson.Title) > 0 {
		lesson1.Title = lesson.Title
	}
	if len(lesson.Content) > 0 {
		lesson1.Content = lesson.Content
	}
	_, err = repo.Db.Exec("update lessons  set course_id=$1,title=$2,content=$3 where id =$4", lesson1.CourseId, lesson1.Title, lesson1.Content, id)
	_, err = repo.Db.Exec("update lessons  set updated_at=$1 where id =$2", time.Now(), id)
	return err
}

// delete lesson lesson ni delete qiladi delete qlingan vaqt ga yozib qoyadi
func (repo *LessonRepository) DeleteLesson(id string) error {
	_, err := repo.Db.Exec("delete from lessons where id=$1", id)
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err = repo.Db.Exec("update  lessons set deleted_at=$1 where id=$2", DeletedAt, id)
	return err
}

// get lesson va filter bunda hamma lesson yoki fildlar boyicha filtr yoki limit offset
func (repo *LessonRepository) GetLessons(f model.Filter) ([]model.Lesson, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id , course_id,title, content,created_at, updated_at, deleted_at
	 	from lessons where true`
	filter := ""
	if len(f.Content) > 0 {
		params["content"] = f.Content
		filter += " and content = :content "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :course_id "
	}
	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
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
	lessons := []model.Lesson{}
	for rows.Next() {
		lesson := model.Lesson{}
		err := rows.Scan(&lesson.Id, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil

}

// get by id lesson id si boyicha qidiradi

func (repo *LessonRepository) GetById(id string) (*model.Lesson, error) {
	lesson := model.Lesson{}
	err := repo.Db.QueryRow("select id,course_id,title,content from lessons where id=$1", id).Scan(&lesson.Id, &lesson.CourseId, &lesson.Title, &lesson.Content)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}
