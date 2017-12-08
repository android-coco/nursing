package model

import (
	"time"
	"fit"
	"strconv"
	"errors"
	"nursing/utils"
)

type tempChart struct {
	Temp1     []string // 口表
	Temp2     []string // 腋表
	Temp3     []string // 肛表
	TempOther []string // 体温事件
	Pulse     []string // 脉搏
	Heartrate []string // 心率
	Breathe   []string // 呼吸
	Intake    []string // 输入液量
	Output1   []string // 排出其他
	Output2   []string // 排出尿量
	Output3   []string // 排出大便
	Pressure  []string // 血压
	Weight    []string // 体重
	Skin      []string // 皮试
	Other     []string // 其他
	Incident  []string // 事件
}

// 体温单

/*
-1=体温事件
1=口温
2=腋温
3=肛温
4=脉搏
5=心率
6=呼吸
7=输入液量
8=排出大便
9=排出尿量
10=排出其他
11=血压
12=体重
13=皮试
14=其他
15=事件
*/
//获取体温表数据根据周
func PCGetTempChartData(pid string, weeks []time.Time) (chartmod tempChart, err error) {
	// 体温
	temp1map, err1 := GetTempChart("101", pid, weeks)
	temp2map, err1 := GetTempChart("102", pid, weeks)
	temp3map, err1 := GetTempChart("103", pid, weeks)
	tempOthermap, err1 := GetTempChart("109", pid, weeks)

	// 脉搏
	pulsemap, err3 := GetTempChart("2", pid, weeks)
	// 呼吸
	breathemap, err2 := GetTempChart("3", pid, weeks)
	// 血压
	pressuremap, err9 := GetTempChart("4", pid, weeks)
	// 心率
	heartratemap, err4 := GetTempChart("5", pid, weeks)
	// 体重
	weightmap, err10 := GetTempChart("8", pid, weeks)
	// 皮试
	skinmap, err11 := GetTempChart("10", pid, weeks)
	// 事件
	incidentmap, err6 := GetTempChart("12", pid, weeks)

	// 输入液量 1：入量，2：出量
	// 排出量
	// 1=其他  2=尿量  3=大
	intakearr, err5 := GetTempChart("601", pid, weeks)
	outputarr1, err6 := GetTempChart("611", pid, weeks)
	outputarr2, err7 := GetTempChart("612", pid, weeks)
	outputarr3, err8 := GetTempChart("613", pid, weeks)

	// 其他
	othermap, err12 := GetTempChart("14", pid, weeks)

	if err1 != nil || err2 != nil || err3 != nil ||
		err4 != nil || err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil ||
		err10 != nil || err11 != nil || err12 != nil {
		fit.Logger().LogError("error", "temp chart", err1)
		fit.Logger().LogError("error", "temp chart", err2)
		fit.Logger().LogError("error", "temp chart", err3)
		fit.Logger().LogError("error", "temp chart", err4)
		fit.Logger().LogError("error", "temp chart", err5, err6, err7, err8, err9)
		fit.Logger().LogError("error", "temp chart", err6)
		fit.Logger().LogError("error", "temp chart", err9)
		fit.Logger().LogError("error", "temp chart", err10)
		fit.Logger().LogError("error", "temp chart", err11, err12)
		err = errors.New("query temp chart error")
		return
	}
	chartmod = tempChart{
		Temp1:     temp1map,
		Temp2:     temp2map,
		Temp3:     temp3map,
		TempOther: tempOthermap,
		Breathe:   breathemap,
		Pulse:     pulsemap,
		Heartrate: heartratemap,
		Incident:  incidentmap,
		Intake:    intakearr,
		Output1:   outputarr1,
		Output2:   outputarr2,
		Output3:   outputarr3,
		Pressure:  pressuremap,
		Weight:    weightmap,
		Skin:      skinmap,
		Other:     othermap,
	}
	return
}


