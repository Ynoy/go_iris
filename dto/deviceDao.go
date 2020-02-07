package dto

import (
	"fmt"
	"lbps/models"
	"lbps/utility/db"
	"time"
)

var DB = db.GetMysql()

// 根据用户ID查询
func QueryDeviceInfosCount(status int) (count int, err error) {
	res := DB.Model(&models.DeviceInfo{}).Where("status = ?", status).Count(&count)
	err = res.Error

	return
}

func QueryDeviceInfos(status int, pageindex int, pagesize int) (deviceInfos []*models.DeviceInfo, err error) {
	fmt.Println("====start select=====")
	start := time.Now()

	res := DB.Where("status = ?", status).Offset(pageindex * pagesize).Limit(pagesize).Find(&deviceInfos)
	err = res.Error
	fmt.Println("====end select=====")
	end := time.Now()
	fmt.Println("select time:=", end.Sub(start))
	return
}
