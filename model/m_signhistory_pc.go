package model

import (
	"fit"
)

type PatientInfoDup struct {
	VAA01      	int64   `json:"patient_id"`
	BCQ04      	string  `json:"bed_coding"`    // 床号 BCQ1表
	VAA05    	string  `json:"patient_name"`  // 病人名字
	VAA101     	string  `json:"patient_age"`   // 病人年龄
}

type SignHistory1 struct{
	PatientInfoDup  `xorm:"extends"`

	TextTime        string  `json:"text_time"`         // 测试时间
	ThmValue       string  `json:"thm_value"`          // 体温值
	ThmType        string  `json:"thm_type"`           // 体温类型
	PulseValue     string  `json:"pulse_value"`        // 脉搏值
	PulseBriefness string  `json:"pulse_briefness"`    // 是否短促
	BreatheValue   string  `json:"breathe_value"`      // 呼吸值
	BreatheScene   string  `json:"breathe_scene"`      //呼吸场景
	PressureDia    string  `json:"pressure_dia"`       //高压
	PressureSys    string  `json:"pressure_sys"`       //低压
	PressureScene  string  `json:"pressure_scene"`     //血压场景
	HeartrateValue string  `json:"heartrate_value"`    //心率值
	Spo2hValue     string  `json:"spo2h_value"`        //血氧值
	GlucoseValue   string  `json:"glucose_value"`      //血糖值
	WeightValue    string  `json:"weight_value"`       //体重值
	WeightScene    string  `json:"weight_scene"`       //体重场景
	HeightValue    string  `json:"height_value"`       //身高值
	HeightScene    string  `json:"height_scene"`       //身高场景
	SkinValue      string  `json:"skin_value"`         //皮试值
	IncidentScene  string  `json:"incident_scene"`     //事件场景
}

func FetchPatientInfoForSignHistory(id string) (PatientInfoDup, error) {
	patient := PatientInfoDup{}
	_, err := fit.MySqlEngine().SQL("select VAA01, BCQ04, VAA05, VAA10 from VAA1 where VAA01 = ?",id).Get(&patient)
	return patient, err
}


