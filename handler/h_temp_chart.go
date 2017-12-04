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
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, _, pInfo, has := c.GetBedsAndUserinfo(w, r, "9")
	if !has {
		return
	}

	// 入院日期
	datestr := pInfo.VAE11.ParseToSecond()
	weeknum, errnum := getweeknum(datestr)
	if errnum != nil {
		fit.Logger().LogError("error get week num", errnum)
	}

	num := r.FormValue("num")
	weekindex, errweek := strconv.Atoi(num)
	if errweek != nil || weekindex >= weeknum || num == "" {
		weekindex = weeknum
	}
	if weekindex <= 0 {
		weekindex = 1
	}
	fmt.Println("weekindex", weekindex, "weeknum", weeknum)


	// 手术或产后日期
	datestr2, err := model.FetchOperationRecordsDate(pInfo.VAA01)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}
	fmt.Println("patient info", pInfo)
	fmt.Println("	入院 date:", datestr)
	fmt.Println("手术或产后 date:", datestr2)


	weeks, dates1, dates2, _, err := getWeeksByOperationDates(datestr, datestr2, weekindex)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}
	//fmt.Println(weeks, dates1, dates2)
	
	var weekstr []string
	for _, val := range weeks {
		weekstr = append(weekstr, val.Format("2006-01-02"))
	}

	// 体温
	// 口温，腋温，肛温
	temp1map, errtemp1 := model.GetTemperatureChatData(1, pInfo.VAA01, weeks)
	temp2map, errtemp12 := model.GetTemperatureChatData(2, pInfo.VAA01, weeks)
	temp3map, errtemp13 := model.GetTemperatureChatData(3, pInfo.VAA01, weeks)
	tempOther, errother := model.GetTemperatureChatData(-1, pInfo.VAA01, weeks)

	// 呼吸
	breathemap, err2 := model.GetTemperatureChatData(6, pInfo.VAA01, weeks)

	// 脉搏
	pulsemap, err3 := model.GetTemperatureChatData(4, pInfo.VAA01, weeks)

	// 心率
	heartratemap, err4 := model.GetTemperatureChatData(5, pInfo.VAA01, weeks)

	// 事件
	incidentmap, err6 := model.GetTemperatureChatData(15, pInfo.VAA01, weeks)

	// 输入液量 1：入量，2：出量
	// 排出量
	// 1=其他  2=尿量  3=大
	intakearr, err5 := model.GetTemperatureChatData(7, pInfo.VAA01, weeks)
	outputarr1, err6 := model.GetTemperatureChatData(10, pInfo.VAA01, weeks)
	outputarr2, err7 := model.GetTemperatureChatData(9, pInfo.VAA01, weeks)
	outputarr3, err8 := model.GetTemperatureChatData(8, pInfo.VAA01, weeks)

	// 血压
	pressuremap, err9 := model.GetTemperatureChatData(11, pInfo.VAA01, weeks)

	// 体重
	weightmap, err10 := model.GetTemperatureChatData(12, pInfo.VAA01, weeks)

	// 皮试
	skinmap, err11 := model.GetTemperatureChatData(13, pInfo.VAA01, weeks)

	// 其他
	othermap, err12 := model.GetTemperatureChatData(14, pInfo.VAA01, weeks)

	flag := checkerr("error temp chart: ", errtemp1, errtemp12, errtemp13, errother, err2, err3, err4, err5, err6, err7, err8, err9, err10, err11, err12)
	if flag {
		return
	}

	c.Data = fit.Data{
		"Userinfo": userinfo, // 护士信息
		"PInfo":    pInfo,    // 病人信息
		"Beds":     beds,     // 床位list

		"Dates1":    dates1,  // 住院日数
		"Dates2":    dates2,  // 手术或产后日数
		"Weeks":     weekstr, // 日期
		"Weeknum":   weeknum, // 当前病人住院周数
		"PageIndex": weekindex, // 当前周数

		"Temp1":     temp1map,
		"Temp2":     temp2map,
		"Temp3":     temp3map,
		"TempOther": tempOther,
		"Breathe":   breathemap,
		"Pulse":     pulsemap,
		"Heartrate": heartratemap,
		"Pressure":  pressuremap,
		"Weight":    weightmap,
		"Output3":   outputarr3,
		"Output2":   outputarr2,
		"Output1":   outputarr1,
		"Intake":    intakearr,
		"Incident":  incidentmap,
		"Skin":      skinmap,
		"Other":     othermap,

		"Menuindex": "5-0",
	}


	//fmt.Printf("pinfo %+v\n\n userinfo: %+v\n\n", pInfo, userinfo)
	//fmt.Println(len(temp1map), c.Data)
	//c.LoadView(w, "pc/v_templist.html", "pc/header_side.html")
	success := c.LoadViewSafely(w, r, "pc/v_templist.html", "pc/header_side.html", "pc/header_top.html")
	if success == false {
		fmt.Fprintln(w, "加载错误！10002")
	}
}

