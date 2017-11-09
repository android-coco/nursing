//  Created by JP

package model

import (
	"fit"
)

/*病人详情*/
type PatientInfo struct {
	VAE1          `xorm:"extends"`
	VAK1          `xorm:"extends"`
	VAA1Dup       `xorm:"extends"`
	VAO2          `xorm:"extends"`
	VCF1          `xorm:"extends"`
	BCK03  string `json:"department_name"` // 科室名称
	Gender string `json:"gender"`          // 性别(结果字符串)
}

/*病人基本信息表*/
type VAA1Dup struct {
	VAA1         `xorm:"extends"`
	BDP02 string `json:"type"`   // type病人类型 BDP1表 自费
	VAA61 int    `json:"status"` // 2: 住院
	//ABQ02 string `json:"nation"` // 民族
}

/*病人登录记录*/
type VAE1 struct {
	//VAA01  int    `json:"patient_id"`   // 病人ID
	BCE03B string `json:"nurse_name"`     // 责任护士ID
	BCE03C string `json:"physician_name"` // 住院医师
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

/*查询病人详情，不限科室，不限住院、出院*/
func GetPatientInfo(patientId string) ([]PatientInfo, error) {
	responseObj := make([]PatientInfo, 0)
	patient := VAA1Dup{}
	// 查询病人详情
	_, err := fit.MySqlEngine().SQL("select VAA01, VAA04, VAA05, ABW01, VAA10, BDP02, VAA61, BCK01B, BCQ04, VAA73, ABQ02 from VAA1 where VAA01 = ?", patientId).Get(&patient)

	if err == nil && patient.VAA01 != 0 {
		patientInfo := PatientInfo{}
		patientInfo.VAA1Dup = patient
		// 费用：VAK1.VAK20, VAK1.VAK21 (VAK1.VAA01 = ? )
		//_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C from VAE1 where VAE1.VAA01 = ?", patientId).Limit(1, 0).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C from VAE1 where VAE1.VAA01 = ? and VAE1.BCQ04B = ? order by VAE1.VAE01 desc", patientId, patient.BCQ04).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BCK03 from BCK1 where BCK01 = ?", patient.BCK01B).Get(&patientInfo)

		// 查询诊断症状
		//illness := make([]VAO2, 0)
		//err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAA01 = ? order by VAO19 desc", patientId).Limit(1, 0).Find(&illness)
		//if length := len(illness); length > 0 {
		//	patientInfo.VAO2 = illness[0]
		//} else {
		//	patientInfo.VAO15 = ""
		//}
		_, err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAA01 = ? order by VAO2.VAO19 desc", patient.VAA01).Get(&patientInfo)

		// 8.判断性别
		if patientInfo.ABW01 == "1" || patientInfo.ABW01 == "M" {
			patientInfo.Gender = "男"
		} else if patient.ABW01 == "2" || patientInfo.ABW01 == "F" {
			patientInfo.Gender = "女"
		} else {
			patientInfo.Gender = "未知"
		}

		// 查询护理级别
		//nursingDegree := VCF1{}
		//_, _ = fit.SQLServerEngine().SQL("select AAG01 from VCF1 where VAA01 = ?", patientId).Desc("VCF01").Get(&nursingDegree)
		//patientInfo.AAG01 = nursingDegree.AAG01
		_, err = fit.SQLServerEngine().SQL("select AAG01 from VCF1 where VAA01 = ? order by VCF01 desc", patient.VAA01).Get(&patientInfo)
		responseObj = append(responseObj, patientInfo)
	}

	return responseObj, err
}

/*查询当前科室在床病人的详情*/
func QueryPatientInfo(patientId, departmentId int) ([]PatientInfo, error) {
	responseObj := make([]PatientInfo, 0)
	patient := VAA1Dup{}
	// 查询病人详情
	_, err := fit.MySqlEngine().SQL("select VAA01, VAA04, VAA05, ABW01, VAA10, BDP02, VAA61, BCK01B, BCQ04, VAA73, ABQ02 from VAA1 where VAA61 = 2 and BCQ04 != 0 and BCK01B = ? and VAA01 = ?", departmentId, patientId).Get(&patient)

	if err == nil && patient.VAA01 != 0 {
		patientInfo := PatientInfo{}
		patientInfo.VAA1Dup = patient
		// 费用：VAK1.VAK20, VAK1.VAK21 (VAK1.VAA01 = ? )
		_, err = fit.SQLServerEngine().SQL("select VAE1.BCE03B, VAE1.BCE03C from VAE1 where VAE1.VAA01 = ?", patientId).Limit(1, 0).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BCK03 from BCK1 where BCK01 = ?", patient.BCK01B).Get(&patientInfo)

		// 查询诊断症状
		illness := make([]VAO2, 0)
		err = fit.SQLServerEngine().SQL("select VAO2.VAO15 from VAO2 where VAO2.VAA01 = ? order by VAO19 desc", patientId).Limit(1, 0).Find(&illness)
		if length := len(illness); length > 0 {
			patientInfo.VAO2 = illness[0]
		} else {
			patientInfo.VAO15 = ""
		}

		// 8.判断性别
		if patientInfo.ABW01 == "1" || patientInfo.ABW01 == "M" {
			patientInfo.Gender = "男"
		} else if patient.ABW01 == "2" || patientInfo.ABW01 == "F" {
			patientInfo.Gender = "女"
		} else {
			patientInfo.Gender = "未知"
		}

		// 查询护理级别
		nursingDegree := VCF1{}
		_, _ = fit.SQLServerEngine().SQL("select AAG01 from VCF1 where VAA01 = ?", patientId).Desc("VCF01").Get(&nursingDegree)
		patientInfo.AAG01 = nursingDegree.AAG01
		responseObj = append(responseObj, patientInfo)
	}

	return responseObj, err
}
