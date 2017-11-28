package model

import (
	"fit"
	"errors"
)

type BCE1Dup struct {
	BCE1         `xorm:"extends"`
	BCK03 string `json:"department_name"` // 科室名
}

type Account struct {
	User         `xorm:"extends"`
	Key   string `json:"key"`             // 密码的明码
	BCK03 string `json:"department_name"` // 科室名
}

// 员工表
type BCE1 struct {
	BCE01 int    `json:"user_id"`       // 员工ID
	BCE02 string `json:"code"`          // 工号
	BCE03 string `json:"name"`          // 姓名
	BCK01 int    `json:"department_id"` // 科室ID
}

/*接口返回的数据*/
type UserInfo struct {
	UID         int          `json:"user_id"`     // 员工ID
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
	Employeeid   int    `json:"user_id"`       // 员工ID,值为0时不连到员工表
	Authority    int    `json:"authority"`     // 权限等级
	DepartmentID int    `json:"department_id"` // 科室ID
	Status       int    `json:"status"`        // 状态，1：正常，0：停用
}

/*用户详情(存入session)*/
type UserInfoDup struct {
	UID            int    // 员工ID
	Code           string // 工号
	Name           string // 姓名
	Password       string // 密码（加密）
	Authority      int    // 权限等级
	DepartmentID   int    // 科室ID（取自Departments）
	DepartmentName string // 科室名
	Status         int    // 状态，1：正常，0：停用
}

// query User table with code and password
func CheckingUserCodeAndPwd(code, password string) ([]User, error) {
	slice_User := make([]User, 0)
	err_User := fit.MySqlEngine().Table("User").Where("Code = ? and Password = ?", code, password).Find(&slice_User)
	return slice_User, err_User
}

/*更改密码*/
func ChangePasswordWith(code, password, key string) error {
	_, err_db := fit.MySqlEngine().Exec("UPDATE User SET User.Password = ?, User.Key = ? WHERE User.Code = ?", password, key, code)
	return err_db
}

/*查询除本身账号以外的本科室已创建的账号*/
func FetchAllOfTheAccountHasCreated(departmentId, authority, status, uid int) (slice []Account, err error) {
	slice = make([]Account, 0)
	if status == 0 {
		err = fit.MySqlEngine().SQL("select * from User where Departmentid = ? and Authority = ? and Employeeid != ?", departmentId, authority, uid).Find(&slice)
	} else {
		err = fit.MySqlEngine().SQL("select * from User where Departmentid = ? and Authority = ? and Status = ? and Employeeid != ?", departmentId, authority, status, uid).Find(&slice)
	}
	if err == nil {
		for i, _ := range slice {
			slice[i].BCK03, err = QueryDepartmentNameWithId(departmentId)
			if err != nil {
				break
			}
		}
	}
	return
}

/*获取本科室未创建的用户*/
func FetchAllOfTheAccountNotBeenCreated(departmentId int) (slice []BCE1Dup, err error) {
	slice = make([]BCE1Dup, 0)
	err = fit.MySqlEngine().SQL("select BCE1.BCE01, BCE1.BCE02, BCE1.BCE03, BCE1.BCK01 from BCE1 where (select count(1) as num from `User` where `User`.departmentid = ? and BCE1.BCE01 = `User`.employeeid) = 0 and BCE1.BCK01 = ? and BCE41 in ('0','1')", departmentId, departmentId).Find(&slice)
	if err == nil {
		for i, _ := range slice {
			slice[i].BCK03, err = QueryDepartmentNameWithId(departmentId)
			if err != nil {
				break
			}
		}
	}
	return
}

/*根据工号查询账号（不包含科室名）*/
func FetchAccountWithCode(code string) (Account, error) {
	user := Account{}
	_, err_User := fit.MySqlEngine().Table("User").Omit("id", "createdate", "BCK03").Where("code = ?", code).Get(&user)
	return user, err_User
}

/*根据UID查询账号（不包含科室名）*/
func FetchAccountWithUid(uid string) (Account, error) {
	user := Account{}
	_, err_User := fit.MySqlEngine().Table("User").Omit("id", "createdate", "BCK03").Where("employeeid = ?", uid).Get(&user)
	return user, err_User
}

/*从His系统获取账号数据*/
func FetchAccountFromHis(code string) (BCE1, error) {
	user := BCE1{}
	_, err_User := fit.MySqlEngine().SQL("select BCE01,BCE02,BCE03,BCK01 from BCE1 where BCE02 = ?", code).Get(&user)
	if err_User != nil {
		return user, err_User
	} else {
		if user.BCE01 == 0 && user.BCE02 != code {
			return user, errors.New("账号不存在")
		} else {
			return user, err_User
		}
	}
}

/*更改账号的权限以及有效状态*/
func (acc Account) UpdateAccountAuthorityAndStatus() error {
	_, err := fit.MySqlEngine().Exec("update User set User.authority = ?, User.status = ? where User.employeeid = ? and User.authority < 2", acc.Authority, acc.Status, acc.Employeeid)
	return err
}

/*创建新账号*/
func CreateAccountWithOriginalUserinfo(bce BCE1, authority, status int) error {
	const pwd = "123456"
	//var pwd_sha = utils.Sha1Encryption(pwd)
	const pwd_sha = "7c4a8d09ca3762af61e59520943dc26494f8941b"
	account := Account{
		User: User{
			Username:     bce.BCE03,
			Code:         bce.BCE02,
			Password:     pwd_sha,
			Employeeid:   bce.BCE01,
			Authority:    authority,
			DepartmentID: bce.BCK01,
			Status:       status,
		},
		Key: pwd,
	}
	_, err := fit.MySqlEngine().Table("User").Omit("BCK03").InsertOne(&account)
	return err
}
