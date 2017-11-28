package model

import (
	"fit"
)

type NursingRecords struct {
	Id          int64  `json:"id"`          //数据ID
	Updated     string `json:"updated"`     //更新时间
	NursType    int64  `json:"nurstype"`    //文书类型，1=护理记录单，2=首次护理记录单
	NursingId   string `json:"nursingid"`   //责任护士id
	NursingName string `json:"nursingname"` //责任护士姓名
	ClassId     string `json:"classid"`     //科室id
	PatientId   string `json:"patientid"`   //病人ID
	RecordId    int64  `json:"recordid"`    //文书id
	Comment     string `json:"comment"`     //备注（动作）
}

//根据病人ID查询文书记录
func QueryNRecords(PatientId string) ([]NursingRecords, error) {
	responseObj := make([]NursingRecords, 0)
	//fit.MySqlEngine().ShowSQL(true)
	err := fit.MySqlEngine().SQL("select * from NursingRecords where PatientId = ? ORDER BY Updated DESC", PatientId).Find(&responseObj)
	return responseObj, err
}

//根据病人ID查询文书记录根据文书分类和时间
func QueryNRecordsByTypeAndTime(PatientId string, NursType string, start, end string) ([]NursingRecords, error) {
	var err error
	responseObj := make([]NursingRecords, 0)
	//fit.MySqlEngine().ShowSQL(true)
	if NursType == "0" {
		err = fit.MySqlEngine().SQL("select * from NursingRecords where PatientId = ? AND Updated  BETWEEN ? AND ? ORDER BY Updated DESC", PatientId, start+" 00:00:00", end+" 23:59:59").Find(&responseObj)
	} else {
		err = fit.MySqlEngine().SQL("select * from NursingRecords where PatientId = ? AND NursType = ? AND Updated  BETWEEN ? AND ? ORDER BY Updated DESC", PatientId, NursType, start+" 00:00:00", end+" 23:59:59").Find(&responseObj)
	}
	return responseObj, err
}

//更新文书记录
func UpadteNRecords(RecordId int64,dateTime string) (int64, error) {
	//affected, err := fit.MySqlEngine().Id(RecordId).Cols("Updated").Update(&nursing)
	affected, err := fit.MySqlEngine().Table(new(NursingRecords)).Where("RecordId = ?", RecordId).Update(&NursingRecords{Id:RecordId, Updated:dateTime})
	return affected, err
}

//插入文书记录
func InsertNRecords(nursing NursingRecords) (int64, error) {
	affected, err := fit.MySqlEngine().Insert(&nursing)
	return affected, err
}
