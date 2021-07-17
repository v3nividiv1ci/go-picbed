package main

import (
	"github.com/gin-gonic/gin"
	"go-picbed/auth"
	"go-picbed/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/test", auth.JwtMiddleWare(), controller.Test)
	return r
}
