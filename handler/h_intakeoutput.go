package handler

import (
	"nursing/model"
	"fit"
	"strconv"
)

type IntakeOutputCollectController struct {
	fit.Controller
}
type IntakeOutputQueryController struct {
	fit.Controller
}

func (c IntakeOutputCollectController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	if length := len(r.Form); length < 7 {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nid := r.FormValue("nid")
	pid := r.FormValue("pid")
	recordtime := r.FormValue("recordTime")
	tp := r.FormValue("type")
	subtype := r.FormValue("subtype")
	value := r.FormValue("value")
	desc := r.FormValue("desc")

	if nid == "" || pid == "" || recordtime == "" || tp == "" || subtype == "" || value == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整 null value"
		c.JsonData.Datas = []interface{}{}
		return
	}

	tp_i, err := strconv.Atoi(tp)
	if err != nil || tp_i < 1 || tp_i > 3 {
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "参数错误: type"
		c.JsonData.Datas = []interface{}{}
		return
	} else {
		subtype_i, err_copy := strconv.Atoi(subtype)
		//time_dup, err_time := time.Parse("2006-01-02 15:04", recordtime)
		value_i, err_value := strconv.Atoi(value)
		if err_copy != nil || err_value != nil {
			c.JsonData.Result = 2
			c.JsonData.Datas = []interface{}{}
			if err_value != nil {
				c.JsonData.ErrorMsg = "参数错误: value"
			} else {
				c.JsonData.ErrorMsg = "参数错误: subtype"
			}
			return
		}

		iot := model.IntakeOutput{
			BaseModel: model.BaseModel{
				PatientId: pid,
				NurseId:   nid,
			},
			Type:       uint8(tp_i),
			Subtype:    uint8(subtype_i),
			RecordTime: recordtime,
			Value:      uint16(value_i),
			Desc:       desc,
		}

		err := iot.CollectIntakeOrOutputVolume()
		if err != nil {
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "Database " + err.Error()
			c.JsonData.Datas = []interface{}{}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "成功"
			c.JsonData.Datas = []interface{}{}
		}
	}
}

func (c IntakeOutputQueryController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	tp := r.FormValue("type")
	page := r.FormValue("page")
	pid := r.FormValue("pid")
	if tp == "" || page == "" || pid == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	} else {
		tp_i, err_tp := strconv.Atoi(tp)
		page_i, err_dup := strconv.Atoi(page)

		if err_tp != nil || tp_i < 1 || tp_i > 3 {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "参数错误: type"
			c.JsonData.Datas = []interface{}{}
		} else if err_dup != nil || page_i < 0 {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "参数错误: page"
			c.JsonData.Datas = []interface{}{}
		} else {
			var slice []interface{}
			var err error

			if tp_i == model.IntakeOutputTypeIntake|model.IntakeOutputTypeOutput {
				slice, err = model.QueryIntakeOrOutputVolumeAll(pid, page_i)
			} else {
				slice, err = model.QueryIntakeOrOutputVolume(pid, tp_i, page_i)
			}

			if err != nil {
				c.JsonData.Result = 3
				c.JsonData.Datas = []interface{}{}
				c.JsonData.ErrorMsg = err.Error()
			} else {
				c.JsonData.Result = 0
				c.JsonData.ErrorMsg = "成功"
				c.JsonData.Datas = slice
			}
		}
	}
}

func (c IntakeOutputQueryController) Post(w *fit.Response, r *fit.Request, p fit.Params)  {
	c.Get(w, r, p)
}