package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type NRL7Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL7Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL7{}
	c.NRLCheck(w, r, p, &nrl)

}

func (c NRL7Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL7{}
	c.NRLEdit(w, r, p, &nrl)
}

// 接口
// 添加护理记录单
func (c NRL7Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL7{}
	c.NRLAddRecord(w, r, p, &mod)
}

// 修改护理记录单
func (c NRL7Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL7{}
	c.NRLUpdateRecord(w, r, p, &mod)
}


func (c NRL7Controller) UpdateTitle(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	mod := model.NRL7Title{}
	c.FitSetStruct(&mod, r)

	// 记录时间
	datetime, err2 := r.FormTimeStruct("datetime")
	mod.DateTime = model.FitTime(datetime)
	if err2 != nil {
		c.RenderingJson(10001, "参数错误！", err2.Error())
		return
	}

	nrl7title, errt := mod.PCUpdateNRT7Title()
	if errt != nil {
		fit.Logger().LogError("hy:", "nrl1 update :", errt)
		c.RenderingJson(2, "DB错误！", errt.Error())
	} else {
		c.RenderingJson(0, "nrl7 title 修改成功！", nrl7title)
	}
}

// 删除护理
func (c NRL7Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL7{}
	c.NRLDeleteRecord(w, r, p, nrl)
}


type PCNRL7Controller struct {
	PCNRLController
}

func (c PCNRL7Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	data, has := c.BaseNRLRecord(w, r, "7")
	if !has {
		return
	}
	// 护理单
	mods, err := model.PCQueryNRL7(data.pid, data.datestr1, data.datestr2, data.pageIndex)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	nrl7Title := model.NRL7Title{PatientId:data.pInfo.VAA01}

	errTitle := nrl7Title.PCQueryNRL7Title()
	if errTitle != nil {
		fit.Logger().LogError("PCQueryNRL1Title error :", errTitle)
	}


	c.Data = fit.Data{
		"Userinfo":  data.userInfo, // 护士信息
		"PInfo":     data.pInfo,    // 病人信息
		"Beds":      data.beds,     // 床位list
		"NRLTitle":  nrl7Title,
		"NRLList":   mods,
		"PageNum":   data.pageNum,
		"PageIndex": data.pageIndex,
		"Menuindex": "7-7",
	}
	c.LoadPCNRLView(w, r, "pcnrl/v_nrl7.html")
}
