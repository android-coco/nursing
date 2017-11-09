package handler

import (
	"fit"
	"fmt"
	"strconv"
	"time"
	"nursing/model"
)

type NRL5Controller struct {
	fit.Controller

}

// 模板 template PDA端
func (c NRL5Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr5, err1 := model.QueryNRL5(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR5", err1)
	}

	pid := nr5.VAA01

	pinfo, err := model.GetPatientInfo(strconv.FormatInt(pid, 10))
	if err != nil {
		fit.Logger().LogError("m_NR1", err)
		fmt.Fprintln(w, "服务器有点繁忙！")
		return
	}

	recordDate := nr5.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo": pinfo[0],
		"NRL":   nr5,
		"RecordDate": recordDate,
	}

	fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "v_nrl5.html")
}

func (c NRL5Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")
	uid := r.FormValue("uid")
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr5 model.NRL5
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
		nr5, err1 = model.QueryNRL5(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr5.VAA01, 10)
		uid = nr5.BCE01A
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
	if len(pinfo) == 0 {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	account, err2 := model.FetchAccountWithUid(uid)
	if err2 != nil {
		fit.Logger().LogError("nrl5", err2)
		fmt.Fprintln(w, "参数错误！", err2)
		return
	}
	recordDate := nr5.DateTime.Format("2006-01-02")

	c.Data = fit.Data{
		"Pinfo": pinfo[0],
		"NRL":   nr5,
		"Type":  ty,
		"Rid": rid,
		"Account": account,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl5_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL5Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	// 科室ID
	BCK01, err5 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 文书ID（修改时需要）
	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02A, err7 := strconv.Atoi(r.FormValue("NRL02A"))
	NRL02B, err72 := strconv.Atoi(r.FormValue("NRL02B"))
	NRL03A, err8 := strconv.Atoi(r.FormValue("NRL03A"))
	NRL03B, err82 := strconv.Atoi(r.FormValue("NRL03B"))
	NRL04A, err9 := strconv.Atoi(r.FormValue("NRL04A"))
	NRL04B, err92 := strconv.Atoi(r.FormValue("NRL04B"))
	NRL05A, err10 := strconv.Atoi(r.FormValue("NRL05A"))
	NRL05B, err102 := strconv.Atoi(r.FormValue("NRL05B"))
	NRL06A, err11 := strconv.Atoi(r.FormValue("NRL06A"))
	NRL06B, err112 := strconv.Atoi(r.FormValue("NRL06B"))
	NRL07A, err12 := strconv.Atoi(r.FormValue("NRL07A"))
	NRL07B, err122 := strconv.Atoi(r.FormValue("NRL07B"))
	NRL08A, err13 := strconv.ParseFloat(r.FormValue("NRL08A"), 32)
	NRL08B, err132 := strconv.ParseFloat(r.FormValue("NRL08B"), 32)
	NRL09A, err14 := strconv.Atoi(r.FormValue("NRL09A"))
	NRL09B, err142 := strconv.Atoi(r.FormValue("NRL09B"))
	NRL10A, err15 := strconv.Atoi(r.FormValue("NRL10A"))
	NRL10B, err152 := strconv.Atoi(r.FormValue("NRL10B"))
	NRL11, err16 := strconv.Atoi(r.FormValue("NRL11"))
	NRL12, err162 := strconv.Atoi(r.FormValue("NRL12"))
	score, err18 := strconv.Atoi(r.FormValue("score"))


	/*if err1 != nil || err4 != nil || err5 != nil || err6 != nil ||
		err7 != nil || err72 != nil || err8 != nil || err82 != nil || err9 != nil || err92 != nil ||
		err10 != nil || err102 != nil || err11 != nil || err112 != nil || err12 != nil || err122 != nil ||
		err13 != nil || err132 != nil || err14 != nil || err142 != nil || err15 != nil || err16 != nil || err162 != nil {
		fit.Logger().LogError("nrl5 add", err1, err4, err5, err6, err7, err72, err8, err82, err9, err92, err10, err102, err11, err112, err12, err122, err13, err132, err14, err142, err15, err152, err16, err162)
	}*/

	checkerr( "nrl5 add", err1, err4, err5, err6, err7,
		err72, err8, err82, err9, err92, err10, err102, err11,
		err112, err12, err122, err13, err132, err14, err142, err15, err152, err16, err162, err18)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl5 := model.NRL5{
		BCK01:    BCK01,
		VAA01:    VAA01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02A:    NRL02A,
		NRL02B:    NRL02B,
		NRL03A:    NRL03A,
		NRL03B:    NRL03B,
		NRL04A:    NRL04A,
		NRL04B:    NRL04B,
		NRL05A:    NRL05A,
		NRL05B:    NRL05B,
		NRL06A:    NRL06A,
		NRL06B:    NRL06B,
		NRL07A:    NRL07A,
		NRL07B:    NRL07B,
		NRL08A:    float32(NRL08A),
		NRL08B:    float32(NRL08B),
		NRL09A:    NRL09A,
		NRL09B:    NRL09B,
		NRL10A:    NRL10A,
		NRL10B:    NRL10B,
		NRL11:    NRL11,
		NRL12:    NRL12,
		Score:    score,
	}

	_, err17 := nrl5.InsertData()
	if err17 != nil {
		fit.Logger().LogError("NRL5 add :", err17)
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
func (c NRL5Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	//BCE01A := r.FormValue("uid")
	// 护士名
	//BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02A, err7 := strconv.Atoi(r.FormValue("NRL02A"))
	NRL02B, err72 := strconv.Atoi(r.FormValue("NRL02B"))
	NRL03A, err8 := strconv.Atoi(r.FormValue("NRL03A"))
	NRL03B, err82 := strconv.Atoi(r.FormValue("NRL03B"))
	NRL04A, err9 := strconv.Atoi(r.FormValue("NRL04A"))
	NRL04B, err92 := strconv.Atoi(r.FormValue("NRL04B"))
	NRL05A, err10 := strconv.Atoi(r.FormValue("NRL05A"))
	NRL05B, err102 := strconv.Atoi(r.FormValue("NRL05B"))
	NRL06A, err11 := strconv.Atoi(r.FormValue("NRL06A"))
	NRL06B, err112 := strconv.Atoi(r.FormValue("NRL06B"))
	NRL07A, err12 := strconv.Atoi(r.FormValue("NRL07A"))
	NRL07B, err122 := strconv.Atoi(r.FormValue("NRL07B"))
	NRL08A, err13 := strconv.ParseFloat(r.FormValue("NRL08A"), 32)
	NRL08B, err132 := strconv.ParseFloat(r.FormValue("NRL08B"), 32)
	NRL09A, err14 := strconv.Atoi(r.FormValue("NRL09A"))
	NRL09B, err142 := strconv.Atoi(r.FormValue("NRL09B"))
	NRL10A, err15 := strconv.Atoi(r.FormValue("NRL10A"))
	NRL10B, err152 := strconv.Atoi(r.FormValue("NRL10B"))
	NRL11, err16 := strconv.Atoi(r.FormValue("NRL11"))
	NRL12, err162 := strconv.Atoi(r.FormValue("NRL12"))
	score, err18 := strconv.Atoi(r.FormValue("score"))

	checkerr( "nrl5 add", err4, err6, err7,
		err72, err8, err82, err9, err92, err10, err102, err11,
		err112, err12, err122, err13, err132, err14, err142, err15, err152, err16, err162, err18)

	nrl5 := model.NRL5{
		//BCE01A:   BCE01A,
		//BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02A:    NRL02A,
		NRL02B:    NRL02B,
		NRL03A:    NRL03A,
		NRL03B:    NRL03B,
		NRL04A:    NRL04A,
		NRL04B:    NRL04B,
		NRL05A:    NRL05A,
		NRL05B:    NRL05B,
		NRL06A:    NRL06A,
		NRL06B:    NRL06B,
		NRL07A:    NRL07A,
		NRL07B:    NRL07B,
		NRL08A:    float32(NRL08A),
		NRL08B:    float32(NRL08B),
		NRL09A:    NRL09A,
		NRL09B:    NRL09B,
		NRL10A:    NRL10A,
		NRL10B:    NRL10B,
		NRL11:    NRL11,
		NRL12:    NRL12,
		Score:    score,
	}

	_, err17 := nrl5.UpdateData(id)
	if err17 != nil {
		fit.Logger().LogError("nrl4 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "修改失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "修改成功！"
		c.JsonData.Datas = []interface{}{}
	}
}

func checkerr(tag string,err ...error)  {
	for _, val := range err {
		if val != nil {
			fit.Logger().LogError(tag, val)
		}
	}
}