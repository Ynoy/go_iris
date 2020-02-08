package dto

import (
	"fmt"
	"github.com/kataras/iris"
)

type DeviceInfo struct {
	ServerIP     string `form:"serverIP" json:"serverIP"`
	Status       int    `form:"status" json:"status"`
	PppoeId      int    `form:"pppoeId" json:"pppoeId"`
	DeviceMac    string `form:"deviceMac" json:"deviceMac"`
	HardwareType string `form:"model" json:"model"`
	Idcode       string `form:"idcode" json:"idcode"`
}

func (u *DeviceInfo) Bind(ctx iris.Context) error {
	err := ctx.ReadForm(u)
	fmt.Println(err)
	return nil
}
