package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"
)

type PCNRL5Controller struct {
	PCController
}



func (c PCNRL5Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "5")
	if !has {
		return
	}
	// 起止时间  页码
	// 时间
	var datestr1, datestr2 string
	var pageindex, pagenum int
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000+60*60*24 - 1, 0).Format("2006-01-02 15:04:05")
	}

	// 护理单
	mods, err13 := model.PCQueryNRL5(pid, datestr1, datestr2, pageindex)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err13)
		return
	}

	//总条数
	count := len(mods)
	peerPage := 5
	//总页数
	pagenum = int((count-1)/ peerPage) + 1
	//当前页数
	index := r.FormValue("num")
	pageindex, errnum := strconv.Atoi(index)
	if errnum != nil {
		pageindex = int(pagenum)
	}
	if pageindex < 1 {
		pageindex = 1
	} else if pageindex > pagenum {
		pageindex = pagenum
	}

	list :=  make([]model.APNModel, peerPage)
	if pageindex == pagenum {
		list = mods[(pageindex - 1) * peerPage:count]
	} else {
		list = mods[(pageindex - 1) * peerPage:pageindex * peerPage]
	}

	//fmt.Printf("mods %+v\n %d\n\n", mods, len(mods))
	//fmt.Printf("list %+v\n %d\n\n", list, len(list))
	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRLList":   list,
		"PageNum":   pagenum,
		"PageIndex": pageindex,
		"Menuindex": "7-5",
	}


	c.LoadViewSafely(w, r, "pc/v_nrl5.html", "pc/header_side.html", "pc/header_top.html")
}

