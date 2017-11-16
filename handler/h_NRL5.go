package handler

import (
	"fit"
	"fmt"
	"strconv"
	"time"
	"nursing/model"
)

type NRL5Controller struct {
	NRLController

}

// 模板 template PDA端
func (c NRL5Controller) Check(w *fit.Response, r *fit.Request, p fit.Params) {
	// 文书id
	rid := r.FormValue("rid")
	if "" == rid {
		fmt.Fprintln(w, "参数不完整！")
		return
	}

	nr5, err1 := model.QueryNRL5(rid)
	if err1 != nil {
		fit.Logger().LogError("m_NR5", err1)
	}

	pid := nr5.VAA01

	// 查询对应病人信息
	patient, has := c.LoadPinfoWithPid(w, r, pid)
	if !has {
		return
	}


	recordDate := nr5.DateTime.Format("2006-01-02")
	c.Data = fit.Data{
		"Pinfo": patient,
		"NRL":   nr5,
		"RecordDate": recordDate,
	}

	fmt.Printf("data %+v\n", c.Data)
	c.LoadView(w, "v_nrl5.html")
}

func (c NRL5Controller) Edit(w *fit.Response, r *fit.Request, p fit.Params) {
	pid := r.FormValue("pid")
	uid := r.FormValue("uid")
	rid := r.FormValue("rid") // 护理记录单id
	ty := r.FormValue("type") // 1=add， 2=edit

	var nr5 model.NRL5
	if ty == "1" {
		if "" == pid || "" == uid {
			fmt.Fprintln(w, "参数错误！")
			return
		}
	} else if ty == "2" {
		if rid == "" {
			fmt.Fprintln(w, "参数错误！")
			return
		}
		var err1 error
		nr5, err1 = model.QueryNRL5(rid)
		if err1 != nil {
			fit.Logger().LogError("m_NR1", err1)
		}
		pid = strconv.FormatInt(nr5.VAA01, 10)
		uid = nr5.BCE01A
	} else {
		fmt.Fprintln(w, "参数错误！")
		return
	}

	// 查询对应病人信息 护士的信息
	patient, account, has := c.LoadPinfoAndAccountWithPidUid(w, r, pid, uid)
	if !has {
		return
	}
	recordDate := nr5.DateTime.Format("2006-01-02")

	c.Data = fit.Data{
		"Pinfo": patient,
		"NRL":   nr5,
		"Type":  ty,
		"Rid": rid,
		"Account": account,
		"RecordDate": recordDate,
	}

	c.LoadView(w, "v_nrl5_edit.html")
}

// 接口
// 添加护理记录单
func (c NRL5Controller) AddRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)
	// 科室ID
	BCK01, err5 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 文书

	NRL01 := r.FormValue("NRL01")
	NRL02A := r.FormValue("NRL02A")
	NRL02B := r.FormValue("NRL02B")
	NRL03A := r.FormValue("NRL03A")
	NRL03B := r.FormValue("NRL03B")
	NRL04A := r.FormValue("NRL04A")
	NRL04B := r.FormValue("NRL04B")
	NRL05A := r.FormValue("NRL05A")
	NRL05B := r.FormValue("NRL05B")
	NRL06A := r.FormValue("NRL06A")
	NRL06B := r.FormValue("NRL06B")
	NRL07A := r.FormValue("NRL07A")
	NRL07B := r.FormValue("NRL07B")
	NRL08A := r.FormValue("NRL08A")
	NRL08B := r.FormValue("NRL08B")
	NRL09A := r.FormValue("NRL09A")
	NRL09B := r.FormValue("NRL09B")
	NRL10A := r.FormValue("NRL10A")
	NRL10B := r.FormValue("NRL10B")
	NRL10C := r.FormValue("NRL10C")
	NRL10D := r.FormValue("NRL10D")
	NRL11 := r.FormValue("NRL11")
	NRL12 := r.FormValue("NRL12")
	score := r.FormValue("score")




	checkerr( "nrl5 add", err1, err4, err5)

	if VAA01 == 0 || BCE01A == "" || BCE03A == "" {
		c.JsonData.Result = 1
		c.JsonData.ErrorMsg = "参数不完整"
		c.JsonData.Datas = []interface{}{}
		return
	}

	nrl5 := model.NRL5{
		BCK01:    BCK01,
		VAA01:    VAA01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02A:    NRL02A,
		NRL02B:    NRL02B,
		NRL03A:    NRL03A,
		NRL03B:    NRL03B,
		NRL04A:    NRL04A,
		NRL04B:    NRL04B,
		NRL05A:    NRL05A,
		NRL05B:    NRL05B,
		NRL06A:    NRL06A,
		NRL06B:    NRL06B,
		NRL07A:    NRL07A,
		NRL07B:    NRL07B,
		NRL08A:    NRL08A,
		NRL08B:    NRL08B,
		NRL09A:    NRL09A,
		NRL09B:    NRL09B,
		NRL10A:    NRL10A,
		NRL10B:    NRL10B,
		NRL10C:    NRL10C,
		NRL10D:    NRL10D,
		NRL11:    NRL11,
		NRL12:    NRL12,
		Score:    score,
	}
	has, errExist := nrl5.IsExistNRL5()
	if has {
		c.RenderingJsonAutomatically(3, "已存在该班次，请勿重复提交。")
		return
	}
	if errExist != nil {
		c.RenderingJsonAutomatically(2, "上传失败！")
		return
	}
	rid, err17 := nrl5.InsertData()

	if err17 != nil {
		fit.Logger().LogError("NRL5 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "上传失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		// 文书记录
		nurseRecord := model.NursingRecords{
			Updated:     r.FormValue("datetime"),
			NursType:    5,
			NursingId:   BCE01A,
			NursingName: BCE03A,
			ClassId:     r.FormValue("did"),
			PatientId:   r.FormValue("pid"),
			RecordId:    rid,
			Comment:     "新增",
		}
		_,errRecord := model.InsertNRecords(nurseRecord)
		checkerr("nurse record err:", errRecord)

		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "上传成功！"
		c.JsonData.Datas = []interface{}{}
	}

}

