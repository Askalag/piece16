package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SimpleResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Time         time.Time
	Data         map[string]interface{}
}

func errorResponse(httpCode int, c *gin.Context, errCode int, errMsg string) {
	res := &SimpleResponse{
		Success:      false,
		ErrorCode:    errCode,
		ErrorMessage: errMsg,
		Time:         time.Now().UTC(),
	}
	c.AbortWithStatusJSON(httpCode, res)
}

func okResponse(c *gin.Context, data map[string]interface{}) {
	res := &SimpleResponse{
		Success:      true,
		ErrorCode:    0,
		ErrorMessage: "",
		Time:         time.Now().UTC(),
		Data:         data,
	}
	c.JSON(http.StatusOK, res)
}
