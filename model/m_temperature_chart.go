package model

import (
	"fit"
	"errors"

	"strings"
	"time"
	"github.com/go-xorm/xorm"
)

//体温单数据模型
type TemperatrureChat struct {
	Id        int64        `json:"id" xorm:"pk autoincr"`
	HeadType  string       `json:"headtype" xorm:"notnull comment(头部类型)"`
	DateTime  fit.JsonTime `json:"testtime" xorm:"notnull comment(日期时间)"`
	TestTime  fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	TypeTime  int          `json:"typetime" xorm:"notnull comment(时间段)"`
	SubType   int          `json:"type" xorm:"notnull comment(类型,)"`
	Other     int          `json:"other" xorm:"notnull comment(其他可能选项,)"`
	Value     string       `json:"value" xorm:"notnull comment(值)"`
	PatientId int64        `json:"patientid" xorm:"notnull comment(病人id)"`
	NurseId   int          `json:"nurseid" xorm:"notnull comment(护士id)"`
	NurseName string       `json:"nursename" xorm:"notnull comment(护士姓名)"`
}

//查询是否有体温数据
func WhetherTemperature(sql string, msg ...interface{}) (bool, error) {
	return fit.MySqlEngine().Table("TemperatrureChat").Where(sql, msg...).Exist()
}

//得到是否有待测体温
func GetTemperatureWhetherMeasured(testtime string, pid int64, interval string) (bool, error) {
	return fit.MySqlEngine().Table("TemperatrureChat").Where("DateTime = ? and PatientId = ? and HeadType = 1 and TypeTime = ?", testtime, pid, interval).Exist()
}

//返回体温单数据
func OutTemperatureHistory(sql string, msg ...interface{}) ([]TemperatrureChat, error) {
	items := make([]TemperatrureChat, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

//根据id删除体温单数据
func DeleteTemperatureHistory(session *xorm.Session, id string) error {
	var items TemperatrureChat
	_, err := session.Table("TemperatrureChat").ID(id).Delete(&items)
	return err
}

//根据id更新体温单数据
func UpdateTemperatrureChat(session *xorm.Session, id string, mas map[string]interface{}) (error) {
	_, err := session.Table(new(TemperatrureChat)).ID(id).Update(mas)
	return err
}

//获取体温表数据根据周
func GetTemperatureChatData(tp int, patienttd int64, weeks []time.Time) ([]string, error) {
	var value []string
	for _, v := range weeks {
		data := v.Format("2006-01-02")
		var sql string
		var msg []interface{}
		typetimes := make([]int, 0)
		switch tp {
		case 1: //口表
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and SubType = 3"
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 4, 8, 12, 16, 20, 24)
		case 2: //掖表
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and SubType = 1"
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 4, 8, 12, 16, 20, 24)
		case 3: //肛表
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and SubType = 4"
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 4, 8, 12, 16, 20, 24)
		case 4: //脉搏
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 2 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 4, 8, 12, 16, 20, 24)
		case 5: //心率
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 5 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 4, 8, 12, 16, 20, 24)
		case 6: //呼吸
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 3 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 4, 8, 12, 16, 20, 24)
		case 7: //输入液量
			sql = "SELECT IntakeA FROM IOStatistics WHERE DateTime2 = ? and VAA01 = ? "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 8: //排出大便
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 13 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 9: //排出尿量
			sql = "SELECT OutputA FROM IOStatistics WHERE DateTime2 = ? and VAA01 = ? "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 10: //排出其他
			sql = "SELECT OutputC FROM IOStatistics WHERE DateTime2 = ? and VAA01 = ? "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 11: //血压
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 4 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 12: //体重
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 8 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 13: //皮试
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 10 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 14: //其他
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 14 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		case 15: //事间
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 12 "
			msg = append(msg, data, patienttd)
			typetimes = append(typetimes, 0)
		}

		for _, k := range typetimes {
			sql1 := sql
			msg1 := msg
			if len(typetimes) != 1 {
				sql1 = sql1 + " and TypeTime = ?"
				msg1 = append(msg1, k)
			}

			maps, err := fit.MySqlEngine().QueryString(sql1, msg1...)
			//fmt.Println("---------", maps)

			if err != nil {
				return value, err
			} else {
				if len(maps) > 0 {
					for _, dict := range maps {
						if v, ok := dict["Value"]; ok {
							value = append(value, v)
						} else {
							value = append(value, "")
						}
					}
				} else {
					value = append(value, "")
				}
			}
		}
	}

	return value, nil
}

type TimeShow interface {
	ShowTime() (time.Time)
}

