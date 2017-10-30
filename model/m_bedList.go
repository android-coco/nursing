package model

import (
	"fit"
	"time"
)

/*PDA端专用 科床位室病人*/
type Beds struct {
	VAA1                `xorm:"extends"`
	NewOrder     int    `json:"new_order"`      // 新医嘱 ？
	StoppedOrder int    `json:"stopped_order"`  // 已停医嘱 ？
	Fever        int    `json:"fever"`          // 是否发热（最后一次测量体温>37.5°）
	Operation    int    `json:"operation"`      // 是否待手术 ？
	Arrearage    int    `json:"arrearage"`      // 是否欠费
	NewPatient   int    `json:"new_patient"`    // 是否是新病人（VAA1.VAA73入院时间判断 今天早上8点到明天早上8点）
	VAK20        int    `json:"-"`              // 预交金额
	VAK21        int    `json:"-"`              // 费用总额
	AAG01        string `json:"nursing_degree"` // 护理级别
	Gender       string `json:"gender"`         // 性别(结果字符串)
}

/*PC端主页专用 科室床位病人*/
type PCBeds struct {
	VAA1Dup             `xorm:"extends"`         // 病人资料
	BCK03        string `json:"department_name"` // 科室名称
	VAE1                `xorm:"extends"`         // 主治医师、责任护士
	VAO2                `xorm:"extends"`         // 诊断病症
	NewOrder     int    `json:"new_order"`       // 新医嘱 ？
	StoppedOrder int    `json:"stopped_order"`   // 已停医嘱 ？
	Fever        int    `json:"fever"`           // 是否发热（最后一次测量体温>37.5°）
	Operation    int    `json:"operation"`       // 是否待手术 ？
	Arrearage    int    `json:"arrearage"`       // 是否欠费
	NewPatient   int    `json:"new_patient"`     // 是否是新病人（VAA1.VAA73入院时间判断 今天早上8点到明天早上8点）
	VAK20        int    `json:"prepay"`          // 预交金额
	VAK21        int    `json:"aggregate_costs"` // 费用总额
	AAG01        string `json:"nursing_degree"`  // 护理级别
	Gender       string `json:"gender"`          // 性别(结果字符串)
}

/*PC端通用 科室床位病人*/
type PCBedDup struct {
	VAA1 `xorm:"extends"` // 病人资料
	BCK03        string   // 科室名称
	BCK01C       int      // 病区ID
	BCK03C       string   // 病区名
	VAO2 `xorm:"extends"` // 诊断病症
	AAG01        string   // 护理级别
	Gender       string   // 性别(结果字符串)
	HospitalDate string   // 入院时间字符串
}

/*病人资料表*/
type VAA1 struct {
	VAA01  int          `json:"patient_id"`    // 病人ID
	VAA04  string       `json:"hosp_num"`      // 住院号
	VAA05  string       `json:"name"`          // 姓名
	ABW01  string       `json:"sex"`           // 性别 1,男  2,女
	VAA10  int          `json:"age"`           // 年龄
	BCK01B int          `json:"department_id"` // 科室ID  BCK1表
	VAA73  Datetime_IOV `json:"hospital_date"` // 入院时间
	BCQ04  string       `json:"bed_coding"`    // 床号 BCQ1表
}

/*病床编制表, 查科室所有床位*/
type BCQ1 struct {
	//ROWNR 次序
	VAA01  int    // 病人ID
	BCK01B int    // 科室ID
	BCQ04  string // 床位号
}

/*病人护理记录, 查护理级别*/
type VCF1 struct {
	AAG01 string `json:"nursing_degree"` // 护理级别
}

