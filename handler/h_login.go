package handler

import (
	"fit"
	"nursing/model"
)


/*登录页面*/
type LoginController struct {
	fit.Controller
}

func (c LoginController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	code := r.FormValue("code")
	password := r.FormValue("password")

	if code == "" || password == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
	} else {
		slice_User, err_User := model.QueryUserTable(code, password)
		if err_User != nil {
			c.RenderingJsonAutomatically(3, "Database "+err_User.Error())
		} else if length := len(slice_User); length == 0 {
			c.RenderingJsonAutomatically(2, "工号或者密码错误，请重新输入")
		} else if auth := slice_User[0].Authorized; auth != 0 {
			c.RenderingJsonAutomatically(4, "此账号未被授权")
		} else if eid := slice_User[0].EmployeeID; eid == 0 {
			c.RenderingJsonAutomatically(4, "无法访问员工数据库")
		} else {
			EID := slice_User[0].EmployeeID
			slice_IAN, err_IAN := model.QueryEmployeeContrastTable(EID)
			if err_IAN != nil {
				c.RenderingJsonAutomatically(3, "Database "+err_IAN.Error())
			} else if length := len(slice_IAN); length == 0 {
				c.RenderingJsonAutomatically(2, "工号不存在，请检查工号")
			} else {
				UID := slice_IAN[0].BCE01
				slice_BCE, err_BCE := model.QueryEmployeeTable(UID, code)
				if err_BCE != nil {
					c.RenderingJsonAutomatically(3, "Database "+err_BCE.Error())
				} else if length := len(slice_BCE); length == 0 {
					c.RenderingJsonAutomatically(2, "账号不存在，请检查工号")
				} else {
					len_BCE := len(slice_BCE)
					respoce := model.User_Response{
						UID:         UID,
						Name:        slice_User[0].Name,
						Password:    password,
						Code:        code,
						Departments: make([]model.Department, len_BCE),
					}
					for i, BCE := range slice_BCE {
						DID := BCE.BCK01
						slice_BCK, err_BCK := model.QueryDepartmentTable(DID)
						if err_BCK != nil {
							c.RenderingJsonAutomatically(3, "Database "+err_BCK.Error())
							return
						} else if length := len(slice_BCK); length == 0 {
							c.RenderingJsonAutomatically(5, "此账号未隶属任何科室")
							return
						} else {
							respoce.Departments[i].DepartmentID = DID
							respoce.Departments[i].DepartmentName = slice_BCK[0].BCK03
						}
					}
					c.RenderingJson(0, "登录成功", respoce)
				}
			}
		}
	}
}

func (c *LoginController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *LoginController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
