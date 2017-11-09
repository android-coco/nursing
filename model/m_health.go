package model

import (
	"fit"
	"errors"
	"time"
	"strconv"
	"github.com/go-xorm/xorm"

)

/*import "fmt"*/

const (
	Temperature_Type string = "Temperature_Type" //体温
	Pulse_Type       string = "Pulse_Type"       //脉搏
	Breathe_Type     string = "Breathe_Type"     //呼吸
	Pressure_Type    string = "Pressure_Type"    //血压
	Heartrate_Type   string = "Heartrate_Type"   //心率
	Spo2h_Type       string = "Spo2h_Type"       //血氧
	Glucose_Type     string = "Glucose_Type"     //血糖
	Weight_Type      string = "Weight_Type"      //体重
	Height_Type      string = "Height_Type"      //身高

	Skin_Type     string = "Skin_Type"     //皮试
	Ache_Type     string = "Ache_Type"     //疼痛
	Incident_Type string = "Incident_Type" //事件
)

type TimeShow interface {
	ShowTime()time.Time
}

/*体温*/
type Temperature struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1物理降温,2药物降温,3冰毯降温,4停冰毯降温,5药物+物理降温,6无降温,7不升,8外出,9检查,10请假,11拒试,12无法侧,13未测)"`
	Ttemptype   string       `json:"ttemptype" xorm:"notnull comment(体温的类型,1腋温,2耳温,3口温,4肛温,5额温)"`
	Coolingvalue  string     `json:"coolingvalue" xorm:"notnull comment(降温值)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Temperature)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputTemperature(session *xorm.Session,strData map[string]string) error {
	var item = &Temperature{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v;
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["ttemptype"]; ok {
		item.Ttemptype = v
	} else {
		return errors.New("ttemptype")
	}

	if v, ok := strData["coolingvalue"]; ok {
		item.Coolingvalue = v
	} else {
		return errors.New("coolingvalue")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v;
	} else {
		return errors.New("没有value")
	}


	return InsertTemperature(session,item)
}

func InsertTemperature(session *xorm.Session,item *Temperature) error{
	var item1 = &Temperature{}

	has,err := session.Table("Temperature").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutTemperature(sql string, msg ...interface{}) ([]Temperature, error) {
	items := make([]Temperature, 0)
	//fit.MySqlEngine().ShowSQL(true)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)

	return items, err
}

/*查询用户是否发热*/
func GetWhetherFever(patientId int) (bool, error) {

	item := Temperature{}
	has, err := fit.MySqlEngine().SQL("select Value from Temperature where patientid = ? order by Testtime desc", patientId).Get(&item)
	if has == false {
		return false, errors.New("査不到数据")
	}

	if err != nil {
		return false, err
	}

	value, err := strconv.ParseFloat(item.Value, 32)

	if err != nil {
		return false, errors.New("最新数据为空")
	}

	if value > 37.5 {
		return true, nil
	}
	return false, nil
}

/*脉搏*/
type Pulse struct {
	BaseModel                     `xorm:"extends"`
	Testtime         fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene      string       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value            string       `json:"value" xorm:"notnull comment(值)"`
	Whetherbriefness string       `json:"whetherbriefness" xorm:"notnull comment(是否短促,0否1是)"`
}

func (v Pulse)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputPulse(session *xorm.Session,strData map[string]string) error {
	var item = &Pulse{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	if v, ok := strData["whetherbriefness"]; ok {
		item.Whetherbriefness = v
	} else {
		return errors.New("没有whetherbriefness")
	}

	return InsertPulse(session,item)
}

func InsertPulse(session *xorm.Session,item *Pulse) error{
	var item1 = &Pulse{}

	has,err := session.Table("Pulse").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutPulse(sql string, msg ...interface{}) ([]Pulse, error) {
	items := make([]Pulse, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*呼吸*/
type Breathe struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1辅助呼吸,2停辅助呼吸)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
	Whethertbm  string       `json:"whethertbm" xorm:"notnull comment(是否上呼吸机,0否1是)"`
}

func (v Breathe)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputBreathe(session *xorm.Session,strData map[string]string) error {
	var item = &Breathe{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	if v, ok := strData["whethertbm"]; ok {
		item.Whethertbm = v
	} else {
		return errors.New("没有whetherbriefness")
	}

	return InsertBreathe(session,item)
}

func InsertBreathe(session *xorm.Session,item *Breathe) error{
	var item1 = &Breathe{}

	has,err := session.Table("Breathe").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutBreathe(sql string, msg ...interface{}) ([]Breathe, error) {
	items := make([]Breathe, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*血压*/
type Pressure struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1外出,2检查,3请假,4拒试,5无法侧,6未测)"`
	Diavalue    string       `json:"diavalue" xorm:"notnull comment(低压值)"`
	Sysvalue    string       `json:"sysvalue" xorm:"notnull comment(高压值)"`
	//Pulsevalue  string       `json:"pulsevalue" xorm:"notnull comment(脉率值)"`
}

