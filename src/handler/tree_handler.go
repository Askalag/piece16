package handler

import (
	"errors"
	"fmt"
	validator "github.com/Askalag/piece16/src/handler/validater"
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TreeHandler struct {
	s service.Tree
}

// GetAll return all trees
func (h *TreeHandler) GetAll(c *gin.Context) {
	res, err := h.s.GetAll()
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
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

// WipeFullTree wipe model.Tree include inner elements
func (h *TreeHandler) WipeFullTree(c *gin.Context) {
	var params = []string{"id"}

	id, err := getTreeParam(params[0], c)
	if err != nil {
		errorBadParamResponse(params[0], c)
		return
	}
	tree, err := h.s.BuildById(id)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 1000, err.Error())
		return
	}

	if err = h.s.DeleteFullTree(tree); err != nil {
		errorResponse(http.StatusInternalServerError, c, 1000, err.Error())
		return
	}
	okResponse(c, nil)
}

// DeleteById delete model.Tree by id
func (h *TreeHandler) DeleteById(c *gin.Context) {
	var params = []string{"treeId"}

	treeId, err := getTreeParam(params[0], c)
	if err != nil {
		errorBadParamResponse(params[0], c)
		return
	}

	err = h.s.DeleteById(treeId)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 0, err.Error())
		return
	}
	okResponse(c, treeId)
}

// Create create tree, return new tree id
func (h *TreeHandler) Create(c *gin.Context) {
	var body *model.Tree
	if err := c.BindJSON(&body); err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, "")
		return
	}

	id, err := h.s.Create(body)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, map[string]interface{}{"id": id})
}

// Update update tree
func (h *TreeHandler) Update(c *gin.Context) {
	m, err := getTreeModelAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}

	if err := h.s.Update(m); err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, nil)
}

// UpdateTI update model.TaskItem and return updated tree witch elements
func (h *TreeHandler) UpdateTI(c *gin.Context) {
	var params = []string{"treeId"}
	treeId, err := getTreeParam(params[0], c)
	if err != nil {
		errorBadParamResponse(params[0], c)
		return
	}

	body, err := getTaskItemModelAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	tree, err := h.s.UpdTI(treeId, body)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	okResponse(c, tree)
}

// UpdateT update model.Task and return updated tree witch elements
func (h *TreeHandler) UpdateT(c *gin.Context) {
	var params = []string{"treeId"}
	treeId, err := getTreeParam(params[0], c)
	if err != nil {
		errorBadParamResponse(params[0], c)
		return
	}

	body, err := getTaskModelAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	tree, err := h.s.UpdT(treeId, body)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	okResponse(c, tree)
}

// getTreeModelAndValidate getting body and validate it
func getTreeModelAndValidate(c *gin.Context) (*model.Tree, error) {
	var input model.Tree
	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}
	if !validator.ValidateTree(input) {
		return nil, errors.New("bad model properties")
	}
	return &input, nil
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
