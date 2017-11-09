package handler

import (
	"fit"
	/*"time"*/
	"nursing/model"
	/*"strconv"*/
	/*"strconv"*/
	"time"
	"encoding/json"
	/*"strconv"*/
)

type MedicalAdviceMessage struct {
	PCController
}

//查询医嘱
func (c MedicalAdviceMessage) Get(w *fit.Response, r *fit.Request, p fit.Params) {

	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadView(w, "pc/v_medicaladvice_message.html", "pc/header_side.html","pc/header_top.html")

		t := time.Now()
		starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		endtime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}
		Data["Patients"] = response
		fit.Logger().LogError("gk dd", len(response))

		advicefits := make([]model.AdviceFit, 0)   //返回总数据


		for _,i := range response{
			sql := "VAA01 = ? and VAF36 >= ? and VAF36 <= ?"
			relusts,err := model.OutAdvice(sql,i.VAA01,starttime,endtime)
			//124
			if err != nil {
				return
			}else{
				for _,h := range relusts{
					advicefits = append(advicefits,model.AdviceFit{h,nil,i})
				}
			}
		}

		Data["advicefits"] = advicefits
		Data["Userinfo"] = userinfo
		Data["Menuindex"] = "3-0"
        c.Data = Data
	}
}

func (c MedicalAdviceMessage) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	//advice_class := r.FormValue("advice_class")         //医嘱类别BDA01  关联BDA1.BDA01
	//advice_class == ""||
	advice_type := r.FormValue("advice_type")           //长期或临时 VAF11
	advice_execution := r.FormValue("advice_execution") //用药方式 VAF53或者BBX01  关联BBX1.BBX01 （执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它）
	advice_state := r.FormValue("advice_state")         //VAF10 状态(1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果)
	patient_id := r.FormValue("patient_id")             //VAA01病人ID

	starttime := r.FormValue("starttime")     //开始时间
	endtime := r.FormValue("endtime")         //结束时间

	if advice_type == ""||advice_execution == ""||advice_state == ""||patient_id == ""||starttime == ""||endtime == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}else{
		patientids := make([]string,0)
		err := json.Unmarshal([]byte(patient_id), &patientids)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "解析错误"
			c.JsonData.Datas = []interface{}{}
			return
		}else{
			advicefits := make([]model.AdviceFit, 0)   //返回总数据

			for _,i := range patientids{
				var sql string
				var msg []interface{}

				sql = sql + "VAA01 = ?"
				msg = append(msg, i)

				if advice_type != "全部" {
					sql = sql + " and VAF11 = ?"
					msg = append(msg, advice_type)
				}
				if advice_execution != "全部" {
					sql = sql + " and VAF53 = ?"
					msg = append(msg, advice_execution)
				}
				if advice_state != "全部" {
					sql = sql + " and VAF10 = ?"
					msg = append(msg, advice_state)
				}

				sql = sql + " and VAF36 >= ?"
				msg = append(msg, starttime)

				sql = sql + " and VAF36 <= ?"
				msg = append(msg, endtime)

				fit.Logger().LogError("gggghhhh",sql)

				relusts,err := model.OutAdvice(sql,msg...)
				fit.Logger().LogError("gggghhhh",err)
				patients, err_patient := model.GetPatientInfo(i)

				if err!=nil || err_patient!=nil{
					c.JsonData.Result = 3
					c.JsonData.ErrorMsg = "记录查询错误"
					c.JsonData.Datas = []interface{}{err,err_patient}
					return
				}else{
					for _,j := range relusts{
						advicefits = append(advicefits,model.AdviceFit{j,nil,patients})
					}
				}
			}

			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询成功"
			c.JsonData.Datas = advicefits
		}
	}
}


type MedicalAdviceDetail struct {
	PCController
}

