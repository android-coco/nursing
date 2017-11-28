package model


import (
	//"fit"
	//"fmt"
	"fmt"
	"fit"
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

func AbsInt(i int64) int64 {
	if i>0{
		return i
	}else{
		return -i
	}
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


/*JP 用法查询是否存在某条记录*/
type IsExist struct {
	Exist int // 是否存在
}

/*
查询满足条件的数据是否存在
mysql:  		数据库引擎,true=MySql,false=SqlServer
tableName:  	表名
where：  		条件
*/
func IsExistRecord(mysql bool, tableName, where string) IsExist {
	isEx := IsExist{}
	if mysql == true {
		sqlStr := fmt.Sprintf("select (count(1) > 0) as Exist from %s where %s", tableName, where)
		_, err := fit.MySqlEngine().SQL(sqlStr).Get(&isEx)
		if err != nil {
			fit.Logger().LogError("***JK***IsExistRecord***",err.Error())
		}
		return isEx
	} else {
		sqlStr := fmt.Sprintf("if exists (select 1 from %s where %s) select '1' as Exist else select '0' as Exist", tableName, where)
		_, err := fit.SQLServerEngine().SQL(sqlStr).Get(&isEx)
		if err != nil {
			fit.Logger().LogError("***JK***IsExistRecord***",err.Error())
		}
		return isEx
	}
}
