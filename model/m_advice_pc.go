package model

import (
	"fit"
	"fmt"
)

type MedicalAdvicePrint struct {
	Madid        int64                  `json:"Madid"`        // 医嘱ID
	Pid          int64                  `json:"Pid"`          // 病人ID
	PatientName  string                 `json:"PatientName"`  // 病人姓名
	Bed          string                 `json:"Bed"`          // 病人床位
	Status       int                    `json:"Status"`       // 医嘱状态
	TypeOf       int                    `json:"TypeOf"`       // 医嘱类型
	Category     string                 `json:"Category"`     // 医嘱类别
	Dosage       string                 `json:"Dosage"`       // 用量、剂量
	Count        string                 `json:"Count"`        // 数量
	Content      string                 `json:"Content"`      // 医嘱内容
	Method       string                 `json:"Method"`       // 医嘱用法
	Times        int                    `json:"Times"`        // 医嘱执行次数(需要打印的次数)
	ExTime       DatetimeWithoutSeconds `json:"-"`            // 执行开始时间
	ExexTime     string                 `json:"ExexTime"`     // 执行开始时间
	Speed        string                 `json:"Speed"`        // 滴速
	PrintType    int                    `json:"PrintType"`    // 打印单类型
	PrintState   string                 `json:"PrintState"`   // 打印状态
	PrintOrdinal int                    `json:"PrintOrdinal"` // 当前打印次数(序号)
}

type PCMedicalAdviceExec struct {
	Madid       int64                  `json:"Madid"`       // 医嘱ID
	Pid         int64                  `json:"Pid"`         // 病人ID
	Bed         string                 `json:"Bed"`         // 病人床位
	PatientName string                 `json:"PatientName"` // 病人姓名
	Category    string                 `json:"Category"`    // 医嘱类别
	Content     string                 `json:"Content"`     // 医嘱内容
	Dosage      string                 `json:"Dosage"`      // 用量、剂量
	Method      string                 `json:"Method"`      // 医嘱用法
	TypeOf      int                    `json:"TypeOf"`      // 医嘱类型
	Entrust     string                 `json:"Entrust"`     // 医师嘱托
	StTime      DatetimeWithoutSeconds `json:"-"`           // 执行开始时间
	StartTime   string                 `json:"StartTime"`   // 执行开始时间
	Frequency   string                 `json:"Frequency"`   // 次数/执行频次
	State       int                    `json:"State"`       // 医嘱执行状态
	Process     string                 `json:"Process"`     // 执行步骤
	Nursename   string                 `json:"Nursename"`   // 执行护士
}

type PCMedicalAdvice struct {
	Madid       int64                  `json:"Madid"`       // 医嘱ID
	Pid         int64                  `json:"Pid"`         // 病人ID
	Bed         string                 `json:"Bed"`         // 病人床位
	PatientName string                 `json:"PatientName"` // 病人姓名
	Category    string                 `json:"Category"`    // 医嘱类别
	Content     string                 `json:"Content"`     // 医嘱内容
	Dosage      string                 `json:"Dosage"`      // 用量、剂量
	Count       string                 `json:"Count"`       // 数量
	Method      string                 `json:"Method"`      // 医嘱用法
	Speed       string                 `json:"Speed"`       // 滴速
	TypeOf      int                    `json:"TypeOf"`      // 医嘱类型
	Entrust     string                 `json:"Entrust"`     // 医师嘱托
	Physician   string                 `json:"Physician"`   // 开嘱医师// 开嘱医师
	StTime      DatetimeWithoutSeconds `json:"StartTime"`   // 执行开始时间
	StartTime   string                 `json:"-"`           // 执行开始时间
	EdTime      DatetimeWithoutSeconds `json:"EndTime"`     // 停嘱/作废时间
	EndTime     string                 `json:"-"`           // 停嘱/作废时间
	Status      int                    `json:"Status"`      // 医嘱状态
	Nurse       string                 `json:"Nurse"`       // 停嘱校对护士
	CkTime      DatetimeWithoutSeconds `json:"CheckTime"`   // 执行停嘱时间/停嘱校对时间
	CheckTime   string                 `json:"-"`           // 执行停嘱时间/停嘱校对时间
}

/*获取医嘱执行记录*/
func FetchMedicalAdviceExecutionRecordForPc(madid int64) ([]MedicalAdviceExecutionRecord, error) {
	response := make([]MedicalAdviceExecutionRecord, 0)
	err := fit.MySqlEngine().SQL("select * from AdviceDetail where Madid = ?", madid).Find(&response)
	return response, err
}

