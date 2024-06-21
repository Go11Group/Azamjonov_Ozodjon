package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

/* bu yerda create user qiladi    */
func (handle *Handler) CreateUser(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user) // blindjson - bu body kelgan malumotni jsonga parse qilib
	fmt.Println(user)
	if err != nil {
		BadRequest(gn, err)
	}
	err = handle.UserRepo.CreateUser(user) // create user sql -query code ga otadi  user struct ti berib yuboriladi
	if err != nil {
		fmt.Println("++++")
		ErrorResponse(gn, err) // agar bazaga saqlansamasa response=Ok
	} else {
		Ok(gn) // agar bazaga saqlansa response=Ok
	}
}

/* bu yerda updated user qiladi */
func (handle *Handler) UpdateUser(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user) // blindjson - bu body kelgan malumotni jsonga parse qilib
	fmt.Println(user)
	if err != nil {
		fmt.Println("+++++++", err)
		BadRequest(gn, err)
	}
	err = handle.UserRepo.UpdateUser(gn.Param("id"), user)
	if err != nil {

		fmt.Println("+++++++++", err)
		ErrorResponse(gn, err) // agar bazaga saqlansamasa response=internalservererror
	} else {
		Ok(gn) // agar bazaga updated response=Ok
	}
}

/* bu yerda deleted user qiladi */

func (handle *Handler) DeleteUser(gn *gin.Context) {
	err := handle.UserRepo.DeleteUser(gn.Param("id"))
	if err != nil {
		ErrorResponse(gn, err) // agar bazaga saqlansamasa response=internalservererror
	} else {
		Ok(gn) // agar bazaga delete response=Ok
	}
}

/* bu yerda filter va getAll user qiladi */
func (handle *Handler) GetUser(gn *gin.Context) {
	userFilter := model.Filter{}
	userFilter.Name = gn.Query("name")
	userFilter.Email = gn.Query("email")
	userFilter.Birthday = gn.Query("birthday")
	userFilter.Password = gn.Query("password")
	Limit, err := strconv.Atoi(gn.Query("limit"))
	Offset, err := strconv.Atoi(gn.Query("offset"))
	fmt.Println("-------------", gn.Query("offset"))
	fmt.Println("-------------", Offset)
	//err := gn.ShouldBindQuery(userFilter)

	userFilter.Limit = Limit
	userFilter.Offset = Offset
	fmt.Println(userFilter)
	users, err := handle.UserRepo.GetUser(userFilter)

	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)

	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"users":   users,
		})
	}
}

/* bu yerda search user_id boyicha  user qiladi */
func (handle *Handler) GetCourseByUserId(gn *gin.Context) {
	fmt.Println(gn.Param("id"))
	userId, courses, err := handle.UserRepo.GetEnrollmentByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {
		gn.IndentedJSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"user_id": userId,
			"Course":  courses,
		})
	}
}

/* bu yerda Email yoki Name boyicha qidiradi user qiladi */

func (handle *Handler) GetUserByEmailOrName(gn *gin.Context) {
	users, err := handle.UserRepo.GetUserByEmailOrName(gn.Param("name"), gn.Param("email"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"users":   users,
		})
	}
}

// user id boyich search qiladi

func (handle *Handler) GetUserById(gn *gin.Context) {
	user, err := handle.UserRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message": err,
			"status":  http.StatusOK,
			"time":    time.Now(),
			"user":    user,
		})
	}

}
