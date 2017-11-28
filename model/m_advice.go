package model

import (
	"fit"
	"time"
	"fmt"
)


/*PDA医嘱 查询结果*/
type MedicalAdvice struct {
	VAF01  int64                  `json:"madid"`              // 医嘱ID
	VAA01  int64                  `json:"pid"`                // 病人ID
	VAF10  int                    `json:"status"`             // 状态
	VAF11  int                    `json:"type"`               // 长期医嘱/临时医嘱
	BDA01  string                 `json:"category"`           // 医嘱类别，诊疗类型
	VAF19  string                 `json:"dosage"`             // 用量
	VAF22  string                 `json:"content"`            // 医嘱内容
	BBX01  int                    `json:"-"`                  // 诊疗类型BBX05，可能跟用药途径关联
	Method string                 `json:"method"`             // 用药途径
	VAF23  string                 `json:"entrust"`            // 医师嘱托
	VAF26  string                 `json:"time"`               // 次数/执行频次
	VAF36  DatetimeWithoutSeconds `json:"starttime"`          // 开始执行时间
	BCE03A string                 `json:"creation_physician"` // 开嘱医师
	BBX20  int                    `json:"classify"`           // 按打印单分类
	VAF60  string                 `json:"speed"`              // 滴速
	//VAF42  DatetimeWithoutSeconds `json:"createdtime"`        // 开嘱时间
}

/*PDA医嘱执行 查询结果*/
type MedicalAdviceExecution struct {
	MedicalAdvice                `xorm:"extends"`
	MedicalAdviceExecutionStatus `xorm:"extends"`
}

/*医嘱执行状态*/
type MedicalAdviceExecutionStatus struct {
	Madid   int64  `json:"-"`
	State   int    `json:"state"`   // 执行状态
	Process string `json:"process"` // 执行步骤
}

/*医嘱执行详细*/
type MedicalAdviceExecutionDetail struct {
	MedicalAdvice                          `xorm:"extends"`
	MedicalAdviceExecutionStatus           `xorm:"extends"`
	Records []MedicalAdviceExecutionRecord `json:"records" xorm:"extends"`
	Master  []VAA1                         `json:"master"` // 病人信息
	Desc    string                         `json:"desc"`   // 关于医嘱的描述
}

type MedicalAdviceDup struct {
	VAF01 int64 // 医嘱记录ID
}

/*医嘱执行状态*/
type AdviceStatus struct {
	Patientid  int64
	Madid      int64
	State      int
	Recordtime string
	Nurseid    int
	Nursename  string
	Period     int
	Process    string
}

/*医嘱执行明细*/
type AdviceDetail struct {
	ExecTime  string `json:"exectime"`  // 执行时间
	Patientid int64  `json:"pid"`       // 病人ID
	Nurseid   int    `json:"nid"`       // 护士ID
	Period    int    `json:"period"`    // 周期
	Madid     int64  `json:"madid"`     // 医嘱ID
	Nursename string `json:"nursename"` // 执行护士
	Process   string `json:"process"`   // 执行步骤
}

/*医嘱执行明细*/
type MedicalAdviceExecutionRecord struct {
	Patientid int64                  `json:"pid"`       // 病人ID
	Nurseid   int                    `json:"nid"`       // 护士ID
	Period    int                    `json:"period"`    // 周期
	Madid     int64                  `json:"madid"`     // 医嘱ID
	Nursename string                 `json:"nursename"` // 执行护士
	Process   string                 `json:"process"`   // 执行步骤
	ExecTime  DatetimeWithoutSeconds `json:"exectime"`  // 执行时间
}

/*JP 今天的起始时间*/
func DatetimeNow() string {
	t := time.Now()
	dateNow := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
	return dateNow
}

/*JP 登记表中的入院时间*/
type AdmissionTime struct {
	VAA01 int64  `json:"patientId"`     // 病人ID
	VAE11 string `json:"admissionTime"` // 入院时间
}

/*JP PDA医嘱查询*/
func SearchMedicalAdvice(typeOf, status int, pid int64, category string) ([]MedicalAdvice, error) {
	condition_type := ""
	if typeOf != 0 {
		// 长期和临时
		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
	}

	condition_catg := ""
	if category != "0" {
		// 医嘱类别
		condition_catg = " and a.BDA01 = '" + category + "'"
	}

	// 医嘱状态所有
	condition_stat := " and a.VAF10 in (3,4,8)"
	if status != 0 {
		// 医嘱状态 未停、已撤销、已停
		condition_stat = fmt.Sprintf(" and a.VAF10 = '%d'", status)
	}
	// 时间  VAF36
	admissionTime := AdmissionTime{}
	mAdvices := make([]MedicalAdvice, 0)

	fit.SQLServerEngine().SQL("select VAA01, VAE11 from VAE1 where VAA01 = ?", pid).Get(&admissionTime)
	if admissionTime.VAA01 != pid {
		return mAdvices, fmt.Errorf("病人登记(VAE1)表中无此病人数据")
	}

	sqlStr := fmt.Sprintf("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAA01 = '%d' and a.VAF32 = 0 and a.VAF42 > '%s'%s%s%s order by a.VAF36", pid, admissionTime.VAE11, condition_type, condition_stat, condition_catg)
	//fit.SQLServerEngine().ShowSQL(true)
	err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	//fit.SQLServerEngine().ShowSQL(false)
	return mAdvices, err
}

