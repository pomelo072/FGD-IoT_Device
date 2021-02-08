package myhandlers

import (
	"IoT-server/database"
	"IoT-server/models"
	"strconv"
	"time"
)

func DataUpload(Data *models.FGDDevice) string {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	record := models.FGDDevice{DEVICEID: Data.DEVICEID, RECORDTIME: nowTime, TEMPERATURE: Data.TEMPERATURE, HUMIDITY: Data.HUMIDITY, LIGHTINTENSITY: Data.LIGHTINTENSITY, SOLIDHUMIDITY: Data.SOLIDHUMIDITY}
	database.Db.Create(&record)
	DataMerge(&record)
	return "Success."
}

func GetDetails(DeviceID string) interface{} {
	record := new(models.FGDDevice)
	database.Db.Where("device_id = ?", DeviceID).Last(&record)
	return record
}

func GetTotalList(DeviceID string) interface{} {
	var list []map[string]interface{}
	t2 := time.Now()
	t, _ := time.ParseDuration("-24h")
	t1 := t2.Add(t)
	database.Db.Model(&models.FGDDeviceDaily{}).Where("device_id = ?", DeviceID).Where("recordtime BETWEEN ? AND ?", t1, t2).Find(&list)
	return list
}

func DataMerge(newRecord *models.FGDDevice) string {
	DeviceID := newRecord.DEVICEID
	oldRecord := new(models.FGDDeviceDaily)
	result := database.Db.Where("device_id = ?", DeviceID).Last(&oldRecord)
	nt, _ := time.ParseInLocation("2006-01-02 15:04:05", newRecord.RECORDTIME, time.Local)
	ot, _ := time.ParseInLocation("2006-01-02 15:04:05", oldRecord.RECORDTIME, time.Local)
	sub := nt.Sub(ot)
	if result.RowsAffected != 1 || sub.Hours() >= 1 {
		record := models.FGDDeviceDaily{DEVICEID: newRecord.DEVICEID, RECORDTIME: newRecord.RECORDTIME, TEMPERATURE: newRecord.TEMPERATURE, HUMIDITY: newRecord.HUMIDITY, LIGHTINTENSITY: newRecord.LIGHTINTENSITY, SOLIDHUMIDITY: newRecord.SOLIDHUMIDITY}
		database.Db.Create(&record)
		return "record created."
	} else {
		// Temperature
		newTemperature, _ := strconv.Atoi(newRecord.TEMPERATURE)
		oldTemperature, _ := strconv.Atoi(oldRecord.TEMPERATURE)
		oldRecord.TEMPERATURE = strconv.Itoa((newTemperature + oldTemperature) / 2)
		// Humidity
		newHumidity, _ := strconv.Atoi(newRecord.HUMIDITY)
		oldHumidity, _ := strconv.Atoi(oldRecord.HUMIDITY)
		oldRecord.HUMIDITY = strconv.Itoa((newHumidity + oldHumidity) / 2)
		// Light_Intensity
		newLight, _ := strconv.Atoi(newRecord.LIGHTINTENSITY)
		oldLight, _ := strconv.Atoi(oldRecord.LIGHTINTENSITY)
		oldRecord.LIGHTINTENSITY = strconv.Itoa((newLight + oldLight) / 2)
		// Solid_Humidity
		newSolid, _ := strconv.Atoi(newRecord.SOLIDHUMIDITY)
		oldSolid, _ := strconv.Atoi(oldRecord.SOLIDHUMIDITY)
		oldRecord.SOLIDHUMIDITY = strconv.Itoa((newSolid + oldSolid) / 2)
		// update Data
		database.Db.Save(&oldRecord)
		return "record updated."
	}
}
