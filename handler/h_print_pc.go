package handler

import (
	"fit"
	"nursing/model"
	"errors"
	"fmt"
	"time"
	//"strings"
	//"encoding/json"
	"encoding/json"
	"strings"
	"strconv"
)

// 腕带打印
type PCWristStrapController struct {
	PCController
}

// 腕带打印
type WristStrap struct {
	Beds       string //床位号
	Name       string // 姓名
	Sex        string //性别
	Age        string //年龄
	ClassName  string //科室
	Pid        string //病人ID
	HospitalId string //住院ID
	Diagnose   string //住院ID
}

// 腕带打印
func (c PCWristStrapController) Get(w *fit.Response, r *fit.Request, p fit.Params) {

	r.ParseForm()
	pid := r.FormValue("pid")
	patients, err := model.GetPatientInfo(pid)
	if len(patients) == 0 || err != nil {
		fit.Logger().LogError("yh", err, errors.New("patient slice is empty"))
		return
	}
	var sex = "未知"
	if patients[0].ABW01 == "1" || patients[0].ABW01 == "M" {
		sex = "男"
	} else if patients[0].ABW01 == "2" || patients[0].ABW01 == "F" {
		sex = "女"
	} else {
		sex = "未知"
	}
	c.Data = fit.Data{"Datas": WristStrap{patients[0].BCQ04, patients[0].VAA05, sex, fmt.Sprint(patients[0].VAA10), patients[0].BCK03C, pid, patients[0].VAA04, patients[0].VAO2.VAO15}}
	c.LoadView(w, "pc/v_wdprint.html")
}

//  医嘱各种打印
type PCBottleStrapController struct {
	PCController
}

// 医嘱各种打印
func (c PCBottleStrapController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	//defer c.LoadView(w, "pc/v_dzprint.html") //输液单,口服单,注射单
	typeint := r.FormIntValue("type")
	switch typeint {
	case 1:
		defer c.LoadView(w, "pc/v_infusionprint.html") //输液单
		break
	case 2:
		defer c.LoadView(w, "pc/v_oralprint.html") //口服单
		break
	case 3:
		defer c.LoadView(w, "pc/v_injectionprint.html") //注射单
		break
	case 4:
		defer c.LoadView(w, "pc/v_ptprint.html") //屏贴
		break
	case 5:
		defer c.LoadView(w, "pc/v_bqprint.html") //标签
		break
	}
	printInfos := r.FormValue("reqdata")
	fmt.Println(printInfos)
	timenow := time.Now().Format("2006-01-02")
	c.Data = fit.Data{
		"PrintInfo": printInfos, // 打印信息
		"Now":       timenow,    // 打印时间
	}
}

//打印交接班
type PCSuccessionController struct {
	PCController
}

