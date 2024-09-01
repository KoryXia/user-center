package router

import (
	"github.com/gin-gonic/gin"
	"user-center/model"
)

func Register(ctx *gin.Context) {
	var request model.RegisterRequest
	if ctx.ShouldBind(&request) != nil {
		return
	}
}
