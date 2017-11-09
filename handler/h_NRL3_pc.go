package handler

import (
	"fit"
	"nursing/model"
	"fmt"
	"strconv"
)

type PCNRL3Controller struct {
	PCController
}

func (c PCNRL3Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
	}
	beds, err := model.QueryDepartmentBeds(userinfo.DepartmentID, false)

	fmt.Printf("bed1 : %+v\n", beds)
	if err != nil {
		fit.Logger().LogError("pc nrl2", err)
	}

	// 病人id  病人信息
	pid := r.FormValue("pid")
	var pInfo model.PCBedDup
	if pid == "" {
		pidnum := beds[0].VAA01
		pid = strconv.Itoa(pidnum)

		url := "/pc/record/nrl3?pid=" + pid
		c.Redirect(w, r, url, 302)
		return
	}

	fmt.Println("pid", pid)
	// 病人信息
	for _, val := range beds {
		if strconv.Itoa(val.VAA01) == pid {
			pInfo = val
			break
		}
	}
	if pInfo.VAA01 == 0 {
		fit.Logger().LogError("pc nrl3 pInfo is empty")
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	count, errCount := model.PCQUeryNRL3PageCount(pid)
	if errCount != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	pageNum := count/9 + 1
	index := r.FormValue("num")
	pageindex, errnum := strconv.Atoi(index)
	if errnum != nil {
		pageindex = pageNum
		fit.Logger().LogError("")
	}
	fmt.Println(count, "---", pageNum, "---", pageindex)

	// 护理单
	mods, errl3 := model.PCQueryNRL3(pid, pageindex)
	if errl3 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}
	fmt.Printf("mods %+v\n\n %d\n\n", mods, len(mods))

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   pageNum,
		"PageIndex": pageindex,
		"Menuindex": "7-3",
	}

	//fmt.Printf("data ------ : %+v\n", c.Data)

	c.LoadViewSafely(w, r, "pc/v_nrl3.html", "pc/header_side.html", "pc/header_top.html")
	//c.LoadView(w, "pc/v_nrl3.html", "pc/header_side.html", "pc/header_top.html")
}

func Test() string {
	return "heheehhe"
}
