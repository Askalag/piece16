package handler

import (
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CmdHandler struct {
	s service.Cmd
}

func (h *CmdHandler) fillFullTree(c *gin.Context) {
	if err := h.s.FillFullTree(); err != nil {
		errorResponse(http.StatusInternalServerError, c, 1000, err.Error())
		return
	}
	okResponse(c, "ok")
}

func (h *CmdHandler) InitTables(c *gin.Context) {
	if err := h.s.InitTables(); err != nil {
		errorResponse(http.StatusInternalServerError, c, 1000, err.Error())
		return
	}
	okResponse(c, "ok")
}

func (h *CmdHandler) DropAll(c *gin.Context) {
	if err := h.s.DropAll(); err != nil {
		errorResponse(http.StatusInternalServerError, c, 1000, err.Error())
		return
	}
	okResponse(c, "ok")
}

func NewCmdHandler(s service.Cmd) *CmdHandler {
	return &CmdHandler{s: s}
}
