package model

import (
	"fit"
	"github.com/go-xorm/xorm"
	"time"
)

const (
	Temperature_Type string = "1"      //体温
	Pulse_Type       string = "2"      //脉搏
	Breathe_Type     string = "3"      //呼吸
	Pressure_Type    string = "4"      //血压
	Heartrate_Type   string = "5"      //心率
	Spo2h_Type       string = "6"      //血氧
	Glucose_Type     string = "7"      //血糖
	Weight_Type      string = "8"      //体重
	Height_Type      string = "9"      //身高
	Skin_Type        string = "10"     //皮试
	Incident_Type    string = "12"     //事件
	Shit_Type        string = "13"     //大便
	Other_Type       string = "14"     //其他
	Intake_Type      string = "15"     //入量
	Output_Type      string = "16"     //出量
)

//护理单数据模型
type NurseChat struct {
	Id        int64        `json:"id" xorm:"pk autoincr"`
	HeadType  string       `json:"headtype" xorm:"notnull comment(头部id,对应头部类型)"`
	TestTime  fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	SubType   int          `json:"type" xorm:"notnull comment(类型,)"`
	Other     int          `json:"other" xorm:"notnull comment(其他可能选项,)"`
	OtherStr  string       `json:"otherstr" xorm:"notnull comment(其他可能选项string存储针对出入量)"`
	Describe  string       `json:"otherstr" xorm:"notnull comment(描述针对出入量)"`
	Value     string       `json:"value" xorm:"notnull comment(值)"`
	PatientId int64        `json:"patientid" xorm:"notnull comment(病人id)"`
	NurseId   int          `json:"nurseid" xorm:"notnull comment(护士id)"`
	NurseName string       `json:"nursename" xorm:"notnull comment(护士姓名)"`
}

//添加护理数据
func InsertNurseChat(session *xorm.Session,item *NurseChat) error{
	has,err := session.QueryString("SELECT id FROM NurseChat WHERE TestTime = ? and PatientId = ? and HeadType = ? and SubType = ?",item.TestTime.String(),item.PatientId,item.HeadType,item.SubType)
	if err !=nil{
		return err
	}

	if len(has)>0 {
		ids := has[0]
		if v, ok := ids["id"]; ok {
			_, err = session.Table("NurseChat").ID(v).Update(item);
			fit.Logger().LogError("ghhhhhhhh",item.HeadType,v,err,*item)
		} else {
			_, err = session.Insert(item);
		}
	}else{
		_, err = session.Insert(item);
	}

	return err
}

