package handler

import (
	"fit"
	"nursing/model"
	"strconv"
	"nursing/utils"
	"fmt"
)

type MedicalAdviceController struct {
	fit.Controller
}


/*查询医嘱API PDA*/
func (c MedicalAdviceController) Search(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	mavType := r.FormValue("type")
	mavCategory := r.FormValue("category")
	mavStatus := r.FormValue("status")
	patient_id := r.FormValue("pid")

	if patient_id == "" || mavType == "" || mavCategory == "" || mavStatus == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	type_i, err_t := strconv.Atoi(mavType)
	if err_t != nil {
		c.RenderingJsonAutomatically(2, "参数错误 type")
		return
	}

	status_i, err_s := strconv.Atoi(mavStatus)
	if err_s != nil {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	pid_i, err_p := utils.Int64Value(patient_id)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 pid")
		return
	}

	mAdvices, err_db := model.SearchMedicalAdvice(type_i, status_i, pid_i, mavCategory)
	if err_db == nil {
		c.RenderingJson(0,"查询成功", mAdvices)
	} else {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	}
}


/*医嘱执行查询PAI PDA*/
func (c MedicalAdviceController) ExecSearch(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	mavType := r.FormValue("type")
	mavCategory := r.FormValue("category")
	mavStatus := r.FormValue("status")
	patient_id := r.FormValue("pid")

	if patient_id == "" || mavType == "" || mavCategory == "" || mavStatus == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	type_i, err_t := strconv.Atoi(mavType)
	if err_t != nil {
		c.RenderingJsonAutomatically(2, "参数错误 type")
		return
	}

	status_i, err_s := strconv.Atoi(mavStatus)
	if err_s != nil {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	pid_i, err_p := utils.Int64Value(patient_id)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 pid")
		return
	}

	mAdvices, err_db := model.SearchMedicalAdviceExecution(type_i, status_i, pid_i, mavCategory)
	if err_db == nil {
		c.RenderingJson(0,"查询成功", mAdvices)
	} else {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	}
}

/*医嘱执行明细PAI PDA*/
func (c MedicalAdviceController) ExecDetail(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	madid := r.FormValue("madid")
	master := r.FormValue("master")

	if madid == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	madid_i, err_p := utils.Int64Value(madid)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 madid")
		return
	}

	master_i, err_m := strconv.Atoi(master)
	if err_m != nil || master_i < 0 || master_i > 1 {
		c.RenderingJsonAutomatically(2, "参数错误 master")
		return
	}


	response, err_db := model.FetchMedicalAdviceExecutionDetail(madid_i, master_i)
	if err_db == nil {
		c.RenderingJson(0,"查询成功", response)
	} else {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	}
}

/*查询新医嘱、已停医嘱PAI PDA*/
func (c MedicalAdviceController) StatusSearch(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	mavStatus := r.FormValue("status")
	patient_id := r.FormValue("pid")

	if patient_id == "" || mavStatus == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	status_i, err_s := strconv.Atoi(mavStatus)
	if err_s != nil {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}
	if status_i != 3 && status_i != 8 {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	pid_i, err_p := utils.Int64Value(patient_id)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 pid")
		return
	}

	if status_i == 8 {
		// 医嘱查询（已停医嘱）
		mAdvices, err_db := model.SearchMedicalAdvice(0, 8, pid_i, "0")
		if err_db == nil {
			c.RenderingJson(0,"查询成功", mAdvices)
		} else {
			c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
		}
	} else if status_i == 3 {
		// 医嘱执行查询（新医嘱）
		mAdvices, err_db := model.SearchMedicalAdviceExecution(0, 1, pid_i, "0")
		if err_db == nil {
			c.RenderingJson(0,"查询成功", mAdvices)
		} else {
			c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
		}
	}
}

