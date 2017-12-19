package handler

import (
	"fit"
	"encoding/json"
	"fmt"
	"nursing/model"
)

type HostConfigController struct {
	PCController
}

func (c HostConfigController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)        //用户信息
	departments, err1 := model.QueryDepartmentList(true) //科室信息
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
		_ = c.LoadViewForAdministrator(true,w, r, "pc/v_host_config.html", "pc/header_side.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}
