package http

import (
	"go-learn/internal/app/wweb/controller"

	"github.com/gin-gonic/gin"
)

// Init route
func route(engine *gin.Engine) {
	//user
	user := controller.NewUserController()
	engine.POST("/user", user.Create)
	engine.DELETE("/user/:id", user.Delete)
	engine.PUT("/user", user.Update)
	engine.GET("/user", user.FindByPage)

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
