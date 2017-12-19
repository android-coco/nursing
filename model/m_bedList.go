//  Created by JP

package model

import (
	"fit"
	"time"
	"fmt"
)

/*PDA端专用 科床位室病人*/
type Beds struct {
	VAA1                                `xorm:"extends"`
	NewOrder     int                    `json:"new_order"`      // 新医嘱 ？
	StoppedOrder int                    `json:"stopped_order"`  // 已停医嘱 ？
	Fever        int                    `json:"fever"`          // 是否发热（最后一次测量体温>37.5°）
	Operation    int                    `json:"operation"`      // 是否待手术 ？
	Arrearage    int                    `json:"arrearage"`      // 是否欠费
	NewPatient   int                    `json:"new_patient"`    // 是否是新病人（VAA1.VAA73入院时间判断 今天早上8点到明天早上8点）
	VAK20        int                    `json:"-"`              // 预交金额
	VAK21        int                    `json:"-"`              // 费用总额
	AAG01        string                 `json:"nursing_degree"` // 护理级别
	Vid          int64                  `json:"vid"`            // 就诊ID
	VAE11        DatetimeWithoutSeconds `json:"hospital_date"`  // 入院时间
	Gender       string                 `json:"gender"`         // 性别(结果字符串)
}


/*PC端通用 科室床位病人*/
type PCBedDup struct {
	VAA1Dup                             `xorm:"extends"`       // 病人资料
	BCK03C       string                                        // 病区名
	VAO2                                `xorm:"extends"`       // 诊断病症
	AAG01        string                                        // 护理级别
	Vid          int64                  `json:"vid"`           // 就诊ID
	VAE11        DatetimeWithoutSeconds `json:"hospital_date"` // 入院时间
	Gender       string                                        // 性别(结果字符串)
	HospitalDate string                                        // 入院时间字符串
}

/*病人资料表*/
type VAA1 struct {
	VAA01  int64  `json:"patient_id"`    // 病人ID
	VAA04  string `json:"hosp_num"`      // 住院号
	VAA05  string `json:"name"`          // 姓名
	ABW01  string `json:"sex"`           // 性别 1,男  2,女
	VAA10  int    `json:"age"`           // 年龄
	BCK01C int    `json:"department_id"` // 病区ID
	BCQ04  string `json:"bed_coding"`    // 床号 BCQ1表
	ABQ02  string `json:"nation"`        // 民族
}

/*病床编制表, 查科室所有床位*/
type BCQ1 struct {
	//ROWNR 次序
	VAA01  int64  // 病人ID
	BCK01A int    // 病区ID
	BCQ04  string // 床位号
	BCK03  string // 病区名
	//BCK01B int    // 科室ID
}

/*病人护理记录, 查护理级别*/
type VCF1 struct {
	AAG01 string `json:"nursing_degree"` // 护理级别
}

/*住院病人，PC端专用*/
type InpatientWard struct {
	Pid   int64  // 病人ID  VAA01
	Vid   int64  // 就诊ID  VAE01
	PName string // 姓名    VAE95
	Bed   string // 床号    BCQ04B
	Did   int    // 病区ID  BCK01C
	//	VAE96 性别   VAE46 年龄   VAE94 住院号
}

