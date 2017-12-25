package model

import (
	"fit"
	"fmt"
	"strings"
	"time"
)

/*医嘱拆分model*/
type MedicalAdviceModal struct {
	Vid         int64                  `json:"Vid"`       // 就诊ID
	Madid       int64                  `json:"Madid"`     // 医嘱ID
	Pid         int64                  `json:"Pid"`       // 病人ID
	Bed         string                 `json:"Bed"`       // 病人床位
	PName       string                 `json:"PName"`     // 病人姓名
	Gender      string                 `json:"Gender"`    // 性别
	Age         string                 `json:"Age"`       // 年龄
	HospNum     string                 `json:"HospNum"`   // 住院号
	ExTime      DatetimeWithoutSeconds `json:"-"`         // 计划执行日期
	ExDay       string                 `json:"ExTime"`    // 计划执行日期[不存储]
	GroupNum    int                    `json:"GroupNum"`  // 组号
	Content     string                 `json:"Content"`   // 医嘱内容
	Dosage      string                 `json:"Dosage"`    // 用量、剂量
	Amount      string                 `json:"Amount"`    // 数量
	Frequency   string                 `json:"Frequency"` // 频次
	Times       int                    `json:"Times"`     // 频次数值(医嘱执行次数)
	Method      string                 `json:"Method"`    // 医嘱用法
	Speed       string                 `json:"Speed"`     // 滴速
	TypeV       int                    `json:"TypeV"`     // 医嘱类型(1=长嘱,2=临嘱)
	TypeOf      string                 `json:"TypeOf"`    // 医嘱类型[不存储]
	StTime      DatetimeWithoutSeconds `json:"-"`         // 开始执行时间
	StDay       string                 `json:"StTime"`    // 开始执行时间[不存储]
	MStatus     string                 `json:"MStatus"`   // 医嘱状态[不存储]
	MStatusV    int                    `json:"MStatusV"`  // 医嘱状态数值
	Category    string                 `json:"Category"`  // 医嘱类别
	CategoryV   string                 `json:"CategoryV"` // 医嘱类别原始值,BDA01
	PtType      string                 `json:"PtType"`    // 打印单类型,执行分类
	PtTypeV     int                    `json:"ptTypeV"`   // 打印单类型值,1=输液单,2=口服单,3=注射单,4=输液瓶签,5=检验标签
	PtNum       int64                  `json:"PtNum"`     // 医嘱单的单号
	PtRownr     int                    `json:"PtRownr"`   // 医嘱单中医嘱的序号
	Entrust     string                 `json:"Entrust"`   // 医嘱嘱托
	Physician   string                 `json:"Physician"` // 开嘱医师VAF2.BCE03A
	EdTime      DatetimeWithoutSeconds `json:"-"`         // 停嘱时间
	EdDay       string                 `json:"EdTime"`    // 停止时间[不存储]
	Sender      string                 `json:"Sender"`    // 发送人VBI2.BCE03A
	ExCycle     int                    `json:"ExCycle"`   // 执行周期(序号)
	ExNurse     string                 `json:"ExNurse"`   // 执行护士
	ExStatus    string                 `json:"ExStatus"`  // 执行状态,0=未执行,1=正在执行,2=已结束[不存储]
	ExStatusV   int                    `json:"ExStatusV"` // 执行状态值
	ExStep      string                 `json:"ExStep"`    // 执行步骤,当前执行步骤
	PtTimes     int                    `json:"PtTimes"`   // 打印次数
	PtStatus    string                 `json:"PtStatus"`  // 打印状态[不存储,PtTimes=0=未打,PtTimes=1=已打]
	HisExStatus int                    `json:"-"`         // His系统VBI2表里面的执行状态VBI13,0:未执行; 1:执行完成; 2:拒绝执行; 3:正在执行;9：作废
}

/*医嘱内容*/
type MedicalAdviceContent struct {
	Madid   int64  `json:"madid"`   // 医嘱ID
	Content string `json:"content"` // 医嘱内容
	Dosage  string `json:"dosage"`  // 用量、剂量
}

/*PC接口返回的医嘱数据*/
type MedicalAdviceResponse struct {
	Gid         int64                  `json:"gid"`       // 医嘱组ID，首条医嘱的Madid
	Vid         int64                  `json:"vid"`       // 就诊ID
	Pid         int64                  `json:"pid"`       // 病人ID
	Bed         string                 `json:"bed"`       // 病人床位
	PName       string                 `json:"pName"`     // 病人姓名
	Gender      string                 `json:"gender"`    // 性别
	Age         string                 `json:"age"`       // 年龄
	HospNum     string                 `json:"hospNum"`   // 住院号
	ExTime      DatetimeWithoutSeconds `json:"-"`         // 计划执行日期
	ExDay       string                 `json:"exTime"`    // 计划执行日期[不存储]
	GroupNum    int                    `json:"groupNum"`  // 组号
	Contents    []MedicalAdviceContent `json:"contents"`  // 医嘱内容
	Amount      string                 `json:"amount"`    // 数量
	Frequency   string                 `json:"frequency"` // 频次
	Times       int                    `json:"times"`     // 频次数值(医嘱执行次数)
	Method      string                 `json:"method"`    // 医嘱用法
	Speed       string                 `json:"speed"`     // 滴速
	TypeV       int                    `json:"typeV"`     // 医嘱类型(1=长嘱,2=临嘱)
	TypeOf      string                 `json:"typeOf"`    // 医嘱类型[不存储]
	StTime      DatetimeWithoutSeconds `json:"-"`         // 开始执行时间
	StDay       string                 `json:"stTime"`    // 开始执行时间[不存储]
	MStatus     string                 `json:"mStatus"`   // 医嘱状态[不存储]
	MStatusV    int                    `json:"mStatusV"`  // 医嘱状态数值
	Category    string                 `json:"category"`  // 医嘱类别
	CategoryV   string                 `json:"categoryV"` // 医嘱类别原始值,BDA01
	PtType      string                 `json:"ptType"`    // 打印单类型,执行分类
	PtTypeV     int                    `json:"ptTypeV"`   // 打印单类型值,1=输液单,2=口服单,3=注射单,4=输液瓶签,5=检验标签
	PtNum       int64                  `json:"ptNum"`     // 医嘱单的单号
	PtRownr     int                    `json:"ptRownr"`   // 医嘱单中医嘱的序号
	Entrust     string                 `json:"entrust"`   // 医嘱嘱托
	Physician   string                 `json:"physician"` // 开嘱医师VAF2.BCE03A
	EdTime      DatetimeWithoutSeconds `json:"-"`         // 停嘱时间
	EdDay       string                 `json:"edTime"`    // 停止时间[不存储]
	Sender      string                 `json:"sender"`    // 发送人VBI2.BCE03A
	ExCycle     int                    `json:"exCycle"`   // 执行周期(序号)
	ExNurse     string                 `json:"exNurse"`   // 执行护士
	ExStatus    string                 `json:"exStatus"`  // 执行状态,0=未执行,1=正在执行,2=已结束[不存储]
	ExStatusV   int                    `json:"exStatusV"` // 执行状态值
	ExStep      string                 `json:"exStep"`    // 执行步骤,当前执行步骤
	PtTimes     int                    `json:"ptTimes"`   // 打印次数
	PtStatus    string                 `json:"ptStatus"`  // 打印状态[不存储,PtTimes=0=未打,PtTimes=1=已打]
	HisExStatus int                    `json:"-"`         // His系统VBI2表里面的执行状态VBI13,0:未执行; 1:执行完成; 2:拒绝执行; 3:正在执行;9：作废
}

