package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type NRL6Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL6Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL6{}
	c.NRLCheck(w, r, p, &nrl)
}

func (c NRL6Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL6{}
	c.NRLEdit(w, r, p, &nrl)
}

// 接口
// 添加护理记录单
func (c NRL6Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL6{}
	c.NRLAddRecord(w, r, p, &mod)
}

// 修改护理记录单
func (c NRL6Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL6{}
	c.NRLUpdateRecord(w, r, p, &mod)
}

// 删除护理
func (c NRL6Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL6{}
	c.NRLDeleteRecord(w, r, p, nrl)
}


type PCNRL6Controller struct {
	PCNRLController
}

func (c PCNRL6Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	data, has := c.BaseNRLRecord(w, r, "6")
	if !has {
		return
	}

	// 护理单
	mods, err := model.PCQueryNRL6(data.pid, data.datestr1, data.datestr2, data.pageIndex)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	c.Data = fit.Data{
		"Userinfo":  data.userInfo, // 护士信息
		"PInfo":     data.pInfo,    // 病人信息
		"Beds":      data.beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   data.pageNum,
		"PageIndex": data.pageIndex,
		"Menuindex": "7-6",
	}
	c.LoadPCNRLView(w, r, "pcnrl/v_nrl6.html")
}