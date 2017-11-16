//  Created by JP

package handler

import (
	"fmt"
	"fit"
	"time"
	"nursing/model"
	"strconv"
)

type PCHistoryController struct {
	PCController
}

/*API 搜索历史病人*/
func (c PCHistoryController) SearchPatients(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	did := r.FormValue("did")
	if did == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	did_i, err_did := strconv.Atoi(did)
	if err_did != nil || did_i <= 0 {
		c.RenderingJsonAutomatically(2, "参数错误 did")
		return
	}
	staDate := r.FormValue("staDate")
	endDate := r.FormValue("endDate")
	if staDate != "" && endDate != "" {
		response, err := model.SearchHistoryParientsWithTimeInterval(staDate, endDate, did_i)
		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			if length := len(response); length == 0 {
				c.RenderingJsonAutomatically(4, "没有查找到该时间段出院的病人")
			} else {
				c.RenderingJson(0, "查找成功", response)
			}
		}
		return
	}
	rgtNum := r.FormValue("rgtNum")
	if rgtNum != "" {
		response, err := model.SearchHistoryParientsWithRegisterNum(rgtNum, did_i)
		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			if length := len(response); length == 0 {
				c.RenderingJsonAutomatically(4, "没有查找到该登记号对应的出院病人")
			} else {
				c.RenderingJson(0, "查找成功", response)
			}
		}
		return
	}
	hospNum := r.FormValue("hospNum")
	if hospNum != "" {
		response, err := model.SearchHistoryParientsWithHospitalizationNum(hospNum, did_i)
		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			if length := len(response); length == 0 {
				c.RenderingJsonAutomatically(4, "没有查找到该住院号对应的出院病人")
			} else {
				c.RenderingJson(0, "查找成功", response)
			}
		}
		return
	}
	name := r.FormValue("name")
	if name != "" {
		response, err := model.SearchHistoryParientsWithName(name, did_i)
		if err != nil {
			c.RenderingJsonAutomatically(3, "Database "+err.Error())
		} else {
			if length := len(response); length == 0 {
				c.RenderingJsonAutomatically(4, "没有查找到该姓名对应的出院病人")
			} else {
				c.RenderingJson(0, "查找成功", response)
			}
		}
	}

}

func (c *PCHistoryController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *PCHistoryController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}

/*历史病人页面（医嘱）*/
func (c PCHistoryController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}

		timeNow := time.Now()
		startTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-2, 0, 0, 0, 0, timeNow.Location())
		patients, err := model.SearchHistoryParientsWithTimeInterval(startTime.Format("2006-01-02 15:04:05"), timeNow.Format("2006-01-02 15:04:05"), userinfo.DepartmentID)
		if err != nil {
			fit.Logger().LogError("JP *", err.Error())
			fmt.Fprintln(w, "网页好像出了点故障"+err.Error())
		}
		c.Data["Patients"] = patients
		_ = c.LoadViewSafely(w, r, "pc/v_history_advice1.html", "pc/t_history_side.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*体温页面*/
func (c PCHistoryController) Temperature(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}

		timeNow := time.Now()
		startTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-2, 0, 0, 0, 0, timeNow.Location())
		patients, err := model.SearchHistoryParientsWithTimeInterval(startTime.Format("2006-01-02 15:04:05"), timeNow.Format("2006-01-02 15:04:05"), userinfo.DepartmentID)
		if err != nil {
			fit.Logger().LogError("JP *", err.Error())
			fmt.Fprintln(w, "网页好像出了点故障"+err.Error())
		}
		c.Data["Patients"] = patients
		_ = c.LoadViewSafely(w, r, "pc/v_history_temperature.html", "pc/t_history_side.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*声明体征页面*/
func (c PCHistoryController) Signs(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}

		timeNow := time.Now()
		startTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-2, 0, 0, 0, 0, timeNow.Location())
		patients, err := model.SearchHistoryParientsWithTimeInterval(startTime.Format("2006-01-02 15:04:05"), timeNow.Format("2006-01-02 15:04:05"), userinfo.DepartmentID)
		if err != nil {
			fit.Logger().LogError("JP *", err.Error())
			fmt.Fprintln(w, "网页好像出了点故障"+err.Error())
		}
		c.Data["Patients"] = patients
		_ = c.LoadViewSafely(w, r, "pc/v_history_signs.html", "pc/t_history_side.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*医嘱页面*/
func (c PCHistoryController) Advice(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		c.Data = fit.Data{
			"Userinfo": userinfo,
		}

		timeNow := time.Now()
		startTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-2, 0, 0, 0, 0, timeNow.Location())
		patients, err := model.SearchHistoryParientsWithTimeInterval(startTime.Format("2006-01-02 15:04:05"), timeNow.Format("2006-01-02 15:04:05"), userinfo.DepartmentID)
		if err != nil {
			fit.Logger().LogError("JP *", err.Error())
			fmt.Fprintln(w, "网页好像出了点故障"+err.Error())
		}
		c.Data["Patients"] = patients
		_ = c.LoadViewSafely(w, r, "pc/v_history_advice2.html", "pc/t_history_side.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}
