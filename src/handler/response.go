package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var version = "1.0"

type SimpleResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Time         time.Time
	Ver          string
	Data         []map[string]interface{}
}

func errorResponse(httpCode int, c *gin.Context, errCode int, errMsg string) {
	res := NewSimpleResponse()
	res.ErrorCode = errCode
	res.ErrorMessage = errMsg
	c.AbortWithStatusJSON(httpCode, res.toJSON())
}

func okResponse(c *gin.Context, data []map[string]interface{}) {
	res := NewSimpleResponse()
	res.Success = true
	res.Data = data
	c.JSON(http.StatusOK, res.toJSON())
}

func (r *SimpleResponse) toJSON() map[string]interface{} {
	return map[string]interface{}{
		"success": r.Success,
		"errCode": r.ErrorCode,
		"errMsg":  r.ErrorMessage,
		"ver":     r.Ver,
		"data":    r.Data,
	}
}

func NewSimpleResponse() *SimpleResponse {
	return &SimpleResponse{
		Success:      false,
		ErrorCode:    0,
		ErrorMessage: "",
		Time:         time.Now().UTC(),
		Ver:          version,
		Data:         nil,
	}
}
