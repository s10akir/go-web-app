package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexRouter struct{}

func (indexRouter IndexRouter) Index(context *gin.Context) {
	context.String(http.StatusOK, "This is /index")
}
