package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"github.com/nu7hatch/gouuid"
	"go-picbed/database"
	"go-picbed/model"
	"net/http"
)

func PicAdd(c *gin.Context) {
	u, _ := c.Get("user")
	DB := database.GetDB()
	//TU := DB.Table("users")
	TP := DB.Table("pics")
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return
	}

	master := u.(model.User).Username
	Pic := model.Pic{
		Uuid:   newUUID,
		Master: master,
	}

	TP.Create(&Pic)
	c.JSON(http.StatusOK, gin.H{"code": 200, "user": master, "uuid": newUUID})
}
