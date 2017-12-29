package handler

import (
	"fit"
	"nursing/model"
	"time"
	"fmt"
	"nursing/utils"
	"strconv"
	"errors"
)

// 体温单
type TempChartController struct {
	PCNRLController
}

var dateLables = []string{"", "", "Ⅱ-", "Ⅲ-", "Ⅳ-", "Ⅴ-", "Ⅵ-", "Ⅶ-", "Ⅷ-", "Ⅸ-", "Ⅹ-"}
//Ⅰ- "",

func (c TempChartController) LoadTable(w *fit.Response, r *fit.Request, p fit.Params) {
	defer utils.Trace("temp list")()
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "9")
	if !has {
		return
	}

	// 入院日期
	datestr := pInfo.VAE11.ParseToSecond()
	num := r.FormValue("num")

	// 手术产后日期时间model
	hospdate, err := getWeeksByOperationDates(datestr, pid, num)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}

	chartmod, errchart := model.PCGetTempChartData(pid, hospdate.weeks)
	if errchart != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", errchart)
		fmt.Fprintln(w, "参数错误！temp ", errchart)
		return
	}

	c.Data = fit.Data{
		"Userinfo": userinfo, // 护士信息
		"PInfo":    pInfo,    // 病人信息
		"Beds":     beds,     // 床位list

		"Dates1":    hospdate.dates1,    // 住院日数
		"Dates2":    hospdate.dates2,    // 手术或产后日数
		"Weeks":     hospdate.weekstr,   // 日期
		"Weeknum":   hospdate.weeknum,   // 当前病人住院周数
		"PageIndex": hospdate.weekindex, // 当前周数

		"Temp1":     chartmod.Temp1,
		"Temp2":     chartmod.Temp2,
		"Temp3":     chartmod.Temp3,
		"TempOther": chartmod.TempOther,
		"Breathe":   chartmod.Breathe,
		"Pulse":     chartmod.Pulse,
		"Heartrate": chartmod.Heartrate,
		"Pressure":  chartmod.Pressure,
		"Weight":    chartmod.Weight,
		"Output3":   chartmod.Output3,
		"Output2":   chartmod.Output2,
		"Output1":   chartmod.Output1,
		"Intake":    chartmod.Intake,
		"Incident":  chartmod.Incident,
		"Skin":      chartmod.Skin,
		"Other":     chartmod.Other,

		"Menuindex": "5-0",
	}

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
	if err != nil || len(pInfos) == 0 {
		fmt.Fprintln(w, "参数错误！")
		fit.Logger().LogError("error print temp chart :", err)
		return
	}
	pInfo := pInfos[0]

	// 入院日期
	datestr := pInfo.VAE11.ParseToSecond()

	// 手术产后日期时间model
	hospdate, err := getWeeksByOperationDates(datestr, pid, num)
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}

	chartmod, errchart := model.PCGetTempChartData(pid, hospdate.weeks)
	if errchart != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", errchart)
		fmt.Fprintln(w, "参数错误！temp ", errchart)
		return
	}

	c.Data = fit.Data{
		"PInfo":     pInfo,              // 病人信息
		"Dates1":    hospdate.dates1,    // 住院日数
		"Dates2":    hospdate.dates2,    // 手术或产后日数
		"Weeks":     hospdate.weekstr,   // 日期
		"Weeknum":   hospdate.weeknum,   // 当前病人住院周数
		"PageIndex": hospdate.weekindex, // 当前周数

		"Temp1":     chartmod.Temp1,
		"Temp2":     chartmod.Temp2,
		"Temp3":     chartmod.Temp3,
		"TempOther": chartmod.TempOther,
		"Breathe":   chartmod.Breathe,
		"Pulse":     chartmod.Pulse,
		"Heartrate": chartmod.Heartrate,
		"Pressure":  chartmod.Pressure,
		"Weight":    chartmod.Weight,
		"Output3":   chartmod.Output3,
		"Output2":   chartmod.Output2,
		"Output1":   chartmod.Output1,
		"Intake":    chartmod.Intake,
		"Incident":  chartmod.Incident,
		"Skin":      chartmod.Skin,
		"Other":     chartmod.Other,
	}

	c.LoadView(w, "pc/v_templist_print.html")
}

func (c TempChartController) Test(w *fit.Response, r *fit.Request, p fit.Params)  {
	defer c.ResponseToJson(w)
	pid := "817242"
	// 手术产后日期时间model
	hospdate, err := getWeeksByOperationDates("2017-12-20 10:02:05", pid, "1")
	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		fmt.Fprintln(w, "参数错误！temp ", err)
		return
	}
	data := model.Test(pid, hospdate.weeks)
	c.RenderingJson(1, "test", data)
}

