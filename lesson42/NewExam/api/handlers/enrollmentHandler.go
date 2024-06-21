package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

// create enrollment - enrollment larni yaratamiz
func (handle *Handler) CreateEnrollment(gn *gin.Context) {
	enrollment := model.Enrollment{} //  enrollment model ichidan strucni olib keladi
	err := gn.BindJSON(&enrollment)  // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		BadRequest(gn, err)
	}
	// create postgres ichida sql yozilgan method ishlaydi creatEnroll
	err = handle.EnrollmentRepo.CreateEnrollment(enrollment)
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
}

// updated qilinadi enrollments
func (handle *Handler) UpdateEnrollment(gn *gin.Context) {
	enrollment := model.Enrollment{} //  enrollment model ichidan strucni olib keladi
	err := gn.BindJSON(&enrollment)  // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		BadRequest(gn, err)
	}
	// update postgres ichida sql yozilgan method ishlaydi updatenroll
	err = handle.EnrollmentRepo.UpdateEnrollment(gn.Param("id"), enrollment)
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
}

func (handle *Handler) DeleteEnrollment(gn *gin.Context) {
	err := handle.EnrollmentRepo.DeleteEnrollment(gn.Param("id"))
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
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
	enrollmentFilter.Offset = offset                              // offset enrollmentfiletrga tenlaymiz
	// get enrollemt filter bunda sql da query yozilgan file chaqirilgan
	enrollments, err := handle.EnrollmentRepo.GetEnrollments(enrollmentFilter) // getenrollment
	if err != nil {
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message":     "Success",     // success 200
			"status":      http.StatusOK, // status 200
			"time":        time.Now(),    // time now
			"enrollments": enrollments,
		})
	}
}

// get by id - search qiladi id boyicha

func (handle *Handler) GetEnrollmentById(gn *gin.Context) {
	// enrollem get by id si boyicha search qiladi
	enrollment, err := handle.EnrollmentRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message":    "Success",     //success
			"status":     http.StatusOK, // status 200
			"time":       time.Now(),    // now time
			"enrollment": enrollment,
		})
	}

}
