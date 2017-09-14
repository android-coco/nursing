package model

import (
	"fit"
	"fmt"
	"time"
)

type AccessType int
type AccessReason int

const (
	AccessTypeBack AccessType = 1 << iota
	AccessTypeOut
)

const (
	AccessReasonCheck     AccessReason = 1 << iota
	AccessReasonOperation
	AccessReasonOther
)
type Access struct {
	BaseModel    		`xorm:"extends"`
	AccessTime time.Time      `json:"access_time" xorm:"notnull comment(出入时间)"`
	AccessType   AccessType       	`json:"access_type" xorm:"notnull"`
	AccessReason AccessReason       `json:"access_reason" xorm:"notnull"`
}

func (m Access) InsertData() (int64, error)  {
	id, err := fit.Engine().Insert(m)
	return id, err
}

func (m Access) ParseAccessType(value string) (AccessType, error) {
	switch value {
	case "1":
		return AccessTypeBack, nil
	case "2":
		return AccessTypeOut, nil
	default:
		return AccessTypeBack, fmt.Errorf("not fund such of AccessType: %v", value)
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

func AccessList(nurse_id, patient_id string) []Access  {
	var mods []Access
	fit.Engine().SQL("select * from warn where nurse_id = ? and patient_id = ?", nurse_id, patient_id).Find(&mods)
	return mods
}