/*PC床位列表*/
type PCPatient struct {
	Pid            int64                  `json:"pid"`              // 病人ID  VAA01
	Vid            int64                  `json:"vid"`              // 就诊ID  VAE01
	PName          string                 `json:"p_name"`           // 姓名    VAE95
	Bed            string                 `json:"bed"`              // 床号    BCQ04B
	Did            int                    `json:"did"`              // 病区ID  BCK01C
	DName          string                 `json:"d_name"`           // 病区		BCK03
	HospNum        string                 `json:"hosp_num"`         // 住院号	VAE94
	Age            string                 `json:"age"`              // 年龄	 	VAE46
	Category       string                 `json:"category"`         // 病人类别	BDP02
	Gender         string                 `json:"gender"`           // 性别     VAE96
	Nurse          string                 `json:"nurse"`            // 责任护士ID  BCE03B
	Physician      string                 `json:"physician"`        // 住院医师    BCE03C
	NursingDegree  string                 `json:"nursing_degree"`   // 护理级别 BCF01
	NursingDegreeV string                 `json:"nursing_degree_v"` // 护理级别
	AAG01          int                    `json:"-"`                // AAG01 --> BBY01
	Diagnosis      string                 `json:"diagnosis"`        // 诊断
	HospTime       DatetimeWithoutSeconds `json:"-"`                // 入院时间
	HospDay        string                 `json:"hosp_time"`        // 入院时间
	NewOrder       int                    `json:"new_order"`        // 新医嘱 ？
	StoppedOrder   int                    `json:"stopped_order"`    // 已停医嘱 ？
	Fever          int                    `json:"fever"`            // 是否发热（最后一次测量体温>37.5°）
	Operation      int                    `json:"operation"`        // 是否待手术 ？
	Arrearage      int                    `json:"arrearage"`        // 是否欠费
	NewPatient     int                    `json:"new_patient"`      // 是否是新病人（VAA1.VAA73入院时间判断 今天早上8点到明天早上8点）
}

/*PDA 接口，返回 []Beds 切片数组*/
func QueryDepartmentBedList(BCK01 int) ([]Beds, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select VAA01, BCQ04, BCK01A from BCQ1 where BCK01A = ? and VAA01 != 0 order by ROWNR", BCK01).Find(&bedArray)
	if err == nil {
		length := len(bedArray)
		slice := make([]Beds, length)
		for i, bed := range bedArray {
			patient := Beds{}
			patient.VAA01 = bed.VAA01
			patient.BCK01C = bed.BCK01A
			patient.BCQ04 = bed.BCQ04

			// 病人资料
			_, err = fit.MySqlEngine().SQL("select ABQ02 from VAA1 where VAA01 = ? order by VAA73 desc", bed.VAA01).Get(&patient)
			// 1.判断是否发热
			_, err = fit.MySqlEngine().SQL("select (value > 37.5) as Fever from (SELECT HeadType,PatientId,TestTime,value from NurseChat UNION ALL SELECT HeadType,PatientId,TestTime,value from TemperatrureChat) alias WHERE HeadType = 1 and PatientId = ? and value != '' order by testtime desc limit 1", bed.VAA01).Get(&patient)

			// 4.查询护理级别
			//_, err = fit.SQLServerEngine().SQL("select BBY1.BCF01 as AAG01,VAE1.VAE11,VAE1.VAE94 as VAA04,VAE1.VAE95 as VAA05,VAE1.VAE96 as ABW01,Case VAE1.VAE96 when 1 then '男' when 2 then '女' else '未知' end as Gender,VAE1.VAE46 as VAA10 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", bed.VAA01).Get(&patient)
			_, err = fit.SQLServerEngine().SQL("SELECT TOP 1 a.VAE95 VAA05, a.VAE01 Vid, a.VAE94 VAA04, a.VAE46 VAA10, a.VAE96 ABW01, CASE a.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, a.VAE11, b.BCF01 AAG01 FROM VAE1 a LEFT JOIN BBY1 b ON b.BBY01 = a.AAG01 WHERE a.VAA01 = ? AND a.VAE44 = 2 ORDER BY a.VAE11 DESC", bed.VAA01).Get(&patient)
			hospitalDate := patient.VAE11.ParseToSecond()

			// 待手术
			fit.SQLServerEngine().SQL("select count(1) as Operation from VAT1 where VAT08 > ? and VAT04 = 2 and VAA01 = ?", hospitalDate, bed.VAA01).Get(&patient)
			if patient.Operation > 0 {
				patient.Operation = 1
			}

			// 3.判断是否为新病人
			timeNow := time.Now()
			hospital_date := time.Time(patient.VAE11)
			// 当前系统时间大于8点（比较今天8点）
			if timeNow.Hour() > 8 {
				reference := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 8, 0, 0, 0, time.Local)
				if wh := hospital_date.After(reference); wh == true {
					patient.NewPatient = 1
				}
				// 当前系统时间小于8点（比较昨天8点到今天8点）
			} else {
				// 参考时间：前一天早上8点
				temp := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 8, 0, 0, 0, time.Local)
				reference := temp.AddDate(0, 0, -1)
				// 入院时间大于前一天早上8点
				if wh := reference.Before(hospital_date); wh == true {
					patient.NewPatient = 1
				}
			}

			// 5. 是否有新医嘱
			patient.NewOrder = IsExistNewMedicalAdvice(patient.Vid, BCK01, hospitalDate)

			// 6. 是否有已停医嘱
			patient.StoppedOrder = IsExistFinishedMedicalAdvice(patient.Vid, BCK01)

			//slice = append(slice, patient)
			slice[i] = patient
		}
		return slice, err
	}
	slice := make([]Beds, 0)
	return slice, err
}

