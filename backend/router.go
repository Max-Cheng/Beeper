package main

import (
	"Beeper/backend/controller"
	"Beeper/backend/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/register", controller.Register)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)
	//TODO: ChatRoom Websocket Router
	return r
}
