package handler

import (
	"fit"
	"nursing/model"
	"encoding/json"
	"time"
	"errors"
	"github.com/go-xorm/xorm"
)

type PCBatvhinputController struct{
	PCController
}

func (c PCBatvhinputController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadViewSafely(w, r, "pc/v_batchInput.html", "pc/header_side.html", "pc/header_top.html")

		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}
		Data["Patients"] = response
		fit.Logger().LogError("gk dd", len(response))
		Data["Userinfo"] = userinfo
		Data["Menuindex"] = "4-0"
		c.Data = Data
	}
}

func (c PCBatvhinputController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	inputs := r.FormValue("value")
    if inputs == ""{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		return
	}

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

	err := session.Begin()
	if err!=nil{
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "事务开始失败"
		return
	}

	var maps []map[string]string
	err_js := json.Unmarshal([]byte(inputs), &maps)
	if err_js !=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "格式错误1"
		c.JsonData.Datas = err_js
		return
	}else{
		for _,str := range maps{
			err_as := BatchAnalysis(session,str)
			if err_as !=nil {
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "解析错误1"
				c.JsonData.Datas = err_as
				return
			}
		}
	}

	err_com := session.Commit()
	if err_com != nil {
		c.JsonData.Result = 4
		c.JsonData.ErrorMsg = "数据库插入失败"
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "录入成功"
	}
}

