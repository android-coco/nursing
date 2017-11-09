package model

import (
	"fit"
	"time"
	"fmt"
)

type Advice struct {
	VAF01 int    `json:"advice_id"`        //医嘱id
	VAA01 int    `json:"patient_id"`       //病人id
	VAF36 string `json:"start_time"`       //开始执行时间
	VAF37 string `json:"end_time"`         //结束执行时间
	VAF11 int    `json:"advice_type"`      //长期或临时
	VAF10 int    `json:"advice_state"`     //状态
	VAF53 int    `json:"advice_execution"` //用药方式

	VAF22 string  `json:"advice_msg"`         //医嘱内容
	VAF17 int     `json:"day_number"`         //天数
	VAF18 float32 `json:"single_dose"`        //剂量, 单次用量
	VAF19 string  `json:"dosage"`             //用量
	VAF20 float32 `json:"single"`             //单量
	VAF21 float32 `json:"number"`             //数量
	VAF26 string  `json:"perform_frequency"`  //执行频次
	VAF27 int     `json:"frequency"`          //频率次数
	VAF28 int     `json:"frequency_interval"` //频率间隔
	VAF29 string  `json:"interval"`           //间隔单位

	BCE03A string `json:"doctor_name"` //开嘱医生名称
	VAF42  string `json:"VAF42"`       //开嘱时间
	BCE03B string `json:"BCE03B"`      //开嘱护士
	BCE03C string `json:"BCE03C"`      //校对护士
	VAF45  string `json:"VAF45"`       //校对时间
	BCE03D string `json:"BCE03D"`      //停嘱医生
	VAF47  string `json:"VAF47"`       //停嘱时间
	BCE03E string `json:"BCE03E"`      //停嘱护士
	BCE03F string `json:"BCE03F"`      //停嘱校对护士
	VAF50  string `json:"VAF50"`       //执行停嘱时间
}

type MedicalAdviceDup struct {
	VAF01 int64 // 医嘱记录ID
}

func OutAdvice(sql string, msg ...interface{}) ([]Advice, error) {
	advices := make([]Advice, 0)
	err := fit.MySqlEngine().Table("VAF2").Where(sql, msg...).Find(&advices)
	return advices, err
}

//得到需要执行的医嘱
func GetNewAdvice(patient_id int) ([]Advice, error) {
	advices := make([]Advice, 0)
	t := time.Now()
	starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	err := fit.SQLServerEngine().Table("VAF2").Where("(VAF11 = ? and VAF36 >= ? and  VAF10 =? and VAA01 = ? and BCE03D = ?) or (VAF11 = ? and  VAF10 =? and VAA01 = ? and BCE03D = ?)", 2, starttime, 8, patient_id, "", 1, 8, patient_id, "").Find(&advices)
	return advices, err
}

//得到未执行的医嘱
func GetNonExecutionAdvice(patient_id int) ([]Advice, error) {
	advices, err := GetNewAdvice(patient_id)

	if err != nil || len(advices) == 0 {
		return advices, err
	} else {
		t := time.Now()
		starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

		results := make([]Advice, 0)
		for _, i := range advices {
			advicestates, err := OutAdviceState("advicestateId = ? and time >= ?", i.VAF01, starttime)
			if err == nil && len(advicestates) == 0 {
				results = append(results, i)
			}
		}
		return results, nil
	}

	return advices, err
}

type IsExist struct {
	Exist int  // 是否存在
}

/*JP 是否存在新医嘱*/
func IsExistNewMedicalAdvice(pid int) int {
	mAdvice := make([]MedicalAdviceDup, 0)
	t := time.Now()
	starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	err_db := fit.SQLServerEngine().SQL("SELECT VAF01 FROM VAF2 WHERE VAA01 = ? and VAF10 = 8 and BCE03D = '' and (VAF11 = 1 or (VAF11 = 2 and VAF36 >= ?))", pid, starttime).Find(&mAdvice)

	length := len(mAdvice)
	if length == 0 {
		if err_db != nil {
			fit.Logger().LogError("**JP**", err_db)
		}
		return 0
	}

	tempStr := ""
	var i int
	for i = 0; i < length; i ++ {
		tempStr = tempStr + fmt.Sprintf("%d", mAdvice[i].VAF01)
		if i < (length - 1) {
			tempStr = tempStr + ","
		}
	}

	isEx := IsExist{}
	fit.MySqlEngine().SQL("select if(count(1) > 0, 0, 1) as Exist from AdviceState where patientid = ? and time >= ? and advicestateId in (?) order by id", pid, starttime, tempStr).Get(&isEx)
	return isEx.Exist
}

//得到停止医嘱但未确认
func GetUncertainOewAdvice(patient_id int) ([]Advice, error) {
	t := time.Now()
	starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	advices := make([]Advice, 0)

	err := fit.SQLServerEngine().Table("VAF2").Where("(VAF11 = ? and VAA01 = ? and VAF10 =? and BCE03D != ?) or (VAF11 = ? and VAA01 = ? and VAF36 >= ? and VAF10 =? and BCE03D != ?)", 1, patient_id, 8, "", 2, patient_id, starttime, 8, "").Find(&advices)
	return advices, err
}

/*JP 是否存在已停医嘱*/
func IsExistFinishedMedicalAdvice(pid int) int  {
	t := time.Now()
	starttime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	isEx := IsExist{}
	_, err_db := fit.SQLServerEngine().SQL("SELECT count(1) as Exist FROM VAF2 WHERE VAA01 = ? and VAF10 = 8 and BCE03D != '' and (VAF11 = 1 or (VAF11 = 2 and VAF36 >= ?))",pid, starttime).Get(&isEx)
	if err_db != nil {
		fit.Logger().LogError("**JP**", err_db)
	}
	if isEx.Exist == 0 {
		return 0
	}
	return 1
}

type AdviceFit struct {
	Advice
	Advicestates []AdviceState `json:"advicestates"` //数量
	Patient      interface{}   `json:"patientinfo"`  //个人信息
}

type AdviceState struct {
	NurseId       string       `json:"nurse_id" xorm:"notnull comment(护士id)"`
	NurseName     string       `json:"nurse_name" xorm:"notnull comment(护士姓名)"`
	PatientId     int          `json:"patient_id" xorm:"notnull comment(病人id)"`
	AdviceStateId int          `json:"advicestate_id" xorm:"notnull comment(医嘱id)"`
	State         string       `json:"state" xorm:"notnull comment(医嘱状态)"`
	Period        string       `json:"period" xorm:"notnull comment(医嘱周期)"`
	Time          fit.JsonTime `json:"testtime" xorm:"notnull comment(打点时间)"`
}

func OutAdviceState(sql string, msg ...interface{}) ([]AdviceState, error) {
	advicestates := make([]AdviceState, 0)
	err := fit.MySqlEngine().Table("AdviceState").Where(sql, msg...).Find(&advicestates)
	return advicestates, err
}

func InsertAdviceState(advicestates ...AdviceState) error {
	_, err := fit.MySqlEngine().Table("AdviceState").Insert(advicestates)
	return err
}

func UpdateAdviceState(sql string, advicestate AdviceState, msg ...interface{}) error {
	_, err := fit.MySqlEngine().Table("AdviceState").Where(sql, msg...).Update(&advicestate)
	return err
}
