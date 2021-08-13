package handler

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskTimeItemHandler struct {
	s service.TaskTimeItem
}

func (h TaskTimeItemHandler) GetAll(c *gin.Context) {
	res, err := h.s.GetAll()
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

// GetById get model.TimeItem by id
func (h TaskTimeItemHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 4004, err.Error())
		return
	}
	res, err := h.s.GetById(id)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, res)
}

// Create create model.TimeItem, return id
func (h TaskTimeItemHandler) Create(c *gin.Context) {
	var body *model.TimeItem
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

// Update update model.TimeItem
func (h TaskTimeItemHandler) Update(c *gin.Context) {
	m, err := getTimeItemModelAndValidate(c)
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

// DeleteById delete model.TimeItem by id
func (h TaskTimeItemHandler) DeleteById(c *gin.Context) {
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

func NewTaskTimeItemHandler(s service.TaskTimeItem) *TaskTimeItemHandler {
	return &TaskTimeItemHandler{s: s}
}
