package handler

import (
	"fit"
	"nursing/model"
	"time"
	"fmt"
	"strconv"
)

// 护理记录单 PDA端

// 模板 template
type NRL1Controller struct {
	NRLController
}

// 查看护理记录单
func (c NRL1Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr1, err1 := model.QueryNRL1(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR1", err1)
	}

	pid := nr1.PatientId
	uid := nr1.NurseId

	patintinfos, err := model.GetPatientInfo(pid)
	if err != nil {
		fit.Logger().LogError("m_NR1", err)
		fmt.Fprintln(w, "服务器有点繁忙！")
		return
	}
	if patintinfos == nil {
		fmt.Fprintln(w, "病人ID错误！")
		return
	}
	_, err1 = strconv.Atoi(uid)
	if err1 != nil {
		fit.Logger().LogError("m_NR1", err1)
		fmt.Fprintln(w, "用户ID错误！")
		return
	}

	nurse, _ := model.FetchAccountWithUid(uid)
	t := patintinfos[0].VAA1.VAA73
	timeStr := time.Time(t).Format("2006-01-02 15:04")

	rdate := nr1.DateTime.Format("2006-01-02")
	rtime := nr1.DateTime.Format("15:04")
	c.Data = fit.Data{
		"Patintinfos": patintinfos,
		"Time":        timeStr,
		"UserName":    nurse.Username,
		"NRL":         nr1,
		"RecordDate":  rdate,
		"RecordTime":  rtime,
	}
	c.LoadView(w, "v_nrl1.html")
}

