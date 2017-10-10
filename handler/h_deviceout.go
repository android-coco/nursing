package handler

import (
	   "fit"
	"strconv"
	"nursing/model"
)

type DeviceoutController struct {
	fit.Controller
}

func (c DeviceoutController) Post(w *fit.Response, r *fit.Request, p fit.Params) {

	devicesclass,err := strconv.Atoi(r.FormValue("devicesclass"))
	//devicesname := r.FormValue("devicesname")

	if err!=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}else{

		items := []model.Devices{}
		err = fit.MySqlEngine().Where(" devicesclass = ?",devicesclass).Find(&items)

		if  err != nil {
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = []interface{}{err}
		}else{

			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询完成"
			c.JsonData.Datas = items
		}
	}

	c.ResponseToJson(w)

}