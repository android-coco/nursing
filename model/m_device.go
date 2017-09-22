package model

import "fit"

type Devices struct {
	Devicesclass     uint16      `json:"devicesclass" xorm:"notnull comment(套餐的科室信息)"`
	Devicesname      string     `json:"devicesname" xorm:"notnull comment(套餐的名字)"`
	Devicelist       string   `json:"devicelist" xorm:"notnull comment(套餐的设备)"`
}

func (mod Devices) InsertData(a interface{}) error {
	_, err := fit.MySqlEngine().Insert(a)
	return err
}