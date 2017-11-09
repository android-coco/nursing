package handler

import (
	"fit"
	"nursing/model"
	"time"
	"fmt"
	"nursing/utils"
	"strconv"

)

// 体温单
type TempChartController struct {
	PCController
}



func (c TempChartController) LoadTable(w *fit.Response, r *fit.Request, p fit.Params) {

	//date := r.FormValue("date")
	//date, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-22 00:00:00", time.Now().Location())

	// 护士信息
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
	}
	beds, err12 := model.QueryDepartmentBeds(userinfo.DepartmentID, false)


	num := r.FormValue("num")
	// 病人id
	var pInfo model.PCBedDup
	pid := r.FormValue("pid")
	if pid == "" || num == "" { // 默认去当前科室第一个床位病人
		if pid == "" {
			pidnum := beds[0].VAA01
			pid = strconv.Itoa(pidnum)
		}

		// 病人信息
		for _, val := range beds {
			if strconv.Itoa(val.VAA01) == pid {
				pInfo = val
				break
			}
		}
		fmt.Println("p info ", pInfo)
		// 入院日期
		datestr := pInfo.VAA73.NormParse()

		weeknum := 1
		if num == "" {
			var errnum error
			weeknum, errnum = getweeknum(datestr)
			if errnum != nil {
				fit.Logger().LogError("error get week num", errnum)
			}
		}

		url := "/pc/templist?pid=" + pid + "&num=" + strconv.Itoa(weeknum)
		c.Redirect(w, r, url, 302)
		return
	}


	weekindex, errweek := strconv.Atoi(num)
	if errweek != nil {
		weekindex = 0
	}
	if weekindex <= 0 {
		weekindex = 1
	}

	// 病人信息
	for _, val := range beds {
		if strconv.Itoa(val.VAA01) == pid {
			pInfo = val
			break
		}
	}
	if pInfo.VAA01 == 0 {
		//url := "/pc/templist"
		//c.Redirect(w, r, url, 302)
		fmt.Fprintln(w, "参数错误！  user info error", err)
	}
	weeknum, errnum := getweeknum(pInfo.VAA73.NormParse())
	if errnum != nil {
		fit.Logger().LogError("error get week num", errnum)
	}
	fmt.Println("weekindex", weekindex, "weeknum", weeknum)
	if weekindex >= weeknum {
		weekindex = weeknum
	}
	// 入院日期
	datestr := pInfo.VAA73.NormParse()
	// 手术或产后日期
	datestr2, err := model.GetOperationDate(pid)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
	}
	fmt.Println("patient info", pInfo)
	fmt.Println("	入院 date:", datestr)
	fmt.Println("operation date:", datestr2)

	weeks, dates1, dates2, weeknum, err := getWeeksByDatestr(datestr, datestr2, weekindex)
	fmt.Println(weeks)
	fmt.Println(dates1, dates2)
	fmt.Println(weeknum, err)
	if weekindex == 0 {
		weekindex = weeknum
	}



	var weekstr []string
	for _, val := range weeks {
		weekstr = append(weekstr, val.Format("2006-01-02"))
	}



	// 体温
	temp1map, err1 := model.GetTempChart("101", pid, weeks)
	temp2map, err1 := model.GetTempChart("102", pid, weeks)
	temp3map, err1 := model.GetTempChart("103", pid, weeks)
	// 呼吸
	breathemap, err2 := model.GetTempChart("2", pid, weeks)

	// 脉搏
	pulsemap, err3 := model.GetTempChart("3", pid, weeks)

	// 心率
	heartratemap, err4 := model.GetTempChart("4", pid, weeks)

	// 事件
	incidentmap, err6 := model.GetTempChart("5", pid, weeks)

	// 输入液量 1：入量，2：出量
	// 排出量
	// 1=其他  2=尿量  3=大
	intakearr, err5 := model.GetTempChart("601", pid, weeks)
	outputarr1, err6 := model.GetTempChart("611", pid, weeks)
	outputarr2, err7 := model.GetTempChart("612", pid, weeks)
	outputarr3, err8 := model.GetTempChart("613", pid, weeks)

	// 血压
	pressuremap, err9 := model.GetTempChart("7", pid, weeks)

	// 体重
	weightmap, err10 := model.GetTempChart("8", pid, weeks)

	// 皮试
	skinmap, err11 := model.GetTempChart("9", pid, weeks)

	if err1 != nil || err2 != nil || err3 != nil ||
		err4 != nil || err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil ||
		err10 != nil || err11 != nil || err12 != nil || err12 != nil {
		fit.Logger().LogError("error", "temp chart", err1)
		fit.Logger().LogError("error", "temp chart", err2)
		fit.Logger().LogError("error", "temp chart", err3)
		fit.Logger().LogError("error", "temp chart", err4)
		fit.Logger().LogError("error", "temp chart", err5, err6, err7, err8, err9)
		fit.Logger().LogError("error", "temp chart", err6)
		fit.Logger().LogError("error", "temp chart", err9)
		fit.Logger().LogError("error", "temp chart", err10)
		fit.Logger().LogError("error", "temp chart", err11)
		fit.Logger().LogError("error", "temp chart", err12)
		fit.Logger().LogError("error", "temp chart", err12)
		fmt.Fprintln(w, "参数错误！  patient data error")
		return
	}


	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo, // 病人信息
		"Beds":      beds, // 床位list

		"Dates1":    dates1, // 住院日数
		"Dates2":    dates2, // 手术或产后日数
		"Weeks":     weekstr, // 日期
		"Weeknum":   weeknum, // 当前病人住院周数
		"Temp1":     temp1map,
		"Temp2":     temp2map,
		"Temp3":     temp3map,
		"Breathe":   breathemap,
		"Pulse":     pulsemap,
		"Heartrate": heartratemap,
		"Pressure":  pressuremap,
		"Weight":    weightmap,
		"Output3":   outputarr3,
		"Output2":   outputarr2,
		"Output1":   outputarr1,
		"Intake":    intakearr,
		"Incident": incidentmap,
		"Skin":     skinmap,
		"Menuindex": "5-0",
	}

	//fmt.Printf("beds %+v\n\n %+v\n\n", beds, incidentmap)
	fmt.Println(len(incidentmap))
	//c.LoadView(w, "pc/v_templist.html", "pc/header_side.html")
	success := c.LoadViewSafely(w, r, "pc/v_templist.html", "pc/header_side.html", "pc/header_top.html")
	if success == false {
		fmt.Fprintln(w,"加载错误！10002")
	}
}

