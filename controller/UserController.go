package controller

import (
	"github.com/gin-gonic/gin"
	"go-picbed/auth"
	"go-picbed/database"
	"go-picbed/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// AutoRegister auto-register if not registered
func AutoRegister(username, password string) bool {
	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	DB := database.GetDB()
	User := model.User{
		Username: username,
		Password: string(EncryptedPassword),
	}
	DB.Create(&User)
	return true
}

func Login(c *gin.Context) {
	DB := database.GetDB()
	// get params
	username := c.PostForm("username")
	password := c.PostForm("password")
	// if registered
	var user model.User
	DB.Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		if AutoRegister(username, password) == false {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "encryption failed"})
			return
		}
	} else {
		// decrypt password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "wrong password"})
			return
		}
	}
	// issue token
	token, err := auth.TokenRelease(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token release failed"})
	}

	c.JSON(http.StatusOK,
		gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "successfully logged in"})

}

func Test(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
