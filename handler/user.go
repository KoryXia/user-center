package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user-center/model"
	"user-center/service"
	"user-center/utils"
)

func Register(c *gin.Context) {
	var request model.RegisterRequest
	if err := utils.Bind(c, &request); err != nil {
		return
	}

	response := model.BaseResponse[any]{}
	if err := service.Register(&request); err != nil {
		response.ErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, nil)
}

func Login(c *gin.Context) {
	var request model.LoginRequest
	if err := utils.Bind(c, &request); err != nil {
		return
	}

	response := model.BaseResponse[any]{}
	if err := service.Login(&request); err != nil {
		response.ErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	c.SetCookie(
		"user_login_id",
		fmt.Sprintf("%s%d", request.Username, time.Now().UnixMilli()),
		86400,
		"/",
		"",
		false,
		true,
	)
	response.SuccessResponse(c, http.StatusOK, nil)
}

func Logout(c *gin.Context) {
	value, _ := c.Cookie("user_login_id")
	var request model.LogoutRequest
	if err := utils.Bind(c, &request); err != nil {
		return
	}

	response := model.BaseResponse[any]{}
	if err := service.Logout(&request); err != nil {
		response.ErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	c.SetCookie(
		"user_login_id",
		value,
		-1,
		"/",
		"",
		false,
		true,
	)
	response.SuccessResponse(c, http.StatusOK, nil)
}

func GetUserInfo(c *gin.Context) {
	username := c.Param("username")

	response := model.BaseResponse[any]{}
	userInfo, err := service.GetUserInfo(username)
	if err != nil {
		response.ErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, userInfo)
}

func UpdateUserInfo(c *gin.Context) {

	var request model.UpdateRequest
	if err := utils.Bind(c, &request); err != nil {
		return
	}

	response := model.BaseResponse[any]{}
	if err := service.UpdateUserInfo(&request); err != nil {
		response.ErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, nil)
}
