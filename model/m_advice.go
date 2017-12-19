package model

import (
	"fit"
	"time"
	"fmt"
	"strings"
)

/*PDA医嘱 查询结果*/
type MedicalAdvice struct {
	VAF01  int64                  `json:"madid"`       // 医嘱ID
	VAA01  int64                  `json:"pid"`         // 病人ID
	VAF10  int                    `json:"status"`      // 状态
	VAF11  int                    `json:"type"`        // 长期医嘱/临时医嘱
	BDA01  string                 `json:"category"`    // 医嘱类别，诊疗类型
	VAF19  string                 `json:"dosage"`      // 用量
	VAF22  string                 `json:"content"`     // 医嘱内容
	BBX01  int                    `json:"-"`           // 诊疗类型BBX05，可能跟用药途径关联
	Method string                 `json:"method"`      // 用药途径
	VAF23  string                 `json:"entrust"`     // 医师嘱托
	VAF26  string                 `json:"time"`        // 次数/执行频次
	VAF27  int                    `json:"count"`       // 次数
	VAF36  DatetimeWithoutSeconds `json:"starttime"`   // 开始执行时间
	BCE03A string                 `json:"physician"`   // 开嘱医师
	BBX20  int                    `json:"classify"`    // 按打印单分类
	VAF60  string                 `json:"speed"`       // 滴速
	VAF42  DatetimeWithoutSeconds `json:"createdtime"` // 开嘱时间
}

/*PDA医嘱执行 查询结果*/
//type MedicalAdviceExecution struct {
//	MedicalAdvice                `xorm:"extends"`
//	MedicalAdviceExecutionStatus `xorm:"extends"`
//}

/*医嘱执行状态*/
//type MedicalAdviceExecutionStatus struct {
//	Madid   int64  `json:"-"`
//	State   int    `json:"state"`   // 执行状态
//	Process string `json:"process"` // 执行步骤
//}

/*医嘱执行详细*/
//type MedicalAdviceExecutionDetail struct {
//	MedicalAdvice                          `xorm:"extends"`
//	MedicalAdviceExecutionStatus           `xorm:"extends"`
//	Records []MedicalAdviceExecutionRecord `json:"records" xorm:"extends"`
//	Master  []VAA1                         `json:"master"` // 病人信息
//	Desc    string                         `json:"desc"`   // 关于医嘱的描述
//}

//type MedicalAdviceDup struct {
//	VAF01 int64 // 医嘱记录ID
//}

/*医嘱执行状态*/
//type AdviceStatus struct {
//	Patientid  int64
//	Madid      int64
//	State      int
//	Recordtime string
//	Nurseid    int
//	Nursename  string
//	Period     int
//	Process    string
//}

/*医嘱执行明细*/
//type AdviceDetail struct {
//	ExecTime  string `json:"exectime"`  // 执行时间
//	Patientid int64  `json:"pid"`       // 病人ID
//	Nurseid   int    `json:"nid"`       // 护士ID
//	Period    int    `json:"period"`    // 周期
//	Madid     int64  `json:"madid"`     // 医嘱ID
//	Nursename string `json:"nursename"` // 执行护士
//	Process   string `json:"process"`   // 执行步骤
//}

/*医嘱执行明细*/
type MedicalAdviceExecutionRecord struct {
	Patientid int64  `json:"pid"`       // 病人ID
	Nurseid   int    `json:"nid"`       // 护士ID
	Period    int    `json:"period"`    // 周期
	Madid     int64  `json:"madid"`     // 医嘱ID
	Nursename string `json:"nursename"` // 执行护士
	Process   string `json:"process"`   // 执行步骤
	ExecTime  string `json:"exectime"`  // 执行时间
	ExCycle   int    `json:"excycle"`   // 同组医嘱的序号
	Plan      string `json:"plan"`      // 计划执行时间
}

/*JP 今天的起始时间*/
func DatetimeNow() string {
	t := time.Now()
	dateNow := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
	return dateNow
}

/*JP 登记表中的入院时间*/
type AdmissionTime struct {
	Vid    int64  `json:"vid"`           // 就诊ID
	VAA01  int64  `json:"patientId"`     // 病人ID
	VAE11  string `json:"admissionTime"` // 入院时间
	BCK01C int    `json:"did"`           // 住院病区ID
}

/*JP PDA医嘱查询*/
//func SearchMedicalAdvice(typeOf, status int, pid int64, category string) ([]MedicalAdvice, error) {
//	condition_type := ""
//	if typeOf != 0 {
//		// 长期和临时
//		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
//	}
//
//	condition_catg := ""
//	if category != "0" {
//		// 医嘱类别
//		condition_catg = " and a.BDA01 = '" + category + "'"
//	}
//
//	// 医嘱状态所有
//	condition_stat := " and a.VAF10 in (3,4,8,9)"
//	if status != 0 {
//		// 医嘱状态 未停、已撤销、已停
//		if status != 8 {
//			condition_stat = fmt.Sprintf(" and a.VAF10 = '%d'", status)
//		} else {
//			condition_stat = " and a.VAF10 in (8,9)"
//		}
//	}
//	// 时间  VAF36
//	admissionTime := AdmissionTime{}
//	mAdvices := make([]MedicalAdvice, 0)
//
//	fit.SQLServerEngine().SQL("select VAA01, VAE11, BCK01C from VAE1 where VAA01 = ?", pid).Get(&admissionTime)
//	if admissionTime.VAA01 != pid {
//		return mAdvices, fmt.Errorf("病人登记(VAE1)表中无此病人数据")
//	}
//
//	sqlStr := fmt.Sprintf("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF27,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAA01 = '%d' and a.VAF32 = 0 and a.VAF42 > '%s'%s%s%s order by a.VAF36", pid, admissionTime.VAE11, condition_type, condition_stat, condition_catg)
//	//fit.SQLServerEngine().ShowSQL(true)
//	err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
//	//fit.SQLServerEngine().ShowSQL(false)
//	return mAdvices, err
//}

