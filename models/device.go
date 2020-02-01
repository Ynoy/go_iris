package models

type DeviceInfo struct {
	Idcode             string `gorm:"column:idcode"`
	ServiceType        int    `gorm:"column:service_type"`
	HardwareType       string `gorm:"column:hardware_type"`
	DeviceMac          string `gorm:"column:device_mac"`
	WifiMac            string `gorm:"column:wifi_mac"`
	DeviceIP           string `gorm:"column:device_ip"`
	ServerIP           string `gorm:"column:server_ip"`
	PppoeId            string `gorm:"column:pppoe_id"`
	PppoePassword      string `gorm:"column:pppoe_password"`
	Ctime              string `gorm:"column:ctime"`
	Description        string `gorm:"column:description"`
	Status             int    `gorm:"column:status"`
	ManualServerForm   string `gorm:"column:manual_server_form"`
	ManualServerTo     string `gorm:"column:manual_server_to"`
	ManualToken        string `gorm:"column:manual_token"`
	Province           string `gorm:"column:province"`
	City               string `gorm:"column:city"`
	Isp                string `gorm:"column:isp"`
	FwVersion          string `gorm:"column:fw_version"`
	PerVersion         string `gorm:"column:per_version"`
	OriginalFilialeId  int    `gorm:"column:original_filiale_id"`
	OriginalCityId     int    `gorm:"column:original_city_id"`
	OriginalProvinceId int    `gorm:"column:original_province_id"`
	CommunityId        int    `gorm:"column:community_id"`
	BusinessHallId     int    `gorm:"column:business_hall_id"`
	DistrictBureauId   int    `gorm:"column:district_bureau_id"`
	BandWith           int    `gorm:"column:band_with"`
}

func (DeviceInfo) TableName() string {
	return "t_device_perception_status"
}

