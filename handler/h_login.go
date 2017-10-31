//  Created by JP

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
		return
	} else {
		// 查询User表（多科室多条数据,前提是BCE1表支持多科室）
		slice_User, err_User := model.QueryUserTable(code, password)
		length := len(slice_User)
		if err_User != nil {
			c.RenderingJsonAutomatically(3, "Database "+err_User.Error())
			return
		} else if length == 0 {
			c.RenderingJsonAutomatically(2, "工号或者密码错误，请重新输入")
			return
		}

		user := slice_User[0]
		if user.Employeeid == 0 {
			// Employeeid == 0 有可能是管理员账号，也可能未授权的普通账号
			if authority := user.Authority; authority != 0 {
				responce := model.UserInfo{
					UID:         user.Employeeid,
					Name:        user.Username,
					Password:    user.Password,
					Code:        user.Code,
					Departments: make([]model.Department, 0),
					Authority:   authority,
					Status:      user.Status,
				}

				if user.Status == 1 {
					if authority == 3 {
						// 设备管理员
						c.RenderingJson(0, "登录成功", responce)
					} else if authority == 2 {
						// 账号管理员
						c.RenderingJson(4, "此账号禁止在PDA端使用", responce)
					}
				} else {
					c.RenderingJson(6, "此账号已停用，请联系管理员", responce)
				}
				return
			} else {
				c.RenderingJsonAutomatically(4, "未被授权，无法访问员工数据")
				return
			}
		} else if user.DepartmentID == 0 {
			c.RenderingJsonAutomatically(5, "此账号未隶属任何科室")
			return
		} else {
			responce := model.UserInfo{
				UID:         user.Employeeid,
				Name:        user.Username,
				Password:    user.Password,
				Code:        user.Code,
				Departments: make([]model.Department, length),
				Authority:   user.Authority,
				Status:      user.Status,
			}
			for i, obj := range slice_User {
				department, err_BCK := model.QueryDepartmentWithDID(obj.DepartmentID)
				if err_BCK != nil {
					c.RenderingJsonAutomatically(3, "Database "+err_BCK.Error())
					return
				} else {
					responce.Departments[i].DepartmentID = department.BCK01
					responce.Departments[i].DepartmentName = department.BCK03
				}
			}

			if user.Status == 1 {
				c.RenderingJson(0, "登录成功", responce)
			} else {
				c.RenderingJson(6, "此账号已停用，请联系管理员", responce)
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
