package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var version = "1.2"

type SimpleResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Time         time.Time
	Ver          string
	Data         interface{}
}

// okIdsResponse response with http status OK 200 and any ids
func okIdsResponse(c *gin.Context, ids ...int) {
	res := NewSimpleResponse()
	res.Success = true
	res.Data = ids
	c.JSON(http.StatusOK, res.toJSON())
}

// errorResponse response with error and any param
func errorResponse(httpCode int, c *gin.Context, errCode int, errMsg string) {
	res := NewSimpleResponse()
	res.ErrorCode = errCode
	res.ErrorMessage = errMsg
	c.AbortWithStatusJSON(httpCode, res.toJSON())
}

// okResponse response with http status OK 200 and any body data
func okResponse(c *gin.Context, data interface{}) {
	res := NewSimpleResponse()
	res.Success = true
	res.Data = data
	c.JSON(http.StatusOK, res.toJSON())
}

func (r *SimpleResponse) toJSON() map[string]interface{} {
	res := map[string]interface{}{
		"success":  r.Success,
		"err_code": r.ErrorCode,
		"err_msg":  r.ErrorMessage,
		"api_ver":  r.Ver,
		"data":     r.Data,
	}
	return res
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