type hospDateModel struct {
	weeks     []time.Time // 日期, 这一周内有哪些天
	weekstr   []string    // 日期, 这一周内有哪些天, 体温单显示用
	dates1    []string    // 住院日数
	dates2    []string    // 手术日数
	weeknum   int         // 入院至今天一共多少周
	weekindex int         // 当前第几周
}

/*
获取手术或产后时间  12-07 优化版
hospitaldate 入院时间
*/
func getWeeksByOperationDates(hospitaldate, pid, weekindexstr string) (hospdate hospDateModel, err error) {
	loc := time.Now().Location()
	hospitaldate = utils.Substr(hospitaldate, 0, 10)

	// 入院日期
	var t0 time.Time
	t0, err = time.ParseInLocation("2006-01-02", hospitaldate, loc)
	if err != nil {
		return
	}

	// 入院日期到今天的总周数
	weeknum := int(time.Since(t0).Hours()/24)/7 + 1
	hospdate.weeknum = weeknum
	// 当前周
	weekindex, errAtoi := strconv.Atoi(weekindexstr)
	if errAtoi != nil || weekindexstr == "" || weekindex <= 0 {
		weekindex = weeknum
	}
	hospdate.weekindex = weekindex
	offset := weekindex*7 - 7

	// 日期的第一天（第weekindex周的第一天）
	t1 := time.Date(t0.Year(), t0.Month(), t0.Day()+offset, 0, 0, 0, 0, loc)

	//手术或产后时间字符串数组
	lables, errlabel := fetchOperationLable(pid, t1)
	if errlabel != nil {
		fit.Logger().LogInfo("info templist", "获取手术时间错误！temp ", errlabel)
		err = errors.New("获取手术时间出错")
		return
	}
	hospdate.dates2 = lables

	var weeks []time.Time
	var dates1 []string
	//入院日数时间字符串数组
	nowdate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, i)
		diffnow := nowdate.Sub(t2).Hours()
		if diffnow >= 0 {
			weeks = append(weeks, t2)
			dates1 = append(dates1, fmt.Sprintln(offset+i+1))
		} else {
			dates1 = append(dates1, "")
		}
	}
	hospdate.weeks = weeks
	hospdate.dates1 = dates1

	//日期
	var weekstr []string
	for _, val := range hospdate.weeks {
		weekstr = append(weekstr, val.Format("2006-01-02"))
	}
	hospdate.weekstr = weekstr
	return hospdate, nil
}

// 获取所有手术的期间，返回在此期间的前缀排列组合
func fetchOperationLable(pid string, sdate time.Time) (strlables []string, err error) {
	var operationTimes []time.Time
	var results []string

	edate := sdate.AddDate(0, 0, 7)
	operationTimes, err = model.GetOperationRecords(pid, edate)

	if err != nil {
		fit.Logger().LogInfo("info templist", "参数错误！temp ", err)
		return nil, err
	} else if len(operationTimes) == 0 {
		strlables = []string{"", "", "", "", "", "", ""}
		return
	}
	operationTimes = append(operationTimes, edate)
	//fmt.Println("手术时间：", operationTimes)
	var t1, t2 time.Time
	var lable string
	var flag = 1
	for index := 0; index < len(operationTimes)-1; index++ {
		t1 = operationTimes[index]
		t2 = operationTimes[index+1]

		offset := int(t2.Sub(t1).Hours() / 24)
		//fmt.Println("offset:", t1, t2, "offset day:", offset, t2.Sub(t1).Hours())
		//if t2.Sub(t1).Hours() <= 24*10 {
		//fmt.Println("index:", index, "flag:",flag, "lable:", dateLables[flag - 1])

		for ii := 0; ii < offset; ii++ {
			lable = dateLables[flag]
			str := fmt.Sprintf("%s%d", lable, ii)
			if ii > 10 {
				str = ""
			}
			results = append(results, str)
			//fmt.Println("flag:", flag, lable, "str:", str, t1.AddDate(0, 0, ii))
		}
		if offset <= 10 {
			flag++
		} else {
			flag = 1
		}

	}
	fmt.Println("results:", results, "len:", len(results))
	length := len(results)
	if length >= 7 {
		strlables = results[length-7:]
	} else {
		// 第一次手术时间， 解决时间错位问题
		t0 := operationTimes[0]
		days := int(t0.Sub(sdate).Hours() / 24)
		var temp []string
		for ii := 0; ii < days; ii++ {
			temp = append(temp, "")
		}
		temp = append(temp, results...)
		strlables = temp
	}

	// 超过的当前时间的不显示
	nowdate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	diffnow := nowdate.Sub(sdate).Hours()
	if diffnow > 0 {
		offset := int(diffnow / 24)
		for ii := 6; ii > offset; ii-- {
			strlables[ii] = ""
		}
	}
	return
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