/*医嘱执行查询*/
//func SearchMedicalAdviceExecution(typeOf, status int, pid int64, category, st, et string) ([]MedicalAdviceResponse, error) {
//condition_type := ""
//if typeOf != 0 {
//	condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
//}
//
//condition_catg := ""
//if category != "0" {
//	condition_catg = " and a.BDA01 = '" + category + "'"
//}
//
//admissionTime := AdmissionTime{}
//mAdvices := make([]MedicalAdvice, 0)
//
//fit.SQLServerEngine().SQL("select VAA01, VAE11, BCK01C from VAE1 where VAA01 = ?", pid).Get(&admissionTime)
//if admissionTime.VAA01 != pid {
//	return make([]MedicalAdviceExecution, 0), fmt.Errorf("病人登记(VAE1)表中无此病人数据")
//}
//
//sqlStr := fmt.Sprintf("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF27,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAA01 = '%d' and a.VAF32 = '0' and a.VAF10 = '3' and a.VAF42 > '%s'%s%s order by a.VAF36", pid, admissionTime.VAE11, condition_type, condition_catg)
//err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
//if err != nil {
//	return make([]MedicalAdviceExecution, 0), err
//}
//
//condition_stat := ""
//if status != 0 && status != 1 {
//	condition_stat = fmt.Sprintf(" and State = '%d'", status)
//}
//response := make([]MedicalAdviceExecution, 0)
//for _, v := range mAdvices {
//	madStatus := make([]MedicalAdviceExecutionStatus, 0)
//	sql := fmt.Sprintf("select Madid,State,Process from AdviceStatus where Madid = %d%s", v.VAF01, condition_stat)
//	err_db := fit.MySqlEngine().SQL(sql).Find(&madStatus)
//	if err_db != nil {
//		fit.Logger().LogError("**JK**", err_db.Error())
//	}
//
//	if leng := len(madStatus); leng == 0 {
//		status := MedicalAdviceExecutionStatus{}
//		status.Process = "未执行"
//		status.State = 1
//		status.Madid = v.VAF01
//		madStatus = append(madStatus, status)
//	}
//	record := madStatus[0]
//	if status == 0 {
//		execution := MedicalAdviceExecution{}
//		execution.MedicalAdvice = v
//		execution.MedicalAdviceExecutionStatus = record
//		response = append(response, execution)
//	} else if status == record.State {
//		execution := MedicalAdviceExecution{}
//		execution.MedicalAdvice = v
//		execution.MedicalAdviceExecutionStatus = record
//		response = append(response, execution)
//	}
//}
//
//return response, err
//}