func GetTempChart(ty, pid string, weeks []time.Time) ([]string, error) {
	//fit.MySqlEngine().ShowSQL(true)

	var sql  = ""
	var results []string
	var jmax int
	switch ty {
	case "109": // 体温事件
		jmax = 6
		sql = "SELECT Other FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1"
	case "101": // 口温
		jmax = 6
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and SubType = 3"
	case "102": // 腋温
		jmax = 6
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and SubType = 1"
	case "103": // 肛温
		jmax = 6
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 1 and SubType = 4"

	case "2": // 脉搏
		jmax = 6
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 2 "

	case "3": // 呼吸
		jmax = 6
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 3 "

	case "4": // 血压
		jmax = 1
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 4 "

	case "5": // 心率
		jmax = 6
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 5 "

	case "8": // 体重
		jmax = 1
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 8 "
	case "10": // 皮试
		jmax = 1
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 10 "

	case "12": // 事件
		jmax = 6
		sql = "SELECT Other FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 12 "

		//  ORDER BY DateTime1 DESC LIMIT 1
	case "601": //   总入量
		jmax = 1
		sql = "SELECT IntakeTotal AS Value FROM IOStatistics WHERE date_format(DateTime1,'%Y-%m-%d') = ? and VAA01 = ? and DataType = 1 ORDER BY DateTime1 DESC "
	case "611": // 出量其他
		jmax = 1
		sql = "SELECT OutputC AS Value FROM IOStatistics WHERE date_format(DateTime1,'%Y-%m-%d') = ? and VAA01 = ? and DataType = 1 ORDER BY DateTime1 DESC "

	case "612": // 出量 排尿/ml,
		jmax = 1
		sql = "SELECT OutputA AS Value FROM IOStatistics WHERE date_format(DateTime1,'%Y-%m-%d') = ? and VAA01 = ? and DataType = 1 ORDER BY DateTime1 DESC "

	case "613": // 出量 大便/次',
		jmax = 1
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 13 "

	case "14": // 其他
		jmax = 1
		sql = "SELECT Value FROM TemperatrureChat WHERE DateTime = ? and PatientId = ? and HeadType = 14 "

	default:
		sql = ""
	}

	if 6 == jmax {
		sql = sql + " and TypeTime = ?"
	}
	for i := 0; i < len(weeks); i++ {
		day := weeks[i]
		daystr := day.Format("2006-01-02")
		for j := 0; j < jmax; j++ {
			var resultmap []map[string]string
			var err error
			if jmax == 6 {
				resultmap, err = fit.MySqlEngine().QueryString(sql, daystr, pid, 4+j*4)
			} else {
				resultmap, err = fit.MySqlEngine().QueryString(sql, daystr, pid)
			}
			//if ty == "601" || ty == "611" || ty == "612" || ty == "612" {
			//	fit.MySqlEngine().ShowSQL(true)
			//	fmt.Println("resultmap:", ty, resultmap, results, len(results))
			//}
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil, err
			} else {
				if len(resultmap) > 0 {

					switch ty {
					case "12"://事件可以多次
						totalvalue := ""
						for _, dict := range resultmap {
							if v, ok := dict["Other"]; ok {
								switch v {
								case "1":
									totalvalue = totalvalue + "入院<br/><br/>"
								case "2":
									totalvalue = totalvalue + "出院<br/><br/>"
								case "3":
									totalvalue = totalvalue + "手术<br/><br/>"
								case "4":
									totalvalue = totalvalue + "分娩<br/><br/>"
								case "5":
									totalvalue = totalvalue + "出生<br/><br/>"
								case "6":
									totalvalue = totalvalue + "转入<br/><br/>"
								case "7":
									totalvalue = totalvalue + "转科<br/><br/>"
								case "8":
									totalvalue = totalvalue + "<span>转院<br/><br/></span>"
								case "9":
									totalvalue = totalvalue + "死亡<br/><br/>"
								case "10":
									totalvalue = totalvalue + "外出<br/><br/>"
								default:
									totalvalue = totalvalue + ""
								}
							} else {
								totalvalue = totalvalue + ""
							}
						}
						results = append(results, totalvalue)

					case "109": // 体温事件
						dict := resultmap[0]
						if v, ok := dict["Other"]; ok {
							switch v {
							case "1":
								results = append(results, "物理降温 ")
							case "2":
								results = append(results, "药物降温 ")
							case "3":
								results = append(results, "冰毯降温 ")
							case "4":
								results = append(results, "停冰毯降温 ")
							case "5":
								results = append(results, "药物+物理降温 ")
							case "6":
								results = append(results, "无降温 ")
							case "7":
								results = append(results, "不升 ")
							case "8":
								results = append(results, "外出 ")
							case "9":
								results = append(results, "检查 ")
							case "10":
								results = append(results, "请假 ")
							case "11":
								results = append(results, "拒试 ")
							case "12":
								results = append(results, "无法侧 ")
							case "13":
								results = append(results, "未测 ")
							default:
								results = append(results, "")
							}
						} else {
							results = append(results, "")
						}

					case "613":// 大便一天多次需要合计
						totalvalue := 0
						for _, dict := range resultmap {
							if v, ok := dict["Value"]; ok {
								it,_:=strconv.Atoi(v)
								totalvalue = totalvalue + it
							} else {
								totalvalue = totalvalue + 0
							}
						}
						results = append(results, strconv.Itoa(totalvalue))

					default:
						dict := resultmap[0]
						if v, ok := dict["Value"]; ok {
							results = append(results, v)
						} else {
							results = append(results, "")
						}
					}

				} else {
					results = append(results, "")
				}
			}
		}
	}
	fit.MySqlEngine().ShowSQL(false)

	//fmt.Println("result:", ty, results, len(results))
	return results, nil
}


