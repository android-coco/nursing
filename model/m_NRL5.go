package model

import (
	"time"
	"fit"
)

//深静脉血栓护理观察单
type NRL5 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	BCK01    int64     `xorm:"comment(classid科室id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCE01A   string    `xorm:"comment(NursingId责任护士ID)"`
	BCE03A   string    `xorm:"comment(NursingName责任护士签名)"`
	DateTime time.Time `xorm:"comment(记录时间)"`
	DateStr  string    `xorm:"-"`
	//,部位左右下肢,1=A左,2=A右,3=P左,4=P右,5=A左,6=A右
	NRL01  string     `xorm:"comment(时间APN,1=A,2=P,3=N)"`
	NRL02A string     `xorm:"comment(腘动脉搏动，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL02B string     `xorm:"comment(腘动脉搏动，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL03A string     `xorm:"comment(足背动脉搏动，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL03B string     `xorm:"comment(足背动脉搏动，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL04A string     `xorm:"comment(下肢皮肤颜色，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL04B string     `xorm:"comment(下肢皮肤颜色，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL05A string     `xorm:"comment(Homans征，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL05B string     `xorm:"comment(Homans征，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL06A string     `xorm:"comment(疼痛尺或表情评分，左下肢)"`
	NRL06B string     `xorm:"comment(疼痛尺或表情评分，右下肢)"`
	NRL07A string     `xorm:"comment(肿胀，左下肢)"`
	NRL07B string     `xorm:"comment(肿胀，右下肢)"`
	NRL08A string 		`xorm:"comment(皮肤温度)"`
	NRL08B string 		`xorm:"comment(皮肤温度)"`
	NRL09A string     `xorm:"comment(主观麻痹感,左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL09B string     `xorm:"comment(主观麻痹感,右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL10A string     `xorm:"comment(大腿周径,左下肢)"`
	NRL10B string     `xorm:"comment(大腿周径,右下肢)"`
	NRL10C string     `xorm:"comment(小腿周径,左下肢)"`
	NRL10D string     `xorm:"comment(小腿周径,右下肢)"`
	NRL11  string     `xorm:"comment(评估意见,1=未发现问题,2=进一步评估,3=采取相应护理措施)"`
	NRL12  string     `xorm:"comment(护理措施，1、卧床休息2、抬高患肢，肢体位置高于心脏水平20~30cm3、膝关节微屈15°，腘窝处避免受压4、严禁按摩及热敷，避免下肢静脉穿刺5、指导踝泵锻炼6、监测外周循环情况7、监测D二聚体等实验室指标8、遵嘱使用抗凝药)"`
	Score  string     `xorm:"comment(总分)"`
}



func (m *NRL5) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL5").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL5").Where("patientid = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m *NRL5) IsExistNRL5() (has bool, err error)  {
	var datestr = m.DateTime.Format("2006-01-02")
	has, err = fit.MySqlEngine().Table(new(NRL5)).Where("VAA01 = ? AND DateTime = ?", m.VAA01, datestr).Exist()
	return
}

func (m NRL5) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(&m)
}

func QueryNRL5(rid string) (NRL5, error) {
	var nr5 NRL5
	_, err := fit.MySqlEngine().Table("NRL5").Where("id = ?", rid).Get(&nr5)
	if err != nil {
		return NRL5{}, err
	} else {
		return nr5, nil
	}
}

// pc端接口
func PCQueryNRL5(pid, datestr1, datestr2 string, pagenum int) ([]NRL5, error) {
	var mods []NRL5
	var err error
	if datestr2 == "" || datestr1 == "" {
		err = fit.MySqlEngine().Table("NRL5").Where("VAA01 = ?", pid).Limit(9, (pagenum - 1) * 9).Asc("datetime", "NRL01").Find(&mods)
	} else {
		err = fit.MySqlEngine().Table("NRL5").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum - 1) * 9).Find(&mods)
	}
	if err != nil {
		return nil, err
	}
	for key,_ := range mods {
		val := mods[key]
		mods[key].DateStr = val.DateTime.Format("2006-01-02")
	}
	return mods, nil
}

func PCQUeryNRL5PageCount(pid, datestr1, datestr2 string) (counts int64, err error)  {
	if datestr2 == "" || datestr1 == "" {
		counts,err = fit.MySqlEngine().Table("NRL5").Where("VAA01 = ?", pid).Count()
	} else {
		counts,err = fit.MySqlEngine().Table("NRL5").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
	}
	return counts, err
}


/*查询是否存在某个班次的深静脉血栓护理观察单*/
func IsExistNRL5Shift(pid, datestr string) ([]string, error) {
	var slice []string
	err := fit.MySqlEngine().Table(new(NRL5)).Select("NRL01").Where("VAA01 = ? AND DateTime = ?", pid, datestr).Find(&slice)
	return slice, err
}
