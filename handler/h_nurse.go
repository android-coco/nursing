package handler

import "fit"

type NurseController struct {
	fit.Controller
}

// 病人基本信息
type PatientInfo struct {
	VAA01 int  // 病人id
	VAA05 string // 姓名
	ABW01 string // 性别 0=未知，1=M=男，2=F=女，9=未说明
	VAA10 int    // 年龄
	BCQ04 string // 床号
	BCE03C string // 住院医师
	BCE03B string // 责任护士
	VAE11 string // 入院日期
	VAJ36 float64 // 全额
	VAT39 int // 费用标志 0=正常，1=自费，2=免费
}

func (c NurseController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	nurse_id := r.FormValue("nurse_id")
	department_id := r.FormValue("department_id")

	var userInfos []PatientInfo
	//err := fit.SQLServerEngine().SQL("select B.VAA01, B.VAA05, B.ABW01, B.VAA10 ,B.BCQ04 from VAE1 as A JOIN VAA1 as B on A.VAA01 = B.VAA01 where VAE04 in (0, 2) and BCK01D = ?  ", department_id ,nurse_id).Find(&userInfos)
	err := fit.SQLServerEngine().SQL("select A.BCE03C, A.BCE03B, B.VAJ36, B.VAT39 from VAE1 as A join VAJ2 as B on A.VAA01 = B.VAA01 where VAE04 in (0, 2) and BCK01D = ?  ", department_id ,nurse_id).Find(&userInfos)
	c.JsonData.Result = 1000
	c.JsonData.ErrorMsg = ""
	c.JsonData.Datas = []interface{}{userInfos, err}
	return



}

/*func (c NurseController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	nurse_id := r.FormValue("nurse_id")
	department_id := r.FormValue("department_id")

	var userInfos []PatientInfo
	err := fit.SQLServerEngine().SQL("select B.VAA01, B.VAA05, B.ABW01, B.VAA10 ,B.BCQ04 from VAE1 as A JOIN VAA1 as B on A.VAA01 = B.VAA01 where VAE04 in (0, 2) and BCK01D = ?  ", department_id,nurse_id).Find(&userInfos)
	c.JsonData.Result = 1000
	c.JsonData.ErrorMsg = ""
	c.JsonData.Datas = []interface{}{userInfos, err}
	return



}*/
