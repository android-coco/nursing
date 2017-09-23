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

type IntakeOutputDup struct {
	BaseModel                  `xorm:"extends"`
	Type          uint8        `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
	Subtype       uint8        `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次)"`
	OperationType uint8        `json:"opertion_type" xorm:"notnull comment(操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便)"`
	Value         uint16       `json:"value" xorm:"notnull comment(采集值)"`
	Desc          string       `json:"desc" xorm:"comment(描述)"`
	NurseName     string       `json:"nurse_name" xorm:"notnull comment(护士姓名)"`
	RecordTime    Datetime_IOV `json:"recordTime" xorm:"notnull comment(采集时间)"`
}

type IntakeOutput struct {
	BaseModel            `xorm:"extends"`
	Type          uint8  `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
	Subtype       uint8  `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次)"`
	OperationType uint8  `json:"opertion_type" xorm:"notnull comment(操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便)"`
	Value         uint16 `json:"value" xorm:"notnull comment(采集值)"`
	Desc          string `json:"desc" xorm:"comment(描述)"`
	NurseName     string `json:"nurse_name" xorm:"notnull comment(护士姓名)"`
	RecordTime    string `xorm:"notnull comment(采集时间)"`
}

// inert one intake or output volume into database
func (iot *IntakeOutput) CollectIntakeOrOutputVolume() error {
	_, err := fit.MySqlEngine().InsertOne(iot)
	return err
}

/*
	patientId: patient's id
	tp： IntakeOutputTypeIntake，IntakeOutputTypeOutput
	page： minValue is 0，maximum return 10 data from db
	return slice of type IntakeOutput
*/
func QueryIntakeOrOutputVolume(patientId string, tp int, page int) ([]IntakeOutputDup, error) {
	count := 10
	responseObj := make([]IntakeOutputDup, 0)
	idx := page * count
	err := fit.MySqlEngine().Omit("id","datetime").Table("IntakeOutput").Where("patientId = ? and type = ?", patientId, tp).Limit(count, idx).Find(&responseObj)
	return responseObj, err
}

// query datas where type is equal to IntakeOutputTypeIntake or IntakeOutputTypeOutput
func QueryIntakeOrOutputVolumeAll(patientId string, page int) ([]IntakeOutputDup, error) {
	count := 10
	responseObj := make([]IntakeOutputDup, 0)
	idx := page * count
	err := fit.MySqlEngine().Omit("id","datetime").Table("IntakeOutput").Where("patientId = ? ", patientId).And("(type = ? or type = ?)", IntakeOutputTypeOutput, IntakeOutputTypeIntake).Limit(count, idx).Find(&responseObj)
	return responseObj, err
}
