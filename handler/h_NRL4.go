package handler

import (
	"fit"
	"fmt"
	"strconv"
	"time"
	"nursing/model"
)

type NRL4Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL4Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr4, err1 := model.QueryNRL4(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR4", err1)
	}

	pid := nr4.VAA01

	// 查询对应病人信息
	patient, has := c.LoadPInfoWithPid(w, r, pid)
	if !has {
		return
	}

	recordDate := nr4.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"PInfo":      patient,
		"NRL":        nr4,
		"RecordDate": recordDate,
	}

	fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "v_nrl4.html")
}

func (c NRL4Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	/*pid := r.FormValue("pid")
	uid := r.FormValue("uid")
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit
	var nr4 model.NRL4
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
		nr4, err1 = model.QueryNRL4(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr4.VAA01, 10)
		uid = nr4.BCE01A
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}*/
	// 文书model，type， 文书id，病人id，护士id，参数是否正确
	nrl, ty, rid, pid, uid, isOk := c.LoadNRLDataWithParm(w, r, "4")
	if !isOk {
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPInfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}

	c.Data = fit.Data{
		"PInfo":      patient,
		"NRL":        nrl,
		"Type":       ty,
		"Rid":        rid,
		"Account":    account,
	}

	c.LoadView(w, "v_nrl4_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL4Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	did := r.FormValue("did")
	BCK01, err5 := strconv.Atoi(did)

	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02, err7 := strconv.Atoi(r.FormValue("NRL02"))
	NRL03, err8 := strconv.Atoi(r.FormValue("NRL03"))
	NRL04, err9 := strconv.Atoi(r.FormValue("NRL04"))
	NRL05, err10 := strconv.Atoi(r.FormValue("NRL05"))
	NRL06, err11 := strconv.Atoi(r.FormValue("NRL06"))
	NRL07, err12 := strconv.Atoi(r.FormValue("NRL07"))
	score := r.FormValue("score")

	checkerr("nrl4 add", err1, err4, err5, err6, err7, err8, err9, err10, err11, err12)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl4 := model.NRL4{
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
		Score:    score,
	}

	rid, err17 := nrl4.InsertData()

	if err17 != nil {
		fit.Logger().LogError("NRL4 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    4,
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
func (c NRL4Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormValue("rid")
	id, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		fit.Logger().LogError("Error", "nrl1 update :", err)
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "rid 错误！"
		c.JsonData.Datas = []interface{}{}
		return
	}
	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 科室ID
	did := r.FormValue("did")
	BCK01, err5 := strconv.Atoi(did)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02, err7 := strconv.Atoi(r.FormValue("NRL02"))
	NRL03, err8 := strconv.Atoi(r.FormValue("NRL03"))
	NRL04, err9 := strconv.Atoi(r.FormValue("NRL04"))
	NRL05, err10 := strconv.Atoi(r.FormValue("NRL05"))
	NRL06, err11 := strconv.Atoi(r.FormValue("NRL06"))
	NRL07, err12 := strconv.Atoi(r.FormValue("NRL07"))
	score := r.FormValue("score")

	checkerr("nrl4 add",err1, err5, err4, err6, err7, err8, err9, err10, err11, err12)

	nrl4 := model.NRL4{
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
		Score:    score,
	}

	_, err17 := nrl4.UpdateData(id)


	if err17 != nil {
		fit.Logger().LogError("nrl4 add :", err17)
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

// 删除护理
func (c NRL4Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormValue("rid")
	if rid == "" {
		c.RenderingJsonAutomatically(3, "参数不完整")
	}
	id, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		fit.Logger().LogError("Error", "nrl1 update :", err)
		c.RenderingJsonAutomatically(3, "rid 错误！")
		return
	}


	nrl := model.NRL4{}

	_, err18 := nrl.DeleteData(id)
	if err18 != nil {
		fit.Logger().LogError("nrl delete :", err18)
		c.RenderingJsonAutomatically(3, "删除失败！")
	} else {
		c.RenderingJsonAutomatically(0, "删除成功！")
	}
}
