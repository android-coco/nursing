//  Created by JP

package model

import (
	"fit"
	"time"
)

const (
	IntakeOutputTypeIntake int = 1 << iota
	IntakeOutputTypeOutput
)

//const (
//	IntakeTypeOther       int = iota + 1
//	IntakeTypeTransfusion
//	IntakeTypeEating
//)
//
//const (
//	OutputTypeOther     int = iota + 1
//	OutputTypeUrination
//	OutputTypeCacation
//)

// customizing time type with special format
type Datetime_IOV time.Time

func (t Datetime_IOV) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04") + `"`), nil
}

func (t Datetime_IOV) NormParse() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t Datetime_IOV) NormParse2() string {
	return time.Time(t).Format("2006-01-02 15:04")
}

type IntakeOutputDup struct {
	BaseModel                  `xorm:"extends"`
	Type          uint8        `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
	Subtype       uint8        `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次)"`
	OperationType uint8        `json:"opertion_type" xorm:"notnull comment(操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便)"`
	Value         uint16       `json:"value" xorm:"notnull comment(采集值)"`
	Desc          string       `json:"desc" xorm:"comment(描述)"`
	NurseName     string       `json:"nurse_name" xorm:"notnull comment(护士姓名)"`
	Testtime      Datetime_IOV `json:"recordTime" xorm:"notnull comment(采集时间)"`
	OtherDesc     string       `json:"other_desc" xorm:"comment(其它出量的补充描述)"`
}

type IntakeOutput struct {
	BaseModel            `xorm:"extends"`
	Type          uint8  `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
	Subtype       uint8  `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次)"`
	OperationType uint8  `json:"opertion_type" xorm:"notnull comment(操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便)"`
	Value         uint16 `json:"value" xorm:"notnull comment(采集值)"`
	Desc          string `json:"desc" xorm:"comment(描述)"`
	Testtime      string `xorm:"notnull comment(采集时间)"`
	OtherDesc     string `xorm:"comment(其它出量的补充描述)"`
}

// 插入
func (iot *IntakeOutput) CollectIntakeOrOutputVolume() error {
	_, err := fit.MySqlEngine().Table("IntakeOutput").InsertOne(iot)
	return err
}

/*
	根据类型查询出入量，page = -1时不分页
*/
func QueryIntakeOrOutputVolume(patientId string, tp int, page int) ([]IntakeOutputDup, error) {
	responseObj := make([]IntakeOutputDup, 0)
	if page >= 0 {
		count := 20
		idx := page * count
		err := fit.MySqlEngine().Omit("id", "datetime").Table("IntakeOutput").Where("patientId = ? and type = ?", patientId, tp).Desc("testtime").Limit(count, idx).Find(&responseObj)
		return responseObj, err
	} else {
		err := fit.MySqlEngine().Omit("id", "datetime").Table("IntakeOutput").Where("patientId = ? and type = ?", patientId, tp).Desc("testtime").Find(&responseObj)
		return responseObj, err
	}
}

// 所有的出入量
func QueryIntakeOrOutputVolumeAll(patientId string, page int) ([]IntakeOutputDup, error) {
	responseObj := make([]IntakeOutputDup, 0)
	if page >= 0 {
		count := 20
		idx := page * count
		err := fit.MySqlEngine().Omit("id", "datetime").Table("IntakeOutput").Where("patientId = ? ", patientId).And("(type = ? or type = ?)", IntakeOutputTypeOutput, IntakeOutputTypeIntake).Desc("testtime").Limit(count, idx).Find(&responseObj)
		return responseObj, err
	} else {
		err := fit.MySqlEngine().Omit("id", "datetime").Table("IntakeOutput").Where("patientId = ? ", patientId).And("(type = ? or type = ?)", IntakeOutputTypeOutput, IntakeOutputTypeIntake).Desc("testtime").Find(&responseObj)
		return responseObj, err
	}
}

/*根据时间段查询出入量*/
func QueryIntakeOrOutputVolumeWithPatient(pid int, startTime string, endTime string) ([]IntakeOutput, error) {
	response := make([]IntakeOutput, 0)
	err := fit.MySqlEngine().SQL("SELECT IntakeOutput.patientid,IntakeOutput.type,IntakeOutput.subtype,IntakeOutput.operationtype,IntakeOutput.`desc`,IntakeOutput.`value`,IntakeOutput.testtime from IntakeOutput WHERE IntakeOutput.patientid = ? and IntakeOutput.testtime BETWEEN ? AND ?",pid, startTime, endTime).Find(&response)
	return response, err
}