/*查询住院期间的手术记录
func FetchOperationRecordsDate(pid int64) ([]string, error) {
	records := make([]string, 0)

	var err error
	// VAT04 = 4 表示已结束手术
	//err = fit.SQLServerEngine().SQL("select VAT08 from VAT1 where VAA01 = ? and VAT04 = 4", pid).Find(&records) .Desc("VAT08")
	resultsmap, err := fit.SQLServerEngine().QueryString("select top 2 VAT08 from VAT1 where VAA01 = ? and VAT04 = 4 ORDER BY VAT08 DESC", pid)
	//err = fit.SQLServerEngine().Table("VAT1").Select("VAT08").Where("VAA01 = ? and VAT04 = 4", pid).Limit(2, 1).Find(&records)  mssql  limit无效

	for _, val := range resultsmap {
		records = append(records, val["VAT08"])
	}
	if len(records) == 2 {
		records[1], records[0] = records[0], records[1]
	}
	return records, err
}
*/
// 获取一周内有效的手术记录
func GetValidOperationDates(pid string, sdate time.Time) ([]string, error) {
	records := make([]string, 0)
	date1 := sdate.Format("2006-01-02")
	date2 := sdate.AddDate(0,0,8).Format("2006-01-02")
	var err error
	// VAT04 = 4 表示已结束手术
	resultsmap, err := fit.SQLServerEngine().QueryString("select VAT08 from VAT1 where VAA01 = ? and VAT08 >= ? and VAT08 < ? and VAT04 = 4 ORDER BY VAT08 ASC", pid, date1, date2)
	for _, val := range resultsmap {
		datestr := utils.Substr(val["VAT08"], 0, 10)
		records = append(records, datestr)
	}
	return records, err
}
// 获取离当前周最近的一次手术
func GetLastOperationDates(pid string, sdate time.Time)  (lastRecords string, err error) {
	date2 := sdate.AddDate(0,0,0).Format("2006-01-02")
	// VAT04 = 4 表示已结束手术
	resultsmap, err := fit.SQLServerEngine().QueryString("select top 1 VAT08 from VAT1 where VAA01 = ? and VAT08 < ? and VAT04 = 4 ORDER BY VAT08 DESC", pid, date2)
	if len(resultsmap) == 0 {
		return "", nil
	}
	lastRecords = utils.Substr(resultsmap[0]["VAT08"], 0, 10)
	return
}

// 获取全部手术记录
func GetOperationRecords(pid string, edate time.Time) ([]time.Time, error) {
	records := make([]time.Time, 0)
	//date1 := sdate.AddDate(0,0, -10).Format("2006-01-02")
	date2 := edate.Format("2006-01-02")
	var err error
	// VAT04 = 4 表示已结束手术
	resultsmap, err := fit.SQLServerEngine().QueryString("select VAT08 from VAT1 where VAA01 = ? and VAT08 < ? and VAT04 = 4 ORDER BY VAT08 ASC", pid, date2)
	for _, val := range resultsmap {
		datestr := utils.Substr(val["VAT08"], 0, 10)
		date, err := time.ParseInLocation("2006-01-02", datestr, time.Local)
		if err != nil || date.Year() < 2000 {
			return nil, err
		}
		records = append(records, date)
	}
	return records, err
}


