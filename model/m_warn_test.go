package model

import (
	"testing"
	"fit"
	"fmt"
	"time"
)

func TestBaseModel_InsertData(t *testing.T) {
	warnModle := Warn{Base:BaseModel{PatientId:"123", NurseId:"888"}, Name:"hehehehehe", WarnTime:time.Now()}

	m, err := fit.Engine().Insert(warnModle)
	fmt.Println(m, err)
}
