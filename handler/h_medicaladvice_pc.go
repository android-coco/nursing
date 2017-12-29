package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"time"
	"strconv"
	"nursing/utils"
	"encoding/json"
	"strings"
)

type PCMedicalAdviceController struct {
	PCController
}

/*医嘱查询 page*/
func (c PCMedicalAdviceController) PCSearch(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		patients := model.FetchInpatientWardPatients(userinfo.DepartmentID)

		//pidStr := ""
		//var index int
		//length := len(patients)
		//for index = 0; index < length; index ++ {
		//	p := patients[index]
		//	if v := length - 1; index < v {
		//		pidStr = fmt.Sprintf("%s%d,", pidStr, p.Vid)
		//	} else {
		//		pidStr = fmt.Sprintf("%s%d", pidStr, p.Vid)
		//	}
		//}
		//mAdvices, _ := model.SearchMedicalAdviceForPC(0, 0, "0", pidStr, "all", "all")
		c.Data = fit.Data{
			"Userinfo":  userinfo,
			"Patients":  patients,
			//"MAdvices":  mAdvices,
			"Menuindex": "3-0",
		}

		_ = c.LoadViewSafely(w, r, "pc/v_medicaladvice_message.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*医嘱查询 API*/
func (c PCMedicalAdviceController) Search(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	patients := r.FormValue("patients")
	st := r.FormValue("starttime")
	et := r.FormValue("endtime")
	category := r.FormValue("category")
	typeOf := r.FormValue("type")
	status := r.FormValue("status")

	if patients == "" || st == "" || et == "" || category == "" || typeOf == "" || status == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	if st != "all" && et != "all" {
		if _, err := time.Parse("2006-01-02 15:04:05", st); err != nil {
			c.RenderingJsonAutomatically(2, "参数错误 starttime")
			return
		} else if _, err := time.Parse("2006-01-02 15:04:05", et); err != nil {
			c.RenderingJsonAutomatically(2, "参数错误 endtime")
			return
		}
	}

	type_i, err_t := strconv.Atoi(typeOf)
	status_i, err_s := strconv.Atoi(status)
	if err_t != nil {
		c.RenderingJsonAutomatically(2, "参数错误 type")
		return
	} else if err_s != nil {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	mAdvices, err_db := model.SearchMedicalAdviceForPC(type_i, status_i, category, patients, st, et)
	if err_db != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	} else {
		c.RenderingJson(0, "成功", mAdvices)
	}
}

/*医嘱执行明细 page*/
func (c PCMedicalAdviceController) PCExecState(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		patients := model.FetchInpatientWardPatients(userinfo.DepartmentID)
		//pidStr := ""
		//var index int
		//length := len(patients)
		//for index = 0; index < length; index ++ {
		//	p := patients[index]
		//	if v := length - 1; index < v {
		//		pidStr = fmt.Sprintf("%s%d,", pidStr, p.Vid)
		//	} else {
		//		pidStr = fmt.Sprintf("%s%d", pidStr, p.Vid)
		//	}
		//}
		//mAdvices, err_db := model.SearchMedicalAdviceExecutionForPC(0, 0, "0", pidStr, "all", "all")
		//if err_db != nil {
		//	fit.Logger().LogError("***JK***", err_db.Error())
		//}
		c.Data = fit.Data{
			"Userinfo":  userinfo,
			"Patients":  patients,
			//"MAdvices":  mAdvices,
			"Menuindex": "3-0",
		}
		_ = c.LoadViewSafely(w, r, "pc/v_medicaladvice_message2.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*医嘱执行明细 API*/
func (c PCMedicalAdviceController) PCExecSearch(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()

	patients := r.FormValue("patients")
	st := r.FormValue("starttime")
	et := r.FormValue("endtime")
	category := r.FormValue("category")
	typeOf := r.FormValue("type")
	status := r.FormValue("state")

	if patients == "" || st == "" || et == "" || category == "" || typeOf == "" || status == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}
	if st != "all" && et != "all" {
		if _, err := time.Parse("2006-01-02 15:04:05", st); err != nil {
			c.RenderingJsonAutomatically(2, "参数错误 starttime")
			return
		} else if _, err := time.Parse("2006-01-02 15:04:05", et); err != nil {
			c.RenderingJsonAutomatically(2, "参数错误 endtime")
			return
		}
	}

	type_i, err_t := strconv.Atoi(typeOf)
	status_i, err_s := strconv.Atoi(status)
	if err_t != nil {
		c.RenderingJsonAutomatically(2, "参数错误 type")
		return
	} else if err_s != nil {
		c.RenderingJsonAutomatically(2, "参数错误 state")
		return
	}

	mAdvices, err_db := model.SearchMedicalAdviceExecutionForPC(type_i, status_i, category, patients, st, et)
	if err_db != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_db.Error())
	} else {
		c.RenderingJson(0, "成功", mAdvices)
	}
}

/*医嘱执行记录 API*/
func (c PCMedicalAdviceController) PCExecDetail(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	madid := r.FormValue("madid")     // 医嘱ID
	extime := r.FormValue("extime")   // 计划执行时间
	excycle := r.FormValue("excycle") // 分组序号
	if madid == "" || extime == "" || excycle == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
	} else {
		mid, err_m := utils.Int64Value(madid)
		if err_m != nil {
			c.RenderingJsonAutomatically(2, "参数错误 madid"+madid)
		} else if _, err := time.Parse("2006-01-02 15:04:05", extime); err != nil {
			c.RenderingJsonAutomatically(2, "参数错误 extime"+extime)
		} else if _, err := strconv.Atoi(excycle); err != nil {
			c.RenderingJsonAutomatically(2, "参数错误 excycle"+excycle)
		} else {
			res, err := model.FetchMedicalAdviceExecutionRecordForPc(mid, excycle, extime)
			if err != nil {
				c.RenderingJsonAutomatically(3, "Database "+err.Error())
			} else {
				c.RenderingJson(0, "成功", res)
			}
		}
	}
}

