package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

// function deleted lessons -lesson ni updated qiladi

func (handle *Handler) CreateLesson(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		BadRequest(gn, err)
	}
	err = handle.LessonRepo.CreateLesson(lesson)
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
}

// function updated lessons -lesson ni updated qiladi
func (handle *Handler) UpdateLesson(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		BadRequest(gn, err)
	}
	err = handle.LessonRepo.UpdateLesson(gn.Param("id"), lesson)
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
}

/* bu functionda deleteLesson da lesson ni delete qiliadi id boyicha delete qilish vazifasini oz ichiga oladi*/
func (handle *Handler) DeleteLesson(gn *gin.Context) {
	err := handle.LessonRepo.DeleteLesson(gn.Param("id"))
	if err != nil {
		BadRequest(gn, err)
	} else {
		Ok(gn)
	}
}

/* filter yoki getAll lessons ga nisbatana bu method ishlatiladi  qaysi fildan query dan kelsa shu paramni search qilib topib beradi */
func (handle *Handler) GetLesson(gn *gin.Context) {
	lessonFilter := model.Filter{}
	lessonFilter.CourseId = gn.Query("course_id") // course id boyicha
	lessonFilter.Title = gn.Query("title")        // title boyicha
	lessonFilter.Content = gn.Query("content")    // content boyicha
	limit, err := strconv.Atoi(gn.Query("limit")) // bu yerda parse qilib beradi  limit stringdan intga
	offset, err := strconv.Atoi(gn.Query("offset"))
	lessonFilter.Limit = limit
	lessonFilter.Offset = offset
	lessons, err := handle.LessonRepo.GetLessons(lessonFilter)
	if err != nil {
		fmt.Println("+++++++++++", err)
		BadRequest(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"lessons": lessons,
		})
	}
}

// getlesson by id search qiladi
func (handle *Handler) GetLessonById(gn *gin.Context) {
	lesson, err := handle.LessonRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"lessons": lesson,
		})
	}

}
