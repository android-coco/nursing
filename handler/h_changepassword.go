//  Created by JP

package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"nursing/utils"
)

type ChangePasswordController struct {
	PCController
}

func (c ChangePasswordController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}
		_ = c.LoadViewSafely(w, r, "pc/v_change_password.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

func (c ChangePasswordController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	orginal := r.FormValue("orginal")
	password := r.FormValue("password")
	verify := r.FormValue("verify")
	code := r.FormValue("code")

	if orginal == "" || password == "" || verify == "" || code == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	orginal_sha := utils.Sha1Encryption(orginal)
	password_sha := utils.Sha1Encryption(password)
	verify_sha := utils.Sha1Encryption(verify)

	if password_sha != verify_sha {
		c.RenderingJsonAutomatically(4, "新密码和确认密码不一致，请重新输入")
		return
	}

	userSlice, err_user := model.QueryUserTable(code, orginal_sha)
	if err_user != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_user.Error())
	} else if len(userSlice) == 0 {
		c.RenderingJsonAutomatically(5, "原始密码错误，请重新输入")
	} else {
		err_db := model.ChangePasswordWith(code, password_sha, password)
		if err_db != nil {
			c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
		} else {
			// 清楚session缓存
			fit.GlobalManager().SessionDestroy(w, r)
			c.RenderingJsonAutomatically(0, "修改成功，请重新登录")
		}
	}
}

func (c *ChangePasswordController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *ChangePasswordController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
