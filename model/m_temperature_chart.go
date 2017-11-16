package model

import (
	"fit"
	"errors"

	"strings"
	"time"
	"fmt"
)

type TemperatrureChat struct {
	Id          int64		  `json:"id" xorm:"pk autoincr"`
	HeadType    string        `json:"headtype" xorm:"notnull comment(头部类型)"`
	DateTime    fit.JsonTime  `json:"testtime" xorm:"notnull comment(日期时间)"`
	TestTime    fit.JsonTime  `json:"testtime" xorm:"notnull comment(测试时间)"`
	TypeTime    int           `json:"typetime" xorm:"notnull comment(时间段)"`
	Type        int           `json:"type" xorm:"notnull comment(类型,)"`
	Other       int           `json:"other" xorm:"notnull comment(其他可能选项,)"`
	Value       string        `json:"value" xorm:"notnull comment(值)"`
	PatientId   int           `json:"patientid" xorm:"notnull comment(病人id)"`
	NurseId     int           `json:"nurseid" xorm:"notnull comment(护士id)"`
	NurseName   string        `json:"nursename" xorm:"notnull comment(护士姓名)"`
}

func OutTemperatureHistory(sql string, msg ...interface{})([]TemperatrureChat, error){
	items := make([]TemperatrureChat, 0)
	//fit.MySqlEngine().ShowSQL(true)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)

	return items, err
}

//获取体温表数据
func GetTemperatureChatData(tp int,patienttd int,weeks []time.Time) (error,[]string) {
    var value []string
    for _,v := range weeks {
		data := v.Format("2006-01-02 15:04:05")
		var sql string
		var msg []interface{}
		typetimes := make([]int,0)
		switch tp {
		case 1:  //口表
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and Type = 3"
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,4,8,12,16,20,24)
		case 2:  //掖表
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and Type = 1"
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,4,8,12,16,20,24)
		case 3:  //肛表
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and Type = 4"
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,4,8,12,16,20,24)
		case 4:  //脉搏
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 2 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,4,8,12,16,20,24)
		case 5:  //心率
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 5 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,4,8,12,16,20,24)
		case 6:  //呼吸
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 3 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,4,8,12,16,20,24)
		case 7:  //输入液量

		case 8:  //派出大便
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 13 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,0)
		case 9:  //派出尿量

		case 10: //派出其他

		case 11: //血压
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 4 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,0)
		case 12: //体重
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 8 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,0)
		case 13: //皮试
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 10 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,0)
		case 14: //其他
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 14 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,0)
		case 15: //事间
			sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 12 "
			msg = append(msg,data,patienttd)
			typetimes = append(typetimes,0)
		}

        for _,k := range typetimes{
        	sql1 := sql
			msg1 := msg
        	if len(typetimes) != 1{
				sql1 = sql1 + " and TypeTime = ?"
				msg1 = append(msg1,k)
			}

			maps,err := fit.MySqlEngine().QueryString(sql1,msg1...)
			fmt.Println("---------", maps)

			if err!= nil{
				return err,value
			}else{
				if len(maps) > 0{
						for _,dict := range  maps {
							if v, ok := dict["Value"]; ok {
								value = append(value,v)
							} else {
								value = append(value,"")
							}
						}
				}else{
					value = append(value,"")
				}
			}
		}
	}

	return nil,value
}

