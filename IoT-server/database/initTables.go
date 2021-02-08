package database

import "IoT-server/models"

func Createtable() {
	Db.AutoMigrate(
		&models.FGDDevice{},
		&models.FGDDeviceDaily{},
	)
}