//体温单返回数据结构
type TemperatrureChatHistory struct {
	PatientAge     string `json:"patient_age"`
	PatientBed     string `json:"patient_bed"`
	PatientId      int64  `json:"patient_id"`
	PatientName    string `json:"patient_name"`
	TestTime       string `json:"test_time"`  //测量时间
	DateTime       string `json:"date_time"`  //日期时间
	TimeFrame      string `json:"time_frame"` //时段
	ThmId          int64  `json:"thm_id"`
	ThmValue       string `json:"thm_value"`
	ThmType        int    `json:"thm_type"`
	ThmScene       int    `json:"thm_scene"`
	PulseId        int64  `json:"pulse_id"`
	PulseValue     string `json:"pulse_value"`
	PulseBriefness int    `json:"pulse_briefness"`
	BreatheId      int64  `json:"breathe_id"`
	BreatheValue   string `json:"breathe_value"`
	BreatheScene   int    `json:"breathe_scene"`
	ShitId         int64  `json:"shit_id"`
	ShitValue      string `json:"shit_value"`
	ShitScene      int    `json:"shit_scene"`
	PressureId     int64  `json:"pressure_id"`
	PressureSys    string `json:"pressure_sys"`
	PressureDia    string `json:"pressure_dia"`
	PressureScene  int    `json:"pressure_scene"`
	HeartrateId    int64  `json:"heartrate_id"`
	HeartrateValue string `json:"heartrate_value"`
	WeightId       int64  `json:"weight_id"`
	WeightValue    string `json:"weight_value"`
	WeightScene    int    `json:"weight_scene"`
	HeightId       int64  `json:"height_id"`
	HeightValue    string `json:"height_value"`
	HeightScene    int    `json:"height_scene"`
	SkinId         int64  `json:"skin_id"`
	SkinValue      string `json:"skin_value"`
	OtherId        int64  `json:"other_id"`
	OtherValue     string `json:"other_value"`
	IncidentId     int64  `json:"incident_id"`
	IncidentScene  int    `json:"incident_scene"`
	IncidentTime   string `json:"incident_time"`
}

func (v TemperatrureChatHistory) ShowTime() (time.Time) {
	test_time, _ := time.ParseInLocation("2006-01-02 15:04:05", v.TestTime, time.Local)
	return test_time
}

type PersonSlice []TemperatrureChatHistory

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].ShowTime().Unix() > a[i].ShowTime().Unix()
}

//数据合并返回
func TransformTemperatrureCH(chat TemperatrureChat, history *TemperatrureChatHistory) (error, bool) {
	if chat.HeadType == "" {
		return errors.New("没有类型"), false
	}

	switch chat.HeadType {
	case Temperature_Type:
		if history.ThmValue != "" || history.ThmScene != 0 {
			return nil, true
		}
		history.ThmId = chat.Id
		history.ThmValue = chat.Value
		history.ThmType = chat.SubType
		history.ThmScene = chat.Other
	case Pulse_Type:
		history.PulseId = chat.Id
		history.PulseValue = chat.Value
		history.PulseBriefness = chat.Other
	case Breathe_Type:
		history.BreatheId = chat.Id
		history.BreatheValue = chat.Value
		history.BreatheScene = chat.Other
	case Pressure_Type:
		history.PressureId = chat.Id
		if chat.Value != "" {
			value := strings.Split(chat.Value, "/")
			if len(value) == 2 {
				history.PressureSys = value[0]
				history.PressureDia = value[1]
			} else {
				history.PressureSys = ""
				history.PressureDia = ""
			}
		} else {
			history.PressureSys = ""
			history.PressureDia = ""
		}
		history.PressureScene = chat.Other
	case Heartrate_Type:
		history.HeartrateId = chat.Id
		history.HeartrateValue = chat.Value
	case Weight_Type:
		history.WeightId = chat.Id
		history.WeightValue = chat.Value
		history.WeightScene = chat.Other
	case Height_Type:
		history.HeightId = chat.Id
		history.HeightValue = chat.Value
		history.HeightScene = chat.Other
	case Skin_Type:
		history.SkinId = chat.Id
		history.SkinValue = chat.Value
	case Incident_Type:
		if history.IncidentScene != 0 {
			return nil, true
		}
		history.IncidentId = chat.Id
		history.IncidentScene = chat.Other
		history.IncidentTime = chat.TestTime.String()
	case Shit_Type:
		if history.ShitValue != "" {
			return nil, true
		}
		history.ShitId = chat.Id
		history.ShitValue = chat.Value
		history.ShitScene = chat.Other
	case Other_Type:
		history.OtherId = chat.Id
		history.OtherValue = chat.Value
	default:
		return errors.New("未知类型"), false
	}
	return nil, false
}

func BackIntervalTime(Interval int) string {
	switch Interval {
	case 4:
		return "04:00:00"
	case 8:
		return "08:00:00"
	case 12:
		return "12:00:00"
	case 16:
		return "16:00:00"
	case 20:
		return "20:00:00"
	case 24:
		return "23:59:59"
	default:
		return "00:00:00"
	}
}

//返回病人结构
type PatientHistory struct {
	PatientId int64  `json:"patient_id"` // 病人ID
	BedCoding string `json:"bed_coding"` // 床号 BCQ1表
	Name      string `json:"name"`       // 姓名
	Checked   int    `json:"checked"`    //是否选中  0没有选中 1选中
	Age       string `json:"age"`        // 年龄
}
