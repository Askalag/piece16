package middleware

import (
	"github.com/Askalag/piece16/src/log"
	"github.com/gin-gonic/gin"
	"time"
)

func LogToConsole() gin.HandlerFunc {
	return func(c *gin.Context) {

		// log level
		level := "info"

		// Start
		startTime := time.Now()

		// Process request
		c.Next()

		// End Time
		endTime := time.Now()

		// Completed in
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := c.Request.Method

		// Request URI
		reqUri := c.Request.RequestURI

		// Status code
		statusCode := c.Writer.Status()

		// Request IP
		reqIP := c.ClientIP()

		// logger
		log.GinLogInfo(level, reqIP, statusCode, reqMethod, reqUri, latencyTime.Seconds())

	}
}

func LogToFile() gin.HandlerFunc {
	return nil
}
