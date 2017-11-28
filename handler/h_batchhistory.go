package handler

import (
	"nursing/model"
	"fit"
	"time"
	"strconv"
	"strings"
	"nursing/utils"
	"sort"
)

//体征历史路由
type PCBatvhHistoryController struct{
	PCController
}

//体征历史获取
func (c PCBatvhHistoryController) TZHistory(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadViewSafely(w, r, "pc/v_batchinput2.html", "pc/header_side.html", "pc/header_top.html")
		c.Data = make(fit.Data)

		c.Data["Userinfo"] = userinfo
		c.Data["Menuindex"] = "4-0"

		classid := userinfo.DepartmentID

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}

		PatientHistorys := make([]model.PatientHistory,0)   //病人信息

		starttime :=  r.FormValue("starttime")
		endtime   :=  r.FormValue("endtime")
		startframe :=  r.FormValue("startframe")
		endframe   :=  r.FormValue("endframe")
		patient_id := r.FormValue("patient_id")   //VAA01病人ID

		if starttime == ""|| endtime == ""{
		    t := time.Now()
		    st := time.Date(t.Year(), t.Month(), t.Day(), 6, 0, 0, 0, t.Location())
		    et := time.Date(t.Year(), t.Month(), t.Day(), 10, 0, 0, 0, t.Location())

		    starttime = st.Format("2006-01-02 15:04:05")
		    endtime   = et.Format("2006-01-02 15:04:05")
		    c.Data["key"] = 1
			c.Data["startframe"] = "8"
			c.Data["endframe"] = "8"

			for _,v := range  response{
				var ph model.PatientHistory
				ph.PatientId = v.VAA01
				ph.BedCoding = v.BCQ04
				ph.Name = v.VAA05
				ph.Age = strconv.Itoa(v.VAA10)
				ph.Checked = 1
				PatientHistorys = append(PatientHistorys,ph)
			}
		}else{
			c.Data["key"] = 2
			c.Data["starttime"] = starttime
			c.Data["endtime"] = endtime
			c.Data["startframe"] = startframe
			c.Data["endframe"] = endframe
			c.Data["patientids"] = patient_id


			st , _:= time.ParseInLocation("2006-01-02 15:04:05", starttime+" 00:00:00" ,time.Local)
			et , _:= time.ParseInLocation("2006-01-02 15:04:05", endtime+" 00:00:00"  ,time.Local)

			day,_ := time.ParseDuration("24h")

			switch startframe {
			case "4":
				st = time.Date(st.Year(), st.Month(), st.Day(), 2, 0, 0, 0, st.Location())
			case "8":
				st = time.Date(st.Year(), st.Month(), st.Day(), 6, 0, 0, 0, st.Location())
			case "12":
				st = time.Date(st.Year(), st.Month(), st.Day(), 10, 0, 0, 0, st.Location())
			case "16":
				st = time.Date(st.Year(), st.Month(), st.Day(), 14, 0, 0, 0, st.Location())
			case "20":
				st = time.Date(st.Year(), st.Month(), st.Day(), 18, 0, 0, 0, st.Location())
			case "24":
				st = time.Date(st.Year(), st.Month(), st.Day(), 22, 0, 0, 0, st.Location())
			}

			switch endframe {
			case "4":
				et = time.Date(et.Year(), et.Month(), et.Day(), 5, 59, 59, 0, st.Location())
			case "8":
				et = time.Date(et.Year(), et.Month(), et.Day(), 9, 59, 59, 0, st.Location())
			case "12":
				et = time.Date(et.Year(), et.Month(), et.Day(), 13, 59, 59, 0, st.Location())
			case "16":
				et = time.Date(et.Year(), et.Month(), et.Day(), 17, 59, 59, 0, st.Location())
			case "20":
				et = time.Date(et.Year(), et.Month(), et.Day(), 21, 59, 59, 0, st.Location())
			case "24":
				et = et.Add(day)
				et = time.Date(et.Year(), et.Month(), et.Day(), 1, 59, 59, 0, st.Location())
			}

			starttime = st.Format("2006-01-02 15:04:05")
			endtime   = et.Format("2006-01-02 15:04:05")
			fit.Logger().LogError("ghffdref",starttime,endtime)

			patientids := strings.Split(patient_id,",")
			map_patientids := make(map[string]string)
			for _,v := range patientids{
				map_patientids[v]=v
			}
			fit.Logger().LogError("ghffdref", patientids)
			for _,v := range  response{
				var ph model.PatientHistory
				ph.PatientId = v.VAA01
				ph.BedCoding = v.BCQ04
				ph.Name = v.VAA05
				ph.Age = strconv.Itoa(v.VAA10)
				if _, ok := map_patientids[utils.FormatInt64(v.VAA01)];ok{
					ph.Checked = 1
				}else{
					ph.Checked = 0
				}
				PatientHistorys = append(PatientHistorys,ph)
			}
		}

		var signhistorys []model.TemperatrureChatHistory     //体温表数据

        //if patient_id == ""{
			for _,ii := range PatientHistorys{

				if ii.Checked == 0 {
					break
				}

				signhistory_map := make(map[string]*model.TemperatrureChatHistory )

				var sql string
				var msg []interface{}

				sql = sql + "testtime >= ? and testtime <= ? and PatientId = ?"
				msg = append(msg,starttime, endtime,ii.PatientId)

				item, err := model.OutTemperatureHistory(sql, msg...)
				if err != nil {
					c.JsonData.Result = 1
					c.JsonData.ErrorMsg = "查询错误"
					fit.Logger().LogError("hhhhhhh",err,ii)
					return
				} else {
					for _,i := range item{
						if v, ok := signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime)]; ok {
							_,boo := model.TransformTemperatrureCH(i,v)
							if boo {
                                for j := 0;j<10;j++ {
									if v, ok := signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime) + strconv.Itoa(j)]; ok {
										_,boo := model.TransformTemperatrureCH(i,v)
										if !boo {
											break
										}
									} else {
										var chat model.TemperatrureChatHistory
										chat.PatientId = ii.PatientId
										chat.PatientBed = ii.BedCoding
										chat.PatientAge = ii.Age
										chat.PatientName = ii.Name
										chat.TestTime =  time.Time(i.DateTime).Format("2006-01-02")+" "+model.BackIntervalTime(i.TypeTime)
										chat.DateTime = time.Time(i.DateTime).Format("2006-01-02")
										chat.TimeFrame = strconv.Itoa(i.TypeTime)
										er,_ := model.TransformTemperatrureCH(i,&chat)
										if er != nil{
											continue
										}
										signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime) + strconv.Itoa(j)] = &chat
										break
									}
								}
							}
						} else {
							var chat model.TemperatrureChatHistory
							chat.PatientId = ii.PatientId
							chat.PatientBed = ii.BedCoding
							chat.PatientAge = ii.Age
							chat.PatientName = ii.Name
							chat.TestTime =  time.Time(i.DateTime).Format("2006-01-02")+" "+model.BackIntervalTime(i.TypeTime)
							chat.DateTime = time.Time(i.DateTime).Format("2006-01-02")
							chat.TimeFrame = strconv.Itoa(i.TypeTime)
							er,_ := model.TransformTemperatrureCH(i,&chat)
							if er != nil{
								continue
							}
							signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime)] = &chat
						}
					}
				}
				fit.Logger().LogError("hhhhhhh",len(item))

				for _,j :=range  signhistory_map{
					signhistorys = append(signhistorys,*j)
				}
			}
		//}else{
				/*for _,ii := range PatientHistorys{

					signhistory_map := make(map[string]*model.TemperatrureChatHistory )

					var sql string
					var msg []interface{}

					sql = sql + "testtime >= ? and testtime <= ? and PatientId = ?"
					msg = append(msg,starttime, endtime,ii.PatientId)

					if ii.Checked == 0 {
						break
					}

					fit.Logger().LogError("bbbbbbbb",ii.PatientId)

					item, err := model.OutTemperatureHistory(sql, msg...)
					if err != nil {
						c.JsonData.Result = 1
						c.JsonData.ErrorMsg = "查询错误0"
						c.JsonData.Datas = err
						return
					} else {
						for _,i := range item{
							if v, ok := signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime)]; ok {
								_,boo := model.TransformTemperatrureCH(i,v)
								if boo {
									for j := 0;j<10;j++ {
										if v, ok := signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime) + strconv.Itoa(j)]; ok {
											_,boo := model.TransformTemperatrureCH(i,v)
											if !boo {
												break
											}
										} else {
											var chat model.TemperatrureChatHistory
											chat.PatientId = ii.PatientId
											chat.PatientBed = ii.BedCoding
											chat.PatientAge = ii.Age
											chat.PatientName = ii.Name
											chat.TestTime =  time.Time(i.DateTime).Format("2006-01-02")+" "+model.BackIntervalTime(i.TypeTime)
											chat.DateTime = time.Time(i.DateTime).Format("2006-01-02")
											chat.TimeFrame = strconv.Itoa(i.TypeTime)
											er,_ := model.TransformTemperatrureCH(i,&chat)
											if er != nil{
												continue
											}
											signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime) + strconv.Itoa(j)] = &chat
											break
										}
									}
								}
							} else {
								var chat model.TemperatrureChatHistory
								chat.PatientId = ii.PatientId
								chat.PatientBed = ii.BedCoding
								chat.PatientAge = ii.Age
								chat.PatientName = ii.Name
								chat.TestTime = time.Time(i.DateTime).Format("2006-01-02")+" "+model.BackIntervalTime(i.TypeTime)
								chat.DateTime = time.Time(i.DateTime).Format("2006-01-02")
								chat.TimeFrame = strconv.Itoa(i.TypeTime)
								er,_ := model.TransformTemperatrureCH(i,&chat)
								if er != nil{
									continue
								}
								signhistory_map[time.Time(i.DateTime).Format("2006-01-02") + " 时段:" + strconv.Itoa(i.TypeTime)] = &chat
							}
						}
					}
					fit.Logger().LogError("hhhhhhh",len(item))

					for _,j :=range  signhistory_map{
						signhistorys = append(signhistorys,*j)
					}
				}*/
		//}


		c.Data["Patients"] = PatientHistorys
		sort.Sort(model.PersonSlice(signhistorys))
		c.Data["signhistorys"]=signhistorys
		fit.Logger().LogError("jjjj",len(PatientHistorys))
	}
}