func (c MedicalAdviceDetail) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadView(w, "pc/v_medicaladvice_message2.html", "pc/header_side.html","pc/header_top.html")

		t := time.Now()
		starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		endtime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}
		Data["Patients"] = response
		fit.Logger().LogError("gk dd", len(response))

		advicefits := make([]model.AdviceFit, 0)   //返回总数据

		for _,i := range response{
			sql := "VAA01 = ? and VAF36 >= ? and VAF36 <= ?"
			relusts,err := model.OutAdvice(sql,124,starttime,endtime)
			//i.VAA01
			sql = "advicestateId = ? and time >= ? and time <= ?"

			if err != nil {
				return
			}else{
				for _,h := range relusts{
					advicestates,_ := model.OutAdviceState(sql,h.VAF01,starttime,endtime)
					advicefits = append(advicefits,model.AdviceFit{h,advicestates,i})
				}
			}
		}

		Data["advicefits"] = advicefits
		Data["Userinfo"] = userinfo
		Data["Menuindex"] = "3-0"
		fit.Logger().LogError("sb  gk dd", len(advicefits))
		c.Data = Data
	}
}

func (c MedicalAdviceDetail) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	//advice_class := r.FormValue("advice_class")         //医嘱类别BDA01  关联BDA1.BDA01
	//advice_class == ""||
	advice_type := r.FormValue("advice_type")           //长期或临时 VAF11
	advice_execution := r.FormValue("advice_execution") //用药方式 VAF53或者BBX01  关联BBX1.BBX01 （执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它）
	advice_state := r.FormValue("advice_state")         //VAF10 状态(1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果)
	patient_id := r.FormValue("patient_id")             //VAA01病人ID

	starttime := r.FormValue("starttime")     //开始时间
	endtime := r.FormValue("endtime")         //结束时间

	if advice_type == ""||advice_execution == ""||advice_state == ""||patient_id == ""||starttime == ""||endtime == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}else{
		patientids := make([]string,0)
		err := json.Unmarshal([]byte(patient_id), &patientids)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "解析错误"
			c.JsonData.Datas = []interface{}{}
			return
		}else{
			advicefits := make([]model.AdviceFit, 0)   //返回总数据

			for _,i := range patientids{
				var sql string
				var msg []interface{}

				sql = sql + "VAA01 = ?"
				msg = append(msg, i)

				if advice_type != "全部" {
					sql = sql + " and VAF11 = ?"
					msg = append(msg, advice_type)
				}
				if advice_execution != "全部" {
					sql = sql + " and VAF53 = ?"
					msg = append(msg, advice_execution)
				}
				if advice_state != "全部" {
					sql = sql + " and VAF10 = ?"
					msg = append(msg, advice_state)
				}

				sql = sql + " and VAF36 >= ?"
				msg = append(msg, starttime)

				sql = sql + " and VAF36 <= ?"
				msg = append(msg, endtime)

				relusts,err := model.OutAdvice(sql,msg...)
				patients, err_patient := model.GetPatientInfo(i)

				if err!=nil || err_patient!=nil{
					c.JsonData.Result = 3
					c.JsonData.ErrorMsg = "记录查询错误"
					c.JsonData.Datas = []interface{}{err,err_patient}
					return
				}else{
					for _,j := range relusts{
						advicestates,_ := model.OutAdviceState( "advicestateId = ? and time >= ? and time <= ?",j.VAF01,starttime,endtime)
						advicefits = append(advicefits,model.AdviceFit{j,advicestates,patients})
					}
				}
			}
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询成功"
			c.JsonData.Datas = advicefits
		}
	}
}


type MedicalAdviceSplit struct {
	PCController
}


func (c MedicalAdviceSplit) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadView(w, "pc/v_medicaladvice_split.html", "pc/header_side.html","pc/header_top.html")

		t := time.Now()
		starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		endtime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}
		Data["Patients"] = response
		fit.Logger().LogError("gk dd", len(response))

		advicefits := make([]model.AdviceFit, 0)   //返回总数据


		for _,i := range response{
			sql := "VAA01 = ? and VAF36 >= ? and VAF36 <= ?"
			relusts,err := model.OutAdvice(sql,i.VAA01,starttime,endtime)
			//124
			if err != nil {
				return
			}else{
				for _,h := range relusts{
					advicefits = append(advicefits,model.AdviceFit{h,nil,i})
				}
			}
		}

		Data["advicefits"] = advicefits
		Data["Userinfo"] = userinfo
		Data["Menuindex"] = "2-0"
		c.Data = Data
	}
}


func (c MedicalAdviceSplit) Post(w *fit.Response, r *fit.Request, p fit.Params) {

}




