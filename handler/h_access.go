package handler

import (
	"fit"
	"time"
	"nursing/model"
)

type  AccessController struct {
	fit.Controller
}

func (c AccessController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	access_type, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))
	access_reason, err2 := model.Access{}.ParseAccessReason(r.FormValue("access_reason"))
	access_time, err := time.Parse("2006-01-02 15:04:05", r.FormValue("access_time"))
	nurse_id := r.FormValue("nurse_id")
	patient_id := r.FormValue("patient_id")

	if nurse_id == "" || patient_id == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		c.ResponseToJson(w)
		return
	} else if  err1 != nil || err2 != nil{
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "参数错误"
		c.JsonData.Datas = []interface{}{}
		c.ResponseToJson(w)
		return
	}

	accessModel := model.Access{
		BaseModel: model.BaseModel{NurseId: nurse_id, PatientId: patient_id},
		AccessType: access_type,
		AccessReason: access_reason,
		AccessTime:access_time}

	_, err = accessModel.InsertData()
	if err != nil {
		fit.Logger().LogError("Error", "warn add :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		//accessModel.Id = id
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{accessModel}
	}
	c.ResponseToJson(w)
}

type AccessListController struct {
	fit.Controller
}

func (c AccessListController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	//access_type, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))
	//access_reason, err2 := model.Access{}.ParseAccessReason(r.FormValue("access_reason"))
	//access_time, err := time.Parse("2006-01-02 15:04:05", r.FormValue("access_time"))
	nurse_id := r.FormValue("nurse_id")
	patient_id := r.FormValue("patient_id")

	if nurse_id == "" || patient_id == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		c.ResponseToJson(w)
	} else {
		accessModel := model.AccessList(nurse_id, patient_id)
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{accessModel}
		c.ResponseToJson(w)
	}
}

type AccessSearchController struct {
	fit.Controller
}

func (c AccessSearchController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	//access_type, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))
	//access_reason, err2 := model.Access{}.ParseAccessReason(r.FormValue("access_reason"))
	//access_time, err := time.Parse("2006-01-02 15:04:05", r.FormValue("access_time"))
	//nurse_id := r.FormValue("nurse_id")
	//patient_id := r.FormValue("patient_id")


}