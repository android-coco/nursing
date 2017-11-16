//  Created by JP

package handler

import (
	"fit"
	"nursing/model"
	"nursing/utils"
)

type PCLoginController struct {
	fit.Controller
}

/*PC 登录页*/
func (c PCLoginController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	c.Data = fit.Data{
		"Baseurl": "http://127.0.0.1:8181/",
	}
	c.LoadView(w, "pc/v_login.html")
}

/*API 退出*/
func (c PCLoginController) Logout(w *fit.Response, r *fit.Request, p fit.Params) {
	fit.GlobalManager().SessionDestroy(w, r)
	//c.Redirect(w, r, "/pc/login", 302)
	c.RenderingJsonAutomatically(0, "已退出账号")
	c.ResponseToJson(w)
}

/*API 登录*/
func (c PCLoginController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	code := r.FormValue("code")
	password := r.FormValue("password")

	if code == "" || password == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	} else {
		password_sha1 := utils.Sha1Encryption(password)
		// 查询User表（多科室多条数据,前提是BCE1表支持多科室）
		slice_User, err_User := model.CheckingUserCodeAndPwd(code, password_sha1)
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

				// 将登录数据存入Session
				userinfo := model.UserInfoDup{
					UID:            user.Employeeid,
					Name:           user.Username,
					Password:       user.Password,
					Code:           user.Code,
					DepartmentID:   user.DepartmentID,
					DepartmentName: "",
					Authority:      authority,
					Status:         user.Status,
				}
				session, _ := fit.GlobalManager().SessionStart(w, r)
				if session != nil {
					session.Set("UserInfo", userinfo)
				}
				// navigate to 管理页面
				c.RenderingJson(0, "登录成功", responce)
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

			// 将登录数据存入Session
			userinfo := model.UserInfoDup{
				UID:            user.Employeeid,
				Name:           user.Username,
				Password:       user.Password,
				Code:           user.Code,
				DepartmentID:   responce.Departments[0].DepartmentID,
				DepartmentName: responce.Departments[0].DepartmentName,
				Authority:      user.Authority,
				Status:         user.Status,
			}
			session, _ := fit.GlobalManager().SessionStart(w, r)
			if session != nil {
				session.Set("UserInfo", userinfo)
			}

			if user.Status == 1 {
				c.RenderingJson(0, "登录成功", responce)
			} else {
				c.RenderingJson(6, "此账号已停用，请联系管理员", responce)
			}
		}
	}
}

func (c *PCLoginController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *PCLoginController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
