package model

import (
	"time"
	"fit"
)

//深静脉血栓DVT形成风险评估表
type NRL4 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId责任护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName责任护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`
	NRL01    int       `xorm:"comment(年龄评分)"`
	NRL02    int       `xorm:"comment(体重评分)"`
	NRL03    int       `xorm:"comment(运动能力评分)"`
	NRL04    int       `xorm:"comment(特殊风险种类)"`
	NRL05    int       `xorm:"comment(创伤风险种类，术前)"`
	NRL06    int       `xorm:"comment(外科干预)"`
	NRL07    int       `xorm:"comment(现有的高风险疾病)"`
	Score    int       `xorm:"comment(总分)"`
}


func (m *NRL4) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL4").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL4").Where("patientid = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m NRL4) UpdateData(id int64) (int64, error) {
	// "BCK01","VAA01","BCE01A","BCE03A",
	return fit.MySqlEngine().Id(id).Cols("DateTime","NRL01","NRL02","NRL03","NRL04","NRL05","NRL06","NRL07","NRL08","Score").Update(&m)
}

func QueryNRL4(rid string) (NRL4, error)  {
	var nr4 NRL4
	_, err := fit.MySqlEngine().Table("NRL4").Where("id = ?", rid).Get(&nr4)
	if err != nil {
		return NRL4{}, err
	} else {
		return nr4, nil
	}
}