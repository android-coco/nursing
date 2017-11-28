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
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}


	nr3, err1 := model.QueryNRL3(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR3", err1)
	}

	pid := nr3.VAA01
	// 查询对应病人信息
	patient, has := c.LoadPInfoWithPid(w, r, pid)
	if !has {
		return
	}

	recordDate := nr3.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"PInfo": patient,
		"NRL":   nr3,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl3.html")
}

func (c NRL3Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")//病人id
	uid := r.FormValue("uid") //护士id
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr3 model.NRL3
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
		nr3, err1 = model.QueryNRL3(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr3.VAA01, 10)
		uid = nr3.BCE01A
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPInfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}


	recordDate := nr3.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"PInfo": patient,
		"NRL":   nr3,
		"Type":  ty,
		"Rid": rid,
		"Account": account,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl3_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL3Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	// 文书ID（修改时需要）
	//recordId := r.FormValue("rid")
	NRL01, err6 := strconv.Atoi(r.FormValue("NRL01"))
	NRL02, err7 := strconv.Atoi(r.FormValue("NRL02"))
	NRL03, err8 := strconv.Atoi(r.FormValue("NRL03"))
	NRL04, err9 := strconv.Atoi(r.FormValue("NRL04"))
	NRL05, err10 := strconv.Atoi(r.FormValue("NRL05"))
	NRL06, err11 := strconv.Atoi(r.FormValue("NRL06"))
	NRL07, err12 := strconv.Atoi(r.FormValue("NRL07"))
	NRL08, err13 := strconv.Atoi(r.FormValue("NRL08"))
	NRL09, err14 := strconv.Atoi(r.FormValue("NRL09"))
	NRL10, err15 := strconv.Atoi(r.FormValue("NRL10"))
	NRL11, err16 := strconv.Atoi(r.FormValue("NRL11"))
	score := r.FormValue("score")

	checkerr("nrl3 add", err1, err4, err5, err6, err7, err8, err9, err10, err11, err12, err13, err14, err15, err16)


	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl3 := model.NRL3{
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
		NRL09:    NRL09,
		NRL10:    NRL10,
		NRL11:    NRL11,
		Score:    score,
	}

	rid, err17 := nrl3.InsertData()


	if err17 != nil {
		fit.Logger().LogError("NRL3 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    3,
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
func (c NRL3Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	NRL08, err13 := strconv.Atoi(r.FormValue("NRL08"))
	NRL09, err14 := strconv.Atoi(r.FormValue("NRL09"))
	NRL10, err15 := strconv.Atoi(r.FormValue("NRL10"))
	NRL11, err16 := strconv.Atoi(r.FormValue("NRL11"))
	score := r.FormValue("score")

	checkerr("nrl3 add",err1, err5, err4, err6, err7, err8, err9, err10, err11, err12, err13, err14, err15, err16)

	nrl3 := model.NRL3{
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
		NRL09:    NRL09,
		NRL10:    NRL10,
		NRL11:    NRL11,
		Score:    score,

	}

	_, err17 := nrl3.UpdateData(id)

	if err17 != nil {
		fit.Logger().LogError("nrl3 add :", err17)
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
func (c NRL3Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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


	nrl := model.NRL3{}

	_, err18 := nrl.DeleteData(id)
	if err18 != nil {
		fit.Logger().LogError("nrl delete :", err18)
		c.RenderingJsonAutomatically(3, "删除失败！")
	} else {
		c.RenderingJsonAutomatically(0, "删除成功！")
	}
}