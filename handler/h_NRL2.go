//  Created by JP

package handler

import (
	"fit"
	"strconv"
	"time"
	"nursing/model"
	"fmt"
	"encoding/json"
	"strings"
)

type NRL2Controller struct {
	NRLController
}

// 模板 template PDA端
func (c NRL2Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr2, err1 := model.QueryNRL2(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR2", err1)
		fmt.Fprintln(w, "无法查询到相应护理单记录", err1.Error())
		return
	}

	pid := nr2.PatientId
	// 查询对应病人信息
	patient, has := c.LoadPInfoWithPid(w, r, pid)
	if !has {
		return
	}

	// 过敏史-过敏源-NRL06-Json字符串，key：对应index，value：对应内容
	NRL06A := ""
	NRL06B := ""
	if nr2.NRL06 == "2" {
		by := []byte(nr2.NRL06A)
		anyObj := make(map[string]string, 0)
		json.Unmarshal(by, &anyObj)
		for k, v := range anyObj {
			NRL06A = k
			NRL06B = v
		}
	}

	// 排便次数- "n,m"，n,m代表2个空的数值，即n次/天，1次/m天
	NRL18A := ""
	NRL18B := ""
	if nr2.NRL18 != "" {
		slice := strings.Split(nr2.NRL18, ",")
		length := len(slice)
		if length == 1 {
			NRL18A = slice[0]
		} else if length == 2 {
			NRL18A = slice[0]
			NRL18B = slice[1]
		}
	}

	// 入院时间
	VAA73 := time.Time(patient.VAE11).Format("2006-01-02 15:04")
	// 拆分护理单录入时间
	tempTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nr2.NRL38, time.Local)
	NRL38A := tempTime.Format("2006-01-02")
	NRL38B := tempTime.Format("15:04")
	//审核时间
	tempTime2, _ := time.ParseInLocation("2006-01-02 15:04:05", nr2.NRL39A, time.Local)
	NRL39B := tempTime2.Format("2006-01-02")
	NRL39C := tempTime2.Format("15:04")
	c.Data = fit.Data{
		"PInfo":  patient, // 病人数据
		"NRL":    nr2,     // 护理单数据
		"VAA73":  VAA73,   // 入院时间
		"NRL06A": NRL06A,  // 过敏源的index
		"NRL06B": NRL06B,  // 过敏源的补充内容
		"NRL18A": NRL18A,  // n次/天
		"NRL18B": NRL18B,  // 1次/m天
		"NRL38A": NRL38A,  // 录入护理单的年月日
		"NRL38B": NRL38B,  // 录入护理单的时分
		"NRL39B": NRL39B,  //审核时间
		"NRL39C": NRL39C,
	}

	//fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "pda/v_nrl2.html")
}

