package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TasksRouter struct{}

func (tasksRouter TasksRouter) Index(context *gin.Context) {
	context.String(http.StatusOK, "This is /tasks")
}

func (tasksRouter TasksRouter) Create(context *gin.Context) {
	context.String(http.StatusOK, "This is /tasks"+" (Create)")
}

func (tasksRouter TasksRouter) Read(context *gin.Context) {
	context.String(http.StatusOK, "This is /tasks/"+context.Param("id")+" (READ)")
}

func (tasksRouter TasksRouter) Update(context *gin.Context) {
	context.String(http.StatusOK, "This is /tasks/"+context.Param("id")+" (UPDATE)")
}

func (tasksRouter TasksRouter) Delete(context *gin.Context) {
	context.String(http.StatusOK, "This is /tasks/"+context.Param("id")+" (DELETE)")
}
