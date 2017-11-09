package handler

import (
	"fit"
	"nursing/model"
	"errors"
	"fmt"
)

// 腕带打印
type PCWristStrapController struct {
	fit.Controller
}

// 腕带打印
type WristStrap struct {
	Beds       string //床位号
	Name       string // 姓名
	Sex        string //性别
	Age        string //年龄
	ClassName  string //科室
	Pid        string //病人ID
	HospitalId string //住院ID
}

// 腕带打印
func (c PCWristStrapController) Get(w *fit.Response, r *fit.Request, p fit.Params) {

	r.ParseForm()
	pid := r.FormValue("pid")
	patients, err := model.GetPatientInfo(pid)
	if len(patients) == 0 || err != nil {
		fit.Logger().LogError("yh", err, errors.New("patient slice is empty"))
		return
	}
	var sex = "未知"
	if patients[0].ABW01 == "1" || patients[0].ABW01 == "M" {
		sex = "男"
	} else if patients[0].ABW01 == "2" || patients[0].ABW01 == "F" {
		sex = "女"
	} else {
		sex = "未知"
	}
	c.Data = fit.Data{"Datas": WristStrap{patients[0].BCQ04, patients[0].VAA05, sex, fmt.Sprint(patients[0].VAA10), patients[0].BCK03, pid, patients[0].VAA04}}
	c.LoadView(w, "pc/v_wdprint.html")
}

// 瓶贴打印
type PCBottleStrapController struct {
	fit.Controller
}

// 瓶贴打印
type BottleStrap struct {
	Beds       string
	Name       string
	Sex        string
	Age        string
	ClassName  string
	Pid        string
	HospitalId string
}

// 瓶贴打印
func (c PCBottleStrapController) Get(w *fit.Response, r *fit.Request, p fit.Params) {

	//r.ParseForm()
	//pid := r.FormValue("pid")
	//patients, err := model.QueryPatientInfo(pid)
	//if len(patients) == 0 || err != nil {
	//	fit.Logger().LogError("yh", err, errors.New("patient slice is empty"))
	//	return
	//}
	//var sex = "未知"
	//if patients[0].ABW01 == "1" || patients[0].ABW01 == "M" {
	//	sex = "男"
	//} else if patients[0].ABW01 == "2" || patients[0].ABW01 == "F" {
	//	sex = "女"
	//} else {
	//	sex = "未知"
	//}
	//c.Data = fit.Data{"Datas": WristStrap{patients[0].BCQ04, patients[0].VAA05, sex, fmt.Sprint(patients[0].VAA10), patients[0].BCK03, pid, patients[0].VAA04}}
	c.LoadView(w, "pc/v_ptprint.html")
}
