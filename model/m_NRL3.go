package model

import (
	"time"
	"fit"
	"errors"
	"strconv"
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


/*3-8 公用查询方法*/
/*pda 端， 编辑页 查看某一个文书*/
func QueryNRLWithRid(nrlType, rid string) (nrl interface{}, pid, uid string, err error) {
	tableName := "NRL" + nrlType
	switch nrlType {
	case "3":
		var nr3 NRL3
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr3)
		if err != nil {
			return
		} else {
			nr3.DateStr = nr3.DateTime.Format("2006-01-02")
			return nr3, strconv.FormatInt(nr3.VAA01, 10), nr3.BCE01A, nil
		}
	case "4":
		var nr4 NRL4
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr4)
		if err != nil {
			return
		} else {
			nr4.DateStr = nr4.DateTime.Format("2006-01-02")
			return nr4, strconv.FormatInt(nr4.VAA01, 10), nr4.BCE01A, nil
		}
	case "5":
		var nr5 NRL5
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr5)
		if err != nil {
			return
		} else {
			nr5.DateStr = nr5.DateTime.Format("2006-01-02")
			return nr5, strconv.FormatInt(nr5.VAA01, 10), nr5.BCE01A, nil
		}
	case "6":
		var nr6 NRL6
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr6)
		if err != nil {
			return
		} else {
			nr6.DateStr = nr6.DateTime.Format("2006-01-02")
			return nr6, strconv.FormatInt(nr6.VAA01, 10), nr6.BCE01A, nil
		}
	case "7":
		var nr7 NRL7
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr7)
		if err != nil {
			return
		} else {
			nr7.DateStr = nr7.DateTime.Format("2006-01-02")
			nr7.TimeStr = nr7.DateTime.Format("15:04")
			return nr7, strconv.FormatInt(nr7.VAA01, 10), nr7.BCE01A, nil
		}
	case "8":
		var nr8 NRL8
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr8)
		if err != nil {
			return
		} else {
			nr8.DateStr = nr8.DateTime.Format("2006-01-02")
			nr8.TimeStr = nr8.DateTime.Format("15:04")
			return nr8, strconv.FormatInt(nr8.VAA01, 10), nr8.BCE01A, nil
		}
	default:
		err = errors.New("QueryNRLWithRid : invalid nrltype type")
		return
	}


}

/*PC端*/
//查询页数
func PCQUeryNRLPageCount(nrlType, pid, datestr1, datestr2 string) (counts int64, err error) {
	tablename := "NRL" + nrlType
	if nrlType == "1" {
		return 0, errors.New("PCQUeryNRLPageCount: invalid type")
	} else if nrlType == "5" {
		return 0, errors.New("PCQUeryNRLPageCount: invalid type")
	} else {
		if datestr2 == "" || datestr1 == "" {
			counts, err = fit.MySqlEngine().Table(tablename).Where("VAA01 = ?", pid).Count()
		} else {
			counts, err = fit.MySqlEngine().Table(tablename).Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
		}
	}
	return counts, err
}


func PCQUeryNRLData(nrlType, pid, datestr1, datestr2 string) (counts int64, err error) {
	tablename := "NRL" + nrlType
	if nrlType == "1" {
		return 0, errors.New("PCQUeryNRLPageCount: invalid type")
	} else if nrlType == "5" {
		return 0, errors.New("PCQUeryNRLPageCount: invalid type")
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
