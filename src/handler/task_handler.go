package handler

import (
	"github.com/Askalag/piece16/src/model"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	s service.Task
}

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

func (h *TaskHandler) GetAllTask(c *gin.Context) {
	arr, err := h.s.GetAll()
	if err != nil {
		errorResponse(http.StatusInternalServerError, c, 0, err.Error())
	}
	okResponse(c, model.ToJSONArr(*arr))
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var body model.Task
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

func NewTaskHandler(s service.Task) *TaskHandler {
	return &TaskHandler{s: s}
}
