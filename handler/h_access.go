package handler

import (
	"fit"
	"nursing/model"
)

type  AccessController struct {
	fit.Controller
}

type VAA1 struct {
	VAA01 int  // 病人id
	VAA05 string // 姓名
	ABW01 string // 性别 0=未知，1=M=男，2=F=女，9=未说明
	VAA10 int    // 年龄
	BCQ04 string // 床号
}

func (c AccessController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	access_type, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))
	access_reason, err2 := model.Access{}.ParseAccessReason(r.FormValue("access_reason"))
	//access_time, err := time.Parse("2006-01-02 15:04:05", r.FormValue("access_time"))
	access_time := r.FormValue("access_time")
	nurse_id := r.FormValue("nurse_id")
	patient_id := r.FormValue("patient_id")

	if nurse_id == "" || patient_id == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	} else if  err1 != nil || err2 != nil{
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "参数错误"
		c.JsonData.Datas = []interface{}{}
		return
	}

	accessModel := model.Access{
		BaseModel: model.BaseModel{NurseId: nurse_id, PatientId: patient_id},
		AccessType: access_type,
		AccessReason: access_reason,
		AccessTime:access_time}

	_, err := accessModel.InsertData()
	if err != nil {
		fit.Logger().LogError("Error", "warn add :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		//accessModel.Id = id
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = accessModel
	}
}

type AccessListController struct {
	fit.Controller
}

func (c AccessListController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	//access_type, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))
	//access_reason, err2 := model.Access{}.ParseAccessReason(r.FormValue("access_reason"))

	nurse_id := r.FormValue("nurse_id")
	patient_id := r.FormValue("patient_id")

	//var userInfos []VAA1
	//err := fit.SQLServerEngine().SQL("select VAA01, VAA05,ABW01 ,VAA10 ,BCQ04 from VAA1 WHERE VAA01 = ?", nurse_id).Find(&userInfos)
	//c.JsonData.Result = 1000
	//c.JsonData.ErrorMsg = "test"
	//c.JsonData.Datas = []interface{}{userInfos, err}
	//return



	if nurse_id == "" || patient_id == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	} else {
		accessModel, err := model.AccessList(nurse_id, patient_id)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "错误"
			c.JsonData.Datas = err
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = ""
			c.JsonData.Datas = accessModel
		}
	}
}

type AccessSearchController struct {
	fit.Controller
}

func (c AccessSearchController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	paramstr := r.FormValue("paramstr")
	mods, err := model.AccessSearch(paramstr)
	if err != nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "错误"
		c.JsonData.Datas = err
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = ""
		c.JsonData.Datas = mods
	}


}