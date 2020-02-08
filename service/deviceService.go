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
	"time"
)

var pageSize = 500000

func QueryDeviceInfo(status int) (int, error) {
	// 根据ID查询用户
	var err error

	var counts int
	if counts, err = dto.QueryDeviceInfosCount(status); err != nil {
		return 0, errors.DBError("query users by id", err)
	}
	// 开始时间点
	fmt.Println("====start=====")
	start := time.Now()

	// 具体处理流程
	dealWithDB(counts, status)

	strFileName := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(strFileName)

	appendCSV(counts, strFileName)

	// 结束时间点
	fmt.Println("====end=====")
	end := time.Now()
	curr := end.Sub(start)
	fmt.Println("run time:", curr)

	return counts, nil
}

/*
	合并多个csv文件
*/
func appendCSV(counts int, newFileName string) {
	size := counts / pageSize
	if 0 < counts % pageSize {
		size++
	}

	for i := 0; i < size; i++ {
		csvBuffer := bytes.NewBuffer(nil)
		csvWriter := csv.NewWriter(csvBuffer)

		csvFile, _ := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		// 延迟关闭
		defer csvFile.Close()

		cntb, _ := os.Open(strconv.Itoa(i+1) + ".csv")

		reader := csv.NewReader(cntb)
		reader.FieldsPerRecord = -1
		content, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
			return
		}

		csvWriter.Flush()
		csvWriter.WriteAll(content)

		csvFile.Write(csvBuffer.Bytes())
	}

}

/*
	协程查询数据库
*/
func worker(jobChan chan int, i int, status int) {
	////开始时间点
	//fmt.Println("====start worker=====")
	//start := time.Now()

	 dto.QueryDeviceInfos(status, i, pageSize)
	//fmt.Printf("index:%d", i)
	//fmt.Printf("zip: %v.csv\n", i+1)
	//
	//csvBuffer := Csv(deviceInfos, i)
	//
	//// 将数据存入csv文件，并压缩
	//csvFile, err := os.Create(strconv.Itoa(i+1) + ".csv")
	//if err != nil {
	//	fmt.Println("open file is failed, err: ", err)
	//}
	//// 延迟关闭
	//defer csvFile.Close()
	//csvFile.WriteString("\xEF\xBB\xBF")
	//
	//csvFile.Write(csvBuffer.Bytes())

	////结束时间点
	//fmt.Println("====end worker=====")
	//end := time.Now()
	//curr := end.Sub(start)
	//fmt.Println("worker run time:", curr)

	jobChan <- i

}

/**
	单线程执行任务，几十万数据效率太低
*/
func dealWithDB(counts int, status int) {
	size := counts / pageSize
	if 0 < counts % pageSize {
		size++
	}

	// 创建多管道进行查询
	jobChan := make(chan int, size)

	for i := 0; i < size; i++ {
		go worker(jobChan, i, status)
	}

	for i := 0; i < size; i++ {
		<-jobChan
	}

}

/*
	初始化csv文件Buffer、Writer
 */
func Csv(deviceInfos [][]string, index int) *bytes.Buffer {
	csvBuffer := bytes.NewBuffer(nil)
	csvWriter := csv.NewWriter(csvBuffer)
	if err := csvWriter.WriteAll(deviceInfos); err != nil {
		panic(err)
	}

	return csvBuffer
}

/*
	定义csv文件数据列
 */
func getData(deviceInfos []*models.DeviceInfo, index int) [][]string {
	data := make([][]string, len(deviceInfos))
	if index == 0 {
		data[0] = []string{
			"Idcode", "DeviceMac",
		}
	}

	for i, v := range deviceInfos {
		if i == 0 {
			continue
		}
		data[i] = []string{
			v.Idcode,
			v.DeviceMac,
		}
	}
	return data
}
