package routers

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.POST("/users", func(context *gin.Context) {
		context.JSON(200, gin.H{"users": "users info"})
	})
	return ginRouter
}