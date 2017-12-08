//  Created by JP

package model

import (
	"fit"
	"time"
	/*"github.com/go-xorm/xorm"*/
)

// customizing time type with special format
type DatetimeWithoutSeconds time.Time

/*精确到分，API专用*/
func (t DatetimeWithoutSeconds) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04") + `"`), nil
}

/*精确到秒*/
func (t DatetimeWithoutSeconds) ParseToSecond() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

/*精确到分*/
func (t DatetimeWithoutSeconds) ParseToMinute() string {
	return time.Time(t).Format("2006-01-02 15:04")
}

//type IntakeOutputDup struct {
//	BaseModel                  `xorm:"extends"`
//	Type          uint8        `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
//	Subtype       uint8        `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次)"`
//	OperationType uint8        `json:"opertion_type" xorm:"notnull comment(操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便)"`
//	Value         uint16       `json:"value" xorm:"notnull comment(采集值)"`
//	Desc          string       `json:"desc" xorm:"comment(描述)"`
//	NurseName     string       `json:"nurse_name" xorm:"notnull comment(护士姓名)"`
//	Testtime      fit.JsonTime `json:"recordTime" xorm:"notnull comment(采集时间)"`
//	OtherDesc     string       `json:"other_desc" xorm:"comment(其它出量的补充描述)"`
//}
//
//type IntakeOutput struct {
//	BaseModel            `xorm:"extends"`
//	Type          uint8  `json:"type" xorm:"notnull comment(出入量类型，1：入量，2：出量)"`
//	Subtype       uint8  `json:"subtype" xorm:"notnull comment(出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次)"`
//	OperationType uint8  `json:"opertion_type" xorm:"notnull comment(操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便)"`
//	Value         uint16 `json:"value" xorm:"notnull comment(采集值)"`
//	Desc          string `json:"desc" xorm:"comment(描述)"`
//	Testtime      string `xorm:"notnull comment(采集时间)"`
//	OtherDesc     string `xorm:"comment(其它出量的补充描述)"`
//}

type IntakeOutput struct {
	HeadType  string `xorm:"notnull comment(一级类型,入量15,出量16)"`
	TestTime  string `xorm:"notnull comment(测量时间)"`
	SubType   int    `xorm:"notnull comment(二级类型)"`
	Other     int    `xorm:"notnull comment(三级类型)"`
	OtherStr  string `xorm:"notnull comment(其它的第三级自定义类型)"`
	Describe  string `xorm:"notnull comment(描述)"`
	Value     string `xorm:"notnull comment(值、量)"`
	PatientId int64    `xorm:"notnull comment(病人id)"`
	NurseId   int    `xorm:"notnull comment(护士id)"`
	NurseName string `xorm:"notnull comment(护士姓名)"`
}

type IntakeOutputDup struct {
	NurseId   string                 `json:"nurse_id"`
	NurseName string                 `json:"nurse_name"`
	PatientId string                 `json:"patient_id"`
	HeadType  string                 `json:"type"`
	SubType   int                    `json:"subtype"`
	Other     int                    `json:"opertion_type"`
	Value     string                 `json:"value"`
	Describe  string                 `json:"desc"`
	Testtime  DatetimeWithoutSeconds `json:"recordTime"`
	OtherStr  string                 `json:"other_desc"`
}

// 插入
func (iot *IntakeOutput) CollectIntakeOrOutputVolume() error {

	_, err := fit.MySqlEngine().Table("NurseChat").InsertOne(iot)
	return err
}

//添加护理数据
func QueryIntakeOrOutput(item *NurseChat) error{
	has,err := fit.MySqlEngine().QueryString("SELECT id FROM NurseChat WHERE TestTime = ? and PatientId = ? and HeadType = ? and SubType = ?",item.TestTime.String(),item.PatientId,item.HeadType,item.SubType)
	if err !=nil{
		return err
	}

	if len(has)>0 {
		ids := has[0]
		if v, ok := ids["id"]; ok {
			_, err = fit.MySqlEngine().Table("NurseChat").ID(v).Update(item);
			fit.Logger().LogError("ghhhhhhhh",item.HeadType,v,err,*item)
		} else {
			_, err = fit.MySqlEngine().Insert(item);
		}
	}else{
		_, err = fit.MySqlEngine().Insert(item);
	}
	fit.Logger().LogError("QueryIntakeOrOutput",has,item.TestTime.String())

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
		err := fit.MySqlEngine().Omit("id").Table("NurseChat").Where("PatientId = ? and HeadType = ?", patientId, tp).Desc("Testtime").Limit(count, idx).Find(&responseObj)
		return responseObj, err
	} else {
		err := fit.MySqlEngine().Omit("id").Table("NurseChat").Where("PatientId = ? and HeadType = ?", patientId, tp).Desc("Testtime").Find(&responseObj)
		return responseObj, err
	}
}

// 所有的出入量
func QueryIntakeOrOutputVolumeAll(patientId string, page int) ([]IntakeOutputDup, error) {
	responseObj := make([]IntakeOutputDup, 0)
	if page >= 0 {
		count := 20
		idx := page * count
		err := fit.MySqlEngine().Omit("id").Table("NurseChat").Where("PatientId = ? ", patientId).And("(HeadType = 15 or HeadType = 16)").Desc("Testtime").Limit(count, idx).Find(&responseObj)
		return responseObj, err
	} else {
		err := fit.MySqlEngine().Omit("id").Table("NurseChat").Where("PatientId = ? ", patientId).And("(HeadType = 15 or HeadType = 16)").Desc("Testtime").Find(&responseObj)
		return responseObj, err
	}
}
