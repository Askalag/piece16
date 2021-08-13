package handler

import (
	"errors"
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskItemHandler struct {
	s service.TaskItem
}

func (h TaskItemHandler) GetAll(c *gin.Context) {
	res, err := h.s.GetAll()
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

// GetById get model.TaskItem by id
func (h TaskItemHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, "")
		return
	}
	res, err := h.s.GetById(id)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

func (h TaskItemHandler) Create(c *gin.Context) {
	var body *model.TaskItem
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

// Update update model.TaskItem
func (h TaskItemHandler) Update(c *gin.Context) {
	m, err := getTaskItemModelAndValidate(c)
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

func (h TaskItemHandler) DeleteById(c *gin.Context) {
	paramId, err := getIdParamAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}

	if err := h.s.DeleteById(paramId); err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okIdsResponse(c, paramId)
}

// getTaskItemModelAndValidate getting body and validate it
func getTaskItemModelAndValidate(c *gin.Context) (*model.TaskItem, error) {
	var m model.TaskItem
	if err := c.Bind(&m); err != nil {
		return &m, err
	}
	if ok := m.Valid(); !ok {
		return &m, errors.New("req body (TimeItem) is not valid")
	}
	return &m, nil
}

func NewTaskItemHandler(s service.TaskItem) *TaskItemHandler {
	return &TaskItemHandler{s: s}
}