/**
datestr 入院时间
datestr2 手术时间
weekindex 第几周

## return
weeks 日期
dates1 住院日数
dates2 手术日数
weeknum 入院至今天一共多少周

 */
func getWeeksByDatestr(datestr, datestr2 string, weekindex int) (weeks []time.Time, dates1, dates2 []int, weeknum int, err error) {
	loc := time.Now().Location()
	datestr = utils.Substr(datestr, 0, 10)
	datestr2 = utils.Substr(datestr2, 0, 10)
	t, err := time.ParseInLocation("2006-01-02", datestr, loc)
	if err != nil {
		return nil, nil, nil, 0, err
	}
	fmt.Println("weeks datestr", datestr, datestr2)
	var hostime time.Time
	if datestr2 != "" {
		hostime, err = time.ParseInLocation("2006-01-02", datestr2, loc)
		//hostime.Add(time.Duration(60))
		if err != nil || t.Sub(hostime) > 0 {
			hostime = time.Now().AddDate(1, 0, 0)
			fmt.Println("------ hos :", hostime)
		}
		fmt.Println("success", hostime)

	} else {
		hostime = time.Now().AddDate(1, 0, 0)
		fmt.Println("hos :", hostime)
	}

	fmt.Println("weeks date", t, hostime)
	// 入院日期到今天的总周数
	weeknum = int(time.Since(t).Hours()/24)/7 + 1

	//weekoffset := int(time.Since(t) / 24) / 7
	fmt.Println("weeknum", weeknum, t)

	offset := weekindex*7 - 7
	if weekindex == 0 {
		offset = weeknum*7 - 7
	}

	t1 := time.Date(t.Year(), t.Month(), t.Day()+offset, 0, 0, 0, 0, loc)
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, +i)
		//fmt.Println("time ---", t2.String())
		//fmt.Println(t2.Sub(hostime).Hours() / 24)
		weeks = append(weeks, t2)
		dates1 = append(dates1, offset+i+1)
		// 手术后或产后日期
		hosoffset := t2.Sub(hostime).Hours()
		if hosoffset >= 0 {
			dates2 = append(dates2, int(hosoffset/24)+1)
		} else {
			dates2 = append(dates2, 0)
		}
	}
	return weeks, dates1, dates2, weeknum, nil
}

