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
	vid := r.FormValue("vid")
	starttime := r.FormValue("st")
	endtime := r.FormValue("et")

	if vid == "" || mavType == "" || mavCategory == "" || mavStatus == "" || starttime == "" || endtime == "" {
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

	_, err_p := utils.Int64Value(vid)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 vid")
		return
	}
	mAdvices, err_db := model.SearchMedicalAdviceForPC(type_i, status_i, mavCategory, vid, starttime, endtime)
	if err_db == nil {
		c.RenderingJson(0, "查询成功", mAdvices)
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
	vid := r.FormValue("vid")
	starttime := r.FormValue("st")
	endtime := r.FormValue("et")

	if vid == "" || mavType == "" || mavCategory == "" || mavStatus == "" {
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

	_, err_p := utils.Int64Value(vid)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 vid")
		return
	}
	mAdvices, err_db := model.SearchMedicalAdviceExecutionForPC(type_i, status_i, mavCategory, vid, starttime, endtime)
	if err_db == nil {
		c.RenderingJson(0, "查询成功", mAdvices)
	} else {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	}
}

/*医嘱执行明细PAI PDA*/
func (c MedicalAdviceController) ExecDetail(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	gid := r.FormValue("gid")
	ext := r.FormValue("ext")
	exc := r.FormValue("exc")


	if gid == "" || ext == "" || exc == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	gid_i, err_p := utils.Int64Value(gid)
	exc_i, err_c := strconv.Atoi(exc)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 gid")
		return
	} else if err_c != nil {
		c.RenderingJsonAutomatically(2, "参数错误 exc")
		return
	}

	mAdviceonse, err_db := model.FetchMedicalAdviceExecutionDetail(gid_i, ext, exc_i)
	if err_db == nil {
		c.RenderingJson(0, "查询成功", mAdviceonse)
	} else {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	}
}

/*查询新医嘱、已停医嘱PAI PDA*/
func (c MedicalAdviceController) StatusSearch(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	mavStatus := r.FormValue("status")
	patient_id := r.FormValue("vid")

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

	vid_i, err_p := utils.Int64Value(patient_id)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 vid")
		return
	}

	if status_i == 8 {
		// 医嘱查询（已停医嘱）
		mAdvices, err_db := model.FetchFinishedMedicalAdvice(vid_i)
		if err_db == nil {
			c.RenderingJson(0, "查询成功", mAdvices)
		} else {
			c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
		}
	} else if status_i == 3 {
		// 医嘱执行查询（新医嘱）
		mAdvices, err_db := model.FetchNewMedicalAdvice(vid_i)
		if err_db == nil {
			c.RenderingJson(0, "查询成功", mAdvices)
		} else {
			c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
		}
	}
}