/*返回病人列表，包括病人ID，床位，姓名，科室ID*/
func FetchInpatientWardPatients(BCK01 int) []InpatientWard {
	bedArray := make([]InpatientWard, 0)
	fit.SQLServerEngine().SQL("select a.VAA01 Pid , a.VAE01 Vid , a.VAE95 PName , a.BCQ04B Bed , a.BCK01C Did from BCQ1 b , VAE1 a where b.BCK01A = ? and b.VAA01 != 0 and a.VAE44 = 2 and a.VAA01 = b.VAA01 order by b.BCQ04", BCK01).Find(&bedArray)
	return bedArray
}

/*
PC端通用func，返回 []PCBedDup 切片数组(不包括空床位)

BCK01: 科室ID
showEmpty(废弃)： 1：包括空床位，0：不包括空床位
*/
func QueryDepartmentBeds(BCK01 int, showEmpty bool) ([]PCBedDup, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select a.VAA01, a.BCQ04, a.BCK01A, b.BCK03 from BCQ1 a JOIN BCK1 b ON a.BCK01A = b.BCK01 where a.BCK01A = ? and a.VAA01 != 0 order by a.BCQ04", BCK01).Find(&bedArray)

	if err == nil {
		length := len(bedArray)
		slice := make([]PCBedDup, length)
		for i, bed := range bedArray {
			patient := PCBedDup{}
			patient.VAA01 = bed.VAA01
			patient.BCK01C = bed.BCK01A
			patient.BCK03C = bed.BCK03
			patient.BCQ04 = bed.BCQ04

			// 1.民族
			_, err = fit.MySqlEngine().SQL("select ABQ02 from VAA1 where VAA01 = ? order by VAA73 desc", bed.VAA01).Get(&patient)

			// 各种基本信息
			//_, err = fit.SQLServerEngine().SQL("select BBY1.BCF01 as AAG01, VAE1.BCE03B, VAE1.BCE03C, VAE1.VAE11, VAE1.VAE94 as VAA04, VAE1.VAE95 as VAA05, VAE1.VAE96 as ABW01, Case VAE1.VAE96 when 1 then '男' when 2 then '女' else '未知' end as Gender, VAE1.VAE46 as VAA10 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 = 2 and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", bed.VAA01).Get(&patient)
			_, err = fit.SQLServerEngine().SQL("SELECT TOP 1 a.VAE95 VAA05, a.VAE01 Vid, a.VAE94 VAA04, a.VAE46 VAA10, a.VAE96 ABW01, a.BDP02, a.VAE44 VAA61, CASE a.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, a.VAE11, b.BCF01 AAG01 FROM VAE1 a LEFT JOIN BBY1 b ON b.BBY01 = a.AAG01 WHERE a.VAA01 = ? AND a.VAE44 = 2 ORDER BY a.VAE11 DESC", bed.VAA01).Get(&patient)
			patient.HospitalDate = patient.VAE11.ParseToMinute()
			hospitalDate := patient.VAE11.ParseToSecond()

			// 4. 查询诊断症状
			_, err = fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT top 1 VAO2.VAO15 FROM VAO2 WHERE VAO2.ACF01 = 2 AND VAO2.VAA01 = %d AND VAO2.VAO19 > '%s' ORDER BY VAO2.VAO19 DESC", bed.VAA01, hospitalDate)).Get(&patient)
			slice[i] = patient
		}
		return slice, err
	}
	slice := make([]PCBedDup, 0)
	return slice, err
}

