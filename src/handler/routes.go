package handler

import (
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	w   *WelcomeHandler
	t   *TaskHandler
	ti  *TaskItemHandler
	tti *TaskTimeItemHandler
}

func MakeHandlers(s *service.Service) *Handler {
	return &Handler{
		w:   NewWelcomeHandler(),
		t:   NewTaskHandler(s.TaskService),
		ti:  NewTaskItemHandler(s.TaskItemService),
		tti: nil,
	}
}

func NewRootHandler(h *Handler) *gin.Engine {
	r := gin.New()

	// Auth
	//aApi := r.Group("/auth")

	// Tree Rest Api
	tApi := r.Group("/api")
	{
		// Welcome group
		wlc := tApi.Group("/wlc")
		{
			wlc.GET("/h", h.w.Hello)
			//wlc.GET("/t", h.w.Test)
		}

		// Task group
		task := tApi.Group("/task")
		{
			task.GET("/", h.t.GetAllTask)
			task.GET("/:id", h.t.GetById)
			task.POST("/", h.t.CreateTask)
			task.PATCH("/:model", h.t.Update)
			task.DELETE("/:id", h.t.DeleteById)
		}
	}
	return r
}
