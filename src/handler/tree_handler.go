package handler

import (
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TreeHandler struct {
	s service.Tree
}

// BuildById Build tree by Id...
func (h *TreeHandler) BuildById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 0, "")
		return
	}
	res, err := h.s.BuildById(id)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

func NewTreeHandler(s service.Tree) *TreeHandler {
	return &TreeHandler{s: s}
}
