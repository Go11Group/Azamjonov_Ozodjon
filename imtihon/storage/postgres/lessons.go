package postgres

import (
	"database/sql"
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
func (l *LessonRepo) Lesson_Get(id string) (model.Lessons, error) {
	var lesson model.Lessons
	err := l.db.QueryRow("SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons WHERE lesson_id = $1", id).Scan(
		&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	return lesson, err
}

// Create new lesson
func (l *LessonRepo) Lesson_Create(lesson model.Lessons) error {
	query := `INSERT INTO lessons (lesson_id, course_id, title, content, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := l.db.Exec(query, lesson.LessonId, lesson.CourseId, lesson.Title, lesson.Content, time.Now(), time.Now(), nil)
	return err
}

// Update existing lesson
func (l *LessonRepo) Lesson_Update(lesson model.Lessons) error {
	query := `UPDATE lessons SET course_id = $1, title = $2, content = $3, updated_at = $4, deleted_at = $5 WHERE lesson_id = $6`
	_, err := l.db.Exec(query, lesson.CourseId, lesson.Title, lesson.Content, time.Now(), lesson.DeletedAt, lesson.LessonId)
	return err
}

// Delete lesson by setting deleted_at timestamp
func (l *LessonRepo) Lesson_Delete(id string) error {
	query := `UPDATE lessons SET deleted_at = $1 WHERE lesson_id = $2`
	_, err := l.db.Exec(query, time.Now(), id)
	return err
}

//                         lesson_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//                         course_id UUID REFERENCES Courses(course_id) ON DELETE CASCADE,
//                         title VARCHAR(100) NOT NULL,
//                         content TEXT,
//                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//                         deleted_at TIMESTAMP
