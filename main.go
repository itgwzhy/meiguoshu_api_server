package main

import (
	"fmt"
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/router"
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/system"
	"github.com/gin-gonic/gin"
)

func main() {

	// init database
	if err := system.LoadConfig(); err != nil {
		fmt.Println(err)
		return
	}
	service := system.Config.Server
	host, port := service.Host, service.Port
	if service.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//router init
	r := router.InitRouter()
	r.Run(host + ":" + port)
}
