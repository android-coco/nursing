package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
)

type PCAccessController struct {
	PCController
}

func (c PCAccessController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	access, err1 := model.AccessALLList(strconv.Itoa(userinfo.DepartmentID))
	if err == nil && err1 == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
			"Accesses": access,
		}
		_ = c.LoadViewSafely(w, r, "pc/v_access_manage.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

func (c PCAccessController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	userinfo, err := c.GetLocalUserinfo(w, r)
	classid := userinfo.DepartmentID
	paramstr := r.FormValue("paramstr")

	var access []model.Access

	if  paramstr == "" {
		access, err = model.AccessALLList(strconv.Itoa(userinfo.DepartmentID))
	}else{
		access, err = model.AccessSearch(strconv.Itoa(classid), paramstr)
	}
	if err == nil {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询成功"
		c.JsonData.Datas = access
		return
	} else {
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "查询出错"
		c.JsonData.Datas = err
		return
	}
}
