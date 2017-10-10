package model

import (
	"fit"
	"fmt"
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
	BaseModel            `xorm:"extends"`
	ClassId      string 		`json:"class_id" xorm:"comment(科室)`
	PatientName  string        `json:"patient_name" xorm:"comment(病人姓名)`
	NurseName    string        `json:"nurse_name" xorm:"comment(护士姓名)`
	BedId        string        `json:"bed_id" xorm:"comment(床号)`
	AccessTime   string      `json:"access_time" xorm:"notnull comment(出入时间)"`
	AccessType   AccessType        `json:"access_type" xorm:"notnull"`
	AccessReason AccessReason       `json:"access_reason" xorm:"notnull"`
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

func AccessList(classId string, accessType AccessType) ([]Access, error) {
	var mods []Access
	var err error
	if accessType == AccessTypeAll {
		err = fit.MySqlEngine().SQL("select * from Access where classId = ?", classId).Find(&mods)
	} else {
		err = fit.MySqlEngine().SQL("select * from Access where classId = ? and accesstype = ?", classId, accessType).Find(&mods)
	}
	return mods, err
}

func AccessSearch(classId, paramstr string) ([]Access, error) {
	var mods []Access
	params := "%" + paramstr + "%"
	err := fit.MySqlEngine().SQL("select * from Access where classId = ? and (patientid like ? or patientName like ?)", classId, params, params).Find(&mods)
	return mods, err
}
