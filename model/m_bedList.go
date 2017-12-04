//  Created by JP

package model

import (
	"fit"
	"time"
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
	VAE11        DatetimeWithoutSeconds `json:"hospital_date"`  // 入院时间
	Gender       string                 `json:"gender"`         // 性别(结果字符串)
}

/*PC端主页专用 科室床位病人*/
type PCBeds struct {
	VAA1Dup                             `xorm:"extends"`         // 病人资料
	BCK03C       string                 `json:"department_name"` // 科室名称
	BCE03B       string                 `json:"nurse_name"`      // 责任护士ID
	BCE03C       string                 `json:"physician_name"`  // 住院医师
	VAE11        DatetimeWithoutSeconds `json:"-"`               // 入院时间
	VAO2                                `xorm:"extends"`         // 诊断病症
	NewOrder     int                    `json:"new_order"`       // 新医嘱 ？
	StoppedOrder int                    `json:"stopped_order"`   // 已停医嘱 ？
	Fever        int                    `json:"fever"`           // 是否发热（最后一次测量体温>37.5°）
	Operation    int                    `json:"operation"`       // 是否待手术 ？
	Arrearage    int                    `json:"arrearage"`       // 是否欠费
	NewPatient   int                    `json:"new_patient"`     // 是否是新病人（VAA1.VAA73入院时间判断 今天早上8点到明天早上8点）
	VAK20        int                    `json:"prepay"`          // 预交金额
	VAK21        int                    `json:"aggregate_costs"` // 费用总额
	AAG01        string                 `json:"nursing_degree"`  // 护理级别
	Gender       string                 `json:"gender"`          // 性别(结果字符串)
	HospitalDate string                 `json:"hospital_date"`   // 入院时间字符串
}

/*PC端通用 科室床位病人*/
type PCBedDup struct {
	VAA1Dup                             `xorm:"extends"`       // 病人资料
	BCK03C       string                                        // 病区名
	VAO2                                `xorm:"extends"`       // 诊断病症
	AAG01        string                                        // 护理级别
	VAE11        DatetimeWithoutSeconds `json:"hospital_date"` // 入院时间
	Gender       string                                        // 性别(结果字符串)
	HospitalDate string                                        // 入院时间字符串
	//BCK01C       int      // 病区ID
	//BCK03        string   // 科室名称
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
	//BCK01B int          `json:"department_id"` // 科室ID  BCK1表
}

/*病床编制表, 查科室所有床位*/
type BCQ1 struct {
	//ROWNR 次序
	VAA01  int64  // 病人ID
	BCK01A int    // 病区ID
	BCQ04  string // 床位号
	//BCK01B int    // 科室ID
}

/*病人护理记录, 查护理级别*/
type VCF1 struct {
	AAG01 string `json:"nursing_degree"` // 护理级别
}

/*住院病人，PC端专用*/
type InpatientWard struct {
	VAA01  int64  // 病人ID
	BCQ04  string // 床位
	VAA05  string // 姓名
	BCK01C int    // 科室ID
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
			_, err = fit.MySqlEngine().SQL("select VAA04, VAA05, ABW01, ABQ02, VAA10 from VAA1 where VAA01 = ? order by VAA73 desc", bed.VAA01).Get(&patient)
			// 1.判断是否发热
			_, err = fit.MySqlEngine().SQL("select (value > 37.5) as Fever from (SELECT HeadType,PatientId,TestTime,value from NurseChat UNION ALL SELECT HeadType,PatientId,TestTime,value from TemperatrureChat) alias WHERE HeadType = 1 and PatientId = ? and value != '' order by testtime desc limit 1", bed.VAA01).Get(&patient)

			// 4.查询护理级别
			_, err = fit.SQLServerEngine().SQL("select BBY1.BCF01 as AAG01,VAE1.VAE11 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", bed.VAA01).Get(&patient)
			hospitalDate := patient.VAE11.ParseToSecond()
			// 2.判断是否欠费 (10.25暂时不做，查不到准确的费用数据)
			//if patient.VAK20 < patient.VAK21 {
			//	patient.Arrearage = 1
			//}

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
			patient.NewOrder = IsExistNewMedicalAdvice(bed.VAA01, BCK01, hospitalDate)

			// 6. 是否有已停医嘱
			patient.StoppedOrder = IsExistFinishedMedicalAdvice(bed.VAA01, BCK01)

			// 7.判断性别
			if patient.ABW01 == "1" || patient.ABW01 == "M" {
				patient.Gender = "男"
			} else if patient.ABW01 == "2" || patient.ABW01 == "F" {
				patient.Gender = "女"
			} else {
				patient.Gender = "未知"
			}

			//slice = append(slice, patient)
			slice[i] = patient
		}
		return slice, err
	}
	slice := make([]Beds, 0)
	return slice, err
}

/*返回病人列表，包括病人ID，床位，姓名，科室ID*/
func FetchInpatientWard(BCK01 int) []InpatientWard {
	bedArray := make([]InpatientWard, 0)
	fit.SQLServerEngine().SQL("select b.VAA01, b.BCQ04, b.BCK01A as BCK01C, a.VAA05 from BCQ1 b left join VAA1 a on a.VAA01 = b.VAA01 where b.BCK01A = ? and b.VAA01 != 0 order by b.ROWNR", BCK01).Find(&bedArray)
	return bedArray
}

