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

	if nurse_name == "" || nid == "" || pid == "" || recordtime == "" || tp == "" || subtype == "" || value == "" {
		c.RenderingJsonAutomatically(1, "参数不完整 null value")
		return
	}

	tp_i, err := strconv.Atoi(tp)
	if err != nil || tp_i < 1 || tp_i > 3 {
		c.RenderingJsonAutomatically(2, "参数错误: type")
		return
	} else {
		subtype_i, err_copy := strconv.Atoi(subtype)
		_, err_time := time.Parse("2006-01-02 15:04:05", recordtime)
		value_i, err_value := strconv.Atoi(value)
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

		iot := model.IntakeOutput{
			BaseModel: model.BaseModel{
				PatientId: pid,
				NurseId:   nid,
			},
			Type:          uint8(tp_i),
			Subtype:       uint8(subtype_i),
			OperationType: uint8(opertion_i),
			RecordTime:    recordtime,
			Value:         uint16(value_i),
			Desc:          desc,
			NurseName:     nurse_name,
		}

		err := iot.CollectIntakeOrOutputVolume()
		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			c.RenderingJsonAutomatically(0, "成功")
		}
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

func (c IntakeOutputQueryController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	tp := r.FormValue("type")
	page := r.FormValue("page")
	pid := r.FormValue("pid")
	if tp == "" || page == "" || pid == "" {
		c.RenderingJsonAutomatically(1, "参数不完整 null value")
		return
	}
	tp_i, err_tp := strconv.Atoi(tp)
	page_i, err_dup := strconv.Atoi(page)

	if err_tp != nil || tp_i < 1 || tp_i > 3 {
		c.RenderingJsonAutomatically(2, "参数错误: type")
	} else if err_dup != nil || page_i < 0 {
		c.RenderingJsonAutomatically(2, "参数错误: page")
	} else {
		var slice []model.IntakeOutputDup
		var err error

		if tp_i == model.IntakeOutputTypeIntake|model.IntakeOutputTypeOutput {
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
