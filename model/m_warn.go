package model

import (
	"fit"
	"fmt"
	"strings"
)

type WarnType int

type Warn struct {
	Id       int64    `json:"id" xorm:"pk autoincr "`
	ClassId  string   `json:"class_id" xorm:"comment(科室id)"`
	Name     string   `json:"name" xorm:"comment(提醒标签)"`
	Desc     string   `json:"desc" xorm:"comment(提醒描述)"`
	WarnTime string   `json:"warn_time" xorm:"notnull comment(提醒时间)"`
	WarnType WarnType `json:"warn_type" xorm:"comment(提醒类型，1=响铃，2=震动，4=响铃+震动)"`
}

const (
	WarnTypeRing  WarnType = 1 << iota
	WarnTypeShake
	WarnTypeAll
)

func (m Warn) ParseWarnType(value string) (WarnType, error) {
	switch value {
	case "1":
		return WarnTypeRing, nil
	case "2":
		return WarnTypeShake, nil
	case "4":
		return WarnTypeAll, nil
	default:
		return WarnTypeAll, fmt.Errorf("not fund such of WarnType: %v", value)
	}
}

func (m Warn) InsertData() (int64, error) {
	fit.Logger().LogInfo("InsertData()InsertData()InsertData()InsertData()")
	has, err1 := fit.MySqlEngine().Where("WarnTime = ?", m.WarnTime).Exist(&Warn{})
	if has {
		return 0, fmt.Errorf("there has the same warn")
	} else if err1 != nil {
		return 0, err1
	}

	id, err := fit.MySqlEngine().Insert(m)
	return id, err
}

func Warnlist(classId string) []Warn {
	var warns []Warn
	//params := time.Now().Format("2006-01-02") + "%"
	fit.MySqlEngine().SQL("select * from Warn where ClassId = ?", classId).Find(&warns)
	//if listType == "1" { // 已完成
	//} else { // 待执行
	//	fit.MySqlEngine().SQL("select * from Warn where ClassId = ? and WarnTime > ?", classId, time.Now().Format("2006-01-02 15:04")).Find(&warns)
	//}
	return warns
}

func WarnAll(classId string) []Warn {
	var warns []Warn
	fit.MySqlEngine().SQL("select * from Warn where ClassId = ? ORDER BY warntime DESC", classId).Find(&warns)
	return warns
}

func (m Warn) DeleteWarn() (int64, error) {
	//// 入院时间
	//var VAA73 time.Time
	//// 出院时间
	//var VAA74 time.Time
	//fit.SQLServerEngine().SQL("select VAA73,VAA74 form VAA1 where")

	// id 更新条数

	valSlice := strings.Split(m.WarnTime, ",")

	//id, err := fit.MySqlEngine().SQL("DELETE FROM Warn WHERE classid = ?", m.ClassId).In("WarnTime", "2017-10-09 08:00:00,2017-10-09 08:00").Delete(&m)
	id, err := fit.MySqlEngine().Where("classid = ?", m.ClassId).In("WarnTime", valSlice).Delete(Warn{})
	return id, err
}

func DelWarn(wid string) (int64, error) {
	valSlice := strings.Split(wid, ",")
	id, err := fit.MySqlEngine().In("id", valSlice).Delete(new(Warn))
	return id, err
}

//更新提醒
func (m Warn) ModifyWarn() (int64, error) {
	//affected, err := fit.MySqlEngine().Id(RecordId).Cols("Updated").Update(&nursing)
	affected, err := fit.MySqlEngine().Table(new(Warn)).ID(m.Id).Update(&m)
	return affected, err
}
