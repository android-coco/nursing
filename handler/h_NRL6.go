package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"
)

type NRL6Controller struct {
	NRLController
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

	// 查询对应病人信息
	patient, has := c.LoadPInfoWithPid(w, r, pid)
	if !has {
		return
	}

	recordDate := nr6.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"PInfo": patient,
		"NRL":   nr6,
		"RecordDate": recordDate,
	}

	fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "v_nrl6.html")
}

func (c NRL6Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	/*pid := r.FormValue("pid")//病人id
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
	}*/
	// 文书model，type， 文书id，病人id，护士id，参数是否正确
	nrl, ty, rid, pid, uid, isOk := c.LoadNRLDataWithParm(w, r, "6")
	if !isOk {
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPInfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}


	c.Data = fit.Data{
		"PInfo": patient,
		"NRL":   nrl,
		"Type":  ty,
		"Rid": rid,
		"Account": account,
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
	NRL01 := r.FormValue("NRL01")
	NRL02 := r.FormValue("NRL02")
	NRL03 := r.FormValue("NRL03")
	NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL06 := r.FormValue("NRL06")
	NRL07 := r.FormValue("NRL07")
	NRL08 := r.FormValue("NRL08")
	NRL09 := r.FormValue("NRL09")
	NRL10 := r.FormValue("NRL10")
	NRL11 := r.FormValue("NRL11")
	NRL12 := r.FormValue("NRL12")
	NRL13 := r.FormValue("NRL13")
	//NRL14 := r.FormValue("NRL14")
	NRL15A := r.FormValue("NRL15A")
	NRL15B := r.FormValue("NRL15B")
	score := r.FormValue("score")


	checkerr("nrl6 add", err1, err4, err5)

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
		NRL07:    NRL07,
		NRL08:    NRL08,
		NRL09:    NRL09,
		NRL10:    NRL10,
		NRL11:    NRL11,
		NRL12:    NRL12,
		NRL13:    NRL13,
		//NRL14:    NRL14,
		NRL15A:    NRL15A,
		NRL15B:    NRL15B,
		Score:    score,
	}

	rid, err17 := nrl6.InsertData()

	if err17 != nil {
		fit.Logger().LogError("NRL6 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    6,
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

	NRL01 := r.FormValue("NRL01")
	NRL02 := r.FormValue("NRL02")
	NRL03 := r.FormValue("NRL03")
	NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL06 := r.FormValue("NRL06")
	NRL07 := r.FormValue("NRL07")
	NRL08 := r.FormValue("NRL08")
	NRL09 := r.FormValue("NRL09")
	NRL10 := r.FormValue("NRL10")
	NRL11 := r.FormValue("NRL11")
	NRL12 := r.FormValue("NRL12")
	NRL13 := r.FormValue("NRL13")
	//NRL14 := r.FormValue("NRL14")
	NRL15A := r.FormValue("NRL15A")
	NRL15B := r.FormValue("NRL15B")
	score := r.FormValue("score")

	checkerr("nrl6 update",err1, err5, err4)

	nrl6 := model.NRL6{
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
		NRL12:    NRL12,
		NRL13:    NRL13,
		//NRL14:    NRL14,
		NRL15A:    NRL15A,
		NRL15B:    NRL15B,
		Score:    score,
	}

	_, err17 := nrl6.UpdateData(id)
	if err17 != nil {
		fit.Logger().LogError("nrl6 add :", err17)
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
func (c NRL6Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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


	nrl := model.NRL6{}

	_, err18 := nrl.DeleteData(id)
	if err18 != nil {
		fit.Logger().LogError("nrl8 add :", err18)
		c.RenderingJsonAutomatically(3, "删除失败！")
	} else {
		c.RenderingJsonAutomatically(0, "删除成功！")
	}
}
