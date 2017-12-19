package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"

)

type NRL3Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL3Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL3{}
	c.NRLCheck(w, r, p, &nrl)
}

func (c NRL3Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL3{}
	c.NRLEdit(w, r, p, &nrl)
}

// 接口
// 添加护理记录单
func (c NRL3Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL3{}
	c.NRLAddRecord(w, r, p, &mod)
}

// 修改护理记录单
func (c NRL3Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	mod := model.NRL3{}
	c.NRLUpdateRecord(w, r, p, &mod)
}


// 删除护理
func (c NRL3Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	nrl := model.NRL3{}
	c.NRLDeleteRecord(w, r, p, nrl)
}


// 修改护理记录单 审核人签名
func (c NRL3Controller) UpdateChecker(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormValue("rid")
	ty := r.FormValue("nrltype")
	nurseName := r.FormValue("nurseName")
	id, err := strconv.ParseInt(rid, 10, 64)
	datestr := r.FormValue("datetime")
	_, err4 := time.ParseInLocation("2006-01-02 15:04:05", datestr, time.Local)

	if err != nil || ty == "" || nurseName == "" || (datestr != "" && err4 != nil) {
		fit.Logger().LogError("hy:", "nrl checker update :", err)
		c.RenderingJsonAutomatically(10002, "参数不完整")
		return
	}

	fmt.Println("Datatime:", datestr)
	errupdate := model.UpdateNRLChcker(ty, rid, datestr, nurseName)

	if errupdate != nil {
		fit.Logger().LogError("hy:","update  checker :", errupdate)
		c.RenderingJson(2, "修改失败！", errupdate.Error())
	} else {
		_, errRecord := model.UpadteNRecords(id, r.FormValue("datetime"))
		checkerr("nurse record update checker err:", errRecord)
		c.RenderingJsonAutomatically(0, "修改成功！")
	}
}

type PCNRL3Controller struct {
	PCNRLController
}

func (c PCNRL3Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {

	data, has := c.BaseNRLRecord(w, r, "3")
	if !has {
		return
	}
	// 护理单
	mods, err13 := model.PCQueryNRL3(data.pid, data.datestr1, data.datestr2, data.pageIndex)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error")
		return
	}


	c.Data = fit.Data{
		"Userinfo":  data.userInfo, // 护士信息
		"PInfo":     data.pInfo,    // 病人信息
		"Beds":      data.beds,     // 床位list
		"NRLList":   mods,
		"PageNum":   data.pageNum,
		"PageIndex": data.pageIndex,
		"Menuindex": "7-3",
	}
	c.LoadPCNRLView(w, r, "pcnrl/v_nrl3.html")
	//c.LoadViewSafely(w, r, "pcnrl/v_nrl3.html", "pc/header_side.html", "pc/header_top.html")
}