func (c PCSuccessionController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadViewSafely(w, r, "pc/v_successionprint.html")
		starttime := r.FormValue("datatime")
		fit.Logger().LogError("gk", starttime)
		if starttime == "" {
			t := time.Now()
			starttime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
		}
		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, errrp := model.QueryDepartmentBeds(classid, false)
		if errrp != nil {
			fit.Logger().LogError("gk", errrp)
			return
		}
		Data["Patients"] = response
		//fit.Logger().LogError("yh 打印交接班", response ,classid)

		successions, err1 := model.OutSuccession("datatime = ? and classid = ?", starttime, classid)
		//fit.Logger().LogError("yh 打印交接班", successions ,classid)
		successiondetails := make([]model.SuccessionDetails, 0)
		if err1 != nil || len(successions) == 0 {
			fit.Logger().LogError("gk", starttime, len(successions), err1)
		} else {
			Data["d_Disabled"] = false
			Data["n_Disabled"] = false
			Data["N_Disabled"] = false
			Data["DateTime"] = starttime
			for _, k := range successions {
				if k.Type == 1 {
					Data["d_NoldPatient"] = k.NoldPatient
					Data["d_NnowPatient"] = k.NnowPatient
					Data["d_NintHospital"] = k.NintHospital
					Data["d_NoutHospital"] = k.NoutHospital
					Data["d_Ninto"] = k.Ninto
					Data["d_Nout"] = k.Nout
					Data["d_Nsurgery"] = k.Nsurgery
					Data["d_Nchildbirth"] = k.Nchildbirth
					Data["d_Ncritically"] = k.Ncritically
					Data["d_Ndeath"] = k.Ndeath
					Data["d_NintensiveCare"] = k.NintensiveCare
					Data["d_NprimaryCare"] = k.NprimaryCare
					Data["d_NursingName"] = k.NursingName

					if userinfo.UID == k.NursingID {
						Data["d_Disabled"] = false
					} else {
						Data["d_Disabled"] = true
					}
				} else if k.Type == 2 {
					Data["n_NoldPatient"] = k.NoldPatient
					Data["n_NnowPatient"] = k.NnowPatient
					Data["n_NintHospital"] = k.NintHospital
					Data["n_NoutHospital"] = k.NoutHospital
					Data["n_Ninto"] = k.Ninto
					Data["n_Nout"] = k.Nout
					Data["n_Nsurgery"] = k.Nsurgery
					Data["n_Nchildbirth"] = k.Nchildbirth
					Data["n_Ncritically"] = k.Ncritically
					Data["n_Ndeath"] = k.Ndeath
					Data["n_NintensiveCare"] = k.NintensiveCare
					Data["n_NprimaryCare"] = k.NprimaryCare
					Data["n_NursingName"] = k.NursingName

					if userinfo.UID == k.NursingID {
						Data["n_Disabled"] = false
					} else {
						Data["n_Disabled"] = true
					}
				} else {
					Data["N_NoldPatient"] = k.NoldPatient
					Data["N_NnowPatient"] = k.NnowPatient
					Data["N_NintHospital"] = k.NintHospital
					Data["N_NoutHospital"] = k.NoutHospital
					Data["N_Ninto"] = k.Ninto
					Data["N_Nout"] = k.Nout
					Data["N_Nsurgery"] = k.Nsurgery
					Data["N_Nchildbirth"] = k.Nchildbirth
					Data["N_Ncritically"] = k.Ncritically
					Data["N_Ndeath"] = k.Ndeath
					Data["N_NintensiveCare"] = k.NintensiveCare
					Data["N_NprimaryCare"] = k.NprimaryCare

					Data["N_NursingName"] = k.NursingName

					if userinfo.UID == k.NursingID {
						Data["N_Disabled"] = false
					} else {
						Data["N_Disabled"] = true
					}
				}
			}

			successiondetails, _ = model.OutSuccessionDetails("datatime = ? and classid = ?", starttime, classid)
		}
		//总条数
		var page int
		pageDataNum := 10 //每页多少数据
		if len(successiondetails)%pageDataNum == 0 {
			page = len(successiondetails) / pageDataNum
		} else {
			page = len(successiondetails)/pageDataNum + 1
		}

		//计算最后一页差多少数据
		//[]NRL3
		//fit.Logger().LogError("nrl page info :", len(mods), mods)
		oldlen := len(successiondetails)
		for i := 0; i < (page*pageDataNum - oldlen); i++ {
			successiondetails = append(successiondetails, model.SuccessionDetails{})
		}

		//fit.Logger().LogError("yh 打印交接班", page,len(successiondetails))
		Data["Userinfo"] = userinfo
		Data["Menuindex"] = "6-0"
		Data["Len"] = len(successiondetails)
		Data["Successiondetails"] = successiondetails
		Data["Page"] = page
		Data["Pages"] = make([]int, page)
		Data["PageDataNum"] = pageDataNum
		c.Data = Data
	}
}

//打印深圳万丰医院护理记录单
type PCNrl1Controller struct {
	PCNRLController
}

