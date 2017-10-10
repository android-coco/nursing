package model

import (
	"fit"
	"errors"
	"time"
	"strconv"
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
)

/*体温*/
type Temperature struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Ttemptype   uint16       `json:"ttemptype" xorm:"notnull comment(体温的类型)"`
	Coolingtype uint16       `json:"coolingtype" xorm:"notnull comment(降温的类型)"`
	Value       float32      `json:"value" xorm:"notnull comment(值)"`
}

func IputTemperature(strData map[string]string) error {
	var item = &Temperature{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["ttemptype"]; ok {
		ttemptype, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("ttemptype")
		} else {
			item.Recordscene = uint16(ttemptype)
		}
	} else {
		return errors.New("ttemptype")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["coolingtype"]; ok {
		coolingtype, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有coolingtype")
		} else {
			item.Ttemptype = uint16(coolingtype)
		}
	} else {
		return errors.New("没有coolingtype")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = float32(value)
		}
	} else {
		return errors.New("没有value")
	}

	_, err := fit.MySqlEngine().Insert(item);
	return err
}

func OutTemperature(sql string, msg ...interface{}) ([]Temperature, error) {
	items := make([]Temperature, 0)
	err := fit.MySqlEngine().Where(sql, msg...).Find(&items)
	return items, err
}

/*查询用户是否发热*/
func GetWhetherFever(PatientId string) (bool, error) {

	item := Temperature{}
	has, err := fit.MySqlEngine().Desc("id").Get(&item)

	if has == false {
		return false, errors.New("差不到数据")
	}

	if err != nil {
		return false, err
	}

	if item.Value > 38 {
		return true, nil
	}
	return false, nil
}

/*脉搏*/
type Pulse struct {
	BaseModel                     `xorm:"extends"`
	Testtime         fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene      uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value            uint16       `json:"value" xorm:"notnull comment(值)"`
	Whetherbriefness bool         `json:"whetherbriefness" xorm:"notnull comment(是否短促)"`
}

func IputPulse(strData map[string]string) error {
	var item = &Pulse{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}
	if v, ok := strData["whetherbriefness"]; ok {
		whetherbriefness, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有whetherbriefness")
		} else {
			if whetherbriefness == 0 {
				item.Whetherbriefness = false
			} else {
				item.Whetherbriefness = true
			}
		}
	} else {
		return errors.New("没有whetherbriefness")
	}

	err := item.InsertData(item);
	return err
}

/*呼吸*/
type Breathe struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       uint16       `json:"value" xorm:"notnull comment(值)"`
	Whethertbm  bool         `json:"whethertbm" xorm:"notnull comment(是否上呼吸机)"`
}

func IputBreathe(strData map[string]string) error {
	var item = &Breathe{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}
	if v, ok := strData["whethertbm"]; ok {
		whethertbm, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("whethertbm")
		} else {
			if whethertbm == 0 {
				item.Whethertbm = false
			} else {
				item.Whethertbm = true
			}
		}
	} else {
		return errors.New("没有whetherbriefness")
	}

	err := item.InsertData(item);
	return err
}

/*血压*/
type Pressure struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Diavalue    uint16       `json:"diavalue" xorm:"notnull comment(低压值)"`
	Sysvalue    uint16       `json:"sysvalue" xorm:"notnull comment(高压值)"`
	Pulsevalue  uint16       `json:"pulsevalue" xorm:"notnull comment(脉率值)"`
}

func IputPressure(strData map[string]string) error {
	var item = &Pressure{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["diavalue"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Diavalue = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}
	if v, ok := strData["sysvalue"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Sysvalue = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}
	if v, ok := strData["pulsevalue"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Pulsevalue = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}

	err := item.InsertData(item);
	return err
}

/*心率*/
type Heartrate struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       uint16       `json:"value" xorm:"notnull comment(值)"`
}

func IputHeartrate(strData map[string]string) error {
	var item = &Heartrate{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}

	err := item.InsertData(item);
	return err
}

/*血氧*/
type Spo2h struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       uint16       `json:"value" xorm:"notnull comment(值)"`
}

func IputSpo2h(strData map[string]string) error {
	var item = &Spo2h{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = uint16(value)
		}
	} else {
		return errors.New("没有value")
	}

	err := item.InsertData(item);
	return err
}

/*血糖*/
type Glucose struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       float32      `json:"value" xorm:"notnull comment(值)"`
}

func IputGlucose(strData map[string]string) error {
	var item = &Glucose{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = float32(value)
		}
	} else {
		return errors.New("没有value")
	}

	err := item.InsertData(item);
	return err
}

/*体重*/
type Weight struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       float32      `json:"value" xorm:"notnull comment(值)"`
}

func IputWeight(strData map[string]string) error {
	var item = &Weight{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = float32(value)
		}
	} else {
		return errors.New("没有value")
	}

	err := item.InsertData(item);
	return err
}

/*身高*/
type Height struct {
	BaseModel                `xorm:"extends"`
	Testtime    fit.JsonTime `json:"testtime" xorm:"notnull comment(测试时间)"`
	Recordscene uint16       `json:"recordscene" xorm:"notnull comment(测试的场景)"`
	Value       float64      `json:"value" xorm:"notnull comment(值)"`
}

func IputHeight(strData map[string]string) error {
	var item = &Height{}

	if v, ok := strData["nurse_id"]; ok {
		item.NurseId = v
	} else {
		return errors.New("没有nurse_id")
	}
	if v, ok := strData["patient_id"]; ok {
		item.PatientId = v
	} else {
		return errors.New("没有patient_id")
	}
	if v, ok := strData["testtime"]; ok {
		texttime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			item.Testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}
	if v, ok := strData["recordscene"]; ok {
		recordscene, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return errors.New("没有recordscene")
		} else {
			item.Recordscene = uint16(recordscene)
		}
	} else {
		return errors.New("没有recordscene")
	}
	if v, ok := strData["value"]; ok {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return errors.New("没有value")
		} else {
			item.Value = value
		}
	} else {
		return errors.New("没有value")
	}

	err := item.InsertData(item);
	return err
}
