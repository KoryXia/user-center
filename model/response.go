package model

import (
	"github.com/gin-gonic/gin"
)

type BaseResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func (response *BaseResponse[string]) ErrorResponse(c *gin.Context, code int, data string) {
	response.Code = code
	response.Status = "error"
	response.Data = data
	c.JSON(code, response)
}

func (response *BaseResponse[T]) SuccessResponse(c *gin.Context, code int, data T) {
	response.Code = code
	response.Status = "ok"
	response.Data = data
	c.JSON(code, response)
}

type UserInfoResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Nickname string `json:"nickname"`
}
