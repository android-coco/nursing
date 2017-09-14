package handler

import (
	"fit"
)

type LoginController struct {
	fit.Controller
}

//json实体struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c LoginController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	} else {
		if username == "admin" && password == "123456" {
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "登录成功"
			c.JsonData.Datas = []interface{}{User{Username: username, Password: password}}
		} else {
			c.JsonData.Result = 2
			c.JsonData.ErrorMsg = "用户名或密码错误！"
			c.JsonData.Datas = []interface{}{}
		}
	}
	c.ResponseToJson(w)
}