/*
PC端 接口 主页专用（返回相应类别的病人列表）

BCK01: 科室ID
typeDup: 0：所有病人，1：新病人，2：新医嘱，3：停止医嘱，4：发热，5：待手术
*/
func GetDepartmentBedsByClassifying(BCK01 int, typeDup int) (map[string]interface{}, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select a.VAA01, a.BCQ04, a.BCK01A, b.BCK03 from BCQ1 a JOIN BCK1 b ON a.BCK01A = b.BCK01 where a.BCK01A = ? order by a.BCQ04", BCK01).Find(&bedArray)
	slice := make([]PCPatient, 0)
	t0 := 0
	t1 := 0
	t2 := 0
	t3 := 0
	t4 := 0
	t5 := 0

	if length := len(bedArray); length != 0 {
		for _, bed := range bedArray {
			// bed.VAA01 == 0 表示空床位
			patient := PCPatient{
				Pid:   bed.VAA01,
				Did:   bed.BCK01A,
				DName: bed.BCK03,
				Bed:   bed.BCQ04,
			}
			if bed.VAA01 != 0 {
				// 基本信息和护理级别
				_, err = fit.SQLServerEngine().SQL("SELECT TOP 1 a.VAE95 PName, a.VAE01 Vid, a.VAE94 HospNum, a.VAE46 Age, a.BDP02 Category, CASE a.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, a.BCE03B Nurse, a.BCE03C Physician, a.AAG01, a.VAE11 HospTime, CASE WHEN b.BCF01 is null THEN '无护理级别' ELSE b.BBY05 END NursingDegree, b.BCF01 NursingDegreeV FROM VAE1 a LEFT JOIN BBY1 b ON b.BBY01 = a.AAG01 WHERE a.VAA01 = ? AND a.VAE44 = 2 ORDER BY a.VAE11 DESC", bed.VAA01).Get(&patient)
				patient.HospDay = patient.HospTime.ParseToMinute()
				hospitalDate := patient.HospTime.ParseToSecond()
				// 诊断信息
				_, err = fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT top 1 VAO2.VAO15 Diagnosis FROM VAO2 WHERE VAO2.ACF01 = 2 AND VAO2.VAA01 = %d AND VAO2.VAO19 > '%s' ORDER BY VAO2.VAO19 DESC", bed.VAA01, hospitalDate)).Get(&patient)

				// 判断是否发热
				_, err = fit.MySqlEngine().SQL("select (value > 37.5) as Fever from (SELECT HeadType,PatientId,TestTime,value from NurseChat UNION ALL SELECT HeadType,PatientId,TestTime,value from TemperatrureChat) alias WHERE HeadType = 1 and PatientId = ? and value != '' order by testtime desc limit 1", bed.VAA01).Get(&patient)

				// 判断是否为新病人
				timeNow := time.Now()
				hospital_date := time.Time(patient.HospTime)
				// 当前系统时间大于8点（比较今天8点）
				if timeNow.Hour() > 8 {
					reference := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 8, 0, 0, 0, time.Local)
					if wh := hospital_date.After(reference); wh == true {
						patient.NewPatient = 1
					}
					// 当前系统时间小于8点（比较昨天8点到今天8点）
				} else {
					// 参考时间：前一天早上8点
					temp := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 8, 0, 0, 0, time.Local)
					reference := temp.AddDate(0, 0, -1)
					// 入院时间大于前一天早上8点
					if wh := reference.Before(hospital_date); wh == true {
						patient.NewPatient = 1
					}
				}

				// 是否有新医嘱
				patient.NewOrder = IsExistNewMedicalAdvice(patient.Vid, BCK01, hospitalDate)

				// 是否有已停医嘱
				patient.StoppedOrder = IsExistFinishedMedicalAdvice(patient.Vid, BCK01)

				// 待手术
				fit.SQLServerEngine().SQL("select count(1) as Operation from VAT1 where VAT08 > ? and VAT04 = 2 and VAA01 = ?", hospitalDate, bed.VAA01).Get(&patient)
				if patient.Operation > 0 {
					patient.Operation = 1
				}
				// 归类
				if patient.NewPatient == 1 {
					t1 += 1
					if typeDup == 1 {
						slice = append(slice, patient)
					}
				}
				if patient.NewOrder == 1 {
					t2 += 1
					if typeDup == 2 {
						slice = append(slice, patient)
					}
				}
				if patient.StoppedOrder == 1 {
					t3 += 1
					if typeDup == 3 {
						slice = append(slice, patient)
					}
				}
				if patient.Fever == 1 {
					t4 += 1
					if typeDup == 4 {
						slice = append(slice, patient)
					}
				}
				if patient.Operation == 1 {
					t5 += 1
					if typeDup == 5 {
						slice = append(slice, patient)
					}
				}
				if typeDup == 0 {
					slice = append(slice, patient)
				}
				t0 += 1

			} else if typeDup == 0 {
				slice = append(slice, patient)
			}
		}
	}
	obj := make(map[string]interface{})
	num := make(map[string]int)
	num["t0"] = t0
	num["t1"] = t1
	num["t2"] = t2
	num["t3"] = t3
	num["t4"] = t4
	num["t5"] = t5
	obj["num"] = num
	obj["bed"] = slice
	return obj, err
}

