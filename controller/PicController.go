package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"strings"

	//"github.com/nu7hatch/gouuid"
	"go-picbed/database"
	"go-picbed/model"
	"go-picbed/pics"
	"net/http"
)

func Cmp(userT, userJ string) bool {
	if strings.Compare(userT, userJ) == 0 {
		return true
	} else {
		return false
	}
}

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

	// find all pics named imgName while master named sender
	var img model.Pic
	result := TP.Where("master = ? AND uuid = ?", sender, imgU).Find(&img)
	if result.RowsAffected == 0 {
		// 这个status code是对的吗,,
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "pic not found"})
		return
	}
	c.FileAttachment(pics.Root+img.Uuid+".jpeg", imgU)

	// 好像会panic..
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successfully downloaded"})

}
