package model

import (
	"time"
	"fit"
)

//患者跌倒风险评估护理单
type NRL7 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId评估护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName评估护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`

	NRL08  int    `xorm:"comment(陪护情况)"`
	NRL08A string `xorm:"comment(陪护情况)"`
	NRL01  int    `xorm:"comment(跌倒史)"`
	NRL02  int    `xorm:"comment(超过1个医学诊断)"`
	NRL03  int    `xorm:"comment(使用助行器具)"`
	NRL04  int    `xorm:"comment(静脉输液/肝素锁)"`
	NRL05  int    `xorm:"comment(步态)"`
	NRL06  int    `xorm:"comment(认知状态)"`
	NRL07  int    `xorm:"comment(护理措施)"`
	Score  int    `xorm:"comment(总分)"`
}

func (m *NRL7) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL7").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL7").Where("patientid = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m NRL7) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).Cols("DateTime","NRL01", "NRL02", "NRL03", "NRL04", "NRL05", "NRL06", "NRL07", "NRL08", "NRL08A", "NRL09", "Score").Update(&m)
}

func QueryNRL7(rid string) (NRL7, error) {
	var nr7 NRL7
	_, err := fit.MySqlEngine().Table("NRL7").Where("id = ?", rid).Get(&nr7)
	if err != nil {
		return NRL7{}, err
	} else {
		return nr7, nil
	}
}
