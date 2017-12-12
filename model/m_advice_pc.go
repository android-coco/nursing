package model

import (
	"fit"
	"fmt"
)

/*医嘱拆分model*/
type MedicalAdviceModal struct {
	Vid       int64                  `json:"Vid"`       // 就诊ID
	Madid     int64                  `json:"Madid"`     // 医嘱ID
	Pid       int64                  `json:"Pid"`       // 病人ID
	Bed       string                 `json:"Bed"`       // 病人床位
	PName     string                 `json:"PName"`     // 病人姓名
	ExTime    DatetimeWithoutSeconds `json:"-"`         // 计划执行日期
	ExDay     string                 `json:"ExTime"`    // 计划执行日期[不存储]
	GroupNum  int                    `json:"GroupNum"`  // 组号
	Content   string                 `json:"Content"`   // 医嘱内容
	Dosage    string                 `json:"Dosage"`    // 用量、剂量
	Amount    string                 `json:"Amount"`    // 数量
	Frequency string                 `json:"Frequency"` // 频次
	Times     int                    `json:"Times"`     // 频次数值(医嘱执行次数)
	Method    string                 `json:"Method"`    // 医嘱用法
	Speed     string                 `json:"Speed"`     // 滴速
	TypeV     int                    `json:"TypeV"`     // 医嘱类型(1=长嘱,2=临嘱)
	TypeOf    string                 `json:"TypeOf"`    // 医嘱类型[不存储]
	StTime    DatetimeWithoutSeconds `json:"-"`         // 开始执行时间
	StDay     string                 `json:"StTime"`    // 开始执行时间[不存储]
	MStatus   string                 `json:"MStatus"`   // 医嘱状态[不存储]
	MStatusV  int                    `json:"MStatusV"`  // 医嘱状态数值
	Category  string                 `json:"Category"`  // 医嘱类别
	CategoryV string                 `json:"CategoryV"` // 医嘱类别原始值,BDA01
	PtType    string                 `json:"PtType"`    // 打印单类型,执行分类
	PtNum     int64                  `json:"PtNum"`     // 医嘱单的单号
	PtRownr   int                    `json:"PtRownr"`   // 医嘱单中医嘱的序号
	Entrust   string                 `json:"Entrust"`   // 医嘱嘱托
	Physician string                 `json:"Physician"` // 开嘱医师VAF2.BCE03A
	EdTime    DatetimeWithoutSeconds `json:"-"`         // 停嘱时间
	EdDay     string                 `json:"EdTime"`    // 停止时间[不存储]
	Sender    string                 `json:"Sender"`    // 发送人VBI2.BCE03A
	ExCycle   int                    `json:"ExCycle"`   // 执行周期(序号)
	ExNurse   string                 `json:"ExNurse"`   // 执行护士
	ExStatus  string                 `json:"ExStatus"`  // 执行状态,0=未执行,1=正在执行,2=已结束[不存储]
	ExStatusV int                    `json:"ExStatusV"` // 执行状态值
	ExStep    string                 `json:"ExStep"`    // 执行步骤,当前执行步骤
	PtTimes   int                    `json:"PtTimes"`   // 打印次数
	PtStatus  string                 `json:"PtStatus"`  // 打印状态[不存储,PtTimes=0=未打,PtTimes=1=已打]
}

/*医嘱内容*/
type MedicalAdviceContent struct {
	Madid   int64  `json:"Madid"`   // 医嘱ID
	Content string `json:"Content"` // 医嘱内容
	Dosage  string `json:"Dosage"`  // 用量、剂量
}