func BatchAnalysis(session *xorm.Session,strData map[string]string) error{

	var nurse_id       string
	var nurse_name     string
	var patient_id     string
	var testtime       fit.JsonTime

	var thm_value      string    //体温
	var thm_type       string

	var pulse_value    string    //脉搏
	var pulse_briefness string

	var breathe_value  string    //呼吸
	var breathe_scene  string

	var pressure_dia   string    //血压
	var pressure_sys   string
	var pressure_scene string

	var heartrate_value string   //心率

	var spo2h_value    string    //血氧

	var glucose_value  string    //血糖

	var weight_value   string    //体重
	var weight_scene   string

	var height_value   string    //身高
	var height_scene   string

	var skin_value     string    //皮试

	var incident_scene string   //事件

	//var intake_value   string   //入量
	//var intake_type    string   //入量类型



	if v, ok := strData["nurse_id"]; ok {
		nurse_id = v
	} else {
		return errors.New("没有nurse_id")
	}

	if v, ok := strData["nurse_name"]; ok {
		nurse_name = v
	} else {
		return errors.New("nurse_name")
	}

	if v, ok := strData["patient_id"]; ok {
		patient_id = v
	} else {
		return errors.New("没有patient_id")
	}

	if v, ok := strData["testtime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05",v,time.Local)
		if err != nil {
			return errors.New("没有testtime")
		} else {
			testtime = fit.JsonTime(texttime)
		}
	} else {
		return errors.New("没有testtime")
	}

	if v, ok := strData["thm_value"]; ok {
		thm_value = v
	} else {
		return errors.New("thm_value")
	}

	if v, ok := strData["thm_type"]; ok {
		thm_type = v
	} else {
		return errors.New("thm_type")
	}

	if v, ok := strData["pulse_value"]; ok {
		pulse_value = v
	} else {
		return errors.New("pulse_value")
	}

	if v, ok := strData["pulse_briefness"]; ok {
		pulse_briefness = v
	} else {
		return errors.New("pulse_briefness")
	}

	if v, ok := strData["breathe_value"]; ok {
		breathe_value = v
	} else {
		return errors.New("breathe_value")
	}

	if v, ok := strData["breathe_scene"]; ok {
		breathe_scene = v
	} else {
		return errors.New("breathe_scene")
	}

	if v, ok := strData["pressure_dia"]; ok {
		pressure_dia = v
	} else {
		return errors.New("pressure_dia")
	}

	if v, ok := strData["pressure_sys"]; ok {
		pressure_sys = v
	} else {
		return errors.New("pressure_sys")
	}

	if v, ok := strData["pressure_scene"]; ok {
		pressure_scene = v
	} else {
		return errors.New("pressure_scene")
	}

	if v, ok := strData["heartrate_value"]; ok {
		heartrate_value = v
	} else {
		return errors.New("heartrate_value")
	}

	if v, ok := strData["spo2h_value"]; ok {
		spo2h_value = v
	} else {
		return errors.New("spo2h_value")
	}

	if v, ok := strData["glucose_value"]; ok {
		glucose_value = v
	} else {
		return errors.New("glucose_value")
	}

	if v, ok := strData["weight_value"]; ok {
		weight_value = v
	} else {
		return errors.New("weight_value")
	}

	if v, ok := strData["weight_scene"]; ok {
		weight_scene = v
	} else {
		return errors.New("weight_scene")
	}

	if v, ok := strData["height_value"]; ok {
		height_value = v
	} else {
		return errors.New("height_value")
	}

	if v, ok := strData["height_scene"]; ok {
		height_scene = v
	} else {
		return errors.New("height_scene")
	}

	if v, ok := strData["skin_value"]; ok {
		skin_value = v
	} else {
		return errors.New("skin_value")
	}

	if v, ok := strData["incident_scene"]; ok {
		incident_scene = v
	} else {
		return errors.New("incident_scene")
	}

	if nurse_id =="" || nurse_name == "" || patient_id == "" {
		return errors.New("参数不完整")
	}


	var tem_item = &model.Temperature{}
	tem_item.NurseId = nurse_id
	tem_item.NurseName = nurse_name
	tem_item.PatientId = patient_id
	tem_item.Testtime = testtime
	tem_item.Value    = thm_value
	tem_item.Ttemptype = thm_type

	if thm_value != ""{
		err := model.InsertTemperature(session,tem_item);
		if err != nil{
			return err
		}
	}

	var pulse_item = &model.Pulse{}
	pulse_item.NurseId = nurse_id
	pulse_item.NurseName = nurse_name
	pulse_item.PatientId = patient_id
	pulse_item.Testtime = testtime
	pulse_item.Value    = pulse_value
	pulse_item.Whetherbriefness = pulse_briefness

	if pulse_value!= ""{
		err := model.InsertPulse(session,pulse_item);
		if err != nil{
			return err
		}
	}

	var breathe_item = &model.Breathe{}
	breathe_item.NurseId = nurse_id
	breathe_item.NurseName = nurse_name
	breathe_item.PatientId = patient_id
	breathe_item.Testtime = testtime
	breathe_item.Value    = breathe_value
	breathe_item.Recordscene = breathe_scene

	if breathe_value!= ""{
		err := model.InsertBreathe(session,breathe_item);
		if err != nil{
			return err
		}
	}

	var pressure_item = &model.Pressure{}
	pressure_item.NurseId = nurse_id
	pressure_item.NurseName = nurse_name
	pressure_item.PatientId = patient_id
	pressure_item.Testtime = testtime
	pressure_item.Diavalue = pressure_dia
	pressure_item.Sysvalue = pressure_sys
	pressure_item.Recordscene = pressure_scene

	if pressure_dia!= "" || pressure_sys !=""{
		err := model.InsertPressure(session,pressure_item);
		if err != nil{
			return err
		}
	}

	var heartrate_item = &model.Heartrate{}
	heartrate_item.NurseId = nurse_id
	heartrate_item.NurseName = nurse_name
	heartrate_item.PatientId = patient_id
	heartrate_item.Testtime = testtime
	heartrate_item.Value =  heartrate_value

	if heartrate_value!= "" {
		err := model.InsertHeartrate(session,heartrate_item);
		if err != nil{
			return err
		}
	}

	var spo2h_item = &model.Spo2h{}
	spo2h_item.NurseId = nurse_id
	spo2h_item.NurseName = nurse_name
	spo2h_item.PatientId = patient_id
	spo2h_item.Testtime = testtime
	spo2h_item.Value = spo2h_value

	if spo2h_value!= "" {
		err := model.InsertSpo2h(session,spo2h_item);
		if err != nil{
			return err
		}
	}

	var glucose_item = &model.Glucose{}
	glucose_item.NurseId = nurse_id
	glucose_item.NurseName = nurse_name
	glucose_item.PatientId = patient_id
	glucose_item.Testtime = testtime
	glucose_item.Value = glucose_value

	if glucose_value!= "" {
		err := model.InsertGlucose(session,glucose_item);
		if err != nil{
			return err
		}
	}

	var weight_item = &model.Weight{}
	weight_item.NurseId = nurse_id
	weight_item.NurseName = nurse_name
	weight_item.PatientId = patient_id
	weight_item.Testtime = testtime
	weight_item.Value = weight_value
	weight_item.Recordscene = weight_scene

	if weight_value!= "" {
		err := model.InsertWeight(session,weight_item);
		if err != nil{
			return err
		}
	}

	var height_item = &model.Height{}
	height_item.NurseId = nurse_id
	height_item.NurseName = nurse_name
	height_item.PatientId = patient_id
	height_item.Testtime = testtime
	height_item.Value = height_value
	height_item.Recordscene = height_scene

	if height_value!= "" {
		err := model.InsertHeight(session,height_item);
		if err != nil{
			return err
		}
	}

	var skin_item = &model.Skin{}
	skin_item.NurseId = nurse_id
	skin_item.NurseName = nurse_name
	skin_item.PatientId = patient_id
	skin_item.Testtime = testtime
	skin_item.Value = skin_value

	if skin_value!= "" {
		err := model.InsertSkin(session,skin_item);
		if err != nil{
			return err
		}
	}

	var incident_item = &model.Incident{}
	incident_item.NurseId = nurse_id
	incident_item.NurseName = nurse_name
	incident_item.PatientId = patient_id
	incident_item.Testtime = testtime
	incident_item.Recordscene = incident_scene

	if incident_scene!= "" {
		err := model.InsertIncident(session,incident_item);
		if err != nil{
			return err
		}
	}

	return nil
}

