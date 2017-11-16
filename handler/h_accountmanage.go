//  Created by JP

package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"encoding/json"
)

type AccountManageController struct {
	PCController
}

/*账号管理页面（创建账号）*/
func (c AccountManageController) Manage(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}

		// 取当前科室
		bck01 := userinfo.DepartmentID
		if userinfo.Authority > 1 {
			// 获取科室列表
			departments, err_dep := model.QueryDepartmentList(true) //科室信息
			if err_dep != nil {
				fmt.Fprintln(w, "接口好像出了一点问题 err_dep:"+err_dep.Error())
			} else {
				// 默认取第一个科室
				bck01 = departments[0].BCK01
				c.Data["Departments"] = departments
			}
		}

		users, err_user := model.FetchAllOfTheAccountNotBeenCreated(bck01)
		if err_user != nil {
			fmt.Fprintln(w, "接口好像出了一点问题 err_user:"+err_user.Error())
		} else {
			c.Data["Users"] = users
		}

		_ = c.LoadViewSafely(w, r, "pc/v_account_manage_create.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*账号管理页面（已创建）*/
func (c AccountManageController) List(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}

		// 取当前科室
		bck01 := userinfo.DepartmentID
		if userinfo.Authority > 1 {
			// 获取科室列表
			departments, err_dep := model.QueryDepartmentList(true) //科室信息
			if err_dep != nil {
				fmt.Fprintln(w, "接口好像出了一点问题 err_dep:"+err_dep.Error())
			} else {
				// 默认取第一个科室
				bck01 = departments[0].BCK01
				c.Data["Departments"] = departments
			}
		}

		uid := int(userinfo.UID)
		users, err_user := model.FetchAllOfTheAccountHasCreated(bck01, 0, 0, uid)
		if err_user != nil {
			fmt.Fprintln(w, "接口好像出了一点问题 err_user:"+err_user.Error())
		} else {
			c.Data["Users"] = users
		}

		_ = c.LoadViewSafely(w, r, "pc/v_account_manage_update.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*API 创建账号*/
func (c AccountManageController) Create(w *fit.Response, r *fit.Request, p fit.Params) {
	// 新建账号
	defer c.ResponseToJson(w)
	r.ParseForm()
	handler := r.FormValue("handler")
	datas := r.FormValue("datas")

	if handler == "" || datas == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	handlerInfo, err_h := model.FetchAccountWithCode(handler)
	if err_h != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_h.Error())
		return
	} else if handlerInfo.Code != handler {
		c.RenderingJsonAutomatically(2, "参数错误： handler")
		return
	}

	by := []byte(datas)
	anyObj := make([]map[string]string, 0)
	err_json := json.Unmarshal(by, &anyObj)
	if err_json != nil {
		fit.Logger().LogError("JP", err_json)
		c.RenderingJsonAutomatically(2, "参数错误 datas")
		return
	}

	failure := make([]map[string]string, 0)
	for _, obj := range anyObj {
		code, ok := obj["code"]
		if !ok || "" == code {
			obj["desc"] = "参数错误"
			failure = append(failure, obj)
			continue
		}
		userinfo, _ := model.FetchAccountWithCode(code)
		// 过滤掉已创建的账号
		if userinfo.Code == code {
			obj["desc"] = "已创建"
			failure = append(failure, obj)
			continue
		}
		status, okk := obj["status"]
		authority, okkk := obj["authority"]

		if !okk || !okkk {
			obj["desc"] = "参数错误"
			failure = append(failure, obj)
			continue
		}

		// 验证权限
		authority_i, err_auth := strconv.Atoi(authority)
		status_i, err_sta := strconv.Atoi(status)
		if err_auth != nil || err_sta != nil {
			obj["desc"] = "参数错误"
			failure = append(failure, obj)
			continue
		} else if handlerInfo.Authority <= authority_i {
			obj["desc"] = "权限不足"
			failure = append(failure, obj)
			continue
		}

		// 从his获取账号信息
		bce, err_bce := model.FetchAccountFromHis(code)
		if err_bce != nil {
			obj["desc"] = "原始账号不存在"
			failure = append(failure, obj)
			continue
		} else if handlerInfo.Authority == 1 && handlerInfo.DepartmentID != bce.BCK01 {
			obj["desc"] = "权账号非护士长账号所属科室"
			failure = append(failure, obj)
			continue
		}
		// 创建账号
		err_db := model.CreateAccountWithOriginalUserinfo(bce, authority_i, status_i)
		if err_db != nil {
			obj["desc"] = "创建失败:" + err_db.Error()
			failure = append(failure, obj)
			continue
		}
	}

	length := len(anyObj) - len(failure)
	c.RenderingJson(0, "创建成功："+fmt.Sprintf("%d", length)+"，创建失败："+fmt.Sprintf("%d", len(failure)), failure)
}

