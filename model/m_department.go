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

func QueryDepartmentList() ([]BCK1, error) {
	slice_BCK := make([]BCK1, 0)
	err_BCK := fit.SQLServerEngine().SQL("select BCK01, BCK02, BCK03 from BCK1").Find(&slice_BCK)
	return slice_BCK, err_BCK
}


func QueryDepartmentNameWithId(id int) (name string , err error) {
	if id == 0 {
		return "", errors.New("科室ID为0")
	} else {
		department := BCK1{}
		_, err = fit.SQLServerEngine().SQL("select BCK03 from BCK1 where BCK01 = ?", id).Get(&department)
		return department.BCK03, err
	}
}