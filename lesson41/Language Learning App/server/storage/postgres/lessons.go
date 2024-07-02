package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"language_learning_app/pkg"
	"time"
)

type LessonRepo struct {
	DB *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{DB: db}
}

// Create
func (l *LessonRepo) CreateLesson(newLesson *model.Lesson) error {
	tx, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `insert into lessons(course_id, title, content, created_at)
		values ($1, $2, $3, $4)`
	_, err = tx.Exec(query, newLesson.CourseId, newLesson.Title,
		newLesson.Content, time.Now())
	return err
}

// Readall
func (l *LessonRepo) GetLessons(filter *model.LessonFilter) (*[]model.Lesson, error) {
	query := `select * from lessons where deleted_at is null`
	params := []interface{}{}
	paramsCount := 1

	if filter.CourseId != nil {
		params = append(params, filter.CourseId)
		query += pkg.CreateNewParams(" and course_id = ", paramsCount)
		paramsCount++
	}
	if filter.Title != nil {
		params = append(params, filter.Title)
		query += pkg.CreateNewParams(" and title = ", paramsCount)
		paramsCount++
	}
	if filter.Content != nil {
		params = append(params, filter.Content)
		query += pkg.CreateNewParams(" and content = ", paramsCount)
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

	lessons := &[]model.Lesson{}
	rows, err := l.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		lesson := model.Lesson{}
		err := rows.Scan(&lesson.LessonId, &lesson.CourseId,
			&lesson.Title, &lesson.Content, &lesson.CreatedAt,
			&lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		*lessons = append(*lessons, lesson)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return lessons, nil
}

// Read by id
func (l *LessonRepo) GetLessonById(lessonId string) (*model.Lesson, error) {
	query := `select * from lessons where lesson_id = $1`
	lesson := model.Lesson{}
	err := l.DB.QueryRow(query, lessonId).Scan(&lesson.LessonId,
		&lesson.CourseId, &lesson.Title, &lesson.Content,
		&lesson.CreatedAt, &lesson.UpdatedAt,
		&lesson.DeletedAt)

	return &lesson, err
}

// Update
func (l *LessonRepo) UpdateLesson(lesson *model.Lesson) error {
	tx, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update lessons set course_id=$1, title=$2, content=$3, 
	updated_at=$4 where lesson_id=$5 and deleted_at is null`

	_, err = tx.Exec(query, lesson.CourseId, lesson.Title, lesson.Content,
		time.Now(), lesson.LessonId)
	return err
}

// Delete
func (l *LessonRepo) DeleteLesson(lessonId string) error {
	tx, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `update lessons set deleted_at=$1 where lesson_id=$2
	and deleted_at is null`

	_, err = tx.Exec(query, time.Now(), lessonId)
	return err
}
