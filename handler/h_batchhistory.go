package handler

import (
	"nursing/model"
	"fit"
	"time"
	/*"encoding/json"*/
	/*"strconv"*/
	"strconv"
	"strings"
)

//体征历史数据
type PCBatvhHistoryController struct{
	PCController
}

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
				et = time.Date(et.Year(), et.Month(), et.Day(), 6, 0, 0, 0, st.Location())
			case "8":
				et = time.Date(et.Year(), et.Month(), et.Day(), 10, 0, 0, 0, st.Location())
			case "12":
				et = time.Date(et.Year(), et.Month(), et.Day(), 14, 0, 0, 0, st.Location())
			case "16":
				et = time.Date(et.Year(), et.Month(), et.Day(), 18, 0, 0, 0, st.Location())
			case "20":
				et = time.Date(et.Year(), et.Month(), et.Day(), 22, 0, 0, 0, st.Location())
			case "24":
				et = et.Add(1)
				et = time.Date(et.Year(), et.Month(), et.Day(), 2, 0, 0, 0, st.Location())
			}

			starttime = st.Format("2006-01-02 15:04:05")
			endtime   = et.Format("2006-01-02 15:04:05")
			fit.Logger().LogError("ghffdref",starttime,endtime)

			patientids := strings.Split(patient_id,",")
			map_patientids := make(map[string]string)
			for _,v := range patientids{
				map_patientids[v]=v
			}
			fit.Logger().LogError("ghffdref",len(patientids))
			for _,v := range  response{
				var ph model.PatientHistory
				ph.PatientId = v.VAA01
				ph.BedCoding = v.BCQ04
				ph.Name = v.VAA05
				ph.Age = strconv.Itoa(v.VAA10)
				if _, ok := map_patientids[strconv.Itoa(v.VAA01)];ok{
					ph.Checked = 1
				}else{
					ph.Checked = 0
				}
				PatientHistorys = append(PatientHistorys,ph)
			}

		}

		var signhistorys []model.TemperatrureChatHistory     //体温表数据

        if patient_id == ""{
			for _,ii := range PatientHistorys{

				signhistory_map := make(map[string]*model.TemperatrureChatHistory )

				var sql string
				var msg []interface{}

				sql = sql + "testtime >= ? and testtime <= ? and PatientId = ?"
				msg = append(msg,starttime, endtime,ii.PatientId)

				item, err := model.OutTemperatureHistory(sql, msg...)
				if err != nil {
					c.JsonData.Result = 1
					c.JsonData.ErrorMsg = "查询错误"
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
		}else{
				for _,ii := range PatientHistorys{

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
				}
		}

		c.Data["Patients"] = PatientHistorys
		c.Data["signhistorys"]=signhistorys
		fit.Logger().LogError("jjjj",len(signhistorys))
	}
}


/*func (c PCBatvhHistoryController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
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
		var signhistorys []model.TemperatrureChatHistory

		for _,ii := range patientids{

			signhistory_map := make(map[string]*model.TemperatrureChatHistory )

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

			item, err := model.OutTemperatureHistory(sql, msg...)
			if err != nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "查询错误0"
				c.JsonData.Datas = err
				return
			} else {
				for _,i := range item{
					if v, ok := signhistory_map[i.TestTime.String()]; ok {
						model.TransformTemperatrureCH(i,v)
					} else {
						var chat model.TemperatrureChatHistory
						chat.PatientId = int(patient.VAA01)
						chat.PatientBed = patient.BCQ04
						chat.PatientAge = patient.VAA101
						chat.PatientName = patient.VAA05
						chat.TestTime = i.TestTime.String()
						er := model.TransformTemperatrureCH(i,&chat)
						if er != nil{
							continue
						}
						signhistory_map[i.TestTime.String()] = &chat
					}
			}
			}
			fit.Logger().LogError("hhhhhhh",len(item))

			for _,j :=range  signhistory_map{
				signhistorys = append(signhistorys,*j)
			}
		}

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "获取成功"
		c.JsonData.Datas = signhistorys
	}
}*/