func (c NRL2Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid") //病人id
	uid := r.FormValue("uid") //护士id
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit
	var nr2 model.NRL2
	if ty == "1" {
		if "" == pid || "" == uid {
			fmt.Fprintln(w, "参数错误！1")
			return
		}
	} else if ty == "2" {
		if rid == "" {
			fmt.Fprintln(w, "参数错误！2")
			return
		}
		var err1 error
		nr2, err1 = model.QueryNRL2(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr2.PatientId, 10)
		uid = nr2.NurseId
	} else {
		fmt.Fprintln(w, "参数错误！3")
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPInfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}

	// 入院时间
	VAA73 := time.Time(patient.VAE11).Format("2006-01-02 15:04")
	//fmt.Printf("account %+v \n\n %+v\n\n", account, pinfo)
	var NRL06A, NRL06B, NRL18A, NRL18B, NRL38A, NRL38B, NRL39B, NRL39C string
	if ty == "2" {
		// 过敏史-过敏源-NRL06-Json字符串，key：对应index，value：对应内容
		if nr2.NRL06 == "2" {
			by := []byte(nr2.NRL06A)
			anyObj := make(map[string]string, 0)
			json.Unmarshal(by, &anyObj)
			for k, v := range anyObj {
				NRL06A = k
				NRL06B = v
			}
		}

		// 排便次数- "n,m"，n,m代表2个空的数值，即n次/天，1次/m天
		if nr2.NRL18 != "" {
			slice := strings.Split(nr2.NRL18, ",")
			length := len(slice)
			if length == 1 {
				NRL18A = slice[0]
			} else if length == 2 {
				NRL18A = slice[0]
				NRL18B = slice[1]
			}
		}
		// 拆分护理单录入时间
		tempTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nr2.NRL38, time.Local)
		NRL38A = tempTime.Format("2006-01-02")
		NRL38B = tempTime.Format("15:04")

		//审核时间
		tempTime2, _ := time.ParseInLocation("2006-01-02 15:04:05", nr2.NRL39A, time.Local)
		NRL39B = tempTime2.Format("2006-01-02")
		NRL39C = tempTime2.Format("15:04")
	}

	c.Data = fit.Data{
		"Type":    ty,
		"Rid":     rid,
		"Account": account,
		"PInfo":   patient, // 病人数据
		"NRL":     nr2,     // 护理单数据
		"VAA73":   VAA73,   // 入院时间

		"NRL06A": NRL06A, // 过敏源的index
		"NRL06B": NRL06B, // 过敏源的补充内容
		"NRL18A": NRL18A, // n次/天
		"NRL18B": NRL18B, // 1次/m天
		"NRL38A": NRL38A, // 录入护理单的年月日
		"NRL38B": NRL38B, // 录入护理单的时分
		"NRL39B": NRL39B, //审核时间
		"NRL39C": NRL39C,
	}

	c.LoadView(w, "pda/v_nrl2_edit.html")
}

