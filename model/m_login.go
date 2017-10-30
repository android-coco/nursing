package model

import (
	"fit"
)

// 员工对照表
//type IAN1 struct {
//	IAN01 int // 表ID
//	BCE01 int // 员工ID
//}

// 用户表
//type SYS_Users struct {
//	Code       string // 工号
//	EmployeeID int    // 员工对照表的IAN01，值为0时不连到员工表
//	Name       string // 姓名
//	Password   string // 密码（加密）
//	Authorized int    // 授权，0：正常  1：无效
//}

// 员工表
type BCE1 struct {
	BCE01 int    // 员工ID
	BCE02 string // 工号
	BCE03 string // 姓名
	BCK01 int    // 科室ID
}

type UserInfo struct {
	UID         uint64       `json:"user_id"`     // 员工ID
	Code        string       `json:"code"`        // 工号
	Name        string       `json:"name"`        // 姓名
	Password    string       `json:"password"`    // 密码（加密）
	Departments []Department `json:"departments"` // 科室
	Authority   int          `json:"authority"`   // 权限等级
}

/*用户表*/
type User struct {
	Username     string `json:"name"`          // 姓名
	Code         string `json:"code"`          // 工号
	Password     string `json:"password"`      // 密码（sha1加密）
	Employeeid   uint64 `json:"user_id"`       // 员工ID,值为0时不连到员工表
	Authority    int    `json:"authority"`     // 权限等级
	DepartmentID int    `json:"department_id"` // 科室ID
}

/*用户详情(存入session)*/
type UserInfoDup struct {
	UID            uint64       // 员工ID
	Code           string       // 工号
	Name           string       // 姓名
	Password       string       // 密码（加密）
	Authority      int          // 权限等级
	DepartmentID   int          // 科室ID（取自Departments）
	DepartmentName string       // 科室名
}

// query SYS_Users table with code and password
func QueryUserTable(code, password string) ([]User, error) {
	slice_User := make([]User, 0)
	err_User := fit.MySqlEngine().Table("User").Where("Code = ? and Password = ?", code, password).Find(&slice_User)
	return slice_User, err_User
}

// query IAN1 table with IAN01 (employeeID)
//func QueryEmployeeContrastTable(employeeID int) ([]IAN1, error) {
//	slice_IAN := make([]IAN1, 0)
//	err_IAN := fit.SQLServerEngine().Table("IAN1").Where("IAN01 = ?", employeeID).Find(&slice_IAN)
//	return slice_IAN, err_IAN
//}

// query BCE1 table with BCE01(user_id) and BCE02(code)
func QueryEmployeeTable(UID int) ([]BCE1, error) {
	slice_BCE := make([]BCE1, 0)
	err_BCE := fit.MySqlEngine().Table("BCE1").Where("BCE01 = ?", UID).Find(&slice_BCE)
	return slice_BCE, err_BCE
}

// query BCK1 table with BCK01(departmentID)
func QueryDepartmentWithDID(departmentID int) (BCK1, error) {
	bck := BCK1{}
	_, err_BCK := fit.SQLServerEngine().Table("BCK1").Where("BCK01 = ?", departmentID).Get(&bck)
	return bck, err_BCK
}

