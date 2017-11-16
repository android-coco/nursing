package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type PCNRL7Controller struct {
	PCController
}

func (c PCNRL7Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "7")
	if !has {
		return
	}
	// 起止时间  页码
	datestr1, datestr2, pageindex, pagenum, err := c.GetPageInfo(w, r, "7", pid)
	if err != nil {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL7(pid, datestr1, datestr2, pageindex)

	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	var pmodel model.NRL7
	if len(mods) > 0 {
		pmodel = mods[len(mods)-1]
	}

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRL08":     pmodel.NRL08,
		"NRL08A":    pmodel.NRL08A,
		"NRL08B":    pmodel.NRL08B,
		"NRLList":   mods,
		"PageNum":   pagenum,
		"PageIndex": pageindex,
		"Menuindex": "7-7",
	}

	c.LoadViewSafely(w, r, "pc/v_nrl7.html", "pc/header_side.html", "pc/header_top.html")
}
