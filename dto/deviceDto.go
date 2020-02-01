package dto

import (
	"github.com/kataras/iris"
	"lbps/errors"
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
	if err := ctx.ReadForm(u); err != nil {
		return errors.ParamError("invalid form format")
	}

	return nil
}
