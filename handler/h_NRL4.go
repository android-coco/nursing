package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type NRL4Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL4Controller) Check(w *fit.Response, r *fit.Request, p fit.Params)  {
	nrl := model.NRL4{}
	c.NRLCheck(w, r, p, &nrl)
}

func (c NRL4Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL4{}
	c.NRLEdit(w, r, p, &nrl)
}

// 接口
// 添加护理记录单
func (c NRL4Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL4{}
	c.NRLAddRecord(w, r, p, &mod)
}

// 修改护理记录单
func (c NRL4Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL4{}
	c.NRLUpdateRecord(w, r, p, &mod)
}

// 删除护理
func (c NRL4Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL4{}
	c.NRLDeleteRecord(w, r, p, nrl)
}


type PCNRL4Controller struct {
	PCNRLController
}

func (c PCNRL4Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	data, has := c.BaseNRLRecord(w, r, "4")
	if !has {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL4(data.pid, data.datestr1, data.datestr2, data.pageIndex)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err13)
		fit.Logger().LogError("hy:", "nrl page info :", err13)
		return
	}

	c.Data = fit.Data{
		"Userinfo":  data.userInfo, // 护士信息
		"PInfo":     data.pInfo,    // 病人信息
		"Beds":      data.beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   data.pageNum,
		"PageIndex": data.pageIndex,
		"Menuindex": "7-4",
	}
	c.LoadPCNRLView(w, r, "pcnrl/v_nrl4.html")
}