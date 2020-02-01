package dto

import (
	"lbps/models"
	"lbps/utility/db"
)

// 根据用户ID查询
func QueryDeviceInfosCount(status int) (count int, err error) {
	res := db.GetMysql().Model(&models.DeviceInfo{}).Where("status = ?", status).Count(&count)
	err = res.Error




	return
}

func QueryDeviceInfos(status int,pageindex int,pagesize int) (deviceInfos []*models.DeviceInfo, err error) {
	res := db.GetMysql().Where("status = ?", status).Offset(pageindex*pagesize).Limit(pagesize).Find(&deviceInfos)
	err = res.Error
	return
}