/*医嘱拆分查询*/
func SearchMedicalAdvicesForSplitting(startTime, endTime, pidArr, printTypeArr string, typeOf, print, did int) ([]MedicalAdvicePrint, error) {
	array := make([]MedicalAdvicePrint, 0)

	typeOfStr := ""
	if typeOf != 0 {
		typeOfStr = fmt.Sprintf(" and a.VAF11 = %d", typeOf)
	}

	sqlStr := fmt.Sprintf("select a.VAF01 as Madid,a.VAA01 as Pid,v.VAA05 as PatientName,v.BCQ04 as Bed,a.VAF10 as Status,a.VAF11 as TypeOf,a.BDA01 as Category,a.VAF19 as Dosage,a.VAF21 as Count,a.VAF22 as Content,b.VAF22 as Method,a.VAF27 as Times,a.VAF36 as ExTime,c.BBX20 PrintType,a.VAF60 Speed from ((VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01) left join VAA1 v on v.VAA01 = a.VAA01 where a.VAA01 in (%s) and a.VAF32 = 0 and a.BDA01 != '0' and a.VAF42 BETWEEN '%s' AND '%s'%s and c.BBX20 in (%s) order by a.VAF36", pidArr, startTime, endTime, typeOfStr, printTypeArr)
	err := fit.SQLServerEngine().SQL(sqlStr).Find(&array)
	fit.Logger().LogDebug("***JK***",sqlStr)
	if err != nil {
		fit.Logger().LogError("**JK**", err.Error())
		return array, err
	}
	response := make([]MedicalAdvicePrint, 0)
	for _, v := range array {
		_, err = fit.MySqlEngine().SQL("select if(Ordinal is not null,Ordinal,0) as PrintOrdinal, if(count(1) > 0,'已打印','未打印') as PrintState from AdvicePrint where Madid = ?", v.Madid).Get(&v)
		if err != nil {
			fit.Logger().LogError("**JK**", err.Error())
		}
		v.ExexTime = v.ExTime.ParseToMinute()
		if print == 0 {
			// all
			response = append(response, v)
		} else if print == 1 && v.PrintOrdinal > 0 {
			// printed
			response = append(response, v)
		} else {
			// unprinted
			response = append(response, v)
		}
	}

	return response, err
}

func SearchMedicalAdviceExecutionForPC(typeOf, status, did int, category, pids, st, et string) ([]PCMedicalAdviceExec, error) {
	condition_type := ""
	if typeOf != 0 {
		// 长期和临时
		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
	}

	condition_catg := ""
	if category != "0" {
		// 医嘱类别
		condition_catg = fmt.Sprintf(" and a.BDA01 = '%s'", category)
	}

	condition_time := " and a.VAF42 > v.VAA73"
	if st != "all" && et != "all" {
		condition_time = fmt.Sprintf(" and a.VAF42 BETWEEN '%s' AND '%s'", st, et)
	}

	mAdvices := make([]PCMedicalAdviceExec, 0)
	sqlStr := fmt.Sprintf("select a.VAF01 as Madid,a.VAA01 as Pid,v.VAA05 as PatientName,v.BCQ04 as Bed,d.BDA02 as Category,a.VAF22 as Content,a.VAF19 as Dosage,b.VAF22 as Method,a.VAF11 as TypeOf,a.VAF23 as Entrust,a.VAF36 as StTime,a.VAF26 as Frequency from (((VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01) left join BDA1 d on a.BDA01 = d.BDA01) left join VAA1 v on v.VAA01 = a.VAA01 where a.VAA01 in (%s) and a.VAF32 = 0 and a.VAF10 = 3%s%s%s order by a.VAF36", pids, condition_type, condition_catg, condition_time)
	err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)

	response := make([]PCMedicalAdviceExec, 0)
	for _, obj := range mAdvices {
		obj.StartTime = obj.StTime.ParseToMinute()
		sqlStr = fmt.Sprintf("select Nursename,State,Process from AdviceStatus where Madid = %d", obj.Madid)
		_, err = fit.MySqlEngine().SQL(sqlStr).Get(&obj)
		if err != nil {
			fit.Logger().LogError("***JK***", err.Error())
		}
		if status == 0 {
			response = append(response, obj)
		} else if status == 1 {
			if obj.State == 0 || obj.State == 1 {
				obj.State = 1
				obj.Process = "未执行"
				response = append(response, obj)
			}
		} else if status == obj.State {
			response = append(response, obj)
		}
	}

	return response, err
}

func SearchMedicalAdviceForPC(typeOf, status, did int, category, pids, st, et string) ([]PCMedicalAdvice, error) {
	condition_type := ""
	if typeOf != 0 {
		// 长期和临时
		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
	}

	condition_catg := ""
	if category != "0" {
		// 医嘱类别
		condition_catg = fmt.Sprintf(" and a.BDA01 = '%s'", category)
	}

	// 医嘱状态所有
	condition_state := " and a.VAF10 in (3,4,8,9)"
	if status != 0 {
		// 医嘱状态 未停、已撤销、已停
		if status == 8 {
			condition_state = " and a.VAF10 in (8,9)"
		} else {
			condition_state = fmt.Sprintf(" and a.VAF10 = '%d'", status)
		}
	}

	condition_time := " and a.VAF42 > v.VAA73"
	if st != "all" && et != "all" {
		condition_time = fmt.Sprintf(" and a.VAF42 BETWEEN '%s' AND '%s'", st, et)
	}

	mAdvices := make([]PCMedicalAdvice, 0)
	sqlStr := fmt.Sprintf("select a.VAF01 as Madid,a.VAA01 as Pid,v.VAA05 as PatientName,v.BCQ04 as Bed,d.BDA02 as Category,a.VAF22 as Content,a.VAF19 as Dosage,a.VAF21 as Count,b.VAF22 as Method,a.VAF60 Speed,a.VAF11 as TypeOf,a.VAF23 as Entrust,a.BCE03A as Physician,a.VAF36 as StTime,a.VAF47 as EdTime,a.VAF10 as Status,a.BCE03F as Nurse,a.VAF50 as CkTime from (((VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01) left join BDA1 d on a.BDA01 = d.BDA01) left join VAA1 v on v.VAA01 = a.VAA01 where a.VAA01 in (%s) and a.VAF32 = 0%s%s%s%s order by a.VAF36", pids, condition_type, condition_catg, condition_state, condition_time)
	err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	return mAdvices, err
}
