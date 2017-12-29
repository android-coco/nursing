//  Created by JP

package handler

import (
	"fit"
	"strconv"
	"nursing/model"
	"fmt"
)

type PCHomeController struct {
	PCController
}

type PCBedController struct {
	PCController
}


/*PC 主页*/
func (c PCHomeController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		num := r.FormValue("num")
		if num == "" {
			num = "0"
		}
		typeDup, errtp := strconv.Atoi(num)
		if errtp != nil {
			fit.Logger().LogError("home", errtp)
			fmt.Fprintln(w, "服务器有点繁忙！")
			return
		}

		departmentId := userinfo.DepartmentID
		response, err := model.GetDepartmentBedsByClassifying(departmentId, typeDup)
		if err != nil {
			fmt.Fprintln(w, "服务器有点繁忙！"+err.Error())
			return
		}

		menu := "t" + num
		len := 0
		if obj, ok := response["num"].(map[string]int); ok {
			 len = obj[menu]
		}

		menuindex := "1-" + num
		c.Data = fit.Data {
			"Userinfo": userinfo,
			"Menuindex": menuindex,
			"Beds": response["bed"],
			"Num": response["num"],
			"Len": len,
		}
		//fit.Logger().LogError("Beds:",response["bed"])
		_ = c.LoadViewSafely(w, r, "pc/v_index.html", "pc/header_side.html","pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}


func (c PCBedController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	r.ParseForm()
	department_id := r.FormValue("department_id")
	type_dup := r.FormValue("type")

	if department_id == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	depid_i, err_dep := strconv.Atoi(department_id)
	if err_dep != nil || depid_i < 0 {
		c.RenderingJsonAutomatically(2, "参数错误： department_id")
		return
	}
	if type_dup == "" {
		type_dup = "0"
	}
	type_i, err_tp := strconv.Atoi(type_dup)
	if err_tp != nil || type_i > 5 {
		c.RenderingJsonAutomatically(2, "参数错误： type")
		return
	}

	//showEmpty_i, err_show := strconv.Atoi(showEmpty)
	//showEmpty_b := false
	//if err_show != nil || showEmpty_i < 0 {
	//	c.RenderingJsonAutomatically(2, "参数错误： showEmpty")
	//	return
	//} else if showEmpty_i > 0 {
	//	showEmpty_b = true
	//}
	response, err := model.GetDepartmentBedsByClassifying(depid_i, type_i)
	if err != nil {
		c.RenderingJsonAutomatically(3, "Database "+err.Error())
	} else {
		c.RenderingJson(0, "成功", response)
	}
}

func (c *PCBedController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *PCBedController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}