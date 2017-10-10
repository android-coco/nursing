package handler

import (
	"fit"
	"nursing/model"
)

type WarnController struct {
	fit.Controller
}

func (c WarnController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	warnName := r.FormValue("warn_name")
	warnDesc := r.FormValue("warn_desc")
	warnType, err1 := model.Warn{}.ParseWarnType(r.FormValue("warn_type"))
	warnTimeStr := r.FormValue("warn_time")
	classId := r.FormValue("class_id")

	if warnName == "" || warnDesc == "" || warnTimeStr == "" || err1 != nil || classId == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{err1.Error()}
		return
	}


	warnModel := model.Warn{
		Name:     warnName,
		Desc:     warnDesc,
		WarnType: warnType,
		ClassId:  classId,
		WarnTime: warnTimeStr}

	id, err := warnModel.InsertData()
	if err != nil || id == 0{
		fit.Logger().LogError("Error", "warn add :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{err.Error()}
	} else {
		//warnModel.Id = id
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = warnModel
	}
}

func (c WarnController) DelWarn(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	warnTime := r.FormValue("warn_time")
	classId := r.FormValue("class_id")

	if warnTime == "" || classId == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	warnModel := model.Warn{ClassId: classId, WarnTime:warnTime}
	id, err := warnModel.DeleteWarn()
	if err != nil {
		fit.Logger().LogError("Error", "warn delete :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "删除失败！"
		c.JsonData.Datas = []interface{}{err.Error()}
		return
	} else {
		if id == 0 {
			fit.Logger().LogError("Error", "warn delete :", err)
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "不存在该条记录！"
			c.JsonData.Datas = []interface{}{}
			return
		}
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "删除成功！"
		//c.JsonData.Datas = []interface{}{}
		return
	}
}

type WarnListController struct {
	fit.Controller
}

func (c WarnListController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	//nurseId := r.FormValue("nurse_id")
	classId := r.FormValue("class_id")
	listType := r.FormValue("type")

	if classId == "" || listType == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}

	} else {
		c.JsonData.Result = 0
		warns := model.Warnlist(classId, listType)
		c.JsonData.Datas = warns
	}

}
