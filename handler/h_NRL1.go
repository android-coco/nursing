package handler

import (
	"fit"
	"strconv"
	"nursing/model"
	"fmt"
	"time"
	"encoding/json"
)

// 护理记录单 PDA端

// 模板 template
type NRL1Controller struct {
	NRLController
}

// 修改护理记录单
func (c NRL1Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 病人ID
	pid := r.FormValue("pid")
	VAA01, err1 := strconv.ParseInt(pid, 10, 64)
	// 科室ID
	//did := r.FormValue("did")
	//BCK01, err2 := strconv.Atoi(did)
	// 护士ID
	BCE01A, err3 := strconv.Atoi(r.FormValue("uid"))
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	flag := checkerr("update record:", err1, err3, err4)
	if flag || BCE03A == ""{
		c.RenderingJsonAutomatically(2, "参数不完整")
		return
	}
	//fmt.Println(VAA01, BCK01, BCE01A, BCE03A, datetime)

	jsonstr := r.FormValue("jsonstr")
	var mods []model.NRLData
	errUnmarshal := json.Unmarshal([]byte(jsonstr), &mods)
	if errUnmarshal != nil {
		fit.Logger().LogError("json.Unmarshal err:", errUnmarshal)
		c.RenderingJsonAutomatically(2, "参数错误")
		return
	}

	//开启事务
	session := fit.MySqlEngine().NewSession()
	defer session.Close()
	errsession := session.Begin()
	if errsession != nil {
		fit.Logger().LogError("session err:", errsession)
		c.RenderingJsonAutomatically(1, "修改失败", )
		return
	}
	for _, mod := range mods {

		id := mod.ID
		if id == 0 {
			mod.TestTime = fit.JsonTime(datetime)
			mod.PatientId = VAA01
			mod.NurseId = BCE01A
			mod.NurseName = BCE03A

			_, errInsert := session.Table("NurseChat").Insert(mod)
			if errInsert != nil {
				fit.Logger().LogError("session insert err:", errInsert)
				c.RenderingJsonAutomatically(1, "修改失败", )
				session.Rollback()
				return
			}
		} else {
			_, errUpdate := session.Table("NurseChat").ID(id).Update(&mod)
			if errUpdate != nil {
				fit.Logger().LogError("session update err:", errUpdate)
				c.RenderingJsonAutomatically(1, "修改失败", )
				session.Rollback()
				return
			}
		}
		//fmt.Printf("---------------model: %+v\n\n", model)
	}
	errsession = session.Commit()
	if errsession != nil {
		fit.Logger().LogError("session Commit err:", errsession)
		c.RenderingJsonAutomatically(1, "修改失败", )
		return
	} else {
		c.RenderingJsonAutomatically(0, "修改成功", )
	}

	//jsonbyte, err12 := json.Marshal(maps)
	//fmt.Println("jsonmap :", string(jsonbyte), err12)

}

func (c NRL1Controller) UpdateTitle(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 科室ID
	BCK01, err3 := strconv.Atoi(r.FormValue("did"))

	// 文书ID（修改时需要）
	NRT01 := r.FormValue("NRT01")
	NRT01V := r.FormValue("NRT01V")
	NRT02 := r.FormValue("NRT02")
	NRT02V := r.FormValue("NRT02V")
	NRT03 := r.FormValue("NRT03")
	NRT03V := r.FormValue("NRT03V")
	NRT04 := r.FormValue("NRT04")
	NRT04V := r.FormValue("NRT04V")
	NRT05 := r.FormValue("NRT05")
	NRT05V := r.FormValue("NRT05V")
	NRT06 := r.FormValue("NRT06")
	NRT06V := r.FormValue("NRT06V")
	NRT07 := r.FormValue("NRT07")
	NRT07V := r.FormValue("NRT07V")


	if flag := checkerr("nrl7 title:", err1, err3); flag {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	title := model.NRL1Title{
		VAA01:  VAA01,
		BCK01:  BCK01,
		NRT01:  NRT01,
		NRT01V: NRT01V,
		NRT02:  NRT02,
		NRT02V: NRT02V,
		NRT03:  NRT03,
		NRT03V: NRT03V,
		NRT04:  NRT04,
		NRT04V: NRT04V,
		NRT05:  NRT05,
		NRT05V: NRT05V,
		NRT06:  NRT06,
		NRT06V: NRT06V,
		NRT07:  NRT07,
		NRT07V: NRT07V,
	}
	errt := title.PCUpdateNRT1Title()
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

// 删除护理
func (c NRL1Controller) DeleteRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 文书ID

	jsonstr := r.FormValue("jsonstr")
	var mods []model.NRLData
	errUnmarshal := json.Unmarshal([]byte(jsonstr), &mods)
	if errUnmarshal != nil {
		fit.Logger().LogError("json.Unmarshal err:", errUnmarshal)
		c.RenderingJsonAutomatically(4, "参数不完整")
		return
	}

	//开启事务
	session := fit.MySqlEngine().NewSession()
	defer session.Close()
	errsession := session.Begin()
	if errsession != nil {
		fit.Logger().LogError("session err:", errsession)
		c.RenderingJsonAutomatically(4, "参数错误")
		return
	}


	for _, mod := range mods {
		fmt.Printf("model: %+v\n\n", mod)
		id := mod.ID
		if id == 0 {
			c.RenderingJsonAutomatically(3, "删除失败 id为空", )
			session.Rollback()
			return
		} else {
			if mod.HeadType == "18" {
				modio := model.IOStatistics{ID:id}

				affected, errUpdate := session.Table("IOStatistics").ID(id).Delete(&modio)
				if errUpdate != nil {
					fit.Logger().LogError("session update err:", errUpdate)
					c.RenderingJsonAutomatically(1, "删除失败", )
					session.Rollback()
					return
				}
				if affected == 0 {
					c.RenderingJsonAutomatically(2, "删除失败, 不存在改条记录", )
					session.Rollback()
					return
				}
			} else {
				affected, errUpdate := session.Table("NurseChat").ID(id).Delete(&mod)
				if errUpdate != nil {
					fit.Logger().LogError("session update err:", errUpdate)
					c.RenderingJsonAutomatically(1, "删除失败", )
					session.Rollback()
					return
				}
				if affected == 0 {
					c.RenderingJsonAutomatically(2, "删除失败, 不存在改条记录", )
					session.Rollback()
					return
				}
			}

			switch mod.HeadType {
			case "18":
			case "1", "2", "3", "4", "6", "7", "15", "16", "17":
			default:
			}

		}
	}
	errsession = session.Commit()
	if errsession != nil {
		fit.Logger().LogError("session Commit err:", errsession)
		c.RenderingJsonAutomatically(1, "删除失败", )
	} else {
		c.RenderingJsonAutomatically(0, "删除成功", )
	}
}

/*
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


*/