/*接口返回的医嘱数据*/
type MedicalAdviceResponse struct {
	Gid       int64                  `json:"Gid"`       // 医嘱组ID，首条医嘱的Madid
	Vid       int64                  `json:"Vid"`       // 就诊ID
	Pid       int64                  `json:"Pid"`       // 病人ID
	Bed       string                 `json:"Bed"`       // 病人床位
	PName     string                 `json:"PName"`     // 病人姓名
	ExTime    DatetimeWithoutSeconds `json:"-"`         // 计划执行日期
	ExDay     string                 `json:"ExTime"`    // 计划执行日期[不存储]
	GroupNum  int                    `json:"GroupNum"`  // 组号
	Contents  []MedicalAdviceContent `json:"Contents"`  // 医嘱内容
	Amount    string                 `json:"Amount"`    // 数量
	Frequency string                 `json:"Frequency"` // 频次
	Times     int                    `json:"Times"`     // 频次数值(医嘱执行次数)
	Method    string                 `json:"Method"`    // 医嘱用法
	Speed     string                 `json:"Speed"`     // 滴速
	TypeV     int                    `json:"TypeV"`     // 医嘱类型(1=长嘱,2=临嘱)
	TypeOf    string                 `json:"TypeOf"`    // 医嘱类型[不存储]
	StTime    DatetimeWithoutSeconds `json:"-"`         // 开始执行时间
	StDay     string                 `json:"StTime"`    // 开始执行时间[不存储]
	MStatus   string                 `json:"MStatus"`   // 医嘱状态[不存储]
	MStatusV  int                    `json:"MStatusV"`  // 医嘱状态数值
	Category  string                 `json:"Category"`  // 医嘱类别
	CategoryV string                 `json:"CategoryV"` // 医嘱类别原始值,BDA01
	PtType    string                 `json:"PtType"`    // 打印单类型,执行分类
	PtNum     int64                  `json:"PtNum"`     // 医嘱单的单号
	PtRownr   int                    `json:"PtRownr"`   // 医嘱单中医嘱的序号
	Entrust   string                 `json:"Entrust"`   // 医嘱嘱托
	Physician string                 `json:"Physician"` // 开嘱医师VAF2.BCE03A
	EdTime    DatetimeWithoutSeconds `json:"-"`         // 停嘱时间
	EdDay     string                 `json:"EdTime"`    // 停止时间[不存储]
	Sender    string                 `json:"Sender"`    // 发送人VBI2.BCE03A
	ExCycle   int                    `json:"ExCycle"`   // 执行周期(序号)
	ExNurse   string                 `json:"ExNurse"`   // 执行护士
	ExStatus  string                 `json:"ExStatus"`  // 执行状态,0=未执行,1=正在执行,2=已结束[不存储]
	ExStatusV int                    `json:"ExStatusV"` // 执行状态值
	ExStep    string                 `json:"ExStep"`    // 执行步骤,当前执行步骤
	PtTimes   int                    `json:"PtTimes"`   // 打印次数
	PtStatus  string                 `json:"PtStatus"`  // 打印状态[不存储,PtTimes=0=未打,PtTimes=1=已打]
}

/*MySql中的医嘱表*/
type MedicalAdviceItem struct {
	Vid       int64  `json:"Vid"`       // 就诊ID
	Madid     int64  `json:"Madid"`     // 医嘱ID
	Pid       int64  `json:"Pid"`       // 病人ID
	Bed       string `json:"Bed"`       // 病人床位
	PName     string `json:"PName"`     // 病人姓名
	ExTime    string `json:"ExTime"`    // 计划执行日期
	GroupNum  int    `json:"GroupNum"`  // 组号
	Content   string `json:"Content"`   // 医嘱内容
	Dosage    string `json:"Dosage"`    // 用量、剂量
	Amount    string `json:"Amount"`    // 数量
	Frequency string `json:"Frequency"` // 频次
	Times     int    `json:"Times"`     // 频次数值(医嘱执行次数)
	Method    string `json:"Method"`    // 医嘱用法
	Speed     string `json:"Speed"`     // 滴速
	TypeV     int    `json:"TypeV"`     // 医嘱类型(1=长嘱,2=临嘱)
	StTime    string `json:"StTime"`    // 开始执行时间
	MStatusV  int    `json:"MStatusV"`  // 医嘱状态数值
	Category  string `json:"Category"`  // 医嘱类别
	CategoryV string `json:"CategoryV"` // 医嘱类别原始值,BDA01
	PtType    string `json:"PtType"`    // 打印单类型,执行分类
	PtNum     int64  `json:"PtNum"`     // 医嘱单的单号
	PtRownr   int    `json:"PtRownr"`   // 医嘱单中医嘱的序号
	Entrust   string `json:"Entrust"`   // 医嘱嘱托
	Physician string `json:"Physician"` // 开嘱医师VAF2.BCE03A
	EdTime    string `json:"EdTime"`    // 停止时间
	Sender    string `json:"Sender"`    // 发送人VBI2.BCE03A
	ExCycle   int    `json:"ExCycle"`   // 执行周期(序号)
	ExNurse   string `json:"ExNurse"`   // 执行护士
	ExStatusV int    `json:"ExStatusV"` // 执行状态值
	ExStep    string `json:"ExStep"`    // 执行步骤,当前执行步骤
	PtTimes   int    `json:"PtTimes"`   // 打印次数
}

