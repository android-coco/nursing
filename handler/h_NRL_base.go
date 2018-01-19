package handler

import (
	"fit"
	"nursing/model"
	"fmt"
	"strconv"
	"time"
	"nursing/utils"
)

type NRLController struct {
	fit.Controller
}

// 模板 template PDA端
func (c *NRLController) LoadPInfoWithPid(w *fit.Response, r *fit.Request, pid int64) (model.PatientInfo, bool) {
	pinfo, err := model.GetPatientInfo(strconv.FormatInt(pid, 10))
	if err != nil {
		fit.Logger().LogError("GetPatientInfo :", err)
		fmt.Fprintln(w, "无法查询到相关病人的信息，err：", err.Error())
		return model.PatientInfo{}, false
	}
	if len(pinfo) == 0 {
		fit.Logger().LogError("PatientInfo is empty", pid)
		fmt.Fprintln(w, "无法查询到相关病人的信息.")
		return model.PatientInfo{}, false
	}
	return pinfo[0], true
}

func (c *NRLController) LoadPInfoAndAccountWithPidUid(w *fit.Response, r *fit.Request, pid, uid string) (model.PatientInfo, model.Account, bool) {
	pinfo, err := model.GetPatientInfo(pid)
	if err != nil {
		fit.Logger().LogError("GetPatientInfo :", err)
		fmt.Fprintln(w, "无法查询到相关病人的信息，err：", err.Error())
		return model.PatientInfo{}, model.Account{}, false
	}
	if len(pinfo) == 0 {
		fit.Logger().LogError("PatientInfo is empty")
		fmt.Fprintln(w, "无法查询到相关病人的信息.")
		return model.PatientInfo{}, model.Account{}, false
	}

	account, err2 := model.FetchAccountWithUid(uid)
	if err2 != nil {
		fit.Logger().LogError("NRL PDA err：", err2)
		fmt.Fprintln(w, "用户信息获取失败！", err2)
		return model.PatientInfo{}, model.Account{}, false
	}
	return pinfo[0], account, true
}

/*
func (c *NRLController) LoadNRLDataWithParm(w *fit.Response, r *fit.Request, nrlType string) (nrl interface{}, ty, rid, pid, uid string, isOk bool) {
	pid = r.FormValue("pid") //病人id
	uid = r.FormValue("uid") //护士id
	rid = r.FormValue("rid") // 护理记录单id
	ty = r.FormValue("type") // 1=add， 2=edit

	if ty == "1" {
		if "" == pid || "" == uid {
			fmt.Fprintln(w, "参数错误！")
			return
		}
		switch nrlType {
		case "3":
			nrl = model.NRL3{}
		case "4":
			nrl = model.NRL4{}
		case "5":
			nrl = model.NRL5{}
		case "6":
			nrl = model.NRL6{}
		case "7":
			nrl = model.NRL7{}
		case "8":
			nrl = model.NRL8{}
		default:
			fmt.Fprintln(w, "参数错误！")
			return
		}
	} else if ty == "2" {
		if rid == "" {
			fmt.Fprintln(w, "参数错误！")
			return
		}
		var err1 error

		nrl, pid, uid, err1 = model.QueryNRLWithRid(nrlType, rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
			fmt.Fprintln(w, "NRL 查询错误！", err1.Error())
			return
		}
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	return nrl, ty, rid, pid, uid, true
}
*/

