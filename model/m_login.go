package model

import "fit"

// 员工对照表
type IAN1 struct {
	IAN01 int // 表ID
	BCE01 int // 员工ID
}

// 用户表
type SYS_Users struct {
	Code       string // 工号
	EmployeeID int    // 员工对照表的IAN01，值为0时不连到员工表
	Name       string // 姓名
	Password   string // 密码（加密）
	Authorized int    // 授权，0：正常  1：无效
}

// 员工表
type BCE1 struct {
	BCE01 int    // 员工ID
	BCE02 string // 工号
	BCE03 string // 姓名
	BCK01 int    // 科室ID
}

// 科室表
type BCK1 struct {
	BCK01 int    // 科室ID
	BCK02 string // 科室编码
	BCK03 string // 科室名称
}

type User_Response struct {
	UID         int          `json:"user_id"`  // 员工ID
	Code        string       `json:"code"`     // 工号
	Name        string       `json:"name"`     // 姓名
	Password    string       `json:"password"` // 密码（加密）
	Departments []Department `json:"departments"`
}

type Department struct {
	DepartmentID   int    `json:"department_id"`   // 科室ID
	DepartmentName string `json:"department_name"` // 科室名
}

// query SYS_Users table with code and password
func QueryUserTable(code, password string) ([]SYS_Users, error) {
	slice_User := make([]SYS_Users, 0)
	err_User := fit.SQLServerEngine().Table("SYS_Users").Where("Code = ? and Password = ?", code, password).Find(&slice_User)
	return slice_User, err_User
}

// query IAN1 table with IAN01 (employeeID)
func QueryEmployeeContrastTable(employeeID int) ([]IAN1, error) {
	slice_IAN := make([]IAN1, 0)
	err_IAN := fit.SQLServerEngine().Table("IAN1").Where("IAN01 = ?", employeeID).Find(&slice_IAN)
	return slice_IAN, err_IAN
}

// query BCE1 table with BCE01(user_id) and BCE02(code)
func QueryEmployeeTable(UID int, code string) ([]BCE1, error) {
	slice_BCE := make([]BCE1, 0)
	err_BCE := fit.SQLServerEngine().Table("BCE1").Where("BCE01 = ? and BCE02 = ?", UID, code).Find(&slice_BCE)
	return slice_BCE, err_BCE
}

// query BCK1 table with BCK01(departmentID)
func QueryDepartmentTable(departmentID int) ([]BCK1, error){
	slice_BCK := make([]BCK1, 0)
	err_BCK := fit.SQLServerEngine().Table("BCK1").Where("BCK01 = ?", departmentID).Find(&slice_BCK)
	return slice_BCK, err_BCK
}