/*医嘱执行/打印字段*/
type MedicalAdvicePrintSubModel struct {
	Madid     int64  `json:"Madid"`     // 医嘱ID
	ExNurse   string `json:"ExNurse"`   // 执行护士
	ExStatus  string `json:"ExStatus"`  // 执行状态,0=未执行,1=正在执行,2=已结束[不存储]
	ExStatusV int    `json:"ExStatusV"` // 执行状态值
	ExStep    string `json:"ExStep"`    // 执行步骤,当前执行步骤
	PtTimes   int    `json:"PtTimes"`   // 打印次数
	PtStatus  string `json:"PtStatus"`  // 打印状态[不存储,PtTimes=0=未打,PtTimes=1=已打]
}

/*忽略打印的参数，照搬*/
func ignorePrintParameters(mtype, mprint int) (paramA, paramB int) {
	paramA = 0
	paramB = 0
	if mtype == 0 && mprint == 0 {
		paramA = 0
		paramB = 2
	} else if mtype == 0 && mprint == 1 {
		paramA = 1
		paramB = 3
	} else if mtype == 1 && mprint == 0 {
		paramA = 0
		paramB = 1
	} else if mtype == 1 && mprint == 1 {
		paramA = 2
		paramB = 3
	} else if mtype > 1 && mprint == 1 {
		paramA = 1
		paramB = 1
	}
	return
}

