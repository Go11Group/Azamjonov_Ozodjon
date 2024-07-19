package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/imtihon/model"
	"time"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

// Get lesson by ID
func (l *LessonRepo) GetById(id string) (model.Lessons, error) {
	var lesson model.Lessons
	err := l.db.QueryRow("SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons WHERE lesson_id = $1", id).Scan(
		&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	return lesson, err
}

// Create new lesson
func (l *LessonRepo) Create(lesson model.Lessons) error {
	query := `INSERT INTO lessons (lesson_id, course_id, title, content, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := l.db.Exec(query, uuid.New(), lesson.CourseId, lesson.Title, lesson.Content, time.Now(), time.Now(), 0)
	return err
}

// Update existing lesson
func (l *LessonRepo) Update(id string, lesson model.Lessons) error {
	fmt.Println(id)
	_, err := l.db.Exec("UPDATE lessons SET course_id = $1, title = $2, content = $3, updated_at = $4 WHERE lesson_id = $5",
		lesson.CourseId, lesson.Title, lesson.Content, time.Now(), id)
	return err
}

// Delete lesson by setting deleted_at timestamp
func (l *LessonRepo) Delete(id string) error {
	query := `update lessons set deleted_at = date_part('epoch', current_timestamp)::INT where lesson_id = $1 and deleted_at = 0`
	_, err := l.db.Exec(query, id)
	return err
}

// filterga get
func (l *LessonRepo) Get(query string, args []interface{}) ([]model.Lessons, error) {
	rows, err := l.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content,
			&lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lessons, nil
}

// Get courses by user_id
func (l *LessonRepo) LessonsByCourseId(lessonID string) ([]model.Lessons, error) {
	query := `
		SELECT l.lesson_id, l.course_id,l.title,l.content, l.created_at, l.updated_at, l.deleted_at
		FROM lessons l
		JOIN  courses c ON l.course_id = c.course_id
		WHERE c.course_id = $1 AND l.deleted_at = 0`
	rows, err := l.db.Query(query, lessonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
