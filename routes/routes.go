package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skncvo/gin-app/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "gin server is running"})
	})

	userGroup := r.Group("/users")
	{
		// 매개변수 생략하는 이유
		// 지금 바로 실행이 아니라 '등록'만 하는 것
		// 즉 "참조"
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PATCH("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
