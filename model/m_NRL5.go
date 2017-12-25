package model

import (
	"fit"
)

type APNModel struct {
	DateStr string
	ModelA  NRL5
	ModelP  NRL5
	ModelN  NRL5
}

//深静脉血栓护理观察单
type NRL5 struct {
	ID        int64        `xorm:"pk autoincr comment(文书id)" fit:"rid"`
	PatientId int64        `xorm:"comment(patientid病人id)" fit:"pid"`
	NurseId   string       `xorm:"comment(NursingId责任护士ID)" fit:"uid"`
	NurseName string       `xorm:"comment(NursingName责任护士签名)" fit:"username"`
	DateTime  FitTime `xorm:"comment(记录时间)"`
	DateStr   string       `xorm:"-" fit:"-"`
	//,部位左右下肢,1=A左,2=A右,3=P左,4=P右,5=A左,6=A右
	NRL01  string `xorm:"comment(时间APN,1=A,2=P,3=N)"`
	NRL02A string `xorm:"comment(腘动脉搏动，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL02B string `xorm:"comment(腘动脉搏动，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL03A string `xorm:"comment(足背动脉搏动，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL03B string `xorm:"comment(足背动脉搏动，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL04A string `xorm:"comment(下肢皮肤颜色，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL04B string `xorm:"comment(下肢皮肤颜色，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL05A string `xorm:"comment(Homans征，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL05B string `xorm:"comment(Homans征，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL06A string `xorm:"comment(疼痛尺或表情评分，左下肢)"`
	NRL06B string `xorm:"comment(疼痛尺或表情评分，右下肢)"`
	NRL07A string `xorm:"comment(肿胀，左下肢)"`
	NRL07B string `xorm:"comment(肿胀，右下肢)"`
	NRL08A string `xorm:"comment(皮肤温度)"`
	NRL08B string `xorm:"comment(皮肤温度)"`
	NRL09A string `xorm:"comment(主观麻痹感,左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL09B string `xorm:"comment(主观麻痹感,右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL10A string `xorm:"comment(大腿周径,左下肢)"`
	NRL10B string `xorm:"comment(大腿周径,右下肢)"`
	NRL10C string `xorm:"comment(小腿周径,左下肢)"`
	NRL10D string `xorm:"comment(小腿周径,右下肢)"`
	NRL11  string `xorm:"comment(评估意见,1=未发现问题,2=进一步评估,3=采取相应护理措施)"`
	NRL12  string `xorm:"comment(护理措施，1、卧床休息2、抬高患肢，肢体位置高于心脏水平20~30cm3、膝关节微屈15°，腘窝处避免受压4、严禁按摩及热敷，避免下肢静脉穿刺5、指导踝泵锻炼6、监测外周循环情况7、监测D二聚体等实验室指标8、遵嘱使用抗凝药)"`
	Score  string `xorm:"comment(总分)" fit:"score"`
}

/*func (m *NRL5) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL5").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL5").Where("PatientId = ?", m.PatientId).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.ID = rid
			return rid, err
		}
	}
	return 0, err
}*/

func (m *NRL5) IsExistNRL5() (has bool, err error) {
	var datestr = m.DateTime.ParseDate()
	has, err = fit.MySqlEngine().Table(new(NRL5)).Where("PatientId = ? AND NRL01 = ? AND DateTime = ?", m.PatientId, m.NRL01, datestr).Exist()
	return
}

/*func (m *NRL5) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(m)
}

func (m *NRL5) DeleteData(id int64) (int64, error) {
	DeleteNRecords(id)
	return fit.MySqlEngine().ID(id).Delete(m)
}

func QueryNRL5(rid string) (NRL5, error) {
	var nr5 NRL5
	_, err := fit.MySqlEngine().Table("NRL5").Where("id = ?", rid).Get(&nr5)
	if err != nil {
		return NRL5{}, err
	} else {
		//nr5.DateStr = nr5.DateTime.Format("2006-01-02")
		//nr5.TimeStr = nr5.DateTime.Format("15:04")
		return nr5, nil
	}
}*/

// pc端接口
func PCQueryNRL5(pid, datestr1, datestr2 string, pagenum int) ([]APNModel, error) {
	var mods []NRL5
	var err error

	if pagenum == -1 { // 打印用，获取全部数据
		err = fit.MySqlEngine().Table("NRL5").Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Asc("datetime", "NRL01").Find(&mods)
	} else { // pc 翻页用
		if datestr2 == "" || datestr1 == "" {
			err = fit.MySqlEngine().Table("NRL5").Where("PatientId = ?", pid).Asc("datetime", "NRL01").Find(&mods) //.Limit(5, (pagenum-1)*5)
		} else {
			err = fit.MySqlEngine().Table("NRL5").Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Asc("datetime", "NRL01").Find(&mods) // .Limit(5, (pagenum-1)*5)
		}
	}

	if err != nil {
		return nil, err
	}
	for key, _ := range mods {
		val := mods[key]
		mods[key].DateStr = val.DateTime.ParseDate()
	}
	list := mateNRL5Data(mods)
	return list, nil
}

func mateNRL5Data(mods []NRL5) []APNModel {
	var list []APNModel

	for ii := 0; ii < len(mods); ii++ {
		mod := mods[ii]
		if ii == 0 {
			var apnmod APNModel
			apnmod.DateStr = mod.DateStr
			mateAPNModel(&mod, &apnmod)

			list = append(list, apnmod)
		} else {
			oldapnmod := &list[len(list) - 1]
			if oldapnmod.DateStr == mod.DateStr {
				mateAPNModel(&mod, oldapnmod)
			} else {
				var apnmod APNModel
				apnmod.DateStr = mod.DateStr
				mateAPNModel(&mod, &apnmod)

				list = append(list, apnmod)
			}
		}
	}
	return list
}

func mateAPNModel(mod *NRL5, apnModel *APNModel)  {
	switch mod.NRL01 {
	case "1": apnModel.ModelA = *mod
	case "2": apnModel.ModelP = *mod
	case "3": apnModel.ModelN = *mod
	default:
		fit.Logger().LogError("error :**************", "FormatNRLData err : invalid type")
	}
}

/*查询是否存在某个班次的深静脉血栓护理观察单*/
func IsExistNRL5Shift(pid, datestr string) ([]string, error) {
	var slice []string
	err := fit.MySqlEngine().Table(new(NRL5)).Select("NRL01").Where("PatientId = ? AND DateTime = ?", pid, datestr).Find(&slice)
	return slice, err
}
