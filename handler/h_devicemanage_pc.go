package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"encoding/json"
)

type DeviceManageController struct {
	PCController
}

func (c DeviceManageController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)        //用户信息
	departments, err1 := model.QueryDepartmentList() //科室信息
	if err == nil && err1 == nil {
		devices, _ := model.GetAllDevices()
		for i := range devices {
			var tcDetails []model.DeviceInfos
			jsonBlob := []byte(devices[i].Devicelist)
			err := json.Unmarshal(jsonBlob, &tcDetails)
			if err != nil {
				fmt.Println("error:", err)
			}
			devices[i].DeviceInfos = tcDetails
			//fit.Logger().LogError("adsfasdfasdf111:",devices[i].DeviceInfos)
		}
		c.Data = fit.Data{
			"Userinfo":    userinfo,
			"Departments": departments,
			"Devices":     devices,
		}
		//fit.Logger().LogError("adsfasdfasdf:",devices[0].DeviceInfos)
		_ = c.LoadViewSafely(w, r, "pc/v_device_manage.html", "pc/header_side.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

func (c DeviceManageController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	classid := r.FormValue("classid")
	if classid == "all" {
		devices, _ := model.GetAllDevices()
		for i := range devices {
			var tcDetails []model.DeviceInfos
			jsonBlob := []byte(devices[i].Devicelist)
			err := json.Unmarshal(jsonBlob, &tcDetails)
			if err != nil {
				fmt.Println("error:", err)
			}
			devices[i].DeviceInfos = tcDetails
			//fit.Logger().LogError("adsfasdfasdf111:",devices[i].DeviceInfos)
		}
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "查询成功"
		c.JsonData.Datas = devices
	} else {
		devicesclass, err := strconv.Atoi(classid)
		if err == nil {
			devices, _ := model.GetDevicesByClass(devicesclass)
			for i := range devices {
				var tcDetails []model.DeviceInfos
				jsonBlob := []byte(devices[i].Devicelist)
				err := json.Unmarshal(jsonBlob, &tcDetails)
				if err != nil {
					fmt.Println("error:", err)
				}
				devices[i].DeviceInfos = tcDetails
			}
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "查询成功"
			c.JsonData.Datas = devices
		} else {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "查询错误"
			c.JsonData.Datas = err.Error()
		}
	}

}
