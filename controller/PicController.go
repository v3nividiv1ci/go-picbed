package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"os"

	//"github.com/nu7hatch/gouuid"
	"go-picbed/database"
	"go-picbed/model"
	"go-picbed/pics"
	"net/http"
)

func PicAdd(c *gin.Context) {
	u, _ := c.Get("user")
	master := u.(model.User).Username

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
	Pic := model.Pic{
		PicName: image.Filename,
		Uuid:    newUUID,
		Master:  master,
	}
	TP.Create(&Pic)

	c.JSON(http.StatusOK, gin.H{"code": 200, "user": master, "uuid": newUUID})
}

func PicDownload(c *gin.Context) {
	u, _ := c.Get("user")
	sender := u.(model.User).Username

	DB := database.GetDB()
	TP := DB.Table("pics")
	// get param
	imgU := c.PostForm("ImageUUID")

	// find where uuid and master both fit
	var img model.Pic
	result := TP.Where("master = ? AND uuid = ?", sender, imgU).Find(&img)
	if result.RowsAffected == 0 {
		// 这个status code是对的吗,,
		// assume that uuid from the front end is legal
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "permission denied"})
		return
	}

	c.FileAttachment(pics.Root+img.Uuid+".jpeg", imgU)

	// 好像会panic..
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successfully downloaded"})

}

func PicDelete(c *gin.Context) {
	u, _ := c.Get("user")
	sender := u.(model.User).Username

	DB := database.GetDB()
	TP := DB.Table("pics")

	imgU := c.PostForm("ImageUUID")
	var img model.Pic
	result := TP.Where("master = ? AND uuid = ?", sender, imgU).Find(&img)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "permission denied"})
		return
	}
	err := os.Remove(pics.Root + imgU + ".jpeg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "delete failed"})
		return
	}

	TP.Where("uuid = ?", imgU).Delete(&img)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successfully deleted"})
}
