package main

import (
	"github.com/gin-gonic/gin"

	"github.com/s10akir/go-web-app/src/router"
)

func main() {
	r := gin.Default()
	indexRouter := router.IndexRouter{}

	r.GET("/", indexRouter.Index)

	r.Run()
}
