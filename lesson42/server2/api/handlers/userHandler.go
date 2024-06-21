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
)

/* bu yerda create user qiladi    */
func (handle *Handler) CreateUser(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user) // blindjson - bu body kelgan malumotni jsonga parse qilib
	if err != nil {
		BadRequest(gn, err)
	}
	userJson, err := json.Marshal(&user)
	if err != nil {

		BadRequest(gn, err)
		return
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/create", bytes.NewBuffer(userJson))
	if err != nil {

		ErrorResponse(gn, err)
		return
	}
	//err = handle.UserRepo.CreateUser(user) // create user sql -query code ga otadi  user struct ti berib yuboriladi
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err) // agar bazaga saqlansamasa response=Ok
		return
	}
	if response.StatusCode == 200 || response.StatusCode == 201 {
		Ok(gn)
	} else {
		fmt.Println("----------", userJson)

		ErrorResponse(gn, err)
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
	userJson, err := json.Marshal(&user)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	client := http.Client{}
	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/user/update/%s", gn.Param("id")), bytes.NewBuffer(userJson))
	//err = handle.UserRepo.UpdateUser(gn.Param("id"), user)
	if err != nil {
		fmt.Println("+++++++++", err)
		ErrorResponse(gn, err) // agar bazaga saqlansamasa response=internalservererror
	}
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

/* bu yerda deleted user qiladi */

func (handle *Handler) DeleteUser(gn *gin.Context) {
	//err := handle.UserRepo.DeleteUser(gn.Param("id"))
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/user/delete/%s", gn.Param("id")), nil)
	if err != nil {
		ErrorResponse(gn, err) // agar bazaga saqlansamasa response=internalservererror
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
	//userFilterJson, err := json.Marshal(&userFilter)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/get/?name=%s&email=%s&birthday=%s&password=%s&limit=%d&offset=%d", &userFilter.Name, userFilter.Email, &userFilter.Password, &userFilter.Limit, &userFilter.Offset), nil)
	//users, err := handle.UserRepo.GetUser(userFilter)
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)

	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		ErrorResponse(gn, err)
	}
	if response.StatusCode == 200 {
		users := []model.User{}
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"users": users,
		})
	} else {
		ErrorResponse(gn, err)
	}

}

/* bu yerda search user_id boyicha  user qiladi */
func (handle *Handler) GetCourseByUserId(gn *gin.Context) {
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/courses/%s", gn.Param("id")), nil)
	//userId, courses, err := handle.UserRepo.GetEnrollmentByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
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

/* bu yerda Email yoki Name boyicha qidiradi user qiladi */

func (handle *Handler) GetUserByEmailOrName(gn *gin.Context) {
	//users, err := handle.UserRepo.GetUserByEmailOrName(gn.Param("name"), gn.Param("email"))
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/get/%s/%s", gn.Param("name"), gn.Param("email")), nil)
	if err != nil {
		fmt.Println("+++++++++++", err)
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
		err := json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"users": users,
		})
	} else {
		ErrorResponse(gn, err)
	}
}

// user id boyich search qiladi

func (handle *Handler) GetUserById(gn *gin.Context) {
	//user, err := handle.UserRepo.GetById(gn.Param("id"))
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/id/%s", gn.Param("id")), nil)
	if err != nil {
		fmt.Println("+++++++++++", err)
		ErrorResponse(gn, err)
	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		return
	}
	if response.StatusCode == 200 {
		users := []model.User{}
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			ErrorResponse(gn, err)
			return
		}
		gn.JSON(200, gin.H{
			"users": users,
		})
	} else {
		ErrorResponse(gn, err)
	}

}
