package model

import (
	"time"
	"fit"
)

type Warn struct {
	BaseModel 				`xorm:"extends"`
	Name     string        	`json:"name" xorm:"comment(提醒名称)"`
	WarnTime time.Time      `json:"warn_time" xorm:"notnull comment(提醒时间)"`
}

func (m Warn) InsertData() (int64, error)  {
	id, err := fit.Engine().Insert(m)
	return id, err
}

func Warnlist(nurse_id, patient_id string) []Warn {
	var warns []Warn
	fit.Engine().SQL("select * from warn where nurse_id = ? and patient_id = ?", nurse_id, patient_id).Find(&warns)
	return warns
}