/*
PC端 接口 主页专用（返回相应类别的病人列表）

BCK01: 科室ID
typeDup: 0：所有病人，1：新病人，2：新医嘱，3：停止医嘱，4：发热，5：待手术
page 页数
pagenum 一页多少个
showEmpty 是否显示空床位
*/
func GetDepartmentBedsByClassifyingByPage(BCK01 int, typeDup int, page, pagenum int, showEmpty bool) (map[string]interface{}, error) {
	bedArray := make([]BCQ1, 0)
	//err := fit.SQLServerEngine().SQL("select VAA01, BCQ04, BCK01A from BCQ1 where BCK01A = ? order by ROWNR", BCK01).Find(&bedArray)
	//t := (page - 1) * pagenum
	var err error
	if showEmpty {
		err = fit.SQLServerEngine().SQL("select CASE  WHEN a.VAA01 > 0 THEN a.VAA01 ELSE 0 END VAA01, a.BCQ04, a.BCK01A, b.BCK03 from BCQ1 a JOIN BCK1 b ON a.BCK01A = b.BCK01 where a.BCK01A = ? order by a.BCQ04", BCK01).Find(&bedArray)
	} else {
		err = fit.SQLServerEngine().SQL("select a.VAA01, a.BCQ04, a.BCK01A, b.BCK03 from BCQ1 a JOIN BCK1 b ON a.BCK01A = b.BCK01 where a.BCK01A = ? AND a.VAA01 > 0 order by a.BCQ04", BCK01).Find(&bedArray)
	}

	slice := make([]PCPatient, 0)
	t0 := 0
	t1 := 0
	t2 := 0
	t3 := 0
	t4 := 0
	t5 := 0

	if length := len(bedArray); length != 0 {
		for _, bed := range bedArray {
			// bed.VAA01 == 0 表示空床位
			patient := PCPatient{
				Pid:   bed.VAA01,
				Did:   bed.BCK01A,
				DName: bed.BCK03,
				Bed:   bed.BCQ04,
			}
			if bed.VAA01 != 0 {
				// 基本信息和护理级别
				_, err = fit.SQLServerEngine().SQL("SELECT TOP 1 a.VAE95 PName, a.VAE01 Vid, a.VAE94 HospNum, a.VAE46 Age, a.BDP02 Category, CASE a.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, a.BCE03B Nurse, a.BCE03C Physician, a.AAG01, a.VAE11 HospTime, CASE WHEN b.BCF01 is null THEN '无护理级别' ELSE b.BBY05 END NursingDegree, b.BCF01 NursingDegreeV FROM VAE1 a LEFT JOIN BBY1 b ON b.BBY01 = a.AAG01 WHERE a.VAA01 = ? AND a.VAE44 = 2 ORDER BY a.VAE11 DESC", bed.VAA01).Get(&patient)
				patient.HospDay = patient.HospTime.ParseToMinute()
				hospitalDate := patient.HospTime.ParseToSecond()
				// 诊断信息
				_, err = fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT top 1 VAO2.VAO15 Diagnosis FROM VAO2 WHERE VAO2.ACF01 = 2 AND VAO2.VAA01 = %d AND VAO2.VAO19 > '%s' ORDER BY VAO2.VAO19 DESC", bed.VAA01, hospitalDate)).Get(&patient)

				// 判断是否发热
				_, err = fit.MySqlEngine().SQL("select (value > 37.5) as Fever from (SELECT HeadType,PatientId,TestTime,value from NurseChat UNION ALL SELECT HeadType,PatientId,TestTime,value from TemperatrureChat) alias WHERE HeadType = 1 and PatientId = ? and value != '' order by testtime desc limit 1", bed.VAA01).Get(&patient)

				// 判断是否为新病人
				timeNow := time.Now()
				hospital_date := time.Time(patient.HospTime)
				// 当前系统时间大于8点（比较今天8点）
				if timeNow.Hour() > 8 {
					reference := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 8, 0, 0, 0, time.Local)
					if wh := hospital_date.After(reference); wh == true {
						patient.NewPatient = 1
					}
					// 当前系统时间小于8点（比较昨天8点到今天8点）
				} else {
					// 参考时间：前一天早上8点
					temp := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 8, 0, 0, 0, time.Local)
					reference := temp.AddDate(0, 0, -1)
					// 入院时间大于前一天早上8点
					if wh := reference.Before(hospital_date); wh == true {
						patient.NewPatient = 1
					}
				}

				// 是否有新医嘱
				patient.NewOrder = IsExistNewMedicalAdvice(patient.Vid, BCK01, hospitalDate)

				// 是否有已停医嘱
				patient.StoppedOrder = IsExistFinishedMedicalAdvice(patient.Vid, BCK01)

				// 待手术
				fit.SQLServerEngine().SQL("select count(1) as Operation from VAT1 where VAT08 > ? and VAT04 = 2 and VAA01 = ?", hospitalDate, bed.VAA01).Get(&patient)
				if patient.Operation > 0 {
					patient.Operation = 1
				}
				// 归类
				if patient.NewPatient == 1 {
					t1 += 1
					if typeDup == 1 {
						slice = append(slice, patient)
					}
				}
				if patient.NewOrder == 1 {
					t2 += 1
					if typeDup == 2 {
						slice = append(slice, patient)
					}
				}
				if patient.StoppedOrder == 1 {
					t3 += 1
					if typeDup == 3 {
						slice = append(slice, patient)
					}
				}
				if patient.Fever == 1 {
					t4 += 1
					if typeDup == 4 {
						slice = append(slice, patient)
					}
				}
				if patient.Operation == 1 {
					t5 += 1
					if typeDup == 5 {
						slice = append(slice, patient)
					}
				}
				if typeDup == 0 {
					slice = append(slice, patient)
				}
				t0 += 1

			} else if typeDup == 0 && showEmpty {
				slice = append(slice, patient)
			}
		}
	}
	obj := make(map[string]interface{})
	num := make(map[string]int)
	var totalPage int
	if len(bedArray)%pagenum == 0 {
		totalPage = len(bedArray) / pagenum
	} else {
		totalPage = len(bedArray)/pagenum + 1
	}
	num["t0"] = t0
	num["t1"] = t1
	num["t2"] = t2
	num["t3"] = t3
	num["t4"] = t4
	num["t5"] = t5
	//obj["num"] = num
	var start, end int
	if page > totalPage {
		page = totalPage
	}
	start = (page - 1) * pagenum
	end = page * pagenum
	if start >= len(slice) {
		start = len(slice) - pagenum
	}
	if end >= len(slice) {
		end = len(slice)
	}
	obj["bed"] = slice[start:end]
	obj["totalpage"] = totalPage
	return obj, err
}
