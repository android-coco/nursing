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

	//,部位左右下肢,1=A左,2=A右,3=P左,4=P右,5=A左,6=A右
	NRL01  int     `xorm:"comment(时间APN,1=A,2=P,3=N)"`
	NRL02A int     `xorm:"comment(腘动脉搏动，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL02B int     `xorm:"comment(腘动脉搏动，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL03A int     `xorm:"comment(足背动脉搏动，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL03B int     `xorm:"comment(足背动脉搏动，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL04A int     `xorm:"comment(下肢皮肤颜色，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL04B int     `xorm:"comment(下肢皮肤颜色，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL05A int     `xorm:"comment(Homans征，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL05B int     `xorm:"comment(Homans征，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL06A int     `xorm:"comment(疼痛尺或表情评分，左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL06B int     `xorm:"comment(疼痛尺或表情评分，右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL07A int     `xorm:"comment(肿胀，左下肢)"`
	NRL07B int     `xorm:"comment(肿胀，右下肢)"`
	NRL08A float32 `xorm:"comment(皮肤温度)"`
	NRL08B float32 `xorm:"comment(皮肤温度)"`
	NRL09A int     `xorm:"comment(主观麻痹感,左下肢，1=正常，2=变弱，3=不能触及)"`
	NRL09B int     `xorm:"comment(主观麻痹感,右下肢，1=正常，2=变弱，3=不能触及)"`
	NRL10A int     `xorm:"comment(大腿周径,左下肢)"`
	NRL10B int     `xorm:"comment(大腿周径,右下肢)"`
	NRL10C int     `xorm:"comment(小腿周径,左下肢)"`
	NRL10D int     `xorm:"comment(小腿周径,右下肢)"`
	NRL11  int     `xorm:"comment(评估意见,1=未发现问题,2=进一步评估,3=采取相应护理措施)"`
	NRL12  int     `xorm:"comment(护理措施，1、卧床休息2、抬高患肢，肢体位置高于心脏水平20~30cm3、膝关节微屈15°，腘窝处避免受压4、严禁按摩及热敷，避免下肢静脉穿刺5、指导踝泵锻炼6、监测外周循环情况7、监测D二聚体等实验室指标8、遵嘱使用抗凝药)"`
	Score  int     `xorm:"comment(总分)"`
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

func (m NRL5) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).Cols("DateTime","NRL01", "NRL02A", "NRL02B", "NRL03A", "NRL03B", "NRL04A", "NRL04B", "NRL05A", "NRL05B", "NRL06A", "NRL06B", "NRL07A", "NRL07B", "NRL08A", "NRL08B", "NRL09A", "NRL09B", "NRL10A", "NRL10B", "NRL10C", "NRL10D", "NRL11", "NRL12", "NRL13","Score").Update(&m)
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