func getweeknum(datestr string) (weeknum int, err error) {
	loc := time.Now().Location()
	datestr = utils.Substr(datestr, 0, 10)
	t, err := time.ParseInLocation("2006-01-02", datestr, loc)
	if err != nil {
		return 0, err
	}
	// 入院日期到今天的总周数
	weeknum = int(time.Since(t).Hours()/24)/7 + 1
	return weeknum, nil
}

func GetWeeksByDate(datestr string) ([]time.Time, error) {
	loc := time.Now().Location()
	t, err := time.ParseInLocation("2006-01-02 15:04:05", datestr, loc)
	if err != nil {
		return nil, err
	}
	//fmt.Println(time.Since(t).Hours() / 24, time.Since(t))
	weeknum := int(time.Since(t).Hours()/24) / 7
	//weekoffset := int(time.Since(t) / 24) / 7
	//fmt.Println("weeknum", weeknum, t)

	t1 := time.Date(t.Year(), t.Month(), t.Day()+7*weeknum, 0, 0, 0, 0, loc)
	var weeks []time.Time
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, +i)
		fmt.Println("time ---", t2.String())
		weeks = append(weeks, t2)
	}
	return weeks, nil
}


//当前周的手术或产后日期

func GetWeeks(datestr string) []time.Time {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", datestr, time.Now().Location())
	t2, err2 := time.ParseInLocation("2006-01-02 15:04:05", datestr, loc)
	t3, _ := time.Parse("2006-01-02 15:04:05", datestr)
	if err != nil || err2 != nil {
		fit.Logger().LogError("error", "temp chart", err, err2)
	}
	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(t3)

	index := t.Weekday()
	fmt.Println(time.Now())

	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	var weeks []time.Time
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, -int(index)+i)
		fmt.Println("time:", t2.String())
		weeks = append(weeks, t2)
	}
	return weeks
}