/*医嘱执行详情*/
func FetchMedicalAdviceExecutionDetail(gid int64, ext string, exc int) ([]MedicalAdviceExecutionDetail, error) {
	//madid int64, master int  gid int64, ext string, exc int
	//mAdvice := make([]MedicalAdvice, 0)
	//err := fit.SQLServerEngine().SQL("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF27,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAF01 = ?", madid).Find(&mAdvice)
	//if err != nil {
	//	fit.Logger().LogError("***JP***", err.Error())
	//	return make([]MedicalAdviceExecutionDetail, 0), err
	//}
	//
	//if length := len(mAdvice); length == 0 {
	//	return make([]MedicalAdviceExecutionDetail, 0), fmt.Errorf("无法查询到对应医嘱,二维码无效")
	//}
	//detail := MedicalAdviceExecutionDetail{}
	//detail.MedicalAdvice = mAdvice[0]
	//detail.Records = make([]MedicalAdviceExecutionRecord, 0)
	//detail.Master = make([]VAA1, 0)
	//detail.State = 0
	//detail.Process = "未知"
	//if master == 1 {
	//	fit.MySqlEngine().Table("VAA1").Where("VAA01 = ?", detail.MedicalAdvice.VAA01).Find(&detail.Master)
	//}
	//
	//response := make([]MedicalAdviceExecutionDetail, 1)
	//if detail.MedicalAdvice.VAF10 < 3 {
	//	detail.Desc = "该医嘱暂未生效"
	//	response[0] = detail
	//	return response, err
	//} else if detail.MedicalAdvice.VAF10 == 4 {
	//	detail.Desc = "该医嘱已被作废"
	//	response[0] = detail
	//	return response, err
	//} else if detail.MedicalAdvice.VAF10 > 4 && detail.MedicalAdvice.VAF10 < 7 {
	//	detail.Desc = "该医嘱已被删除或暂停"
	//	response[0] = detail
	//	return response, err
	//} else if detail.MedicalAdvice.VAF10 > 8 && detail.MedicalAdvice.VAF10 < 10 {
	//	detail.Desc = "该医嘱已停止"
	//	response[0] = detail
	//	return response, err
	//} else if detail.MedicalAdvice.VAF11 == 2 {
	//	timeMark := time.Time(detail.MedicalAdvice.VAF36)
	//	timeNow := time.Now()
	//	if timeNow.Sub(timeMark).Hours() <= 24 {
	//		if timeNow.Day()-timeMark.Day() > 0 {
	//			detail.Desc = "该医嘱已失效"
	//			response[0] = detail
	//			return response, err
	//		}
	//	} else {
	//		detail.Desc = "该医嘱已失效"
	//		response[0] = detail
	//		return response, err
	//	}
	//}
	//
	//err = fit.MySqlEngine().SQL("select * from AdviceDetail where Madid = ?", madid).Find(&detail.Records)
	//if leng := len(detail.Records); leng == 0 {
	//	detail.Desc = "该医嘱暂未执行"
	//	detail.Process = "未执行"
	//	detail.State = 1
	//} else {
	//	fit.MySqlEngine().SQL("select Madid,State,Process from AdviceStatus where Madid = ?", detail.MedicalAdvice.VAF01).Get(&detail.MedicalAdviceExecutionStatus)
	//}
	//response[0] = detail
	//return response, err

	orgin := MedicalAdviceModal{}
	_, err_res := fit.SQLServerEngine().SQL("select a.VAF01 Madid, a.VAF06 Vid, a.CBM01 PtNum, a.VAF59 GroupNum from VAF2 a where a.VAF01 = ?", gid).Get(&orgin)
	response := make([]MedicalAdviceExecutionDetail, 0)
	if err_res != nil {
		fit.Logger().LogError("***JK***", err_res.Error())
	}
	if orgin.Madid != gid {
		return response, err_res
	}
	mAdvices := make([]MedicalAdviceModal, 0)
	err_res = fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT d.* FROM( SELECT CASE WHEN(( a2.BBX20 = 2 OR a2.BBX20 = 4 OR a2.BBX20 = 5) AND(a.BDA01 = '1' OR a.BDA01 = '2') AND a2.BDA01 = 'T' AND a2.BBX13 = '2') THEN 4 ELSE 0 END PtTypeV, a.VAF06 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName,c.VAE46 Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 LEFT JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF06 = %d AND a.VAF32 = 0 AND a.VAF04 = 2 AND a.CBM01 = %d AND a.VAF59 = %d AND b.VBI10 = '%s') d ORDER BY d.TypeV, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr ", orgin.Vid, orgin.PtNum, orgin.GroupNum, ext)).Find(&mAdvices)
	if err_res != nil {
		fit.Logger().LogError("***JK***", err_res.Error())
		return response, err_res
	} else if length := len(mAdvices); length == 0 {
		return response, err_res
	}
	arrA := make([]MedicalAdviceResponse, 0)

	// 取出第一条医嘱数据
	temp := mAdvices[0]
	temp.ExDay = temp.ExTime.ParseToSecond()
	temp.StDay = temp.StTime.ParseToMinute()
	temp.EdDay = temp.EdTime.ParseToMinute()

	// 用temp的数据实例第一个object
	object := MedicalAdviceResponse{
		Vid:       temp.Vid,
		Pid:       temp.Pid,
		Bed:       temp.Bed,
		PName:     temp.PName,
		Age:       temp.Age,
		HospNum:   temp.HospNum,
		Gender:    temp.Gender,
		ExTime:    temp.ExTime,
		ExDay:     temp.ExDay,
		GroupNum:  temp.GroupNum,
		Amount:    temp.Amount,
		Frequency: temp.Frequency,
		Times:     temp.Times,
		Method:    temp.Method,
		Speed:     temp.Speed,
		TypeV:     temp.TypeV,
		TypeOf:    temp.TypeOf,
		StTime:    temp.StTime,
		StDay:     temp.StDay,
		MStatus:   temp.MStatus,
		MStatusV:  temp.MStatusV,
		Category:  temp.Category,
		CategoryV: temp.CategoryV,
		PtType:    temp.PtType,
		PtTypeV:   temp.PtTypeV,
		PtNum:     temp.PtNum,
		PtRownr:   temp.PtRownr,
		Entrust:   temp.Entrust,
		Physician: temp.Physician,
		EdTime:    temp.EdTime,
		EdDay:     temp.EdDay,
		Sender:    temp.Sender,
		ExCycle:   1,
	}
	object.Gid = temp.Madid
	object.Contents = make([]MedicalAdviceContent, 1)
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content," ")[0], temp.Dosage}

	length := len(mAdvices)
	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage})
		} else {
			//	按次拆分
			for idx := 1; idx <= object.Times; idx ++ {
				if idx == 1 {
					object.ExCycle = idx
					if object.Vid == orgin.Vid && object.PtNum == orgin.PtNum && object.GroupNum == orgin.GroupNum && idx == exc {
						arrA = append(arrA, object)
					}
				} else {
					if object.Vid == orgin.Vid && object.PtNum == orgin.PtNum && object.GroupNum == orgin.GroupNum && idx == exc {
						obj := MedicalAdviceResponse{
							Vid:       object.Vid,
							Pid:       object.Pid,
							Bed:       object.Bed,
							PName:     object.PName,
							Age:       object.Age,
							HospNum:   object.HospNum,
							Gender:    object.Gender,
							ExTime:    object.ExTime,
							ExDay:     object.ExDay,
							GroupNum:  object.GroupNum,
							Amount:    object.Amount,
							Frequency: object.Frequency,
							Times:     object.Times,
							Method:    object.Method,
							Speed:     object.Speed,
							TypeV:     object.TypeV,
							TypeOf:    object.TypeOf,
							StTime:    object.StTime,
							StDay:     object.StDay,
							MStatus:   object.MStatus,
							MStatusV:  object.MStatusV,
							Category:  object.Category,
							CategoryV: object.CategoryV,
							PtType:    object.PtType,
							PtTypeV:   object.PtTypeV,
							PtNum:     object.PtNum,
							PtRownr:   object.PtRownr,
							Entrust:   object.Entrust,
							Physician: object.Physician,
							EdTime:    object.EdTime,
							EdDay:     object.EdDay,
							Sender:    object.Sender,
							ExCycle:   idx,
						}
						obj.Gid = object.Contents[0].Madid
						obj.Contents = make([]MedicalAdviceContent, 0)
						obj.Contents = append(obj.Contents, object.Contents...)
						arrA = append(arrA, object)
					}
				}
			}

			object = MedicalAdviceResponse{
				Vid:       v.Vid,
				Pid:       v.Pid,
				Bed:       v.Bed,
				PName:     v.PName,
				Age:       v.Age,
				HospNum:   v.HospNum,
				Gender:    v.Gender,
				ExTime:    v.ExTime,
				ExDay:     v.ExDay,
				GroupNum:  v.GroupNum,
				Amount:    v.Amount,
				Frequency: v.Frequency,
				Times:     v.Times,
				Method:    v.Method,
				Speed:     v.Speed,
				TypeV:     v.TypeV,
				TypeOf:    v.TypeOf,
				StTime:    v.StTime,
				StDay:     v.StDay,
				MStatus:   v.MStatus,
				MStatusV:  v.MStatusV,
				Category:  v.Category,
				CategoryV: v.CategoryV,
				PtType:    v.PtType,
				PtTypeV:   v.PtTypeV,
				PtNum:     v.PtNum,
				PtRownr:   v.PtRownr,
				Entrust:   v.Entrust,
				Physician: v.Physician,
				EdTime:    v.EdTime,
				EdDay:     v.EdDay,
				Sender:    v.Sender,
				ExCycle:   v.ExCycle,
			}
			object.Gid = v.Madid
			object.Contents = make([]MedicalAdviceContent, 1)
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage}
		}
	}
	// 继续拆分最后一次创建的object
	if length == 1 {
		arrA = append(arrA, object)
	} else {
		for idx := 1; idx <= object.Times; idx ++ {
			if idx == 1 {
				if object.Vid == orgin.Vid && object.PtNum == orgin.PtNum && object.GroupNum == orgin.GroupNum && idx == exc {
					object.ExCycle = idx
					arrA = append(arrA, object)
				}
			} else {
				if object.Vid == orgin.Vid && object.PtNum == orgin.PtNum && object.GroupNum == orgin.GroupNum && idx == exc {
					obj := MedicalAdviceResponse{
						Vid:       object.Vid,
						Pid:       object.Pid,
						Bed:       object.Bed,
						PName:     object.PName,
						Age:       object.Age,
						HospNum:   object.HospNum,
						Gender:    object.Gender,
						ExTime:    object.ExTime,
						ExDay:     object.ExDay,
						GroupNum:  object.GroupNum,
						Amount:    object.Amount,
						Frequency: object.Frequency,
						Times:     object.Times,
						Method:    object.Method,
						Speed:     object.Speed,
						TypeV:     object.TypeV,
						TypeOf:    object.TypeOf,
						StTime:    object.StTime,
						StDay:     object.StDay,
						MStatus:   object.MStatus,
						MStatusV:  object.MStatusV,
						Category:  object.Category,
						CategoryV: object.CategoryV,
						PtType:    object.PtType,
						PtTypeV:   object.PtTypeV,
						PtNum:     object.PtNum,
						PtRownr:   object.PtRownr,
						Entrust:   object.Entrust,
						Physician: object.Physician,
						EdTime:    object.EdTime,
						EdDay:     object.EdDay,
						Sender:    object.Sender,
						ExCycle:   idx,
					}
					obj.Gid = object.Contents[0].Madid
					obj.Contents = make([]MedicalAdviceContent, 0)
					obj.Contents = append(obj.Contents, object.Contents...)
					arrA = append(arrA, obj)
				}
			}
		}
	}

	length = len(arrA)
	if length == 0 {
		return response, err_res
	}

	// 取arrA的首个元素
	res := MedicalAdviceExecutionDetail{
		MedicalAdviceResponse: arrA[0],
		Records:               make([]MedicalAdviceExecutionRecord, 0),
	}

	exec := MedicalAdvicePrintSubModel{}
	// 用组里的第一个数据代替组 查询执行和打印记录
	_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", res.Gid, res.ExTime.ParseToSecond(), res.ExCycle).Get(&exec)
	if err_pt != nil {
		fit.Logger().LogError("***JK***", "医嘱执行明细-err_pt", err_pt.Error())
	}
	if exec.Madid != res.Gid {
		exec.Madid = res.Gid
		exec.ExNurse = ""
		exec.ExStatus = "未执行"
		exec.ExStatusV = 1
		exec.ExStep = ""
		exec.PtTimes = 0
		exec.PtStatus = "未打印"
	}

	if exec.ExStatusV == 0 {
		exec.ExStatusV = 1
	}

	res.ExStatusV = exec.ExStatusV
	res.ExStatus = exec.ExStatus
	res.ExNurse = exec.ExNurse
	res.ExStep = exec.ExStep
	res.PtTimes = exec.PtTimes
	res.PtStatus = exec.PtStatus

	if res.MStatusV < 3 {
		res.Desc = "医嘱暂未生效"
	} else if res.MStatusV == 4 {
		res.Desc = "该医嘱已作废"
	} else if res.MStatusV > 4 && res.MStatusV < 7 {
		res.Desc = "该医嘱已被删除或暂停"
	} else if res.MStatusV >= 8 {
		res.Desc = "该医嘱已停"
	} else if res.TypeV == 2 {
		timeMark := time.Time(res.StTime)
		timeNow := time.Now()
		if timeNow.Sub(timeMark).Hours() <= 24 {
			if timeNow.Day()-timeMark.Day() > 0 {
				res.Desc = "该医嘱已失效"
			}
		} else {
			res.Desc = "该医嘱已失效"
		}
	}

	if res.Desc != "" {
		return response, err_res
	}

	err_res = fit.MySqlEngine().SQL("select * from AdviceDetail where Madid = ? and Plan = ? And ExCycle = ?", gid, ext, exc).Find(&res.Records)
	if len(res.Records) == 0 {
		res.Desc = "empty"
		//	 无执行记录，但是可执行
	}
	if err_res != nil {
		fit.Logger().LogError("***JK***", err_res.Error())
	}
	response = append(response, res)
	return response, err_res
}

