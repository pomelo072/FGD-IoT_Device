package controllers

import (
	"IoT-server/models"
	"IoT-server/myhandlers"
	"IoT-server/utils"
	"github.com/kataras/iris/v12"
)

func FGDDetails(ctx iris.Context) {
	DeviceID := ctx.URLParam("DeviceID")
	list := myhandlers.GetDetails(DeviceID)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

func FGDTotalList(ctx iris.Context) {
	DeviceID := ctx.URLParam("DeviceID")
	list := myhandlers.GetTotalList(DeviceID)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

func FGDUpload(ctx iris.Context) {
	DeviceID := ctx.URLParam("device_id")
	Temperature := ctx.URLParam("temperature")
	Humidity := ctx.URLParam("humidity")
	LightIntensity := ctx.URLParam("light")
	SolidHumidity := ctx.URLParam("solid")
	Data := models.FGDDevice{DEVICEID: DeviceID, TEMPERATURE: Temperature, HUMIDITY: Humidity, LIGHTINTENSITY: LightIntensity, SOLIDHUMIDITY: SolidHumidity}
	list := myhandlers.DataUpload(&Data)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
