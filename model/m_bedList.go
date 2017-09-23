package model

import "fit"

type VAA1 struct {
	VAA01  int    `json:"patient_id"`                // 病人ID
	VAA04  string `json:"hosp_num"`                  // 住院号
	VAA05  string `json:"name"`                      // 姓名
	ABW01  int    `json:"sex"`                       // 性别
	VAA10  int    `json:"age"`                       // 年龄
	BDP02  string `json:"type"`                      // 病人类型 BDP1表
	VAA61  int    `json:"status"`                    // 就诊状态 2：住院
	BCK01B int    `json:"department_id"`             // 科室ID  BCK1表
	BCQ04  string `json:"bed_coding"`                // 床号 BCQ1表
	AAG01  int    `json:"nursing_degree"`            // 护理级别  AAG1表
	VAF10  int    `json:"order_status"`              // 医嘱状态1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果
	CBM07  int    `json:"order_desc" xorm:"'CBM07'"` // 1=普通, 2=急诊, 3=儿童, 4=麻醉, 5=第一类精神药品, 6=第二类精神药品, 7=放射药品,8=毒性药品,9=检查,10=检验,11=手术,12=治疗,99=其它
}

func QueryDepartmentBedList(BCK01 int, page int) ([]VAA1, error) {
	count := 10
	responseObj := make([]VAA1, 0)
	idx := page * count
	// 查找护理级别名称， and CBP1.AAG01 = AAG1.AAG01
	err := fit.SQLServerEngine().SQL("select * from VAA1, CBP1, VAF2, CBM2 where VAA1.BCK01B = ? and VAA1.VAA01 = CBP1.VAA01 and VAA1.BCK01B = CBP1.BCK01 and VAA1.VAA01 = VAF2.VAA01 and VAA1.BCK01B = VAF2.BCK01A and VAA1.VAA01 = CBM2.CBM02", BCK01).Limit(count, idx).Find(&responseObj)
	return responseObj, err
}