/*JP 是否存在新医嘱*/
func IsExistNewMedicalAdvice(vid int64, did int, hospitalDate string) int {
	// 获取校验过的医嘱
	today := DatetimeNow()
	//mAdvices := make([]MedicalAdviceDup, 0)
	//err_db := fit.SQLServerEngine().SQL("SELECT a.VAF01 from VAF2 a where a.VAF06 = ? and a.BCK01B = ? and a.VAF32 = 0 and a.VAF10 = 3 and ((a.VAF11 = 1 and a.VAF42 > ?) or (a.VAF11 = 2 and a.VAF42 > ?))", vid, did, hospitalDate, today).Find(&mAdvices)

	mAdvices := make([]MedicalAdviceModal, 0)
	err_db := fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT d.* FROM( SELECT a.VAF06 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 LEFT JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF06 = %d AND c.BCK01C = %d AND a.VAF32 = 0 AND a.VAF04 = 2 AND a.VAF10 = 3 AND((a.VAF11 = 1 AND b.VBI10 > '%s') OR(a.VAF11 = 2 AND b.VBI10 > '%s'))) d ORDER BY d.TypeV, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", vid, did, hospitalDate, today)).Find(&mAdvices)
	if err_db != nil {
		fit.Logger().LogError("**JP**", err_db.Error())
	}

	length := len(mAdvices)
	if length == 0 {
		return 0
	}

	arrA := make([]MedicalAdviceResponse, 0)

	// 取出第一条医嘱数据
	temp := mAdvices[0]
	temp.ExDay = temp.ExTime.ParseToSecond()
	temp.StDay = temp.StTime.ParseToMinute()
	temp.EdDay = temp.EdTime.ParseToMinute()

	// 用temp的数据实例第一个object
	object := MedicalAdviceResponse{
		Vid:       temp.Vid,
		Pid:       temp.Pid,
		Bed:       temp.Bed,
		PName:     temp.PName,
		Age:       temp.Age,
		HospNum:   temp.HospNum,
		Gender:    temp.Gender,
		ExTime:    temp.ExTime,
		ExDay:     temp.ExDay,
		GroupNum:  temp.GroupNum,
		Amount:    temp.Amount,
		Frequency: temp.Frequency,
		Times:     temp.Times,
		Method:    temp.Method,
		Speed:     temp.Speed,
		TypeV:     temp.TypeV,
		TypeOf:    temp.TypeOf,
		StTime:    temp.StTime,
		StDay:     temp.StDay,
		MStatus:   temp.MStatus,
		MStatusV:  temp.MStatusV,
		Category:  temp.Category,
		CategoryV: temp.CategoryV,
		PtType:    temp.PtType,
		PtNum:     temp.PtNum,
		PtRownr:   temp.PtRownr,
		Entrust:   temp.Entrust,
		Physician: temp.Physician,
		EdTime:    temp.EdTime,
		EdDay:     temp.EdDay,
		Sender:    temp.Sender,
		ExCycle:   temp.ExCycle,
	}
	object.Gid = temp.Madid
	object.Contents = make([]MedicalAdviceContent, 1)
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content," ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage})
		} else {
			//	按次拆分
			for idx := 1; idx <= object.Times; idx ++ {
				if idx == 1 {
					object.ExCycle = idx
					arrA = append(arrA, object)
				} else {
					obj := MedicalAdviceResponse{
						Vid:       object.Vid,
						Pid:       object.Pid,
						Bed:       object.Bed,
						PName:     object.PName,
						Age:       object.Age,
						HospNum:   object.HospNum,
						Gender:    object.Gender,
						ExTime:    object.ExTime,
						ExDay:     object.ExDay,
						GroupNum:  object.GroupNum,
						Amount:    object.Amount,
						Frequency: object.Frequency,
						Times:     object.Times,
						Method:    object.Method,
						Speed:     object.Speed,
						TypeV:     object.TypeV,
						TypeOf:    object.TypeOf,
						StTime:    object.StTime,
						StDay:     object.StDay,
						MStatus:   object.MStatus,
						MStatusV:  object.MStatusV,
						Category:  object.Category,
						CategoryV: object.CategoryV,
						PtType:    object.PtType,
						PtNum:     object.PtNum,
						PtRownr:   object.PtRownr,
						Entrust:   object.Entrust,
						Physician: object.Physician,
						EdTime:    object.EdTime,
						EdDay:     object.EdDay,
						Sender:    object.Sender,
						ExCycle:   idx,
					}
					obj.Gid = object.Contents[0].Madid
					obj.Contents = make([]MedicalAdviceContent, 0)
					obj.Contents = append(obj.Contents, object.Contents...)
					arrA = append(arrA, obj)
				}
			}

			object = MedicalAdviceResponse{
				Vid:       v.Vid,
				Pid:       v.Pid,
				Bed:       v.Bed,
				PName:     v.PName,
				Age:       v.Age,
				HospNum:   v.HospNum,
				Gender:    v.Gender,
				ExTime:    v.ExTime,
				ExDay:     v.ExDay,
				GroupNum:  v.GroupNum,
				Amount:    v.Amount,
				Frequency: v.Frequency,
				Times:     v.Times,
				Method:    v.Method,
				Speed:     v.Speed,
				TypeV:     v.TypeV,
				TypeOf:    v.TypeOf,
				StTime:    v.StTime,
				StDay:     v.StDay,
				MStatus:   v.MStatus,
				MStatusV:  v.MStatusV,
				Category:  v.Category,
				CategoryV: v.CategoryV,
				PtType:    v.PtType,
				PtNum:     v.PtNum,
				PtRownr:   v.PtRownr,
				Entrust:   v.Entrust,
				Physician: v.Physician,
				EdTime:    v.EdTime,
				EdDay:     v.EdDay,
				Sender:    v.Sender,
				ExCycle:   v.ExCycle,
			}
			object.Gid = v.Madid
			object.Contents = make([]MedicalAdviceContent, 1)
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage}
		}
	}
	// 继续拆分最后一次创建的object
	if length == 1 {
		arrA = append(arrA, object)
	} else {
		for idx := 1; idx <= object.Times; idx ++ {
			if idx == 1 {
				object.ExCycle = idx
				arrA = append(arrA, object)
			} else {
				obj := MedicalAdviceResponse{
					Vid:       object.Vid,
					Pid:       object.Pid,
					Bed:       object.Bed,
					PName:     object.PName,
					Age:       object.Age,
					HospNum:   object.HospNum,
					Gender:    object.Gender,
					ExTime:    object.ExTime,
					ExDay:     object.ExDay,
					GroupNum:  object.GroupNum,
					Amount:    object.Amount,
					Frequency: object.Frequency,
					Times:     object.Times,
					Method:    object.Method,
					Speed:     object.Speed,
					TypeV:     object.TypeV,
					TypeOf:    object.TypeOf,
					StTime:    object.StTime,
					StDay:     object.StDay,
					MStatus:   object.MStatus,
					MStatusV:  object.MStatusV,
					Category:  object.Category,
					CategoryV: object.CategoryV,
					PtType:    object.PtType,
					PtNum:     object.PtNum,
					PtRownr:   object.PtRownr,
					Entrust:   object.Entrust,
					Physician: object.Physician,
					EdTime:    object.EdTime,
					EdDay:     object.EdDay,
					Sender:    object.Sender,
					ExCycle:   idx,
				}
				obj.Gid = object.Contents[0].Madid
				obj.Contents = make([]MedicalAdviceContent, 0)
				obj.Contents = append(obj.Contents, object.Contents...)
				arrA = append(arrA, obj)
			}
		}
	}

	// 查询 组医嘱的执行记录
	length = len(arrA)
	//response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		//contentObj := resp.Contents[0]
		madid := resp.Gid
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
		}
		if exec.Madid != madid {
			exec.Madid = madid
			exec.ExNurse = ""
			exec.ExStatus = "未执行"
			exec.ExStatusV = 1
			exec.ExStep = ""
			exec.PtTimes = 0
			exec.PtStatus = "未打印"
		}
		if exec.ExStatusV == 0 {
			exec.ExStatusV = 1
		}

		//resp.ExStatusV = exec.ExStatusV
		//resp.ExStatus = exec.ExStatus
		//resp.ExNurse = exec.ExNurse
		//resp.ExStep = exec.ExStep
		//resp.PtTimes = exec.PtTimes
		//resp.PtStatus = exec.PtStatus

		if exec.ExStatusV == 1 {
			return 1
		}

		//if exec.Madid != madid {
		//	//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
		//	obj := MedicalAdviceItem{
		//		Vid:       resp.Vid,
		//		Madid:     madid,
		//		Pid:       resp.Pid,
		//		Bed:       resp.Bed,
		//		PName:     resp.PName,
		//		Gender:    resp.Gender,
		//		Age:       resp.Age,
		//		HospNum:   resp.HospNum,
		//		ExTime:    resp.ExTime.ParseToSecond(),
		//		GroupNum:  resp.GroupNum,
		//		Content:   content,
		//		Dosage:    dosage,
		//		Amount:    resp.Amount,
		//		Frequency: resp.Frequency,
		//		Times:     resp.Times,
		//		Method:    resp.Method,
		//		Speed:     resp.Speed,
		//		TypeV:     resp.TypeV,
		//		StTime:    resp.StTime.ParseToSecond(),
		//		MStatusV:  resp.MStatusV,
		//		Category:  resp.Category,
		//		CategoryV: resp.CategoryV,
		//		PtType:    resp.PtType,
		//		PtNum:     resp.PtNum,
		//		PtRownr:   resp.PtRownr,
		//		Entrust:   resp.Entrust,
		//		Physician: resp.Physician,
		//		EdTime:    resp.EdTime.ParseToSecond(),
		//		Sender:    resp.Sender,
		//		ExCycle:   resp.ExCycle,
		//		ExNurse:   resp.ExNurse,
		//		ExStatusV: resp.ExStatusV,
		//		ExStep:    resp.ExStep,
		//		PtTimes:   resp.PtTimes,
		//	}
		//	_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
		//	if err_in != nil {
		//		fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
		//		err_ss = err_in
		//		return make([]MedicalAdviceResponse, 0), err_ss
		//	}
		//} else {
		//	//	有记录
		//	sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
		//	_, err_up := fit.MySqlEngine().Exec(sqlStr)
		//	if err_up != nil {
		//		fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
		//		err_ss = err_up
		//		return make([]MedicalAdviceResponse, 0), err_ss
		//	}
		//}
	}

	return 0
}

