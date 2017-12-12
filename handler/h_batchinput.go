package handler

import (
	"fit"
	"nursing/model"
	"encoding/json"
	"time"
	"errors"
	"github.com/go-xorm/xorm"
	"strconv"
	"nursing/utils"
	/*"strings"*/
	"strings"
)

//体征批量录入路由
type PCBatvhinputController struct{
	PCController
}

func (c PCBatvhinputController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadViewSafely(w, r, "pc/v_batchInput.html", "pc/header_side.html", "pc/header_top.html")

		classid := userinfo.DepartmentID

		c.Data = make(fit.Data)
		c.Data["Userinfo"] = userinfo
		c.Data["Menuindex"] = "4-0"

		classresponse, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}

		measure_type := r.FormValue("measuretype")
		test_time := r.FormValue("testtime")
		interval := r.FormValue("timeblock")

		testtime := time.Now()
		selecttime := time.Now()

		if test_time == ""{
			testtime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 0, 0, 0, 0, testtime.Location())
			selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 8, 0, 0, 0, testtime.Location())
		}else{
			testtime, err_rp = time.ParseInLocation("2006-01-02",test_time,time.Local)
			/*if err_rp != nil{
				testtime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 0, 0, 0, 0, testtime.Location())
			}*/
			switch interval {
			case "4":
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 5, 59, 59, 0, testtime.Location())
			case "8":
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 9, 59, 59, 0, testtime.Location())
			case "12":
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 13, 59, 59, 0, testtime.Location())
			case "16":
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 17, 59, 59, 0, testtime.Location())
			case "20":
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 21, 59, 59, 0, testtime.Location())
			case "24":
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 23, 59, 59, 0, testtime.Location())
			default:
				selecttime = time.Date(testtime.Year(), testtime.Month(), testtime.Day(), 5, 59, 59, 0, testtime.Location())
			}
		}

		var response []model.PCBedDup

		for _,k := range classresponse{
			hospitaldate := strings.Split(k.HospitalDate," ")
			admission_time, err1 := time.ParseInLocation("2006-01-02",hospitaldate[0],time.Local)

			if err1 != nil {
				fit.Logger().LogError("gk", err1)
				continue
			}

			if !admission_time.After(testtime){
				response = append(response,k)
			}
		}


		if measure_type == ""{
			c.Data["Patients"] = response
		}else if measure_type == "1"{
			c.Data["Patients"] = response
		}else if measure_type == "2"{
			var PCBitems []model.PCBedDup

			for _,v := range response{

				whether := false

				hospitaldate := strings.Split(v.HospitalDate," ")
				hospitaltime, err := time.ParseInLocation("2006-01-02",hospitaldate[0],time.Local)
				//testtime, err1 := time.ParseInLocation("2006-01-02",test_time,time.Local)
				if err != nil {
					//|| err1 != nil
					fit.Logger().LogError("待测体温1",v.VAA01,err,test_time)
					return
				}


				if !whether {  //获取发热体温
				whetherfever,err := model.GetWhetherFever(selecttime.Format("2006-01-02 15:04:05"),strconv.FormatInt(v.VAA01,10),38.5);
                if err != nil{
					return
				}

				if whetherfever{
                    bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
					if err != nil {
						return
					}
					if !bo && !whether{
						whether = true
						PCBitems = append(PCBitems,v)
					}
				}else{
					whetherfever,err := model.GetWhetherFever(selecttime.Format("2006-01-02 15:04:05"),strconv.FormatInt(v.VAA01,10),37.5)
					if err != nil{
						return
					}
					fit.Logger().LogError("待测体温2",v.VAA01,v.VAA05,whetherfever,selecttime)
					if whetherfever{
						if interval == "8" || interval == "12" || interval == "16" || interval == "20"{
							bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
							if err != nil {
								return
							}
							if !bo && !whether{
								whether = true
								PCBitems = append(PCBitems,v)
							}
						}
					}
				}}

				if !whether{  //获取最近三天是否发热
				    day,_ := time.ParseDuration("24h")
					recordbefore := testtime.Add(-day*3)
					//recordlater  := testtime.Add(day)
					fit.Logger().LogError("三天待测体温",recordbefore,selecttime)
					whetherfever,err := model.GetWhetherFeverThree(recordbefore.Format("2006-01-02 15:04:05"),selecttime.Format("2006-01-02 15:04:05"),strconv.FormatInt(v.VAA01,10),38.5);
					if err != nil{
						return
					}

					if whetherfever{
						bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
						if err != nil {
							return
						}
						if !bo && !whether{
							whether = true
							PCBitems = append(PCBitems,v)
						}
					}else{
						whetherfever,err := model.GetWhetherFeverThree(recordbefore.Format("2006-01-02 15:04:05"),selecttime.Format("2006-01-02 15:04:05"),strconv.FormatInt(v.VAA01,10),37.5)
						if err != nil{
							return
						}
						fit.Logger().LogError("待测体温2",v.VAA01,v.VAA05,whetherfever)
						if whetherfever{
							if interval == "8" || interval == "12" || interval == "16" || interval == "20"{
								bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
								if err != nil {
									return
								}
								if !bo && !whether{
									whether = true
									PCBitems = append(PCBitems,v)
								}
							}
						}
					}
				}

				if !whether { //获取新病人
					if GetWhetherNew(hospitaltime, testtime) {
						if interval == "8" || interval == "16" || interval == "20" {
							bo, err := model.GetTemperatureWhetherMeasured(testtime.String(), v.VAA01, interval)
							if err != nil {
								return
							}
							if !bo && !whether {
								whether = true
								PCBitems = append(PCBitems, v)
							}
						}
					}
				}

				if !whether { //获取手术
				day,_ := time.ParseDuration("24h")
				recordbefore := testtime.Add(-day*3)
				recordlater  := testtime.Add(day*2)
				records,err := model.FetchOperationRecordsDuringHospitalization(v.VAA01,recordbefore.Format("2006-01-02 15:04:05"),recordlater.Format("2006-01-02 15:04:05"))
				if err != nil {
					return
				}
				if len(records)>0{
					record := records[0]
					if record.VAT04 == 2{
						if  interval == "20"{
							bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
							if err != nil {
								return
							}
							if !bo && !whether{
								whether = true
								PCBitems = append(PCBitems,v)
							}
						}
					}
					if record.VAT04 == 3{
						if  interval == "8"{
							bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
							if err != nil {
								return
							}
							if !bo && !whether{
								whether = true
								PCBitems = append(PCBitems,v)
							}
						}
					}
					if record.VAT04 == 4{
						if  interval == "4" || interval == "8" || interval == "16" || interval == "20"{
							bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
							if err != nil {
								return
							}
							if !bo && !whether{
								whether = true
								PCBitems = append(PCBitems,v)
							}
						}
					}
				}}

				if !whether { //正常
				if v.VAA10 < 15{  //儿童每天四次
					if  interval == "4" || interval == "8" || interval == "16" || interval == "20"{
						bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,interval)
						if err != nil {
							return
						}
						if !bo && !whether{
							whether = true
							PCBitems = append(PCBitems,v)
						}
					}
				}else{          //正常人每天一次
					bo,err := model.GetTemperatureWhetherMeasured(testtime.String(),v.VAA01,"")
					if err != nil {
						return
					}
					if !bo && !whether{
						whether = true
						PCBitems = append(PCBitems,v)
					}
				}
				}
			}
			c.Data["Patients"] = PCBitems
		}else if measure_type == "3"{
			var PCBitems []model.PCBedDup
			for _,v := range response{
				hospitaldate := strings.Split(v.HospitalDate," ")
				hospitaltime, err := time.ParseInLocation("2006-01-02",hospitaldate[0],time.Local)
				//testtime, err1 := time.ParseInLocation("2006-01-02",test_time,time.Local)
				if err != nil {
					//|| err1 != nil
					fit.Logger().LogError("待测体温1",err)
					return
				}

				if GetWeekOntime(hospitaltime,testtime) {
					spl := "HeadType = 8 and DateTime = ? and PatientId = ?"
					var msg []interface{}
					msg = append(msg,testtime.String(),v.VAA01)
					bo, err := model.WhetherTemperature(spl,msg...)
					fit.Logger().LogError("待测体温1",v.VAA01,hospitaltime,GetWeekOntime(hospitaltime,testtime),bo)
					if err!= nil{
						return
					}
					if !bo {
						PCBitems = append(PCBitems,v)
					}
				}
			}
			c.Data["Patients"] = PCBitems
		}

		fit.Logger().LogError("gk dd", len(response))
	}
}

