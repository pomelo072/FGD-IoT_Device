package utils

import (
	"IoT-server/models"
	"crypto/sha1"
	"encoding/hex"
	"gorm.io/gorm"
	"io"
	"strconv"
)

func GetSHAEncode(str string) string {
	raw := sha1.New()
	io.WriteString(raw, str)
	safe := raw.Sum(nil)
	shastr := hex.EncodeToString(safe)
	return shastr
}
func GetReturnData(dt interface{}, msgstring string) *models.Result {
	result := new(models.Result)
	result.Data = dt
	if msgstring == "SUCCESS" {
		result.Code = 200
	} else {
		result.Code = 400
	}
	result.Msg = msgstring
	return result
}

// 分页器
// 分页数据请求分页大小为30, 45, 60三种大小
func Paginate(pages string, pagesizes string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(pages)
		if page == 0 {
			page = 1
		}
		pagesize, _ := strconv.Atoi(pagesizes)
		switch {
		case pagesize <= 30:
			pagesize = 30
		case pagesize < 60:
			pagesize = 45
		case pagesize >= 60:
			pagesize = 60
		}

		offset := (page - 1) * pagesize
		return db.Offset(offset).Limit(pagesize)
	}
}
