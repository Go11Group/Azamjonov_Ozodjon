package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* create course method dida gin contextni oladi  bu yerda request ni qabul qilib create methoddga berib yuboradi u method table yozadi*/
func (handle *Handler) CreateCourse(gn *gin.Context) {
	course := model.Course{}    //  enrollment model ichidan strucni olib keladi
	err := gn.BindJSON(&course) // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		BadRequest(gn, err)
	}
	courseJson, err := json.Marshal(&course)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	s, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:8080/api/course/create"), bytes.NewBuffer(courseJson))
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

// updated course bu updated qiliadi
func (handle *Handler) UpdateCourse(gn *gin.Context) {
	course := model.Course{}    //  enrollment model ichidan strucni olib keladi
	err := gn.BindJSON(&course) // blindjson bu body oqib olib strucga parse qiladi
	if err != nil {
		BadRequest(gn, err)
	}
	coursetJson, err := json.Marshal(&course)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/course/update/%s", gn.Param("id")), bytes.NewBuffer(coursetJson))
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

// delete course qilinadi
func (handle *Handler) DeleteCourse(gn *gin.Context) {
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/course/delete/%s", gn.Param("id")), nil)
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

// getcourse filter yoki getAt course
func (handle *Handler) GetCourse(gn *gin.Context) {
	courseFilter := model.Filter{} /// course course filter modeldan chaqiriladi
	courseFilter.Title = gn.Query("title")
	limit, err := strconv.Atoi(gn.Query("limit"))   // limitni  int ga parse qiliadi
	offset, err := strconv.Atoi(gn.Query("offset")) // offsetni int ga parse qiliadi
	courseFilter.Limit = limit                      // limit  bunda filterda boladi
	courseFilter.Offset = offset                    // offset bunda filterda boladi
	courseFilter.Description = gn.Query("description")
	//courseFilterJson, err := json.Marshal(&courseFilter)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/course/get/?title=%s&description=%s&limit=%d&offset=%d", &courseFilter.Title, &courseFilter.Description, &courseFilter.Limit, &courseFilter.Offset), nil)
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
	courses := []model.Course{}
	err = json.NewDecoder(response.Body).Decode(&courses)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"enrollments": courses,
	})
}

// course id si orqali qidirib lessons larni olib keladi
func (handle *Handler) GetLessonsByCourseId(gn *gin.Context) {
	// course id,lessons lar slice olamiz bunda shu couese id teng bolgan lessons lar slice bor
	//courseId, lessons, err := handle.CourseRepo.GetLessonByCourseId(gn.Param("id"))
	//if err != nil {
	//	fmt.Println("+++++++++++", err)
	//	ErrorResponse(gn, err)
	//} else {
	//	gn.JSON(200, gin.H{
	//		"message":   err,
	//		"status":    http.StatusOK,
	//		"time":      time.Now(),
	//		"course_id": courseId,
	//		"Lessons":   lessons,
	//	})
	//}
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/course/lessons/%s", gn.Param("id")), nil)
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
	lessons := []model.Lesson{}
	err = json.NewDecoder(response.Body).Decode(&lessons)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"lessons": lessons,
	})
}

func (handle *Handler) GetUserWithEnrollmentByCourseId(gn *gin.Context) {
	// course id,user lar slice ,err qaytaradi  course da enrollment bilen inner join
	//qilinidai user_id topiladi keyin user_id boyicha user ni topadi
	//courseId, users, err := handle.CourseRepo.GetUserByCourseIdWithEnrollment(gn.Param("id"))
	courseId := gn.Param("id")
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/course/enrollments/%s", gn.Param("id")), nil)
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
		users := []model.User{}
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(http.StatusOK, gin.H{
			"message":   err,
			"status":    http.StatusOK,
			"time":      time.Now(),
			"course_id": courseId,
			"User":      users,
		})
	} else {
		ErrorResponse(gn, err)
	}

}

// course larni idsi boyicha qidiradi
func (handle *Handler) GetCourseById(gn *gin.Context) {
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/course/id/%s", gn.Param("id")), nil)
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
		courses := []model.Course{}
		err = json.NewDecoder(response.Body).Decode(&courses)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"courses": courses,
		})
	} else {
		ErrorResponse(gn, err)
	}
}

// show popularni courselarni qidiradi
func (handle *Handler) ShowPopularCourse(gn *gin.Context) {
	//startTime := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	//endTime := time.Date(2030, 12, 31, 23, 59, 59, 999, time.UTC)

	time1 := gn.Query("start_time")
	time2 := gn.Query("end_time")

	startTime1 := strings.Split(time1, "-")
	startYear, err := strconv.Atoi(startTime1[0])
	startMonth, err := strconv.Atoi(startTime1[1])
	startDay, err := strconv.Atoi(startTime1[2])
	if err != nil {
		panic(err)
	}
	endTime1 := strings.Split(time2, "-")

	endYear, err := strconv.Atoi(endTime1[0])
	endMonth, err := strconv.Atoi(endTime1[1])
	endDay, err := strconv.Atoi(endTime1[2])
	if err != nil {
		panic(err)
	}

	startTime := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endYear, time.Month(endMonth), endDay, 00, 00, 00, 0, time.UTC)

	//courses, err := handle.CourseRepo.GetPopularyCourse(startTime, endTime)

	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/?start_time=%s&end_time=%s", gn.Query("start_time"), gn.Query("end_time")), nil)
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
		courses := []model.Course{}
		err = json.NewDecoder(response.Body).Decode(&courses)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"time_period": gin.H{
				"start_date": startTime,
				"end_date":   endTime,
			},
			"popular_courses": courses,
		},
		)
	} else {
		ErrorResponse(gn, err)
	}

}