type TemperatrureChatHistory struct{
	//NurseId             int       `json:"nurse_id"`
	//NurseName           string    `json:"nurse_name"`
	PatientAge          string    `json:"patient_age"`
	PatientBed          string    `json:"patient_bed"`
	PatientId           int       `json:"patient_id"`
	PatientName         string    `json:"patient_name"`
	TestTime            string    `json:"test_time"`       //测量时间
	DateTime            string    `json:"date_time"`       //日期时间
	TimeFrame           string    `json:"time_frame"`      //时段
	ThmValue            string    `json:"thm_value"`
	ThmType             int       `json:"thm_type"`
	ThmScene            int       `json:"thm_scene"`
	PulseValue          string    `json:"pulse_value"`
	PulseBriefness      int       `json:"pulse_briefness"`
	BreatheValue        string    `json:"breathe_value"`
	BreatheScene        int       `json:"breathe_scene"`
	ShitValue           string    `json:"shit_value"`
	ShitScene           int       `json:"shit_scene"`
	PressureSys       string      `json:"pressure_sys"`
	PressureDia      string       `json:"pressure_dia"`
	PressureScene       int       `json:"pressure_scene"`
	HeartrateValue      string    `json:"heartrate_value"`
	Spo2hValue          string    `json:"spo2h_value"`
	GlucoseValue        string    `json:"glucose_value"`
	GlucoseType         int       `json:"glucose_type"`
	GlucoseScene        int       `json:"glucose_scene"`
	WeightValue         string    `json:"weight_value"`
	WeightScene         int       `json:"weight_scene"`
	HeightValue         string    `json:"height_value"`
	HeightScene         int       `json:"height_scene"`
	SkinValue           string    `json:"skin_value"`
	OtherValue          string    `json:"other_value"`
	IncidentScene       int       `json:"incident_scene"`
	IncidentTime        string    `json:"incident_time"`
}

func TransformTemperatrureCH(chat TemperatrureChat,history *TemperatrureChatHistory) (error,bool) {
     if chat.HeadType == ""{
     	return errors.New("没有类型"),false
	 }

	switch chat.HeadType {
	case Temperature_Type :
		if history.ThmValue != "" || history.ThmScene != 0{
			return nil,true
		}
		history.ThmValue = chat.Value
		history.ThmType = chat.Type
		history.ThmScene = chat.Other
	case Pulse_Type:
		history.PulseValue = chat.Value
		history.PulseBriefness = chat.Other
	case Breathe_Type:
		history.BreatheValue = chat.Value
		history.BreatheScene = chat.Other
	case Pressure_Type:
		if chat.Value != ""{
			value := strings.Split(chat.Value,"/")
			if len(value) == 2{
				history.PressureSys= value[0]
				history.PressureDia = value[1]
			}else{
				history.PressureSys= ""
				history.PressureDia = ""
			}
		}else{
			history.PressureSys= ""
			history.PressureDia = ""
		}
		history.PressureScene = chat.Other
	case Heartrate_Type:
		history.HeartrateValue = chat.Value
	/*case Spo2h_Type:
		history.Spo2hValue = chat.Value
	case Glucose_Type:
		history.GlucoseValue = chat.Value
		history.GlucoseType = chat.Type
		history.GlucoseScene = chat.Other*/
	case Weight_Type:
		history.WeightValue = chat.Value
		history.WeightScene = chat.Other
	case Height_Type:
		history.HeightValue = chat.Value
		history.HeightScene = chat.Other
	case Skin_Type:
		history.SkinValue = chat.Value
	case Incident_Type:
		if  history.IncidentScene != 0{
			return nil,true
		}
		history.IncidentScene = chat.Other
		history.IncidentTime = chat.TestTime.String()
	case Shit_Type:
		history.ShitValue = chat.Value
		history.ShitScene = chat.Other
	case Other_Type:
		history.OtherValue = chat.Value
	default:
		return errors.New("未知类型"),false
	}
	return nil,false
}

func BackIntervalTime(Interval int) string{
	switch Interval {
	case 4:
		return "04:00:00"
	case 8:
		return "05:00:00"
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

type PatientHistory struct{
	PatientId  int          `json:"patient_id"`    // 病人ID
	BedCoding  string       `json:"bed_coding"`    // 床号 BCQ1表
	Name       string       `json:"name"`          // 姓名
	Checked    int          `json:"checked"`       //是否选中  0没有选中 1选中
	Age        string       `json:"age"`           // 年龄
}