//  PDA  加载模板
// 查看
func (c NRLController) NRLCheck(w *fit.Response, r *fit.Request, p fit.Params, nrlMod interface{}) {
	// 文书id
	rid := r.FormInt64Value("rid")
	if 0 == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}
	c.Data = make(fit.Data)
	tplname := ""
	switch obj := nrlMod.(type) {
	case *model.NRL3:
		tplname = "v_nrl3.html"
		obj.ID = rid
	case *model.NRL4:
		tplname = "v_nrl4.html"
		obj.ID = rid
	case *model.NRL5:
		tplname = "v_nrl5.html"
		obj.ID = rid
	case *model.NRL6:
		tplname = "v_nrl6.html"
		obj.ID = rid
	case *model.NRL7:
		tplname = "v_nrl7.html"
		obj.ID = rid
	case *model.NRL8:
		tplname = "v_nrl8.html"
		obj.ID = rid

	default:
		fmt.Fprintln(w, "NRLQueryNRLModel : invalid nrlMod")
		return
	}

	pid, _, err1 := model.NRLQueryNRLModel(nrlMod)
	if err1 != nil {
		fit.Logger().LogError("m_NRL", err1)
		fmt.Fprintln(w, " 无法查询到相关病人的信息. error :", err1.Error())
		return
	}
	//fmt.Println("pid:", pid)

	if obj, ok := nrlMod.(*model.NRL7); ok {
		nrl7Title := model.NRL7Title{PatientId: obj.PatientId}
		errt := nrl7Title.PCQueryNRL7Title()
		if errt != nil {
			fit.Logger().LogError("m_NR7", errt)
			fmt.Fprintln(w, " 无法查询到相关病人的信息. error :", errt.Error())
			return
		}

		c.Data["NRLTitle"] = nrl7Title
	}

	// 查询对应病人信息
	patient, has := c.LoadPInfoWithPid(w, r, pid)
	if !has {
		return
	}
	c.Data["PInfo"] = patient
	c.Data["NRL"] = nrlMod

	tplname = "pda/" + tplname
	c.LoadView(w, tplname)
}

// 1=添加、2=编辑
func (c NRLController) NRLEdit(w *fit.Response, r *fit.Request, p fit.Params, nrlMod interface{}) {

	pid := r.FormValue("pid")      //病人id
	uid := r.FormValue("uid")      //护士id
	rid := r.FormInt64Value("rid") // 护理记录单id
	ty := r.FormValue("type")      // 1=add， 2=edit

	c.Data = make(fit.Data)
	tplname := ""

	if ty == "1" {
		if "" == pid || "" == uid {
			fmt.Fprintln(w, "参数错误！")
			return
		}
		switch nrlMod.(type) {
		case *model.NRL3:
			tplname = "v_nrl3_edit.html"
		case *model.NRL4:
			tplname = "v_nrl4_edit.html"
		case *model.NRL5:
			tplname = "v_nrl5_edit.html"
		case *model.NRL6:
			tplname = "v_nrl6_edit.html"
		case *model.NRL7:
			tplname = "v_nrl7_edit.html"
		case *model.NRL8:
			tplname = "v_nrl8_edit.html"
		default:
			fmt.Fprintln(w, "NRLQueryNRLModel : invalid nrlMod")
			return
		}
	} else if ty == "2" {
		if rid == 0 {
			fmt.Fprintln(w, "参数错误！")
			return
		}
		switch obj := nrlMod.(type) {
		case *model.NRL3:
			tplname = "v_nrl3_edit.html"
			obj.ID = rid

		case *model.NRL4:
			tplname = "v_nrl4_edit.html"
			obj.ID = rid
		case *model.NRL5:
			tplname = "v_nrl5_edit.html"
			obj.ID = rid
		case *model.NRL6:
			tplname = "v_nrl6_edit.html"
			obj.ID = rid
		case *model.NRL7:
			tplname = "v_nrl7_edit.html"
			obj.ID = rid
		case *model.NRL8:
			tplname = "v_nrl8_edit.html"
			obj.ID = rid

		default:
			fmt.Fprintln(w, "NRLQueryNRLModel : invalid nrlMod")
			return
		}

		pidnum, uidstr, err1 := model.NRLQueryNRLModel(nrlMod)
		if err1 != nil {
			fit.Logger().LogError("m_NRL", err1)
			fmt.Fprintln(w, " 无法查询到相关病人的信息. error :", err1.Error())
			return
		}
		pid = utils.FormatInt64(pidnum)
		uid = uidstr
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	pinfo, err := model.GetPatientInfo(pid)
	if err != nil {
		fit.Logger().LogError("GetPatientInfo :", err)
		fmt.Fprintln(w, "无法查询到相关病人的信息，err：", err.Error())
		return
	}
	if len(pinfo) == 0 {
		fit.Logger().LogError("PatientInfo is empty")
		fmt.Fprintln(w, "无法查询到相关病人的信息.PatientInfo is empty")
		return
	}
	patient := pinfo[0]
	fmt.Println("-----------", pid, "uid:", uid)
	account, err2 := model.FetchAccountWithUid(uid)
	if err2 != nil {
		fit.Logger().LogError("NRL PDA err：", err2)
		fmt.Fprintln(w, "用户信息获取失败！", err2.Error())
		return
	}

	if obj, ok := nrlMod.(*model.NRL7); ok {
		nrl7Title := model.NRL7Title{PatientId: obj.PatientId}
		errt := nrl7Title.PCQueryNRL7Title()
		if errt != nil {
			fit.Logger().LogError("m_NR7", errt)
			fmt.Fprintln(w, " 无法查询到相关病人的信息. error :", errt.Error())
			return
		}
		c.Data["NRLTitle"] = nrl7Title
	}

	c.Data["PInfo"] = patient
	c.Data["NRL"] = nrlMod
	c.Data["Type"] = ty
	c.Data["Rid"] = rid
	c.Data["Account"] = account

	tplname = "pda/" + tplname
	c.LoadView(w, tplname)
}

// 接口
// 添加护理记录单
func (c NRLController) NRLAddRecord(w *fit.Response, r *fit.Request, p fit.Params, nrlMod interface{}) {
	defer c.ResponseToJson(w)
	// 获取fit.Request中的参数赋值到mod中
	errflag := c.FitSetStruct(nrlMod, r)
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	var (
		patientId          int64
		nurseId, nurseName string
	)
	var nrlType int64
	switch obj := nrlMod.(type) {
	case *model.NRL3:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
		nrlType = 3
	case *model.NRL4:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
		nrlType = 4
	case *model.NRL5:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
		nrlType = 5
	case *model.NRL6:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
		nrlType = 6
	case *model.NRL7:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
		nrlType = 7
	case *model.NRL8:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
		nrlType = 8
	default:
		fit.Logger().LogError("NRLAddRecord : invalid nrlMod")
		c.RenderingJsonAutomatically(10001, "参数错误！")
		return
	}
	if errflag != nil || err4 != nil || patientId == 0 || nurseId == "" || nurseName == "" {
		fit.Logger().LogError("NRL3 :", err4, errflag)
		c.RenderingJsonAutomatically(10002, "参数不完整")
		return
	}

	rid, erradd := model.NRLInsertData(nrlMod)
	//fmt.Println("rid:", rid, erradd, "pid:", patientId)

	if erradd != nil {
		fit.Logger().LogError("DB NRL add :", erradd.Error())
		c.RenderingJson(2, "DB错误！", erradd.Error())
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    nrlType,
			NursingId:   nurseId,
			NursingName: nurseName,
			ClassId:     r.FormValue("did"),
			PatientId:   r.FormValue("pid"),
			RecordId:    rid,
			Comment:     "新增",
		}
		_, errRecord := model.InsertNRecords(nurseRecord)
		checkerr("nurse record err:", errRecord)

		c.RenderingJsonAutomatically(0, "添加成功！")
	}

}