/*医嘱拆分-瓶签*/
func SearchSplitMedicalAdviceForBottlePost(startTime, endTime, vids string, typeOf, print int) ([]MedicalAdviceResponse, error) {
	paramA, paramB := ignorePrintParameters(0, 0)
	sqlHeader := "SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, a.VAF21 Amount, a.VAF26 Frequency, a.VAF27 Times, a2.BBX05 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, case when a.BDA01 = '3' then a1.VAF23 else a.VAF23 end Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender "
	sqlTable := "FROM VAF2 a JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 "
	sqlStr := fmt.Sprintf("%s%sWHERE a.VAF04 = 2 AND b.VBI07 > 0 AND((%d = 0) OR(%d > 0 AND a.VAF11 = %d)) AND c.VAE01 IN(%s) AND a2.BBX20 IN(2, 4, 5) AND( isnull(b.VBI29, 0) = %d OR isnull(b.VBI29, 0) = %d) AND(a.VAF10 >= 8 OR a.VAF10 = 3) AND(a.BDA01 >= '1' AND a.BDA01 <= '3') AND a2.BDA01 = 'T' AND a2.BBX13 IN('2', '4') AND datediff(MINUTE ,'%s', b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,'%s') >= 0) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", sqlHeader, sqlTable, typeOf, typeOf, typeOf, vids, paramA, paramB, startTime, endTime)
	fmt.Println("***JK***瓶签", sqlStr)
	mAdvices := make([]MedicalAdviceModal, 0)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	if err_ss != nil {
		fit.Logger().LogError("***JK***", "医嘱拆分-瓶签", err_ss.Error())
	}

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 || err_ss != nil {
		if err_ss != nil {
			fit.Logger().LogError("***JK***", err_ss.Error())
		}
		return arrA, err_ss
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, temp.Content, temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, v.Content, v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, v.Content, v.Dosage}
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
	response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		contentObj := resp.Contents[0]
		madid := resp.Gid
		content := contentObj.Content
		dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
			response = append(response, *resp)
		}
		if exec.Madid != madid {
			//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
			obj := MedicalAdviceItem{
				Vid:       resp.Vid,
				Madid:     madid,
				Pid:       resp.Pid,
				Bed:       resp.Bed,
				PName:     resp.PName,
				ExTime:    resp.ExTime.ParseToSecond(),
				GroupNum:  resp.GroupNum,
				Content:   content,
				Dosage:    dosage,
				Amount:    resp.Amount,
				Frequency: resp.Frequency,
				Times:     resp.Times,
				Method:    resp.Method,
				Speed:     resp.Speed,
				TypeV:     resp.TypeV,
				StTime:    resp.StTime.ParseToSecond(),
				MStatusV:  resp.MStatusV,
				Category:  resp.Category,
				CategoryV: resp.CategoryV,
				PtType:    resp.PtType,
				PtNum:     resp.PtNum,
				PtRownr:   resp.PtRownr,
				Entrust:   resp.Entrust,
				Physician: resp.Physician,
				EdTime:    resp.EdTime.ParseToSecond(),
				Sender:    resp.Sender,
				ExCycle:   resp.ExCycle,
				ExNurse:   resp.ExNurse,
				ExStatusV: resp.ExStatusV,
				ExStep:    resp.ExStep,
				PtTimes:   resp.PtTimes,
			}
			_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
			if err_in != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
				err_ss = err_in
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		} else {
			//	有记录
			sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
			_, err_up := fit.MySqlEngine().Exec(sqlStr)
			if err_up != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
				err_ss = err_up
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		}
	}
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱拆分-输液卡*/
func SearchSplitMedicalAdviceForInfusion(startTime, endTime, vids string, typeOf, print int) ([]MedicalAdviceResponse, error) {
	paramA, paramB := ignorePrintParameters(0, 0)
	sqlHeader := "SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '-' Dosage, a.VAF21 Amount, a.VAF26 Frequency, a.VAF27 Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a.VAF53 = 0 THEN '口服单' WHEN a.VAF53 = 1 THEN '注射单' when (a.VAF53 = 2) or (a.VAF53 = 4) then '输液单' When a.VAF53 = 3 then '治疗单' When a.VAF53 = 5 Then '输血单' When a.VAF53 = 6 Then '护理单' end as PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender "
	sqlTable := "From VAF2 a JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 "
	sqlStr := fmt.Sprintf("%s%s where a.BDA01 = '0' and ((%d = 0)or(%d > 0 and a.VAF11 = %d)) and a.VAF04 = 2 and c.VAE01 in (%s) and a.VAF53 in (2,4,5) And (isnull(b.VBI29,0) = %d or isnull(b.VBI29,0) = %d) and (a.VAF10 >= 8 OR a.VAF10 = 3) and datediff(MINUTE, '%s', b.VBI10) >= 0 and datediff(MINUTE, b.VBI10, '%s') >= 0) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", sqlHeader, sqlTable, typeOf, typeOf, typeOf, vids, paramA, paramB, startTime, endTime)
	mAdvices := make([]MedicalAdviceModal, 0)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	fmt.Println("***JK***输液卡", sqlStr)

	if err_ss != nil {
		fit.Logger().LogError("***JK***", "医嘱拆分-输液单", err_ss.Error())
	}

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 || err_ss != nil {
		if err_ss != nil {
			fit.Logger().LogError("***JK***", err_ss.Error())
		}
		return arrA, err_ss
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, temp.Content, temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, v.Content, v.Dosage})
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
					object.Gid = object.Contents[0].Madid
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, v.Content, v.Dosage}
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
	response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		content := contentObj.Content
		dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
			response = append(response, *resp)
		}
		if exec.Madid != madid {
			//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
			obj := MedicalAdviceItem{
				Vid:       resp.Vid,
				Madid:     madid,
				Pid:       resp.Pid,
				Bed:       resp.Bed,
				PName:     resp.PName,
				ExTime:    resp.ExTime.ParseToSecond(),
				GroupNum:  resp.GroupNum,
				Content:   content,
				Dosage:    dosage,
				Amount:    resp.Amount,
				Frequency: resp.Frequency,
				Times:     resp.Times,
				Method:    resp.Method,
				Speed:     resp.Speed,
				TypeV:     resp.TypeV,
				StTime:    resp.StTime.ParseToSecond(),
				MStatusV:  resp.MStatusV,
				Category:  resp.Category,
				CategoryV: resp.CategoryV,
				PtType:    resp.PtType,
				PtNum:     resp.PtNum,
				PtRownr:   resp.PtRownr,
				Entrust:   resp.Entrust,
				Physician: resp.Physician,
				EdTime:    resp.EdTime.ParseToSecond(),
				Sender:    resp.Sender,
				ExCycle:   resp.ExCycle,
				ExNurse:   resp.ExNurse,
				ExStatusV: resp.ExStatusV,
				ExStep:    resp.ExStep,
				PtTimes:   resp.PtTimes,
			}
			_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
			if err_in != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
				err_ss = err_in
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		} else {
			//	有记录
			sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
			_, err_up := fit.MySqlEngine().Exec(sqlStr)
			if err_up != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
				err_ss = err_up
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		}
	}
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱拆分-口服单*/
func SearchSplitMedicalAdviceForOralMedical(startTime, endTime, vids string, typeOf, print int) ([]MedicalAdviceResponse, error) {
	paramA, paramB := ignorePrintParameters(2, 0)
	sqlHeader := "SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, a.VAF21 Amount, a.VAF26 Frequency, a.VAF27 Times, a2.BBX05 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, case when a.BDA01 = '3' then a1.VAF23 else a.VAF23 end as Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender "
	sqlTable := "FROM VAF2 a JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 "
	sqlStr := fmt.Sprintf("%s%sWHERE a.VAF04 = 2 and b.VBI07 > 0 and ((%d = 0)or(%d > 0 and a.VAF11 = %d)) AND c.VAE01 IN (%s) AND a2.BBX20 IN (0) and (isnull(b.VBI29,0) = %d or isnull(b.VBI29,0) = %d) and (a.VAF10 >= 8 OR a.VAF10 = 3) AND (a.BDA01 >= '1' And a.BDA01 <= '3') AND a2.BDA01 = 'T' AND a2.BBX13 in ('2','4') and datediff(MINUTE, '%s', b.VBI10) >= 0 and datediff(MINUTE, b.VBI10, '%s') >= 0) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", sqlHeader, sqlTable, typeOf, typeOf, typeOf, vids, paramA, paramB, startTime, endTime)
	fmt.Println("***JK***口服单", sqlStr)
	mAdvices := make([]MedicalAdviceModal, 0)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	if err_ss != nil {
		fit.Logger().LogError("***JK***", "医嘱拆分-口服单", err_ss.Error())
	}

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 || err_ss != nil {
		if err_ss != nil {
			fit.Logger().LogError("***JK***", err_ss.Error())
		}
		return arrA, err_ss
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, temp.Content, temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, v.Content, v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, v.Content, v.Dosage}
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
	response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		content := contentObj.Content
		dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
			response = append(response, *resp)
		}
		if exec.Madid != madid {
			//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
			obj := MedicalAdviceItem{
				Vid:       resp.Vid,
				Madid:     madid,
				Pid:       resp.Pid,
				Bed:       resp.Bed,
				PName:     resp.PName,
				ExTime:    resp.ExTime.ParseToSecond(),
				GroupNum:  resp.GroupNum,
				Content:   content,
				Dosage:    dosage,
				Amount:    resp.Amount,
				Frequency: resp.Frequency,
				Times:     resp.Times,
				Method:    resp.Method,
				Speed:     resp.Speed,
				TypeV:     resp.TypeV,
				StTime:    resp.StTime.ParseToSecond(),
				MStatusV:  resp.MStatusV,
				Category:  resp.Category,
				CategoryV: resp.CategoryV,
				PtType:    resp.PtType,
				PtNum:     resp.PtNum,
				PtRownr:   resp.PtRownr,
				Entrust:   resp.Entrust,
				Physician: resp.Physician,
				EdTime:    resp.EdTime.ParseToSecond(),
				Sender:    resp.Sender,
				ExCycle:   resp.ExCycle,
				ExNurse:   resp.ExNurse,
				ExStatusV: resp.ExStatusV,
				ExStep:    resp.ExStep,
				PtTimes:   resp.PtTimes,
			}
			_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
			if err_in != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
				err_ss = err_in
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		} else {
			//	有记录
			sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
			_, err_up := fit.MySqlEngine().Exec(sqlStr)
			if err_up != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
				err_ss = err_up
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		}
	}
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱拆分-注射单*/
func SearchSplitMedicalAdviceForOralInjection(startTime, endTime, vids string, typeOf, print int) ([]MedicalAdviceResponse, error) {
	paramA, paramB := ignorePrintParameters(2, 1)
	sqlHeader := "SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '-' Dosage, a.VAF21 Amount, a.VAF26 Frequency, a.VAF27 Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a.VAF53 = 0 THEN '口服单' WHEN a.VAF53 = 1 THEN '注射单' when (a.VAF53 = 2) or (a.VAF53 = 4) then '输液单' When a.VAF53 = 3 then '治疗单' When a.VAF53 = 5 Then '输血单' When a.VAF53 = 6 Then '护理单' end as PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender "
	sqlTable := "From VAF2 a JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 "
	sqlStr := fmt.Sprintf("%s%swhere a.BDA01 = '0' and ((%d = 0)or(%d > 0 and a.VAF11 = %d)) and a.VAF04 = 2 And c.VAE01 in (%s) and a.VAF53 in (1) And (isnull(b.VBI29,0) = %d or isnull(b.VBI29,0) = %d) and (a.VAF10 >= 8 OR a.VAF10 = 3) and datediff(MINUTE, '%s', b.VBI10) >= 0 and datediff(MINUTE, b.VBI10, '%s') >= 0) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", sqlHeader, sqlTable, typeOf, typeOf, typeOf, vids, paramA, paramB, startTime, endTime)
	fmt.Println("***JK***注射单", sqlStr)
	mAdvices := make([]MedicalAdviceModal, 0)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	if err_ss != nil {
		fit.Logger().LogError("***JK***", "医嘱拆分-注射单", err_ss.Error())
	}

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 || err_ss != nil {
		if err_ss != nil {
			fit.Logger().LogError("***JK***", err_ss.Error())
		}
		return arrA, err_ss
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, temp.Content, temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, v.Content, v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, v.Content, v.Dosage}
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
				object.Gid = object.Contents[0].Madid
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
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		content := contentObj.Content
		dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
			response = append(response, *resp)
		}
		if exec.Madid != madid {
			//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
			obj := MedicalAdviceItem{
				Vid:       resp.Vid,
				Madid:     madid,
				Pid:       resp.Pid,
				Bed:       resp.Bed,
				PName:     resp.PName,
				ExTime:    resp.ExTime.ParseToSecond(),
				GroupNum:  resp.GroupNum,
				Content:   content,
				Dosage:    dosage,
				Amount:    resp.Amount,
				Frequency: resp.Frequency,
				Times:     resp.Times,
				Method:    resp.Method,
				Speed:     resp.Speed,
				TypeV:     resp.TypeV,
				StTime:    resp.StTime.ParseToSecond(),
				MStatusV:  resp.MStatusV,
				Category:  resp.Category,
				CategoryV: resp.CategoryV,
				PtType:    resp.PtType,
				PtNum:     resp.PtNum,
				PtRownr:   resp.PtRownr,
				Entrust:   resp.Entrust,
				Physician: resp.Physician,
				EdTime:    resp.EdTime.ParseToSecond(),
				Sender:    resp.Sender,
				ExCycle:   resp.ExCycle,
				ExNurse:   resp.ExNurse,
				ExStatusV: resp.ExStatusV,
				ExStep:    resp.ExStep,
				PtTimes:   resp.PtTimes,
			}
			_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
			if err_in != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
				err_ss = err_in
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		} else {
			//	有记录
			sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
			_, err_up := fit.MySqlEngine().Exec(sqlStr)
			if err_up != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
				err_ss = err_up
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		}
	}
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱执行查询*/
func SearchMedicalAdviceExecutionForPC(typeOf, status int, category, vids, st, et string) ([]MedicalAdviceResponse, error) {
	condition_type := ""
	if typeOf != 0 {
		// 长期和临时
		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
	}

	condition_catg := " and a.BDA01 != '0'"
	if category != "0" {
		// 医嘱类别
		condition_catg = fmt.Sprintf(" and a.BDA01 = '%s'", category)
	}

	condition_time := " AND b.VBI10 > c.VAE11"
	if st != "all" && et != "all" {
		condition_time = fmt.Sprintf(" AND b.VBI10 > c.VAE11 AND b.VBI10 BETWEEN '%s' AND '%s'", st, et)
	}

	mAdvices := make([]MedicalAdviceModal, 0)
	//sqlStr := fmt.Sprintf("select a.VAF01 as Madid,a.VAA01 as Pid,v.VAA05 as PatientName,v.BCQ04 as Bed,d.BDA02 as Category,a.VAF22 as Content,a.VAF19 as Dosage,a.VAF21 as Count,b.VAF22 as Method,a.VAF60 Speed,a.VAF11 as TypeOf,a.VAF23 as Entrust,a.BCE03A as Physician,a.VAF36 as StTime,a.VAF47 as EdTime,a.VAF10 as Status,a.BCE03F as Nurse,a.VAF50 as CkTime from (((VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01) left join BDA1 d on a.BDA01 = d.BDA01) left join VAA1 v on v.VAA01 = a.VAA01 where a.VAA01 in (%s) and a.VAF32 = 0%s%s%s%s order by a.VAF36", pids, condition_type, condition_catg, condition_state, condition_time)
	sqlStr := fmt.Sprintf("SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, a.VAF21 Amount, a.VAF26 Frequency, a.VAF27 Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' WHEN e.BDA01 = 'T' THEN '治疗单' WHEN e.BDA01 = 'N' THEN '护理单' ELSE '其它' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, CASE WHEN a.VAF10 >= 8 THEN a.VAF47 ELSE '' END AS EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF04 = 2 AND a.VAF32 = 0 %s AND b.VBI07 > 0 AND c.VAE01 IN(%s) %s and a.VAF10 = 3 %s) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", condition_catg, vids, condition_type, condition_time)
	fit.Logger().LogDebug("***JK***医嘱查询", sqlStr)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 || err_ss != nil {
		if err_ss != nil {
			fit.Logger().LogError("***JK***", err_ss.Error())
		}
		return arrA, err_ss
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, temp.Content, temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, v.Content, v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, v.Content, v.Dosage}
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
	response := make([]MedicalAdviceResponse, 0)
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		content := contentObj.Content
		dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		// 执行状态归类
		if status == 1 && exec.ExStatusV == 1 {
			response = append(response, *resp)
		} else if status == 2 && exec.ExStatusV == 2 {
			response = append(response, *resp)
		} else if status == 3 && exec.ExStatusV == 3 {
			response = append(response, *resp)
		}

		if exec.Madid != madid {
			//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
			obj := MedicalAdviceItem{
				Vid:       resp.Vid,
				Madid:     madid,
				Pid:       resp.Pid,
				Bed:       resp.Bed,
				PName:     resp.PName,
				ExTime:    resp.ExTime.ParseToSecond(),
				GroupNum:  resp.GroupNum,
				Content:   content,
				Dosage:    dosage,
				Amount:    resp.Amount,
				Frequency: resp.Frequency,
				Times:     resp.Times,
				Method:    resp.Method,
				Speed:     resp.Speed,
				TypeV:     resp.TypeV,
				StTime:    resp.StTime.ParseToSecond(),
				MStatusV:  resp.MStatusV,
				Category:  resp.Category,
				CategoryV: resp.CategoryV,
				PtType:    resp.PtType,
				PtNum:     resp.PtNum,
				PtRownr:   resp.PtRownr,
				Entrust:   resp.Entrust,
				Physician: resp.Physician,
				EdTime:    resp.EdTime.ParseToSecond(),
				Sender:    resp.Sender,
				ExCycle:   resp.ExCycle,
				ExNurse:   resp.ExNurse,
				ExStatusV: resp.ExStatusV,
				ExStep:    resp.ExStep,
				PtTimes:   resp.PtTimes,
			}
			_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
			if err_in != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
				err_ss = err_in
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		} else {
			//	有记录
			sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
			_, err_up := fit.MySqlEngine().Exec(sqlStr)
			if err_up != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
				err_ss = err_up
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		}
	}
	if status == 0 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱查询*/
