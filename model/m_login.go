//  Created by JP

package model

import (
	"fit"
)

// 员工表
type BCE1 struct {
	BCE01 int    `json:"user_id"`       // 员工ID
	BCE02 string `json:"code"`          // 工号
	BCE03 string `json:"name"`          // 姓名
	BCK01 int    `json:"department_id"` // 科室ID
}

/*接口返回的数据*/
type UserInfo struct {
	UID         uint64       `json:"user_id"`     // 员工ID
	Code        string       `json:"code"`        // 工号
	Name        string       `json:"name"`        // 姓名
	Password    string       `json:"password"`    // 密码（加密）
	Departments []Department `json:"departments"` // 科室
	Authority   int          `json:"authority"`   // 权限等级
	Status      int          `json:"status"`      // 状态，1：正常，0：停用
}

/*用户表*/
type User struct {
	Username     string `json:"name"`          // 姓名
	Code         string `json:"code"`          // 工号
	Password     string `json:"-"`             // 密码（sha1加密）
	Employeeid   uint64 `json:"user_id"`       // 员工ID,值为0时不连到员工表
	Authority    int    `json:"authority"`     // 权限等级
	DepartmentID int    `json:"department_id"` // 科室ID
	Status       int    `json:"status"`        // 状态，1：正常，0：停用
}

/*用户详情(存入session)*/
type UserInfoDup struct {
	UID            uint64 // 员工ID
	Code           string // 工号
	Name           string // 姓名
	Password       string // 密码（加密）
	Authority      int    // 权限等级
	DepartmentID   int    // 科室ID（取自Departments）
	DepartmentName string // 科室名
	Status         int    // 状态，1：正常，0：停用
}

// query User table with code and password
func QueryUserTable(code, password string) ([]User, error) {
	slice_User := make([]User, 0)
	err_User := fit.MySqlEngine().Table("User").Where("Code = ? and Password = ?", code, password).Find(&slice_User)
	return slice_User, err_User
}

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

func ChangePasswordWith(code, password, key string) error {
	_, err_db := fit.MySqlEngine().Exec("UPDATE User SET User.password = ?, User.key = ? WHERE User.code = ?", password, key, code)
	return err_db
}