// 修改护理记录单
func (c NRLController) NRLUpdateRecord(w *fit.Response, r *fit.Request, p fit.Params, nrlMod interface{}) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormInt64Value("rid")
	if rid == 0 {
		fit.Logger().LogError("Error", "nrl update : rid in empty")
		c.RenderingJsonAutomatically(10001, "参数错误！")
		return
	}

	errflag := c.FitSetStruct(nrlMod, r)
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	var (
		patientId          int64
		nurseId, nurseName string
	)
	switch obj := nrlMod.(type) {
	case *model.NRL3:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
	case *model.NRL4:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
	case *model.NRL5:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
	case *model.NRL6:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
	case *model.NRL7:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
	case *model.NRL8:
		obj.DateTime = model.FitTime(datetime)
		patientId = obj.PatientId
		nurseId = obj.NurseId
		nurseName = obj.NurseName
	default:
		fit.Logger().LogError("hy:", "NRLAddRecord : invalid nrlMod")
		c.RenderingJsonAutomatically(10001, "参数错误！")
		return
	}
	if errflag != nil || err4 != nil || patientId == 0 || nurseId == "" || nurseName == "" {
		fit.Logger().LogError("NRL :", err4, errflag)
		c.RenderingJsonAutomatically(10002, "参数不完整")
		return
	}

	_, errUpdate := model.NRLUpdateData(rid, nrlMod)
	if errUpdate != nil {
		fit.Logger().LogError("DB nrl update :", errUpdate)
		c.RenderingJson(2, "DB错误！", errUpdate.Error())
	} else {
		_, errRecord := model.UpadteNRecords(rid, r.FormValue("datetime"))
		checkerr("nurse record update err:", errRecord)

		c.RenderingJsonAutomatically(0, "修改成功！")
	}
}

