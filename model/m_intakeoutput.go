package model

import (
	"fit"
)


const (
	IntakeOutputTypeIntake int = 1 << iota
	IntakeOutputTypeOutput
)

const (
	IntakeTypeOther       int = iota + 1
	IntakeTypeTransfusion
	IntakeTypeEating
)

const (
	OutputTypeOther     int = iota + 101
	OutputTypeUrination
	OutputTypeCacation
)

type IntakeOutput struct {
	BaseModel               `xorm:"extends"`
	Type       uint8        `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
	Subtype    uint8        `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量/ml，2：输液入量/ml，3：饮食入量/ml，101：其他出量/ml，102：排尿出量/ml，103：大便出量/次)"`
	RecordTime string    	`json:"recordTime" xorm:"notnull comment(采集时间)"`
	Value      uint16       `json:"value" xorm:"notnull comment(采集值)"`
	Desc       string       `json:"desc" xorm:"comment(描述)"`
}

// inert one intake or output volume into database
func (iot *IntakeOutput) CollectIntakeOrOutputVolume() error {
	_, err := fit.Engine().InsertOne(iot)
	return err
}

/*
	patientId: patient's id
	tp： IntakeOutputTypeIntake，IntakeOutputTypeOutput
	page： minValue is 0，maximum return 10 data from db
	return slice of type IntakeOutput
*/
func QueryIntakeOrOutputVolume(patientId string, tp int, page int) ([]interface{}, error) {
	count := 10
	temp := make([]IntakeOutput, 0)
	idx := page * count
	err := fit.Engine().Omit("id","date_time").Where("patient_id = ? and type = ?", patientId, tp).Limit(count, idx).Find(&temp)
	return convertTypeToInterface(temp), err
}


func QueryIntakeOrOutputVolumeAll(patientId string, page int) ([]interface{}, error) {
	count := 10
	temp := make([]IntakeOutput, 0)
	idx := page * count
	err := fit.Engine().Omit("id","date_time").Where("patient_id = ?", patientId).And("type = ? or type = ?", IntakeOutputTypeOutput, IntakeOutputTypeIntake).Limit(count, idx).Find(&temp)
	return convertTypeToInterface(temp), err
}


func convertTypeToInterface(slice []IntakeOutput) []interface{} {
	var ifSlice []interface{} = make([]interface{}, len(slice))
	for idx, v := range slice {
		ifSlice[idx] = v
	}
	return ifSlice
}
