package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeHandler struct {
}

func (h *WelcomeHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "HELLO IM HERE...")
}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}
