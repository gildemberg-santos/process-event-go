package controller

import "github.com/gin-gonic/gin"

func EventController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Event",
	})
}