//体征历史数据删除
func (c PCBatvhHistoryController) TZDelete(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	delete_id := r.FormValue("delete_id")   //VAA01病人ID
	if delete_id == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

	err := session.Begin()
	if err!=nil{
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "事务开始失败"
		return
	}

	ids := strings.Split(delete_id,",")
	fit.Logger().LogError("delete_id",ids)
	for _,i := range ids{
		if i == "0"{
			continue
		}
		fit.Logger().LogError("delete_idaaaaa",i)
		err := model.DeleteTemperatureHistory(session,i)
		if err != nil {
			session.Rollback()
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "删除错误2"
			c.JsonData.Datas = err
			return
		}
	}

	err_com := session.Commit()
	if err_com != nil {
		c.JsonData.Result = 4
		c.JsonData.ErrorMsg = "数据库插入失败"
	}else{
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "删除成功"
	}
}

//体征历史数据更新
func (c PCBatvhHistoryController) TZUpdate(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

	fit.Logger().LogError("PCBatvhUpdateController","PCBatvhUpdateController")

	err := session.Begin()
	if err!=nil{
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "事务开始失败"
		return
	}

	nurse_id,err_nid     := strconv.Atoi(r.FormValue("nurse_id"))
	nurse_name           := r.FormValue("nurse_name")
	patient_id,err_pid   := utils.Int64Value(r.FormValue("patient_id"))
	test_time ,err_tm    := time.ParseInLocation("2006-01-02 15:04:05",r.FormValue("test_time"),time.Local)

	if err_nid !=nil || err_pid !=nil|| nurse_name == "" || err_tm !=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整1"
		//c.JsonData.Datas = err.Error()
		return
	}

	thm_id              := r.FormValue("thm_id")
	thm_value           := r.FormValue("thm_value")
	thm_type,err_tp     := strconv.Atoi(r.FormValue("thm_type"))
	thm_scene,err_to    := strconv.Atoi(r.FormValue("thm_scene"))
	if err_tp!=nil || err_to!=nil || thm_id == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整2"
		//c.JsonData.Datas = err_tp.Error()
		return
	}
	if thm_id != "0"{
		if thm_value != "" || (thm_scene != 0 && thm_scene != 7){
			maps := make(map[string]interface{})
			maps["value"] = thm_value
			maps["other"] = thm_scene
			err := model.UpdateTemperatrureChat(session,thm_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				fit.Logger().LogError("PCBatvhUpdateController",err)
				return
			}
		}
	}else{
	if thm_value != "" || (thm_scene != 0 && thm_scene != 7) {
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}
	}

	pulse_id      := r.FormValue("pulse_id")
	pulse_value := r.FormValue("pulse_value")
	pulse_briefness,err_bf := strconv.Atoi(r.FormValue("pulse_briefness"))
	if  err_bf!=nil || pulse_id == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整3"
		c.JsonData.Datas = err_bf.Error()
		return
	}
	if pulse_id != "0"{
		if pulse_value != "" || pulse_briefness == 1{
			maps := make(map[string]interface{})
			maps["Value"] = pulse_value
			maps["Other"] = pulse_briefness
			err := model.UpdateTemperatrureChat(session,pulse_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if pulse_value != "" || pulse_briefness == 1{
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Pulse_Type
			item.Value = pulse_value
			item.SubType = 0
			item.Other = pulse_briefness

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	breathe_id      := r.FormValue("breathe_id")
	breathe_value  := r.FormValue("breathe_value")
	breathe_scene,err_bs  := strconv.Atoi(r.FormValue("breathe_scene"))
	if  err_bs!=nil || breathe_id == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整4"
		c.JsonData.Datas = err_bs.Error()
		return
	}
	if breathe_id != "0"{
		if breathe_value != "" || breathe_scene !=0{
			maps := make(map[string]interface{})
			maps["Value"] = breathe_value
			maps["Other"] = breathe_scene
			err := model.UpdateTemperatrureChat(session,breathe_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if breathe_value != "" || breathe_scene !=0 {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Breathe_Type
			item.Value = breathe_value
			item.SubType = 0
			item.Other = breathe_scene

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	shit_id      := r.FormValue("shit_id")
	shit_value := r.FormValue("shit_value")
	shit_scene,err_ss := strconv.Atoi(r.FormValue("shit_scene"))
	if  err_ss!=nil || shit_id == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整5"
		c.JsonData.Datas = err_ss.Error()
		return
	}
	if shit_id != "0"{
		if shit_value != "" || shit_scene!=0 {
			maps := make(map[string]interface{})
			maps["Value"] = shit_value
			maps["Other"] = shit_scene
			err := model.UpdateTemperatrureChat(session,shit_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if shit_value != "" || shit_scene!=0 {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Shit_Type
			item.Value = shit_value
			item.SubType = 0
			item.Other = shit_scene

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	pressure_id      := r.FormValue("pressure_id")
	pressure_sys := r.FormValue("pressure_sys")
	pressure_dia := r.FormValue("pressure_dia")
	pressure_scene,err_ps := strconv.Atoi(r.FormValue("pressure_scene"))
	if  err_ps!=nil || pressure_id == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整6"
		c.JsonData.Datas = err_ps.Error()
		return
	}
	if pressure_id != "0"{
		if (pressure_sys != "" && pressure_dia != "") || pressure_scene != 0{
			maps := make(map[string]interface{})
			maps["Value"] =  pressure_sys + "/" + pressure_dia
			maps["Other"] = pressure_scene
			err := model.UpdateTemperatrureChat(session,pressure_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if (pressure_sys != "" && pressure_dia != "") || pressure_scene != 0 {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Pressure_Type
			item.Value =  pressure_sys + "/" + pressure_dia
			item.SubType = 0
			item.Other = pressure_scene

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	heartrate_id      := r.FormValue("heartrate_id")
	heartrate_value := r.FormValue("heartrate_value")
	if heartrate_id != "0"{
		if heartrate_value != "" {
			maps := make(map[string]interface{})
			maps["Value"] = heartrate_value
			err := model.UpdateTemperatrureChat(session,heartrate_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if heartrate_value != "" {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Heartrate_Type
			item.Value = heartrate_value
			item.SubType = 0
			item.Other = 0

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	weight_id      := r.FormValue("weight_id")
	weight_value  := r.FormValue("weight_value")
	weight_scene,err_ws  := strconv.Atoi(r.FormValue("weight_scene"))
	if  err_ws!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整8"
		c.JsonData.Datas = err_ws.Error()
		return
	}
	if weight_id != "0"{
		if weight_value != "" || weight_scene != 0 {
			maps := make(map[string]interface{})
			maps["Value"] = weight_value
			maps["Other"] = weight_scene
			err := model.UpdateTemperatrureChat(session,weight_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if weight_value != "" || weight_scene != 0 {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Weight_Type
			item.Value = weight_value
			item.SubType = 0
			item.Other = weight_scene

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	height_id      := r.FormValue("height_id")
	height_value   := r.FormValue("height_value")
	height_scene,err_hs   := strconv.Atoi(r.FormValue("height_scene"))
	if  err_hs!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整9"
		c.JsonData.Datas = err_hs.Error()
		return
	}
	if height_id != "0"{
		if height_value != "" || height_scene != 0 {
			maps := make(map[string]interface{})
			maps["Value"] = height_value
			maps["Other"] = height_scene
			err := model.UpdateTemperatrureChat(session,height_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if height_value != "" || height_scene != 0 {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Height_Type
			item.Value = height_value
			item.SubType = 0
			item.Other = height_scene

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	skin_id      := r.FormValue("skin_id")
	skin_value := r.FormValue("skin_value")
	if skin_id != "0"{
		if skin_value != "" {
			maps := make(map[string]interface{})
			maps["Value"] = skin_value
			err := model.UpdateTemperatrureChat(session,skin_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if skin_value != "" {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Skin_Type
			item.Value = skin_value
			item.SubType = 0
			item.Other = 0

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	other_id      := r.FormValue("other_id")
	other_value := r.FormValue("other_value")
	if other_id != "0"{
		if  other_value != "" {
			maps := make(map[string]interface{})
			maps["Value"] = other_value
			err := model.UpdateTemperatrureChat(session,other_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if other_value != "" {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Other_Type
			item.Value = other_value
			item.SubType = 0
			item.Other = 0

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	incident_id      := r.FormValue("incident_id")
	incident_scene,err_is := strconv.Atoi(r.FormValue("incident_scene"))
	if  err_is!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整10"
		c.JsonData.Datas = err_is.Error()
		return
	}
	if incident_id != "0"{
		if  incident_scene != 0 {
			maps := make(map[string]interface{})
			maps["Other"] = incident_scene
			err := model.UpdateTemperatrureChat(session,incident_id,maps)
			if(err!=nil){
				session.Rollback()
				c.JsonData.Result = 3
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}else {
		if incident_scene != 0 {
			var item model.NurseChat
			item.NurseName = nurse_name
			item.NurseId = nurse_id
			item.PatientId = patient_id
			item.TestTime = fit.JsonTime(test_time)

			item.HeadType = model.Incident_Type
			item.Value = ""
			item.SubType = 0
			item.Other = incident_scene

			msg, err := model.IputChat(session, item)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = msg
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err.Error()
				return
			}
		}
	}

	err_com := session.Commit()
	if err_com != nil {
		c.JsonData.Result = 4
		c.JsonData.ErrorMsg = "数据库插入失败"
	}else{
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "录入成功"
	}
}



