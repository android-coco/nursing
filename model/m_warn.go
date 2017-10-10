package model

import (

	"fit"
	"fmt"
	"time"
)

type WarnType int

type Warn struct {
	BaseModel 				`xorm:"extends"`
	ClassId        string        `json:"class_id" xorm:"comment(科室id)`
	Name     string        	`json:"name" xorm:"comment(提醒标签)"`
	Desc     string        	`json:"desc" xorm:"comment(提醒描述)"`
	WarnTime string      `json:"warn_time" xorm:"notnull comment(提醒时间)"`
	WarnType WarnType		`json:"warn_type" xorm:"comment(提醒类型，1=响铃，2=震动，3=响铃+震动)`
}

const (
	WarnTypeRing WarnType = 1 << iota
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

func (m Warn) InsertData() (int64, error)  {
	has, err1 := fit.MySqlEngine().Where("WarnTime = ?", m.WarnTime).Exist(&m)
	if has {
		return 0, fmt.Errorf("there has the same warn")
	} else if err1 != nil {
		return 0, err1
	}

	id, err := fit.MySqlEngine().Insert(m)
	return id, err
}

func Warnlist(classId, listType string) []Warn {
	var warns []Warn
	if listType == "1" { // 已完成
		fit.MySqlEngine().SQL("select * from Warn where ClassId = ? and WarnTime < ?", classId, time.Now().Format("2006-01-02 15:04")).Find(&warns)
	} else { // 待执行
		fit.MySqlEngine().SQL("select * from Warn where ClassId = ? and WarnTime > ?", classId, time.Now().Format("2006-01-02 15:04")).Find(&warns)
	}
	return warns
}

func (m Warn)DeleteWarn()(int64, error)  {
	// id 更新条数
	id, err := fit.MySqlEngine().SQL("DELETE FROM Warn WHERE classid = ? and WarnTime = ?", m.ClassId, m.WarnTime).Delete(&m)
	return id, err
}