/*医嘱执行查询*/
func SearchMedicalAdviceExecution(typeOf, status int, pid int64, category string) ([]MedicalAdviceExecution, error) {
	condition_type := ""
	if typeOf != 0 {
		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
	}

	condition_catg := ""
	if category != "0" {
		condition_catg = " and a.BDA01 = '" + category + "'"
	}

	admissionTime := AdmissionTime{}
	mAdvices := make([]MedicalAdvice, 0)

	fit.SQLServerEngine().SQL("select VAA01, VAE11 from VAE1 where VAA01 = ?", pid).Get(&admissionTime)
	if admissionTime.VAA01 != pid {
		return make([]MedicalAdviceExecution, 0), fmt.Errorf("病人登记(VAE1)表中无此病人数据")
	}

	sqlStr := fmt.Sprintf("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAA01 = '%d' and a.VAF32 = '0' and a.VAF10 = '3' and a.VAF42 > '%s'%s%s order by a.VAF36", pid, admissionTime.VAE11, condition_type, condition_catg)
	err := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	if err != nil {
		return make([]MedicalAdviceExecution, 0), err
	}

	condition_stat := ""
	if status != 0 && status != 1 {
		condition_stat = fmt.Sprintf(" and State = '%d'", status)
	}
	response := make([]MedicalAdviceExecution, 0)
	for _, v := range mAdvices {
		madStatus := make([]MedicalAdviceExecutionStatus, 0)
		sql := fmt.Sprintf("select Madid,State,Process from AdviceStatus where Madid = %d%s", v.VAF01, condition_stat)
		err_db := fit.MySqlEngine().SQL(sql).Find(&madStatus)
		if err_db != nil {
			fit.Logger().LogError("**JK**",err_db.Error())
		}

		if leng := len(madStatus); leng == 0 {
			status := MedicalAdviceExecutionStatus{}
			status.Process = "未执行"
			status.State = 1
			madStatus = append(madStatus, status)
		}
		record := madStatus[0]
		if status == 0 {
			execution := MedicalAdviceExecution{}
			execution.MedicalAdvice = v
			execution.MedicalAdviceExecutionStatus = record
			response = append(response, execution)
		} else if status == record.State {
			execution := MedicalAdviceExecution{}
			execution.MedicalAdvice = v
			execution.MedicalAdviceExecutionStatus = record
			response = append(response, execution)
		}
	}

	return response, err
}

