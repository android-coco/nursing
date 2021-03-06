package model

import (
	"fit"
)

//压疮风险因素评估表
type NRL6 struct {
	ID        int64        `xorm:"pk autoincr comment(文书id)" fit:"rid"`
	PatientId int64        `xorm:"comment(patientid病人id)" fit:"pid"`
	NurseId   string       `xorm:"comment(NursingId责任护士ID)" fit:"uid"`
	NurseName string       `xorm:"comment(NursingName责任护士签名)" fit:"username"`
	DateTime  FitTime `xorm:"comment(记录时间)"`
	DateStr   string       `xorm:"-" fit:"-"`
	NRL01    string    `xorm:"comment(感觉知觉程度)"`
	NRL02    string    `xorm:"comment(潮湿情况)"`
	NRL03    string    `xorm:"comment(活动能力)"`
	NRL04    string    `xorm:"comment(体位变换能力)"`
	NRL05    string    `xorm:"comment(营养进食状况)"`
	NRL06    string    `xorm:"comment(摩擦力和剪切力)"`
	NRL07    string    `xorm:"comment(护理措施，低度风险)"`
	NRL08    string    `xorm:"comment(护理措施，中度风险)"`
	NRL09    string    `xorm:"comment(护理措施，高度风险)"`
	NRL10    string    `xorm:"comment(护理措施，非常风险)"`
	NRL11    string    `xorm:"comment(护理措施，潮湿管理)"`
	NRL12    string    `xorm:"comment(护理措施，营养管理)"`
	NRL13    string    `xorm:"comment(护理措施，摩擦力和剪切力的预防)"`
	//NRL14    string       `xorm:"comment(护理措施，其他护理注意事项)"`

	NRL15A string `xorm:"comment(审核护士id)"`
	NRL15B string `xorm:"comment(审核护士签名)"`

	Score string `xorm:"comment(总分)" fit:"score"`
}

/*func (m *NRL6) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL6").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL6").Where("PatientId = ?", m.PatientId).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m *NRL6) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(m)
}

func (m *NRL6) DeleteData(id int64) (int64, error) {
	DeleteNRecords(id)
	return fit.MySqlEngine().ID(id).Delete(m)
}

func QueryNRL6(rid string) (NRL6, error) {
	var nr6 NRL6
	_, err := fit.MySqlEngine().Table("NRL6").Where("id = ?", rid).Get(&nr6)
	if err != nil {
		return NRL6{}, err
	} else {
		//nr6.DateStr = nr6.DateTime.Format("2006-01-02")
		//nr6.TimeStr = nr6.DateTime.Format("15:04")
		return nr6, nil
	}
}*/

// pc端接口
func PCQueryNRL6(pid, datestr1, datestr2 string, pagenum int) ([]NRL6, error) {
	var mods []NRL6
	var err error

	if pagenum == -1 {
		err = fit.MySqlEngine().Table("NRL6").Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Asc("DateTime").Find(&mods)
	} else {
		if datestr2 == "" || datestr1 == "" {
			err = fit.MySqlEngine().Table("NRL6").Where("PatientId = ?", pid).Limit(4, (pagenum-1)*4).Asc("DateTime").Find(&mods)
		} else {
			err = fit.MySqlEngine().Table("NRL6").Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(4, (pagenum-1)*4).Asc("DateTime").Find(&mods)
		}
	}
	if err != nil {
		return nil, err
	}
	for key, _ := range mods {
		val := mods[key]
		mods[key].DateStr = val.DateTime.ParseDate()
	}
	return mods, nil
}