/*
PC端通用func，返回 []PCBedDup 切片数组(不包括空床位)

BCK01: 科室ID
showEmpty(废弃)： 1：包括空床位，0：不包括空床位
*/
func QueryDepartmentBeds(BCK01 int, showEmpty bool) ([]PCBedDup, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select VAA01, BCQ04, BCK01A from BCQ1 where BCK01A = ? and VAA01 != 0 order by ROWNR", BCK01).Find(&bedArray)

	if err == nil {
		length := len(bedArray)
		slice := make([]PCBedDup, length)
		// 获取科室名
		bck := BCK1{}
		_, err = fit.SQLServerEngine().SQL("select BCK03 from BCK1 where BCK01 = ?", BCK01).Get(&bck)
		for i, bed := range bedArray {
			patient := PCBedDup{}
			patient.VAA01 = bed.VAA01
			patient.BCK01C = bed.BCK01A
			patient.BCK03C = bck.BCK03
			patient.BCQ04 = bed.BCQ04
			// 1.查询病人资料
			_, err = fit.MySqlEngine().SQL("select VAA04, VAA05, ABW01, ABQ02, VAA10 from VAA1 where VAA01 = ? order by VAA73 desc", bed.VAA01).Get(&patient)
			// 医生、护士、护理级别
			_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C, VAE1.VAE11, BBY1.BCF01 as AAG01 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", bed.VAA01).Get(&patient)
			patient.HospitalDate = patient.VAE11.ParseToMinute()
			hospitalDate := patient.VAE11.ParseToSecond()

			// 4. 查询诊断症状
			_, err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAO19 > ? and VAO2.ACF01 = 2 and VAO2.VAA01 = ? order by VAO2.VAO19 desc", hospitalDate, bed.VAA01).Get(&patient)
			// 5.判断性别
			if patient.ABW01 == "1" || patient.ABW01 == "M" {
				patient.Gender = "男"
			} else if patient.ABW01 == "2" || patient.ABW01 == "F" {
				patient.Gender = "女"
			} else {
				patient.Gender = "未知"
			}
			if err != nil {
				fit.Logger().LogError("**JP**", err.Error())
			}

			//slice = append(slice, patient)
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
	err := fit.SQLServerEngine().SQL("select VAA01, BCQ04, BCK01A from BCQ1 where BCK01A = ? order by ROWNR", BCK01).Find(&bedArray)
	slice := make([]PCBeds, 0)
	t0 := 0
	t1 := 0
	t2 := 0
	t3 := 0
	t4 := 0
	t5 := 0

	if length := len(bedArray); length != 0 {
		// 获取科室名
		bck := BCK1{}
		_, err = fit.SQLServerEngine().SQL("select BCK03 from BCK1 where BCK01 = ?", BCK01).Get(&bck)
		for _, bed := range bedArray {
			// bed.VAA01 == 0 表示空床位
			if bed.VAA01 != 0 {
				patient := PCBeds{}
				patient.VAA01 = bed.VAA01
				patient.BCK01C = bed.BCK01A
				patient.BCK03C = bck.BCK03
				patient.BCQ04 = bed.BCQ04

				// 病人资料
				_, err = fit.MySqlEngine().SQL("select VAA04, VAA05, ABW01, VAA10, BDP02, ABQ02 from VAA1 where VAA01 = ? order by VAA73 desc", bed.VAA01).Get(&patient)
				// 住院时间
				// 医生、护士、护理级别
				_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C, VAE1.VAE11, BBY1.BCF01 as AAG01 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", bed.VAA01).Get(&patient)
				patient.HospitalDate = patient.VAE11.ParseToMinute()
				hospitalDate := patient.VAE11.ParseToSecond()

				// 1.判断是否发热
				_, err = fit.MySqlEngine().SQL("select (value > 37.5) as Fever from (SELECT HeadType,PatientId,TestTime,value from NurseChat UNION ALL SELECT HeadType,PatientId,TestTime,value from TemperatrureChat) alias WHERE HeadType = 1 and PatientId = ? and value != '' order by testtime desc limit 1", bed.VAA01).Get(&patient)

				// 2.判断是否欠费 (10.25暂时不做，查不到准确的费用数据)
				//_, err = fit.SQLServerEngine().SQL("select VAK20, VAK21 from VAK1 where VAA01 = ?", bed.VAA01).Get(&patient)
				//if patient.VAK20 < patient.VAK21 {
				//	patient.Arrearage = 1
				//}

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

				//// 5. 是否有新医嘱
				patient.NewOrder = IsExistNewMedicalAdvice(bed.VAA01, BCK01, hospitalDate)

				// 6. 是否有已停医嘱
				patient.StoppedOrder = IsExistFinishedMedicalAdvice(bed.VAA01, BCK01)

				// 7. 查询诊断症状
				_, err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAO19 > ? and VAO2.ACF01 = 2 and VAO2.VAA01 = ? order by VAO2.VAO19 desc", hospitalDate, bed.VAA01).Get(&patient)

				// 8.判断性别
				if patient.ABW01 == "1" || patient.ABW01 == "M" {
					patient.Gender = "男"
				} else if patient.ABW01 == "2" || patient.ABW01 == "F" {
					patient.Gender = "女"
				} else {
					patient.Gender = "未知"
				}

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
				patient := PCBeds{}
				patient.VAA01 = 0 // 空床位
				patient.BCK01C = bed.BCK01A
				patient.BCQ04 = bed.BCQ04
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
