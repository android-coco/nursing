package handler

import (
	"fit"
	"fmt"
	"strconv"
	"time"
)

// 护理记录单 PC端
type PCNRL1Controller struct {
	PCController
}

// 护理记录单
// 查看  list

func (c PCNRL1Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {

	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w,r, "1")
	if !has {
		return
	}


	// 时间
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	datestr1 := time.Unix(date1 / 1000 - 60 * 60 * 8, 0).Format("2006-01-02 15:04:05")
	datestr2 := time.Unix(date2 / 1000 + 60 * 60 * 16, 0).Format("2006-01-02 15:04:05")
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	}
	fmt.Println("-----", date1,date2, datestr1, datestr2, pid)

	// 总条数
	/*count, errCount := model.PCQueryNRL1PageCount(pid, datestr1, datestr2)
	if errCount != nil {
		fmt.Fprintln(w, "参数错误！ page count:", err)
		return
	}

	fmt.Println("count:", count)
	//总页数
	pageNum := int(count / 9) + 1
	//当前页数
	index := r.FormValue("num")
	pageindex, errnum := strconv.Atoi(index)
	if errnum != nil {
		pageindex = int(pageNum)
		fit.Logger().LogError("parse index error :", errnum)
	}
	if pageindex < 1 {
		pageindex = 1
	} else if pageindex > pageNum {
		pageindex = pageNum
	}
	fmt.Println("count:", count, "pageNum:", pageNum, "pageindex:", pageindex)

	// 护理单
	mods, err13 := model.PCQueryNRL1(pid, datestr1, datestr2, pageindex)

	if err13 != nil {
		fmt.Fprintln(w, "参数错误！ pc get nrl8 model error", err13)
		return
	}
	fmt.Printf("mods %+v\n %d\n\n", mods, len(mods))*/

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		//"NRLList":   mods,
		"PageNum":   2,
		"PageIndex": 2,
		"Menuindex": "7-1",
	}

	c.LoadViewSafely(w, r, "pc/v_nrl1.html", "pc/header_side.html", "pc/header_top.html")
}

/*func (c PCNRL1Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params)  {
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

}*/

// 护理记录单 录入
func (c PCNRL1Controller) NRL1TypeIn(w *fit.Response, r *fit.Request, p fit.Params)  {

}