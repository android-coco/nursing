//  Created by JP

package model

import (
	"fit"
	"time"
	"fmt"
)

/*病人详情*/
type PatientInfo struct {
	VAE1             `xorm:"extends"`
	VAK1             `xorm:"extends"`
	VAA1Dup          `xorm:"extends"`
	VAO2             `xorm:"extends"`
	VCF1             `xorm:"extends"`
	BCK03C    string `json:"department_name"` // 病区名称
	Gender    string `json:"gender"`          // 性别(结果字符串)
	EntryDate string `json:"entry_date"`      // 入科日期
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
	Vid    int64                  `json:"vid"`            // 就诊ID
	BCE03B string                 `json:"nurse_name"`     // 责任护士ID
	BCE03C string                 `json:"physician_name"` // 住院医师
	VAE11  DatetimeWithoutSeconds `json:"hospital_date"`  // 入院日期
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

/*入科日期绑定*/
type EntryDepartment struct {
	IsExist   int    // 是否已录入入科日期，是：1，否：0
	Pid       int64  // 病人ID  VAA01
	EntryDate string // 入科日期+时间
	EntryDay  string // 入科日期
	EntryTime string // 入科日期
}

/*查询病人详情，不限科室，不限住院、出院*/
func GetPatientInfo(patientId string) ([]PatientInfo, error) {
	responseObj := make([]PatientInfo, 0)
	// 查询病人详情
	//fit.SQLServerEngine().ShowSQL(true)
	patientInfo := PatientInfo{}
	sqlStr := fmt.Sprintf("SELECT TOP 1 a.VAE01 Vid, a.VAE95 VAA05, a.VAA01, a.BCK01C, a.VAE44 VAA61, a.BCQ04B BCQ04, a.VAE96 ABW01, a.VAE94 VAA04, a.VAE46 VAA10, a.BDP02, CASE a.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, a.BCE03B, a.BCE03C, a.VAE11, b.BCF01 AAG01 FROM VAE1 a LEFT JOIN BBY1 b ON b.BBY01 = a.AAG01 WHERE a.VAA01 = %s AND a.VAE44 = 2 ORDER BY a.VAE11 DESC", patientId)
	_, err := fit.SQLServerEngine().SQL(sqlStr).Get(&patientInfo)
	if err == nil && patientInfo.VAA01 != 0 {
		_, err = fit.MySqlEngine().SQL("select ABQ02 from VAA1 where VAA01 = ? order by VAA73 desc", patientId).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BCK03 as BCK03C from BCK1 where BCK01 = ?", patientInfo.BCK01C).Get(&patientInfo)
		hospitalDate := patientInfo.VAE11.ParseToSecond()
		// 查询诊断症状

		// 入科日期
		_, err = fit.MySqlEngine().SQL("select DATE_FORMAT(EntryDate,'%Y-%m-%d %H:%i') as EntryDate from departmenttransfer where Pid = ?",patientId).Get(&patientInfo)

		sqlStr = fmt.Sprintf("SELECT top 1 VAO2.VAO15 FROM VAO2 WHERE VAO2.ACF01 = 2 AND VAO2.VAA01 = %s AND VAO2.VAO19 > '%s' ORDER BY VAO2.VAO19 DESC", patientId, hospitalDate)
		_, err = fit.SQLServerEngine().SQL(sqlStr).Get(&patientInfo)
		responseObj = append(responseObj, patientInfo)
	}
	//fit.SQLServerEngine().ShowSQL(false)

	return responseObj, err
}

/*查询当前科室在床病人的详情*/
func QueryPatientInfo(patientId, departmentId int) ([]PatientInfo, error) {
	responseObj := make([]PatientInfo, 0)
	// 查询病人详情
	patientInfo := PatientInfo{}
	_, err := fit.SQLServerEngine().SQL("SELECT TOP 1 a.VAE01 Vid, a.VAE95 VAA05, a.VAA01, a.BCK01C, a.VAE44 VAA61, a.BCQ04B BCQ04, a.VAE96 ABW01, a.VAE94 VAA04, a.VAE46 VAA10, a.BDP02, CASE a.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, a.BCE03B, a.BCE03C, a.VAE11, b.BCF01 AAG01 FROM VAE1 a LEFT JOIN BBY1 b ON b.BBY01 = a.AAG01 WHERE a.BCK01C = ? and a.VAA01 = ? AND a.VAE44 = 2 ORDER BY a.VAE11 DESC", departmentId, patientId).Get(&patientInfo)
	if err == nil && patientInfo.VAA01 != 0 {
		_, err = fit.MySqlEngine().SQL("select ABQ02 from VAA1 where VAA01 = ? order by VAA73 desc", patientId).Get(&patientInfo)
		_, err = fit.SQLServerEngine().SQL("select BCK03 as BCK03C from BCK1 where BCK01 = ?", patientInfo.BCK01C).Get(&patientInfo)
		hospitalDate := patientInfo.VAE11.ParseToSecond()
		// 入科日期
		_, err = fit.MySqlEngine().SQL("select DATE_FORMAT(EntryDate,'%Y-%m-%d %H:%i') as EntryDate from departmenttransfer where Pid = ?",patientId).Get(&patientInfo)

		// 查询诊断症状
		_, err = fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT top 1 VAO2.VAO15 FROM VAO2 WHERE VAO2.ACF01 = 2 AND VAO2.VAA01 = %d AND VAO2.VAO19 > '%s' ORDER BY VAO2.VAO19 DESC", patientId, hospitalDate)).Get(&patientInfo)
		responseObj = append(responseObj, patientInfo)
	}
	return responseObj, err
}

/*查询住院期间的么个时间前后的手术记录 */
func FetchOperationRecordsDuringHospitalization(pid int64, startdate string, enddate string) ([]VAT1, error) {
	records := make([]VAT1, 0)
	// VAT04 = 4 表示已结束手术  3正在手术中  2准备手术
	err := fit.SQLServerEngine().SQL("select VAA01, VAT08, VAT04 from VAT1 where VAA01 = ? and VAT08 > ? and VAT08 < ? and VAT04 in(2,3,4)", pid, startdate, enddate).Find(&records)
	return records, err
}

/*获取病人的入科日期*/
func FetchEntryDepartmentDate(pid int64) EntryDepartment {
	obj := EntryDepartment{}
	_, err := fit.MySqlEngine().SQL("SELECT if(count(1) > 0, 1, 0) as IsExist, Pid, DATE_FORMAT(EntryDate,'%Y-%m-%d %H:%i') as EntryDate, DATE_FORMAT(EntryDate,'%Y-%m-%d') as EntryDay, DATE_FORMAT(EntryDate,'%H:%i') as EntryTime from departmenttransfer where Pid = ?", pid).Get(&obj)
	if err != nil {
		fit.Logger().LogDebug("***JK***", "查询病人入科日期出错,Pid:"+fmt.Sprintf("%d", pid)+err.Error())
	}
	return obj
}

/*录入入科日期*/
func EnteringEntryDepartmentDate(pid int64, date string) error {
	if exist := IsExistRecord(true, "departmenttransfer", fmt.Sprintf("Pid = %d", pid)); exist.Exist == 1 {
		_, err := fit.MySqlEngine().Exec("update departmenttransfer set EntryDate = ? where Pid = ?", date, pid)
		return err
	} else {
		_, err := fit.MySqlEngine().Exec("insert into departmenttransfer (Pid, EntryDate) VALUES(?,?)", pid, date)
		return err
	}
	return nil
}
