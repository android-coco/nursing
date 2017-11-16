package model

import (
	"time"
	"fit"
	"fmt"
)

//患者跌倒风险评估护理单 表头
type NRL7Title struct {
	ID    int64 `xorm:"pk autoincr comment(id)"`
	BCK01 int64 `xorm:"comment(classid科室id)"`
	VAA01 int64 `xorm:"comment(patientid病人id)"`
	DateTime time.Time `xorm:"comment(记录时间)"`


	NRT01  string `xorm:"comment(表头25，11=固定，其他空值)"`
	NRT01V string `xorm:"comment(表头25，自定义内容)"`
	NRT02  string `xorm:"comment(表头26，11=固定，其他空值)"`
	NRT02V string `xorm:"comment(表头26，自定义内容)"`
}

//患者跌倒风险评估护理单
type NRL7 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId评估护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName评估护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`
	DateStr  string    `xorm:"-"`
	NRL08    string    `xorm:"comment(陪护情况，有无)"`
	NRL08A   string    `xorm:"comment(陪护情况，家属，保姆，护工)"`
	NRL08B   string    `xorm:"comment(陪护情况,其他)"`
	NRL01    string    `xorm:"comment(跌倒史)"`
	NRL02    string    `xorm:"comment(超过1个医学诊断)"`
	NRL03    string    `xorm:"comment(使用助行器具)"`
	NRL04    string    `xorm:"comment(静脉输液/肝素锁)"`
	NRL05    string    `xorm:"comment(步态)"`
	NRL06    string    `xorm:"comment(认知状态)"`
	NRL07    string    `xorm:"comment(护理措施)"`
	NRL09A   string    `xorm:"comment(审核护士id)"`
	NRL09B   string    `xorm:"comment(审核护士签名)"`

	Score string `xorm:"comment(总分)"`
}

func (m *NRL7) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL7").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL7").Where("VAA01 = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m *NRL7) UpdateData(id int64) (int64, error) {
	//.Cols("DateTime", "NRL01", "NRL02", "NRL03", "NRL04", "NRL05", "NRL06", "NRL07", "NRL08", "NRL08A", "NRL09", "Score")
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(m)
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

// pc端接口
func (m *NRL7Title) PCQueryNRL7Title() (err error) {
	_, err = fit.MySqlEngine().Table("NRL7Title").Where("VAA01 = ? AND BCK01 = ?", m.VAA01, m.BCK01).Get(m)
	if err != nil {
		return err
	}
	return nil
}

/*func (m *NRL7Title) InsertTitle() error {
	_, err := fit.MySqlEngine().Table("NRL7Title").Insert(m)
	if err != nil {
		return err
	}
	return nil
}*/

func (m *NRL7Title) UpdateTitle()  error {
	has, err := fit.MySqlEngine().Table("NRL7Title").Where("VAA01 = ? AND BCK01 = ?", m.VAA01, m.BCK01).Exist()
	if err != nil {
		return err
	}
	if has {
		var mod NRL7Title
		_, err = fit.MySqlEngine().Table("NRL7Title").Where("VAA01 = ? AND BCK01 = ?", m.VAA01, m.BCK01).Get(&mod)
		if err != nil {
			return err
		}
		if mod.NRT01 == "11" {
			_, err = fit.MySqlEngine().Table("NRL7Title").Cols( "NRT02","NRT02V").Update(m)
		} else if mod.NRT02 == "11" {
			_, err = fit.MySqlEngine().Table("NRL7Title").Cols("NRT01","NRT01V").Update(m)
		} else if mod.NRT01 == "11" && mod.NRT02 == "11" {
			_, err = fit.MySqlEngine().Table("NRL7Title").Cols("NRT01", "NRT01V", "NRT02", "NRT02V").Update(m)
		} else {
			return fmt.Errorf("can not update NRL7Title")
		}
	} else {
		_, err = fit.MySqlEngine().Table("NRL7Title").Insert(m)
	}

	if err != nil {
		return err
	}
	return nil
}

func PCQueryNRL7(pid, datestr1, datestr2 string, pagenum int) ([]NRL7, error) {
	var mods []NRL7
	var err error
	if datestr2 == "" || datestr1 == "" {
		err = fit.MySqlEngine().Table("NRL7").Where("VAA01 = ?", pid).Limit(9, (pagenum-1)*9).Find(&mods)
	} else {
		err = fit.MySqlEngine().Table("NRL7").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum-1)*9).Find(&mods)
	}
	if err != nil {
		return nil, err
	}
	for key, _ := range mods {
		val := mods[key]
		mods[key].DateStr = val.DateTime.Format("2006-01-02")
	}
	return mods, nil
}

func PCQUeryNRL7PageCount(pid, datestr1, datestr2 string) (counts int64, err error) {
	if datestr2 == "" || datestr1 == "" {
		counts, err = fit.MySqlEngine().Table("NRL7").Where("VAA01 = ?", pid).Count()
	} else {
		counts, err = fit.MySqlEngine().Table("NRL7").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
	}
	return counts, err
}
