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
	DateStr  string    `xorm:"-"`
	TimeStr  string    `xorm:"-"`
	NRL01    string       `xorm:"comment(口诉言词量表VRS)"`
	NRL02    string       `xorm:"comment(数字评定量表NRS)"`
	NRL03    string       `xorm:"comment(面部表情量表WongBaker)"`
	NRL04    string       `xorm:"comment(面部表情量表FPSR)"`

	NRL05    string       `xorm:"comment(评分量表选择)"`
	NRL05A    string       `xorm:"comment(评分量值索引)"`
	Score    string       `xorm:"comment(评分量值)"`

	NRL06A   string       `xorm:"comment(审核护士id)"`
	NRL06B   string       `xorm:"comment(审核护士签名)"`
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
	//.Cols("DateTime","NRL01","NRL02","NRL03","NRL04","NRL05","NRL06A","NRL06B", "Score")
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(&m)
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

// pc端接口
func PCQueryNRL8(pid, datestr1, datestr2 string, pagenum int) ([]NRL8, error) {
	var mods []NRL8
	var err error
	if datestr2 == "" || datestr1 == "" {
		err = fit.MySqlEngine().Table("NRL8").Where("VAA01 = ?", pid).Limit(9, (pagenum - 1) * 9).Find(&mods)
	} else {
		err = fit.MySqlEngine().Table("NRL8").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum - 1) * 9).Find(&mods)
	}
	if err != nil {
		return nil, err
	}
	for key,_ := range mods {
		val := mods[key]
		mods[key].DateStr = val.DateTime.Format("2006-01-02")
		mods[key].TimeStr = val.DateTime.Format("15:04")
	}
	return mods, nil
}

func PCQUeryNRL8PageCount(pid, datestr1, datestr2 string) (counts int64, err error)  {
	if datestr2 == "" || datestr1 == "" {
		counts,err = fit.MySqlEngine().Table("NRL8").Where("VAA01 = ?", pid).Count()
	} else {
		counts,err = fit.MySqlEngine().Table("NRL8").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
	}
	return counts, err
}