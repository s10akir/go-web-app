package main

import (
	"github.com/gin-gonic/gin"

	"github.com/s10akir/go-web-app/src/models"
	"github.com/s10akir/go-web-app/src/router"
)

func main() {
	r := gin.Default()
	indexRouter := router.IndexRouter{}
	tasksRouter := router.TasksRouter{}

	models.Init()

	r.GET("/", indexRouter.Index)

	tasks := r.Group("/tasks")
	{
		tasks.GET("", tasksRouter.Index)
		tasks.POST("", tasksRouter.Create)
		tasks.GET("/:id", tasksRouter.Read)
		tasks.POST("/:id", tasksRouter.Update)
		tasks.DELETE("/:id", tasksRouter.Delete)
	}

	r.Run()
}