/*医嘱执行PAI PDA*/
func (c MedicalAdviceController) Execute(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	vid := r.FormValue("vid")
	madid := r.FormValue("gid")
	ext := r.FormValue("ext")
	execCycle := r.FormValue("exc")

	exectime := r.FormValue("exectime")
	state := r.FormValue("state")
	nid := r.FormValue("nid")
	period := r.FormValue("period")
	process := r.FormValue("process")


	if vid == "" || madid == "" || state == "" || ext == "" || execCycle == "" || nid == "" || process == "" || period == "" || exectime == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	vid_i, err_pid := utils.Int64Value(vid)
	if err_pid != nil {
		c.RenderingJsonAutomatically(2, "参数错误 vid")
		return
	}

	mid, err_mid := utils.Int64Value(madid)
	if err_mid != nil {
		c.RenderingJsonAutomatically(2, "参数错误 madid")
		return
	}
	if isExist := model.IsExistRecord(false, "VAF2", fmt.Sprintf("VAF06 = %d and VAF01 = %d", vid_i, mid)); isExist.Exist == 0 {
		c.RenderingJsonAutomatically(4, "该医嘱与病人不匹配，无法查询到对应的医嘱记录")
		return
	}

	uid, err_uid := strconv.Atoi(nid)
	if err_uid != nil {
		c.RenderingJsonAutomatically(2, "参数错误 nid")
		return
	}
	account, err_acc := model.FetchAccountWithUid(nid)
	if err_acc != nil || account.Employeeid != uid {
		fmt.Println("***JK", uid, nid, account, err_acc)
		c.RenderingJsonAutomatically(4, "无法查询到对应工作人员")
		return
	}

	if isExist := model.IsExistRecord(false, "VAE1", fmt.Sprintf("VAE01 = %d and BCK01C = %d", vid_i, account.DepartmentID)); isExist.Exist == 0 {
		c.RenderingJsonAutomatically(4, "病人ID与用户所在病区不匹配，无法查询到对应的病人信息")
		return
	}

	period_i, err_per := strconv.Atoi(period)
	if err_per != nil {
		c.RenderingJsonAutomatically(2, "参数错误 period")
		return
	}
	state_i, err_st := strconv.Atoi(state)
	if err_st != nil || state_i < 1 || state_i > 3 {
		c.RenderingJsonAutomatically(2, "参数错误 state")
		return
	}

	exc, err_exc := strconv.Atoi(execCycle)
	if err_exc != nil {
		c.RenderingJsonAutomatically(2, "参数错误 exc")
		return
	}

	mAdvice := model.MedicalAdviceModal{}
	_, err_sql := fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT d.* FROM( SELECT a.VAF06 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName,CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, a.VAF21 Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 LEFT JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF06 = %d AND a.VAF01 = %d AND b.VBI10 = '%s') d ORDER BY d.TypeV, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr ", vid_i, mid, ext)).Get(&mAdvice)
	if err_sql != nil {
		fit.Logger().LogError("***JK***", err_sql)
	}

	if mAdvice.Vid != vid_i && mAdvice.Madid != mid {
		c.RenderingJsonAutomatically(4, "医嘱不存在")
		return
	}

	obj := model.MedicalAdviceItem{
		Vid:       mAdvice.Vid,
		Madid:     mid,
		Pid:       mAdvice.Pid,
		Bed:       mAdvice.Bed,
		PName:     mAdvice.PName,
		Gender:    mAdvice.Gender,
		Age:       mAdvice.Age,
		HospNum:   mAdvice.HospNum,
		ExTime:    mAdvice.ExTime.ParseToSecond(),
		GroupNum:  mAdvice.GroupNum,
		Content:   mAdvice.Content,
		Dosage:    mAdvice.Dosage,
		Amount:    mAdvice.Amount,
		Frequency: mAdvice.Frequency,
		Times:     mAdvice.Times,
		Method:    mAdvice.Method,
		Speed:     mAdvice.Speed,
		TypeV:     mAdvice.TypeV,
		StTime:    mAdvice.StTime.ParseToSecond(),
		MStatusV:  mAdvice.MStatusV,
		Category:  mAdvice.Category,
		CategoryV: mAdvice.CategoryV,
		PtType:    mAdvice.PtType,
		PtNum:     mAdvice.PtNum,
		PtRownr:   mAdvice.PtRownr,
		Entrust:   mAdvice.Entrust,
		Physician: mAdvice.Physician,
		EdTime:    mAdvice.EdTime.ParseToSecond(),
		Sender:    mAdvice.Sender,
		ExCycle:   exc,
		ExNurse:   account.Username,
		ExStatusV: state_i,
		ExStep:    process,
		PtTimes:   mAdvice.PtTimes,
	}

	execution := model.MedicalAdviceExecutionRecord{
		Patientid: obj.Pid,
		Nurseid:   uid,
		Period:    period_i,
		Madid:     obj.Madid,
		Nursename: obj.ExNurse,
		Process:   obj.ExStep,
		ExecTime:  exectime,
		ExCycle:   obj.ExCycle,
		Plan:      obj.ExTime,
	}

	if isExist := model.IsExistRecord(true, "medicaladvice", fmt.Sprintf("Madid = %d and ExTime = '%s' and ExCycle = %d ", mid, ext, exc)); isExist.Exist == 0 {
		_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
		if err_in != nil {
			fit.Logger().LogError("***JK***", "医嘱执行-Insert", err_in.Error())
		}
	} else {
		//	有记录
		sqlStr := fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExStatusV = '%d', ExStep = '%s', ExNurse = '%s' where Madid = %d and ExTime = '%s' and ExCycle = %d", obj.Bed, obj.GroupNum, obj.Times, obj.MStatusV, obj.PtNum, obj.EdTime, obj.Sender, obj.ExStatusV, obj.ExStep, obj.ExNurse, obj.Madid, obj.ExTime, obj.ExCycle)
		_, err_up := fit.MySqlEngine().Exec(sqlStr)
		if err_up != nil {
			fit.Logger().LogError("***JK***", "医嘱执行-update", err_up.Error())
		}
	}

	_, err_in := fit.MySqlEngine().Table("advicedetail").Insert(&execution)
	if err_in != nil {
		fit.Logger().LogError("***JK***", err_in.Error())
		c.RenderingJsonAutomatically(3, "DataBase "+err_in.Error())
		return
	} else {
		res := make([]model.MedicalAdviceExecutionRecord, 1)
		res[0] = execution
		c.RenderingJson(0, "执行成功", res)
	}

	//status := model.AdviceStatus {
	//	Patientid:pid_i,
	//	Madid:mid,
	//	State:state_i,
	//	Recordtime:exectime,
	//	Nurseid:account.Employeeid,
	//	Nursename:account.Username,
	//	Period:period_i,
	//	Process:process,
	//}
	//detail := model.AdviceDetail {
	//	ExecTime: exectime,
	//	Process:process,
	//	Madid:mid,
	//	Nursename:account.Username,
	//	Nurseid:account.Employeeid,
	//	Patientid:pid_i,
	//	Period:period_i,
	//}
	//
	//// 1. 存在执行记录即更新AdviceStatus，否则往AdviceStatus插入记录。
	//// 2. 往AdviceDetail插入执行明细
	//if isExist := model.IsExistRecord(true, "AdviceStatus", fmt.Sprintf("Madid = %d",mid)); isExist.Exist == 0 {
	//	_, err_db := fit.MySqlEngine().Table("AdviceStatus").InsertOne(&status)
	//	if err_db != nil {
	//		c.RenderingJsonAutomatically(3,"Database（AdviceStatus）"+err_db.Error())
	//	} else {
	//		_, err_db = fit.MySqlEngine().Table("AdviceDetail").InsertOne(&detail)
	//		if err_db != nil {
	//			c.RenderingJsonAutomatically(3,"Database（AdviceDetail）"+err_db.Error())
	//		} else {
	//			res := make([]model.AdviceDetail, 1)
	//			res[0] = detail
	//			c.RenderingJson(0, "执行成功",res)
	//		}
	//	}
	//} else {
	//	_, err_db := fit.MySqlEngine().Table("AdviceStatus").Where("Madid = ?",mid).Update(&status)
	//	if err_db != nil {
	//		c.RenderingJsonAutomatically(3,"Database（AdviceStatus）"+err_db.Error())
	//	} else {
	//		_, err_db = fit.MySqlEngine().Table("AdviceDetail").InsertOne(&detail)
	//		if err_db != nil {
	//			c.RenderingJsonAutomatically(3,"Database（AdviceDetail）"+err_db.Error())
	//		} else {
	//			res := make([]model.AdviceDetail, 1)
	//			res[0] = detail
	//			c.RenderingJson(0, "执行成功",res)
	//		}
	//	}
	//}
}

func (c *MedicalAdviceController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *MedicalAdviceController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