func (v Pressure)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputPressure(session *xorm.Session,strData map[string]string) error {
	var item = &Pressure{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["diavalue"]; ok {
		item.Diavalue = v
	} else {
		return errors.New("没有value")
	}

	if v, ok := strData["sysvalue"]; ok {
		item.Sysvalue = v
	} else {
		return errors.New("没有value")
	}

	return InsertPressure(session,item)
}

func InsertPressure(session *xorm.Session,item *Pressure) error{
	var item1 = &Pressure{}

	has,err := session.Table("Pressure").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutPressure(sql string, msg ...interface{}) ([]Pressure, error) {
	items := make([]Pressure, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*心率*/
type Heartrate struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Heartrate)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputHeartrate(session *xorm.Session,strData map[string]string) error {
	var item = &Heartrate{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertHeartrate(session,item)
}

func InsertHeartrate(session *xorm.Session,item *Heartrate) error{
	var item1 = &Heartrate{}

	has,err := session.Table("Heartrate").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutHeartrate(sql string, msg ...interface{}) ([]Heartrate, error) {
	items := make([]Heartrate, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*血氧*/
type Spo2h struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Spo2h)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputSpo2h(session *xorm.Session,strData map[string]string) error {
	var item = &Spo2h{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertSpo2h(session,item)
}

func InsertSpo2h(session *xorm.Session,item *Spo2h) error{
	var item1 = &Spo2h{}

	has,err := session.Table("Spo2h").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutSpo2h(sql string, msg ...interface{}) ([]Spo2h, error) {
	items := make([]Spo2h, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*血糖*/
type Glucose struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1外出,2检查,3请假,4拒试,5无法侧,6不在)"`
	Teststate   string       `json:"teststate" xorm:"notnull comment(测试的状态,1空腹,2早餐后1h,3早餐后2h,4中餐前,5中餐后1h,6中餐后2h,7晚餐前,8晚餐后1h,9晚餐后2h,10睡前)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Glucose)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputGlucose(session *xorm.Session,strData map[string]string) error {
	var item = &Glucose{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["teststate"]; ok {
		item.Teststate = v
	} else {
		return errors.New("teststate")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertGlucose(session,item)
}

func InsertGlucose(session *xorm.Session,item *Glucose) error{
	var item1 = &Glucose{}

	has,err := session.Table("Glucose").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutGlucose(sql string, msg ...interface{}) ([]Glucose, error) {
	items := make([]Glucose, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*体重*/
type Weight struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1卧床,2轮椅,3平车)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Weight)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputWeight(session *xorm.Session,strData map[string]string) error {
	var item = &Weight{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertWeight(session,item)
}

func InsertWeight(session *xorm.Session,item *Weight) error{
	var item1 = &Weight{}

	has,err := session.Table("Weight").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutWeight(sql string, msg ...interface{}) ([]Weight, error) {
	items := make([]Weight, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*身高*/
type Height struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1卧床,2轮椅,3平车)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Height)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputHeight(session *xorm.Session,strData map[string]string) error {
	var item = &Height{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertHeight(session,item)
}

func InsertHeight(session *xorm.Session,item *Height) error{
	var item1 = &Height{}

	has,err := session.Table("Height").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutHeight(sql string, msg ...interface{}) ([]Height, error) {
	items := make([]Height, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*皮试*/
type Skin struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Skin)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputSkin(session *xorm.Session,strData map[string]string) error {
	var item = &Skin{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertSkin(session,item)
}

func InsertSkin(session *xorm.Session,item *Skin) error{
	var item1 = &Skin{}

	has,err := session.Table("Skin").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutSkin(sql string, msg ...interface{}) ([]Skin, error) {
	items := make([]Skin, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*其他*/
type Ache struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Ache)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputAche(session *xorm.Session,strData map[string]string) error {
	var item = &Ache{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertAche(session,item)
}

func InsertAche(session *xorm.Session,item *Ache) error{
	var item1 = &Ache{}

	has,err := session.Table("Ache").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutAche(sql string, msg ...interface{}) ([]Ache, error) {
	items := make([]Ache, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*事件*/
type Incident struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene string       `json:"recordscene" xorm:"notnull comment(测试的场景,1入院,2出院,3手术,4分娩,5出生,6转入,7转科,8转院,9死亡,10外出)"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
}

func (v Incident)ShowTime()time.Time{
	return time.Time(v.Testtime)
}

func IputIncident(session *xorm.Session,strData map[string]string) error {
	var item = &Incident{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		item.NurseName = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["recordscene"]; ok {
		item.Recordscene = v
	} else {
		return errors.New("没有recordscene")
	}

	if v, ok := strData["value"]; ok {
		item.Value = v
	} else {
		return errors.New("没有value")
	}

	return InsertIncident(session,item)
}

func InsertIncident(session *xorm.Session,item *Incident) error{
	var item1 = &Incident{}

	has,err := session.Table("Incident").Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Get(item1)

	if err !=nil{
		return err
	}

	if has {
		_, err = session.Where("Testtime = ? and PatientId = ?",item.Testtime.String(),item.PatientId).Update(item);
	}else{
		_, err = session.Insert(item);
	}

	return err
}

func OutIncident(sql string, msg ...interface{}) ([]Incident, error) {
	items := make([]Incident, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

type TimeShowSlice []TimeShow

func (s TimeShowSlice) Len() int {
	return len(s)
	}

func (s TimeShowSlice) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
	}

func (s TimeShowSlice) Less(i, j int) bool {
	return s[i].ShowTime().Unix() < s[j].ShowTime().Unix()
	}


