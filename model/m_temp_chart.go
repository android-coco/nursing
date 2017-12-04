package model

import (
	"time"
	"fit"
)

// 体温单

/*
1=体温
2=呼吸
3=脉搏
4=心率
5=事件
6=输入量/排出量
7=血压
8=体重
9=皮试
*/
func GetTempChart(ty, pid string, weeks []time.Time) ([]string, error) {
	//fit.MySqlEngine().ShowSQL(true)
	loc, _ := time.LoadLocation("Local")
	var tablename, condition string = "", ""
	var results []string
	var jmax int
	switch ty {
	case "101":
		jmax = 6
		tablename = "SELECT Value FROM `TemperatrureChat` "
		condition = "AND `HeadType` = 1 AND `SubType` = 1" // 腋温
	case "102":
		jmax = 6
		tablename = "SELECT value FROM `TemperatrureChat` "
		condition = "AND `HeadType` = 1 AND `SubType` = 3" // 口温
	case "103":
		jmax = 6
		tablename = "SELECT value FROM `TemperatrureChat` "
		condition = "AND `HeadType` = 1 AND `SubType` = 4" // 肛温
	case "2": // 呼吸
		jmax = 6
		tablename = "SELECT value FROM `TemperatrureChat` "
		condition = "AND `HeadType` = 3"
	case "3": // 脉搏
		jmax = 6
		tablename = "SELECT value FROM `TemperatrureChat` "
	case "4": // 心率
		jmax = 6
		tablename = "SELECT value FROM `TemperatrureChat` "
	case "5": // 事件
		jmax = 6
		tablename = "SELECT value FROM `TemperatrureChat` "
	case "601":
		jmax = 1
		tablename = "SELECT value FROM `IOStatistics` "
		condition = "AND `type` = 1 " //   入量
	case "611":
		jmax = 1
		tablename = "SELECT value FROM `IOStatistics` "
		condition = "AND `type` = 2 AND `subtype` = 1 " // 出量 1：其他入量或其他出量/ml,
	case "612":
		jmax = 1
		tablename = "SELECT value FROM `IOStatistics` "
		condition = "AND `type` = 2 AND `subtype` = 2 " // 出量 2：输液入量或排尿出量/ml,
	case "613":
		jmax = 1
		tablename = "SELECT value FROM `TemperatrureChat` "
		condition = "AND `type` = 2 AND `subtype` = 3 " // 出量 3：饮食入量/ml或大便出量/次',
	case "7": // 血压
		jmax = 1
		tablename = "SELECT diavalue, sysvalue FROM `TemperatrureChat` "
	case "8": // 体重
		jmax = 1
		tablename = "SELECT value FROM `TemperatrureChat` "
	case "9": // 皮试
		jmax = 1
		tablename = "SELECT value FROM `TemperatrureChat` "
	default:
		tablename = ""
	}

	sqlstr := tablename + "WHERE `PatientId` = ? AND `DateTime` >= ? AND `DateTime` < ? " + condition + "ORDER BY `DateTime` DESC LIMIT 1"

	//if "5" == ty {
	//	sqlstr = "SELECT * FROM " + tablename + "WHERE `type` = 1 AND `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1"
	//} else if "6" == ty {
	//	sqlstr = "SELECT * FROM " + tablename + "WHERE `type` = 2 AND `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1"
	//}

	var offset int
	if 6 == jmax {
		offset = 2
	} else {
		offset = 0
	}
	for i := 0; i < 7; i++ {
		day := weeks[i]
		for j := 0; j < jmax; j++ {
			time1 := time.Date(day.Year(), day.Month(), day.Day(), (24/jmax)*j+offset, 0, 0, 0, loc)
			time2 := time.Date(day.Year(), day.Month(), day.Day(), (24/jmax)*(j+1)+offset, 0, 0, 0, loc)
			if "6" == ty {
				time1 = time1.Add(time.Duration(60 * 60 * 8))
				time2 = time2.Add(time.Duration(60 * 60 * 8))
			}
			resultmap, err := fit.MySqlEngine().QueryString(sqlstr, pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil, err
			}
			if len(resultmap) > 0 {
				switch ty {
				case "101", "102", "103", "2", "3", "4", "5", "6", "8", "9":
					dict := resultmap[0]
					results = append(results, dict["value"])
				case "7":
					dict := resultmap[0]
					results = append(results, dict["diavalue"])
					results = append(results, dict["sysvalue"])
				default:
					results = append(results, resultmap[0]["testtime"])
				}

			} else {
				if ty == "7" {
					results = append(results, "")
					results = append(results, "")
				} else {
					results = append(results, "")
				}

			}
		}
	}
	//fit.MySqlEngine().ShowSQL(false)
	return results, nil
}

