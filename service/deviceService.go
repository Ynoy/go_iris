package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"lbps/dto"
	"lbps/errors"
	"lbps/models"
	"os"
	"strconv"
)


func QueryDeviceInfo(status int) (int, error) {
	// 根据ID查询用户
	var err error
	//var users []*models.DeviceInfo
	var counts int
	if counts, err = dto.QueryDeviceInfosCount(status); err != nil {
		return 0, errors.DBError("query users by id", err)
	}
	dealWith(counts,status)

	return counts, nil
}

/**
	单线程执行任务，几十万数据效率太低
 */
func dealWith(counts int, status int) {
	size := counts / 100000
	if 0 < counts%100000 {
		size++
	}
	for i := 0; i < size; i++ {
		deviceInfos ,_:= dto.QueryDeviceInfos(status, i, 100000)
		fmt.Printf("index:%d",i)
		fmt.Printf("zip: %v.csv\n", i+1)
		csvBuffer := Csv(deviceInfos)
		// 将数据存入csv文件，并压缩
		csvFile, err := os.Create(strconv.Itoa(i+1) + ".csv")
		if err != nil {
			fmt.Println("open file is failed, err: ", err)
		}
		// 延迟关闭
		defer csvFile.Close()
		csvFile.WriteString("\xEF\xBB\xBF")

		csvFile.Write(csvBuffer.Bytes())
	}
}

func Csv(deviceInfos []*models.DeviceInfo) *bytes.Buffer {
	csvBuffer := bytes.NewBuffer(nil)
	csvWriter := csv.NewWriter(csvBuffer)
	if err := csvWriter.WriteAll(getData(deviceInfos)); err != nil {
		panic(err)
	}

	return csvBuffer
}

func getData(deviceInfos []*models.DeviceInfo) [][]string {
	data := make([][]string, len(deviceInfos))
	data[0] = []string{
		"Idcode", "DeviceMac",
	}
	for i, v := range deviceInfos {
		if i == 0 {
			continue
		}
		data[i]=[]string{
			v.Idcode,
			v.DeviceMac,
		}
	}
	return data
}