/*医嘱拆分 page*/
func (c PCMedicalAdviceController) PCSplit(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err == nil {
		patients := model.FetchInpatientWardPatients(userinfo.DepartmentID)

		//pidStr := ""
		//var index int
		//length := len(patients)
		//for index = 0; index < length; index ++ {
		//	p := patients[index]
		//	if v := length - 1; index < v {
		//		pidStr = fmt.Sprintf("%s%d,", pidStr, p.Vid)
		//	} else {
		//		pidStr = fmt.Sprintf("%s%d", pidStr, p.Vid)
		//	}
		//}

		//t := time.Now()
		//st := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
		//et := time.Date(t.Year(), t.Month(), t.Day(), 23, 55, 59, 1, t.Location()).Format("2006-01-02 15:04:05")

		//var temp []model.MedicalAdviceResponse
		//response := make([]model.MedicalAdviceResponse, 0)
		//temp, _ = model.SearchSplitMedicalAdviceForInfusion(st, et, pidStr, 0, 2, userinfo.DepartmentID)

		//response = append(response, temp...)
		//temp, _ = model.SearchSplitMedicalAdviceForOralMedical(st, et, pidStr, 0, 2, userinfo.DepartmentID)
		//response = append(response, temp...)
		//temp, _ = model.SearchSplitMedicalAdviceForOralInjection(st, et, pidStr, 0, 2, userinfo.DepartmentID)
		//response = append(response, temp...)
		//mAdvices, _ := model.SearchMedicalAdvicesForSplitting(st, et, pidStr, "0,1,2,3", 0, 2, userinfo.DepartmentID)
		c.Data = fit.Data{
			"Userinfo":  userinfo,
			"Patients":  patients,
			//"MAdvices":  response,
			"Menuindex": "2-0",
		}
		_ = c.LoadViewSafely(w, r, "pc/v_medicaladvice_split.html", "pc/header_side.html", "pc/header_top.html")
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}

/*医嘱拆分查询 API*/
func (c PCMedicalAdviceController) SpiltSearch(w *fit.Response, r *fit.Request, p fit.Params) {
	//select a.VAF01,a.VAA01,v.VAA05,v.BCQ04,a.VAF10 状态,a.VAF11 长临,a.BDA01 类别,a.VAF19 用剂量,a.VAF21 数量,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF27 次数,a.VAF36 开始执行时间,c.BBX20 打印单,a.VAF60 滴速from ((VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01) left join VAA1 v on v.VAA01 = a.VAA01 where a.VAA01 in (select q.VAA01 from BCQ1 q where q.VAA01 = a.VAA01) and a.VAF32 = 0 and a.VAF42 BETWEEN '2017-11-08 00:00:00' AND '2017-11-08 23:59:59' and a.VAF11 = '1' order by a.VAF36
	defer c.ResponseToJson(w)
	r.ParseForm()

	patients := r.FormValue("patients")
	dt := r.FormValue("time")
	typeOf := r.FormValue("term")
	print := r.FormValue("print")
	did := r.FormValue("did")

	if patients == "" || dt == "" || typeOf == "" || print == "" || did == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	by := []byte(print)
	anyObj := make(map[string]string, 0)
	err_json := json.Unmarshal(by, &anyObj)
	if err_json != nil {
		fit.Logger().LogError("**JP**", err_json)
		c.RenderingJsonAutomatically(2, "参数错误 print"+print)
		return
	}

	t, err_t := time.Parse("2006-01-02", dt)
	if err_t != nil {
		c.RenderingJsonAutomatically(2, "参数错误 time")
		return
	}
	st := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
	et := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 1, t.Location()).Format("2006-01-02 15:04:05")

	type_i, err_t := strconv.Atoi(typeOf)
	did_i, err_d := strconv.Atoi(did)
	if err_t != nil {
		c.RenderingJsonAutomatically(2, "参数错误 term")
		return
	} else if err_d != nil {
		c.RenderingJsonAutomatically(2, "参数错误 did")
		return
	}

	printType := anyObj["type"]
	if printType != "excution" && printType != "label" && printType != "temporary" {
		c.RenderingJsonAutomatically(2, "参数错误 type")
		return
	}

	status := anyObj["status"]
	print_i, err_p := strconv.Atoi(status)
	if err_p != nil {
		c.RenderingJsonAutomatically(2, "参数错误 status")
		return
	}

	//printStr := anyObj["value"]
	prints := strings.Split(anyObj["value"], ",")

	var err_re error
	response := make([]model.MedicalAdviceResponse, 0)
	if printType == "excution" {
		for _, v := range prints {
			if v == "2" {
				// 输液单
				temp, err_db := model.SearchSplitMedicalAdviceForInfusion(st, et, patients, type_i, print_i, did_i)
				err_re = err_db
				response = append(response, temp...)
			} else if v == "0" {
				// 口服单
				temp, err_db := model.SearchSplitMedicalAdviceForOralMedical(st, et, patients, type_i, print_i, did_i)
				err_re = err_db
				response = append(response, temp...)
			} else if v == "1" {
				// 注射单
				temp, err_db := model.SearchSplitMedicalAdviceForOralInjection(st, et, patients, type_i, print_i, did_i)
				err_re = err_db
				response = append(response, temp...)
			} else if v == "3" {
				// 治疗单无数据
			}
		}
	} else if printType == "label" {
		for _, v := range prints {
			if v == "0" {
				temp, err_db := model.SearchSplitMedicalAdviceForBottlePost(st, et, patients, type_i, print_i)
				err_re = err_db
				response = append(response, temp...)
			} else {
				// 检验标签无数据
			}
		}
	} else  {
		// 口服+临嘱
		temp, err_db := model.SearchSplitMedicalAdviceForOralMedical(st, et, patients, type_i, print_i, did_i)
		err_re = err_db
		response = append(response, temp...)
	}
	//mAdvices, err_db := model.SearchMedicalAdvicesForSplitting(st, et, patients, printStr, type_i, print_i,did_i)

	if err_re != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_re.Error())
	} else {
		c.RenderingJson(0, "成功", response)
	}
}

