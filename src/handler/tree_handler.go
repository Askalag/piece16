package handler

import (
	"errors"
	"fmt"
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TreeHandler struct {
	s service.Tree
}

// BuildById Build tree by Id, return model.Tree with elements
func (h *TreeHandler) BuildById(c *gin.Context) {
	id, err := getTreeParam("id", c)
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

// DeleteTIById delete model.TaskItem, return updated model.Tree with elements
func (h *TreeHandler) DeleteTIById(c *gin.Context) {
	var params = []string{"treeId", "tiId"}
	treeId, err := getTreeParam(params[0], c)
	if err != nil {
		errorBadParamResponse(params[0], c)
		return
	}
	tiId, err := getTreeParam(params[1], c)
	if err != nil {
		errorBadParamResponse(params[1], c)
		return
	}

	res, err := h.s.DelTI(treeId, tiId, false)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

// UpdateTTI update model.TimeItem, return updated model.Tree with elements
func (h *TreeHandler) UpdateTTI(c *gin.Context) {
	var params = []string{"treeId"}
	treeId, err := getTreeParam(params[0], c)
	if err != nil {
		errorBadParamResponse(params[0], c)
		return
	}

	body, err := getTimeItemModelAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	tree, err := h.s.UpdTTI(treeId, body)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	okResponse(c, tree)
}

// getTimeItemModelAndValidate getting body and validate it
func getTimeItemModelAndValidate(c *gin.Context) (*model.TimeItem, error) {
	var m model.TimeItem
	if err := c.Bind(&m); err != nil {
		return &m, err
	}
	if ok := m.Valid(); !ok {
		return &m, errors.New("req body (TimeItem) is not valid")
	}
	return &m, nil
}

// getTreeParam getting param by key and try to parse it
func getTreeParam(key string, c *gin.Context) (int, error) {
	param, ok := c.Params.Get(key)
	value, err := strconv.Atoi(param)
	if !ok || err != nil {
		return 0, err
	}
	return value, nil
}

// errorBadParamResponse sending 400 code if param is not valid
func errorBadParamResponse(param string, c *gin.Context) {
	errorResponse(http.StatusBadRequest, c, 40004, fmt.Sprintf("The key param: '%s' is invalid - BadRequest", param))
}

func NewTreeHandler(s service.Tree) *TreeHandler {
	return &TreeHandler{s: s}
}
