package model

import (
	"time"
	"fit"
)

//压疮风险因素评估表
type NRL6 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId责任护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName责任护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`
	NRL01    int       `xorm:"comment(感觉知觉程度)"`
	NRL02    int       `xorm:"comment(潮湿情况)"`
	NRL03    int       `xorm:"comment(活动能力)"`
	NRL04    int       `xorm:"comment(体位变换能力)"`
	NRL05    int       `xorm:"comment(营养进食状况)"`
	NRL06    int       `xorm:"comment(摩擦力和剪切力)"`
	NRL07A   int       `xorm:"comment(责任护士id)"`
	NRL07B   int       `xorm:"comment(责任护士签名)"`
	Score    int       `xorm:"comment(总分)"`
}


func (m *NRL6) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL6").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL6").Where("patientid = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m NRL6) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).Cols("DateTime","NRL01","NRL02","NRL03","NRL04","NRL05","NRL06","NRL07A","NRL07B","NRL08","Score").Update(&m)
}

func QueryNRL6(rid string) (NRL6, error)  {
	var nr6 NRL6
	_, err := fit.MySqlEngine().Table("NRL6").Where("id = ?", rid).Get(&nr6)
	if err != nil {
		return NRL6{}, err
	} else {
		return nr6, nil
	}
}