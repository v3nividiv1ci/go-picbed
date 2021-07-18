package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"

	//"github.com/nu7hatch/gouuid"
	"go-picbed/database"
	"go-picbed/model"
	"go-picbed/pics"
	"net/http"
)

func PicAdd(c *gin.Context) {
	u, _ := c.Get("user")
	DB := database.GetDB()
	//1. download pic
	newUUID := uuid.NewString()
	image, _ := c.FormFile("image")
	log.Println(image.Filename)
	err := c.SaveUploadedFile(image, pics.Root+newUUID+".jpeg")
	if err != nil {
		// 这里http状体码应该用什么呢ovo 官方文档没写怎么处理错误ovo
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "upload failed"})
		return
	}

	// 2. uuid created in pics
	TP := DB.Table("pics")
	master := u.(model.User).Username
	Pic := model.Pic{
		Uuid:   newUUID,
		Master: master,
	}
	TP.Create(&Pic)

	c.JSON(http.StatusOK, gin.H{"code": 200, "user": master, "uuid": newUUID})
}
