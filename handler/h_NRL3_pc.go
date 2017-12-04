package handler

import (
	"fit"
	"nursing/model"
	"fmt"
)

type PCNRL3Controller struct {
	PCController
}

func (c PCNRL3Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "3")
	if !has {
		return
	}

	datestr1, datestr2, pageindex, pagenum, err := c.GetPageInfo(w, r, "3", pid)
	if err != nil {
		fit.Logger().LogError("nrl page info :", err)
		return
	}

	fmt.Println("--------------:", datestr1, datestr2)
	/*// 时间
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	datestr1 := time.Unix(date1 / 1000 - 60 * 60 * 8, 0).Format("2006-01-02 15:04:05")
	datestr2 := time.Unix(date2 / 1000 + 60 * 60 * 16, 0).Format("2006-01-02 15:04:05")
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	}
	fmt.Println("-----", date1,date2, datestr1, datestr2)

	// 总条数
	count, errCount := model.PCQUeryNRL3PageCount(pid, datestr1, datestr2)
	if errCount != nil {
		fmt.Fprintln(w, "参数错误！  user info error")
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
		fit.Logger().LogError("")
	}
	if pageindex < 1 {
		pageindex = 1
	} else if pageindex > pageNum {
		pageindex = pageNum
	}
	fmt.Println("count:", count, "pageNum:", pageNum, "pageindex:", pageindex)

	// 护理单
	mods, err13 := model.PCQueryNRL3(pid, datestr1, datestr2, pageindex)

	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error")
		return
	}
	fmt.Printf("mods %+v\n\n %d\n\n", mods, len(mods))*/


	// 护理单
	mods, err13 := model.PCQueryNRL3(pid, datestr1, datestr2, pageindex)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error")
		return
	}


	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   pagenum,
		"PageIndex": pageindex,
		"Menuindex": "7-3",
	}

	c.LoadViewSafely(w, r, "pc/v_nrl3.html", "pc/header_side.html", "pc/header_top.html")
}
