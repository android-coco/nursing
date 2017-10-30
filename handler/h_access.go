package handler

import (
	"fit"
	"nursing/model"
	"fmt"
)

type AccessController struct {
	fit.Controller
}

//type VAA1 struct {
//	VAA01 int    // 病人id
//	VAA05 string // 姓名
//	ABW01 string // 性别 0=未知，1=M=男，2=F=女，9=未说明
//	VAA10 int    // 年龄
//	BCQ04 string // 床号
//}

func (c AccessController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	accessType, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))
	accessReason, err2 := model.Access{}.ParseAccessReason(r.FormValue("access_reason"))
	accessTime := r.FormValue("access_time")
	nurseId := r.FormValue("nurse_id")
	nurseName := r.FormValue("nurse_name")
	patientId := r.FormValue("patient_id")
	patientName := r.FormValue("patient_name")
	bedId := r.FormValue("bed_id")
	classId := r.FormValue("class_id")

	fmt.Println("classid", classId)
	if nurseId == "" || patientId == "" || nurseName == "" || patientName == "" || bedId == "" || classId == "" || err1 != nil {
		if accessType == model.AccessTypeOut && err2 != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "参数错误"
			c.JsonData.Datas = []interface{}{}
			return
		}
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	accessModel := model.Access{
		BaseModel: model.BaseModel{
			NurseId:   nurseId,
			PatientId: patientId,
			NurseName:    nurseName},
		AccessType:   accessType,
		AccessReason: accessReason,
		AccessTime:   accessTime,
		PatientName:  patientName,
		BedId:        bedId,
		ClassId:      classId}

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

	classId := r.FormValue("class_id")
	page := r.FormValue("page")
	accessType, err1 := model.Access{}.ParseAccessType(r.FormValue("access_type"))

	if classId == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	} else if err1 != nil {
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "参数错误"
		c.JsonData.Datas = []interface{}{}
		return
	} else {
		accessModel, err := model.AccessList(classId, page, accessType)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "错误"
			c.JsonData.Datas = err
		} else {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = ""
			if accessModel == nil {
				c.JsonData.Datas = []interface{}{}
			} else {
				c.JsonData.Datas = accessModel
			}
		}
	}
}

type AccessSearchController struct {
	fit.Controller
}

func (c AccessSearchController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	classId := r.FormValue("class_id")
	paramstr := r.FormValue("paramstr")
	mods, err := model.AccessSearch(classId, paramstr)
	if classId == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	} else if err != nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "错误"
		c.JsonData.Datas = err
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = ""
		c.JsonData.Datas = mods
	}

}
