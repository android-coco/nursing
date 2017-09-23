package handler

import (
	"fit"
	"nursing/model"
	"strconv"
	"time"
)

type SignsiputController struct {
	fit.Controller
}

func (c SignsiputController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
}

func (c SignsiputController) Post(w *fit.Response, r *fit.Request, p fit.Params) {

	fit.Logger().LogAssert("", "post")

	measuretype, err := strconv.Atoi(r.FormValue("measuretype"))
	if err != nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	switch measuretype {
	case model.Temperature_Type:
		c.iputTemperature(r)
	case model.Pulse_Type:
		c.iputPulse(r)
	case model.Breathe_Type:
		c.iputBreathe(r)
	case model.Pressure_Type:
		c.iputPressure(r)
	case model.Heartrate_Type:
		c.iputHeartrate(r)
	case model.Spo2h_Type:
		c.iputSpo2h(r)
	case model.Glucose_Type:
		c.iputGlucose(r)
	case model.Weight_Type:
		c.iputWeight(r)
	case model.Height_Type:
		c.iputHeight(r)
	}

	c.ResponseToJson(w)
}

func (c *SignsiputController) iputTemperature(r *fit.Request) {

	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err4 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	temptype, err5 := strconv.ParseUint(r.FormValue("temptype"), 10, 16)
	coolingtype, err6 := strconv.ParseUint(r.FormValue("coolingtype"), 10, 16)
	value, err7 := strconv.ParseFloat(r.FormValue("v1"), 32)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err4 == nil && err5 == nil && err6 == nil && err7 == nil {

		var item model.Temperature
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Ttemptype = uint16(temptype)
		item.Coolingtype = uint16(coolingtype)
		item.Value = float32(value)
		item.Testtime = fit.JsonTime(texttime)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetTemperature")
}

func (c *SignsiputController) iputPulse(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	value, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && err2 == nil {

		var item model.Pulse
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Value = uint16(value)
		item.Testtime = fit.JsonTime(texttime)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetPulse")
}

func (c *SignsiputController) iputBreathe(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	value, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)
	whethertbm := r.FormValue("whethertbm")

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && len(whethertbm) == 0 && err2 == nil {

		var item model.Breathe
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Value = uint16(value)
		item.Testtime = fit.JsonTime(texttime)

		if whethertbm == "true" {
			item.Whethertbm = true
		} else {
			item.Whethertbm = false
		}

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetBreathe")
}

func (c *SignsiputController) iputPressure(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err5 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)
	v2, err3 := strconv.ParseUint(r.FormValue("v2"), 10, 16)
	v3, err4 := strconv.ParseUint(r.FormValue("v3"), 10, 16)

	if len(patientid) != 0 && len(nurseid) != 0 && err5 == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil {

		var item model.Pressure
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Sysvalue = uint16(v1)
		item.Diavalue = uint16(v2)
		item.Pulsevalue = uint16(v3)
		item.Testtime = fit.JsonTime(texttime)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetPressure")
}

func (c *SignsiputController) iputHeartrate(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && err2 == nil {

		var item model.Heartrate
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Testtime = fit.JsonTime(texttime)
		item.Value = uint16(v1)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetHeartrate")
}

func (c *SignsiputController) iputSpo2h(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseUint(r.FormValue("v1"), 10, 16)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && err2 == nil {

		var item model.Spo2h
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Testtime = fit.JsonTime(texttime)
		item.Value = uint16(v1)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetSpo2h")
}

func (c *SignsiputController) iputGlucose(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseFloat(r.FormValue("v1"), 32)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && err2 == nil {

		var item model.Glucose
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Testtime = fit.JsonTime(texttime)
		item.Value = float32(v1)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetGlucose")
}

func (c *SignsiputController) iputWeight(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseFloat(r.FormValue("v1"), 32)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && err2 == nil {

		var item model.Weight
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Testtime = fit.JsonTime(texttime)
		item.Recordscene = uint16(recordscene)
		item.Value = float32(v1)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", r.FormValue("Time"))
}

func (c *SignsiputController) iputHeight(r *fit.Request) {
	patientid := r.FormValue("patientid")
	nurseid := r.FormValue("nurseid")
	texttime, err3 := time.Parse("2006-01-02 15:04:05", r.FormValue("texttime"))
	recordscene, err1 := strconv.ParseUint(r.FormValue("recordscene"), 10, 16)
	v1, err2 := strconv.ParseFloat(r.FormValue("v1"), 64)

	if len(patientid) != 0 && len(nurseid) != 0 && err3 == nil && err1 == nil && err2 == nil {

		var item model.Height
		item.PatientId = patientid
		item.NurseId = nurseid

		item.Recordscene = uint16(recordscene)
		item.Testtime = fit.JsonTime(texttime)
		item.Value = float64(v1)

		err := item.InsertData(item);
		if err != nil {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "录入失败"
			c.JsonData.Datas = []interface{}{err}
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "录入成功"
			c.JsonData.Datas = []interface{}{item}
		}
	} else {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}

	fit.Logger().LogDebug("", "GetHeight")
}