/*PDA端 接口，返回 []Beds 切片数组*/
func QueryDepartmentBedList(BCK01 int) ([]Beds, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select BCK01B, VAA01, BCQ04 from BCQ1 where BCK01B = ? order by ROWNR", BCK01).Find(&bedArray)
	slice := make([]Beds, 0)
	if err == nil {
		if length := len(bedArray); length != 0 {
			for _, bed := range bedArray {
				// bed.VAA01 == 0 表示空床位
				if bed.VAA01 != 0 {
					patient := Beds{}
					patient.VAA01 = bed.VAA01
					patient.BCK01B = bed.BCK01B
					patient.BCQ04 = bed.BCQ04
					_, err = fit.MySqlEngine().SQL("select VAA04, VAA05, ABW01, VAA10, VAA73 from VAA1 where BCK01B = ? and VAA61 = 2 and VAA01 = ?", BCK01, bed.VAA01).Get(&patient)
					_, err = fit.SQLServerEngine().SQL("select VAK20, VAK21 from VAK1 where VAA01 = ?", bed.VAA01).Get(&patient)

					// 1.判断是否发热
					if bl, _ := GetWhetherFever(bed.VAA01); bl == true {
						patient.Fever = 1
					}

					// 2.判断是否欠费 (10.25暂时不做，查不到准确的费用数据)
					//if patient.VAK20 < patient.VAK21 {
					//	patient.Arrearage = 1
					//}

					// 3.判断是否为新病人
					timeNow := time.Now()
					hospital_date := time.Time(patient.VAA73)
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
					// 4.查询护理级别
					nursingDegree := VCF1{}
					_, _ = fit.SQLServerEngine().SQL("select AAG01 from VCF1 where VAA01 = ?", bed.VAA01).Desc("VCF01").Get(&nursingDegree)
					patient.AAG01 = nursingDegree.AAG01

					// 5. 是否有新医嘱
					advices, _ := GetNonExecutionAdvice(patient.VAA01)
					if length_ad := len(advices); length_ad > 0 {
						patient.NewOrder = 1
					}

					// 6. 是否有已停医嘱
					advices_stoped, _ := GetUncertainOewAdvice(patient.VAA01)
					if length_ad := len(advices_stoped); length_ad > 0 {
						patient.StoppedOrder = 1
					}

					// 7.判断性别
					if patient.ABW01 == "1" || patient.ABW01 == "M" {
						patient.Gender = "男"
					} else if patient.ABW01 == "2" || patient.ABW01 == "F" {
						patient.Gender = "女"
					} else {
						patient.Gender = "未知"
					}

					slice = append(slice, patient)
				}
			}
		}
	}
	return slice, err
}

/*
PC端通用func，返回 []PCBedDup 切片数组

BCK01: 科室ID
showEmpty： 1：包括空床位，0：不包括空床位
*/
func QueryDepartmentBeds(BCK01 int, showEmpty bool) ([]PCBedDup, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select BCK01B, VAA01, BCQ04 from BCQ1 where BCK01B = ? order by ROWNR", BCK01).Find(&bedArray)
	slice := make([]PCBedDup, 0)

	if err == nil {
		if length := len(bedArray); length != 0 {
			for _, bed := range bedArray {
				// bed.VAA01 == 0 表示空床位
				if bed.VAA01 != 0 {
					patient := PCBedDup{}
					patient.VAA01 = bed.VAA01
					patient.BCK01B = bed.BCK01B
					patient.BCQ04 = bed.BCQ04
					// 1.查询病人资料
					_, err = fit.MySqlEngine().SQL("select VAA04, VAA05, ABW01, VAA10, VAA73, BCK01C from VAA1 where BCK01B = ? and VAA61 = 2 and VAA01 = ?", BCK01, bed.VAA01).Get(&patient)
					// 2.查询科室名
					patient.BCK03,_ = QueryDepartmentNameWithId(bed.BCK01B)
					patient.BCK03C, _ = QueryDepartmentNameWithId(patient.BCK01C)
					// 3.查询护理级别
					nursingDegree := VCF1{}
					_, _ = fit.SQLServerEngine().SQL("select AAG01 from VCF1 where VAA01 = ?", bed.VAA01).Desc("VCF01").Get(&nursingDegree)
					patient.AAG01 = nursingDegree.AAG01
					// 4. 查询诊断症状
					illness := make([]VAO2, 0)
					err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAA01 = ? order by VAO19 desc", bed.VAA01).Limit(1, 0).Find(&illness)
					if length := len(illness); length > 0 {
						patient.VAO2 = illness[0]
					} else {
						patient.VAO15 = ""
					}

					// 5.判断性别
					if patient.ABW01 == "1" || patient.ABW01 == "M" {
						patient.Gender = "男"
					} else if patient.ABW01 == "2" || patient.ABW01 == "F" {
						patient.Gender = "女"
					} else {
						patient.Gender = "未知"
					}
					patient.HospitalDate = time.Time(patient.VAA73).Format("2006-01-02 15:04")
					slice = append(slice, patient)

				} else if showEmpty == true {
					patient := PCBedDup{}
					patient.VAA01 = 0 // 空床位
					patient.BCK01B = bed.BCK01B
					patient.BCQ04 = bed.BCQ04
					slice = append(slice, patient)
				}
			}
		}
	}
	return slice, err
}