/*医嘱打印*/
func (c PCMedicalAdviceController) PCPrint(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	jsonStr := r.FormValue("json")
	if jsonStr == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	var params []map[string]string
	err_js := json.Unmarshal([]byte(jsonStr), &params)
	if err_js != nil {
		c.RenderingJsonAutomatically(2, "参数错误 json"+err_js.Error())
		return
	}

	for i, v := range params {
		madid := v["gid"]       // 医嘱ID
		extime := v["ext"]      // 计划执行时间
		excycle := v["exc"]     // 分组序号
		ptType := v["ptType"]  // 打印单类型1=输液单，2=口服单，3=注射单，4=瓶签，5=检验标签


		if madid == "" || extime == "" || excycle == "" || ptType == "" {
			c.RenderingJsonAutomatically(1, "参数不完整 "+fmt.Sprintf("%d", i))
			return
		}
		gid, err_m := utils.Int64Value(madid)
		if err_m != nil || gid <= 0 {
			c.RenderingJsonAutomatically(2, "参数错误 gid "+fmt.Sprintf("%d", i))
			return
		}

		exc, err_exc := strconv.Atoi(excycle)
		if err_exc != nil || exc < 1 {
			c.RenderingJsonAutomatically(2, "参数错误 exc "+fmt.Sprintf("%d", i))
			return
		}

		if _, err_t := c.CheckingTimeFormat(extime); err_t != nil {
			c.RenderingJsonAutomatically(2, "参数错误 ext "+fmt.Sprintf("%d", i))
			return
		}

		t, err_t := strconv.Atoi(ptType)
		if err_t != nil || (t < 1 && t > 5)  {
			c.RenderingJsonAutomatically(2, "参数错误 ptType "+fmt.Sprintf("%d", i))
			return
		}

		obj, err_ch := model.CheckingMedicalAdvice(gid, extime, exc)

		if err_ch != nil || obj.Madid != gid {
			c.RenderingJsonAutomatically(4, "医嘱不存在")
			return
		}


		if err_db := model.UpdateMedicalAdvicePrintStatus(gid, excycle, extime, ptType, obj); err_db != nil {
			c.RenderingJsonAutomatically(3, "DataBase "+err_db.Error())
			return
		}
	}

	c.RenderingJsonAutomatically(0, "成功")
}

func (c *PCMedicalAdviceController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *PCMedicalAdviceController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
