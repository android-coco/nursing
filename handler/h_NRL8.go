package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"
)

type NRL8Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL8Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr8, err1 := model.QueryNRL8(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR8", err1)
	}

	pid := nr8.VAA01

	// 查询对应病人信息
	patient, has := c.LoadPinfoWithPid(w, r, pid)
	if !has {
		return
	}

	recordDate := nr8.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo": patient,
		"NRL":   nr8,
		"RecordDate": recordDate,
	}

	fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "v_nrl8.html")
}

func (c NRL8Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")//病人id
	uid := r.FormValue("uid") //护士id
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr8 model.NRL8
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
		nr8, err1 = model.QueryNRL8(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr8.VAA01, 10)
		uid = nr8.BCE01A
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPinfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}

	recordDate := nr8.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo": patient,
		"NRL":   nr8,
		"Type":  ty,
		"Rid": rid,
		"Account": account,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl8_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL8Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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

	NRL05 := r.FormValue("NRL05")
	NRL05A := r.FormValue("NRL05A")
	NRL06A := r.FormValue("NRL06A")
	NRL06B := r.FormValue("NRL06B")
	score := r.FormValue("score")


	checkerr("nrl8 add", err1, err4, err5)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl8 := model.NRL8{
		BCK01:    BCK01,
		VAA01:    VAA01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		NRL05:    NRL05,
		NRL05A:    NRL05A,
		NRL06A:    NRL06A,
		NRL06B:    NRL06B,
		Score:    score,
	}

	rid, err18 := nrl8.InsertData()
	if err18 != nil {
		fit.Logger().LogError("NRL8 add :", err18)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    8,
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
func (c NRL8Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
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
	BCE01A := r.FormValue("nid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	//NRL01:= r.FormValue("NRL01")
	//NRL02 := r.FormValue("NRL02")
	//NRL03 := r.FormValue("NRL03")
	//NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL05A := r.FormValue("NRL05A")
	NRL06A := r.FormValue("NRL06A")
	NRL06B := r.FormValue("NRL06B")
	score := r.FormValue("score")

	checkerr("nrl8 update", err1, err5, err4)

	nrl8 := model.NRL8{
		VAA01:    VAA01,
		BCK01:    BCK01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		//NRL01:    NRL01,
		//NRL02:    NRL02,
		//NRL03:    NRL03,
		//NRL04:    NRL04,
		NRL05:    NRL05,
		NRL05A:    NRL05A,
		NRL06A:    NRL06A,
		NRL06B:    NRL06B,
		Score:    score,
	}

	_, err18 := nrl8.UpdateData(id)
	if err18 != nil {
		fit.Logger().LogError("nrl8 add :", err18)
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


