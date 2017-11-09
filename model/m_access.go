package model

import (
	"fit"
	"fmt"
	"strconv"
)

type AccessType int
type AccessReason int

const (
	AccessTypeBack AccessType = 1 << iota
	AccessTypeOut
	AccessTypeAll
)

const (
	AccessReasonCheck     AccessReason = 1 << iota
	AccessReasonOperation
	AccessReasonOther
)

type Access struct {
	BaseModel                 `xorm:"extends"`
	ClassId      string       `json:"class_id" xorm:"comment(科室)"`
	PatientName  string       `json:"patient_name" xorm:"comment(病人姓名)"`
	//NurseName    string       `json:"nurse_name" xorm:"comment(责任护士姓名)"`
	BedId        string       `json:"bed_id" xorm:"comment(床号)"`
	AccessTime   string       `json:"access_time" xorm:"notnull comment(出入时间)"`
	AccessType   AccessType   `json:"access_type" xorm:"notnull"`
	AccessReason AccessReason `json:"access_reason" xorm:"notnull"`
}

func (m Access) InsertData() (int64, error) {
	id, err := fit.MySqlEngine().Insert(m)
	return id, err
}

func (m Access) ParseAccessType(value string) (AccessType, error) {
	switch value {
	case "1":
		return AccessTypeBack, nil
	case "2":
		return AccessTypeOut, nil
	case "4":
		return AccessTypeAll, nil
	default:
		return AccessTypeAll, fmt.Errorf("not fund such of AccessType: %v", value)
	}
}

func (m Access) ParseAccessReason(value string) (AccessReason, error) {
	switch value {
	case "1":
		return AccessReasonCheck, nil
	case "2":
		return AccessReasonOperation, nil
	case "3":
		return AccessReasonOther, nil

	default:
		return AccessReasonOther, fmt.Errorf("not fund such of AccessReason: %v", value)
	}
}

func AccessList(classId, page string, accessType AccessType) ([]Access, error) {
	var mods []Access
	var err error
	pageInt,_ := strconv.ParseInt(page, 10, 32)
	if pageInt == -1 { // 全部
		pageInt = 0
		if accessType == AccessTypeAll {
			err = fit.MySqlEngine().SQL("select * from Access where classId = ? ORDER BY `AccessTime` DESC", classId).Limit(20, int(pageInt)).Find(&mods)
		} else {
			err = fit.MySqlEngine().SQL("select * from Access where classId = ? and accesstype = ? ORDER BY `AccessTime` DESC", classId, accessType).Limit(20, int(pageInt)).Find(&mods)
		}
	} else {
		pageInt *= 20
		if accessType == AccessTypeAll {
			err = fit.MySqlEngine().SQL("select * from Access where classId = ? ORDER BY `AccessTime` DESC", classId).Limit(20, int(pageInt)).Find(&mods)
		} else {
			err = fit.MySqlEngine().SQL("select * from Access where classId = ? and accesstype = ? ORDER BY `AccessT ime` DESC", classId, accessType).Limit(20, int(pageInt)).Find(&mods)
		}
	}

	return mods, err
}

func AccessSearch(classId, paramstr string) ([]Access, error) {
	mods := make([]Access,0)
	params := "%" + paramstr + "%"
	err := fit.MySqlEngine().SQL("select * from Access where classId = ? and (bedId like ? or patientName like ?)", classId, params, params).Find(&mods)
	return mods, err
}


func AccessALLList(classId string) ([]Access, error) {
	var mods []Access
	var err error
	err = fit.MySqlEngine().SQL("select * from Access where classId = ?  ORDER BY `AccessTime` DESC", classId).Find(&mods)
	return mods, err
}