/*
PC端接口主页专用（返回相应类别的病人列表）

BCK01: 科室ID
typeDup: 0：所有病人，1：新病人，2：新医嘱，3：停止医嘱，4：发热，5：待手术
*/
func GetDepartmentBedsByClassifying(BCK01 int, typeDup int) (map[string]interface{}, error) {
	bedArray := make([]BCQ1, 0)
	err := fit.SQLServerEngine().SQL("select BCK01B, VAA01, BCQ04 from BCQ1 where BCK01B = ? order by ROWNR", BCK01).Find(&bedArray)
	slice := make([]PCBeds, 0)
	t0 := 0
	t1 := 0
	t2 := 0
	t3 := 0
	t4 := 0
	t5 := 0
	if err == nil {
		if length := len(bedArray); length != 0 {
			t0 = length
			for _, bed := range bedArray {
				// bed.VAA01 == 0 表示空床位
				if bed.VAA01 != 0 {
					patient := PCBeds{}
					patient.VAA01 = bed.VAA01
					patient.BCK01B = bed.BCK01B
					patient.BCQ04 = bed.BCQ04
					_, err = fit.MySqlEngine().SQL("select VAA04, VAA05, ABW01, VAA10, VAA73, BDP02, ABQ02, VAA61 from VAA1 where BCK01B = ? and VAA61 = 2 and VAA01 = ?", BCK01, bed.VAA01).Get(&patient)

					_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C, BCK1.BCK03  from VAE1, BCK1 where VAE1.VAA01 = ? and BCK1.BCK01 = ?", bed.VAA01, bed.BCK01B).Limit(1, 0).Get(&patient)
					// 1.判断是否发热
					if bl, _ := GetWhetherFever(bed.VAA01); bl == true {
						patient.Fever = 1
					}

					// 2.判断是否欠费 (10.25暂时不做，查不到准确的费用数据)
					//_, err = fit.SQLServerEngine().SQL("select VAK20, VAK21 from VAK1 where VAA01 = ?", bed.VAA01).Get(&patient)
					//if patient.VAK20 < patient.VAK21 {
					//	patient.Arrearage = 1
					//}

					// 3.判断是否为新病人
					timeNow := time.Now()
					hospital_date := time.Time(patient.VAA73)
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
					// 4.查询护理级别
					nursingDegree := VCF1{}
					_, _ = fit.SQLServerEngine().SQL("select AAG01 from VCF1 where VAA01 = ?", bed.VAA01).Desc("VCF01").Get(&nursingDegree)
					patient.AAG01 = nursingDegree.AAG01

					// 5. 是否有新医嘱
					advices, _ := GetNonExecutionAdvice(patient.VAA01)
					if length_ad := len(advices); length_ad > 0 {
						patient.NewOrder = 1
					}

					// 6. 是否有已停医嘱
					advices_stoped, _ := GetUncertainOewAdvice(patient.VAA01)
					if length_ad := len(advices_stoped); length_ad > 0 {
						patient.StoppedOrder = 1
					}

					// 7. 查询诊断症状
					illness := make([]VAO2, 0)
					err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAA01 = ? order by VAO19 desc", bed.VAA01).Limit(1, 0).Find(&illness)
					if length := len(illness); length > 0 {
						patient.VAO2 = illness[0]
					} else {
						patient.VAO15 = ""
					}

					// 8.判断性别
					if patient.ABW01 == "1" || patient.ABW01 == "M" {
						patient.Gender = "男"
					} else if patient.ABW01 == "2" || patient.ABW01 == "F" {
						patient.Gender = "女"
					} else {
						patient.Gender = "未知"
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

				} else if typeDup == 0 {
					patient := PCBeds{}
					patient.VAA01 = 0 // 空床位
					patient.BCK01B = bed.BCK01B
					patient.BCQ04 = bed.BCQ04
					patient.VAA61 = 2
					slice = append(slice, patient)
				}
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