// 删除护理
func (c NRLController) NRLDeleteRecord(w *fit.Response, r *fit.Request, p fit.Params, nrlMod interface{}) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormInt64Value("rid")
	if rid == 0 {
		fit.Logger().LogError("Error", "nrl delete : rid in empty")
		c.RenderingJsonAutomatically(10001, "参数错误！")
		return
	}
	_, errdel := model.NRLDeleteData(rid, nrlMod)
	if errdel != nil {
		fit.Logger().LogError("nrl delete :", errdel)
		c.RenderingJsonAutomatically(3, "删除失败！")
	} else {
		c.RenderingJsonAutomatically(0, "删除成功！")
	}
}

type PCNRLController struct {
	PCController
}

/*
护理单 PC用
*/
// 护士信息 床位表   病人id  病人信息
func (c *PCNRLController) GetBedsAndUserinfo(w *fit.Response, r *fit.Request, nrlType string) (userinfo model.UserInfoDup, beds []model.PCBedDup, pid string, pInfo model.PCBedDup, isHas bool) {
	// 护士信息
	isHas = false
	var err error
	userinfo, err = c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		fit.Logger().LogError("Error", "参数错误！  user info error", err)
		return
	}

	beds, err = model.QueryDepartmentBeds(userinfo.DepartmentID, false)
	if err != nil {
		fmt.Fprintln(w, "query beds err:", err)
		fit.Logger().LogError("query beds err:", err)
		return
	}

	pid = r.FormValue("pid")
	if pid == "" {
		if len(beds) == 0 {
			fit.Logger().LogError("beds is empty")
			fmt.Fprintln(w, "beds is empty")
			return
		}
		pidnum := beds[0].VAA01
		pid = strconv.FormatInt(pidnum, 10)
		url := "/pc/record/nrl" + nrlType + "?pid=" + pid
		if nrlType == "9" {
			url = "/pc/templist" + "?pid=" + pid
		} else if nrlType == "91" {
			url = "/pc/templist/print?pid=" + pid
		}

		c.Redirect(w, r, url, 302)
		return userinfo, beds, pid, pInfo, false
	}

	// 病人信息
	for _, val := range beds {
		if strconv.FormatInt(val.VAA01, 10) == pid {
			pInfo = val
			break
		}
	}

	if pInfo.VAA01 == 0 {
		fit.Logger().LogError("pc nrl pInfo is empty")
		fmt.Fprintln(w, "pc nrl pInfo is empty")
		return userinfo, beds, pid, pInfo, false
	}

	return userinfo, beds, pid, pInfo, true
}

// pc 文书 翻页处理时间的页码的
func (c *PCNRLController) GetPageInfo(w *fit.Response, r *fit.Request, nrlType, pid string) (datestr1, datestr2 string, pageindex, pagenum int, err error) {
	// 时间
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000+60*60*24-1, 0).Format("2006-01-02 15:04:05")
	}

	// 总条数
	count, errCount := model.PCQueryNRLPageCount(nrlType, pid, datestr1, datestr2)
	if errCount != nil {
		fmt.Fprintln(w, "nrl list err :", errCount)
		fit.Logger().LogError("nrl page info :", errCount)
		err = errCount
		return
	}

	var peerPage int64 = 9
	switch nrlType {
	case "1":
		peerPage = 9
	case "3":
		peerPage = 9
	case "4":
		peerPage = 9
	case "5":
		peerPage = 5
	case "6":
		peerPage = 4
	case "7":
		peerPage = 8
	case "8":
		peerPage = 9
	default:
		peerPage = 9
	}

	//总页数
	pagenum = int((count-1)/peerPage) + 1
	//当前页数
	index := r.FormValue("num")
	pageindex, errnum := strconv.Atoi(index)
	if errnum != nil {
		pageindex = int(pagenum)
	}
	if pageindex < 1 {
		pageindex = 1
	} else if pageindex > pagenum {
		pageindex = pagenum
	}
	fmt.Println("count:", count, "pageNum:", pagenum, "pageindex:", pageindex)

	return
}

