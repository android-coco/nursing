package handler

import (
	"fit"
	"time"
	"nursing/model"
	"encoding/json"
	"fmt"
)

type PCSuccessController struct {
	PCController
}

func (c PCSuccessController) Get(w *fit.Response, r *fit.Request, p fit.Params) {

	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		defer c.LoadViewSafely(w, r, "pc/v_succession.html", "pc/header_side.html", "pc/header_top.html")

		starttime :=  r.FormValue("datatime")

		if starttime == "" {
			t := time.Now()
			starttime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
		}
		classid := userinfo.DepartmentID

		Data := make(fit.Data)

		response, err_rp := model.QueryDepartmentBeds(classid, false)
		if err_rp != nil {
			fit.Logger().LogError("gk", err_rp)
			return
		}
		Data["Patients"] = response
		//fit.Logger().LogError("gk dd", response ,classid)

		successions, err1 := model.OutSuccession("datatime = ? and classid = ?", starttime, classid)

		if (err1 != nil || len(successions) == 0 ) {
			fit.Logger().LogError("gk", starttime, len(successions), err1)
		} else {
			Data["d_Disabled"] = false
			Data["n_Disabled"] = false
			Data["N_Disabled"] = false

			for _, k := range successions {
				if k.Type == 1 {
					Data["d_NoldPatient"] = k.NoldPatient
					Data["d_NnowPatient"] = k.NnowPatient
					Data["d_NintHospital"] = k.NintHospital
					Data["d_NoutHospital"] = k.NoutHospital
					Data["d_Ninto"] = k.Ninto
					Data["d_Nout"] = k.Nout
					Data["d_Nsurgery"] = k.Nsurgery
					Data["d_Nchildbirth"] = k.Nchildbirth
					Data["d_Ncritically"] = k.Ncritically
					Data["d_Ndeath"] = k.Ndeath
					Data["d_NintensiveCare"] = k.NintensiveCare
					Data["d_NprimaryCare"] = k.NprimaryCare
					Data["d_NursingName"] = k.NursingName

					if userinfo.UID == k.NursingID {
						Data["d_Disabled"] = false
					}else{
						Data["d_Disabled"] = true
					}
				} else if k.Type == 2 {
					Data["n_NoldPatient"] = k.NoldPatient
					Data["n_NnowPatient"] = k.NnowPatient
					Data["n_NintHospital"] = k.NintHospital
					Data["n_NoutHospital"] = k.NoutHospital
					Data["n_Ninto"] = k.Ninto
					Data["n_Nout"] = k.Nout
					Data["n_Nsurgery"] = k.Nsurgery
					Data["n_Nchildbirth"] = k.Nchildbirth
					Data["n_Ncritically"] = k.Ncritically
					Data["n_Ndeath"] = k.Ndeath
					Data["n_NintensiveCare"] = k.NintensiveCare
					Data["n_NprimaryCare"] = k.NprimaryCare
					Data["n_NursingName"] = k.NursingName

					if userinfo.UID == k.NursingID {
						Data["n_Disabled"] = false
					}else{
						Data["n_Disabled"] = true
					}
				} else {
					Data["N_NoldPatient"] = k.NoldPatient
					Data["N_NnowPatient"] = k.NnowPatient
					Data["N_NintHospital"] = k.NintHospital
					Data["N_NoutHospital"] = k.NoutHospital
					Data["N_Ninto"] = k.Ninto
					Data["N_Nout"] = k.Nout
					Data["N_Nsurgery"] = k.Nsurgery
					Data["N_Nchildbirth"] = k.Nchildbirth
					Data["N_Ncritically"] = k.Ncritically
					Data["N_Ndeath"] = k.Ndeath
					Data["N_NintensiveCare"] = k.NintensiveCare
					Data["N_NprimaryCare"] = k.NprimaryCare

					Data["N_NursingName"] = k.NursingName

					if userinfo.UID == k.NursingID {
						Data["N_Disabled"] = false
					}else{
						Data["N_Disabled"] = true
					}
				}
			}

			successiondetails, _ := model.OutSuccessionDetails("datatime = ? and classid = ?", starttime, classid)
			if len(successiondetails) != 0 {
				Data["Successiondetails"] = successiondetails
			}
			fit.Logger().LogError("gk succ", len(successiondetails))
		}
		Data["Userinfo"] = userinfo
		Data["Menuindex"] = "6-0"
		c.Data = Data
	}
}

func (c PCSuccessController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	session := fit.MySqlEngine().NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "事务开始失败"
		return
	}

	dayshift := r.FormValue(model.Day_Shift)
	if len(dayshift) != 0 {
		var maps map[string]string
		err := json.Unmarshal([]byte(dayshift), &maps)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误1"
			c.JsonData.Datas = err
			return
		} else {
			err := model.IputSuccession(session, maps)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "参数错误1"
				c.JsonData.Datas = err
				return
			}
		}
	}

	nightshift := r.FormValue(model.Night_Shift)
	if len(nightshift) != 0 {
		var maps map[string]string
		err := json.Unmarshal([]byte(nightshift), &maps)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误2"
			c.JsonData.Datas = err
			return
		} else {
			err := model.IputSuccession(session, maps)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "参数错误2"
				c.JsonData.Datas = err
				return
			}
		}
	}

	graveyardshift := r.FormValue(model.Graveyard_Shift)
	if len(graveyardshift) != 0 {
		var maps map[string]string
		err := json.Unmarshal([]byte(graveyardshift), &maps)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误3"
			c.JsonData.Datas = err
			return
		} else {
			err := model.IputSuccession(session, maps)
			if (err != nil) {
				session.Rollback()
				c.JsonData.Result = 2
				c.JsonData.ErrorMsg = "参数错误3"
				c.JsonData.Datas = err
				return
			}
		}
	}

	commentshift := r.FormValue(model.Comment_Shift)
	if len(commentshift) != 0 {
		var maps []map[string]string
		err := json.Unmarshal([]byte(commentshift), &maps)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误4"
			c.JsonData.Datas = err
			return
		} else {
			for _, str := range maps {
				err := model.IputSuccessionDetails(session, str)
				if (err != nil) {
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误4"
					c.JsonData.Datas = err
					fit.Logger().LogError("fffffff",commentshift,err)
					return
				}
			}
		}
	}


	delets := r.FormValue("Comment_Delet")
	fmt.Println(len(delets))
	if len(delets) != 0 {
		var maps []string
		err := json.Unmarshal([]byte(delets), &maps)
		if err != nil {
			c.JsonData.Result = 1
			c.JsonData.ErrorMsg = "格式错误5"
			c.JsonData.Datas = err
			return
		} else {
			for _, str := range maps {
				err := model.DeletsSuccessionDetails(session, str)
				if (err != nil) {
					session.Rollback()
					c.JsonData.Result = 2
					c.JsonData.ErrorMsg = "参数错误5"
					c.JsonData.Datas = err
					fit.Logger().LogError("fffffff",commentshift,err)
					return
				}
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