func (c TempChartController) TempChart(w *fit.Response, r *fit.Request, p fit.Params) {
	defer utils.Trace("chart")()
	defer c.ResponseToJson(w)
	// 病人id
	pid := r.FormValue("pid")

	datestr, err := model.GetOperationDate(pid)
	if err != nil {
		fmt.Fprintln(w, "参数错误！")
	}
	fmt.Println("operation date:", datestr)

	date := r.FormValue("date")
	//weeks, _ := GetWeeksByDate(date)
	weeks := GetWeeks(date)

	//datetime, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Now().Location())

	/*// 体温
	temp := model.Temperature{}
	tempmap, err1 := temp.GetTemp(pid, weeks)

	// 呼吸
	breathemap, err2 := model.GetBreathe(pid, weeks)

	// 脉搏
	pulsemap, err3 := model.GetPulse(pid, weeks)

	// 心率
	heartratemap, err4 := model.GetHeartrate(pid, weeks)

	// 输入液量
	intakemap, err5 := model.GetIntake(pid, weeks)

	// 排出量
	// 1=其他  2=尿量  3=大便
	output1map, err6 := model.GetOutput1(pid, weeks)
	output2map, err7 := model.GetOutput2(pid, weeks)
	output3map, err8 := model.GetOutput3(pid, weeks)

	// 血压
	perssuremap, err9 := model.GetPressure(pid, weeks)

	// 体重
	weightmap, err10 := model.GetPressure(pid, weeks)

	// 皮试
	skinmap, err11 := model.GetSkin(pid, weeks)*/
	// 体温
	temp1map, err1 := model.GetTempChart("101", pid, weeks)
	temp2map, err1 := model.GetTempChart("102", pid, weeks)
	temp3map, err1 := model.GetTempChart("103", pid, weeks)

	// 呼吸
	breathemap, err2 := model.GetTempChart("2", pid, weeks)

	// 脉搏
	pulsemap, err3 := model.GetTempChart("3", pid, weeks)

	// 心率
	heartratemap, err4 := model.GetTempChart("4", pid, weeks)

	// 事件
	incidentmap, err6 := model.GetTempChart("6", pid, weeks)

	// 输入液量 1：入量，2：出量
	// 排出量
	// 1=其他  2=尿量  3=大
	//intakemap, err5 := model.GetTempChart("5", pid, weeks)

	// 血压
	perssuremap, err9 := model.GetTempChart("7", pid, weeks)

	// 体重
	weightmap, err10 := model.GetTempChart("8", pid, weeks)

	// 皮试
	skinmap, err11 := model.GetTempChart("9", pid, weeks)

	if err1 != nil || err2 != nil || err3 != nil ||
		err4 != nil || err9 != nil ||
		err10 != nil || err11 != nil || err != nil {
		fit.Logger().LogError("error", "temp chart", err1)
		fit.Logger().LogError("error", "temp chart", err2)
		fit.Logger().LogError("error", "temp chart", err3)
		fit.Logger().LogError("error", "temp chart", err4)
		//fit.Logger().LogError("error", "temp chart", err5)
		fit.Logger().LogError("error", "temp chart", err6)
		fit.Logger().LogError("error", "temp chart", err9)
		fit.Logger().LogError("error", "temp chart", err10)
		fit.Logger().LogError("error", "temp chart", err11)
		return
	}

	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "测试!"

	/*var intakearr, outputarr1, outputarr2, outputarr3 []map[string]string
	var nullmap map[string]string
	for _, val := range intakemap {

		if val != nil {
			if "1" == val["type"] {
				intakearr = append(intakearr, val)
				outputarr1 = append(outputarr1, nullmap)
				outputarr2 = append(outputarr2, nullmap)
				outputarr3 = append(outputarr3, nullmap)
			} else if "2" == val["type"] {
				intakearr = append(intakearr, nullmap)
				// 1=其他  2=尿量  3=大便
				if "1" == val["subtype"] {
					outputarr1 = append(outputarr1, val)
					outputarr2 = append(outputarr2, nullmap)
					outputarr3 = append(outputarr3, nullmap)
				} else if "2" == val["subtype"] {
					outputarr1 = append(outputarr1, nullmap)
					outputarr2 = append(outputarr2, val)
					outputarr3 = append(outputarr3, nullmap)
				} else if "3" == val["subtype"] {
					outputarr1 = append(outputarr1, nullmap)
					outputarr2 = append(outputarr2, nullmap)
					outputarr3 = append(outputarr3, val)
				}
			}
		} else {
			intakearr = append(intakearr, val)
			outputarr1 = append(outputarr1, val)
			outputarr2 = append(outputarr2, val)
			outputarr3 = append(outputarr3, val)
		}
	}*/
	//fmt.Println("----------", intakearr, outputarr1, outputarr2, outputarr3)
	//fmt.Println(tempmap,breathemap, pulsemap, heartratemap, output1map, intakemap, perssuremap, weightmap, skinmap)

	c.JsonData.Datas = map[string][]string{
		"incident":  incidentmap,
		"Temp1":     temp1map,
		"Temp2":     temp2map,
		"Temp3":     temp3map,
		"breathe":   breathemap,
		"pulse":     pulsemap,
		"heartrate": heartratemap,
		//"intake":    intakearr,
		//"output1":   outputarr1,
		//"output2":   outputarr2,
		//"output3":   outputarr3,
		"perssure": perssuremap,
		"weight":   weightmap,
		"skin":     skinmap,
	}

	//success := c.LoadViewSafely(w, r, "pc/v_templist", "pc/header_side.html")
	//if success == false {
	//	fmt.Fprintln(w,"加载错误！10002")
	//}
}