func (c TempChartController) PrintTempChart(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")
	num := r.FormValue("num")
	if pid == "" || num == "" {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	pInfos, err := model.GetPatientInfo(pid)
	if err != nil || len(pInfos) == 0{
		fmt.Fprintln(w, "参数错误！")
		return
	}
	pInfo := pInfos[0]

	// 入院日期
	datestr := pInfo.VAE11.ParseToSecond()
	weeknum, errnum := getweeknum(datestr)
	if errnum != nil {
		fit.Logger().LogError("error get week num", errnum)
	}


	weekindex, errweek := strconv.Atoi(num)
	if errweek != nil || weekindex >= weeknum {
		weekindex = weeknum
	}
	if weekindex <= 0 {
		weekindex = 1
	}


	// 手术或产后日期
	datestr2, err := model.FetchOperationRecordsDate(pInfo.VAA01)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}



	weeks, dates1, dates2, _, err := getWeeksByOperationDates(datestr, datestr2, weekindex)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}

	var weekstr []string
	for _, val := range weeks {
		weekstr = append(weekstr, val.Format("2006-01-02"))
	}

	// 体温

	temp1map, err1 := model.GetTemperatureChatData(1, pInfo.VAA01, weeks)
	temp2map, err1 := model.GetTemperatureChatData(2, pInfo.VAA01, weeks)
	temp3map, err1 := model.GetTemperatureChatData(3, pInfo.VAA01, weeks)
	// 呼吸
	breathemap, err2 := model.GetTemperatureChatData(6, pInfo.VAA01, weeks)

	// 脉搏
	pulsemap, err3 := model.GetTemperatureChatData(4, pInfo.VAA01, weeks)

	// 心率
	heartratemap, err4 := model.GetTemperatureChatData(5, pInfo.VAA01, weeks)

	// 事件
	incidentmap, err6 := model.GetTemperatureChatData(15, pInfo.VAA01, weeks)

	// 输入液量 1：入量，2：出量
	// 排出量
	// 1=其他  2=尿量  3=大
	intakearr, err5 := model.GetTemperatureChatData(7, pInfo.VAA01, weeks)
	outputarr1, err6 := model.GetTemperatureChatData(10, pInfo.VAA01, weeks)
	outputarr2, err7 := model.GetTemperatureChatData(9, pInfo.VAA01, weeks)
	outputarr3, err8 := model.GetTemperatureChatData(8, pInfo.VAA01, weeks)

	// 血压
	pressuremap, err9 := model.GetTemperatureChatData(11, pInfo.VAA01, weeks)

	// 体重
	weightmap, err10 := model.GetTemperatureChatData(12, pInfo.VAA01, weeks)

	// 皮试
	skinmap, err11 := model.GetTemperatureChatData(13, pInfo.VAA01, weeks)

	// 其他
	othermap, err12 := model.GetTemperatureChatData(14, pInfo.VAA01, weeks)

	flag := checkerr("error temp chart: ", err1, err2, err3, err4, err5, err6, err7, err8, err9, err10, err11, err12)
	if flag {
		return
	}

	c.Data = fit.Data{
		"PInfo":    pInfo,    // 病人信息
		"Dates1":    dates1,  // 住院日数
		"Dates2":    dates2,  // 手术或产后日数
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
		"Incident":  incidentmap,
		"Skin":      skinmap,
		"Other":     othermap,

		"PageIndex": weekindex,
	}
	c.LoadView(w, "pc/v_templist_print.html")
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
func getWeeksByOperationDates(hospitaldate string, operationTimes []string, weekindex int) (weeks []time.Time, dates1, dates2 []string, weeknum int, err error) {
	loc := time.Now().Location()
	hospitaldate = utils.Substr(hospitaldate, 0, 10)
	operationstr1 := ""
	operationstr2 := ""
	//operationstr3 := ""
	switch len(operationTimes) {
	case 0:
	case 1:
		operationstr1 = operationTimes[0]
	case 2:
		operationstr1 = operationTimes[0]
		operationstr2 = operationTimes[1]
	default:
		operationstr1 = operationTimes[0]
		operationstr2 = operationTimes[1]
		//operationstr3 = operationTimes[2]
	}
	operationstr1 = utils.Substr(operationstr1, 0, 10)
	operationstr2 = utils.Substr(operationstr2, 0, 10)
	//operationstr3 = utils.Substr(operationstr3, 0, 10)

	//fmt.Println("hehe------", operationstr1, operationstr2)

	// 入院日期
	t0, err := time.ParseInLocation("2006-01-02", hospitaldate, loc)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	// 入院日期到今天的总周数
	weeknum = int(time.Since(t0).Hours()/24)/7 + 1


	offset := weekindex*7 - 7
	if weekindex == 0 {
		offset = weeknum*7 - 7
	}

	t1 := time.Date(t0.Year(), t0.Month(), t0.Day()+offset, 0, 0, 0, 0, loc)
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, +i)
		weeks = append(weeks, t2)
		dates1 = append(dates1, fmt.Sprintln(offset+i+1))
		// 手术后或产后日期
		var operatime1 time.Time
		var operatime2 time.Time
		//var operatime3 time.Time
		if operationstr1 != "" {
			operatime1, err = time.ParseInLocation("2006-01-02", operationstr1, loc)
			if err != nil {
				operatime1 = time.Now().AddDate(1, 0, 0)
				//fmt.Println("------ hos :", operatime1, err)
			}

			if operationstr2 != "" {
				operatime2, err = time.ParseInLocation("2006-01-02", operationstr2, loc)
				if err != nil {
					operatime2 = time.Now().AddDate(1, 0, 0)
					//fmt.Println("------ hos :", operatime2, err)
				}
			}
			//if operationstr3 != "" {
			//	operatime3, err = time.ParseInLocation("2006-01-02", operationstr3, loc)
			//	if err != nil {
			//		operatime3 = time.Now().AddDate(1, 0, 0)
			//		fmt.Println("------ hos :", operatime3, err)
			//	}
			//}
			//fmt.Println("success", operatime1, operatime2)

			operaoffset1 := t2.Sub(operatime1).Hours()

			difftime1 := operatime2.Sub(operatime1).Hours()
			//difftime2 := operatime3.Sub(operatime2).Hours()
			//fmt.Println("offset :", operaoffset1, difftime1)

			if operaoffset1 >= 0 && operaoffset1 < 24 * 10 { // 手术十天以内记录时间
				if difftime1 < 0 || difftime1 > 24 * 10 {
					operstr := fmt.Sprintf(" %.0f", operaoffset1/24)
					fmt.Println("--------------------", operstr)
					dates2 = append(dates2, operstr)
				} else { // 十天以内做过第二次手术 则记为 Ⅱ-1
					operaoffset2 := t2.Sub(operatime2).Hours()
					if operaoffset2 >= 0 && operaoffset2 < 24 * 10 {
						aaa := int(operaoffset2/24)
						dates2 = append(dates2, fmt.Sprintf("Ⅱ-%d",aaa))
					} else {
						dates2 = append(dates2, fmt.Sprintf(" %.0f", operaoffset1/24))
					}
				}
			} else {
				str := " "
				if difftime1 < 24 * 10 {
					str = "Ⅱ-"
				}
				operaoffset2 := t2.Sub(operatime2).Hours()
				if operaoffset2 >= 0 && operaoffset2 <= 24 * 10 {
					aaa := int(operaoffset2/24)
					dates2 = append(dates2, fmt.Sprintf("%s%d", str,aaa))
				} else {
					dates2 = append(dates2, "")
				}
			}
		} else {
			for ii := 0; ii < 7; ii++ {
				dates2 = append(dates2, "")
			}
			operatime1 = time.Now().AddDate(1, 0, 0)
			//fmt.Println("hos :", operatime1)
		}

	}
	fmt.Println("date2 :", dates2, len(dates2))
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


//当前周的手术或产后日期
/*func getWeeksByDatestr(datestr, datestr2 string, weekindex int) (weeks []time.Time, dates1, dates2 []int, weeknum int, err error) {
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
}*/
