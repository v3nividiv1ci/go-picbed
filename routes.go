package main

import (
	"github.com/gin-gonic/gin"
	"go-picbed/auth"
	"go-picbed/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/test", auth.JwtMiddleWare(), controller.Test)
	r.POST("/api/pic/upload", auth.JwtMiddleWare(), controller.PicAdd)
	r.GET("/api/pic/download", auth.JwtMiddleWare(), controller.PicDownload)
	// 需要鉴权还需要比对用户所以感觉不是很方便用static？
	//r.Static("/image", pics.Root)
	r.DELETE("/api/pic/delete", auth.JwtMiddleWare(), controller.PicDelete)
	r.GET("/api/pic/reveal", auth.JwtMiddleWare(), controller.PicReveal)
	return r
}
