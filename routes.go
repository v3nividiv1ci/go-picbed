package main

import (
	"github.com/gin-gonic/gin"
	"go-picbed/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/login", controller.Login)
	//r.POST("/test", func(c *gin.Context) {
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	c.JSON(200, gin.H{"name": username, "psd": password})
	//})
	return r
}