func GetOperationDate(pid string) (string, error) {
	resultmap, err := fit.SQLServerEngine().QueryString("SELECT VAT08 FROM VAT1 WHERE VAA01 = ?", pid)
	if err != nil {
		fit.Logger().LogError("temp chart", err)
		return "", err
	}
	if resultmap != nil {
		datestr := resultmap[0]["VAT08"]
		return datestr, nil
	} else {
		return "", nil
	}
}

/*查询住院期间的手术记录 */
func FetchOperationRecordsDate(pid int64) ([]string, error) {
	records := make([]string, 0)

	var err error
	// VAT04 = 4 表示已结束手术
	//err = fit.SQLServerEngine().SQL("select VAT08 from VAT1 where VAA01 = ? and VAT04 = 4", pid).Find(&records) .Desc("VAT08")
	resultsmap, err := fit.SQLServerEngine().QueryString("select top 2 VAT08 from VAT1 where VAA01 = ? and VAT04 = 4 ORDER BY VAT08 DESC", pid)
	//err = fit.SQLServerEngine().Table("VAT1").Select("VAT08").Where("VAA01 = ? and VAT04 = 4", pid).Limit(2, 1).Find(&records)  mssql  limit无效

	for _, val := range resultsmap {
		records= append(records, val["VAT08"])
	}
	if len(records) == 2 {
		records[1], records[0] = records[0], records[1]
	}
	return records, err
}

