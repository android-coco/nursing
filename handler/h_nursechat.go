package handler

import (
	"fit"
	"nursing/model"
	"strconv"
	"time"
	"nursing/utils"
)

//体征PDA路由
type NurseChatController struct {
	fit.Controller
}

//pda体征录入
func (c NurseChatController) NurseChatInput(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

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

	thm_value           := r.FormValue("thm_value")
	thm_type,err_tp     := strconv.Atoi(r.FormValue("thm_type"))
	thm_scene,err_to    := strconv.Atoi(r.FormValue("thm_scene"))
	if err_tp!=nil || err_to!=nil{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整2"
		//c.JsonData.Datas = err_tp.Error()
		return
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	pulse_value := r.FormValue("pulse_value")
	pulse_briefness,err_bf := strconv.Atoi(r.FormValue("pulse_briefness"))
	if  err_bf!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整3"
		c.JsonData.Datas = err_bf.Error()
		return
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	breathe_value  := r.FormValue("breathe_value")
	breathe_scene,err_bs  := strconv.Atoi(r.FormValue("breathe_scene"))
	if  err_bs!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整4"
		c.JsonData.Datas = err_bs.Error()
		return
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	shit_value := r.FormValue("shit_value")
	shit_scene,err_ss := strconv.Atoi(r.FormValue("shit_scene"))
	if  err_ss!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整5"
		c.JsonData.Datas = err_ss.Error()
		return
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	pressure_value := r.FormValue("pressure_value")
	pressure_scene,err_ps := strconv.Atoi(r.FormValue("pressure_scene"))
	if  err_ps!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整6"
		c.JsonData.Datas = err_ps.Error()
		return
	}
	if pressure_value != "" || pressure_scene != 0{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Pressure_Type
		item.Value = pressure_value
		item.SubType = 0
		item.Other = pressure_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	heartrate_value := r.FormValue("heartrate_value")
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	spo2h_value   := r.FormValue("spo2h_value")
	if spo2h_value != "" {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Spo2h_Type
		item.Value = spo2h_value
		item.SubType = 0
		item.Other = 0

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	glucose_value  := r.FormValue("glucose_value")
	glucose_type,err_gt  := strconv.Atoi(r.FormValue("glucose_type"))
	glucose_scene,err_gs  := strconv.Atoi(r.FormValue("glucose_scene"))
	if  err_gt!=nil && err_gs != nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整7"
		//c.JsonData.Datas = []interface{}{err_gt.Error(),err_gs.Error()}
		return
	}
	if glucose_value != "" || glucose_scene != 0 || glucose_type!=0{
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Glucose_Type
		item.Value = glucose_value
		item.SubType = glucose_type
		item.Other = glucose_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	weight_value  := r.FormValue("weight_value")
	weight_scene,err_ws  := strconv.Atoi(r.FormValue("weight_scene"))
	if  err_ws!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整8"
		c.JsonData.Datas = err_ws.Error()
		return
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	height_value   := r.FormValue("height_value")
	height_scene,err_hs   := strconv.Atoi(r.FormValue("height_scene"))
	if  err_hs!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整9"
		c.JsonData.Datas = err_hs.Error()
		return
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	skin_value := r.FormValue("skin_value")
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	other_value := r.FormValue("other_value")
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
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
		}
	}

	incident_scene,err_is := strconv.Atoi(r.FormValue("incident_scene"))
	if  err_is!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整10"
		c.JsonData.Datas = err_is.Error()
		return
	}
	if incident_scene != 0 {
		var item model.NurseChat
		item.NurseName = nurse_name
		item.NurseId  =  nurse_id
		item.PatientId = patient_id
		item.TestTime = fit.JsonTime(test_time)

		item.HeadType = model.Incident_Type
		item.Value = ""
		item.SubType = 0
		item.Other = incident_scene

		msg,err := model.IputChat(session,item)
		if(err!=nil){
			session.Rollback()
			c.JsonData.Result = msg
			c.JsonData.ErrorMsg = "参数错误1"
			c.JsonData.Datas = err.Error()
			return
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

//pda体征获取
func (c NurseChatController) NurseChatOutput(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	patientid := r.FormValue("patientid")
	page,err_page  := strconv.Atoi(r.FormValue("page"))
	headtype := r.FormValue("headtype")

	var sql string
	var msg []interface{}

	if len(patientid) == 0{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		return
	}

	fit.Logger().LogError("ghhhh","kkkk")

	sql = sql + "patientid = ?"
	msg = append(msg, patientid)


	if len(headtype) != 0 {
		sql = sql + " and HeadType = ?"
		msg = append(msg, headtype)
	}else{
		sql = sql + " and HeadType in(1,2,3,4,5,6,7,8,9,10,12,14,13)"  //不返回出量入量
	}

	if err_page == nil{
		sql = sql + " order by Testtime desc,Id desc limit ?,10"
		msg = append(msg, 10*page)
	}

	items := make([]model.NurseChat, 0)

	err := fit.MySqlEngine().SQL("select HeadType, TestTime, SubType, Other, Value, PatientId, NurseId, NurseName from (SELECT HeadType, TestTime, SubType, Other, Value, PatientId, NurseId, NurseName, Id from NurseChat UNION ALL SELECT HeadType, TestTime, SubType, Other, Value, PatientId, NurseId, NurseName , Id from TemperatrureChat) alias WHERE " + sql ,msg...).Find(&items)

	if err != nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = err.Error()
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

}

