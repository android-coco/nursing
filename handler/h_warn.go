package handler

import (
	"fit"
	"nursing/model"
	"time"
)

type WarnController struct {
	fit.Controller
}

func (c WarnController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	warnName := r.FormValue("warn_name")
	warnTimeStr := r.FormValue("warn_time")
	nurse_id := r.FormValue("nurse_id")
	patient_id := r.FormValue("patient_id")

	if warnName == "" || warnTimeStr == "" || nurse_id == "" || patient_id == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		c.ResponseToJson(w)
		return
	}

	warnTime, _ := time.Parse("2006-01-02 15:04:05", warnTimeStr)
	//warnTime = time.Now()
	//fit.Logger().LogInfo("", warnTime)

	warnModel := model.Warn{BaseModel: model.BaseModel{NurseId: nurse_id, PatientId: patient_id}, Name: warnName, WarnTime: warnTime}
	_, err := warnModel.InsertData()
	if err != nil {
		fit.Logger().LogError("Error", "warn add :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		//warnModel.Id = id
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{warnModel}
	}
	c.ResponseToJson(w)
}

type WarnListController struct {
	fit.Controller
}

func (c WarnListController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	nurse_id := r.FormValue("nurse_id")
	patient_id := r.FormValue("patient_id")
	if nurse_id == "" || patient_id == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		c.ResponseToJson(w)

	} else {
		c.JsonData.Result = 0
		warns := model.Warnlist(nurse_id, patient_id)
		c.JsonData.Datas = []interface{}{warns}
		c.ResponseToJson(w)
	}


}
