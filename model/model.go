package model


import (
	//"fit"
	//"fmt"
)

type I_BaseMoel interface {
	InsertData(a interface{}) error
	GetData(a interface{}) (b interface{})
	UpdateData(a interface{}) (b interface{}, err error)
}

type BaseModel struct {
	//Id        int64		`json:"id" xorm:"pk autoincr ->"`
	NurseId   string	`json:"nurse_id" xorm:"notnull comment(护士id)"`
	NurseName  string	`json:"nurse_name" xorm:"notnull comment(护士名字)"`
	PatientId string	`json:"patient_id" xorm:"notnull comment(病人id)"`
	//DateTime  time.Time `json:"date_time" xorm:"created"`
}

type IdModel struct {
	//Id        int64		`json:"id" xorm:"pk autoincr ->"`
	NurseId   string	`json:"nurse_id" xorm:"notnull comment(护士id)"`
	PatientId string	`json:"patient_id" xorm:"notnull comment(病人id)"`
	//DateTime  time.Time `json:"date_time" xorm:"created"`
}



/*func (mod BaseModel) InsertData(a interface{}) error {
	_, err := fit.MySqlEngine().Insert(a)
	return err
}

func (mod BaseModel) GetData(a interface{}) (b interface{}) {
	has , err := fit.MySqlEngine().Get(&a)
	if err != nil || !has {
		fmt.Println(err)
		return nil
	}
	return a
}

func (mod BaseModel) UpdateData(a interface{}) (b interface{}, err error)  {
	return fit.MySqlEngine().Update(a)
}*/
