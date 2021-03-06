package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeHandler struct {
}

// Hello welcome page
func (h *WelcomeHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Here we go...")
}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}
