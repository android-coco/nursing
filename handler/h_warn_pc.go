package handler

import (
	"fit"

	"fmt"
	"nursing/model"
	"strconv"
	"nursing/utils"
)

type PCWarnController struct {
	PCController
}
type WarnDatas struct {
	Warn model.Warn
	Type string //1,未执行，2,已经执行
}

func (c PCWarnController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		warns := model.WarnAll(strconv.Itoa(userinfo.DepartmentID))
		tcDetails := make([]WarnDatas, len(warns))
		for i := range warns {
			if utils.CompareTimeNow(warns[i].WarnTime+":00") {
				tcDetails[i].Type = "2"
			} else {
				tcDetails[i].Type = "1"
			}
			tcDetails[i].Warn = warns[i]
			fit.Logger().LogError("fasdf1111:", utils.CompareTimeNow(warns[i].WarnTime))
		}
		fit.Logger().LogError("fasdf", tcDetails)
		c.Data = fit.Data{
			"Userinfo": userinfo,
			"Warns":    tcDetails,
		}
		_ = c.LoadViewSafely(w, r, "pc/v_remind_admin.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}

}

func (c PCWarnController) DelWarn(w *fit.Response, r *fit.Request, p fit.Params) {
	wid := p.ByName("id")
	_, err := model.DelWarn(wid)
	if err != nil {
		fmt.Fprintln(w, "服务器有点繁忙！")
	} else {
		c.Redirect(w, r, "/pc/warn", 302)
	}

}

func (c PCWarnController) ModifyWarn(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	id := r.FormValue("id")
	name := r.FormValue("warn_name")
	desc := r.FormValue("warn_desc")
	time := r.FormValue("warn_time")
	warnType, err := model.Warn{}.ParseWarnType(r.FormValue("warn_type"))
	if id == "" || name == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	wid, err1 := strconv.Atoi(id)
	if err != nil || err1 != nil {
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "参数错误"
		c.JsonData.Datas = []interface{}{}
		return
	}
	warnModel := model.Warn{
		Id:       int64(wid),
		Name:     name,
		Desc:     desc,
		WarnType: warnType,
		WarnTime: time}
	_, err = warnModel.ModifyWarn()
	if err != nil || err1 != nil {
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "修改失败！" + err.Error()
		c.JsonData.Datas = []interface{}{}
		return
	}
	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "修改成功"
	c.JsonData.Datas = []interface{}{}
}
