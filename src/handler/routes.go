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

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		w:   NewWelcomeHandler(),
		t:   NewTaskHandler(s.Task),
		ti:  NewTaskItemHandler(s.TaskItem),
		tti: nil,
	}
}

func AddRoutes(h *Handler) *gin.Engine {
	r := gin.New()

	// Auth
	//aApi := r.Group("/auth")

	// Tree Rest Api
	tApi := r.Group("/api")
	{
		// Welcome
		wlc := tApi.Group("/hello")
		{
			wlc.GET("/", h.w.Hello)
		}

		// Task
		task := tApi.Group("/tree")
		{
			task.GET("/")
			task.POST("/", h.t.CreateTask)
		}
	}
	return r
}