// 接口
// 添加、修改护理记录单
func (c NRL2Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	pid := r.FormValue("pid")
	// 病人ID
	VAA01, err1 := strconv.ParseInt(pid, 10, 64)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	NRL38 := r.FormValue("datetime")
	//NRL38, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	// 科室ID
	did := r.FormValue("did")
	BCK01, err2 := strconv.Atoi(did)

	// 文书ID（修改时需要）
	recordId := r.FormValue("rid")

	hasErr := checkerr("nrl2 add", err1, err2)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" || hasErr == true {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	NRL02 := r.FormValue("NRL02")
	NRL03 := r.FormValue("NRL03")
	NRL04 := r.FormValue("NRL04")
	NRL05 := r.FormValue("NRL05")
	NRL06 := r.FormValue("NRL06")
	NRL06A := r.FormValue("NRL06A")
	NRL07 := r.FormValue("NRL07")
	NRL07A := r.FormValue("NRL07A")
	NRL08 := r.FormValue("NRL08")
	NRL09 := r.FormValue("NRL09")

	NRL10 := r.FormValue("NRL10")
	NRL11 := r.FormValue("NRL11")
	NRL12 := r.FormValue("NRL12")
	NRL12A := r.FormValue("NRL12A")
	NRL13 := r.FormValue("NRL13")
	NRL14 := r.FormValue("NRL14")
	NRL14A := r.FormValue("NRL14A")
	NRL15 := r.FormValue("NRL15")
	NRL16 := r.FormValue("NRL16")
	NRL16A := r.FormValue("NRL16A")
	NRL17 := r.FormValue("NRL17")
	NRL17A := r.FormValue("NRL17A")
	NRL18 := r.FormValue("NRL18")
	NRL19 := r.FormValue("NRL19")
	LimbsId := r.FormValue("LimbsId")

	NRL20 := r.FormValue("NRL20")
	NRL20A := r.FormValue("NRL20A")
	NRL21 := r.FormValue("NRL21")
	NRL22 := r.FormValue("NRL22")
	NRL23 := r.FormValue("NRL23")
	NRL24 := r.FormValue("NRL24")
	NRL25 := r.FormValue("NRL25")
	NRL26 := r.FormValue("NRL26")
	NRL27 := r.FormValue("NRL27")
	NRL28 := r.FormValue("NRL28")
	NRL29 := r.FormValue("NRL29")

	NRL30 := r.FormValue("NRL30")
	NRL30A := r.FormValue("NRL30A")
	NRL31 := r.FormValue("NRL31")
	NRL32 := r.FormValue("NRL32")
	NRL33 := r.FormValue("NRL33")
	NRL34 := r.FormValue("NRL34")
	NRL35 := r.FormValue("NRL35")
	NRL36 := r.FormValue("NRL36")
	NRL37 := r.FormValue("NRL37")
	NRL39A := r.FormValue("NRL39A")
	NRL39B := r.FormValue("NRL39B")
	NRL39C := r.FormValue("NRL39C")
	NRL40 := r.FormValue("NRL40")
	NRL41 := r.FormValue("NRL41")

	nrl2 := model.NRL2{
		PatientId: VAA01,
		BCK01:     BCK01,
		NRL02:     NRL02,
		NRL03:     NRL03,
		NRL04:     NRL04,
		NRL05:     NRL05,
		NRL06:     NRL06,
		NRL06A:    NRL06A,
		NRL07:     NRL07,
		NRL07A:    NRL07A,
		NRL08:     NRL08,
		NRL09:     NRL09,
		NRL10:     NRL10,
		NRL11:     NRL11,
		NRL12:     NRL12,
		NRL12A:    NRL12A,
		NRL13:     NRL13,
		NRL14:     NRL14,
		NRL14A:    NRL14A,
		NRL15:     NRL15,
		NRL16:     NRL16,
		NRL16A:    NRL16A,
		NRL17:     NRL17,
		NRL17A:    NRL17A,
		NRL18:     NRL18,
		NRL19:     NRL19,
		LimbsId:   LimbsId,
		NRL20:     NRL20,
		NRL20A:    NRL20A,
		NRL21:     NRL21,
		NRL22:     NRL22,
		NRL23:     NRL23,
		NRL24:     NRL24,
		NRL25:     NRL25,
		NRL26:     NRL26,
		NRL27:     NRL27,
		NRL28:     NRL28,
		NRL29:     NRL29,
		NRL30:     NRL30,
		NRL30A:    NRL30A,
		NRL31:     NRL31,
		NRL32:     NRL32,
		NRL33:     NRL33,
		NRL34:     NRL34,
		NRL35:     NRL35,
		NRL36:     NRL36,
		NRL37:     NRL37,
		NRL38:     NRL38,
		NurseId:   BCE01A,
		NurseName: BCE03A,
		NRL39A:    NRL39A,
		NRL39B:    NRL39B,
		NRL39C:    NRL39C,
		NRL40:     NRL40,
		NRL41:     NRL41,
	}

	if recordId == "" { // 添加
		isExist, _ := model.IsExistNRL2(pid)
		if isExist == true {
			c.RenderingJsonAutomatically(4, "已存在该病人的首次护理单，请勿重复添加")
			return
		}

		newNrlId, err17 := nrl2.InsertToDatabase()
		if err17 != nil {
			fit.Logger().LogError("NRL2 add :", err17)
			c.RenderingJsonAutomatically(3, "上传失败！ "+err17.Error())
			return
		}

		// 在护理单记录总表中插入一条记录
		// timeNow := time.Now().Format("2006-01-02 15:04:05")
		nr := model.NursingRecords{
			Updated:     NRL38,
			NursType:    2,
			NursingId:   BCE01A,
			NursingName: BCE03A,
			ClassId:     did,
			PatientId:   pid,
			RecordId:    int64(newNrlId),
			Comment:     "新增",
		}
		_, err := model.InsertNRecords(nr)
		if err != nil {
			c.RenderingJsonAutomatically(0, "Database: "+err.Error())
		} else {
			c.RenderingJsonAutomatically(0, "添加成功")
		}

	} else { // 修改
		recordIdInt, errNrlId := strconv.ParseInt(recordId, 10, 64)
		if errNrlId != nil {
			c.RenderingJsonAutomatically(2, "参数错误 nrl_id")
			return
		}
		// 修改数据库
		err := nrl2.UpdateRecordDatas(recordIdInt)

		// 在护理单记录总表更新修改记录
		_, err = model.UpadteNRecords(recordIdInt, NRL38)

		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			c.RenderingJsonAutomatically(0, "修改成功")
		}
	}

}

