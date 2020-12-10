package main

import (
	"Beeper/backend/common"
	"Beeper/backend/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	//Init Database
	common.InitDB()
	//Create gin Engine
	r := gin.Default()
	//Router
	r = CollectRouter(r)
	//Server run
	err := r.Run(":" + conf.ServerPort)
	if err != nil {
		panic(err)
	}
}
