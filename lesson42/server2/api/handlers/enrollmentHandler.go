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

// create enrollment - enrollment larni yaratamiz
func (handle *Handler) CreateEnrollment(gn *gin.Context) {
	enrollment := model.Enrollment{} //  enrollment model ichidan strucni olib keladi
	err := gn.BindJSON(&enrollment)  // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		BadRequest(gn, err)
	}
	// create postgres ichida sql yozilgan method ishlaydi creatEnroll
	//err = handle.EnrollmentRepo.CreateEnrollment(enrollment)
	enrollmentJson, err := json.Marshal(&enrollment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/enrollment/create", bytes.NewBuffer(enrollmentJson))
	if err != nil {
		ErrorResponse(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	if response.StatusCode == 200 {
		Ok(gn)
	} else {
		ErrorResponse(gn, err)
	}
}

// updated qilinadi enrollments
func (handle *Handler) UpdateEnrollment(gn *gin.Context) {
	enrollment := model.Enrollment{} //  enrollment model ichidan strucni olib keladi
	err := gn.BindJSON(&enrollment)  // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		BadRequest(gn, err)
	}
	enrollmentJson, err := json.Marshal(&enrollment)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/enrollment/update/%s", gn.Param("id")), bytes.NewBuffer(enrollmentJson))
	if err != nil {
		ErrorResponse(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	if response.StatusCode == 200 {
		Ok(gn)
	} else {
		ErrorResponse(gn, err)
	}
}

func (handle *Handler) DeleteEnrollment(gn *gin.Context) {
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/enrollment/delete/%s", gn.Param("id")), nil)
	if err != nil {
		ErrorResponse(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	if response.StatusCode == 200 {
		Ok(gn)
	} else {
		ErrorResponse(gn, err)
	}
}

// filter shu qisimda boladi
func (handle *Handler) GetEnrollment(gn *gin.Context) {
	enrollmentFilter := model.Filter{}
	//enrollmentFilter.User_id = gn.Query("id") // user id boyicha
	enrollmentFilter.UserId = gn.Query("user_id")                 // user id si boyicha
	enrollmentFilter.CourseId = gn.Query("course_id")             // course idsi boyicha
	enrollmentFilter.EnrollmentDate = gn.Query("enrollment_date") // controllment date boyicha
	limit, err := strconv.Atoi(gn.Query("limit"))                 // limit   query
	offset, err := strconv.Atoi(gn.Query("offset"))               // offset query
	enrollmentFilter.Limit = limit                                // limit enrollmentfiletrga tenlaymiz
	enrollmentFilter.Offset = offset
	//enrollmentFilterJson, err := json.Marshal(&enrollmentFilter)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/enrollments/get/?user_id=%s&course_id=%s&enrollment_date=%s&limit=%d&offset=%d", &enrollmentFilter.CourseId, &enrollmentFilter.UserId, &enrollmentFilter.EnrollmentDate, &enrollmentFilter.Limit, &enrollmentFilter.Offset), nil)
	if err != nil {
		ErrorResponse(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	if response.StatusCode == 200 {
		enrollments := []model.Enrollment{}
		err = json.NewDecoder(response.Body).Decode(&enrollments)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"enrollments": enrollments,
		})
	} else {
		ErrorResponse(gn, err)
	}

}

// get by id - search qiladi id boyicha

func (handle *Handler) GetEnrollmentById(gn *gin.Context) {
	// enrollem get by id si boyicha search qiladi
	//enrollment, err := handle.EnrollmentRepo.GetById(gn.Param("id"))
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/enrollment/id/%s", gn.Param("id")), nil)
	if err != nil {
		ErrorResponse(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	if response.StatusCode == 200 {
		enrollments := []model.Enrollment{}
		err = json.NewDecoder(response.Body).Decode(&enrollments)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"enrollments": enrollments,
		})
	} else {
		ErrorResponse(gn, err)
	}

}
