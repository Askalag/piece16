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
	id, err := getIntParamByKey("id", c)
	if err != nil {
		return
	}
	res, err := h.s.BuildById(id)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

// DeleteTIById delete model.TaskItem and rebuild tree
func (h *TreeHandler) DeleteTIById(c *gin.Context) {
	treeId, err := getIntParamByKey("id", c)
	if err != nil {
		return
	}

	tiId, err := getIntParamByKey("tiId", c)
	if err != nil {
		return
	}

	res, err := h.s.DelTI(treeId, tiId, false)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

func getIntParamByKey(key string, c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 0, "")
		return 0, err
	}
	return id, nil
}

func NewTreeHandler(s service.Tree) *TreeHandler {
	return &TreeHandler{s: s}
}
