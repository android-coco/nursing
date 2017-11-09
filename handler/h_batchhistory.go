package handler

import (
	"nursing/model"
	"fit"
	"time"
	"encoding/json"
	/*"strconv"*/
)

/*//体征数据
type signvalue struct{
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	model.IntakeOutput  `xorm:"extends"`  // 出入量

}*/

type PCBatvhHistoryController struct{
	PCController
}

func (c PCBatvhHistoryController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadViewSafely(w, r, "pc/v_batchinput2.html", "pc/header_side.html", "pc/header_top.html")

		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}
		Data["Patients"] = response

		starttime :=  r.FormValue("starttime")
		endtime   :=  r.FormValue("endtime")

		if starttime == ""|| endtime == ""{
		t := time.Now()
		st := time.Date(t.Year(), t.Month(), 3, 0, 0, 0, 0, t.Location())
		et := time.Date(t.Year(), t.Month(), 3, 23, 59, 59, 0, t.Location())
		//t.Day()
		starttime = st.Format("2006-01-02 15:04:05")
		endtime   = et.Format("2006-01-02 15:04:05")
		}

		var signhistorys []SignHistory

        for _,ii := range response{

        	signhistory_map := make(map[string]*SignHistory )

		var sql string
		var msg []interface{}

		sql = sql + "testtime >= ? and testtime <= ? and PatientId = ?"
		msg = append(msg,starttime, endtime,ii.VAA01)

			item, err := model.OutTemperature(sql, msg...)
			if err != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误"
				c.JsonData.Datas = err
				return
			} else {
				for _,i := range item{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Thm = i
					} else {
						var  signhistory SignHistory
						signhistory.Thm = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}
			fit.Logger().LogError("hhhhhhh",len(item))


			item1, err1 := model.OutPulse(sql, msg...)
			if err1 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误1"
				c.JsonData.Datas = err1
				return
			} else {
				for _,i := range item1{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Pulse = i
					} else {
						var  signhistory SignHistory
						signhistory.Pulse = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item2, err2 := model.OutBreathe(sql, msg...)
			if err2 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误2"
				c.JsonData.Datas = err2
				return
			} else {
				for _,i := range item2{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Breathe = i
					} else {
						var  signhistory SignHistory
						signhistory.Breathe = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item3, err3 := model.OutPressure(sql, msg...)
			if err3 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误3"
				c.JsonData.Datas = err3
				return
			} else {
				for _,i := range item3{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Pressure = i
					} else {
						var  signhistory SignHistory
						signhistory.Pressure = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}



			item4, err4 := model.OutHeartrate(sql, msg...)
			if err4 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误4"
				c.JsonData.Datas = err4
				return
			} else {
				for _,i := range item4{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Heartrate = i
					} else {
						var  signhistory SignHistory
						signhistory.Heartrate = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}



			item5, err5 := model.OutSpo2h(sql, msg...)
			if err5 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误5"
				c.JsonData.Datas = err5
				return
			} else {
				for _,i := range item5{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Spo2h = i
					} else {
						var  signhistory SignHistory
						signhistory.Spo2h = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item6, err6 := model.OutGlucose(sql, msg...)
			if err6 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误6"
				c.JsonData.Datas = err6
				return
			} else {
				for _,i := range item6{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Glucose = i
					} else {
						var  signhistory SignHistory
						signhistory.Glucose = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item7, err7 := model.OutWeight(sql, msg...)
			if err7 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误7"
				c.JsonData.Datas = err7
				return
			} else {
				for _,i := range item7{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Weight = i
					} else {
						var  signhistory SignHistory
						signhistory.Weight = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item8, err8 := model.OutHeight(sql, msg...)
			if err8 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误8"
				c.JsonData.Datas = err8
				return
			} else {
				for _,i := range item8{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Height = i
					} else {
						var  signhistory SignHistory
						signhistory.Height = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item9, err9 := model.OutSkin(sql, msg...)
			if err9 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误9"
				c.JsonData.Datas = err9
				return
			} else {
				for _,i := range item9{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Skin = i
					} else {
						var  signhistory SignHistory
						signhistory.Skin = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}

			item11, err11 := model.OutIncident(sql, msg...)
			if err11 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误11"
				c.JsonData.Datas = err11
				return
			} else {
				for _,i := range item11{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Incident = i
					} else {
						var  signhistory SignHistory
						signhistory.Incident = i
						signhistory.Texttime = i.Testtime.String()
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}

			item12, err_iov :=  model.QueryIntakeOrOutputVolumeWithPatient(ii.VAA01, starttime, endtime)
			if err_iov != nil {
				// sql出错
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误11"
				c.JsonData.Datas = err11
				return
			} else {
				for _,i := range item12{
					if v, ok := signhistory_map[i.Testtime]; ok {
						v.IntakeOutput = append(v.IntakeOutput,i)

					} else {
						var  signhistory SignHistory
						signhistory.IntakeOutput = append(signhistory.IntakeOutput,i)
						signhistory.Texttime = i.Testtime
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime] = &signhistory
					}
				}
			}

		  for _,j :=range  signhistory_map{
			  signhistorys = append(signhistorys,*j)
			  fit.Logger().LogError("gggggg",j.IntakeOutput)
		  }
		}

		Data["signhistorys"]=signhistorys
		Data["Menuindex"] = "4-0"
		Data["Userinfo"] = userinfo
		c.Data = Data
	}
}

func (c PCBatvhHistoryController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	patient_id := r.FormValue("patient_id")   //VAA01病人ID
	starttime := r.FormValue("starttime")     //开始时间
	endtime := r.FormValue("endtime")         //结束时间

	if patient_id == ""||starttime == ""||endtime == ""{
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "参数不完整"
			c.JsonData.Datas = []interface{}{}
			return
	}

	patientids := make([]string,0)
	err := json.Unmarshal([]byte(patient_id), &patientids)
	if err !=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "解析错误"
		c.JsonData.Datas = []interface{}{}
		return
	}else{
		var signhistorys []model.SignHistory1

		for _,ii := range patientids{

			signhistory_map := make(map[string]*model.SignHistory1 )

			var sql string
			var msg []interface{}

			sql = sql + "testtime >= ? and testtime <= ? and PatientId = ?"
			msg = append(msg,starttime, endtime,ii)

			patient, err_patient := model.FetchPatientInfoForSignHistory(ii)
			if err_patient != nil{
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误"
				c.JsonData.Datas = err_patient
				return
			}

			item, err := model.OutTemperature(sql, msg...)
			if err != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误0"
				c.JsonData.Datas = err
				return
			} else {
				for _,i := range item{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.ThmValue = i.Value
						v.ThmType  = i.Ttemptype
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.ThmValue= i.Value
						signhistory.ThmType = i.Ttemptype
						signhistory.PatientInfoDup = patient
						signhistory_map[i.Testtime.String()] = &signhistory
					}
			}
			}
			fit.Logger().LogError("hhhhhhh",len(item))


			item1, err1 := model.OutPulse(sql, msg...)
			if err1 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误1"
				c.JsonData.Datas = err1
				return
			} else {
				for _,i := range item1{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.PulseValue = i.Value
						v.PulseBriefness = i.Whetherbriefness
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.PulseValue = i.Value
						signhistory.PulseBriefness = i.Whetherbriefness
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item2, err2 := model.OutBreathe(sql, msg...)
			if err2 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误2"
				c.JsonData.Datas = err2
				return
			} else {
				for _,i := range item2{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.BreatheValue = i.Value
						v.BreatheScene = i.Recordscene
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.BreatheValue = i.Value
						signhistory.BreatheScene = i.Recordscene
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item3, err3 := model.OutPressure(sql, msg...)
			if err3 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误3"
				c.JsonData.Datas = err3
				return
			} else {
				for _,i := range item3{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.PressureDia = i.Diavalue
						v.PressureSys = i.Sysvalue
						v.PressureScene = i.Recordscene
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.PressureDia = i.Diavalue
						signhistory.PressureSys = i.Sysvalue
						signhistory.PressureScene = i.Recordscene
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}



			item4, err4 := model.OutHeartrate(sql, msg...)
			if err4 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误4"
				c.JsonData.Datas = err4
				return
			} else {
				for _,i := range item4{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.HeartrateValue = i.Value
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.HeartrateValue = i.Value
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}



			item5, err5 := model.OutSpo2h(sql, msg...)
			if err5 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误5"
				c.JsonData.Datas = err5
				return
			} else {
				for _,i := range item5{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.Spo2hValue = i.Value
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.Spo2hValue = i.Value
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item6, err6 := model.OutGlucose(sql, msg...)
			if err6 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误6"
				c.JsonData.Datas = err6
				return
			} else {
				for _,i := range item6{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.GlucoseValue = i.Value
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.GlucoseValue = i.Value
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item7, err7 := model.OutWeight(sql, msg...)
			if err7 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误7"
				c.JsonData.Datas = err7
				return
			} else {
				for _,i := range item7{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.WeightValue = i.Value
						v.WeightScene = i.Recordscene
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.WeightValue = i.Value
						signhistory.WeightScene = i.Recordscene
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item8, err8 := model.OutHeight(sql, msg...)
			if err8 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误8"
				c.JsonData.Datas = err8
				return
			} else {
				for _,i := range item8{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.HeightValue = i.Value
						v.HeightScene = i.Recordscene
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.HeightValue = i.Value
						signhistory.HeightScene = i.Recordscene
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}


			item9, err9 := model.OutSkin(sql, msg...)
			if err9 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误9"
				c.JsonData.Datas = err9
				return
			} else {
				for _,i := range item9{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.SkinValue = i.Value
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.SkinValue = i.Value
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}

			item11, err11 := model.OutIncident(sql, msg...)
			if err11 != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误11"
				c.JsonData.Datas = err11
				return
			} else {
				for _,i := range item11{
					if v, ok := signhistory_map[i.Testtime.String()]; ok {
						v.IncidentScene = i.Recordscene
					} else {
						var  signhistory model.SignHistory1
						signhistory.TextTime= i.Testtime.String()
						signhistory.PatientInfoDup = patient
						signhistory.IncidentScene = i.Recordscene
						signhistory_map[i.Testtime.String()] = &signhistory
					}
				}
			}

			/*vv,_:=strconv.Atoi(ii)
			item12, err_iov :=  model.QueryIntakeOrOutputVolumeWithPatient(vv, starttime, endtime)
			if err_iov != nil {
				// sql出错
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误11"
				c.JsonData.Datas = err11
				return
			} else {
				for _,i := range item12{
					if v, ok := signhistory_map[i.Testtime]; ok {
						v.IntakeOutput = append(v.IntakeOutput,i)

					} else {
						var  signhistory SignHistory
						signhistory.IntakeOutput = append(signhistory.IntakeOutput,i)
						signhistory.Texttime = i.Testtime
						signhistory.PCBedDup = ii
						signhistory_map[i.Testtime] = &signhistory
					}
				}

		    }*/
			for _,j :=range  signhistory_map{
				signhistorys = append(signhistorys,*j)
			}
		}

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "获取成功"
		c.JsonData.Datas = signhistorys
	}
}




type SignHistory struct{
	PCBedDup          interface{}        //个人信息
	Texttime          string             //测试时间
	Thm               model.Temperature  //体温数据
	Pulse             model.Pulse        //脉搏
	Breathe           model.Breathe      //呼吸
	Pressure          model.Pressure     //血压
	Heartrate         model.Heartrate    //心率
	Spo2h             model.Spo2h        //血氧
	Glucose           model.Glucose      //血糖
	Weight            model.Weight       //体重
	Height            model.Height       //身高
	Skin              model.Skin         //皮试
	Incident          model.Incident     //事件
	IntakeOutput      []model.IntakeOutput //出入量

}




