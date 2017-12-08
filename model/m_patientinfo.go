//  Created by JP

package model

import (
	"fit"
	"time"
)

/*病人详情*/
type PatientInfo struct {
	VAE1          `xorm:"extends"`
	VAK1          `xorm:"extends"`
	VAA1Dup       `xorm:"extends"`
	VAO2          `xorm:"extends"`
	VCF1          `xorm:"extends"`
	BCK03C string `json:"department_name"` // 病区名称
	Gender string `json:"gender"`          // 性别(结果字符串)
	//BCK03  string `json:"department_name"` // 科室名称
}

/*病人基本信息表*/
type VAA1Dup struct {
	VAA1         `xorm:"extends"`
	BDP02 string `json:"type"`   // type病人类型 BDP1表 自费
	VAA61 int    `json:"status"` // 2: 住院
}

/*病人登录记录*/
type VAE1 struct {
	//VAA01
	BCE03B string                 `json:"nurse_name"`     // 责任护士ID
	BCE03C string                 `json:"physician_name"` // 住院医师
	VAE11  DatetimeWithoutSeconds `json:"hospital_date"`  // 入院时间
}

/*住院病人结账记录*/
type VAK1 struct {
	//VAA01
	VAK20 int `json:"prepay"`          // 预交金额
	VAK21 int `json:"aggregate_costs"` // 费用总额
}

/*住院病人诊断记录*/
type VAO2 struct {
	VAO15 string `json:"diagnose_name"` // 诊断名称
}

/*手术记录*/
type VAT1 struct {
	VAA01 int64     // 病人ID
	VAT04 int       // 手术状态
	VAT08 time.Time // 手术日期
}

/*查询病人详情，不限科室，不限住院、出院*/
func GetPatientInfo(patientId string) ([]PatientInfo, error) {
	responseObj := make([]PatientInfo, 0)
	patient := VAA1Dup{}
	// 查询病人详情
	_, err := fit.MySqlEngine().SQL("select VAA01, BDP02, BCK01C, BCQ04, ABQ02 from VAA1 where VAA01 = ?", patientId).Get(&patient)

	if err == nil && patient.VAA01 != 0 {
		patientInfo := PatientInfo{}
		patientInfo.VAA1Dup = patient
		// 费用：VAK1.VAK20, VAK1.VAK21 (VAK1.VAA01 = ? )
		//_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C from VAE1 where VAE1.VAA01 = ?", patientId).Limit(1, 0).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BBY1.BCF01 as AAG01,VAE1.VAE11,VAE1.VAE94 as VAA04,VAE1.VAE95 as VAA05,VAE1.VAE96 as ABW01,Case VAE1.VAE96 when 1 then '男' when 2 then '女' else '未知' end as Gender,VAE1.VAE46 as VAA10 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", patient.VAA01).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C, VAE1.VAE11, BBY1.BCF01 as AAG01 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", patient.VAA01).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BCK03 as BCK03C from BCK1 where BCK01 = ?", patient.BCK01C).Get(&patientInfo)
		hospitalDate := patientInfo.VAE11.ParseToSecond()
		// 查询诊断症状
		_, err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAO19 > ? and VAO2.ACF01 = 2 and VAO2.VAA01 = ? order by VAO2.VAO19 desc", hospitalDate, patient.VAA01).Get(&patientInfo)
		// 8.判断性别
		//if patientInfo.ABW01 == "1" || patientInfo.ABW01 == "M" {
		//	patientInfo.Gender = "男"
		//} else if patient.ABW01 == "2" || patientInfo.ABW01 == "F" {
		//	patientInfo.Gender = "女"
		//} else {
		//	patientInfo.Gender = "未知"
		//}
		responseObj = append(responseObj, patientInfo)
	}

	return responseObj, err
}

/*查询当前科室在床病人的详情*/
func QueryPatientInfo(patientId, departmentId int) ([]PatientInfo, error) {
	responseObj := make([]PatientInfo, 0)
	patient := VAA1Dup{}
	// 查询病人详情
	_, err := fit.MySqlEngine().SQL("select VAA01, BDP02, BCK01C, BCQ04, ABQ02 from VAA1 where VAA01 = ?", patientId).Get(&patient)

	if err == nil && patient.VAA01 != 0 {
		patientInfo := PatientInfo{}
		patientInfo.VAA1Dup = patient
		// 费用：VAK1.VAK20, VAK1.VAK21 (VAK1.VAA01 = ? )
		//_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C, VAE1.VAE11, BBY1.BCF01 as AAG01 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", patient.VAA01).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BBY1.BCF01 as AAG01,VAE1.VAE11,VAE1.VAE94 as VAA04,VAE1.VAE95 as VAA05,VAE1.VAE96 as ABW01,Case VAE1.VAE96 when 1 then '男' when 2 then '女' else '未知' end as Gender,VAE1.VAE46 as VAA10 from VAE1, BBY1 where VAE1.VAA01 = ? and VAE1.VAE44 in (2,1) and BBY1.BBY01 = VAE1.AAG01 order by VAE1.VAE11 desc", patient.VAA01).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BCK03 as BCK03C from BCK1 where BCK01 = ?", departmentId).Get(&patientInfo)
		hospitalDate := patientInfo.VAE11.ParseToSecond()
		// 查询诊断症状
		_, err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAO19 > ? and VAO2.ACF01 = 2 and VAO2.VAA01 = ? order by VAO2.VAO19 desc", hospitalDate, patient.VAA01).Get(&patientInfo)

		// 8.判断性别
		//if patientInfo.ABW01 == "1" || patientInfo.ABW01 == "M" {
		//	patientInfo.Gender = "男"
		//} else if patient.ABW01 == "2" || patientInfo.ABW01 == "F" {
		//	patientInfo.Gender = "女"
		//} else {
		//	patientInfo.Gender = "未知"
		//}
		responseObj = append(responseObj, patientInfo)
	}

	return responseObj, err
}

/*查询住院期间的么个时间前后的手术记录 */
func FetchOperationRecordsDuringHospitalization(pid int64, startdate string, enddate string ) ([]VAT1, error) {
	records := make([]VAT1, 0)
	// VAT04 = 4 表示已结束手术  3正在手术中  2准备手术
	err := fit.SQLServerEngine().SQL("select VAA01, VAT08, VAT04 from VAT1 where VAA01 = ? and VAT08 > ? and VAT08 < ? and VAT04 in(2,3,4)", pid, startdate, enddate).Find(&records)
	return records, err
}
