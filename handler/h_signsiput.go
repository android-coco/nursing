package handler

import (
	"fit"
	"nursing/model"
	"encoding/json"
)

type SignsiputController struct {
	fit.Controller
}

func (c SignsiputController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

	err := session.Begin()
	if err!=nil{
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "事务开始失败"
		return
	}

	temperature := r.FormValue(model.Temperature_Type )
	if len(temperature) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(temperature), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误1"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputTemperature(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误1"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	pulse := r.FormValue(model.Pulse_Type )
	if len(pulse) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(pulse), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误2"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputPulse(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误2"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	breathe := r.FormValue(model.Breathe_Type)
	if len(breathe) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(breathe), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误3"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputBreathe(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误3"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	pressure := r.FormValue(model.Pressure_Type)
	if len(pressure) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(pressure), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误4"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputPressure(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误4"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	heartrate := r.FormValue(model.Heartrate_Type)
	if len(heartrate) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(heartrate), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误5"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputHeartrate(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误5"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	spo2h := r.FormValue(model.Spo2h_Type)
	if len(spo2h) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(spo2h), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误6"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputSpo2h(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误6"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	glucose := r.FormValue(model.Glucose_Type)
	if len(glucose) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(glucose), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误7"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputGlucose(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误7"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	weight := r.FormValue(model.Weight_Type)
	if len(weight) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(weight), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误8"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputWeight(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误8"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	height := r.FormValue(model.Height_Type)
	if len(height) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(height), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误9"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputHeight(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误9"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	skin := r.FormValue(model.Skin_Type)
	if len(skin) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(skin), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误9"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputSkin(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误10"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	ache := r.FormValue(model.Ache_Type)
	if len(ache) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(ache), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误9"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputAche(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误11"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	incident := r.FormValue(model.Incident_Type)
	if len(incident) != 0{
		var maps []map[string]string
		err := json.Unmarshal([]byte(incident), &maps)
		if err !=nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误9"
			c.JsonData.Datas = err
			return
		}else{
			for _,str := range maps{
				err := model.IputIncident(session,str)
				if(err!=nil){
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误12"
					c.JsonData.Datas = err
					return
				}
			}
		}
	}

	err_com := session.Commit()
	if err_com != nil {
		c.JsonData.Result = 4
		c.JsonData.ErrorMsg = "数据库插入失败"
	}else{
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "录入成功"
	}



}
