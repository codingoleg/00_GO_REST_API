package router

import (
	"07_restapi/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	baseRouter.GET("/user/:tagId", tagsController.FindByUserId)
	baseRouter.GET("/order/:tagId", tagsController.FindByOrderId)
	baseRouter.POST("/order", tagsController.Create)

	return router
}
