package models

type FGDDevice struct {
	ID             int    `gorm:"primaryKey"`
	DEVICEID       string `gorm:"type:varchar(3);index"`
	RECORDTIME     string `gorm:"type:varchar(25);default:"2006-01-02 15:04:05""`
	TEMPERATURE    string `gorm:"type:varchar(5)"`
	HUMIDITY       string `gorm:"type:varchar(3)"`
	LIGHTINTENSITY string `gorm:"type:varchar(5)"`
	SOLIDHUMIDITY  string `gorm:"type:varchar(4)"`
}

type FGDDeviceDaily struct {
	ID             int    `gorm:"primaryKey"`
	DEVICEID       string `gorm:"type:varchar(3);index"`
	RECORDTIME     string `gorm:"type:varchar(25);default:"2006-01-02 15:04:05""`
	TEMPERATURE    string `gorm:"type:varchar(10)"`
	HUMIDITY       string `gorm:"type:varchar(10)"`
	LIGHTINTENSITY string `gorm:"type:varchar(10)"`
	SOLIDHUMIDITY  string `gorm:"type:varchar(10)"`
}
