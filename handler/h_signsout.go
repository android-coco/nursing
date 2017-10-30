package handler

import (
	"fit"
	"nursing/model"
	"time"

)

type SignsoutController struct {
	fit.Controller
}

func (c SignsoutController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	fit.Logger().LogAssert("", "post")
	defer c.ResponseToJson(w)

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
		sql = sql + "testtime > ?  "
		msg = append(msg, starttime)
	}
	if err9 == nil {
		if err8 == nil {
			sql = sql + " and "
		}
		sql = sql + "testtime < ?"
		msg = append(msg, endtime)
	}

	items := make(map[string]interface{})
	temperature := r.FormValue(model.Temperature_Type )
	if len(temperature) != 0{
		item, err := model.OutTemperature(sql, msg...)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err
			return
		} else {
			items[model.Temperature_Type] = item
		}
	}

	pulse := r.FormValue(model.Pulse_Type )
	if len(pulse) != 0{
		item1, err1 := model.OutPulse(sql, msg...)
		if err1 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误1"
			c.JsonData.Datas = err1
			return
		} else {
			items[model.Pulse_Type] = item1
		}
	}

	breathe := r.FormValue(model.Breathe_Type )
	if len(breathe) != 0{
		item2, err2 := model.OutBreathe(sql, msg...)
		if err2 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误2"
			c.JsonData.Datas = err2
			return
		} else {
			items[model.Breathe_Type] = item2
		}
	}

	pressure := r.FormValue(model.Pressure_Type )
	if len(pressure) != 0{
		item3, err3 := model.OutPressure(sql, msg...)
		if err3 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误3"
			c.JsonData.Datas = err3
			return
		} else {
			items[model.Pressure_Type] = item3
		}
	}

	heartrate := r.FormValue(model.Heartrate_Type )
	if len(heartrate) != 0{
		item4, err4 := model.OutHeartrate(sql, msg...)
		if err4 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误4"
			c.JsonData.Datas = err4
			return
		} else {
			items[model.Heartrate_Type] = item4
		}
	}

	spo2h := r.FormValue(model.Spo2h_Type )
	if len(spo2h) != 0{
		item5, err5 := model.OutSpo2h(sql, msg...)
		if err5 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误5"
			c.JsonData.Datas = err5
			return
		} else {
			items[model.Spo2h_Type] = item5
		}
	}

	glucose := r.FormValue(model.Glucose_Type )
	if len(glucose) != 0{
		item6, err6 := model.OutGlucose(sql, msg...)
		if err6 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误6"
			c.JsonData.Datas = err6
			return
		} else {
			items[model.Glucose_Type] = item6
		}
	}

	weight := r.FormValue(model.Weight_Type )
	if len(weight) != 0{
		item7, err7 := model.OutWeight(sql, msg...)
		if err7 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误7"
			c.JsonData.Datas = err7
			return
		} else {
			items[model.Weight_Type] = item7
		}
	}

	height := r.FormValue(model.Height_Type )
	if len(height) != 0{
		item8, err8 := model.OutHeight(sql, msg...)
		if err8 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误8"
			c.JsonData.Datas = err8
			return
		} else {
			items[model.Height_Type] = item8
		}
	}

	skin := r.FormValue(model.Skin_Type )
	if len(skin) != 0{
		item9, err9 := model.OutSkin(sql, msg...)
		if err9 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误9"
			c.JsonData.Datas = err9
			return
		} else {
			items[model.Skin_Type] = item9
		}
	}

	ache := r.FormValue(model.Ache_Type )
	if len(ache) != 0{
		item10, err10 := model.OutAche(sql, msg...)
		if err10 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误10"
			c.JsonData.Datas = err10
			return
		} else {
			items[model.Ache_Type] = item10
		}
	}

	incident := r.FormValue(model.Incident_Type )
	if len(incident) != 0{
		item11, err11 := model.OutIncident(sql, msg...)
		if err11 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误11"
			c.JsonData.Datas = err11
			return
		} else {
			items[model.Incident_Type] = item11
		}
	}

	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "查询完成"
	c.JsonData.Datas = items

}
