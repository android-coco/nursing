package handler

import (
	"fit"
	"strconv"
	"nursing/model"
	"time"

)

type SignsoutController struct {
	fit.Controller
}

func (c SignsoutController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	fit.Logger().LogAssert("", "get")
}

func (c SignsoutController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	fit.Logger().LogAssert("", "post")
	//sql := "WHERE "

	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var sql string
	var msg []interface{}

	if len(patientid) != 0 {
		sql = sql + "patientid = ?"
		msg = append(msg, patientid)
	}
	if len(nurseid) != 0 {
		if len(patientid) != 0 {
			sql = sql + " and "
		}
		sql = sql + "nurseid = ?"
		msg = append(msg, nurseid)
	}
	if err8 == nil {
		if len(nurseid) != 0 {
			sql = sql + " and "
		}
		sql = sql + "starttime > ?  "
		msg = append(msg, starttime)
	}
	if err9 == nil {
		if err8 == nil {
			sql = sql + " and "
		}
		sql = sql + "endtime > ?"
		msg = append(msg, endtime)
	}

	measuretype := r.FormValue("measuretype")
	if len(measuretype) == 0 {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	} else {
		items, err := model.OutTemperature(sql, msg...)
		if err != nil {
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "查询错误" + err.Error()
			c.JsonData.Datas = []interface{}{}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询完成"
			c.JsonData.Datas = items
		}
	}

	/*measuretype := r.FormValue("measuretype")
	if len(measuretype)==0 {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	switch measuretype {
	case model.Temperature_Type :
		c.outTemperature(r)
	case model.Pulse_Type :
		c.outPulse(r)
	case model.Breathe_Type :
		c.outBreathe(r)
	case model.Pressure_Type :
		c.outPulse(r)
	case model.Heartrate_Type :
		c.outHeartrate(r)
	case model.Spo2h_Type:
		c.outSpo2h(r)
	case model.Glucose_Type:
		c.outGlucose(r)
	case model.Weight_Type:
		c.outWeight(r)
	case model.Height_Type:
		c.outHeight(r)
	}*/

	c.ResponseToJson(w)
}

func (c *SignsoutController) outTemperature(r *fit.Request) {

	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err4 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	ttemptype, err5 := strconv.ParseUint(r.FormValue("ttemptype"), 10, 16)
	coolingtype, err6 := strconv.ParseUint(r.FormValue("coolingtype"), 10, 16)
	v1, err7 := strconv.ParseFloat(r.FormValue("v1"), 32)

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Temperature

	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err4 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err5 == nil {
		item.Ttemptype = uint16(ttemptype)
	}
	if err6 == nil {
		item.Coolingtype = uint16(coolingtype)
	}
	if err7 == nil {
		item.Value = float32(v1)
	}

	var err error
	items := make([]model.Temperature, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}

func (c *SignsoutController) outPulse(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Pulse;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = uint16(v1)
	}

	var err error
	items := make([]model.Pulse, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}

func (c *SignsoutController) outBreathe(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)
	whethertbm := r.FormValue("whethertbm")

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Breathe;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = uint16(v1)
	}
	if len(whethertbm) != 0 {
		if whethertbm == "true" {
			item.Whethertbm = true
		} else {
			item.Whethertbm = false
		}
	}

	var err error
	items := make([]model.Breathe, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)

}

func (c *SignsoutController) outPressure(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err5 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)
	v2, err3 := strconv.ParseUint(r.FormValue("v2"), 10, 16)
	v3, err4 := strconv.ParseUint(r.FormValue("v3"), 10, 16)

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Pressure;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err5 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Sysvalue = uint16(v1)
	}
	if err3 == nil {
		item.Diavalue = uint16(v2)
	}
	if err4 == nil {
		item.Pulsevalue = uint16(v3)
	}

	var err error
	items := make([]model.Pressure, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}

func (c *SignsoutController) outHeartrate(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Heartrate;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = uint16(v1)
	}

	var err error
	items := make([]model.Heartrate, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}

func (c *SignsoutController) outSpo2h(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Heartrate;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = uint16(v1)
	}

	var err error
	items := make([]model.Spo2h, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}

func (c *SignsoutController) outGlucose(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseFloat(r.FormValue("v1"), 32)

	starttime, err8 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err9 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Heartrate;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = uint16(v1)
	}

	var err error
	items := make([]model.Glucose, 0)
	if err8 == nil && err9 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err8 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err9 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}

func (c *SignsoutController) outWeight(r *fit.Request) {

	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseFloat(r.FormValue("v1"), 32)

	starttime, err4 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err5 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Weight;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = float32(v1)
	}

	var err error
	items := make([]model.Weight, 0)
	if err4 == nil && err5 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err4 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err5 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)

}

func (c *SignsoutController) outHeight(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseFloat(r.FormValue("v1"), 64)

	starttime, err4 := time.Parse("2006-01-02 15:04:05", r.FormValue("starttime"))
	endtime, err5 := time.Parse("2006-01-02 15:04:05", r.FormValue("endtime"))

	var item model.Height;
	if len(patientid) != 0 {
		item.PatientId = patientid
	}
	if len(nurseid) != 0 {
		item.NurseId = nurseid
	}
	if err3 == nil {
		item.Testtime = fit.JsonTime(texttime)
	}
	if err1 == nil {
		item.Recordscene = uint16(recordscene)
	}
	if err2 == nil {
		item.Value = float64(v1)
	}

	var err error
	items := make([]model.Height, 0)
	if err4 == nil && err5 == nil {
		err = fit.MySqlEngine().Where("Testtime > ? and Testtime < ?", starttime, endtime).Find(&items, item)
	} else if err4 == nil {
		err = fit.MySqlEngine().Where("Testtime > ?", starttime).Find(&items, item)
	} else if err5 == nil {
		err = fit.MySqlEngine().Where("Testtime < ?", endtime).Find(&items, item)
	} else {
		err = fit.MySqlEngine().Find(&items, item)
	}

	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "查询错误"
		c.JsonData.Datas = []interface{}{err}
	} else {

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询完成"
		c.JsonData.Datas = items
	}

	fit.Logger().LogVerbose("", items)
}
