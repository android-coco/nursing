package handler

import (
	"fit"
	"nursing/model"
	/*"strconv"
	"time"*/
	"encoding/json"
)

type SignsiputController struct {
	fit.Controller
}

func (c SignsiputController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
}

func (c SignsiputController) Post(w *fit.Response, r *fit.Request, p fit.Params) {

	strData := r.FormValue("strData")
	fit.Logger().LogAssert("strData", strData)
	if len(strData) == 0{
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数错误"
		c.JsonData.Datas = []interface{}{}
	}

	var jsmap map[string]map[string]string
	err := json.Unmarshal([]byte(strData), &jsmap)
	if err !=nil {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数错误1"
		c.JsonData.Datas = []interface{}{err}
	}else{
		c.JsonData.Result = 0
		Datas := []interface{}{}

		for key,str := range jsmap{
			switch key {
			case model.Temperature_Type :
				err := model.IputTemperature(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				 }
			case model.Pulse_Type :
				err := model.IputPulse(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Breathe_Type :
				err := model.IputBreathe(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Pressure_Type :
				err := model.IputPressure(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Heartrate_Type :
				err := model.IputHeartrate(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Spo2h_Type:
				err := model.IputSpo2h(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Glucose_Type:
				err := model.IputGlucose(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Weight_Type:
				err := model.IputWeight(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			case model.Height_Type:
				err := model.IputHeight(str)
				if(err!=nil){
					c.JsonData.Result = 1
					Datas = append(Datas, err)
				}
			default:
				c.JsonData.Result = 1
				c.JsonData.ErrorMsg = "参数错误2"
				c.JsonData.Datas = []interface{}{}
			}
		}
		c.JsonData.Datas = Datas
	}

	c.ResponseToJson(w)
}
