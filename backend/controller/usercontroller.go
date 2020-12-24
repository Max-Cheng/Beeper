package controller

import (
	"Beeper/backend/common"
	"Beeper/backend/dto"
	"Beeper/backend/model"
	"Beeper/backend/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	db := common.Get_db()
	var user model.User
	var requestuser model.User
	ctx.Bind(&requestuser)
	//Login Verify
	if len(requestuser.Username) <= 0 {
		response.Failed(ctx,nil,"ID Can't Be empty")
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"msg": "ID Can't Be empty",
		//})
		return
	}

	db.Where("username = ?", requestuser.Username).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestuser.Password)); err != nil {
		response.Failed(ctx,nil,"Password Not compete")
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"msg": "Password Not compete",
		//})
		//TODO: Redis cache
		return
	}
	token,err:=common.ReleaseToken(user)
	if err != nil {
		log.Panic("Generate Token Error:",err)
	}
	response.Success(ctx,gin.H{"token":token},"Login successful")
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"data":gin.H{
	//		"token":token,
	//	},
	//	"msg": "Login successful",
	//})
}
func Register(ctx *gin.Context) {
	db := common.Get_db()
	var requestuser model.User
	ctx.Bind(&requestuser)
	if len(requestuser.Username) <= 0 {
		response.Failed(ctx,nil,"Username Too Short")
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"code": http.StatusBadRequest,
		//	"msg": "Username Too Short",
		//})
		return
	}
	if len(requestuser.Password) < 6 {
		response.Failed(ctx,nil,"Password Too Short")
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"code": http.StatusBadRequest,
		//	"msg": "Password Too Short",
		//})
		return
	}
	var user model.User
	db.Where("username = ?",requestuser.Username).First(&user)
	if user.ID!=0 {
		response.Failed(ctx,nil,"Username Invalid")
		return
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(requestuser.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx,http.StatusInternalServerError,http.StatusServiceUnavailable,nil,"Hash Password Faild")
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//	"code": http.StatusServiceUnavailable,
		//	"msg": "Hash Password Faild",
		//})
		return
	}
	newUser := model.User{
		Username: requestuser.Username,
		Password: string(hasePassowrd),
	}
	db.Create(&newUser)
	response.Success(ctx,gin.H{"ID":  newUser.ID},"Create Successful")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "Create Successful",
	//	"ID":  newUser.ID,
	//})
}
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx,gin.H{"user": dto.ToUserdto(user.(model.User))},"")

	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"user": dto.ToUserdto(user.(model.User)),
	//})
}