package service

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Query(ctx iris.Context) string {

	sql := constructionSQL(ctx)
	fmt.Println("SQL:{",sql,"}")
	path := execShell(sql)

	return path
}

/*
	执行mysql导出命令
*/
func execShell(sql string) string {
	// 本地路径
	localPath := "/data/mysql_export/csv/"
	// nginx Download Path
	nginxPath := "https://lbps.ezlink-wifi.com/csv/"
	// 文件名称(时间戳命令)
	strFileName := strconv.FormatInt(time.Now().Unix(), 10) + ".csv"

	// shell命令
	shellString := `mysql -h 172.16.11.1  -ulbps --password='lbps(7&I*@hiveview' --database=lbps --default-character-set=utf8 -ss  -e ` + sql + " > " + localPath + strFileName

	//shellString := `mysql -h 210.74.14.26 -P 6000 -ulbps --password='rVMcISav%84n$tCp' --database=lbps --default-character-set=utf8 -ss  -e ` + sql + " > " + localPath + strFileName

	fmt.Println("execShell:", shellString)

	//执行shell命令

	command := exec.Command("/bin/bash","-c",shellString) //初始化Cmd
	fmt.Println(command.Path,command.Args)

	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()//运行脚本

	if nil != err {
		fmt.Println("execShell-start:",err)
	}
	fmt.Println("Process PID:", command.Process.Pid)
	fmt.Println("Process Stderr:", command.Stderr)
	fmt.Println("Process Stdout:", command.Stdout)
	//err = command.Wait() //等待执行完成
	//if nil != err {
	//	fmt.Println("execShell-wait:",err)
	//}
	fmt.Println("ProcessState PID:", command.ProcessState.Pid())


	result :=nginxPath+strFileName
	fmt.Println(result)
	return result
}

/*
	拼接sql语句
*/
func constructionSQL(ctx iris.Context) string {

	sql := "select * from `t_device_perception_status`"

	if  status,err :=ctx.PostValueInt("status"); err==nil&&(status==0||status==1) {
		sql = sql + " where status = "+ strconv.Itoa(status)
	}

	if serverIP :=ctx.PostValue("serverIP"); len(serverIP) != 0 {
		if strings.Contains(sql, "where") {
			sql =sql + " and server_ip = \""+serverIP+"\""
		}else {
			sql = sql + " where server_ip = \""+serverIP+"\""
		}

	}

	if pppoeId :=ctx.PostValue("pppoeId"); len(pppoeId) != 0 {
		if strings.Contains(sql, "where") {
			sql =sql + " and pppoe_id = \""+pppoeId+"\""
		}else {
			sql = sql + " where pppoe_id = \""+pppoeId+"\""
		}

	}

	if deviceMac :=ctx.PostValue("deviceMac"); len(deviceMac) != 0 {
		if strings.Contains(sql, "where") {
			sql =sql + " and device_mac like \"%"+deviceMac+"%\""
		}else {
			sql = sql + " where device_mac like \"%"+deviceMac+"%\""
		}

	}

	if model :=ctx.PostValue("model"); len(model) != 0 {
		if strings.Contains(sql, "where") {
			sql =sql + " and hardware_type = \""+model+"\""
		}else {
			sql = sql + " where hardware_type = \""+model+"\""
		}

	}

	if idcode :=ctx.PostValue("idcode"); len(idcode) != 0 {
		if strings.Contains(sql, "where") {
			sql =sql + " and idcode = \""+idcode+"\""
		}else {
			sql = sql + " where idcode = \""+idcode+"\""
		}

	}

	fmt.Println("SQL:{",sql,"}")
	return "'"+sql+"'"
}
