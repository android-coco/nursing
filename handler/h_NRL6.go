package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"
)

type NRL6Controller struct {
	fit.Controller
}

// 模板 template PDA端
func (c NRL6Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr6, err1 := model.QueryNRL6(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR6", err1)
	}

	pid := nr6.VAA01

	pinfo, err := model.GetPatientInfo(strconv.FormatInt(pid, 10))
	if err != nil {
		fit.Logger().LogError("m_NR1", err)
		fmt.Fprintln(w, "服务器有点繁忙！")
		return
	}

	recordDate := nr6.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo": pinfo[0],
		"NRL":   nr6,
		"RecordDate": recordDate,
	}

	fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "v_nrl6.html")
}

func (c NRL6Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")//病人id
	uid := r.FormValue("uid") //护士id
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr6 model.NRL6
	if ty == "1" {
		if "" == pid || "" == uid {
			fmt.Fprintln(w, "参数错误！")
			return
		}
	} else if ty == "2" {
		if rid == "" {
			fmt.Fprintln(w, "参数错误！")
			return
		}
		var err1 error
		nr6, err1 = model.QueryNRL6(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr6.VAA01, 10)
		uid = nr6.BCE01A
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	pinfo, err := model.GetPatientInfo(pid)
	if err != nil {
		fit.Logger().LogError("m_NR1", err)
		fmt.Fprintln(w, "服务器有点繁忙！")
		return
	}

	account, err2 := model.FetchAccountWithUid(uid)
	if err2 != nil {
		fit.Logger().LogError("nrl5", err2)
		fmt.Fprintln(w, "参数错误！", err2)
		return
	}
	//fmt.Printf("account %+v \n\n %+v\n\n", account, pinfo)


	if len(pinfo) == 0 {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	recordDate := nr6.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo": pinfo[0],
		"NRL":   nr6,
		"Type":  ty,
		"Rid": rid,
		"Account": account,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl6_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL6Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 护士ID
	BCE01A := r.FormValue("nid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	// 科室ID
	BCK01, err5 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 文书ID（修改时需要）
	//recordId := r.FormValue("rid")
	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02, err7 := strconv.Atoi(r.FormValue("NRL02"))
	NRL03, err8 := strconv.Atoi(r.FormValue("NRL03"))
	NRL04, err9 := strconv.Atoi(r.FormValue("NRL04"))
	NRL05, err10 := strconv.Atoi(r.FormValue("NRL05"))
	NRL06, err11 := strconv.Atoi(r.FormValue("NRL06"))
	NRL07A, err12 := strconv.Atoi(r.FormValue("NRL07A"))
	NRL07B, err13 := strconv.Atoi(r.FormValue("NRL07B"))
	score, err14 := strconv.Atoi(r.FormValue("score"))

	/*if err1 != nil || err4 != nil ||
		err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil ||
		err10 != nil || err11 != nil || err12 != nil || err13 != nil  {
		fit.Logger().LogError("nrl3 add", err1, err4, err5, err6, err7, err8, err9, err10, err11, err12, err13, err14, err15, err16)
	}*/
	checkerr("nrl6 add", err1, err4, err5, err6, err7, err8, err9, err10, err11, err12, err13, err14)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl6 := model.NRL6{
		BCK01:    BCK01,
		VAA01:    VAA01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02:    NRL02,
		NRL03:    NRL03,
		NRL04:    NRL04,
		NRL05:    NRL05,
		NRL06:    NRL06,
		NRL07A:    NRL07A,
		NRL07B:    NRL07B,
		Score:    score,
	}

	_, err17 := nrl6.InsertData()
	if err17 != nil {
		fit.Logger().LogError("NRL6 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{}
	}

}

// 修改护理记录单
func (c NRL6Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormValue("rid")
	id, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		fit.Logger().LogError("Error", "nrl1 update :", err)
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "rid 错误！"
		c.JsonData.Datas = []interface{}{}
	}

	// 护士ID
	//BCE01A := r.FormValue("nid")
	// 护士名
	//BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02, err7 := strconv.Atoi(r.FormValue("NRL02"))
	NRL03, err8 := strconv.Atoi(r.FormValue("NRL03"))
	NRL04, err9 := strconv.Atoi(r.FormValue("NRL04"))
	NRL05, err10 := strconv.Atoi(r.FormValue("NRL05"))
	NRL06, err11 := strconv.Atoi(r.FormValue("NRL06"))
	NRL07A, err12 := strconv.Atoi(r.FormValue("NRL07A"))
	NRL07B, err13 := strconv.Atoi(r.FormValue("NRL07B"))
	score, err14 := strconv.Atoi(r.FormValue("score"))

	checkerr("nrl6 update", err4, err6, err7, err8, err9, err10, err11, err12, err13, err14)

	nrl6 := model.NRL6{
		//BCE01A:   BCE01A,
		//BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02:    NRL02,
		NRL03:    NRL03,
		NRL04:    NRL04,
		NRL05:    NRL05,
		NRL06:    NRL06,
		NRL07A:    NRL07A,
		NRL07B:    NRL07B,
		Score:    score,
	}

	_, err17 := nrl6.UpdateData(id)
	if err17 != nil {
		fit.Logger().LogError("nrl6 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "修改失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "修改成功！"
		c.JsonData.Datas = []interface{}{}
	}
}