/*医嘱执行详情*/

type MedicalAdviceExecutionDetail struct {
	MedicalAdviceResponse                  `xorm:"extends"`
	Records []MedicalAdviceExecutionRecord `json:"records" xorm:"extends"`
	Desc    string                         `json:"desc"` // 关于医嘱的描述
}

/*MySql中的医嘱表字段*/
type MedicalAdviceItem struct {
	Vid       int64  `json:"vid"`       // 就诊ID
	Madid     int64  `json:"madid"`     // 医嘱ID
	Pid       int64  `json:"pid"`       // 病人ID
	Bed       string `json:"bed"`       // 病人床位
	PName     string `json:"pName"`     // 病人姓名
	Gender    string `json:"gender"`    // 性别
	Age       string `json:"age"`       // 年龄
	HospNum   string `json:"hospNum"`   // 住院号
	ExTime    string `json:"exTime"`    // 计划执行日期
	GroupNum  int    `json:"groupNum"`  // 组号
	Content   string `json:"content"`   // 医嘱内容
	Dosage    string `json:"dosage"`    // 用量、剂量
	Amount    string `json:"amount"`    // 数量
	Frequency string `json:"frequency"` // 频次
	Times     int    `json:"times"`     // 频次数值(医嘱执行次数)
	Method    string `json:"method"`    // 医嘱用法
	Speed     string `json:"speed"`     // 滴速
	TypeV     int    `json:"typeV"`     // 医嘱类型(1=长嘱,2=临嘱)
	StTime    string `json:"stTime"`    // 开始执行时间
	MStatusV  int    `json:"mStatusV"`  // 医嘱状态数值
	Category  string `json:"category"`  // 医嘱类别
	CategoryV string `json:"categoryV"` // 医嘱类别原始值,BDA01
	PtType    string `json:"ptType"`    // 打印单类型,执行分类
	PtNum     int64  `json:"ptNum"`     // 医嘱单的单号
	PtRownr   int    `json:"ptRownr"`   // 医嘱单中医嘱的序号
	Entrust   string `json:"entrust"`   // 医嘱嘱托
	Physician string `json:"physician"` // 开嘱医师VAF2.BCE03A
	EdTime    string `json:"edTime"`    // 停止时间
	Sender    string `json:"sender"`    // 发送人VBI2.BCE03A
	ExCycle   int    `json:"exCycle"`   // 执行周期(序号)
	ExNurse   string `json:"exNurse"`   // 执行护士
	ExStatusV int    `json:"exStatusV"` // 执行状态值
	ExStep    string `json:"exStep"`    // 执行步骤,当前执行步骤
	PtTimes   int    `json:"ptTimes"`   // 打印次数
}

/*医嘱执行/打印字段*/
type MedicalAdvicePrintSubModel struct {
	Madid     int64  `json:"madid"`     // 医嘱ID
	ExNurse   string `json:"exNurse"`   // 执行护士
	ExStatus  string `json:"exStatus"`  // 执行状态,0=未执行,1=正在执行,2=已结束[不存储]
	ExStatusV int    `json:"exStatusV"` // 执行状态值
	ExStep    string `json:"exStep"`    // 执行步骤,当前执行步骤
	PtTimes   int    `json:"ptTimes"`   // 打印次数
	PtStatus  string `json:"ptStatus"`  // 打印状态[不存储,PtTimes=0=未打,PtTimes=1=已打]
}

