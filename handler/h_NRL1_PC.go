package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"time"
)

// 护理记录单 PC端
type PCNRL1Controller struct {
	PCController
}

// 护理记录单
// 查看  list

func (c PCNRL1Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {

	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "1")
	if !has {
		return
	}

	// 时间
	var datestr1, datestr2 string
	var pageindex, pagenum int
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000, 0).Format("2006-01-02 15:04:05")
	}
	//fmt.Println("--------------:", datestr1, datestr2)
	mods, errdata := model.GetNRL1Data(pid, datestr1, datestr2)
	if errdata != nil {
		fit.Logger().LogError("PCQueryNRL1Title error :", errdata)
		fmt.Fprintln(w, "查询NRL 错误！  NRL error", errdata)
		return
	}

	//总条数
	count := len(mods)
	peerPage := 9
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

	list := make([]model.NRLModel, peerPage)
	if pageindex == pagenum {
		list = mods[(pageindex-1)*peerPage:count]
	} else {
		list = mods[(pageindex-1)*peerPage:pageindex*peerPage]
	}

	nrl1Title := model.NRL1Title{VAA01: pInfo.VAA01}
	errTitle := nrl1Title.PCQueryNRL1Title()
	if errTitle != nil {
		fit.Logger().LogError("PCQueryNRL1Title error :", errTitle)
		fmt.Fprintln(w, "查询NRL 错误！  NRL error", errTitle)
		return
	}

	//fmt.Printf("mods %+v\n length：%d\n\n", mods, len(mods))
	//fmt.Printf("list %+v\n length：%d\n\n", list, len(list))

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRLList":   list,
		"NRLTitle":  nrl1Title,
		"PageNum":   pagenum,
		"PageIndex": pageindex,
		"Menuindex": "7-1",
	}

	c.LoadViewSafely(w, r, "pc/v_nrl1.html", "pc/header_side.html", "pc/header_top.html")
}

// 护理记录单 出入量统计
func (c PCNRL1Controller) NRLIOStatistcs(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 病人id
	pid := r.FormValue("pid")
	// 起止时间  如果传默认为全部
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")

	pinfo, err := model.GetPatientInfo(pid)
	if len(pinfo) == 0 || err != nil || pid == "" || startDate == "" || endDate == "" {
		fit.Logger().LogError("无法查询到相关病人的信息, 参数不完整 :", err)
		c.RenderingJsonAutomatically(3, "参数错误")
		return
	}
	//patient := pinfo[0]
	//timenow := time.Now().Format("2006-01-02 15:04:05")
	//hospitalDate := patient.VAA73.ParseToSecond()

	//fmt.Printf("pinfo :%+v\n\n", pinfo)
	//fmt.Println("--- date", startDate, endDate)
	mods15, errdata1 := model.PCQueryNRLIntakeOutputData(pid, "15", startDate, endDate)
	mods16, errdata2 := model.PCQueryNRLIntakeOutputData(pid, "16", startDate, endDate)
	//fmt.Println(mods15, mods16)
	if errdata1 != nil || errdata2 != nil {
		c.RenderingJsonAutomatically(2, "数据查询出现错误")
	} else {
		var intakeA, intakeB, intakeC, intakeD, intakeTotal int
		var outputA, outputB, outputC, outputTotal int
		var intakeDArr, outputCArr []model.NRLData
		for _, mod := range mods15 {
			intvalue, errParse := strconv.Atoi(mod.Value)
			if errParse != nil {
				c.RenderingJsonAutomatically(2, "数据出现错误")
			}
			switch mod.SubType {
			case 1:
				intakeA += intvalue
			case 2:
				intakeB += intvalue
			case 3:
				intakeC += intvalue
			case 4:
				intakeD += intvalue
				intakeDArr = append(intakeDArr, mod)
			default:
			}
		}

		intakeTotal = intakeA + intakeB + intakeC + intakeD
		for _, mod := range mods16 {
			intvalue, errParse := strconv.Atoi(mod.Value)
			if errParse != nil {
				c.RenderingJsonAutomatically(2, "数据出现错误")
			}
			switch mod.SubType {
			case 1:
				outputA += intvalue
			case 2:
				outputB += intvalue
			case 3:
				outputC += intvalue
				outputCArr = append(outputCArr, mod)
			default:
			}
		}
		outputTotal = outputA + outputB + outputC

		datas := map[string]interface{}{
			"intake":      mods15,
			"intakeA":     intakeA,
			"intakeB":     intakeB,
			"intakeC":     intakeC,
			"intakeD":     intakeD,
			"intakeDArr":  intakeDArr,
			"intakeTotal": intakeTotal,

			"output":      mods16,
			"outputA":     outputA,
			"outputB":     outputB,
			"outputC":     outputC,
			"outputCArr":  outputCArr,
			"outputTotal": outputTotal,
		}
		c.RenderingJson(0, "成功", datas)
	}
}

// 护理记录单 出入量 插入 同步到体温单
func (c PCNRL1Controller) NRLIOTypeIn(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 病人id
	pid := r.FormValue("pid")
	VAA01, err1 := strconv.ParseInt(pid, 10, 64)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 科室ID
	did := r.FormValue("did")
	BCK01, err5 := strconv.Atoi(did)

	// 起止时间  如果传默认为全部
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")
	DateTime1, errdate1 := time.ParseInLocation("2006-01-02 15:04:05", startDate, time.Local)
	DateTime2, errdate2 := time.ParseInLocation("2006-01-02 15:04:05", endDate, time.Local)

	dataType := r.FormValue("dataType")
	pinfo, err := model.GetPatientInfo(pid)
	if len(pinfo) == 0 || err != nil || pid == "" || startDate == "" || endDate == "" ||
		err1 != nil || err5 != nil || errdate1 != nil || errdate2 != nil {
		fit.Logger().LogError("无法查询到相关病人的信息, 参数不完整 :", err)
		c.RenderingJsonAutomatically(3, "参数错误")
		return
	}
	//fmt.Printf("pinfo :%+v\n\n", pinfo)
	//fmt.Println("--- date", startDate, endDate)
	// 1输液2饮食3饮水4其他
	intakeA := r.FormValue("intakeA")
	intakeB := r.FormValue("intakeB")
	intakeC := r.FormValue("intakeC")
	intakeD := r.FormValue("intakeD")
	intakeDV := r.FormValue("intakeDV")
	intakeTotal := r.FormValue("intakeTotal")

	//1尿量 3其他
	outputA := r.FormValue("outputA")
	outputB := r.FormValue("outputB")
	outputC := r.FormValue("outputC")
	outputCV := r.FormValue("outputCV")
	outputTotal := r.FormValue("outputTotal")

	mod := model.IOStatistics{
		VAA01:       VAA01,
		BCK01:       BCK01,
		BCE01A:      BCE01A,
		BCE03A:      BCE03A,
		DateTime1:   DateTime1,
		DateTime2:   DateTime2,
		DataType:    dataType,
		IntakeA:     intakeA,
		IntakeB:     intakeB,
		IntakeC:     intakeC,
		IntakeD:     intakeD,
		IntakeDV:    intakeDV,
		IntakeTotal: intakeTotal,
		OutputA:     outputA,
		OutputB:     outputB,
		OutputC:     outputC,
		OutputCV:    outputCV,
		OutputTotal: outputTotal,
	}
	//fmt.Printf("io statistics :%+v\n\n", mod)
	_, errinsert := mod.InsertData()
	if errinsert != nil {
		c.RenderingJsonAutomatically(2, "数据出现错误")
	} else {
		c.RenderingJson(0, "成功", []interface{}{})
	}


}
