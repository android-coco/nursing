package handler

import (
	"fit"
	"nursing/model"
	"fmt"
	"encoding/json"
	"time"
	"strconv"
)

type TvController struct {
	//fit.Controller
	PCController
}

//床位列表
func (c TvController) BedList(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	classId := r.FormIntValue("class_id")      //科室ID
	page := r.FormIntValue("page")             //页数
	pagenumber := r.FormIntValue("pagenumber") //每页多少数据
	showEmptystr := r.FormValue("showempty")   //是否显示空病床

	if classId == 0 || page == 0 || pagenumber == 0 {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	var showEmpty bool
	if showEmptystr == "" || showEmptystr == "0" {
		showEmpty = true
	}

	response, err := model.GetDepartmentBedsByClassifyingByPage(classId, 0, page, pagenumber, showEmpty)
	if err != nil {
		fit.Logger().LogError("Error", "TVBedList :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "查询失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "OK！"
		c.JsonData.Datas = response
	}
}

//TV时间
func (c TvController) Time(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "OK！"
	c.JsonData.Datas = time.Now().Format("2006年01月02日 15:04")
}

//查看信息
func (c TvController) List(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	classId := r.FormValue("class_id")
	classIdInt := r.FormIntValue("class_id")

	if classId == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	response, err := model.TVQueryMonitor(classId)
	// 获取所有病人
	amount, bedStr ,_,_ := model.FetchInpatientsForTV(classIdInt)
	if len(response) == 0 {
		response = make([]model.MonitorInfo, 1)
		response[0].MonitorNotifys = make([]model.MonitorNotify, 0)
		response[0].ClassId = classId
		response[0].Display = "1"
	}
	//病人总数
	response[0].V1 = strconv.Itoa(amount)
	// 获取一级护理床位
	response[0].V4 = bedStr
	//新入
	//response[0].V2 = newInPut
	//出院
	// response[0].V3 = newOutPut
	if err != nil {
		fit.Logger().LogError("Error", "TVList :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "查询失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "OK！"
		c.JsonData.Datas = response
	}
}

//TV编辑页面删除通知
func (c TvController) DelNotify(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	nid := r.FormInt64Value("nid")

	if nid <= 0 {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}
	err := model.TVDelMonitorNotify(model.MonitorNotify{ID: nid})
	if err != nil {
		fit.Logger().LogError("Error", "DelNotify :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "删除失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "删除成功！"
		c.JsonData.Datas = []interface{}{}
	}
}

//TV编辑页面修改信息
func (c TvController) UpdateMonitorInfo(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	var err error
	//userinfo, err := c.GetLocalUserinfo(w, r)
	//if err != nil {
	//	c.JsonData.Result = 3
	//	c.JsonData.ErrorMsg = "服务器有点繁忙!"
	//	c.JsonData.Datas = []interface{}{}
	//	return
	//}
	mid := r.FormInt64Value("mid")          //数据ID
	v1 := r.FormValue("v1")                 //病人总数
	v2 := r.FormValue("v2")                 //新入
	v3 := r.FormValue("v3")                 //出院
	v4 := r.FormValue("v4")                 //一级护理
	v5 := r.FormValue("v5")                 //病重
	v6 := r.FormValue("v6")                 //病危
	v7 := r.FormValue("v7")                 //心电监护
	v8 := r.FormValue("v8")                 //吸氧
	v9 := r.FormValue("v9")                 //神志瞳孔监测
	v10 := r.FormValue("v10")               //雾化吸入
	v11 := r.FormValue("v11")               //口腔护理
	v12 := r.FormValue("v12")               //尿道口护理
	v13 := r.FormValue("v13")               //留置导尿
	v14 := r.FormValue("v14")               //计24小时尿量
	v15 := r.FormValue("v15")               //计24小时出入量
	v16 := r.FormValue("v16")               //血压监测
	v17 := r.FormValue("v17")               //血糖检测
	v18 := r.FormValue("v18")               //自定义1
	v19 := r.FormValue("v19")               //自定义2
	v20 := r.FormValue("v20")               //
	nursename := r.FormValue("nursename")   //值班护士
	doctorname := r.FormValue("doctorname") //值班医生
	//mid := r.FormValue("Speed") //轮播速度等级1,2,3,4
	speed := "1"                        //轮播速度等级1,2,3,4
	display := r.FormValue("display")   //是否显示床位
	username := r.FormValue("username") //最后编辑人
	//username := userinfo.Name //最后编辑人
	classname := r.FormValue("classname") //科室名称
	//classname := userinfo.DepartmentName //科室名称
	classid := r.FormValue("classid") //科室ID
	//classid := userinfo.DepartmentID //科室ID
	//通知字符串Json格式[{"id": 1, "mid": 1,"notifyinfo": "就弗拉季德胜路附11近"},{"id": 2, "mid": 1,"notifyinfo": "就弗拉季德胜路附近11"}]
	infostr := r.FormValue("infos") //通知内容

	if classid == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	var inofs []model.MonitorNotify
	if infostr != "" {
		err = json.Unmarshal([]byte(infostr), &inofs)
	}
	data := model.MonitorInfo{}
	data.ID = mid
	data.V1 = v1
	data.V2 = v2
	data.V3 = v3
	data.V4 = v4
	data.V5 = v5
	data.V6 = v6
	data.V7 = v7
	data.V8 = v8
	data.V9 = v9
	data.V10 = v10
	data.V11 = v11
	data.V12 = v12
	data.V13 = v13
	data.V14 = v14
	data.V15 = v15
	data.V16 = v16
	data.V17 = v17
	data.V18 = v18
	data.V19 = v19
	data.V20 = v20
	data.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	data.NurseName = nursename
	data.DoctorName = doctorname
	data.Speed = speed
	data.Display = display
	data.UserName = username
	data.ClassName = classname
	data.ClassId = classid
	data.MonitorNotifys = inofs

	//fit.Logger().LogError("TV更新数据：", data)

	err = model.TVUpdataMonitorInfo(data)
	if err != nil {
		fit.Logger().LogError("Error", "TVList :", err)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "更新失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "更新成功！"
		c.JsonData.Datas = []interface{}{}
	}
}

//TV 管理界面
func (c TvController) Manage(w *fit.Response, r *fit.Request, p fit.Params) {
	userinfo, err1 := c.GetLocalUserinfo(w, r)                                  //用户信息
	response, err2 := model.TVQueryMonitor(strconv.Itoa(userinfo.DepartmentID)) //TV 信息
	if err1 == nil && err2 == nil {
		defer c.LoadViewSafely(w, r, "pc/v_tv_manage.html", "pc/header_side.html", "pc/header_top.html")

		c.Data = fit.Data{
			"Userinfo":     userinfo,
			"MonitorInfos": response,
		}
	} else {
		fmt.Fprintln(w, "服务器有点繁忙！")
	}
}