//获取从一时间开始计算后面每周一的对比
func GetWeekOntime(t1 time.Time,t2 time.Time) bool {
	if  t1.After(t2){
		return false
	}
	d:=t2.Sub(t1).Hours()
	fit.Logger().LogError("GetWeekOntime",t1,t2,d)
	if int64(d) % (24*7) == 0{
		return true
	}else{
		return false
	}
}

//是否新入院，刚入院三天
func GetWhetherNew(t1 time.Time,t2 time.Time) bool {
	if  t1.After(t2){
		return false
	}
	d:=t2.Sub(t1).Hours()
	if d < 24*3{
		return true
	}else{
		return false
	}
}

//体征批量数据录入
func (c PCBatvhinputController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	inputs := r.FormValue("value")
    if inputs == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		return
	}

	fit.Logger().LogError("PCBatvhinputController",inputs)

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

	err := session.Begin()
	if err!=nil{
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "事务开始失败"
		return
	}

	var maps []map[string]string
	err_js := json.Unmarshal([]byte(inputs), &maps)
	if err_js !=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "格式错误1"
		c.JsonData.Datas = err_js
		return
	}else{
		for _,str := range maps{
			mag,err_as := BatchAnalysis(session,str)
			if err_as !=nil {
				session.Rollback()
				c.JsonData.Result = mag
				c.JsonData.ErrorMsg = "解析错误1"
				c.JsonData.Datas = err_as.Error()
				return
			}
		}
	}

	err_com := session.Commit()
	if err_com != nil {
		c.JsonData.Result = 4
		c.JsonData.ErrorMsg = "数据库插入失败"
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "录入成功"
	}
}