/*医嘱执行PAI PDA*/
func (c MedicalAdviceController) Execute(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	pid := r.FormValue("pid")
	madid := r.FormValue("madid")
	state := r.FormValue("state")
	exectime := r.FormValue("exectime")
	nid := r.FormValue("nid")
	period := r.FormValue("period")
	process := r.FormValue("process")

	if pid == "" || madid == "" || state == "" || exectime == "" || nid == "" || period == "" || process == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	pid_i, err_pid := utils.Int64Value(pid)
	if err_pid != nil {
		c.RenderingJsonAutomatically(2,"参数错误 pid")
		return
	}

	mid, err_mid := utils.Int64Value(madid)
	if err_mid != nil {
		c.RenderingJsonAutomatically(2,"参数错误 madid")
		return
	}
	if isExist := model.IsExistRecord(false, "VAF2", fmt.Sprintf("VAA01 = %d and VAF01 = %d",pid_i,mid)); isExist.Exist == 0 {
		c.RenderingJsonAutomatically(4,"该医嘱与病人ID不匹配，无法查询到对应的医嘱记录")
		return
	}

	uid, err_uid := strconv.Atoi(nid)
	if err_uid != nil {
		c.RenderingJsonAutomatically(2,"参数错误 nid")
		return
	}
	account,err_acc := model.FetchAccountWithUid(nid)
	if err_acc != nil || account.Employeeid != uid {
		fmt.Println("***JK",uid, nid ,account, err_acc)
		c.RenderingJsonAutomatically(4,"无法查询到对应工作人员")
		return
	}

	if isExist := model.IsExistRecord(false, "VAE1", fmt.Sprintf("VAA01 = %d and BCK01C = %d",pid_i, account.DepartmentID)); isExist.Exist == 0 {
		c.RenderingJsonAutomatically(4,"病人ID与用户所在病区不匹配，无法查询到对应的病人信息")
		return
	}

	period_i, err_per := strconv.Atoi(period)
	if err_per != nil {
		c.RenderingJsonAutomatically(2,"参数错误 period")
		return
	}
	state_i, err_st := strconv.Atoi(state)
	if err_st != nil || state_i < 1 || state_i > 3 {
		c.RenderingJsonAutomatically(2,"参数错误 state")
		return
	}



	status := model.AdviceStatus {
		Patientid:pid_i,
		Madid:mid,
		State:state_i,
		Recordtime:exectime,
		Nurseid:account.Employeeid,
		Nursename:account.Username,
		Period:period_i,
		Process:process,
	}
	detail := model.AdviceDetail {
		ExecTime: exectime,
		Process:process,
		Madid:mid,
		Nursename:account.Username,
		Nurseid:account.Employeeid,
		Patientid:pid_i,
		Period:period_i,
	}

	// 1. 存在执行记录即更新AdviceStatus，否则往AdviceStatus插入记录。
	// 2. 往AdviceDetail插入执行明细
	if isExist := model.IsExistRecord(true, "AdviceStatus", fmt.Sprintf("Madid = %d",mid)); isExist.Exist == 0 {
		_, err_db := fit.MySqlEngine().Table("AdviceStatus").InsertOne(&status)
		if err_db != nil {
			c.RenderingJsonAutomatically(3,"Database（AdviceStatus）"+err_db.Error())
		} else {
			_, err_db = fit.MySqlEngine().Table("AdviceDetail").InsertOne(&detail)
			if err_db != nil {
				c.RenderingJsonAutomatically(3,"Database（AdviceDetail）"+err_db.Error())
			} else {
				res := make([]model.AdviceDetail, 1)
				res[0] = detail
				c.RenderingJson(0, "执行成功",res)
			}
		}
	} else {
		_, err_db := fit.MySqlEngine().Table("AdviceStatus").Where("Madid = ?",mid).Update(&status)
		if err_db != nil {
			c.RenderingJsonAutomatically(3,"Database（AdviceStatus）"+err_db.Error())
		} else {
			_, err_db = fit.MySqlEngine().Table("AdviceDetail").InsertOne(&detail)
			if err_db != nil {
				c.RenderingJsonAutomatically(3,"Database（AdviceDetail）"+err_db.Error())
			} else {
				res := make([]model.AdviceDetail, 1)
				res[0] = detail
				c.RenderingJson(0, "执行成功",res)
			}
		}
	}
}

func (c *MedicalAdviceController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *MedicalAdviceController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
