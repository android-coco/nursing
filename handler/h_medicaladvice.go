package handler

import (
	"fit"
	"nursing/model"
	"time"
	"strconv"
)

type MedicalAdviceQuery struct {
	fit.Controller
}

//查询医嘱
func (c MedicalAdviceQuery) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	r.ParseForm()

	//advice_class := r.FormValue("advice_class")         //医嘱类别BDA01  关联BDA1.BDA01
	//advice_type := r.FormValue("advice_type")           //长期或临时 VAF11
	//advice_execution := r.FormValue("advice_execution") //用药方式 VAF53或者BBX01  关联BBX1.BBX01 （执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它）
	//advice_state := r.FormValue("advice_state")         //VAF10 状态(1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果)
	patient_id := r.FormValue("patient_id")             //VAA01病人ID


	if  patient_id =="" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}else{
	    var sql string
	    var msg []interface{}

	    sql = "VAA01 = ? and (VAF10 = ? or VAF10 = ? or VAF10 = ?)"
	    msg = append(msg, patient_id,4,5,8)

		relusts,err := model.OutAdvice(sql,msg...)
		if err != nil{
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = []interface{}{err}
		}else{
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询完成"
			c.JsonData.Datas = relusts
		}

		fit.Logger().LogError("医嘱：",relusts,err)
	}

	c.ResponseToJson(w)
}

type MedicalAdviceStateQuery struct {
	fit.Controller
}

//查询执行医嘱
func (c MedicalAdviceStateQuery) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	patient_id := r.FormValue("patient_id")   //VAA01病人ID
	starttime := r.FormValue("starttime")     //开始时间
	endtime := r.FormValue("endtime")         //结束时间

	advice_id := r.FormValue("advice_id")     //医嘱id

	if patient_id == "" && (advice_id=="" && starttime=="" && endtime==""){
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	advicefits := make([]model.AdviceFit, 0)   //返回总数据
	//var relusts []model.Advice
	//var err error

	if advice_id != ""{
		sql :=  "VAF01 = ?"
		relusts,err := model.OutAdvice(sql,advice_id)

		if err != nil{
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err
			return
		}else{
			for k,i := range relusts{
				sql := "advicestateId = ?"
				advicestates,err := model.OutAdviceState(sql,i.VAF01)
				patients, err_patient := model.GetPatientInfo(strconv.Itoa(i.VAA01))

				if err != nil || err_patient!= nil{
					c.JsonData.Result = 3
					c.JsonData.ErrorMsg = "记录查询错误"
					c.JsonData.Datas = err
					return
				}else{
					advicefits = append(advicefits,model.AdviceFit{relusts[k],advicestates,patients})
				}
			}
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询成功"
			c.JsonData.Datas = advicefits
		}
	}else{
		var sql string

		sql = sql + "(VAF11 = ? and VAA01 = ? and VAF36 >= ? and VAF36 <= ? ) or (VAF11 = ? and VAA01 = ?)"
		relusts,err := model.OutAdvice(sql,2,patient_id,starttime,endtime,1,patient_id)

		if err != nil{
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err
			return
		}else{
			for _,i := range relusts{
				sql := "advicestateId = ? and time >= ? and time <= ?"
				advicestates,err := model.OutAdviceState(sql,i.VAF01,starttime,endtime)
				patients, err_patient := model.GetPatientInfo(strconv.Itoa(i.VAA01))

				if err != nil || err_patient!= nil{
					c.JsonData.Result = 3
					c.JsonData.ErrorMsg = "记录查询错误"
					c.JsonData.Datas = []interface{}{err,err_patient}
					return
				}else{
					if len(advicestates) !=0{
						advicefits = append(advicefits,model.AdviceFit{i,advicestates,patients})
					}else{
						if i.VAF10 == 8 || i.VAF10 == 10{
							advicefits = append(advicefits,model.AdviceFit{i,advicestates,patients})
						}
					}
				}
			}
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询成功"
			c.JsonData.Datas = advicefits
		}
	}
}


type MedicalAdviceExecute struct {
	fit.Controller
}

func (c MedicalAdviceExecute) Post(w *fit.Response, r *fit.Request, p fit.Params) {
    defer c.ResponseToJson(w)
	advice_id := r.FormValue("advice_id")
	nurse_id := r.FormValue("nurse_id")
	nurse_name := r.FormValue("nurse_name")
	state := r.FormValue("state")
	period := r.FormValue("period")

	if advice_id == ""||nurse_id == ""||nurse_name == ""||state == ""||period == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}else{
		var sql string
		var msg []interface{}
		sql = "VAF01 = ?"
		msg = append(msg, advice_id)
		relusts,err := model.OutAdvice(sql,msg...)
		if err != nil{
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err
			return
		}else if len(relusts) == 0{
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "没有医嘱记录"
			c.JsonData.Datas = err
		}else{
			var advicestate  model.AdviceState
			advicestate.PatientId = relusts[0].VAA01
			advicestate.AdviceStateId = relusts[0].VAF01
			advicestate.Time = fit.JsonTime(time.Now())
			advicestate.NurseId = nurse_id
			advicestate.NurseName = nurse_name
			advicestate.State = state
			advicestate.Period = period

			err1 :=model.InsertAdviceState(advicestate)
			if err1!= nil{
				c.JsonData.Result = 4
				c.JsonData.ErrorMsg = "记录失败"
				c.JsonData.Datas = err1
			}else{
				c.JsonData.Result = 0
				c.JsonData.ErrorMsg =  "记录成功"
				c.JsonData.Datas = advicestate
			}
		}
	}
}


type MedicalAdviceNewPause struct {
	fit.Controller
}

func (c MedicalAdviceNewPause) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	newpause := r.FormValue("newpause")
	patient_id,err :=  strconv.Atoi(r.FormValue("patient_id"))

	if err != nil || len(newpause) == 0{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = err
		return
	}

	if newpause == "new"{
		newresult,err1 := model.GetNonExecutionAdvice(patient_id)
		if err1 != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err1
			return
		}
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "请求成功"
		c.JsonData.Datas = newresult
	}else{
		pauseresult,err2 := model.GetUncertainOewAdvice(patient_id)
		if err2 != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err2
			return
		}
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "请求成功"
		c.JsonData.Datas = pauseresult
	}
}