/*医嘱拆分-瓶签*/
func SearchSplitMedicalAdviceForBottlePost(startTime, endTime, vids string, typeOf, print int) ([]MedicalAdviceResponse, error) {
	sqlDeclare := fmt.Sprintf("Declare @lDt1 varchar(30) ,@lDt2 varchar(30) ,@asign tinyint Set @lDt1 = Convert( varchar(16), Cast('%s' as datetime), 121) + ':00' Set @lDt2 = Convert( varchar(16), Cast('%s' as datetime), 121) + ':59' set @asign = 0 if DATEDIFF(DAY ,@lDt1 ,@lDt2) = 0 set @asign = 1 ", startTime, endTime)
	sqlHeader := "SELECT d.* FROM( SELECT a.VAF06 Vid, 4 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a2.BBX05 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, '瓶签' PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender "
	sqlTable := "FROM VAF2 a JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 "
	sqlStr := fmt.Sprintf("%s%s%sWHERE a.VAF04 = 2 AND b.VBI07 > 0 AND((%d = 0) OR(%d > 0 AND a.VAF11 = %d)) AND c.VAE01 IN(%s) AND a2.BBX20 IN(2, 4, 5) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND(a.VAF10 = 3 OR a.VAF10 >= 8))) AND(a.BDA01 = '1' OR a.BDA01 = '2') AND a2.BDA01 = 'T' AND a2.BBX13 = '2' AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 AND( @asign = 0 OR(@asign = 1 AND a.VAF11 = 2) OR( a.VAF11 = 1 AND @asign = 1 AND DATEDIFF(DAY, a.VAF36, b.VBI10) <> 0) OR( a.VAF11 = 1 AND @asign = 1 AND DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 AND a.VAF61 > 0.000001))) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", sqlDeclare, sqlHeader, sqlTable, typeOf, typeOf, typeOf, vids)
	//fmt.Println("***JK***瓶签", sqlStr)
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content, " ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage}
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
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case when PtTimes = 2 or PtTimes = 3 then '已打印' else '未打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
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
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱拆分-输液卡*/
func SearchSplitMedicalAdviceForInfusion(startTime, endTime, vids string, typeOf, print, did int) ([]MedicalAdviceResponse, error) {
	sqlStr := fmt.Sprintf("DECLARE @lVAF11 INT, @lBCK01 INT, @lType INT DECLARE  @lStop TINYINT DECLARE @lDt1 VARCHAR(30) ,@lDt2 VARCHAR(30), @p262 VARCHAR(10) SET @lBCK01 = %d SET @lVAF11 = %d SET @lDt1 = CONVERT( VARCHAR(16), CAST( '%s' AS DATETIME), 121) + ':00' SET @lDt2 = CONVERT( VARCHAR(16), CAST( '%s' AS DATETIME), 121) + ':59' SET @lType = 0 SET @lStop = 0 SELECT d.* FROM( SELECT a.VAF06 Vid, 1 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, a2.BBX05 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, CASE WHEN a.BDA01 = '3' THEN a1.VAF23 ELSE a.VAF23 END Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF04 = 2 AND b.VBI07 > 0 AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND c.VAE01 IN(%s) AND a2.BBX20 IN(2, 4, 5) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND(a.BDA01 >= '1' AND a.BDA01 <= '3') AND a2.BDA01 = 'T' AND a2.BBX13 IN('2', '4') AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 UNION ALL SELECT a.VAF06 Vid, 1 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '' Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a1.BBX20 = 0 THEN '口服单' WHEN a1.BBX20 = 1 THEN '注射单' WHEN(a1.BBX20 = 2) OR(a1.BBX20 = 4) THEN '输液单' WHEN a1.BBX20 = 3 THEN '治疗单' WHEN a1.BBX20 = 5 THEN '输血单' WHEN a1.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN BBX1 a1 ON a.BBX01 = a1.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.BDA01 IN('B', 'D', 'N', 'O', 'T', 'Z') AND a.VAF32 = 0 AND a.VAF01A = 0 AND( (@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11) ) AND a.VAF04 = 2 AND c.VAE01 IN(%s) AND a1.BBX20 IN(2, 4, 5)  AND( (a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND( (@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8) ) ) ) AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 UNION ALL SELECT a.VAF06 Vid, 1 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '' Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a.VAF53 = 0 THEN '口服单' WHEN a.VAF53 = 1 THEN '注射单' WHEN(a.VAF53 = 2) OR(a.VAF53 = 4) THEN '输液单' WHEN a.VAF53 = 3 THEN '治疗单' WHEN a.VAF53 = 5 THEN '输血单' WHEN a.VAF53 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.BDA01 = '0' AND( (@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11) ) AND a.VAF04 = 2 AND c.VAE01 IN(%s) AND a.VAF53 IN(2, 4, 5)  AND( (a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND( (@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8) ) ) ) AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 ) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", did, typeOf, startTime, endTime, vids, vids, vids)
	mAdvices := make([]MedicalAdviceModal, 0)
	err_ss := fit.SQLServerEngine().SQL(sqlStr).Find(&mAdvices)
	//fmt.Println("***JK***输液卡", sqlStr)

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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content, " ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage}
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
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case when PtTimes = 1 or PtTimes = 3 then '已打印' else '未打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
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
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱拆分-口服单*/
func SearchSplitMedicalAdviceForOralMedical(startTime, endTime, vids string, typeOf, print, did int) ([]MedicalAdviceResponse, error) {
	sqlStr := fmt.Sprintf("DECLARE @lVAF11 INT, @lBCK01 INT, @lType INT DECLARE  @lStop TINYINT DECLARE @lDt1 VARCHAR(30) ,@lDt2 VARCHAR(30), @p262 VARCHAR(10) SET @lBCK01 = %d SET @lVAF11 = %d SET @lDt1 = CONVERT( VARCHAR(16), CAST( '%s' AS DATETIME), 121) + ':00' SET @lDt2 = CONVERT( VARCHAR(16), CAST( '%s' AS DATETIME), 121) + ':59' SET @lType = 2 SET @lStop = 0 SELECT d.* FROM( SELECT a.VAF06 Vid, 2 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, a2.BBX05 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, CASE WHEN a.BDA01 = '3' THEN a1.VAF23 ELSE a.VAF23 END Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF04 = 2 AND b.VBI07 > 0 AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND c.VAE01 IN( %s) AND a2.BBX20 IN(0) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND(a.BDA01 >= '1' AND a.BDA01 <= '3') AND a2.BDA01 = 'T' AND a2.BBX13 IN('2', '4') AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 UNION ALL SELECT a.VAF06 Vid, 2 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '' Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a1.BBX20 = 0 THEN '口服单' WHEN a1.BBX20 = 1 THEN '注射单' WHEN(a1.BBX20 = 2) OR(a1.BBX20 = 4) THEN '输液单' WHEN a1.BBX20 = 3 THEN '治疗单' WHEN a1.BBX20 = 5 THEN '输血单' WHEN a1.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN BBX1 a1 ON a.BBX01 = a1.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.BDA01 IN('B', 'D', 'N', 'O', 'T', 'Z') AND a.VAF32 = 0 AND a.VAF01A = 0 AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND a.VAF04 = 2 AND c.VAE01 IN( %s) AND a1.BBX20 IN(0) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 UNION ALL SELECT a.VAF06 Vid, 2 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '' Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a.VAF53 = 0 THEN '口服单' WHEN a.VAF53 = 1 THEN '注射单' WHEN(a.VAF53 = 2) OR(a.VAF53 = 4) THEN '输液单' WHEN a.VAF53 = 3 THEN '治疗单' WHEN a.VAF53 = 5 THEN '输血单' WHEN a.VAF53 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.BDA01 = '0' AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND a.VAF04 = 2 AND c.VAE01 IN( %s) AND a.VAF53 IN(0) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", did, typeOf, startTime, endTime, vids, vids, vids)
	//fmt.Println("***JK***口服单", sqlStr)
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content, " ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage}
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
		contentObj := resp.Contents[0]
		madid := contentObj.Madid
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
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
	if print == 2 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱拆分-注射单*/
func SearchSplitMedicalAdviceForOralInjection(startTime, endTime, vids string, typeOf, print, did int) ([]MedicalAdviceResponse, error) {
	sqlStr := fmt.Sprintf("DECLARE @lVAF11 INT, @lBCK01 INT, @lType INT DECLARE  @lStop TINYINT DECLARE @lDt1 VARCHAR(30) ,@lDt2 VARCHAR(30), @p262 VARCHAR(10) SET @lBCK01 = %d SET @lVAF11 = %d SET @lDt1 = CONVERT( VARCHAR(16), CAST( '%s' AS DATETIME), 121) + ':00' SET @lDt2 = CONVERT( VARCHAR(16), CAST( '%s' AS DATETIME), 121) + ':59' SET @lType = 2 SET @lStop = 0 SELECT d.* FROM( SELECT a.VAF06 Vid, 3 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, a2.BBX05 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, CASE WHEN a.BDA01 = '3' THEN a1.VAF23 ELSE a.VAF23 END Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF04 = 2 AND b.VBI07 > 0 AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND c.VAE01 IN(%s) AND a2.BBX20 IN(1) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND(a.BDA01 >= '1' AND a.BDA01 <= '3') AND a2.BDA01 = 'T' AND a2.BBX13 IN('2', '4') AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 UNION ALL SELECT a.VAF06 Vid, 3 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '' Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a1.BBX20 = 0 THEN '口服单' WHEN a1.BBX20 = 1 THEN '注射单' WHEN(a1.BBX20 = 2) OR(a1.BBX20 = 4) THEN '输液单' WHEN a1.BBX20 = 3 THEN '治疗单' WHEN a1.BBX20 = 5 THEN '输血单' WHEN a1.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN BBX1 a1 ON a.BBX01 = a1.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.BDA01 IN('B', 'D', 'N', 'O', 'T', 'Z') AND a.VAF32 = 0 AND a.VAF01A = 0 AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND a.VAF04 = 2 AND c.VAE01 IN(%s) AND a1.BBX20 IN(1) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0 UNION ALL SELECT a.VAF06 Vid, 3 PtTypeV, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, '' Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CAST(a.VAF27 AS INT) Times, '' Method, '0' Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a.VAF53 = 0 THEN '口服单' WHEN a.VAF53 = 1 THEN '注射单' WHEN(a.VAF53 = 2) OR(a.VAF53 = 4) THEN '输液单' WHEN a.VAF53 = 3 THEN '治疗单' WHEN a.VAF53 = 5 THEN '输血单' WHEN a.VAF53 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.BDA01 = '0' AND((@lVAF11 = 0) OR(@lVAF11 > 0 AND a.VAF11 = @lVAF11)) AND a.VAF04 = 2 AND c.VAE01 IN(%s) AND a.VAF53 IN(1) AND((a.VAF11 = 2 AND a.VAF10 >= 8) OR( a.VAF11 = 1 AND((@lStop = 0 AND a.VAF10 = 3) OR(@lStop = 1 AND a.VAF10 >= 8)))) AND datediff(MINUTE ,@lDt1, b.VBI10) >= 0 AND datediff(MINUTE, b.VBI10 ,@lDt2) >= 0) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", did, typeOf, startTime, endTime, vids, vids, vids)
	//fmt.Println("***JK***注射单", sqlStr)
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
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content, " ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage})
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
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage}
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
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		if print == 0 && exec.PtTimes == 0 {
			response = append(response, *resp)
		} else if print == 1 && exec.PtTimes >= 1 {
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

	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 1, 0, t.Location()).Format("2006-01-02 15:04:05")

	mAdvices := make([]MedicalAdviceModal, 0)
	//sqlStr := fmt.Sprintf("SELECT d.* FROM( SELECT CASE WHEN(( a2.BBX20 = 2 OR a2.BBX20 = 4 OR a2.BBX20 = 5) AND(a.BDA01 = '1' OR a.BDA01 = '2') AND a2.BDA01 = 'T' AND a2.BBX13 = '2') THEN 4 ELSE 0 END PtTypeV, c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, a.VAF27 Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' WHEN e.BDA01 = 'T' THEN '治疗单' WHEN e.BDA01 = 'N' THEN '护理单' ELSE '其它' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, CASE WHEN a.VAF10 >= 8 THEN a.VAF47 ELSE '' END AS EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF04 = 2 AND a.VAF32 = 0 %s AND b.VBI07 > 0 AND c.VAE01 IN(%s) %s and a.VAF10 = 3 %s) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", condition_catg, vids, condition_type, condition_time)
	sqlStr := fmt.Sprintf("SELECT d.* FROM( SELECT CASE WHEN(( a2.BBX20 = 2 OR a2.BBX20 = 4 OR a2.BBX20 = 5) AND(a.BDA01 = '1' OR a.BDA01 = '2') AND a2.BDA01 = 'T' AND a2.BBX13 = '2') THEN 4 ELSE 0 END PtTypeV, c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 AS VARCHAR(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 AND a.VAF61 > 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 OR b.VBI13 = 9 THEN '已作废' WHEN a.VAF10 >= 8 AND a.VAF11 = 1 THEN '已停' WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) = 0 THEN '未停' WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) < 0 THEN '已停' ELSE '其它' END AS MStatus, CASE WHEN a.VAF10 = 3 THEN 3 WHEN a.VAF10 = 4 OR b.VBI13 = 9 THEN 4 WHEN a.VAF10 >= 8 AND a.VAF11 = 1 THEN 8 WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) = 0 THEN 3 WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) < 0 THEN 8 ELSE a.VAF10 END AS MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' WHEN e.BDA01 = 'T' THEN '治疗单' WHEN e.BDA01 = 'N' THEN '护理单' ELSE '其它' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, CASE WHEN a.VAF10 >= 8 THEN a.VAF47 ELSE '' END AS EdTime, b.BCE03A Sender, b.VBI13 HisExStatus FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE c.VAE01 IN(%s) AND a.VAF04 = 2 AND a.VAF32 = 0%s AND b.VBI07 > 0%s AND(( a.VAF11 = 1 AND a.VAF10 = 3 AND b.VBI13 != 9) OR( a.VAF11 = 2 AND a.VAF10 = 8 AND DATEDIFF( DAY, '%s', b.VBI10) = 0))%s) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", today, today, today, today, vids, condition_catg, condition_type, today, condition_time)
	//fit.Logger().LogDebug("***JK***医嘱执行查询", sqlStr)
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
		Vid:         temp.Vid,
		Pid:         temp.Pid,
		Bed:         temp.Bed,
		PName:       temp.PName,
		Age:         temp.Age,
		HospNum:     temp.HospNum,
		Gender:      temp.Gender,
		ExTime:      temp.ExTime,
		ExDay:       temp.ExDay,
		GroupNum:    temp.GroupNum,
		Amount:      temp.Amount,
		Frequency:   temp.Frequency,
		Times:       temp.Times,
		Method:      temp.Method,
		Speed:       temp.Speed,
		TypeV:       temp.TypeV,
		TypeOf:      temp.TypeOf,
		StTime:      temp.StTime,
		StDay:       temp.StDay,
		MStatus:     temp.MStatus,
		MStatusV:    temp.MStatusV,
		Category:    temp.Category,
		CategoryV:   temp.CategoryV,
		PtType:      temp.PtType,
		PtTypeV:     temp.PtTypeV,
		PtNum:       temp.PtNum,
		PtRownr:     temp.PtRownr,
		Entrust:     temp.Entrust,
		Physician:   temp.Physician,
		EdTime:      temp.EdTime,
		EdDay:       temp.EdDay,
		Sender:      temp.Sender,
		ExCycle:     temp.ExCycle,
		HisExStatus: temp.HisExStatus,
	}
	object.Gid = temp.Madid
	object.Contents = make([]MedicalAdviceContent, 1)
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content, " ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage})
		} else {
			//	按次拆分
			for idx := 1; idx <= object.Times; idx ++ {
				if idx == 1 {
					object.ExCycle = idx
					arrA = append(arrA, object)
				} else {
					obj := MedicalAdviceResponse{
						Vid:         object.Vid,
						Pid:         object.Pid,
						Bed:         object.Bed,
						PName:       object.PName,
						Age:         object.Age,
						HospNum:     object.HospNum,
						Gender:      object.Gender,
						ExTime:      object.ExTime,
						ExDay:       object.ExDay,
						GroupNum:    object.GroupNum,
						Amount:      object.Amount,
						Frequency:   object.Frequency,
						Times:       object.Times,
						Method:      object.Method,
						Speed:       object.Speed,
						TypeV:       object.TypeV,
						TypeOf:      object.TypeOf,
						StTime:      object.StTime,
						StDay:       object.StDay,
						MStatus:     object.MStatus,
						MStatusV:    object.MStatusV,
						Category:    object.Category,
						CategoryV:   object.CategoryV,
						PtType:      object.PtType,
						PtTypeV:     object.PtTypeV,
						PtNum:       object.PtNum,
						PtRownr:     object.PtRownr,
						Entrust:     object.Entrust,
						Physician:   object.Physician,
						EdTime:      object.EdTime,
						EdDay:       object.EdDay,
						Sender:      object.Sender,
						ExCycle:     idx,
						HisExStatus: object.HisExStatus,
					}
					obj.Gid = object.Contents[0].Madid
					obj.Contents = make([]MedicalAdviceContent, 0)
					obj.Contents = append(obj.Contents, object.Contents...)
					arrA = append(arrA, obj)
				}
			}

			object = MedicalAdviceResponse{
				Vid:         v.Vid,
				Pid:         v.Pid,
				Bed:         v.Bed,
				PName:       v.PName,
				Age:         v.Age,
				HospNum:     v.HospNum,
				Gender:      v.Gender,
				ExTime:      v.ExTime,
				ExDay:       v.ExDay,
				GroupNum:    v.GroupNum,
				Amount:      v.Amount,
				Frequency:   v.Frequency,
				Times:       v.Times,
				Method:      v.Method,
				Speed:       v.Speed,
				TypeV:       v.TypeV,
				TypeOf:      v.TypeOf,
				StTime:      v.StTime,
				StDay:       v.StDay,
				MStatus:     v.MStatus,
				MStatusV:    v.MStatusV,
				Category:    v.Category,
				CategoryV:   v.CategoryV,
				PtType:      v.PtType,
				PtTypeV:     v.PtTypeV,
				PtNum:       v.PtNum,
				PtRownr:     v.PtRownr,
				Entrust:     v.Entrust,
				Physician:   v.Physician,
				EdTime:      v.EdTime,
				EdDay:       v.EdDay,
				Sender:      v.Sender,
				ExCycle:     v.ExCycle,
				HisExStatus: v.HisExStatus,
			}
			object.Gid = v.Madid
			object.Contents = make([]MedicalAdviceContent, 1)
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage}
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
					Vid:         object.Vid,
					Pid:         object.Pid,
					Bed:         object.Bed,
					PName:       object.PName,
					Age:         object.Age,
					HospNum:     object.HospNum,
					Gender:      object.Gender,
					ExTime:      object.ExTime,
					ExDay:       object.ExDay,
					GroupNum:    object.GroupNum,
					Amount:      object.Amount,
					Frequency:   object.Frequency,
					Times:       object.Times,
					Method:      object.Method,
					Speed:       object.Speed,
					TypeV:       object.TypeV,
					TypeOf:      object.TypeOf,
					StTime:      object.StTime,
					StDay:       object.StDay,
					MStatus:     object.MStatus,
					MStatusV:    object.MStatusV,
					Category:    object.Category,
					CategoryV:   object.CategoryV,
					PtType:      object.PtType,
					PtTypeV:     object.PtTypeV,
					PtNum:       object.PtNum,
					PtRownr:     object.PtRownr,
					Entrust:     object.Entrust,
					Physician:   object.Physician,
					EdTime:      object.EdTime,
					EdDay:       object.EdDay,
					Sender:      object.Sender,
					ExCycle:     idx,
					HisExStatus: object.HisExStatus,
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
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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

		// 执行状态归类
		if status == 1 && exec.ExStatusV == 1 {
			response = append(response, *resp)
		} else if status == 2 && exec.ExStatusV == 2 {
			response = append(response, *resp)
		} else if status == 3 && exec.ExStatusV == 3 {
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
	if status == 0 {
		return arrA, err_ss
	} else {
		return response, err_ss
	}
}

/*医嘱查询*/
func SearchMedicalAdviceForPC(typeOf, status int, category, vids, st, et string) ([]MedicalAdviceResponse, error) {
	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 1, 0, t.Location()).Format("2006-01-02 15:04:05")

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
	var condition_state string
	if status != 0 {
		// 医嘱状态 未停、已撤销、已停
		if status >= 8 {
			// 已停：>8 or 长嘱=8 or 临嘱：今天以前
			condition_state = fmt.Sprintf(" AND ((a.VAF11 = 1 and a.VAF10 >= 8) or (a.VAF11 = 2 and a.VAF10 >= 8 and DATEDIFF(DAY, '%s', b.VBI10) < 0))", today)
		} else if status == 3 {
			//  未停：长嘱=3，临嘱当天有效
			condition_state = fmt.Sprintf(" AND ((a.VAF11 = 1 and a.VAF10 = 3 and b.VBI13 != 9) or (a.VAF11 = 2 and a.VAF10 = 8 and DATEDIFF(DAY, '%s', b.VBI10) = 0))", today)
		} else if status == 4 {
			condition_state = " AND (a.VAF10 = 4 or b.VBI13 = 9)"
		}
	} else {
		condition_state = fmt.Sprintf(" AND (((a.VAF11 = 1 and a.VAF10 >= 8) or (a.VAF11 = 2 and a.VAF10 >= 8 and DATEDIFF(DAY, '%s', b.VBI10) < 0)) or ((a.VAF11 = 1 and a.VAF10 = 3 and b.VBI13 != 9) or (a.VAF11 = 2 and a.VAF10 = 8 and DATEDIFF(DAY, '%s', b.VBI10) = 0)) or (a.VAF10 = 4 or b.VBI13 = 9))", today, today)
	}

	condition_time := " AND b.VBI10 > c.VAE11"
	if st != "all" && et != "all" {
		condition_time = fmt.Sprintf(" AND b.VBI10 > c.VAE11 AND b.VBI10 BETWEEN '%s' AND '%s'", st, et)
	}

	mAdvices := make([]MedicalAdviceModal, 0)
	//sqlStr := fmt.Sprintf("SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, a.VAF27 Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' WHEN e.BDA01 = 'T' THEN '治疗单' WHEN e.BDA01 = 'N' THEN '护理单' ELSE '其它' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, CASE WHEN a.VAF10 >= 8 THEN a.VAF47 ELSE '' END AS EdTime, b.BCE03A Sender, b.VBI13 HisExStatus FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE c.VAE01 IN(%s) AND a.VAF04 = 2 AND a.VAF32 = 0 %s AND b.VBI07 > 0 %s %s %s) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", vids, condition_catg, condition_type, condition_state, condition_time)
	sqlStr := fmt.Sprintf("SELECT d.* FROM( SELECT c.VAE01 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName, CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, CAST(a.VAF21 AS INT) Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 AND a.VAF61 > 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 OR b.VBI13 = 9 THEN '已作废' WHEN a.VAF10 >= 8 AND a.VAF11 = 1 THEN '已停' WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) = 0 THEN '未停' WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) < 0 THEN '已停' ELSE '其它' END AS MStatus, CASE WHEN a.VAF10 = 3 THEN 3 WHEN a.VAF10 = 4 OR b.VBI13 = 9 THEN 4 WHEN a.VAF10 >= 8 AND a.VAF11 = 1 THEN 8 WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) = 0 THEN 3 WHEN a.VAF10 = 8 AND a.VAF11 = 2 AND DATEDIFF( DAY, '%s', b.VBI10) < 0 THEN 8 ELSE a.VAF10 END AS MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' WHEN e.BDA01 = 'T' THEN '治疗单' WHEN e.BDA01 = 'N' THEN '护理单' ELSE '其它' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, CASE WHEN a.VAF10 >= 8 THEN a.VAF47 ELSE '' END AS EdTime, b.BCE03A Sender, b.VBI13 HisExStatus FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE c.VAE01 IN(%s) AND a.VAF04 = 2 AND a.VAF32 = 0 %s AND b.VBI07 > 0 %s %s %s) d ORDER BY d.Bed, d.TypeV, d.Vid, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", today, today, today, today, vids, condition_catg, condition_type, condition_state, condition_time)
	//fit.Logger().LogDebug("***JK***医嘱查询", sqlStr)
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
		Vid:         temp.Vid,
		Pid:         temp.Pid,
		Bed:         temp.Bed,
		PName:       temp.PName,
		Age:         temp.Age,
		HospNum:     temp.HospNum,
		Gender:      temp.Gender,
		ExTime:      temp.ExTime,
		ExDay:       temp.ExDay,
		GroupNum:    temp.GroupNum,
		Amount:      temp.Amount,
		Frequency:   temp.Frequency,
		Times:       temp.Times,
		Method:      temp.Method,
		Speed:       temp.Speed,
		TypeV:       temp.TypeV,
		TypeOf:      temp.TypeOf,
		StTime:      temp.StTime,
		StDay:       temp.StDay,
		MStatus:     temp.MStatus,
		MStatusV:    temp.MStatusV,
		Category:    temp.Category,
		CategoryV:   temp.CategoryV,
		PtType:      temp.PtType,
		PtNum:       temp.PtNum,
		PtRownr:     temp.PtRownr,
		Entrust:     temp.Entrust,
		Physician:   temp.Physician,
		EdTime:      temp.EdTime,
		EdDay:       temp.EdDay,
		Sender:      temp.Sender,
		ExCycle:     temp.ExCycle,
		HisExStatus: temp.HisExStatus,
	}
	object.Gid = temp.Madid
	object.Contents = make([]MedicalAdviceContent, 1)
	object.Contents[0] = MedicalAdviceContent{temp.Madid, strings.Split(temp.Content, " ")[0], temp.Dosage}

	// 按组合并，最后一次创建的Object不会被拆分
	for i := 1; i < length; i ++ {
		v := mAdvices[i]
		v.ExDay = v.ExTime.ParseToSecond()
		v.StDay = v.StTime.ParseToMinute()
		v.EdDay = v.EdTime.ParseToMinute()

		// 同 人+时+单+组 则合并
		if object.Vid == v.Vid && object.ExTime.ParseToSecond() == v.ExTime.ParseToSecond() && object.PtNum == v.PtNum && object.GroupNum == v.GroupNum {
			object.Contents = append(object.Contents, MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage})
		} else {
			//	按次拆分
			for idx := 1; idx <= object.Times; idx ++ {
				if idx == 1 {
					object.ExCycle = idx
					arrA = append(arrA, object)
				} else {
					obj := MedicalAdviceResponse{
						Vid:         object.Vid,
						Pid:         object.Pid,
						Bed:         object.Bed,
						PName:       object.PName,
						Age:         object.Age,
						HospNum:     object.HospNum,
						Gender:      object.Gender,
						ExTime:      object.ExTime,
						ExDay:       object.ExDay,
						GroupNum:    object.GroupNum,
						Amount:      object.Amount,
						Frequency:   object.Frequency,
						Times:       object.Times,
						Method:      object.Method,
						Speed:       object.Speed,
						TypeV:       object.TypeV,
						TypeOf:      object.TypeOf,
						StTime:      object.StTime,
						StDay:       object.StDay,
						MStatus:     object.MStatus,
						MStatusV:    object.MStatusV,
						Category:    object.Category,
						CategoryV:   object.CategoryV,
						PtType:      object.PtType,
						PtNum:       object.PtNum,
						PtRownr:     object.PtRownr,
						Entrust:     object.Entrust,
						Physician:   object.Physician,
						EdTime:      object.EdTime,
						EdDay:       object.EdDay,
						Sender:      object.Sender,
						ExCycle:     idx,
						HisExStatus: object.HisExStatus,
					}
					obj.Gid = object.Contents[0].Madid
					obj.Contents = make([]MedicalAdviceContent, 0)
					obj.Contents = append(obj.Contents, object.Contents...)
					arrA = append(arrA, obj)
				}
			}

			object = MedicalAdviceResponse{
				Vid:         v.Vid,
				Pid:         v.Pid,
				Bed:         v.Bed,
				PName:       v.PName,
				Age:         v.Age,
				HospNum:     v.HospNum,
				Gender:      v.Gender,
				ExTime:      v.ExTime,
				ExDay:       v.ExDay,
				GroupNum:    v.GroupNum,
				Amount:      v.Amount,
				Frequency:   v.Frequency,
				Times:       v.Times,
				Method:      v.Method,
				Speed:       v.Speed,
				TypeV:       v.TypeV,
				TypeOf:      v.TypeOf,
				StTime:      v.StTime,
				StDay:       v.StDay,
				MStatus:     v.MStatus,
				MStatusV:    v.MStatusV,
				Category:    v.Category,
				CategoryV:   v.CategoryV,
				PtType:      v.PtType,
				PtNum:       v.PtNum,
				PtRownr:     v.PtRownr,
				Entrust:     v.Entrust,
				Physician:   v.Physician,
				EdTime:      v.EdTime,
				EdDay:       v.EdDay,
				Sender:      v.Sender,
				ExCycle:     v.ExCycle,
				HisExStatus: v.HisExStatus,
			}
			object.Gid = v.Madid
			object.Contents = make([]MedicalAdviceContent, 1)
			object.Contents[0] = MedicalAdviceContent{v.Madid, strings.Split(v.Content, " ")[0], v.Dosage}
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
					Vid:         object.Vid,
					Pid:         object.Pid,
					Bed:         object.Bed,
					PName:       object.PName,
					Age:         object.Age,
					HospNum:     object.HospNum,
					Gender:      object.Gender,
					ExTime:      object.ExTime,
					ExDay:       object.ExDay,
					GroupNum:    object.GroupNum,
					Amount:      object.Amount,
					Frequency:   object.Frequency,
					Times:       object.Times,
					Method:      object.Method,
					Speed:       object.Speed,
					TypeV:       object.TypeV,
					TypeOf:      object.TypeOf,
					StTime:      object.StTime,
					StDay:       object.StDay,
					MStatus:     object.MStatus,
					MStatusV:    object.MStatusV,
					Category:    object.Category,
					CategoryV:   object.CategoryV,
					PtType:      object.PtType,
					PtNum:       object.PtNum,
					PtRownr:     object.PtRownr,
					Entrust:     object.Entrust,
					Physician:   object.Physician,
					EdTime:      object.EdTime,
					EdDay:       object.EdDay,
					Sender:      object.Sender,
					ExCycle:     idx,
					HisExStatus: object.HisExStatus,
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
		//content := contentObj.Content
		//dosage := contentObj.Dosage

		exec := MedicalAdvicePrintSubModel{}
		// 用组里的第一个数据代替组 查询执行和打印记录
		_, err_pt := fit.MySqlEngine().SQL("select Madid, ExNurse, case ExStatusV when 0 then '未执行' when 1 then '未执行' when 2 then ExStep when 3 then '已结束' else '未执行' end as ExStatus, ExStatusV, ExStep, PtTimes, case PtTimes when 0 then '未打印' else '已打印' end as PtStatus from medicaladvice where Madid = ? and ExTime = ? and ExCycle = ?", madid, resp.ExTime.ParseToSecond(), resp.ExCycle).Get(&exec)
		if err_pt != nil {
			fit.Logger().LogError("***JK***", "医嘱查询-err_pt", err_pt.Error())
			err_ss = err_pt
			return make([]MedicalAdviceResponse, 0), err_ss
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
	return arrA, err_ss
}

/*获取医嘱执行记录*/
func FetchMedicalAdviceExecutionRecordForPc(madid int64, excycle, extime string) ([]MedicalAdviceExecutionRecord, error) {
	response := make([]MedicalAdviceExecutionRecord, 0)
	err := fit.MySqlEngine().SQL("SELECT DATE_FORMAT(Plan,'%Y-%m-%d %H:%i:%s') Plan, DATE_FORMAT(ExecTime,'%Y-%m-%d %H:%i') ExecTime,AdviceDetail.* from AdviceDetail where Madid = ? and Plan = ? and ExCycle = ?", madid, extime, excycle).Find(&response)
	return response, err
}

/*医嘱打印登记*/
func UpdateMedicalAdvicePrintStatus(gid int64, exc, ext, ptType string, obj MedicalAdviceItem) error {
	// 1:打印,0：未打印,(输液单已打印=1;瓶签已打印＝2;输液单和瓶签都已打印＝3)   输液单和瓶签可能共用数据
	obj.PtTimes = 1
	if ptType == "4" {
		obj.PtTimes = 2
	}

	if isExist := IsExistRecord(true, "medicaladvice", fmt.Sprintf("Madid = %d and ExTime = '%s' and ExCycle = %s ", gid, ext, exc)); isExist.Exist == 0 {
		_, err_in := fit.MySqlEngine().Table("medicaladvice").InsertOne(&obj)
		if err_in != nil {
			fit.Logger().LogError("***JK***", "医嘱打印-Insert", err_in.Error())
		}
		return err_in
	} else {
		if ptType == "1" { // 输液单，判断 瓶签是否已打印，如果已打印赋值 PtTimes = 3
			if isExist := IsExistRecord(true, "medicaladvice", fmt.Sprintf("Madid = %d and ExTime = '%s' and ExCycle = %s and PtTimes = 2", gid, ext, exc)); isExist.Exist >= 1 {
				obj.PtTimes = 3
			}
		} else if ptType == "4" { // 瓶签，判断 输液单是否已打印，如果已打印赋值 PtTimes = 3
			if isExist := IsExistRecord(true, "medicaladvice", fmt.Sprintf("Madid = %d and ExTime = '%s' and ExCycle = %s and PtTimes = 1", gid, ext, exc)); isExist.Exist >= 1 {
				obj.PtTimes = 3
			}
		}

		//	有记录
		_, err_up := fit.MySqlEngine().Exec("UPDATE medicalAdvice SET PtTimes = ? where Madid = ? and ExTime = ? and ExCycle = ?", obj.PtTimes, gid, ext, exc)
		if err_up != nil {
			fit.Logger().LogError("***JK***", "医嘱打印-update", err_up.Error())
		}
		return err_up
	}
}

/*查询该组医嘱是否存在， 病人返回MedicalAdviceItem*/
func CheckingMedicalAdvice(gid int64, ext string, exc int) (MedicalAdviceItem, error) {
	mAdvice := MedicalAdviceModal{}
	_, err_sql := fit.SQLServerEngine().SQL(fmt.Sprintf("SELECT d.* FROM( SELECT a.VAF06 Vid, a.VAF01 MadId, c.VAA01 Pid, c.BCQ04B Bed, c.VAE95 PName,CAST(c.VAE46 as varchar(10)) Age, c.VAE94 HospNum, CASE c.VAE96 WHEN 1 THEN '男' WHEN 2 THEN '女' ELSE '未知' END AS Gender, b.VBI10 ExTime, a.VAF59 GroupNum, a.VAF22 Content, CASE a.VAF19 WHEN '' THEN '-' ELSE a.VAF19 END Dosage, a.VAF21 Amount, a.VAF26 Frequency, CASE WHEN DATEDIFF(DAY, a.VAF36, b.VBI10) = 0 AND a.VAF61 > 0 THEN CAST(a.VAF61 AS INT) ELSE CAST(a.VAF27 AS INT) END Times, a1.VAF22 Method, a.VAF60 Speed, a.VAF11 TypeV, CASE a.VAF11 WHEN 1 THEN '长嘱' WHEN 2 THEN '临嘱' END AS TypeOf, a.VAF36 StTime, CASE WHEN a.VAF10 = 3 THEN '未停' WHEN a.VAF10 = 4 THEN '已作废' WHEN a.VAF10 >= 8 THEN '已停' ELSE '其它' END AS MStatus, a.VAF10 MStatusV, e.BDA02 Category, a.BDA01 CategoryV, CASE WHEN a2.BBX20 = 0 THEN '口服单' WHEN a2.BBX20 = 1 THEN '注射单' WHEN(a2.BBX20 = 2) OR(a2.BBX20 = 4) THEN '输液单' WHEN a2.BBX20 = 3 THEN '治疗单' WHEN a2.BBX20 = 5 THEN '输血单' WHEN a2.BBX20 = 6 THEN '护理单' END AS PtType, a.CBM01 PtNum, a.Rownr PtRownr, a.VAF23 Entrust, a.BCE03A Physician, a.VAF47 EdTime, b.BCE03A Sender FROM VAF2 a LEFT JOIN VAF2 a1 ON a.VAF01A = a1.VAF01 LEFT JOIN BBX1 a2 ON a1.BBX01 = a2.BBX01 JOIN VBI2 b ON a.VAF01 = b.VAF01 JOIN VAE1 c ON a.VAF06 = c.VAE01 LEFT JOIN BDA1 e ON a.BDA01 = e.BDA01 WHERE a.VAF01 = %d AND b.VBI10 = '%s') d ORDER BY d.TypeV, d.ExTime, d.PtNum, d.GroupNum, d.PtRownr", gid, ext)).Get(&mAdvice)
	if err_sql != nil {
		fit.Logger().LogError("***JK***", err_sql)
		return MedicalAdviceItem{}, err_sql
	}
	obj := MedicalAdviceItem{
		Vid:       mAdvice.Vid,
		Madid:     mAdvice.Madid,
		Pid:       mAdvice.Pid,
		Bed:       mAdvice.Bed,
		PName:     mAdvice.PName,
		Gender:    mAdvice.Gender,
		Age:       mAdvice.Age,
		HospNum:   mAdvice.HospNum,
		ExTime:    ext,
		GroupNum:  mAdvice.GroupNum,
		Content:   mAdvice.Content,
		Dosage:    mAdvice.Dosage,
		Amount:    mAdvice.Amount,
		Frequency: mAdvice.Frequency,
		Times:     mAdvice.Times,
		Method:    mAdvice.Method,
		Speed:     mAdvice.Speed,
		TypeV:     mAdvice.TypeV,
		StTime:    mAdvice.StTime.ParseToSecond(),
		MStatusV:  mAdvice.MStatusV,
		Category:  mAdvice.Category,
		CategoryV: mAdvice.CategoryV,
		PtType:    mAdvice.PtType,
		PtNum:     mAdvice.PtNum,
		PtRownr:   mAdvice.PtRownr,
		Entrust:   mAdvice.Entrust,
		Physician: mAdvice.Physician,
		EdTime:    mAdvice.EdTime.ParseToSecond(),
		Sender:    mAdvice.Sender,
		ExCycle:   exc,
		ExNurse:   "",
		ExStatusV: 0,
		ExStep:    "",
		PtTimes:   0,
	}
	return obj, err_sql
}