func BatchAnalysis(session *xorm.Session,strData map[string]string) (int,error){

	var nurse_id        int
	var nurse_name      string
	var patient_id      int64
	var test_time       fit.JsonTime

	if v, ok := strData["nurse_id"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 1,err
		}else{
			nurse_id = k
		}
	} else {
		return 2,errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		nurse_name = v
	} else {
		return 3,errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		k,err  := utils.Int64Value(v)
		if err != nil{
			return 4,err
		}else{
			patient_id = k
		}
	} else {
		return 5,errors.New("没有patient_id")
	}

	if v, ok := strData["test_time"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return 6,err
		} else {
			test_time = fit.JsonTime(texttime)
		}
	} else {
		return 7,errors.New("没有test_time")
	}


	var thm_value       string    //体温
	var thm_type        int
	var thm_scene       int

	if v, ok := strData["thm_value"]; ok {
		thm_value = v
	} else {
		return 8,errors.New("thm_value")
	}

	if v, ok := strData["thm_type"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 9,err
		}else{
			thm_type = k
		}
	} else {
		return 10,errors.New("thm_type")
	}

	if v, ok := strData["thm_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 11,err
		}else{
			thm_scene = k
		}
	} else {
		return 12,errors.New("thm_scene")
	}

	if thm_value != "" || thm_scene != 0 {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Temperature_Type
		item.Value = thm_value
		item.SubType = thm_type
		item.Other = thm_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}


	var pulse_value     string    //脉搏
	var pulse_briefness int

	if v, ok := strData["pulse_value"]; ok {
		pulse_value = v
	} else {
		return 13,errors.New("pulse_value")
	}

	if v, ok := strData["pulse_briefness"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 14,err
		}else{
			pulse_briefness = k
		}
	} else {
		return 15,errors.New("pulse_briefness")
	}

	if pulse_value != "" || pulse_briefness == 1{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Pulse_Type
		item.Value = pulse_value
		item.SubType = 0
		item.Other = pulse_briefness

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}


	var breathe_value   string    //呼吸
	var breathe_scene   int

	if v, ok := strData["breathe_value"]; ok {
		breathe_value = v
	} else {
		return 16,errors.New("breathe_value")
	}

	if v, ok := strData["breathe_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 17,err
		}else{
			breathe_scene = k
		}
	} else {
		return 18,errors.New("breathe_scene")
	}

	if breathe_value != "" || breathe_scene !=0{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Breathe_Type
		item.Value = breathe_value
		item.SubType = 0
		item.Other = breathe_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var shit_value   string     //大便
	var shit_scene   int

	if v, ok := strData["shit_value"]; ok {
		shit_value = v
	} else {
		return 19,errors.New("shit_value")
	}

	if v, ok := strData["shit_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 20,err
		}else{
			shit_scene = k
		}
	} else {
		return 21,errors.New("shit_scene")
	}

	if shit_value != "" || shit_scene!=0 {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Shit_Type
		item.Value = shit_value
		item.SubType = 0
		item.Other = shit_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}


	var pressure_sys      string    //高压
	var pressure_dia      string    //低压
	var pressure_scene    int

	if v, ok := strData["pressure_sys"]; ok {
		pressure_sys = v
	} else {
		return 22,errors.New("pressure_value")
	}

	if v, ok := strData["pressure_dia"]; ok {
		pressure_dia = v
	} else {
		return 22,errors.New("pressure_dia")
	}

	if v, ok := strData["pressure_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 23,err
		}else{
			pressure_scene = k
		}
	} else {
		return 24,errors.New("pressure_scene")
	}

	if (pressure_sys != "" && pressure_dia != "") || pressure_scene != 0{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Pressure_Type
		if (pressure_sys != "" && pressure_dia != ""){
			item.Value = pressure_sys + "/" + pressure_dia
		}
		item.SubType = 0
		item.Other = pressure_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var heartrate_value string   //心率

	if v, ok := strData["heartrate_value"]; ok {
		heartrate_value = v
	} else {
		return 25,errors.New("heartrate_value")
	}

	if heartrate_value != "" {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Heartrate_Type
		item.Value = heartrate_value
		item.SubType = 0
		item.Other = 0

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var weight_value    string    //体重
	var weight_scene    int

	if v, ok := strData["weight_value"]; ok {
		weight_value = v
	} else {
		return 26,errors.New("weight_value")
	}

	if v, ok := strData["weight_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 27,err
		}else{
			weight_scene = k
		}
	} else {
		return 28,errors.New("weight_scene")
	}

	if weight_value != "" || weight_scene != 0{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Weight_Type
		item.Value = weight_value
		item.SubType = 0
		item.Other = weight_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var height_value    string    //身高
	var height_scene    int

	if v, ok := strData["height_value"]; ok {
		height_value = v
	} else {
		return 29,errors.New("height_value")
	}

	if v, ok := strData["height_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 30,err
		}else{
			height_scene = k
		}
	} else {
		return 31,errors.New("height_scene")
	}

	if height_value != "" || height_scene != 0{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Height_Type
		item.Value = height_value
		item.SubType = 0
		item.Other = height_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var skin_value     string    //皮试
	if v, ok := strData["skin_value"]; ok {
		skin_value = v
	} else {
		return 32,errors.New("skin_value")
	}

	if skin_value != "" {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Skin_Type
		item.Value = skin_value
		item.SubType = 0
		item.Other = 0

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var other_value   string    //其他
	if v, ok := strData["other_value"]; ok {
		other_value = v
	} else {
		return 33,errors.New("other_value")
	}

	if other_value != "" {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Other_Type
		item.Value = other_value
		item.SubType = 0
		item.Other = 0

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	var incident_scene int   //事件
	var incident_time fit.JsonTime  //事件时间

	if v, ok := strData["incident_scene"]; ok {
		k,err  := strconv.Atoi(v)
		if err != nil{
			return 34,err
		}else{
			incident_scene = k
		}
	} else {
		return 35,errors.New("incident_scene")
	}

	if v, ok := strData["incident_time"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return 36,errors.New("没有incident_time")
		} else {
			incident_time = fit.JsonTime(texttime)
		}
	} else {
		return 1,errors.New("没有incident_time")
	}

	if incident_scene != 0 {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(incident_time)

		item.HeadType = model.Incident_Type
		item.Value = ""
		item.SubType = 0
		item.Other = incident_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			return msg,err
		}
	}

	return 0,nil
}

