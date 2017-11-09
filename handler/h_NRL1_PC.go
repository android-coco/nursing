package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
)

// 护理记录单 PC端
type PCNRL1Controller struct {
	PCController
}

// 护理记录单
// 查看  list
func (c PCNRL1Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params)  {
	// 病人id
	pid := r.FormValue("pid")

	// 起止时间  如果传默认为全部
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")
	// 页码
	page := r.FormValue("page")
	if pid == "" {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	pageInt, err1 := strconv.ParseInt(page, 10, 32)
	if err1 != nil {
		fit.Logger().LogError("error", "parse page error", err1)
		pageInt = 0
	}

	//patintinfos, err := model.QueryPatientInfo(pid)
	//if err != nil {
	//	fmt.Fprintln(w, "服务器有点繁忙！")
	//	return
	//}

	nrl1 := model.NRL1{PatientId:pid}
	list, err := nrl1.QueryNRL1ByDate(startDate, endDate, int(pageInt))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("length", len(list))
	fmt.Printf("%#v", list)

	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "上传成功!"
	c.JsonData.Datas = list

	defer c.ResponseToJson(w)

}

// 护理记录单 录入
func (c PCNRL1Controller) NRL1TypeIn(w *fit.Response, r *fit.Request, p fit.Params)  {

}