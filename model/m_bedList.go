package model

import (
	"fit"
)

type Beds struct {
	VAA1 `xorm:"extends"`
	CBP1 `xorm:"extends"`
	VAF2 `xorm:"extends"`
	CBM1 `xorm:"extends"`
	//VAL []VAL1 `json:"allergic_drug"`
	VAL []VAL1 `json:"allergic_drug"` // 过敏药物
}

/*病人资料表*/
type VAA1 struct {
	VAA01  int          `json:"patient_id"`    // 病人ID
	VAA04  string       `json:"hosp_num"`      // 住院号
	VAA05  string       `json:"name"`          // 姓名
	ABW01  int          `json:"sex"`           // 性别
	VAA10  int          `json:"age"`           // 年龄
	BDP02  string       `json:"type"`          // 病人类型 BDP1表
	VAA61  int          `json:"status"`        // 就诊状态 2：住院
	BCK01B int          `json:"department_id"` // 科室ID  BCK1表
	VAA73  Datetime_IOV `json:"hospital_date"` // 入院时间
	BCQ04  string       `json:"bed_coding"`    // 床号 BCQ1表
}

/*病人护理记录*/
type CBP1 struct {
	//VAA01  病人ID
	AAG01 int `json:"nursing_degree"` // 护理级别  AAG1表
}

/*住院病人医嘱记录*/
type VAF2 struct {
	//VAA01  病人ID
	//BCK01A 病人科室ID
	VAF10 int `json:"order_status"` // 医嘱状态1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果
}

/*住院病人医嘱单*/
type CBM1 struct {
	//VAA01 病人ID
	//BCK01B 科室ID
	CBM07 int `json:"order_desc" xorm:"'CBM07'"` // 1=普通, 2=急诊, 3=儿童, 4=麻醉, 5=第一类精神药品, 6=第二类精神药品, 7=放射药品,8=毒性药品,9=检查,10=检验,11=手术,12=治疗,99=其它
}

/*过敏药物*/
type VAL1 struct {
	//VAA01 int // 病人ID
	BBX01 int    `json:"drug_id"`   //药物ID
	BBX05 string `json:"drug_name"` //过敏药物
}

func QueryDepartmentBedList(BCK01 int) ([]Beds, error) {
	responseObj := make([]Beds, 0)
	// 查找护理级别名称， and CBP1.AAG01 = AAG1.AAG01
	err := fit.SQLServerEngine().SQL("select VAA1.VAA01, VAA1.VAA04, VAA1.VAA05, VAA1.ABW01, VAA1.VAA10, VAA1.BDP02, VAA1.VAA61, VAA1.BCK01B, VAA1.BCQ04, VAA1.VAA73, CBP1.AAG01, CBP1.BCK01, VAF2.VAA01, VAF2.VAF10, VAF2.BCK01A, CBM1.VAA01, CBM1.CBM07, CBM1.BCK01B from VAA1, CBP1, VAF2, CBM1 where VAA1.VAA61 = '2' and VAA1.BCK01B = ? and VAA1.VAA01 = CBP1.VAA01 and VAA1.BCK01B = CBP1.BCK01 and VAA1.VAA01 = VAF2.VAA01 and VAA1.BCK01B = VAF2.BCK01A and VAA1.VAA01 = CBM1.VAA01 and VAA1.BCK01B = CBM1.BCK01B", BCK01).Find(&responseObj)
	for i, bed := range responseObj {
		VAL := make([]VAL1, 0)
		err = fit.SQLServerEngine().SQL("select VAL1.VAA01, VAL1.BBX01, VAL1.BBX05 from VAL1 where VAL1.VAA01 = ?", bed.VAA01).Find(&VAL)
		if length := len(VAL); length > 0 {
			for _, v := range VAL {
				responseObj[i].VAL = append(responseObj[i].VAL, v)
			}
		} else {
			responseObj[i].VAL = VAL
		}
	}
	return responseObj, err
}