func SearchMedicalAdviceForPC(typeOf, status int, category, vids, st, et string) ([]MedicalAdviceResponse, error) {
	condition_type := ""
	if typeOf != 0 {
		// 长期和临时
		condition_type = fmt.Sprintf(" and a.VAF11 = '%d'", typeOf)
	}

	condition_catg := " AND a.BDA01 != '0'"
	if category != "0" {
		// 医嘱类别
		condition_catg = fmt.Sprintf(" AND a.BDA01 = '%s'", category)
	}

	// 医嘱状态所有
	condition_state := " AND (a.VAF10 >= 8 or a.VAF10 in (3,4))"
	if status != 0 {
		// 医嘱状态 未停、已撤销、已停
		if status >= 8 {
			condition_state = " AND a.VAF10 >= 8"
		} else {
			condition_state = fmt.Sprintf(" AND a.VAF10 = '%d'", status)
		}
	}

	condition_time := " AND b.VBI10 > c.VAE11"
	if st != "all" && et != "all" {
		condition_time = fmt.Sprintf(" AND b.VBI10 > c.VAE11 AND b.VBI10 BETWEEN '%s' AND '%s'", st, et)
	}

	mAdvices := make([]MedicalAdviceModal, 0)
	//sqlStr := fmt.Sprintf("select a.VAF01 as Madid,a.VAA01 as Pid,v.VAA05 as PatientName,v.BCQ04 as Bed,d.BDA02 as Category,a.VAF22 as Content,a.VAF19 as Dosage,a.VAF21 as Count,b.VAF22 as Method,a.VAF60 Speed,a.VAF11 as TypeOf,a.VAF23 as Entrust,a.BCE03A as Physician,a.VAF36 as StTime,a.VAF47 as EdTime,a.VAF10 as Status,a.BCE03F as Nurse,a.VAF50 as CkTime from (((VAF2 a left join VAF2 b on a.VAF01A = b.VAF01) left join BBX1 c on c.BBX01 = a.BBX01) left join BDA1 d on a.BDA01 = d.BDA01) left join VAA1 v on v.VAA01 = a.VAA01 where a.VAA01 in (%s) and a.VAF32 = 0%s%s%s%s order by a.VAF36", pids, condition_type, condition_catg, condition_state, condition_time)
	sqlStr := fmt.Sprintf("SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, a.VAF21 Amount, a.VAF26 Frequency, a.VAF27 Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' WHEN e.BDA01 = 'T' THEN '治疗单' WHEN e.BDA01 = 'N' THEN '护理单' ELSE '其它' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, CASE WHEN a.VAF10 >= 8 THEN a.VAF47 ELSE '' END AS EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF04 = 2 AND a.VAF32 = 0 %s AND b.VBI07 > 0 AND c.VAE01 IN(%s) %s %s %s) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", condition_catg, vids, condition_type, condition_state, condition_time)
	fit.Logger().LogDebug("***JK***医嘱查询", sqlStr)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)

	arrA := make([]MedicalAdviceResponse, 0)
	length := len(mAdvices)
	if length == 0 || err_ss != nil {
		if err_ss != nil {
			fit.Logger().LogError("***JK***", err_ss.Error())
		}
		return arrA, err_ss
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, temp.Content, temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, v.Content, v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, v.Content, v.Dosage}
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
	for i := 0; i < length; i ++ {
		resp := &arrA[i]
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		content := contentObj.Content
		dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if exec.Madid != madid {
			//	 无记录,将第一条医嘱插入MySQL，代替整组医嘱
			obj := MedicalAdviceItem{
				Vid:       resp.Vid,
				Madid:     madid,
				Pid:       resp.Pid,
				Bed:       resp.Bed,
				PName:     resp.PName,
				ExTime:    resp.ExTime.ParseToSecond(),
				GroupNum:  resp.GroupNum,
				Content:   content,
				Dosage:    dosage,
				Amount:    resp.Amount,
				Frequency: resp.Frequency,
				Times:     resp.Times,
				Method:    resp.Method,
				Speed:     resp.Speed,
				TypeV:     resp.TypeV,
				StTime:    resp.StTime.ParseToSecond(),
				MStatusV:  resp.MStatusV,
				Category:  resp.Category,
				CategoryV: resp.CategoryV,
				PtType:    resp.PtType,
				PtNum:     resp.PtNum,
				PtRownr:   resp.PtRownr,
				Entrust:   resp.Entrust,
				Physician: resp.Physician,
				EdTime:    resp.EdTime.ParseToSecond(),
				Sender:    resp.Sender,
				ExCycle:   resp.ExCycle,
				ExNurse:   resp.ExNurse,
				ExStatusV: resp.ExStatusV,
				ExStep:    resp.ExStep,
				PtTimes:   resp.PtTimes,
			}
			_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
			if err_in != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-Insert", err_in.Error())
				err_ss = err_in
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		} else {
			//	有记录
			sqlStr = fmt.Sprintf("update medicaladvice set Bed='%s', GroupNum='%d', Times='%d', MStatusV='%d', PtNum='%d', EdTime='%s', Sender='%s', ExCycle='%d' where Madid = %d and ExTime = '%s' and ExCycle = %d", resp.Bed, resp.GroupNum, resp.Times, resp.MStatusV, resp.PtNum, resp.EdTime.ParseToSecond(), resp.Sender, resp.ExCycle, madid, resp.ExTime.ParseToSecond(), resp.ExCycle)
			_, err_up := fit.MySqlEngine().Exec(sqlStr)
			if err_up != nil {
				fit.Logger().LogError("***JK***", "医嘱查询-update", err_up.Error())
				err_ss = err_up
				return make([]MedicalAdviceResponse, 0), err_ss
			}
		}
	}
	return arrA, err_ss
}

/*获取医嘱执行记录*/
func FetchMedicalAdviceExecutionRecordForPc(madid int64, excycle, extime string) ([]MedicalAdviceExecutionRecord, error) {
	response := make([]MedicalAdviceExecutionRecord, 0)
	err := fit.MySqlEngine().SQL("SELECT DATE_FORMAT(Plan,'%Y-%m-%d %H:%i:%s') Plan, DATE_FORMAT(ExecTime,'%Y-%m-%d %H:%i') ExecTime,AdviceDetail.* from AdviceDetail where Madid = ? and Plan = ? and ExCycle = ?", madid, extime, excycle).Find(&response)
	return response, err
}