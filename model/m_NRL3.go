package model

import (
	"time"
	"fit"
)

//基本生活活动能力BADL
type NRL3 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId责任护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName责任护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`
	DateStr  string    `xorm:"-"`
	NRL01    int       `xorm:"comment(大便评分)"`
	NRL02    int       `xorm:"comment(小便评分)"`
	NRL03    int       `xorm:"comment(装饰评分)"`
	NRL04    int       `xorm:"comment(如厕评分)"`
	NRL05    int       `xorm:"comment(进食评分)"`
	NRL06    int       `xorm:"comment(转移评分)"`
	NRL07    int       `xorm:"comment(活动评分)"`
	NRL08    int       `xorm:"comment(穿衣评分)"`
	NRL09    int       `xorm:"comment(上下楼梯评分)"`
	NRL10    int       `xorm:"comment(洗澡评分)"`
	NRL11    int       `xorm:"comment(建议护理级别)"`
	Score    int       `xorm:"comment(总分)"`
}


func (m *NRL3) NormParse() string {
	return m.DateTime.Format("2006-01-02 15:04:05")
}

func (m *NRL3) NormParse2() string {
	return m.DateTime.Format("2006-01-02 15:04")
}

func (m *NRL3) NormParse3() string {
	return m.DateTime.Format("2006-01-02")
}


func (m *NRL3) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL3").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL3").Where("patientid = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m NRL3) UpdateData(id int64) (int64, error) {
	// "BCK01", "VAA01", "BCE01A", "BCE03A",
	return fit.MySqlEngine().ID(id).Cols( "DateTime", "NRL01", "NRL02", "NRL03", "NRL04", "NRL05", "NRL06", "NRL07", "NRL08", "NRL09", "NRL10", "NRL11", "NRL12","Score").Update(&m)
}

func QueryNRL3(rid string) (NRL3, error) {
	var nr3 NRL3
	_, err := fit.MySqlEngine().Table("NRL3").Where("id = ?", rid).Get(&nr3)
	if err != nil {
		return NRL3{}, err
	} else {
		return nr3, nil
	}
}

// pc端接口
func PCQueryNRL3(pid string, pagenum int) ([]NRL3, error) {
	var mods []NRL3
	err := fit.MySqlEngine().Table("NRL3").Where("VAA01 = ?", pid).Limit(9, (pagenum - 1) * 9).Find(&mods)
	if err != nil {
		return nil, err
	}
	for key,_ := range mods {
		mods[key].DateStr = mods[key].NormParse3()
	}
	return mods, nil
}

func PCQUeryNRL3PageCount(pid string) (int, error)  {
	counts,err := fit.MySqlEngine().Table("NRL3").Where("VAA01 = ?", pid).Count()
	return int(counts), err
}
