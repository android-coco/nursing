package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"
)

type NRL7Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL7Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr7, err1 := model.QueryNRL7(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR7", err1)
	}
	nrl7Title := model.NRL7Title{
		VAA01: nr7.VAA01,
		BCK01: nr7.BCK01,
	}
	pid := nr7.VAA01

	errt := nrl7Title.PCQueryNRL7Title()
	if errt != nil {
		fit.Logger().LogError("m_NR7", err1)
	}
	fmt.Println("title:", nrl7Title)

	// 查询对应病人信息
	patient, has := c.LoadPinfoWithPid(w, r, pid)
	if !has {
		return
	}

	rdate := nr7.DateTime.Format("2006-01-02")
	rtime := nr7.DateTime.Format("15:04")
	c.Data = fit.Data{
		"Pinfo":      patient,
		"NRL":        nr7,
		"Title":      nrl7Title, // 表头 25，26
		"RecordDate": rdate,
		"RecordTime": rtime,
	}

	c.LoadView(w, "v_nrl7.html")
}

func (c NRL7Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid") //病人id
	uid := r.FormValue("uid") //护士id
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr7 model.NRL7
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
		nr7, err1 = model.QueryNRL7(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr7.VAA01, 10)
		uid = nr7.BCE01A
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPinfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}

	nrl7Title := model.NRL7Title{
		VAA01: nr7.VAA01,
		BCK01: nr7.BCK01,
	}
	if ty == "1" {
		nrl7Title = model.NRL7Title{
			VAA01: int64(patient.VAA01),
			BCK01: int64(patient.BCK01C),
		}
	}

	errt := nrl7Title.PCQueryNRL7Title()
	if errt != nil {
		fit.Logger().LogError("m_NR7 title:", errt)
	}


	recordDate := nr7.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo":      patient,
		"NRL":        nr7,
		"Title":      nrl7Title, // 表头 25，26
		"Type":       ty,
		"Rid":        rid,
		"Account":    account,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl7_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL7Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	NRL01 := r.FormValue("NRL01")
	NRL02 := r.FormValue("NRL02")
	NRL03 := r.FormValue("NRL03")
	NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL06 := r.FormValue("NRL06")
	NRL07 := r.FormValue("NRL07")
	NRL08 := r.FormValue("NRL08")
	NRL08A := r.FormValue("NRL08A")
	NRL08B := r.FormValue("NRL08B")
	NRL09A := r.FormValue("NRL09A")
	NRL09B := r.FormValue("NRL09B")
	score := r.FormValue("score")

	checkerr("nrl6 add", err1, err4, err5)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl7 := model.NRL7{
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
		NRL07:    NRL07,
		NRL08:    NRL08,
		NRL08A:   NRL08A,
		NRL08B:   NRL08B,
		NRL09A:   NRL09A,
		NRL09B:   NRL09B,
		Score:    score,
	}

	rid, err17 := nrl7.InsertData()
	if err17 != nil {
		fit.Logger().LogError("NRL7 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    7,
			NursingId:   BCE01A,
			NursingName: BCE03A,
			ClassId:     r.FormValue("did"),
			PatientId:   r.FormValue("pid"),
			RecordId:    rid,
			Comment:     "新增",
		}
		_,errRecord := model.InsertNRecords(nurseRecord)
		checkerr("nurse record err:", errRecord)

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{}
	}

}

// 修改护理记录单
func (c NRL7Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 科室ID
	BCK01, err5 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	NRL01 := r.FormValue("NRL01")
	NRL02 := r.FormValue("NRL02")
	NRL03 := r.FormValue("NRL03")
	NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL06 := r.FormValue("NRL06")
	NRL07 := r.FormValue("NRL07")
	NRL08 := r.FormValue("NRL08")
	NRL08A := r.FormValue("NRL08A")
	NRL08B := r.FormValue("NRL08B")
	NRL09A := r.FormValue("NRL09A")
	NRL09B := r.FormValue("NRL09B")
	score := r.FormValue("score")

	checkerr("nrl7 update", err1, err5, err4)

	nrl7 := model.NRL7{
		VAA01:    VAA01,
		BCK01:    BCK01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02:    NRL02,
		NRL03:    NRL03,
		NRL04:    NRL04,
		NRL05:    NRL05,
		NRL06:    NRL06,
		NRL07:    NRL07,
		NRL08:    NRL08,
		NRL08A:   NRL08A,
		NRL08B:   NRL08B,
		NRL09A:   NRL09A,
		NRL09B:   NRL09B,
		Score:    score,
	}

	_, err17 := nrl7.UpdateData(id)
	if err17 != nil {
		fit.Logger().LogError("nrl7 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "修改失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		_, errRecord := model.UpadteNRecords(id, r.FormValue("datetime"))
		checkerr("nurse record update err:", errRecord)

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "修改成功！"
		c.JsonData.Datas = []interface{}{}
	}
}

/*func (c NRL7Controller) InsertTitle(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 记录时间
	datetime, err2 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	// 科室ID
	BCK01, err3 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 文书ID（修改时需要）
	NRT01 := r.FormValue("NRT01")
	NRT01V := r.FormValue("NRT01V")
	NRT02 := r.FormValue("NRT02")
	NRT02V := r.FormValue("NRT02V")

	checkerr("nrl7 title:", err1, err2, err3)

	title := model.NRL7Title{
		VAA01:    VAA01,
		BCK01:    BCK01,
		DateTime: datetime,
		NRT01:    NRT01,
		NRT01V:   NRT01V,
		NRT02:    NRT02,
		NRT02V:   NRT02V,
	}
	errt := title.InsertTitle()
	if errt != nil {
		fit.Logger().LogError("Error", "nrl1 update :", errt)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = " 错误！"
		c.JsonData.Datas = []interface{}{errt.Error()}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "nrl7 title 添加成功！"
		c.JsonData.Datas = []interface{}{}
	}
}*/

func (c NRL7Controller) UpdateTitle(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 记录时间
	datetime, err2 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	// 科室ID
	BCK01, err3 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 文书ID（修改时需要）
	NRT01 := r.FormValue("NRT01")
	NRT01V := r.FormValue("NRT01V")
	NRT02 := r.FormValue("NRT02")
	NRT02V := r.FormValue("NRT02V")

	checkerr("nrl7 title:", err1, err2, err3)

	title := model.NRL7Title{
		VAA01:    VAA01,
		BCK01:    BCK01,
		DateTime: datetime,
		NRT01:    NRT01,
		NRT01V:   NRT01V,
		NRT02:    NRT02,
		NRT02V:   NRT02V,
	}
	errt := title.UpdateTitle()
	if errt != nil {
		fit.Logger().LogError("Error", "nrl1 update :", errt)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = " 错误！"
		c.JsonData.Datas = []interface{}{errt.Error()}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "nrl7 title 添加成功！"
		c.JsonData.Datas = []interface{}{}
	}
}
