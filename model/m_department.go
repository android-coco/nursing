//  Created by JP

package model

import (
	"fit"
	"errors"
)

type Department struct {
	DepartmentID   int    `json:"department_id"`   // 科室ID
	DepartmentName string `json:"department_name"` // 科室名
}

// 科室表
type BCK1 struct {
	BCK01 int    `json:"department_id"`   // 科室ID
	BCK02 string `json:"department_code"` // 科室编码
	BCK03 string `json:"department_name"` // 科室名称
}

/*
获取医院的科室列表
wardOnly = true 只返回病区科室，  wardOnly = false  返回全院所有的科室
*/
func QueryDepartmentList(wardOnly bool) ([]BCK1, error) {
	slice_BCK := make([]BCK1, 0)
	if wardOnly == true {
		err_BCK := fit.SQLServerEngine().SQL("select BCK01, BCK02, BCK03 from BCK1 where BCK1.BCK01A = 141").Find(&slice_BCK)
		return slice_BCK, err_BCK
	}

	err_BCK := fit.SQLServerEngine().SQL("select BCK01, BCK02, BCK03 from BCK1").Find(&slice_BCK)
	return slice_BCK, err_BCK
}

/*根据科室ID获取科室名*/
func QueryDepartmentNameWithId(id int) (name string , err error) {
	if id == 0 {
		return "", errors.New("科室ID为0")
	} else {
		department := BCK1{}
		_, err = fit.SQLServerEngine().SQL("select BCK03 from BCK1 where BCK01 = ?", id).Get(&department)
		return department.BCK03, err
	}
}


// query BCK1 table with BCK01(departmentID)
func QueryDepartmentWithDID(departmentID int) (BCK1, error) {
	bck := BCK1{}
	_, err_BCK := fit.SQLServerEngine().Table("BCK1").Where("BCK01 = ?", departmentID).Get(&bck)
	return bck, err_BCK
}