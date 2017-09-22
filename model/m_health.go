package model

import (
	"fit"
)

/*import "fmt"*/

const (
   Temperature_Type int = 1      //体温
   Pulse_Type       int = 2      //脉搏
   Breathe_Type     int = 3      //呼吸
   Pressure_Type    int = 4      //血压
   Heartrate_Type   int = 5      //心率
   Spo2h_Type       int = 6      //血氧
   Glucose_Type     int = 7      //血糖
   Weight_Type      int = 8      //体重
   Height_Type      int = 9      //身高
)

func init() {
	//SyncTable()
}

type Temperature struct{
	BaseModel                   `xorm:"extends"`
	Testtime    fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Ttemptype   uint16          `json:"ttemptype" xorm:"notnull comment(体温的类型)"`
	Coolingtype uint16          `json:"coolingtype" xorm:"notnull comment(降温的类型)"`
	Value       float32         `json:"value" xorm:"notnull comment(值)"`
}

type Pulse struct {
	BaseModel                    `xorm:"extends"`
	Testtime         fit.JsonTime   `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene      uint16      `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value            uint16      `json:"value" xorm:"notnull comment(值)"`
	Whetherbriefness bool        `json:"whetherbriefness" xorm:"notnull comment(是否短促)"`
}

type Breathe struct {
	BaseModel		            `xorm:"extends"`
	Testtime     fit.JsonTime      `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene  uint16         `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value        uint16         `json:"value" xorm:"notnull comment(值)"`
	Whethertbm   bool           `json:"whethertbm" xorm:"notnull comment(是否上呼吸机)"`
}

type Pressure struct {
	BaseModel		            `xorm:"extends"`
	Testtime    fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Diavalue    uint16          `json:"diavalue" xorm:"notnull comment(低压值)"`
	Sysvalue    uint16          `json:"sysvalue" xorm:"notnull comment(高压值)"`
	Pulsevalue  uint16          `json:"pulsevalue" xorm:"notnull comment(脉率值)"`
}

type Heartrate struct{
	BaseModel		            `xorm:"extends"`
	Testtime    fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       uint16          `json:"value" xorm:"notnull comment(值)"`
}

type Spo2h struct {
	BaseModel		            `xorm:"extends"`
	Testtime    fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       uint16          `json:"value" xorm:"notnull comment(值)"`
}

type Glucose struct {
	BaseModel		            `xorm:"extends"`
	Testtime    fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       float32         `json:"value" xorm:"notnull comment(值)"`
}

type Weight struct {
	BaseModel		            `xorm:"extends"`
	Testtime    fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       float32         `json:"value" xorm:"notnull comment(值)"`
}

type Height struct {
	BaseModel		            `xorm:"extends"`
	Testtime   fit.JsonTime       `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16          `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value      	float64         `json:"value" xorm:"notnull comment(值)"`
}