/*医嘱执行详情*/
func FetchMedicalAdviceExecutionDetail(madid int64, master int) ([]MedicalAdviceExecutionDetail, error) {
	mAdvice := make([]MedicalAdvice, 0)
	err := fit.SQLServerEngine().SQL("select a.VAF01,a.VAA01,a.VAF10,a.VAF11,a.BDA01,a.VAF19,a.VAF22,a.BBX01,b.VAF22 as Method,a.VAF23,a.VAF26,a.VAF36,a.BCE03A,c.BBX20,a.VAF60 from (VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01 where a.VAF01 = ?", madid).Find(&mAdvice)
	if err != nil {
		fit.Logger().LogError("***JP***", err.Error())
		return make([]MedicalAdviceExecutionDetail, 0), err
	}

	if length := len(mAdvice); length == 0 {
		return make([]MedicalAdviceExecutionDetail, 0), fmt.Errorf("无法查询到对应医嘱,二维码无效")
	}
	detail := MedicalAdviceExecutionDetail{}
	detail.MedicalAdvice = mAdvice[0]
	detail.Records = make([]MedicalAdviceExecutionRecord, 0)
	detail.Master = make([]VAA1, 0)
	detail.State = 0
	detail.Process = "未知"
	if master == 1 {
		fit.MySqlEngine().Table("VAA1").Where("VAA01 = ?", detail.MedicalAdvice.VAA01).Find(&detail.Master)
	}

	response := make([]MedicalAdviceExecutionDetail, 1)
	if detail.MedicalAdvice.VAF10 < 3 {
		detail.Desc = "该医嘱暂未生效"
		response[0] = detail
		return response, err
	} else if detail.MedicalAdvice.VAF10 == 4 {
		detail.Desc = "该医嘱已被作废"
		response[0] = detail
		return response, err
	} else if detail.MedicalAdvice.VAF10 > 4 && detail.MedicalAdvice.VAF10 < 7 {
		detail.Desc = "该医嘱已被删除或暂停"
		response[0] = detail
		return response, err
	} else if detail.MedicalAdvice.VAF10 > 8 && detail.MedicalAdvice.VAF10 < 10 {
		detail.Desc = "该医嘱已停止"
		response[0] = detail
		return response, err
	} else if detail.MedicalAdvice.VAF11 == 2 {
		timeMark := time.Time(detail.MedicalAdvice.VAF36)
		timeNow := time.Now()
		if timeNow.Sub(timeMark).Hours() <= 24 {
			if timeNow.Day()-timeMark.Day() > 0 {
				detail.Desc = "该医嘱已失效"
				response[0] = detail
				return response, err
			}
		} else {
			detail.Desc = "该医嘱已失效"
			response[0] = detail
			return response, err
		}
	}

	// 口服单只支持扫码执行，其它只支持手动结束
	//else if detail.BBX20 != 2 {
	//	detail.Desc = "该医嘱不支持被扫码执行"
	//	response[0] = detail
	//	return response, err
	//}

	err = fit.MySqlEngine().SQL("select * from AdviceDetail where Madid = ?", madid).Find(&detail.Records)
	if leng := len(detail.Records); leng == 0 {
		detail.Desc = "该医嘱暂未执行"
		detail.Process = "未执行"
		detail.State = 1
	} else {
		fit.MySqlEngine().SQL("select Madid,State,Process from AdviceStatus where Madid = ?", detail.MedicalAdvice.VAF01).Get(&detail.MedicalAdviceExecutionStatus)
	}
	response[0] = detail
	return response, err
}

/*JP 是否存在新医嘱*/
func IsExistNewMedicalAdvice(pid int64, hospitalDate string) int {
	// 获取校验过的医嘱
	mAdvices := make([]MedicalAdviceDup, 0)
	today := DatetimeNow()
	err_db := fit.SQLServerEngine().SQL("SELECT a.VAF01 from VAF2 a LEFT OUTER JOIN VAF2 b ON a.VAF42 > ? and b.VAF42 > ? and a.VAA01 = ? and b.VAA01 = ? and a.VAF01 = b.VAF01A where b.VAF01 is null and a.VAA01 = ? and a.VAF10 = '3' and (a.VAF11 = '1' or (a.VAF11 = '2' and a.VAF36 >= ?)) order by a.VAF36 desc", hospitalDate, hospitalDate, pid, pid, pid, today).Find(&mAdvices)

	for _, v := range mAdvices {
		// 在MySql数据库中查询当天的执行记录，有执行记录代表已被执行，即非新医嘱。改医嘱没被执行，即代表是新医嘱
		// 无论是长嘱还是临时医嘱，至少每天执行一次
		isEx := IsExist{}
		fit.MySqlEngine().SQL("select count(1) as Exist from AdviceState where time >= ? and Madid = ? order by id desc", today, v.VAF01).Get(&isEx)
		if isEx.Exist == 0 {
			return 1
		}
	}

	length := len(mAdvices)
	if length == 0 {
		if err_db != nil {
			fit.Logger().LogError("**JP**", err_db.Error())
		}
		return 0
	}
	return 0
}

/*JP 是否存在已停医嘱*/
func IsExistFinishedMedicalAdvice(pid int64, hospitalDate string) int {
	mAdvices := make([]MedicalAdviceDup, 0)
	//_, err_db := fit.SQLServerEngine().SQL("SELECT count(1) as Exist FROM VAF2 WHERE VAA01 = ? and VAF10 = 8 and BCE03D != '' and (VAF11 = 1 or (VAF11 = 2 and VAF36 >= ?))",pid, starttime).Get(&isEx)

	// 获取"已发生或已停止"的医嘱
	err_db := fit.SQLServerEngine().SQL("SELECT a.VAF01 from VAF2 a LEFT OUTER JOIN VAF2 b ON a.VAF42 > ? and b.VAF42 > ? and a.VAA01 = ? and b.VAA01 = ? and a.VAF01 = b.VAF01A where b.VAF01 is null and a.VAA01 = ? and a.VAF04 = '2' and a.VAF10 = '8' and (a.VAF11 = '1' or (a.VAF11 = '2' and a.VAF36 >= ?)) order by a.VAF36 desc", hospitalDate, hospitalDate, pid, pid, pid, DatetimeNow()).Find(&mAdvices)
	if length := len(mAdvices); length == 0 {
		if err_db != nil {
			fit.Logger().LogError("**JP**", err_db)
		}
		return 0
	}
	return 1
}

