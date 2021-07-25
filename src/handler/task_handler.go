package handler

import (
	"errors"
	validator "github.com/Askalag/piece16/src/handler/validater"
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	s service.Task
}

// GetById Get Task by Id...
func (h *TaskHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 0, "")
		return
	}
	res, err := h.s.GetById(id)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, []map[string]interface{}{
		res.ToJSON(),
	})
}

// GetAllTask Get all Tasks...
func (h *TaskHandler) GetAllTask(c *gin.Context) {
	arr, err := h.s.GetAll()
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
	}
	okResponse(c, model.ToJSONArr(*arr))
}

// CreateTask crate task...
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var body *model.Task
	if err := c.BindJSON(&body); err != nil {
		errorResponse(http.StatusBadRequest, c, 0, "")
		return
	}

	id, err := h.s.Create(body)
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, []map[string]interface{}{
		{"id": id},
	})
}

// Update update by model...
func (h *TaskHandler) Update(c *gin.Context) {
	m, err := getTaskModelAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 0, err.Error())
		return
	}

	if err := h.s.Update(m); err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okResponse(c, nil)
}

// DeleteById delete task by id...
func (h *TaskHandler) DeleteById(c *gin.Context) {
	paramId, err := getIdParamAndValidate(c)
	if err != nil {
		errorResponse(http.StatusBadRequest, c, 0, err.Error())
		return
	}

	if err := h.s.DeleteById(paramId); err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
		return
	}
	okIdsResponse(c, paramId)
}

func getTaskModelAndValidate(c *gin.Context) (*model.Task, error) {
	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}
	if !validator.ValidTaskModel(input) {
		return nil, errors.New("bad model properties")
	}
	return &input, nil
}

func getIdParamAndValidate(c *gin.Context) (int, error) {
	paramId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	if !validator.ValidId(&paramId) {
		return 0, errors.New("bad id param")
	}
	return paramId, nil
}

func NewTaskHandler(s service.Task) *TaskHandler {
	return &TaskHandler{s: s}
}