type recordData struct {
	userInfo  model.UserInfoDup
	pInfo     model.PCBedDup
	beds      []model.PCBedDup
	pageNum   int
	pageIndex int
	pid       string
	datestr1  string
	datestr2  string
}

func (c *PCNRLController) BaseNRLRecord(w *fit.Response, r *fit.Request, nrlType string) (data recordData, isHas bool) {
	//isHas = false
	c.Data = make(fit.Data)
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		fit.Logger().LogError("hy:", "参数错误！  user info error", err)
		return
	}
	data.userInfo = userinfo

	beds, err := model.QueryDepartmentBeds(userinfo.DepartmentID, false)
	if err != nil {
		fmt.Fprintln(w, "query beds err:", err)
		fit.Logger().LogError("hy:", "query beds err:", err)
		return
	}
	if len(beds) == 0 {
		fit.Logger().LogError("hy:", "beds is empty")
		fmt.Fprintln(w, "beds is empty")
		return
	}
	data.beds = beds

	pid := r.FormValue("pid")
	if pid == "" {
		pidnum := beds[0].VAA01
		pid = strconv.FormatInt(pidnum, 10)
		url := "/pc/record/nrl" + nrlType + "?pid=" + pid
		if nrlType == "9" {
			url = "/pc/templist" + "?pid=" + pid
		} else if nrlType == "91" {
			url = "/pc/templist/print?pid=" + pid
		}
		c.Redirect(w, r, url, 302)
		return
	}

	var pInfo model.PCBedDup
	// 病人信息
	for _, val := range beds {
		if strconv.FormatInt(val.VAA01, 10) == pid {
			pInfo = val
			break
		}
	}
	if pInfo.VAA01 == 0 {
		fit.Logger().LogError("hy:", "pc nrl pInfo is empty")
		fmt.Fprintln(w, "pc nrl pInfo is empty")
		return
	}
	data.pInfo = pInfo
	data.pid = pid

	var datestr1, datestr2 string
	// 时间
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		//datestr1 = pInfo.VAE11.ParseToSecond()
		//datestr2 = time.Now().Format("2006-01-02 15:04:05")
		//
		//date1 := time.Time(pInfo.VAE11).Unix() * 1000
		//date2 := time.Now().Unix() * 1000
		//paramstr := fmt.Sprintf("&sdate=%d&edate=%d", date1, date2)
		//urlstr := r.URL.String() + paramstr
		//c.Redirect(w, r, urlstr, 302)
		//return
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000+60*60*24-1, 0).Format("2006-01-02 15:04:05")
	}

	data.datestr1 = datestr1
	data.datestr2 = datestr2
	// 总条数
	count, errCount := model.PCQueryNRLPageCount(nrlType, pid, datestr1, datestr2)
	if errCount != nil {
		fmt.Fprintln(w, "nrl list err :", errCount)
		fit.Logger().LogError("hy:", "nrl page info :", errCount)
		return
	}

	var peerPage int64 = 9
	switch nrlType {
	case "1":
		peerPage = 9
	case "3":
		peerPage = 9
	case "4":
		peerPage = 9
	case "5":
		peerPage = 5
	case "6":
		peerPage = 4
	case "7":
		peerPage = 8
	case "8":
		peerPage = 9
	default:
		peerPage = 9
	}

	//总页数
	pagenum := int((count-1)/peerPage) + 1
	//当前页数
	pageindex := r.FormIntValue("num")
	if pageindex < 1 || pageindex > pagenum {
		pageindex = pagenum
	}
	//fmt.Println("count:", count, "pageNum:", pagenum, "pageindex:", pageindex)
	data.pageIndex = pageindex
	data.pageNum = pagenum

	return data, true
}

func (c *PCNRLController) LoadPCNRLView(w *fit.Response, r *fit.Request, tplname string)  {
	c.LoadViewSafely(w, r, tplname, "pc/header_side.html", "pc/header_top.html")
}