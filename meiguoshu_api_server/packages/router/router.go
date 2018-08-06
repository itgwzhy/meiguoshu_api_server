package router

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/controller"
	"github.com/gin-gonic/gin"
)

var (
	product  *controller.Product
	category *controller.Category
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	addRouter(router)
	return router
}

func addRouter(r *gin.Engine) {
	r.GET("/:id", product.Get)
	r.POST("/:id", product.Post)
	r.PUT("/:id", product.Put)
	r.DELETE("/:id", product.Delete)
}