func FetchNewMedicalAdvice(vid int64) ([]MedicalAdviceResponse, error) {
	//admissionTime := AdmissionTime{}
	//fit.SQLServerEngine().SQL("select VAA01, VAE11, BCK01C from VAE1 where VAA01 = ?", pid).Get(&admissionTime)
	//if admissionTime.VAA01 != pid {
	//	return make([]MedicalAdviceExecution, 0), fmt.Errorf("病人登记(VAE1)表中无此病人数据")
	//}
	//
	//today := DatetimeNow()
	//sqlStr := fmt.Sprintf("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF27,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAA01 = %d and a.VAF32 = '0' and a.VAF10 = '3' and ((a.VAF11 = 1 and a.VAF42 > '%s') or (a.VAF11 = 2 and a.VAF42 > '%s')) order by a.VAF36", pid, admissionTime.VAE11, today)
	//mAdvices := make([]MedicalAdvice, 0)
	//err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	//if err != nil {
	//	return make([]MedicalAdviceExecution, 0), err
	//}
	//
	//response := make([]MedicalAdviceExecution, 0)
	//for _, v := range mAdvices {
	//	isEx := IsExist{}
	//
	//	_, err := fit.MySqlEngine().SQL("select count(1) as Exist from AdviceStatus where Madid = ? and (Recordtime >= ? or State in (2,3))", v.VAF01, today).Get(&isEx)
	//	if err != nil {
	//		fit.Logger().LogError("***JK***", err.Error())
	//	}
	//	if isEx.Exist == 0 {
	//		status := MedicalAdviceExecutionStatus{}
	//		status.Process = "未执行"
	//		status.State = 1
	//
	//		execution := MedicalAdviceExecution{}
	//		execution.MedicalAdvice = v
	//		execution.MedicalAdviceExecutionStatus = status
	//		response = append(response, execution)
	//	}
	//}
	//
	//return response, err

	today := DatetimeNow()
	mAdvices := make([]MedicalAdviceModal, 0)
	err_db := fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT d.* FROM( SELECT CASE WHEN(( a2.BBX20 = 2 OR a2.BBX20 = 4 OR a2.BBX20 = 5) AND(a.BDA01 = '1' OR a.BDA01 = '2') AND a2.BDA01 = 'T' AND a2.BBX13 = '2') THEN 4 ELSE 0 END PtTypeV, a.VAF06 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 LEFT JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF06 = %d AND a.VAF32 = 0 AND a.VAF04 = 2 AND a.VAF10 = 3 AND((a.VAF11 = 1 AND b.VBI10 > c.VAE11) OR(a.VAF11 = 2 AND b.VBI10 > '%s'))) d ORDER BY d.TypeV, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", vid, today)).Find(&mAdvices)
	if err_db != nil {
		fit.Logger().LogError("**JP**", err_db.Error())
	}

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 {
		return arrA, err_db
	}

	// 取出第一条医嘱数据
	temp := mAdvices[0]
	temp.ExDay = temp.ExTime.ParseToSecond()
	temp.StDay = temp.StTime.ParseToMinute()
	temp.EdDay = temp.EdTime.ParseToMinute()

	// 用temp的数据实例第一个object
	object := MedicalAdviceResponse{
		Vid:       temp.Vid,
		Pid:       temp.Pid,
		Bed:       temp.Bed,
		PName:     temp.PName,
		Age:       temp.Age,
		HospNum:   temp.HospNum,
		Gender:    temp.Gender,
		ExTime:    temp.ExTime,
		ExDay:     temp.ExDay,
		GroupNum:  temp.GroupNum,
		Amount:    temp.Amount,
		Frequency: temp.Frequency,
		Times:     temp.Times,
		Method:    temp.Method,
		Speed:     temp.Speed,
		TypeV:     temp.TypeV,
		TypeOf:    temp.TypeOf,
		StTime:    temp.StTime,
		StDay:     temp.StDay,
		MStatus:   temp.MStatus,
		MStatusV:  temp.MStatusV,
		Category:  temp.Category,
		CategoryV: temp.CategoryV,
		PtType:    temp.PtType,
		PtTypeV:   temp.PtTypeV,
		PtNum:     temp.PtNum,
		PtRownr:   temp.PtRownr,
		Entrust:   temp.Entrust,
		Physician: temp.Physician,
		EdTime:    temp.EdTime,
		EdDay:     temp.EdDay,
		Sender:    temp.Sender,
		ExCycle:   temp.ExCycle,
	}
	object.Gid = temp.Madid
	object.Contents = make([]MedicalAdviceContent, 1)
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content," ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage})
		} else {
			//	按次拆分
			for idx := 1; idx <= object.Times; idx ++ {
				if idx == 1 {
					object.ExCycle = idx
					arrA = append(arrA, object)
				} else {
					obj := MedicalAdviceResponse{
						Vid:       object.Vid,
						Pid:       object.Pid,
						Bed:       object.Bed,
						PName:     object.PName,
						Age:       object.Age,
						HospNum:   object.HospNum,
						Gender:    object.Gender,
						ExTime:    object.ExTime,
						ExDay:     object.ExDay,
						GroupNum:  object.GroupNum,
						Amount:    object.Amount,
						Frequency: object.Frequency,
						Times:     object.Times,
						Method:    object.Method,
						Speed:     object.Speed,
						TypeV:     object.TypeV,
						TypeOf:    object.TypeOf,
						StTime:    object.StTime,
						StDay:     object.StDay,
						MStatus:   object.MStatus,
						MStatusV:  object.MStatusV,
						Category:  object.Category,
						CategoryV: object.CategoryV,
						PtType:    object.PtType,
						PtTypeV:   object.PtTypeV,
						PtNum:     object.PtNum,
						PtRownr:   object.PtRownr,
						Entrust:   object.Entrust,
						Physician: object.Physician,
						EdTime:    object.EdTime,
						EdDay:     object.EdDay,
						Sender:    object.Sender,
						ExCycle:   idx,
					}
					obj.Gid = object.Contents[0].Madid
					obj.Contents = make([]MedicalAdviceContent, 0)
					obj.Contents = append(obj.Contents, object.Contents...)
					arrA = append(arrA, obj)
				}
			}

			object = MedicalAdviceResponse{
				Vid:       v.Vid,
				Pid:       v.Pid,
				Bed:       v.Bed,
				PName:     v.PName,
				Age:       v.Age,
				HospNum:   v.HospNum,
				Gender:    v.Gender,
				ExTime:    v.ExTime,
				ExDay:     v.ExDay,
				GroupNum:  v.GroupNum,
				Amount:    v.Amount,
				Frequency: v.Frequency,
				Times:     v.Times,
				Method:    v.Method,
				Speed:     v.Speed,
				TypeV:     v.TypeV,
				TypeOf:    v.TypeOf,
				StTime:    v.StTime,
				StDay:     v.StDay,
				MStatus:   v.MStatus,
				MStatusV:  v.MStatusV,
				Category:  v.Category,
				CategoryV: v.CategoryV,
				PtType:    v.PtType,
				PtTypeV:   v.PtTypeV,
				PtNum:     v.PtNum,
				PtRownr:   v.PtRownr,
				Entrust:   v.Entrust,
				Physician: v.Physician,
				EdTime:    v.EdTime,
				EdDay:     v.EdDay,
				Sender:    v.Sender,
				ExCycle:   v.ExCycle,
			}
			object.Gid = v.Madid
			object.Contents = make([]MedicalAdviceContent, 1)
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage}
		}
	}
	// 继续拆分最后一次创建的object
	if length == 1 {
		arrA = append(arrA, object)
	} else {
		for idx := 1; idx <= object.Times; idx ++ {
			if idx == 1 {
				object.ExCycle = idx
				arrA = append(arrA, object)
			} else {
				obj := MedicalAdviceResponse{
					Vid:       object.Vid,
					Pid:       object.Pid,
					Bed:       object.Bed,
					PName:     object.PName,
					Age:       object.Age,
					HospNum:   object.HospNum,
					Gender:    object.Gender,
					ExTime:    object.ExTime,
					ExDay:     object.ExDay,
					GroupNum:  object.GroupNum,
					Amount:    object.Amount,
					Frequency: object.Frequency,
					Times:     object.Times,
					Method:    object.Method,
					Speed:     object.Speed,
					TypeV:     object.TypeV,
					TypeOf:    object.TypeOf,
					StTime:    object.StTime,
					StDay:     object.StDay,
					MStatus:   object.MStatus,
					MStatusV:  object.MStatusV,
					Category:  object.Category,
					CategoryV: object.CategoryV,
					PtType:    object.PtType,
					PtTypeV:   object.PtTypeV,
					PtNum:     object.PtNum,
					PtRownr:   object.PtRownr,
					Entrust:   object.Entrust,
					Physician: object.Physician,
					EdTime:    object.EdTime,
					EdDay:     object.EdDay,
					Sender:    object.Sender,
					ExCycle:   idx,
				}
				obj.Gid = object.Contents[0].Madid
				obj.Contents = make([]MedicalAdviceContent, 0)
				obj.Contents = append(obj.Contents, object.Contents...)
				arrA = append(arrA, obj)
			}
		}
	}

	// 查询 组医嘱的执行记录
	length = len(arrA)
	response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		//contentObj := resp.Contents[0]
		madid := resp.Gid
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
		}
		if exec.Madid != madid {
			exec.Madid = madid
			exec.ExNurse = ""
			exec.ExStatus = "未执行"
			exec.ExStatusV = 1
			exec.ExStep = ""
			exec.PtTimes = 0
			exec.PtStatus = "未打印"
		}
		if exec.ExStatusV == 0 {
			exec.ExStatusV = 1
		}

		resp.ExStatusV = exec.ExStatusV
		resp.ExStatus = exec.ExStatus
		resp.ExNurse = exec.ExNurse
		resp.ExStep = exec.ExStep
		resp.PtTimes = exec.PtTimes
		resp.PtStatus = exec.PtStatus

		if exec.ExStatusV == 1 {
			response = append(response, *resp)
		}

		//if exec.Madid != madid {
		//	//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
		//	obj := MedicalAdviceItem{
		//		Vid:       resp.Vid,
		//		Madid:     madid,
		//		Pid:       resp.Pid,
		//		Bed:       resp.Bed,
		//		PName:     resp.PName,
		//		Gender:    resp.Gender,
		//		Age:       resp.Age,
		//		HospNum:   resp.HospNum,
		//		ExTime:    resp.ExTime.ParseToSecond(),
		//		GroupNum:  resp.GroupNum,
		//		Content:   content,
		//		Dosage:    dosage,
		//		Amount:    resp.Amount,
		//		Frequency: resp.Frequency,
		//		Times:     resp.Times,
		//		Method:    resp.Method,
		//		Speed:     resp.Speed,
		//		TypeV:     resp.TypeV,
		//		StTime:    resp.StTime.ParseToSecond(),
		//		MStatusV:  resp.MStatusV,
		//		Category:  resp.Category,
		//		CategoryV: resp.CategoryV,
		//		PtType:    resp.PtType,
		//		PtNum:     resp.PtNum,
		//		PtRownr:   resp.PtRownr,
		//		Entrust:   resp.Entrust,
		//		Physician: resp.Physician,
		//		EdTime:    resp.EdTime.ParseToSecond(),
		//		Sender:    resp.Sender,
		//		ExCycle:   resp.ExCycle,
		//		ExNurse:   resp.ExNurse,
		//		ExStatusV: resp.ExStatusV,
		//		ExStep:    resp.ExStep,
		//		PtTimes:   resp.PtTimes,
		//	}
		//	_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
		//	if err_in != nil {
		//		fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
		//		err_ss = err_in
		//		return make([]MedicalAdviceResponse, 0), err_ss
		//	}
		//} else {
		//	//	有记录
		//	sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
		//	_, err_up := fit.MySqlEngine().Exec(sqlStr)
		//	if err_up != nil {
		//		fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
		//		err_ss = err_up
		//		return make([]MedicalAdviceResponse, 0), err_ss
		//	}
		//}
	}
	return response, err_db
}