/*API 更改账号状态*/
func (c AccountManageController) Update(w *fit.Response, r *fit.Request, p fit.Params) {

	// 修改账号的权限和状态
	defer c.ResponseToJson(w)
	r.ParseForm()
	handler := r.FormValue("handler")
	code := r.FormValue("code")
	status := r.FormValue("status")
	authority := r.FormValue("authority")

	if handler == "" || code == "" || status == "" || authority == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	authority_i, err_auth := strconv.Atoi(authority)
	if err_auth != nil {
		c.RenderingJsonAutomatically(2, "参数错误 authority")
		return
	}
	status_i, err_sta := strconv.Atoi(status)
	if err_sta != nil {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	userinfo, err_u := model.FetchAccountWithCode(code)
	if err_u != nil {
		c.RenderingJsonAutomatically(3, "Database_1 "+err_u.Error())
		return
		//	userinfo.Employeeid == 0 也可能是管理员账号，需要过滤掉
	} else if userinfo.Employeeid == 0 && userinfo.Code != code {
		c.RenderingJsonAutomatically(4, "无法查询到此工号员工")
		return
	}

	handlerInfo, err_h := model.FetchAccountWithCode(handler)
	if err_h != nil {
		c.RenderingJsonAutomatically(3, "Database_2 "+err_h.Error())
		return
	} else if handlerInfo.Code != handler {
		c.RenderingJsonAutomatically(2, "参数错误： handler")
		return
	} else if handlerInfo.Authority <= authority_i {
		c.RenderingJsonAutomatically(5, "操作被禁止，没有足够的权限")
		return
	}

	if userinfo.Authority == authority_i && userinfo.Status == status_i {
		c.RenderingJsonAutomatically(0, "操作成功")
	} else {
		userinfo.Authority = authority_i
		userinfo.Status = status_i
		err_db := userinfo.UpdateAccountAuthorityAndStatus()
		if err_db != nil {
			c.RenderingJsonAutomatically(3, "Database_3 "+err_db.Error())
		} else {
			c.RenderingJsonAutomatically(0, "操作成功")
		}
	}
}

/*API 获取未创建的账号*/
func (c AccountManageController) Uncreated(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	departmentId := r.FormValue("did")
	if departmentId == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	did, err_did := strconv.Atoi(departmentId)
	if err_did != nil {
		c.RenderingJsonAutomatically(2, "参数错误"+err_did.Error())
		return
	}

	response, err_res := model.FetchAllOfTheAccountNotBeenCreated(did)
	if err_res != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_res.Error())
	} else {
		c.RenderingJson(0, "请求成功", response)
	}
}

/*API 获取已创建的账号*/
func (c AccountManageController) Created(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	departmentId := r.FormValue("did")
	authority := r.FormValue("authority")
	status := r.FormValue("status")

	if departmentId == "" || authority == "" || status == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	did, err_did := strconv.Atoi(departmentId)

	if err_did != nil {
		c.RenderingJsonAutomatically(2, "参数错误 department_id")
		return
	}
	authority_i, err_auth := strconv.Atoi(authority)
	if err_auth != nil || authority_i < 0 || authority_i > 2 {
		c.RenderingJsonAutomatically(2, "参数错误 authority")
		return
	}
	status_i, err_sta := strconv.Atoi(status)
	if err_sta != nil || status_i < 0 || status_i > 2 {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	userinfo, err_u := c.GetLocalUserinfo(w, r)
	if err_u == nil {
		uid := int(userinfo.UID)
		response, err_res := model.FetchAllOfTheAccountHasCreated(did, authority_i, status_i, uid)
		if err_res != nil {
			c.RenderingJsonAutomatically(3, "Database "+err_res.Error())
		} else {
			c.RenderingJson(0, "请求成功", response)
		}
	}
}

func (c *AccountManageController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *AccountManageController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
