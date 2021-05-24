package router

import (
	"github.com/gin-gonic/gin"

	"restful/handler"
)


func InitRouter() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("")
	{
		userRouter.GET("/user/:id", handler.GetOne)
		userRouter.GET("/users", handler.GetAll)
		userRouter.POST("/user", handler.Insert)
		userRouter.DELETE("/user/:id", handler.DeleteOne)
	}

	return router
}