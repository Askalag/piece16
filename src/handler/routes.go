package handler

import (
	"github.com/Askalag/piece16/src/middleware"
	"github.com/Askalag/piece16/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cmd *CmdHandler
	w   *WelcomeHandler
	tr  *TreeHandler
	t   *TaskHandler
	ti  *TaskItemHandler
	tti *TaskTimeItemHandler
}

func MakeHandlers(s *service.Service) *Handler {
	return &Handler{
		cmd: NewCmdHandler(s.CmdService),
		w:   NewWelcomeHandler(),
		tr:  NewTreeHandler(s.TreeService),
		t:   NewTaskHandler(s.TaskService),
		ti:  NewTaskItemHandler(s.TaskItemService),
		tti: NewTaskTimeItemHandler(s.TaskTimeItemService),
	}
}

func NewEngine(h *Handler) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(middleware.LogToConsole())

	// Rest Api group
	restApi := r.Group("/api")
	{
		// Welcome group
		wlc := restApi.Group("/wlc")
		{
			wlc.GET("/h", h.w.Hello)
		}

		// Commands group
		cmd := restApi.Group("/cmd")
		{
			cmd.GET("/initTables", h.cmd.InitTables)
			cmd.GET("/fillFullTree", h.cmd.fillFullTree)
			cmd.GET("/dropAllTables", h.cmd.DropAll)
		}

		// Tree group
		tree := restApi.Group("/tree")
		{
			tree.GET("/", h.tr.GetAll)
			tree.GET("/build/:id", h.tr.BuildById)
			tree.DELETE("/:treeId", h.tr.DeleteById)
			tree.POST("/", h.tr.Create)
			tree.PATCH("/", h.tr.Update)

			tree.DELETE("/wipe/:id", h.tr.WipeFullTree)

			tree.PATCH("/:treeId/task", h.tr.UpdateT)
			tree.PATCH("/:treeId/taskItem", h.tr.UpdateTI)
			tree.PATCH("/:treeId/timeItem", h.tr.UpdateTTI)
		}

		// Task group
		task := restApi.Group("/task")
		{
			task.GET("/", h.t.GetAll)
			task.GET("/:id", h.t.GetById)
			task.POST("/", h.t.CreateTask)
			task.PATCH("/", h.t.Update)
			task.DELETE("/:id", h.t.DeleteById)
		}

		// TaskItem group
		taskItem := restApi.Group("/taskItem")
		{
			taskItem.GET("/", h.ti.GetAll)
			taskItem.GET("/:id", h.ti.GetById)
			taskItem.POST("/", h.ti.Create)
			taskItem.PATCH("/", h.ti.Update)
			taskItem.DELETE("/:id", h.ti.DeleteById)
		}

		// TaskTimeItem group
		taskTimeItem := restApi.Group("/taskTimeItem")
		{
			taskTimeItem.GET("/", h.tti.GetAll)
			taskTimeItem.GET("/:id", h.tti.GetById)
			taskTimeItem.POST("/", h.tti.Create)
			taskTimeItem.PATCH("/", h.tti.Update)
			taskTimeItem.DELETE("/:id", h.tti.DeleteById)
		}
	}
	return r
}
