package model

import "fit"

type Devices struct {
	Id           int    `json:"id" xorm:"notnull comment(套餐的id)"`
	Devicesclass uint16 `json:"devicesclass" xorm:"notnull comment(套餐的科室信息)"`
	Devicesname  string `json:"devicesname" xorm:"notnull comment(套餐的名字)"`
	Devicelist   string `json:"devicelist" xorm:"notnull comment(套餐的设备)"`
	DeviceInfos []DeviceInfos `xorm:"-"`
}
//[{"name":"FitThm_X03_1wu","address":"DC:0D:30:00:0C:79","devicename":"ghjmkk","deviceinformation":"额温"}]
type DeviceInfos struct {
	Name string
	Address string
	DeviceName string
	DeviceInforMation string
	Type string
}
/**
查询套餐是否存在
 */
func GetDevicesByClassAndName(devicesclass int, devicesname string) (bool, error) {
	var devices Devices
	has, err := fit.MySqlEngine().Table("Devices").Where("devicesclass = ? and devicesname = ?", devicesclass, devicesname).Get(&devices)
	return has, err
}

/**
添加设备
 */
func InsertDevices(devices Devices) (int64, error) {
	has, err := fit.MySqlEngine().Insert(&devices)
	return has, err
}

/**
查询设备 根据科室ID
 */
func GetDevicesByClass(devicesclass int) ([]Devices, error) {
	devices := make([]Devices, 0)
	err := fit.MySqlEngine().Table("Devices").Where("devicesclass = ? ", devicesclass).Find(&devices)
	return devices, err
}

/**
查询设备 根据科室ID
 */
func GetAllDevices() ([]Devices, error) {
	devices := make([]Devices, 0)
	err := fit.MySqlEngine().Table("Devices").Find(&devices)
	return devices, err
}
