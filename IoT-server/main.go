package main

import (
	"IoT-server/config"
	"IoT-server/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func main() {
	IotServer := iris.New()

	//Iot_server.PartyFunc("/admin", func(admin router.Party) {
	//	admin.Get("/", controllers.Admin)
	//})

	IotServer.PartyFunc("/FGD", func(FGD router.Party) {
		FGD.Get("/details", controllers.FGDDetails)
		FGD.Get("/upload", controllers.FGDUpload)
		FGD.Get("/list", controllers.FGDTotalList)
	})

	IotServer.Run(iris.Addr(config.Sysconfig.Port))
}