func (c PCNrl1Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl1print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "1")
	if !has {
		return
	}
	r.ParseForm()
	pid = r.FormValue("pid")
	// 时间
	var datestr1, datestr2 string
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000+60*60*24, 0).Format("2006-01-02 15:04:05")
	}

	mods, errdata := model.GetNRL1Data(pid, datestr1, datestr2)
	if errdata != nil {
		fit.Logger().LogError("PCQueryNRL1Title error :", errdata)
		fmt.Fprintln(w, "查询NRL 错误！  NRL error", errdata)
		return
	}

	//总条数
	var page int
	pageDataNum := 8 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	//计算最后一页差多少数据
	//[]NRL3
	//fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.NRLModel{})
	}

	nrl1Title := model.NRL1Title{PatientId: pInfo.VAA01}
	errTitle := nrl1Title.PCQueryNRL1Title()
	if errTitle != nil {
		fit.Logger().LogError("PCQueryNRL1Title error :", errTitle)
		fmt.Fprintln(w, "查询NRL 错误！  NRL error", errTitle)
		return
	}

	fmt.Printf("mods %+v\n length：%d\n\n", mods, len(mods))

	c.Data = fit.Data{
		"Userinfo":    userinfo, // 护士信息
		"PInfo":       pInfo,    // 病人信息
		"Beds":        beds,     // 床位list
		"NRLList":     mods,
		"NRLTitle":    nrl1Title,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-1",
	}

}

//打印深圳万丰医院首次护理记录单
type PCNrl2Controller struct {
	PCNRLController
}

func (c PCNrl2Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl2print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "2")
	if !has {
		return
	}
	var err error
	r.ParseForm()
	pid1 := r.FormValue("pid")
	if pid1 != "" {
		pid = pid1
	}

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

}

//打印深圳万丰医院基本生活活动能力(BADL)
type PCNrl3Controller struct {
	PCNRLController
}

func (c PCNrl3Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl3print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "3")
	if !has {
		return
	}
	pid = r.FormValue("pid")
	datestr1, datestr2, _, _, err := c.GetPageInfo(w, r, "3", pid)
	if err != nil {
		fit.Logger().LogError("nrl page info :", err)
		return
	}
	// 护理单
	mods, err13 := model.PCQueryNRL3(pid, datestr1, datestr2, -1)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error")
		return
	}

	var page int
	pageDataNum := 7 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	//datas := make([][]model.NRL3, 0)
	//for i := 0; i < page; i++ {
	//	x := i*7 + 7
	//	if x > len(mods) {
	//		x = len(mods)
	//	}
	//	datas = append(datas, mods[i*7:x])
	//}
	//计算最后一页差多少数据
	//[]NRL3
	fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.NRL3{})
	}
	c.Data = fit.Data{
		"Userinfo":    userinfo, // 护士信息
		"PInfo":       pInfo,    // 病人信息
		"Beds":        beds,     // 床位list
		"NRLList":     mods,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-3",
	}
	fit.Logger().LogError("nrl page info :", len(mods), mods)
}

//打印深圳万丰医院深静脉血栓(DVT)形成风险评估表
type PCNrl4Controller struct {
	PCNRLController
}

func (c PCNrl4Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl4print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "4")
	if !has {
		return
	}
	pid = r.FormValue("pid")
	// 起止时间  页码
	datestr1, datestr2, _, _, err := c.GetPageInfo(w, r, "4", pid)
	if err != nil {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL4(pid, datestr1, datestr2, -1)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err13)
		fit.Logger().LogError("nrl page info :", err13)
		return
	}
	var page int
	pageDataNum := 6 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.NRL4{})
	}
	c.Data = fit.Data{
		"Userinfo":    userinfo, // 护士信息
		"PInfo":       pInfo,    // 病人信息
		"Beds":        beds,     // 床位list
		"NRLList":     mods,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-4",
	}
	fit.Logger().LogError("nrl page info :", len(mods), mods)
}

//打印深静脉血栓护理观察表
type PCNrl5Controller struct {
	PCNRLController
}

