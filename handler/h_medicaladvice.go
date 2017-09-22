package handler

import "fit"

type MedicalAdvice struct {
	fit.Controller
}

type AdviceRulst struct {
	VAF36  string `json:"start_time"`  //开始执行时间
	VAF11  int    `json:"advice_type"` //长期或临时
	VAF22  string `json:"advice_msg"`  //医嘱内容
	BCE03A string `json:"doctor_name"` //医生名称
	VAF18  int    `json:"single_dose"` //剂量, 单次用量
}

type AdviceClass struct {
	//BDA01字段关联
	BDA02 string //"诊疗项目"
}

func (c MedicalAdvice) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	fit.MySqlEngine().ShowSQL(true)
	relusts := make([]AdviceRulst, 0)
	defer c.ResponseToJson(w)
	defer setMsg(c,relusts)
	r.ParseForm()
	advice_class := r.FormValue("advice_class")         //医嘱类别BDA01  关联BDA1.BDA01
	advice_type := r.FormValue("advice_type")           //长期或临时 VAF11
	advice_execution := r.FormValue("advice_execution") //BBX01  关联BBX1.BBX01用药方式（执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它）
	advice_state := r.FormValue("advice_state")         //VAF53状态(1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果)
	patient_id := r.FormValue("patient_id")             //VAA01病人ID

	if advice_class == "" || advice_type == "" || advice_execution == "" || advice_state == "" || patient_id ==""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	var whereStr string
	whereStr ="BDA01=? AND VAF11=? AND BBX01=? AND VAF53=? AND VAA01=?"
	if advice_class == "all" {
		whereStr ="VAF11=? AND BBX01=? AND VAF53=? AND VAA01=?"
		fit.MySqlEngine().SQL("SELECT * FROM VAF2").Where(whereStr,advice_type, advice_execution, advice_state, patient_id).Find(&relusts)
		return
	}
	if advice_execution == "all" {
		whereStr ="BDA01=? AND VAF11=? AND VAF53=? AND VAA01=?"
		fit.MySqlEngine().SQL("SELECT * FROM VAF2").Where(whereStr,advice_class, advice_type, advice_state, patient_id).Find(&relusts)
		return
	}
	if advice_state == "all" {
		whereStr ="BDA01=? AND VAF11=? AND BBX01=? AND VAA01=?"
		fit.MySqlEngine().SQL("SELECT * FROM VAF2").Where(whereStr,advice_class, advice_type, advice_execution, patient_id).Find(&relusts)
		return
	}
	fit.MySqlEngine().Table("VAF2").Where(whereStr,advice_class, advice_type, advice_execution, advice_state, patient_id).Find(&relusts)
	fit.Logger().LogError("医嘱：",relusts)
}

func setMsg(c MedicalAdvice,results interface{})  {
	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "OK"
	c.JsonData.Datas = results
}