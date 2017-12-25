//  Created by JP

package handler

import (
	"nursing/model"
	"fit"
	"strconv"
	"time"
)

/*出入量管理页面*/
type IntakeOutputCollectController struct {
	fit.Controller
}
type IntakeOutputQueryController struct {
	fit.Controller
}

/*API 提交出入量*/
func (c IntakeOutputCollectController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	if length := len(r.Form); length < 9 {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	nid := r.FormValue("nurse_id")
	pid := r.FormValue("patient_id")
	recordtime := r.FormValue("recordTime")
	tp := r.FormValue("type")
	subtype := r.FormValue("subtype")
	value := r.FormValue("value")
	desc := r.FormValue("desc")
	nurse_name := r.FormValue("nurse_name")
	opertion_type := r.FormValue("opertion_type")
	other_desc := r.FormValue("other_desc")

	if nurse_name == "" || nid == "" || pid == "" || recordtime == "" || tp == "" || subtype == "" || value == "" {
		c.RenderingJsonAutomatically(1, "参数不完整 null value")
		return
	}

	tp_i, err := strconv.Atoi(tp)
	if err != nil || tp_i < 15 || tp_i > 16 {
		c.RenderingJsonAutomatically(2, "参数错误: type")
		return
	}

	patientID, err_pid := strconv.ParseInt(pid, 10, 64)
	if err_pid != nil {
		c.RenderingJsonAutomatically(2, "参数错误: patient_id")
		return
	}
	nurseId, err_nid := strconv.Atoi(nid)
	if err_nid != nil {
		c.RenderingJsonAutomatically(2, "参数错误: nurse_id")
		return
	}

	subtype_i, err_copy := strconv.Atoi(subtype)
	testtime, err_time := time.ParseInLocation("2006-01-02 15:04:05", recordtime,time.Local)
	_, err_value := strconv.Atoi(value)
	opertion_i, err_op := strconv.Atoi(opertion_type)

	if err_time != nil {
		c.RenderingJsonAutomatically(2, "参数错误: recordtime")
		return
	} else if err_value != nil {
		c.RenderingJsonAutomatically(2, "参数错误: value")
		return
	} else if err_copy != nil {
		c.RenderingJsonAutomatically(2, "参数错误: subtype")
		return
	} else if err_op != nil {
		c.RenderingJsonAutomatically(2, "参数错误: opertion_type")
		return
	}

	// 出量
	iov := model.NurseChat{
		PatientId: patientID,
		NurseId:   nurseId,
		NurseName: nurse_name,
		HeadType:  tp,
		SubType:      subtype_i,
		Other:     opertion_i,
		OtherStr:  other_desc,
		Describe:  desc,
		TestTime:  model.FitTime(testtime),
		Value:     value,
	}

	fit.Logger().LogError("IntakeOutputCollectController",iov,iov.TestTime,recordtime)

	err = model.QueryIntakeOrOutput(&iov)

	//err = iov.CollectIntakeOrOutputVolume()

	if err != nil {
		c.RenderingJsonAutomatically(3, "Database "+err.Error())
	} else {
		c.RenderingJsonAutomatically(0, "成功")
	}
}

func (c *IntakeOutputCollectController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *IntakeOutputCollectController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}

/*查询出入量*/
func (c IntakeOutputQueryController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	tp := r.FormValue("type")
	page := r.FormValue("page")
	pid := r.FormValue("patient_id")
	if tp == "" || page == "" || pid == "" {
		c.RenderingJsonAutomatically(1, "参数不完整 null value")
		return
	}
	tp_i, err_tp := strconv.Atoi(tp)
	page_i, err_dup := strconv.Atoi(page)

	if err_tp != nil || tp_i < 15 || tp_i > 17 {
		c.RenderingJsonAutomatically(2, "参数错误: type")
	} else if err_dup != nil {
		c.RenderingJsonAutomatically(2, "参数错误: page")
	} else {
		var slice []model.IntakeOutputDup
		var err error

		if tp_i == 17 {
			slice, err = model.QueryIntakeOrOutputVolumeAll(pid, page_i)
		} else {
			slice, err = model.QueryIntakeOrOutputVolume(pid, tp_i, page_i)
		}

		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			c.RenderingJson(0, "成功", slice)
		}
	}
}

func (c *IntakeOutputQueryController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *IntakeOutputQueryController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
