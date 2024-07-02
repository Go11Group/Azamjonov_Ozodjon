package model

type Lesson struct {
	LessonId string `json:"lesson_id"`
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Time
}

type LessonFilter struct {
	CourseId *string
	Title    *string
	Content  *string
	Limit    *string
	Offset   *string
}
