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

/*
查询除本身账号以外的账号
*/
func FetchAllOfTheAccountHasCreated(departmentId, authority, status, uid int) (slice []Account, err error) {
	slice = make([]Account, 0)
	if status == 0 {
		err = fit.MySqlEngine().SQL("select * from User where departmentid = ? and authority = ? and employeeid != ?", departmentId, authority, uid).Find(&slice)
	} else {
		err = fit.MySqlEngine().SQL("select * from User where departmentid = ? and authority = ? and status = ? and employeeid != ?", departmentId, authority, status, uid).Find(&slice)
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

func FetchAllOfTheAccountNotBeenCreated(departmentId int) (slice []BCE1Dup, err error) {
	slice = make([]BCE1Dup, 0)
	err = fit.MySqlEngine().SQL("select BCE1.BCE01, BCE1.BCE02, BCE1.BCE03, BCE1.BCK01 from BCE1 where (select count(1) as num from `User` where `User`.departmentid = ? and BCE1.BCE01 = `User`.employeeid) = 0 and BCE1.BCK01 = ?", departmentId, departmentId).Find(&slice)
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

func FetchAccountWithCode(code string) (Account, error) {
	user := Account{}
	_, err_User := fit.MySqlEngine().Table("User").Omit("id","createdate","BCK03").Where("code = ?", code).Get(&user)
	return user, err_User
}

func FetchAccountWithUid(uid string) (Account, error) {
	user := Account{}
	_, err_User := fit.MySqlEngine().Table("User").Omit("id","createdate","BCK03").Where("employeeid = ?", uid).Get(&user)
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
func (acc Account)UpdateAccountAuthorityAndStatus() error {
	_, err := fit.MySqlEngine().Exec("update User set User.authority = ?, User.status = ? where User.employeeid = ? and User.authority < 2",acc.Authority, acc.Status, acc.Employeeid)
	return err
}


func CreateAccountWithOriginalUserinfo(bce BCE1,authority, status int) error {
	const pwd = "123456"
	//var pwd_sha = utils.Sha1Encryption(pwd)
	const pwd_sha = "7c4a8d09ca3762af61e59520943dc26494f8941b"
	account := Account{
		User:User{
			Username:bce.BCE03,
			Code:bce.BCE02,
			Password:pwd_sha,
			Employeeid:uint64(bce.BCE01),
			Authority:authority,
			DepartmentID:bce.BCK01,
			Status:status,
		},
		Key:pwd,
	}
	_, err := fit.MySqlEngine().Table("User").Omit("BCK03").InsertOne(&account)
	return err
}
