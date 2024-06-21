package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
)

// function deleted lessons -lesson ni updated qiladi

func (handle *Handler) CreateLesson(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		BadRequest(gn, err)
	}
	//err = handle.LessonRepo.CreateLesson(lesson)
	lessonJson, err := json.Marshal(&lesson)
	if err != nil {
		ErrorResponse(gn, err)
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/lesson/create", bytes.NewBuffer(lessonJson))
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
	}
	if response.StatusCode == 201 || response.StatusCode == 200 {
		Ok(gn)
	} else {
		ErrorResponse(gn, err)
	}
}

// function updated lessons -lesson ni updated qiladi
func (handle *Handler) UpdateLesson(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		BadRequest(gn, err)
	}
	client := http.Client{}
	lessonJson, err := json.Marshal(&lesson)
	if err != nil {
		return
	}
	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/lesson/update/%s", gn.Param("id")), bytes.NewBuffer(lessonJson))
	response, err := client.Do(s)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	if response.StatusCode == 200 {
		Ok(gn)
	}

}

/* bu functionda deleteLesson da lesson ni delete qiliadi id boyicha delete qilish vazifasini oz ichiga oladi*/
func (handle *Handler) DeleteLesson(gn *gin.Context) {
	client := http.Client{}

	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/lesson/delete/%s", gn.Param("id")), nil)
	if err != nil {
		BadRequest(gn, err)
	} else {
		Ok(gn)
	}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
	}
	if response.StatusCode == 200 {
		Ok(gn)
	} else {
		ErrorResponse(gn, err)
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
	if err != nil {
		BadRequest(gn, err)
		return

	}
	client := http.Client{}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/lesson/get/?course_id=%s&title=%s&content=%s&limit=%d&offset=%d", &lessonFilter.CourseId, &lessonFilter.Title, &lessonFilter.Content, &lessonFilter.Limit, &lessonFilter.Offset), nil)
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
	}
	if response.StatusCode == 200 {
		lessons := []model.Lesson{}
		err = json.NewDecoder(response.Body).Decode(&lessons)
		if err != nil {
			ErrorResponse(gn, err)
		}
		gn.JSON(200, gin.H{
			"lessons": lessons,
		})
	} else {
		ErrorResponse(gn, err)
	}

}

// getlesson by id search qiladi
func (handle *Handler) GetLessonById(gn *gin.Context) {
	//lesson, err := handle.LessonRepo.GetById(gn.Param("id"))
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("localhost:8080/api/lesson/id/%s", gn.Param("id")), nil)
	if err != nil {
		BadRequest(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
	}
	if response.StatusCode == 200 {
		lessons := []model.Lesson{}
		err = json.NewDecoder(response.Body).Decode(&lessons)
		if err != nil {
			ErrorResponse(gn, err)
		} else {
			gn.JSON(200, gin.H{
				"lessons": lessons,
			})
		}
	} else {
		ErrorResponse(gn, err)
	}

}