func (c PCNrl5Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl5print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "5")
	if !has {
		return
	}
	fmt.Printf("bbbbb", "dgrgghr2")
	pid = r.FormValue("pid")

	// 起止时间  页码
	// 时间
	var datestr1, datestr2 string
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000+60*60*24-1, 0).Format("2006-01-02 15:04:05")
	}

	// 护理单
	mods, err13 := model.PCQueryNRL5(pid, datestr1, datestr2, -1)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err13)
		return
	}

	var page int
	pageDataNum := 3 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	//datas := make([][]model.NRL3, 0)
	//for i := 0; i < page; i++ {
	//	x := i*7 + 7
	//	if x > len(mods) {
	//		x = len(mods)
	//	}
	//	datas = append(datas, mods[i*7:x])
	//}
	//计算最后一页差多少数据
	//[]NRL3
	fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.APNModel{})
	}
	c.Data = fit.Data{
		"Userinfo":    userinfo, // 护士信息
		"PInfo":       pInfo,    // 病人信息
		"Beds":        beds,     // 床位list
		"NRLList":     mods,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-5",
	}
}

//打印深圳万丰医院压疮风险因素评估表（Braden评分）
type PCNrl6Controller struct {
	PCNRLController
}

func (c PCNrl6Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl6print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "6")
	if !has {
		return
	}
	pid = r.FormValue("pid")
	// 起止时间  页码
	datestr1, datestr2, _, _, err := c.GetPageInfo(w, r, "6", pid)
	if err != nil {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL6(pid, datestr1, datestr2, -1)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}
	var page int
	pageDataNum := 4 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.NRL6{})
	}
	c.Data = fit.Data{
		"Userinfo":    userinfo, // 护士信息
		"PInfo":       pInfo,    // 病人信息
		"Beds":        beds,     // 床位list
		"NRLList":     mods,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-6",
	}

}

//打印深圳万丰医院患者跌到风险评估护理单
type PCNrl7Controller struct {
	PCNRLController
}

func (c PCNrl7Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl7print.html")
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "7")
	if !has {
		return
	}
	pid = r.FormValue("pid")
	// 起止时间  页码
	datestr1, datestr2, _, _, err := c.GetPageInfo(w, r, "7", pid)
	if err != nil {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL7(pid, datestr1, datestr2, -1)

	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	nrl7Title := model.NRL7Title{
		PatientId: pInfo.VAA01,
	}

	errTitle := nrl7Title.PCQueryNRL7Title()
	if errTitle != nil {
		fit.Logger().LogError("PCQueryNRL1Title error :", errTitle)
	}
	var page int
	pageDataNum := 7 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.NRL7{})
	}
	c.Data = fit.Data{
		"Userinfo": userinfo, // 护士信息
		"PInfo":    pInfo,    // 病人信息
		"Beds":     beds,     // 床位list
		"NRLTitle": nrl7Title,
		//"NRL08":     pmodel.NRL08,
		//"NRL08A":    pmodel.NRL08A,
		//"NRL08B":    pmodel.NRL08B,
		"NRLList":     mods,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-7",
	}

	fit.Logger().LogError("PCNRL7Controller  :", mods, nrl7Title)
}

//打印深圳万丰医院疼痛强度评分量表
type PCNrl8Controller struct {
	PCNRLController
}

func (c PCNrl8Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.LoadViewSafely(w, r, "pcnrl/v_nrl8print.html")

	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "8")
	if !has {
		return
	}
	pid = r.FormValue("pid")
	// 起止时间  页码
	datestr1, datestr2, _, _, err := c.GetPageInfo(w, r, "8", pid)
	if err != nil {
		return
	}

	// 护理单
	mods, err13 := model.PCQueryNRL8(pid, datestr1, datestr2, -1)
	if err13 != nil {
		fmt.Fprintln(w, "参数错误！ pc get nrl8 model error", err13)
		return
	}
	var page int
	pageDataNum := 16 //每页多少数据
	if len(mods)%pageDataNum == 0 {
		page = len(mods) / pageDataNum
	} else {
		page = len(mods)/pageDataNum + 1
	}
	fit.Logger().LogError("nrl page info :", len(mods), mods)
	oldlen := len(mods)
	for i := 0; i < (page*pageDataNum - oldlen); i++ {
		mods = append(mods, model.NRL8{})
	}
	c.Data = fit.Data{
		"Userinfo":    userinfo, // 护士信息
		"PInfo":       pInfo,    // 病人信息
		"Beds":        beds,     // 床位list
		"NRLList":     mods,
		"PageNum":     page,
		"PageDataNum": pageDataNum,
		"Menuindex":   "7-8",
	}
}