func (c NRL2Controller) Exist(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	pid := r.FormValue("pid")
	if pid == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	patient_id, err_p := strconv.Atoi(pid)
	if err_p != nil || patient_id == 0 {
		c.RenderingJsonAutomatically(2, "参数错误 pid")
		return
	}
	isExist, err_exs := model.IsExistNRL2(pid)
	if err_exs != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_exs.Error())
		return
	}

	if isExist == true {
		c.RenderingJsonAutomatically(0, "已存在该病人的首次护理单")
	} else {
		c.RenderingJsonAutomatically(4, "不存在该病人的首次护理单")
	}
}

type PCNRL2Controller struct {
	PCNRLController
}

func (c PCNRL2Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {

	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "2")
	if !has {
		return
	}
	var err error

	// 护理单
	flag, errExist := model.IsExistNRL2(pid)
	if errExist != nil {
		fit.Logger().LogError("pc nr2 is exist nrl2?", errExist)
		fmt.Fprintln(w, "参数错误！  user info error", errExist)
		return
	}

	var nrl2 model.NRL2
	if !flag {
		nrl2 = model.NRL2{}
	} else {
		var errn2 error
		nrl2, errn2 = model.QueryNRL2WithPid(pid)
		if errn2 != nil {
			fit.Logger().LogError("pc nr2 query nrl2:", err)
			fmt.Fprintln(w, "参数错误！  user info error", err)
			return
		}
	}
	// 过敏史-过敏源-NRL06-Json字符串，key：对应index，value：对应内容
	NRL06A := ""
	NRL06B := ""
	if nrl2.NRL06 == "2" {
		by := []byte(nrl2.NRL06A)
		anyObj := make(map[string]string, 0)
		json.Unmarshal(by, &anyObj)
		for k, v := range anyObj {
			NRL06A = k
			NRL06B = v
		}
	}

	// 排便次数- "n,m"，n,m代表2个空的数值，即n次/天，1次/m天
	NRL18A := ""
	NRL18B := ""
	if nrl2.NRL18 != "" {
		slice := strings.Split(nrl2.NRL18, ",")
		length := len(slice)
		if length == 1 {
			NRL18A = slice[0]
		} else if length == 2 {
			NRL18A = slice[0]
			NRL18B = slice[1]
		}
	}
	// 拆分护理单录入时间
	tempTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nrl2.NRL38, time.Local)
	NRL38A := tempTime.Format("2006-01-02")
	NRL38B := tempTime.Format("15:04")

	tempTime, _ = time.ParseInLocation("2006-01-02 15:04:05", nrl2.NRL39A, time.Local)
	NRL39B := tempTime.Format("2006-01-02")
	NRL39C := tempTime.Format("15:04")

	//fmt.Printf("user info %+v:", userinfo)

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRL":       nrl2,
		"Menuindex": "7-2",
		"NRL06A":    NRL06A, // 过敏源的index
		"NRL06B":    NRL06B, // 过敏源的补充内容
		"NRL18A":    NRL18A, // n次/天
		"NRL18B":    NRL18B, // 1次/m天
		"NRL38A":    NRL38A, // 录入护理单的年月日
		"NRL38B":    NRL38B, // 录入护理单的时分
		"NRL39B":    NRL39B,
		"NRL39C":    NRL39C,
	}

	c.LoadViewSafely(w, r, "pcnrl/v_nrl2.html", "pc/header_side.html", "pc/header_top.html")
}
