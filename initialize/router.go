package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerAndRouters() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
