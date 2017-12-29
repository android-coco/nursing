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
	pid := r.FormInt64Value("pid")
	// 护士ID
	uid := r.FormIntValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := r.FormTimeStruct("datetime")

	if err4 != nil || pid == 0 || uid == 0 || BCE03A == "" {
		c.RenderingJsonAutomatically(10002, "参数不完整")
		return
	}
	//fmt.Println(VAA01, BCK01, BCE01A, BCE03A, datetime)

	jsonstr := r.FormValue("jsonstr")
	var mods []model.NRLData
	errUnmarshal := json.Unmarshal([]byte(jsonstr), &mods)
	if errUnmarshal != nil {
		fit.Logger().LogError("json.Unmarshal err:", errUnmarshal)
		c.RenderingJsonAutomatically(10003, "参数错误")
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
			mod.TestTime = model.FitTime(datetime)
			mod.PatientId = pid
			mod.NurseId = uid
			mod.NurseName = BCE03A

			_, errInsert := session.Table("NurseChat").Insert(&mod)
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
}

func (c NRL1Controller) UpdateTitle(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	mod := model.NRL1Title{}
	c.FitSetStruct(&mod, r)

	if mod.PatientId == 0 {
		c.RenderingJsonAutomatically(10002, "参数不完整")
		return
	}

	errt := mod.PCUpdateNRT1Title()
	if errt != nil {
		fit.Logger().LogError("hy:", "nrl1 update :", errt)
		c.RenderingJson(2, "DB  错误！", errt.Error())
	} else {
		c.RenderingJsonAutomatically(0, "添加成功！")
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
				modio := model.IOStatistics{ID: id}

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

// 护理记录单 PC端
type PCNRL1Controller struct {
	PCNRLController
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
		//hospDate := time.Time(pInfo.VAE11)
		//local1, _ := time.LoadLocation("Local")
		//local2 := time.Local
		//fmt.Println("local:", hospDate, hospDate.In(local1), local1.String(), local2.String())

		date1 := time.Time(pInfo.VAE11).Unix() * 1000
		date2 := time.Now().Unix() * 1000
		paramstr := fmt.Sprintf("&sdate=%d&edate=%d", date1, date2)
		datestr1 = pInfo.VAE11.ParseToSecond()
		datestr2 = time.Now().Format("2006-01-02 15:04:05")
		urlstr := r.URL.String() + paramstr
		fmt.Println("-------str:", paramstr, "datestr:", datestr1, datestr2)
		c.Redirect(w, r, urlstr, 302)
		return
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
		ls := mods[(pageindex-1)*peerPage:count]
		length := len(ls)
		for index := 0; index < peerPage; index++ {
			if index < length {
				list[index] = ls[index]
			} else {
				list[index] = model.NRLModel{
					State:     "empty",
					PatientId: pInfo.VAA01,
					NurseId:   strconv.Itoa(userinfo.UID),
					NurseName: userinfo.Name,
				}
			}

		}
	} else {
		list = mods[(pageindex-1)*peerPage:pageindex*peerPage]
	}


	nrl1Title := model.NRL1Title{PatientId: pInfo.VAA01}
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

	c.LoadViewSafely(w, r, "pcnrl/v_nrl1.html", "pc/header_side.html", "pc/header_top.html")
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
	_, err5 := strconv.Atoi(did)

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
		PatientId: VAA01,
		//BCK01:       BCK01,
		NurseId:     BCE01A,
		NurseName:   BCE03A,
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
