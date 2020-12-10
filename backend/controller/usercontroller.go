package controller

import (
	"Chat-Room/backend/common"
	"Chat-Room/backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(ctx *gin.Context) {
	db := common.Get_db()
	var user model.User
	password := ctx.PostForm("password")
	username := ctx.PostForm("username")
	if len(username) < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "ID Too Short",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Password Too Short",
		})
		return
	}
	db.Where("username = ?", username).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("DB", user.Password, "form", password)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Password Not compete",
		})
		//TODO: Redis cache
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "Login successful",
	})
}
func Register(ctx *gin.Context) {
	db := common.Get_db()
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if len(username) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Username Too Short",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Password Too Short",
		})
		return
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Hash Password Faild",
		})
		return
	}
	newUser := model.User{
		Username: username,
		Password: string(hasePassowrd),
	}
	db.Create(&newUser)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create Successful",
		"ID":  newUser.ID,
	})
}
