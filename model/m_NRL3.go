package model

import (
	"time"
	"fit"
	"fmt"
)

//基本生活活动能力BADL
type NRL3 struct {
	ID       int64     `xorm:"pk autoincr comment(文书id)"`
	VAA01    int64     `xorm:"comment(patientid病人id)"`
	BCK01    int       `xorm:"comment(classid科室id)"`
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
	NRL11    int       `xorm:"comment(建议护理级别,1=一级护理，2=二级护理，3=三级护理，9=特护)"`
	Score    string    `xorm:"comment(总分)"`
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
		_, err1 := fit.MySqlEngine().Table("NRL3").Where("VAA01 = ?", m.VAA01).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}

func (m *NRL3) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(m)
}

func (m *NRL3) DeleteData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).Delete(m)
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
func PCQueryNRL3(pid, datestr1, datestr2 string, pagenum int) ([]NRL3, error) {
	var mods []NRL3
	var err error

	if pagenum == -1 {
		err = fit.MySqlEngine().Table("NRL3").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Find(&mods)
	} else {
		if datestr2 == "" || datestr1 == "" {
			err = fit.MySqlEngine().Table("NRL3").Where("VAA01 = ?", pid).Limit(9, (pagenum-1)*9).Find(&mods)
		} else {
			err = fit.MySqlEngine().Table("NRL3").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum-1)*9).Find(&mods)
		}
	}
	if err != nil {
		return nil, err
	}
	for key, _ := range mods {
		mods[key].DateStr = mods[key].NormParse3()
	}
	return mods, nil
}

func PCQUeryNRLPageCount(nrlType, pid, datestr1, datestr2 string) (counts int64, err error) {
	tablename := "NRL" + nrlType
	if nrlType == "1" {
		//tablename = "NurseChat"
		if datestr2 == "" || datestr1 == "" {
			var count1, count2 int64
			count1, err = fit.MySqlEngine().Table("NurseChat").Where("PatientId = ?", pid).Distinct("TestTime", "NurseId").Count()
			if err != nil {
				return
			}
			count2, err = fit.MySqlEngine().Table("IOStatistics").Where("VAA01 = ?", pid).Count() // NurseId
			if err != nil {
				return
			}
			counts = count1 + count2
			fmt.Println("----------count :", count1, count2, counts)
		} else {
			counts, err = fit.MySqlEngine().Table(tablename).Where("PatientId = ? AND TestTime >= ? AND TestTime < ?", pid, datestr1, datestr2).Distinct("TestTime", "NurseId").Count()
			if err != nil {
				return
			}
		}
	} else if nrlType == "5" {
		if datestr2 == "" || datestr1 == "" {
			counts, err = fit.MySqlEngine().Table(tablename).Where("VAA01 = ?", pid).Distinct("DateTime").Count()
		} else {
			counts, err = fit.MySqlEngine().Table(tablename).Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Distinct("DateTime").Count()
		}
	} else {
		if datestr2 == "" || datestr1 == "" {
			counts, err = fit.MySqlEngine().Table(tablename).Where("VAA01 = ?", pid).Count()
		} else {
			counts, err = fit.MySqlEngine().Table(tablename).Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
		}
	}
	return counts, err
}

/*
func PCQueryNRL(pid, datestr1, datestr2, nrlType string, pagenum int, mods interface{}) (err error) {

	t := reflect.TypeOf(mods)
	if k := t.Kind(); k != reflect.Slice {  //判断是否是为切片类型
		err = errors.New("invalid type, you should type in slice")
		return
	}
	fmt.Println("**************", reflect.TypeOf(mods).Kind().String())
	v := reflect.ValueOf(mods)
	var mods11 interface{}
	fmt.Println("*******", mods11)
	switch nrlType {
	case "3":
		mods11 := v.Interface().([]NRL3)
		fmt.Println("#####", mods11)
	default:
		err = errors.New("invalid nrlType")
		return
	}
	fmt.Println("#####******", mods11)
	if datestr2 == "" || datestr1 == "" {
		err = fit.MySqlEngine().Table("NRL3").Where("VAA01 = ?", pid).Limit(9, (pagenum - 1) * 9).Find(&mods11)
	} else {
		err = fit.MySqlEngine().Table("NRL3").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum - 1) * 9).Find(&mods)
	}
	//if err != nil {
	//	return
	//}

	fmt.Println("+++",mods11, err)
	//for key,_ := range mods {
	//	mods[key].DateStr = mods[key].NormParse3()
	//}
	return
}*/
