package initialize

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"user-center/handler"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		method := c.Request.Method
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		path := c.Request.URL.Path
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		var logBuffer bytes.Buffer
		logBuffer.WriteString(fmt.Sprintf("%3d | %v | %s | %s | %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		))

		if errorMessage != "" {
			logBuffer.WriteString(fmt.Sprintf(" | %s", errorMessage))
		}

		log.Info(logBuffer.String())
	}
}

func ServerAndRouters() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CustomLogger())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/user", handler.Register)
	router.POST("/user/login", handler.Login)
	router.POST("/user/logout", handler.Logout)
	router.GET("/user/:username", handler.GetUserInfo)
	router.PUT("/user", handler.UpdateUserInfo)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
