package dto

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lbps/models"
	"lbps/utility/db"
	"time"
)


// 根据用户ID查询
func QueryDeviceInfosCount(status int) (count int, err error) {
	res := db.GetMysql().Model(&models.DeviceInfo{}).Where("status = ?", status).Count(&count)
	err = res.Error

	return
}

func QueryDeviceInfos(status int, pageindex int, pagesize int)(err error)  {
	db, err := sql.Open("mysql", "lbps:rVMcISav%84n$tCp@tcp(210.74.14.26:6000)/lbps?charset=utf8&parseTime=true&loc=Local")

	fmt.Println("====start select=====")
	start := time.Now()

	//rows,err := db.GetMysql().Model(&models.DeviceInfo{}).Where("status = ?", status).Limit(pagesize).Offset(pageindex * pagesize).Rows()
	//err = res.Error

	db.Exec("select * from t_device_perception_status where status = 0 into outfile ~/app.csv")


	fmt.Println("====end select=====")
	endSelect := time.Now()
	fmt.Println("select time:=", endSelect.Sub(start))


	fmt.Println("====end select=====")
	end := time.Now()

	fmt.Println("dealWith time:=", end.Sub(start))
	return
}

