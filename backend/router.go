package main

import (
	"Beeper/backend/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/login", controller.Login)
	r.POST("/api/register", controller.Register)
	//TODO: ChatRoom Websocket Router
	return r
}
