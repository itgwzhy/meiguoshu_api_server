package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MEIGUOSHU/web-api/packages/controller"
)

var (
	Index *controller.Index
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	addRouter(router)
	return router
}

func addRouter(r *gin.Engine) {
	r.GET("/:id", Index.Get)
}