//添加 编辑护理记录单
func (c NRL1Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")
	uid := r.FormValue("uid")
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr1 model.NRL1
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
		nr1, err1 = model.QueryNRL1(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = nr1.PatientId
		uid = nr1.NurseId
	}
	patintinfos, err := model.GetPatientInfo(pid)
	if err != nil {
		fmt.Fprintln(w, "服务器有点繁忙！")
		return
	}
	if patintinfos == nil {
		fmt.Fprintln(w, "病人ID错误！")
		return
	}
	_, err1 := strconv.Atoi(uid)
	if err1 != nil {
		fmt.Fprintln(w, "用户ID错误！")
		return
	}
	nurse, _ := model.FetchAccountWithUid(uid)
	t := patintinfos[0].VAA1.VAA73
	timeStr := time.Time(t).Format("2006-01-02 15:04")
	cid := strconv.Itoa(patintinfos[0].VAA1.BCK01C)

	rdate := time.Time(t).Format("2006-01-02")
	rtime := time.Time(t).Format("15:04")
	c.Data = fit.Data{
		"Patintinfos": patintinfos,
		"Time":        timeStr,
		"RecordDate":  rdate,
		"RecordTime":  rtime,
		"UserName":    nurse.Username,
		"Ty":          ty,
		"Rid":         rid,
		"NRL":         nr1,
		"Nid":         uid,
		"Pid":         pid,
		"Cid":         cid,
	}
	c.LoadView(w, "v_nrl1_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL1Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	//strconv.ParseInt(r.FormValue("nurse_id"), 10, 64)
	nurseId := r.FormValue("nurse_id")
	classId := r.FormValue("class_id")
	nurseName := r.FormValue("nurse_name")
	patientId := r.FormValue("patient_id")

	datetime, err1 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	temperature := r.FormValue("temperature")
	pulse := r.FormValue("pulse")
	heartrate := r.FormValue("heartrate")
	breathe := r.FormValue("breathe")
	preDia := r.FormValue("pre_dia")
	preSys := r.FormValue("pre_sys")
	vNRA5 := r.FormValue("NRA5")
	vNRA6A := r.FormValue("NRA6A")
	vNRA6B := r.FormValue("NRA6B")
	vNRA7A := r.FormValue("NRA7A")
	vNRA7B := r.FormValue("NRA7B")
	vNRA8 := r.FormValue("NRA8")
	vNRA9A := r.FormValue("NRA9A")
	vNRA9B := r.FormValue("NRA9B")
	vNRA10A := r.FormValue("NRA10A")
	vNRA10B := r.FormValue("NRA10B")
	vNRA11A := r.FormValue("NRA11A")
	vNRA11B := r.FormValue("NRA11B")
	vNRA12A := r.FormValue("NRA12A")
	vNRA12B := r.FormValue("NRA12B")
	vNRA13A := r.FormValue("NRA13A")
	vNRA13B := r.FormValue("NRA13B")
	vNRA14A := r.FormValue("NRA14A")
	vNRA14B := r.FormValue("NRA14B")
	vNRA15A := r.FormValue("NRA15A")
	vNRA15B := r.FormValue("NRA15B")

	vNRA16B := r.FormValue("NRA16B")

	if patientId == "" || nurseId == "" || nurseName == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	if err1 != nil {
		fit.Logger().LogError("error", err1.Error())
	}

	record := model.NRL1{
		PatientId:   patientId,
		ClassId:     classId,
		NurseName:   nurseName,
		NurseId:     nurseId,
		DateTime:    datetime,
		Temperature: temperature,
		Pulse:       pulse,
		Heartrate:   heartrate,
		Breathe:     breathe,
		PressureDIA: preDia,
		PressureSYS: preSys,
		NRA5:        vNRA5,
		NRA6A:       vNRA6A,
		NRA6B:       vNRA6B,
		NRA7A:       vNRA7A,
		NRA7B:       vNRA7B,
		NRA8:        vNRA8,
		NRA9A:       vNRA9A,
		NRA9B:       vNRA9B,
		NRA10A:      vNRA10A,
		NRA10B:      vNRA10B,
		NRA11A:      vNRA11A,
		NRA11B:      vNRA11B,
		NRA12A:      vNRA12A,
		NRA12B:      vNRA12B,
		NRA13A:      vNRA13A,
		NRA13B:      vNRA13B,
		NRA14A:      vNRA14A,
		NRA14B:      vNRA14B,
		NRA15A:      vNRA15A,
		NRA15B:      vNRA15B,
		NRA16B:      vNRA16B,
	}

	_, err := record.InsertData()

	fmt.Printf("%+v", record)
	// 文书记录
	nurseRecord := model.NursingRecords{
		Updated:     r.FormValue("datetime"),
		NursType:    1,
		NursingId:   nurseId,
		NursingName: nurseName,
		ClassId:     classId,
		PatientId:   patientId,
		RecordId:    record.Id,
		Comment:     "新增",
	}

	_, err2 := model.InsertNRecords(nurseRecord)

	if err != nil {
		fit.Logger().LogError("Error", "NRL1 add :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		if err2 != nil {
			fit.Logger().LogError("Error", "NRL1 add :", err2)
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "上传成功! (文书记录更新失败！)"
			c.JsonData.Datas = []interface{}{}
		}
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{}

	}
}

// 修改护理记录单
func (c NRL1Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	//strconv.ParseInt(r.FormValue("nurse_id"), 10, 64)
	rid := r.FormValue("rid")
	id, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		fit.Logger().LogError("Error", "nrl1 update :", err)
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "rid 错误！"
		c.JsonData.Datas = []interface{}{}
	}
	nurseId := r.FormValue("nurse_id")
	classId := r.FormValue("class_id")
	nurseName := r.FormValue("nurse_name")
	patientId := r.FormValue("patient_id")
	updatestr := r.FormValue("datetime")
	datetime, err1 := time.ParseInLocation("2006-01-02 15:04:05", updatestr, time.Local)
	temperature := r.FormValue("temperature")
	pulse := r.FormValue("pulse")
	heartrate := r.FormValue("heartrate")
	breathe := r.FormValue("breathe")
	preDia := r.FormValue("pre_dia")
	preSys := r.FormValue("pre_sys")
	vNRA5 := r.FormValue("NRA5")
	vNRA6A := r.FormValue("NRA6A")
	vNRA6B := r.FormValue("NRA6B")
	vNRA7A := r.FormValue("NRA7A")
	vNRA7B := r.FormValue("NRA7B")
	vNRA8 := r.FormValue("NRA8")
	vNRA9A := r.FormValue("NRA9A")
	vNRA9B := r.FormValue("NRA9B")
	vNRA10A := r.FormValue("NRA10A")
	vNRA10B := r.FormValue("NRA10B")
	vNRA11A := r.FormValue("NRA11A")
	vNRA11B := r.FormValue("NRA11B")
	vNRA12A := r.FormValue("NRA12A")
	vNRA12B := r.FormValue("NRA12B")
	vNRA13A := r.FormValue("NRA13A")
	vNRA13B := r.FormValue("NRA13B")
	vNRA14A := r.FormValue("NRA14A")
	vNRA14B := r.FormValue("NRA14B")
	vNRA15A := r.FormValue("NRA15A")
	vNRA15B := r.FormValue("NRA15B")
	vNRA16B := r.FormValue("NRA16B")

	if rid == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	if err1 != nil {
		fit.Logger().LogError("error", err1.Error())
	}

	record := model.NRL1{
		PatientId:   patientId,
		ClassId:     classId,
		NurseName:   nurseName,
		NurseId:     nurseId,
		DateTime:    datetime,
		Temperature: temperature,
		Pulse:       pulse,
		Heartrate:   heartrate,
		Breathe:     breathe,
		PressureDIA: preDia,
		PressureSYS: preSys,
		NRA5:        vNRA5,
		NRA6A:       vNRA6A,
		NRA6B:       vNRA6B,
		NRA7A:       vNRA7A,
		NRA7B:       vNRA7B,
		NRA8:        vNRA8,
		NRA9A:       vNRA9A,
		NRA9B:       vNRA9B,
		NRA10A:      vNRA10A,
		NRA10B:      vNRA10B,
		NRA11A:      vNRA11A,
		NRA11B:      vNRA11B,
		NRA12A:      vNRA12A,
		NRA12B:      vNRA12B,
		NRA13A:      vNRA13A,
		NRA13B:      vNRA13B,
		NRA14A:      vNRA14A,
		NRA14B:      vNRA14B,
		NRA15A:      vNRA15A,
		NRA15B:      vNRA15B,
		NRA16B:      vNRA16B,
	}

	_, err3 := record.UpdateData(id)
	// 文书记录


	_, err2 := model.UpadteNRecords(id, updatestr)
	if err3 != nil {
		fit.Logger().LogError("Error", "warn add :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "修改失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		if err2 != nil {
			fit.Logger().LogError("Error", "NRL add :", err2)
			c.JsonData.Result = 0
			c.JsonData.ErrorMsg = "修改成功! (文书记录更新失败！)"
			c.JsonData.Datas = []interface{}{}
		}
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "修改成功！"
		c.JsonData.Datas = []interface{}{}
	}
}
