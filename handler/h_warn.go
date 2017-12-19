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

	if warnName == "" || warnTimeStr == "" || err1 != nil || classId == "" {
		if err1 != nil {
			fit.Logger().LogError("Error", "warn add :", err1)
		}
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	warnModel := model.Warn{
		Name:     warnName,
		Desc:     warnDesc,
		WarnType: warnType,
		ClassId:  classId,
		WarnTime: warnTimeStr}

	id, err := warnModel.InsertData()
	if err != nil || id == 0 {
		fit.Logger().LogError("Error", "warn add :", err)
		if id == 0 {
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "不能创建同一时间的提醒！"
			c.JsonData.Datas = []interface{}{}
		} else {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "上传失败！"
			c.JsonData.Datas = []interface{}{}
		}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = warnModel
	}
}

func (c WarnController) DelWarn(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	//warnTime := r.FormValue("warn_time")
	//classId := r.FormValue("class_id")
	ids := r.FormValue("ids")

	if ids == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	//warnModel := model.Warn{ClassId: classId, WarnTime: warnTime}
	id, err := model.DelWarn(ids)
	if err != nil {
		fit.Logger().LogError("Error", "warn delete :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "删除失败！"
		c.JsonData.Datas = []interface{}{}
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


func (c WarnController) WarnList(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	classId := r.FormValue("class_id")

	if classId == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}

	} else {
		c.JsonData.Result = 0
		warns := model.Warnlist(classId)
		if warns == nil {
			c.JsonData.Datas = []interface{}{}
		} else {
			c.JsonData.Datas = warns
		}
	}

}
