package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type NRL8Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL8Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL8{}
	c.NRLCheck(w, r, p, &nrl)
}

func (c NRL8Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL8{}
	c.NRLEdit(w, r, p, &nrl)
}

// 接口
// 添加护理记录单
func (c NRL8Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL8{}
	c.NRLAddRecord(w, r, p, &mod)
}

// 修改护理记录单
func (c NRL8Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL8{}
	c.NRLUpdateRecord(w, r, p, &mod)
}
// 删除护理
func (c NRL8Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL8{}
	c.NRLDeleteRecord(w, r, p, nrl)
}

type PCNRL8Controller struct {
	PCNRLController
}

func (c PCNRL8Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	data, has := c.BaseNRLRecord(w, r, "8")
	if !has {
		return
	}

	// 护理单
	mods, err := model.PCQueryNRL8(data.pid, data.datestr1, data.datestr2, data.pageIndex)
	if err != nil {
		fmt.Fprintln(w, "参数错误！ pc get nrl8 model error", err)
		return
	}

	c.Data = fit.Data{
		"Userinfo":  data.userInfo, // 护士信息
		"PInfo":     data.pInfo,    // 病人信息
		"Beds":      data.beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   data.pageNum,
		"PageIndex": data.pageIndex,
		"Menuindex": "7-8",
	}
	c.LoadPCNRLView(w, r, "pcnrl/v_nrl8.html")
}