//输入数据分流护理单和体温单
func IputChat(session *xorm.Session,strData NurseChat) (int,error) {

	day,_ := time.ParseDuration("24h")

	if   strData.HeadType == Glucose_Type || strData.HeadType == Spo2h_Type{  //血糖和血氧存入护理单
		err := InsertNurseChat(session,&strData)
			return 59,err
	}else if strData.HeadType == Incident_Type || strData.HeadType == Shit_Type {  //事件和大便一天可以存很多次
		var test_time         = time.Time(strData.TestTime)
		var test_headtype     = strData.HeadType
		var test_type         = strData.SubType
		var test_other        = strData.Other
		var test_value        = strData.Value
		var test_patientid    = strData.PatientId
		var text_nurseid      = strData.NurseId
		var text_nursename    = strData.NurseName

		var text_typetime     int
		var text_datetime     time.Time

		text_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),0,0,0,0,time.Local)

		hour := test_time.Hour()
		if hour>=2 && hour<6{
			text_typetime = 4
		}else if hour>=6 && hour<10{
			text_typetime = 8
		}else if hour>=10 && hour<14{
			text_typetime = 12
		}else if hour>=14 && hour<18{
			text_typetime = 16
		}else if hour>=18 && hour<22{
			text_typetime = 20
		}else if hour>=22 && hour<24{
			text_typetime = 24
		}else if hour<2{
			text_typetime = 24
			text_datetime = text_datetime.Add(-day)
		}

		var item = &TemperatrureChat{}
		var sql string
		var msg []interface{}

		sql = "TestTime = ? and PatientId = ? and HeadType = ? and SubType = ?"
		msg = append(msg,test_time.Format("2006-01-02 15:04:05"),test_patientid,test_headtype,test_type)

		has,err := session.Table("TemperatrureChat").Where(sql,msg...).Get(item)

		if err !=nil{
			return 31,err
		}

		if has {
			old_time := time.Time(item.TestTime)
			if test_time.Format("2006-01-02 15:04:05") == old_time.Format("2006-01-02 15:04:05") {
				item.HeadType = test_headtype
				item.DateTime = fit.JsonTime(text_datetime)
				item.TestTime = fit.JsonTime(test_time)
				item.TypeTime = text_typetime
				item.SubType     = test_type
				item.Other    = test_other
				item.Value    = test_value
				item.PatientId = test_patientid
				item.NurseId =  text_nurseid
				item.NurseName = text_nursename
				fit.Logger().LogError("ghhhhhhhh",item.HeadType,test_value,item.Id)
				_,err = session.Table("TemperatrureChat").ID(item.Id).Update(item)
				return 32,err
			}else {
				item.HeadType = test_headtype
				item.DateTime = fit.JsonTime(text_datetime)
				item.TestTime = fit.JsonTime(test_time)
				item.TypeTime = text_typetime
				item.SubType     = test_type
				item.Other    = test_other
				item.Value    = test_value
				item.PatientId = test_patientid
				item.NurseId =  text_nurseid
				item.NurseName = text_nursename
				_, err = session.Insert(item);
				return 34,err
			}
		}else{
			item.HeadType = test_headtype
			item.DateTime = fit.JsonTime(text_datetime)
			item.TestTime = fit.JsonTime(test_time)
			item.TypeTime = text_typetime
			item.SubType     = test_type
			item.Other    = test_other
			item.Value    = test_value
			item.PatientId = test_patientid
			item.NurseId =  text_nurseid
			item.NurseName = text_nursename
			_, err = session.Insert(item);
			return 34,err
		}
	}else if ( strData.HeadType == Pressure_Type || strData.HeadType == Weight_Type || strData.HeadType == Skin_Type || strData.HeadType == Other_Type || strData.HeadType == Height_Type){//血压体重身高皮试其他 一天存一次
		var test_time         = time.Time(strData.TestTime)
		var test_headtype     = strData.HeadType
		var test_type         = strData.SubType
		var test_other        = strData.Other
		var test_value        = strData.Value
		var test_patientid    = strData.PatientId
		var text_nurseid      = strData.NurseId
		var text_nursename    = strData.NurseName

		var text_typetime     int
		var text_datetime     time.Time

		text_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),0,0,0,0,time.Local)
		hour := test_time.Hour()
		if hour>=2 && hour<6{
			text_typetime = 4
		}else if hour>=6 && hour<10{
			text_typetime = 8
		}else if hour>=10 && hour<14{
			text_typetime = 12
		}else if hour>=14 && hour<18{
			text_typetime = 16
		}else if hour>=18 && hour<22{
			text_typetime = 20
		}else if hour>=22 && hour<24{
			text_typetime = 24
		}else if hour<2{
			text_typetime = 24
			text_datetime = text_datetime.Add(-day)
		}

		var item = &TemperatrureChat{}
		var sql string
		var msg []interface{}

		sql = "DateTime = ? and PatientId = ? and HeadType = ? and SubType = ?"
		msg = append(msg,text_datetime.String(),test_patientid,test_headtype,test_type)

		has,err := session.Table("TemperatrureChat").Where(sql,msg...).Get(item)

		if err !=nil{
			return 31,err
		}

		if has {
			old_time := time.Time(item.TestTime)
			var item1 = &NurseChat{}
			if test_time.Format("2006-01-02 15:04:05") == old_time.Format("2006-01-02 15:04:05") {
				item.HeadType = test_headtype
				item.DateTime = fit.JsonTime(text_datetime)
				item.TestTime = fit.JsonTime(test_time)
				item.TypeTime = text_typetime
				item.SubType     = test_type
				item.Other    = test_other
				item.Value    = test_value
				item.PatientId = test_patientid
				item.NurseId =  text_nurseid
				item.NurseName = text_nursename
				fit.Logger().LogError("ghhhhhhhh",item.HeadType,test_value,item.Id)
				_,err = session.Table("TemperatrureChat").ID(item.Id).Update(item)
				return 32,err
			}else if test_time.Unix() < old_time.Unix(){
				item1.HeadType = item.HeadType
				item1.TestTime = item.TestTime
				item1.SubType = item.SubType
				item1.Other = item.Other
				item1.Value = item.Value
				item1.PatientId = item.PatientId
				item1.NurseId = item.NurseId
				item1.NurseName = item.NurseName

				err = InsertNurseChat(session,item1)
				if err !=nil{
					return 33,err
				}

				item.HeadType = test_headtype
				item.DateTime = fit.JsonTime(text_datetime)
				item.TestTime = fit.JsonTime(test_time)
				item.TypeTime = text_typetime
				item.SubType     = test_type
				item.Other    = test_other
				item.Value    = test_value
				item.PatientId = test_patientid
				item.NurseId =  text_nurseid
				item.NurseName = text_nursename
				_,err = session.Table("TemperatrureChat").Where("ID = ?",item.Id).Update(item)
				return 33,err
			}else{
				item1.HeadType = test_headtype
				item1.TestTime = fit.JsonTime(test_time)
				item1.SubType = test_type
				item1.Other = test_other
				item1.Value = test_value
				item1.PatientId = test_patientid
				item1.NurseId = text_nurseid
				item1.NurseName = text_nursename

				err = InsertNurseChat(session,item1)
				return 33,err
			}
		}else{
			item.HeadType = test_headtype
			item.DateTime = fit.JsonTime(text_datetime)
			item.TestTime = fit.JsonTime(test_time)
			item.TypeTime = text_typetime
			item.SubType     = test_type
			item.Other    = test_other
			item.Value    = test_value
			item.PatientId = test_patientid
			item.NurseId =  text_nurseid
			item.NurseName = text_nursename
			_, err = session.Insert(item);
			return 34,err
		}
	}else {  //温度呼吸按时段存
		var test_time         = time.Time(strData.TestTime)
		var test_headtype     = strData.HeadType
		var test_type         = strData.SubType
		var test_other        = strData.Other
		var test_value        = strData.Value
		var test_patientid    = strData.PatientId
		var text_nurseid      = strData.NurseId
		var text_nursename    = strData.NurseName

		var text_typetime     int
		var text_datetime     time.Time

		text_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),0,0,0,0,time.Local)
		var centre_datetime time.Time
		hour := test_time.Hour()
		if hour>=2 && hour<6{
			text_typetime = 4
			centre_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),4,0,0,0,time.Local)
		}else if hour>=6 && hour<10{
			text_typetime = 8
			centre_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),8,0,0,0,time.Local)
		}else if hour>=10 && hour<14{
			text_typetime = 12
			centre_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),12,0,0,0,time.Local)
		}else if hour>=14 && hour<18{
			text_typetime = 16
			centre_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),16,0,0,0,time.Local)
		}else if hour>=18 && hour<22{
			text_typetime = 20
			centre_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),20,0,0,0,time.Local)
		}else if hour>=22 && hour<24 {
			text_typetime = 24
			centre_datetime = time.Date(test_time.Year(),test_time.Month(),test_time.Day(),23,59,59,0,time.Local)
		}else if hour<2{
			text_typetime = 24
			text_datetime = text_datetime.Add(-day)
			centre_datetime = time.Date(text_datetime.Year(),text_datetime.Month(),text_datetime.Day(),23,59,59,0,time.Local)
		}

		var item = &TemperatrureChat{}
		var sql string
		var msg []interface{}


		sql = "DateTime = ? and TypeTime = ? and PatientId = ? and HeadType = ? and SubType = ?"
		msg = append(msg,text_datetime.String(),text_typetime,test_patientid,test_headtype,test_type)

		has,err := session.Table("TemperatrureChat").Where(sql,msg...).Get(item)

		if err !=nil{
			return 31,err
		}

		if has {
			old_time := time.Time(item.TestTime)
			var item1 = &NurseChat{}
			if test_time.Format("2006-01-02 15:04:05") == old_time.Format("2006-01-02 15:04:05") {
				item.HeadType = test_headtype
				item.DateTime = fit.JsonTime(text_datetime)
				item.TestTime = fit.JsonTime(test_time)
				item.TypeTime = text_typetime
				item.SubType     = test_type
				item.Other    = test_other
				item.Value    = test_value
				item.PatientId = test_patientid
				item.NurseId =  text_nurseid
				item.NurseName = text_nursename
				fit.Logger().LogError("ghhhhhhhh",item.HeadType,test_value,item.Id)
				_,err = session.Table("TemperatrureChat").ID(item.Id).Update(item)
				return 32,err
			}else if AbsInt(test_time.Unix() - centre_datetime.Unix()) <= AbsInt(old_time.Unix() - centre_datetime.Unix()){
				item1.HeadType = item.HeadType
				item1.TestTime = item.TestTime
				item1.SubType = item.SubType
				item1.Other = item.Other
				item1.Value = item.Value
				item1.PatientId = item.PatientId
				item1.NurseId = item.NurseId
				item1.NurseName = item.NurseName

				err = InsertNurseChat(session,item1)
				if err !=nil{
					return 33,err
				}

				item.HeadType = test_headtype
				item.DateTime = fit.JsonTime(text_datetime)
				item.TestTime = fit.JsonTime(test_time)
				item.TypeTime = text_typetime
				item.SubType     = test_type
				item.Other    = test_other
				item.Value    = test_value
				item.PatientId = test_patientid
				item.NurseId =  text_nurseid
				item.NurseName = text_nursename
				_,err = session.Table("TemperatrureChat").Where("ID = ?",item.Id).Update(item)
				return 33,err
			}else{
				item1.HeadType = test_headtype
				item1.TestTime = fit.JsonTime(test_time)
				item1.SubType = test_type
				item1.Other = test_other
				item1.Value = test_value
				item1.PatientId = test_patientid
				item1.NurseId = text_nurseid
				item1.NurseName = text_nursename

				err = InsertNurseChat(session,item1)
				return 33,err
			}
		}else{
			item.HeadType = test_headtype
			item.DateTime = fit.JsonTime(text_datetime)
			item.TestTime = fit.JsonTime(test_time)
			item.TypeTime = text_typetime
			item.SubType     = test_type
			item.Other    = test_other
			item.Value    = test_value
			item.PatientId = test_patientid
			item.NurseId =  text_nurseid
			item.NurseName = text_nursename
			_, err = session.Insert(item);
			return 34,err
		}
	}
}

//得到护理单数据
func OutNurseChat(sql string, msg ...interface{}) ([]NurseChat, error) {
	items := make([]NurseChat, 0)
	//fit.MySqlEngine().ShowSQL(true)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)

	return items, err
}

//得到病人是否发热
func GetWhetherFever(PatientId string,t float32)(bool,error){
	return fit.MySqlEngine().SQL("select (value > ?) as Fever from (SELECT HeadType,PatientId,TestTime,value from NurseChat UNION ALL SELECT HeadType,PatientId,TestTime,value from TemperatrureChat) alias WHERE HeadType = 1 and PatientId = ? order by testtime desc limit 1",t,PatientId).Exist()
}
