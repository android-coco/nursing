package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type PCNRL6Controller struct {
	PCController
}

func (c PCNRL6Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "6")
	if !has {
		return
	}
	// 起止时间  页码
	datestr1, datestr2, pageindex, pagenum, err := c.GetPageInfo(w, r, "6", pid)
	if err != nil {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL6(pid, datestr1, datestr2, pageindex)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   pagenum,
		"PageIndex": pageindex,
		"Menuindex": "7-6",
	}

	c.LoadViewSafely(w, r, "pc/v_nrl6.html", "pc/header_side.html", "pc/header_top.html")
}