package handler

import (
	"fit"
	"strconv"
	"nursing/model"
)

type DeviceiputController struct {
	fit.Controller
}

func (c DeviceiputController) Post(w *fit.Response, r *fit.Request, p fit.Params) {

	actiontype := r.FormValue("actiontype")
    if len(actiontype) == 0{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}else{
		if actiontype == "0"{
			c.addDevices(r)
		}else{
			c.amendDevices(r)
		}
	}

	c.ResponseToJson(w)

	fit.Logger().LogDebug("","Post")
}

func (c *DeviceiputController) addDevices(r *fit.Request) {
	devicesclass,err := strconv.Atoi(r.FormValue("devicesclass"))
	devicesname := r.FormValue("devicesname")
	devicelist  := r.FormValue("devicelist")

	if err != nil || len(devicesname) ==0 || len(devicelist) == 0 {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}else{
		var devices model.Devices

		has, err := fit.MySqlEngine().Where("devicesclass = ? and devicesname = ?",devicesclass,devicesname).Get(&devices)
		if has == false || err!=nil {
			devices.Devicesclass = uint16(devicesclass)
			devices.Devicesname = devicesname
			devices.Devicelist = devicelist

			_, err := fit.MySqlEngine().Insert(&devices)
			if err!= nil{
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "插入失败"
				c.JsonData.Datas = []interface{}{err}
			}else{
				c.JsonData.Result = 0
				c.JsonData.ErrorMsg = "插入成功"
				c.JsonData.Datas = []interface{}{}
			}
		}else{
			c.JsonData.Result = 3
			c.JsonData.ErrorMsg = "套餐已存在"
			c.JsonData.Datas = []interface{}{err}
		}

	}
}

func (c *DeviceiputController) amendDevices(r *fit.Request) {
	//devicesclass,err := strconv.Atoi(r.FormValue("devicesclass"))
	id,err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	altername := r.FormValue("devicesname")
	alterlist  := r.FormValue("devicelist")

	if  err!= nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
	}
	var devices model.Devices
	//devices.Devicesclass = uint16(devicesclass)
	//devices.Id = id
	//.Where("id = ?",id)
	has, err1 := fit.MySqlEngine().Where("id = ?",id).Get(&devices)

	//fit.Logger().LogDebug("aaa",devices)
	if has == false || err1!=nil {
		c.JsonData.Result = 4
		c.JsonData.ErrorMsg = "套餐不存在"
		c.JsonData.Datas = []interface{}{}
	}else{
		if len(altername) ==0 && len(alterlist) ==0{
			_,err2 := fit.MySqlEngine().Where("id = ?",id).Delete(&devices)
			if err2!= nil{
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "删除失败"
				c.JsonData.Datas = []interface{}{err2}
			}else{
				c.JsonData.Result = 0
				c.JsonData.ErrorMsg = "删除成功"
				c.JsonData.Datas = []interface{}{}
			}
		}else if len(altername) !=0 && len(alterlist) !=0{
			if devices.Devicesname != altername{

				has, err := fit.MySqlEngine().Where("devicesclass = ? and devicesname = ?",devices.Devicesclass,altername).Get(&model.Devices{})
				if has == false && err == nil{
					devices.Devicelist = alterlist
					devices.Devicesname = altername

					_,err3 := fit.MySqlEngine().Where("id = ? ",id).Update(&devices)
					if err3!= nil{
						c.JsonData.Result = 2
						c.JsonData.ErrorMsg = "更新失败"
						c.JsonData.Datas = []interface{}{err3}
					}else{
						c.JsonData.Result = 0
						c.JsonData.ErrorMsg =  "更新成功"
						c.JsonData.Datas = []interface{}{}
					}
				}else{
					if err == nil{
						c.JsonData.Result = 3
						c.JsonData.ErrorMsg = "套餐名字冲突"
						c.JsonData.Datas = []interface{}{}
					}else{
						c.JsonData.Result = 3
						c.JsonData.ErrorMsg = "数据库请求出错"
						c.JsonData.Datas = []interface{}{err}
					}
				}
			}else{
				devices.Devicelist = alterlist
				_,err3 := fit.MySqlEngine().Where("id = ? ",id).Update(&devices)
				if err3!= nil{
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "更新失败"
					c.JsonData.Datas = []interface{}{err3}
				}else{
					c.JsonData.Result = 0
					c.JsonData.ErrorMsg =  "更新成功"
					c.JsonData.Datas = []interface{}{}
				}
			}
		}else{
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "修改参数不完整"
			c.JsonData.Datas = []interface{}{}
		}


		/*else if len(altername) !=0{
			devices.Devicesname = altername
			_,err4 :=fit.MySqlEngine().Where("id = ?",id).Update(&devices)

			if err4!= nil{
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "更新失败"
				c.JsonData.Datas = []interface{}{err4}
			}else{
				c.JsonData.Result = 0
				c.JsonData.ErrorMsg = "更新成功"
				c.JsonData.Datas = []interface{}{}
			}
		}else {
			devices.Devicelist = alterlist
			_,err5 :=fit.MySqlEngine().Where("id = ?",id).Update(&devices)

			if err5!= nil{
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "更新失败"
				c.JsonData.Datas = []interface{}{err5}
			}else{
				c.JsonData.Result = 0
				c.JsonData.ErrorMsg = "更新成功"
				c.JsonData.Datas = []interface{}{}
			}
		}*/
	}
}


