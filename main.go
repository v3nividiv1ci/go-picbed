package main

import (
	"github.com/gin-gonic/gin"
	"go-picbed/database"
	//"net/http"
)

func main() {
	database.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()

}
