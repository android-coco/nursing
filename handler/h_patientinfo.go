//  Created by JP

package handler

import (
	"fit"
	"nursing/model"
	"strconv"
	"time"
	"nursing/utils"
)

type PatientInfoController struct {
	fit.Controller
}

/*API 获取病人信息*/
func (c PatientInfoController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	patient_id := r.FormValue("patient_id")
	department_id := r.FormValue("department_id")
	if patient_id == "" || department_id == "" {
		c.RenderingJsonAutomatically(1, "参数不完整, null value: patient_id or department_id")
		return
	}
	pid, err_pid := strconv.Atoi(patient_id)
	if err_pid != nil || pid == 0 {
		c.RenderingJsonAutomatically(2, "参数错误, patient_id")
		return
	}

	did, err_did := strconv.Atoi(department_id)
	if err_did != nil || did == 0 {
		c.RenderingJsonAutomatically(2, "参数错误, department_id")
		return
	}

	patients, err := model.QueryPatientInfo(pid, did)
	if err != nil {
		c.RenderingJsonAutomatically(3, "Database "+err.Error())
	} else {

		if length := len(patients); length > 0 {
			c.RenderingJson(0, "成功", patients)
		} else  {
			c.RenderingJson(4, "非当前科室在床病人，无法查找相应患者信息", make([]interface{}, 0))
		}
	}
}

/*修改入科日期*/
func (c PatientInfoController) UpdateEntry(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	entryDay := r.FormValue("day")
	entryTime := r.FormValue("time")
	pid := r.FormValue("pid")
	if "" == entryDay || "" == entryTime || pid == "" {
		c.RenderingJsonAutomatically(1,"参数不完整")
		return
	}

	entry := entryDay + " " + entryTime + ":00"
	if _, err := time.Parse("2006-01-02 15:04:05", entry);err != nil {
		c.RenderingJsonAutomatically(2,"参数错误 day+time "+entryDay+" " + entryTime)
		return
	}

	pid_i, err := utils.Int64Value(pid)
	if pid_i == 0 || err != nil {
		 c.RenderingJsonAutomatically(2,"参数错误 pid "+pid)
		 return
	}

	err = model.EnteringEntryDepartmentDate(pid_i, entry)
	if err != nil {
		c.RenderingJsonAutomatically(3, "Database "+err.Error())
	} else {
		c.RenderingJsonAutomatically(0,"录入成功")
	}
}


func (c *PatientInfoController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *PatientInfoController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
