package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"user-center/model"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Bind[T model.Request](c *gin.Context, request T) error {
	if err := c.ShouldBind(request); err != nil {
		response := model.BaseResponse[string]{}
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	return nil
}