/*JP 是否存在已停医嘱*/
func IsExistFinishedMedicalAdvice(vid int64, did int) int {
	// 获取已停的长嘱
	exist := IsExist{}
	_, err_db := fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT CASE WHEN COUNT(1) > 0 THEN 1 ELSE 0 END Exist FROM VAF2 a WHERE a.VAF06 = %d AND a.VAF32 = 0 AND a.VAF04 = 0 AND a.VAF11 = 1 AND a.VAF10 = 8 AND a.VAF47 > '%s'", vid, DatetimeNow())).Get(&exist)
	if err_db != nil {
		fit.Logger().LogError("**JP**", err_db)
	}
	return exist.Exist
}

/*获取已停医嘱 JKWARN -- LEFT JOIN BBX1 c ON c.BBX01 = a.BBX01 可能要改成 LEFT JOIN BBX1 c ON c.BBX01 = a1.BBX01  */
func FetchFinishedMedicalAdvice(vid int64) ([]MedicalAdviceResponse, error) {
	//mAdvices := make([]MedicalAdvice, 0)
	//if vid == 0 {
	//	return mAdvices, fmt.Errorf("病人登记(VAE1)表中无此病人数据")
	//}
	//
	//sqlStr := fmt.Sprintf("SELECT a.VAF01, a.VAA01, a.VAF10, a.VAF11, a.BDA01, a.VAF19, a.VAF22, a.BBX01, b.VAF22 AS Method, a.VAF23, a.VAF26, a.VAF27, a.VAF36, a.BCE03A, c.BBX20, a.VAF60 FROM( VAF2 a LEFT JOIN VAF2 b ON a.VAF01A = b.VAF01) LEFT JOIN BBX1 c ON c.BBX01 = a.BBX01 WHERE a.VAF06 = '%d' AND a.VAF32 = 0 AND VAF04 = 2 AND a.VAF11 = 1 AND a.VAF10 = 8 AND a.VAF47 > '%s' ORDER BY a.VAF36", vid, DatetimeNow())
	//err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	//return mAdvices, err

	today := DatetimeNow()
	mAdvices := make([]MedicalAdviceModal, 0)
	err_db := fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT d.* FROM( SELECT a.VAF06 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 LEFT JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF06 = %d AND a.VAF32 = 0 AND a.VAF04 = 2 AND a.VAF11 = 1 AND a.VAF10 = 8 AND a.VAF47 > '%s') d ORDER BY d.TypeV, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", vid, today)).Find(&mAdvices)
	if err_db != nil {
		fit.Logger().LogError("**JP**", err_db.Error())
	}

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 {
		return arrA, err_db
	}

	// 取出第一条医嘱数据
	temp := mAdvices[0]
	temp.ExDay = temp.ExTime.ParseToSecond()
	temp.StDay = temp.StTime.ParseToMinute()
	temp.EdDay = temp.EdTime.ParseToMinute()

	// 用temp的数据实例第一个object
	object := MedicalAdviceResponse{
		Vid:       temp.Vid,
		Pid:       temp.Pid,
		Bed:       temp.Bed,
		PName:     temp.PName,
		Age:       temp.Age,
		HospNum:   temp.HospNum,
		Gender:    temp.Gender,
		ExTime:    temp.ExTime,
		ExDay:     temp.ExDay,
		GroupNum:  temp.GroupNum,
		Amount:    temp.Amount,
		Frequency: temp.Frequency,
		Times:     temp.Times,
		Method:    temp.Method,
		Speed:     temp.Speed,
		TypeV:     temp.TypeV,
		TypeOf:    temp.TypeOf,
		StTime:    temp.StTime,
		StDay:     temp.StDay,
		MStatus:   temp.MStatus,
		MStatusV:  temp.MStatusV,
		Category:  temp.Category,
		CategoryV: temp.CategoryV,
		PtType:    temp.PtType,
		PtNum:     temp.PtNum,
		PtRownr:   temp.PtRownr,
		Entrust:   temp.Entrust,
		Physician: temp.Physician,
		EdTime:    temp.EdTime,
		EdDay:     temp.EdDay,
		Sender:    temp.Sender,
		ExCycle:   temp.ExCycle,
	}
	object.Gid = temp.Madid
	object.Contents = make([]MedicalAdviceContent, 1)
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content," ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage})
		} else {
			//	按次拆分
			for idx := 1; idx <= object.Times; idx ++ {
				if idx == 1 {
					object.ExCycle = idx
					arrA = append(arrA, object)
				} else {
					obj := MedicalAdviceResponse{
						Vid:       object.Vid,
						Pid:       object.Pid,
						Bed:       object.Bed,
						PName:     object.PName,
						Age:       object.Age,
						HospNum:   object.HospNum,
						Gender:    object.Gender,
						ExTime:    object.ExTime,
						ExDay:     object.ExDay,
						GroupNum:  object.GroupNum,
						Amount:    object.Amount,
						Frequency: object.Frequency,
						Times:     object.Times,
						Method:    object.Method,
						Speed:     object.Speed,
						TypeV:     object.TypeV,
						TypeOf:    object.TypeOf,
						StTime:    object.StTime,
						StDay:     object.StDay,
						MStatus:   object.MStatus,
						MStatusV:  object.MStatusV,
						Category:  object.Category,
						CategoryV: object.CategoryV,
						PtType:    object.PtType,
						PtNum:     object.PtNum,
						PtRownr:   object.PtRownr,
						Entrust:   object.Entrust,
						Physician: object.Physician,
						EdTime:    object.EdTime,
						EdDay:     object.EdDay,
						Sender:    object.Sender,
						ExCycle:   idx,
					}
					obj.Gid = object.Contents[0].Madid
					obj.Contents = make([]MedicalAdviceContent, 0)
					obj.Contents = append(obj.Contents, object.Contents...)
					arrA = append(arrA, obj)
				}
			}

			object = MedicalAdviceResponse{
				Vid:       v.Vid,
				Pid:       v.Pid,
				Bed:       v.Bed,
				PName:     v.PName,
				Age:       v.Age,
				HospNum:   v.HospNum,
				Gender:    v.Gender,
				ExTime:    v.ExTime,
				ExDay:     v.ExDay,
				GroupNum:  v.GroupNum,
				Amount:    v.Amount,
				Frequency: v.Frequency,
				Times:     v.Times,
				Method:    v.Method,
				Speed:     v.Speed,
				TypeV:     v.TypeV,
				TypeOf:    v.TypeOf,
				StTime:    v.StTime,
				StDay:     v.StDay,
				MStatus:   v.MStatus,
				MStatusV:  v.MStatusV,
				Category:  v.Category,
				CategoryV: v.CategoryV,
				PtType:    v.PtType,
				PtNum:     v.PtNum,
				PtRownr:   v.PtRownr,
				Entrust:   v.Entrust,
				Physician: v.Physician,
				EdTime:    v.EdTime,
				EdDay:     v.EdDay,
				Sender:    v.Sender,
				ExCycle:   v.ExCycle,
			}
			object.Gid = v.Madid
			object.Contents = make([]MedicalAdviceContent, 1)
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content," ")[0], v.Dosage}
		}
	}
	// 继续拆分最后一次创建的object
	if length == 1 {
		arrA = append(arrA, object)
	} else {
		for idx := 1; idx <= object.Times; idx ++ {
			if idx == 1 {
				object.ExCycle = idx
				arrA = append(arrA, object)
			} else {
				obj := MedicalAdviceResponse{
					Vid:       object.Vid,
					Pid:       object.Pid,
					Bed:       object.Bed,
					PName:     object.PName,
					Age:       object.Age,
					HospNum:   object.HospNum,
					Gender:    object.Gender,
					ExTime:    object.ExTime,
					ExDay:     object.ExDay,
					GroupNum:  object.GroupNum,
					Amount:    object.Amount,
					Frequency: object.Frequency,
					Times:     object.Times,
					Method:    object.Method,
					Speed:     object.Speed,
					TypeV:     object.TypeV,
					TypeOf:    object.TypeOf,
					StTime:    object.StTime,
					StDay:     object.StDay,
					MStatus:   object.MStatus,
					MStatusV:  object.MStatusV,
					Category:  object.Category,
					CategoryV: object.CategoryV,
					PtType:    object.PtType,
					PtNum:     object.PtNum,
					PtRownr:   object.PtRownr,
					Entrust:   object.Entrust,
					Physician: object.Physician,
					EdTime:    object.EdTime,
					EdDay:     object.EdDay,
					Sender:    object.Sender,
					ExCycle:   idx,
				}
				obj.Gid = object.Contents[0].Madid
				obj.Contents = make([]MedicalAdviceContent, 0)
				obj.Contents = append(obj.Contents, object.Contents...)
				arrA = append(arrA, obj)
			}
		}
	}

	// 查询 组医嘱的执行记录
	length = len(arrA)
	//response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		//contentObj := resp.Contents[0]
		madid := resp.Gid
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
		}
		if exec.Madid != madid {
			exec.Madid = madid
			exec.ExNurse = ""
			exec.ExStatus = "未执行"
			exec.ExStatusV = 1
			exec.ExStep = ""
			exec.PtTimes = 0
			exec.PtStatus = "未打印"
		}
		if exec.ExStatusV == 0 {
			exec.ExStatusV = 1
		}

		resp.ExStatusV = exec.ExStatusV
		resp.ExStatus = exec.ExStatus
		resp.ExNurse = exec.ExNurse
		resp.ExStep = exec.ExStep
		resp.PtTimes = exec.PtTimes
		resp.PtStatus = exec.PtStatus

		//if exec.Madid != madid {
		//	//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
		//	obj := MedicalAdviceItem{
		//		Vid:       resp.Vid,
		//		Madid:     madid,
		//		Pid:       resp.Pid,
		//		Bed:       resp.Bed,
		//		PName:     resp.PName,
		//		Gender:    resp.Gender,
		//		Age:       resp.Age,
		//		HospNum:   resp.HospNum,
		//		ExTime:    resp.ExTime.ParseToSecond(),
		//		GroupNum:  resp.GroupNum,
		//		Content:   content,
		//		Dosage:    dosage,
		//		Amount:    resp.Amount,
		//		Frequency: resp.Frequency,
		//		Times:     resp.Times,
		//		Method:    resp.Method,
		//		Speed:     resp.Speed,
		//		TypeV:     resp.TypeV,
		//		StTime:    resp.StTime.ParseToSecond(),
		//		MStatusV:  resp.MStatusV,
		//		Category:  resp.Category,
		//		CategoryV: resp.CategoryV,
		//		PtType:    resp.PtType,
		//		PtNum:     resp.PtNum,
		//		PtRownr:   resp.PtRownr,
		//		Entrust:   resp.Entrust,
		//		Physician: resp.Physician,
		//		EdTime:    resp.EdTime.ParseToSecond(),
		//		Sender:    resp.Sender,
		//		ExCycle:   resp.ExCycle,
		//		ExNurse:   resp.ExNurse,
		//		ExStatusV: resp.ExStatusV,
		//		ExStep:    resp.ExStep,
		//		PtTimes:   resp.PtTimes,
		//	}
		//	_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
		//	if err_in != nil {
		//		fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
		//		err_ss = err_in
		//		return make([]MedicalAdviceResponse, 0), err_ss
		//	}
		//} else {
		//	//	有记录
		//	sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
		//	_, err_up := fit.MySqlEngine().Exec(sqlStr)
		//	if err_up != nil {
		//		fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
		//		err_ss = err_up
		//		return make([]MedicalAdviceResponse, 0), err_ss
		//	}
		//}
	}
	return arrA, err_db
}
