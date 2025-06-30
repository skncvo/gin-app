package handlers

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.String(200, "hello gin-gonic")
}