// 修改护理记录单
func (c NRL5Controller) UpdateRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	// 文书ID
	rid := r.FormValue("rid")
	id, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		fit.Logger().LogError("Error", "nrl1 update :", err)
		c.JsonData.Result = 3
		c.JsonData.ErrorMsg = "rid 错误！"
		c.JsonData.Datas = []interface{}{}
	}
	// 病人ID
	VAA01, err1 := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	// 科室ID
	BCK01, err5 := strconv.ParseInt(r.FormValue("did"), 10, 64)
	// 护士ID
	BCE01A := r.FormValue("uid")
	// 护士名
	BCE03A := r.FormValue("username")
	// 记录时间
	datetime, err4 := time.ParseInLocation("2006-01-02 15:04:05", r.FormValue("datetime"), time.Local)

	NRL01 := r.FormValue("NRL01")
	NRL02A := r.FormValue("NRL02A")
	NRL02B := r.FormValue("NRL02B")
	NRL03A := r.FormValue("NRL03A")
	NRL03B := r.FormValue("NRL03B")
	NRL04A := r.FormValue("NRL04A")
	NRL04B := r.FormValue("NRL04B")
	NRL05A := r.FormValue("NRL05A")
	NRL05B := r.FormValue("NRL05B")
	NRL06A := r.FormValue("NRL06A")
	NRL06B := r.FormValue("NRL06B")
	NRL07A := r.FormValue("NRL07A")
	NRL07B := r.FormValue("NRL07B")
	NRL08A := r.FormValue("NRL08A")
	NRL08B := r.FormValue("NRL08B")
	NRL09A := r.FormValue("NRL09A")
	NRL09B := r.FormValue("NRL09B")
	NRL10A := r.FormValue("NRL10A")
	NRL10B := r.FormValue("NRL10B")
	NRL10C := r.FormValue("NRL10C")
	NRL10D := r.FormValue("NRL10D")
	NRL11 := r.FormValue("NRL11")
	NRL12 := r.FormValue("NRL12")
	score := r.FormValue("score")

	checkerr( "nrl5 add",err1, err5, err4)

	nrl5 := model.NRL5{
		VAA01:    VAA01,
		BCK01:    BCK01,
		BCE01A:   BCE01A,
		BCE03A:   BCE03A,
		DateTime: datetime,
		NRL01:    NRL01,
		NRL02A:    NRL02A,
		NRL02B:    NRL02B,
		NRL03A:    NRL03A,
		NRL03B:    NRL03B,
		NRL04A:    NRL04A,
		NRL04B:    NRL04B,
		NRL05A:    NRL05A,
		NRL05B:    NRL05B,
		NRL06A:    NRL06A,
		NRL06B:    NRL06B,
		NRL07A:    NRL07A,
		NRL07B:    NRL07B,
		NRL08A:    NRL08A,
		NRL08B:    NRL08B,
		NRL09A:    NRL09A,
		NRL09B:    NRL09B,
		NRL10A:    NRL10A,
		NRL10B:    NRL10B,
		NRL10C:    NRL10C,
		NRL10D:    NRL10D,
		NRL11:    NRL11,
		NRL12:    NRL12,
		Score:    score,
	}

	_, err17 := nrl5.UpdateData(id)


	if err17 != nil {
		fit.Logger().LogError("nrl4 add :", err17)
		c.JsonData.Result = 2
		c.JsonData.ErrorMsg = "修改失败！"
		c.JsonData.Datas = []interface{}{}
	} else {
		_, errRecord := model.UpadteNRecords(id, r.FormValue("datetime"))
		checkerr("nurse record update err:", errRecord)
		c.JsonData.Result = 0
		c.JsonData.ErrorMsg = "修改成功！"
		c.JsonData.Datas = []interface{}{}
	}
}


func (c NRL5Controller) Exist(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	pid := r.FormValue("pid")
	datestr := r.FormValue("date")
	if pid == "" || datestr == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	shifts, err_exs := model.IsExistNRL5Shift(pid, datestr)
	if err_exs != nil {
		c.RenderingJsonAutomatically(3, "Database "+err_exs.Error())
		return
	}

	c.RenderingJson(0,"成功！", shifts)
}


func checkerr(tag string,err ...error) bool  {
	var flag bool = false
	for _, val := range err {
		if val != nil {
			fit.Logger().LogError(tag, val)
			flag = true
		}
	}
	return flag
}