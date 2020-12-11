package middleware

import (
	"Beeper/backend/common"
	"Beeper/backend/model"
	"Beeper/backend/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		//Get Token
		tokenString:=ctx.GetHeader("Authorization")
		//Verify Token
		if tokenString==""||!strings.HasPrefix(tokenString,"Bearer ") {
			response.Response(ctx,http.StatusBadRequest,http.StatusBadRequest,nil,"Permission denied")
			//ctx.JSON(http.StatusBadRequest,gin.H{
			//	"msg":"Permission denied",
			//})
			ctx.Abort()
			return
		}
		//Get After Bearer String
		tokenString=tokenString[7:]
		//Parse Token
		token,claims,err:=common.ParseToken(tokenString)
		if err != nil||!token.Valid {
			log.Panic("Parse Token Failed:",err)
			response.Response(ctx,http.StatusUnauthorized,400,nil,"Parse Token Failed")
			//ctx.JSON(http.StatusUnauthorized,gin.H{
			//	"code":http.StatusUnauthorized,
			//	"msg":"Parse Token Failed",
			//})
			ctx.Abort()
			return
		}
		//If Pass Verify Get Claims User Id
		userid:=claims.Id
		DB:=common.Get_db()
		var user model.User
		DB.First(&user,userid)
		//Verify User Exists
		if user.ID==0{
			response.Response(ctx,http.StatusUnauthorized,http.StatusUnauthorized,nil,"User Not Exists")
			//ctx.JSON(http.StatusUnauthorized,gin.H{
			//	"code":http.StatusUnauthorized,
			//	"msg":"User Not Exists",
			//})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)

		ctx.Next()
	}
}
