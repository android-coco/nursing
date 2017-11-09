package model

import (
	"time"
	"fit"
)

//疼痛强度评分量表
type NRL8 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId评估护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName评估护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`

	Score    int       `xorm:"comment(面部表情量表FPSR)"`
	NRL01    int       `xorm:"comment(口诉言词量表VRS)"`
	NRL02    int       `xorm:"comment(数字评定量表NRS)"`
	NRL03    int       `xorm:"comment(面部表情量表WongBaker)"`
	NRL04    int       `xorm:"comment(面部表情量表FPSR)"`
}

func (m *NRL8) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL8").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL8").Where("patientid = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m NRL8) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).Cols("DateTime","NRL01","NRL02","NRL03","NRL04","NRL05", "Score").Update(&m)
}

func QueryNRL8(rid string) (NRL8, error)  {
	var nr8 NRL8
	_, err := fit.MySqlEngine().Table("NRL8").Where("id = ?", rid).Get(&nr8)
	if err != nil {
		return NRL8{}, err
	} else {
		return nr8, nil
	}
}