/*
switch ty {
	case "1":
		jmax = 6
		tablename = "`Temperature` "
	case "2":
		jmax = 6
		tablename = "`Breathe` "
	case "3":
		jmax = 6
		tablename = "`Pulse` "
	case "4":
		jmax = 6
		tablename = "`Heartrate` "
	case "5":
		jmax = 6
		tablename = "`Incident` "
	case "6":
		jmax = 1
		tablename = "`IntakeOutput` "
	case "7":
		jmax = 1
		tablename = "`Pressure` "
	case "8":
		jmax = 1
		tablename = "`Weight` "
	case "9":
		jmax = 1
		tablename = "`Skin` "
	default:
		tablename = ""
	}

// 体温
func (m Temperature) GetTemp(pid string, weeks []time.Time)  ([]map[string]string, error) {
	var results []map[string]string
	for i := 0; i < 7 ;i++  {
		day := weeks[i]
		//daystr := day.Format("2006-01-02")
		//time1 := day.Add(time.Duration(60 * 4 * i))
		//var time1, time2 time.Time
		for j := 0;j < 6 ;j++  {
			time1 := time.Date(day.Year(), day.Month(), day.Day(), 4 * j, 0, 0, 0, loc)
			time2 := time.Date(day.Year(), day.Month(), day.Day(), 4 * j + 4, 0, 0, 0, loc)
			//var resultmap []map[string]string
			//var err error
			fmt.Println("get temp", i, time1, time2)
			resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Temperature` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil,err
			}
			//fmt.Println(resultmap)
			if len(resultmap) > 0 {
				results = append(results, resultmap[0])
			} else {
				results = append(results, nil)
			}
		}
	}
	return results, nil
}

// 呼吸
func GetBreathe(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		for j := 0;j < 6 ;j++  {
			time1 = time.Date(day.Year(), day.Month(), day.Day(), 4 * j, 0, 0, 0, loc)
			time2 = time.Date(day.Year(), day.Month(), day.Day(), 4 * j + 4, 0, 0, 0, loc)
			resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Breathe` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil,err
			}
			if len(resultmap) > 0 {
				results = append(results, resultmap[0])
			} else {
				results = append(results, nil)
			}
		}
	}
	return results, nil
}

// 脉搏
func GetPulse(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		for j := 0;j < 6 ;j++  {
			time1 = time.Date(day.Year(), day.Month(), day.Day(), 4 * j, 0, 0, 0, loc)
			time2 = time.Date(day.Year(), day.Month(), day.Day(), 4 * j + 4, 0, 0, 0, loc)
			resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Pulse` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil,err
			}
			if len(resultmap) > 0 {
				results = append(results, resultmap[0])
			} else {
				results = append(results, nil)
			}
		}
	}
	return results, nil
}

// 心率
func GetHeartrate(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		for j := 0;j < 6 ;j++  {
			time1 = time.Date(day.Year(), day.Month(), day.Day(), 4 * j, 0, 0, 0, loc)
			time2 = time.Date(day.Year(), day.Month(), day.Day(), 4 * j + 4, 0, 0, 0, loc)
			resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Heartrate` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil,err
			}
			if len(resultmap) > 0 {
				results = append(results, resultmap[0])
			} else {
				results = append(results, nil)
			}
		}
	}
	return results, nil
}

//输入液量
func GetIntake(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		time1 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		time2 = time.Date(day.Year(), day.Month(), day.Day() + 1, 0, 0, 0, 0, loc)
		resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `IntakeOutput` WHERE `type` = 1 AND `patientid` = ? AND `datetime` >= ? AND `datetime` < ? ORDER BY `datetime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
		if err != nil {
			fit.Logger().LogError("temp chart", err)
			return nil,err
		}
		if len(resultmap) > 0 {
			results = append(results, resultmap[0])
		} else {
			results = append(results, nil)
		}
	}
	return results, nil
}

// 排出量 其他
func GetOutput1(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		time1 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		time2 = time.Date(day.Year(), day.Month(), day.Day() + 1, 0, 0, 0, 0, loc)
		resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `IntakeOutput` WHERE `type` = 2 AND `subtype` = 1 AND `patientid` = ? AND `datetime` >= ? AND `datetime` < ? ORDER BY `datetime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
		if err != nil {
			fit.Logger().LogError("temp chart", err)
			return nil,err
		}
		if len(resultmap) > 0 {
			results = append(results, resultmap[0])
		} else {
			results = append(results, nil)
		}
	}
	return results, nil
}
// 排出量 尿量
func GetOutput2(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		time1 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		time2 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `IntakeOutput` WHERE `type` = 2 AND `subtype` = 1 AND `patientid` = ? AND `datetime` >= ? AND `datetime` < ? ORDER BY `datetime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
		if err != nil {
			fit.Logger().LogError("temp chart", err)
			return nil,err
		}
		if len(resultmap) > 0 {
			results = append(results, resultmap[0])
		} else {
			results = append(results, nil)
		}
	}
	return results, nil
}
// 排出量 大便
func GetOutput3(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		time1 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		time2 = time.Date(day.Year(), day.Month(), day.Day() + 1, 0, 0, 0, 0, loc)
		resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `IntakeOutput` WHERE `type` = 2 AND `subtype` = 1 AND `patientid` = ? AND `datetime` >= ? AND `datetime` < ? ORDER BY `datetime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
		if err != nil {
			fit.Logger().LogError("temp chart", err)
			return nil,err
		}
		if len(resultmap) > 0 {
			results = append(results, resultmap[0])
		} else {
			results = append(results, nil)
		}
	}
	return results, nil
}

// 血压
func GetPressure(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		for j := 0;j < 2 ;j++  {
			time1 = time.Date(day.Year(), day.Month(), day.Day(), 12 * j, 0, 0, 0, loc)
			time2 = time.Date(day.Year(), day.Month(), day.Day(), 12 * j + 12, 0, 0, 0, loc)
			resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Pressure` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
			if err != nil {
				fit.Logger().LogError("temp chart", err)
				return nil,err
			}
			if len(resultmap) > 0 {
				results = append(results, resultmap[0])
			} else {
				results = append(results, nil)
			}
		}
	}
	return results, nil
}

// 体重
func GetWeight(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		time1 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		time2 = time.Date(day.Year(), day.Month(), day.Day() + 1, 0, 0, 0, 0, loc)
		resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Weight` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
		if err != nil {
			fit.Logger().LogError("temp chart", err)
			return nil,err
		}
		if len(resultmap) > 0 {
			results = append(results, resultmap[0])
		} else {
			results = append(results, nil)
		}
	}
	return results, nil
}

// 皮试
func GetSkin(pid string, weeks []time.Time) ([]map[string]string, error)  {
	var results []map[string]string
	var time1, time2 time.Time
	var day time.Time
	for i := 0; i < 7 ;i++  {
		day = weeks[i]
		time1 = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		time2 = time.Date(day.Year(), day.Month(), day.Day() + 1, 0, 0, 0, 0, loc)
		resultmap, err := fit.MySqlEngine().QueryString("SELECT * FROM `Skin` WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1", pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
		if err != nil {
			fit.Logger().LogError("temp chart", err)
			return nil,err
		}
		if len(resultmap) > 0 {
			results = append(results, resultmap[0])
		} else {
			results = append(results, nil)
		}
	}
	return results, nil
}
*/

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
