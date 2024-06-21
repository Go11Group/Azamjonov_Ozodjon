package handlers

import (
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
	course := model.Course{}    // yangi course struct ochish
	err := gn.BindJSON(&course) // requesdan kelayotgan malumot ni course ga parse qilish
	if err != nil {
		fmt.Println("++++++", err)
		BadRequest(gn, err)

	}
	// create course bu postgres ichidagi method borgan malumot ni tablega yozadi
	err = handle.CourseRepo.CreateCourse(course)
	if err != nil {
		fmt.Println("_____________", err)
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
}

// updated course bu updated qiliadi
func (handle *Handler) UpdateCourse(gn *gin.Context) {
	course := model.Course{}    //// yangi course struct ochish
	err := gn.BindJSON(&course) // requesdan kelayotgan malumot ni course ga parse qilish
	if err != nil {
		BadRequest(gn, err)
	}
	// bu yerda param da olib method ga berib yuboriladi paramda id  keladi
	err = handle.CourseRepo.UpdateCourse(gn.Param("id"), course)
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
	}
}

// delete course qilinadi
func (handle *Handler) DeleteCourse(gn *gin.Context) {
	fmt.Println(gn.Param("id"))
	err := handle.CourseRepo.DeleteCourse(gn.Param("id")) // param dan id oladi
	if err != nil {
		ErrorResponse(gn, err)
	} else {
		Ok(gn)
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
	courses, err := handle.CourseRepo.GetCourses(courseFilter) // course lar slice qaytaradi ,error
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK, // time 200
			"time":    time.Now(),
			"courses": courses,
		})
	}
}

// course id si orqali qidirib lessons larni olib keladi
func (handle *Handler) GetLessonsByCourseId(gn *gin.Context) {
	// course id,lessons lar slice olamiz bunda shu couese id teng bolgan lessons lar slice bor
	courseId, lessons, err := handle.CourseRepo.GetLessonByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {
		gn.JSON(200, gin.H{
			"message":   err,
			"status":    http.StatusOK,
			"time":      time.Now(),
			"course_id": courseId,
			"Lessons":   lessons,
		})
	}
}

func (handle *Handler) GetUserWithEnrollmentByCourseId(gn *gin.Context) {
	// course id,user lar slice ,err qaytaradi  course da enrollment bilen inner join
	//qilinidai user_id topiladi keyin user_id boyicha user ni topadi
	courseId, users, err := handle.CourseRepo.GetUserByCourseIdWithEnrollment(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err) //
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message":   err,
			"status":    http.StatusOK,
			"time":      time.Now(),
			"course_id": courseId,
			"User":      users,
		})
	}
}

// course larni idsi boyicha qidiradi
func (handle *Handler) GetCourseById(gn *gin.Context) {
	// paramdan id ni beradi course lar slice ni,err qaytaradi
	course, err := handle.CourseRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"courses": course, // course id

		})
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

	courses, err := handle.CourseRepo.GetPopularyCourse(startTime, endTime)
	if err != nil {
		ErrorResponse(gn, err)
		return
	}

	response := gin.H{
		"time_period": gin.H{
			"start_date": startTime,
			"end_date":   endTime,
		},
		"popular_courses": courses,
	}
	gn.JSON(http.StatusOK, response)
}
