package handler

import (
	"fit"
	"nursing/model"
)

type QueryRecordController struct {
	fit.Controller
}

func (c QueryRecordController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	pid := r.FormValue("pid")
	nursType := r.FormValue("nurstype") //文书类型
	startTime := r.FormValue("start")   //开始时间
	endTime := r.FormValue("end")       //开始时间
	if "" == pid {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	var items []model.NursingRecords
	var err error
	if nursType != "" || startTime != "" || endTime != "" {
		items, err = model.QueryNRecordsByTypeAndTime(pid, nursType, startTime, endTime)
	}else
	{
		items, err = model.QueryNRecords(pid)
	}
	if err != nil {
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "查询出错"+err.Error()
		c.JsonData.Datas = []interface{}{}
		return
	}
	if nil == items {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "暂无数据"
		c.JsonData.Datas = []interface{}{}
		return
	}
	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "OK"
	c.JsonData.Datas = items
}
