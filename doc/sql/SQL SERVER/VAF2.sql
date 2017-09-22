CREATE TABLE VAF2 --住院病人医嘱记录
(
	VAF01 INT --ID
	,
	VAF01A INT --相关ID, 关联字段：VAF1.VAF01
	,
	VAF01B INT --前提ID, 关联字段：VAF1.VAF01
	,
	VAF04 TINYINT --1：门诊;2：住院
	,
	VAA01 INT --病人ID, 关联字段：VAA1.VAA01
	,
	VAF06 INT --就诊ID, 主页ID
	,
	VAF07 INT --婴儿ID
	,
	BCK01A INT --病人科室ID, 关联字段：BCK1.BCK01
	,
	ROWNR INT --次序
	,
	VAF10 TINYINT --1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果
	,
	VAF11 TINYINT --医嘱类型, 1=长期医嘱, 2=临时医嘱
	,
	BDA01 VARCHAR (2) --诊疗类型, 关联字段：BDA1.BDA01
	,
	BBX01 INT --诊疗项目ID, 关联字段：BBX1.BBX01
	,
	VAF14 VARCHAR (60) --标本部位
	,
	VAF15 VARCHAR (30) --检查方法
	,
	BBY01 INT --收费项目ID, 关联字段：BBY1.BBY01
	,
	VAF17 INT --天数, day number
	,
	VAF18 NUMERIC (18, 4) --剂量, 单次用量
	,
	VAF19 VARCHAR (10) --用量
	,
	VAF20 NUMERIC (18, 4) --单量
	,
	VAF21 NUMERIC (18, 4) --数量
	,
	VAF22 VARCHAR (1024) --医嘱
	,
	VAF23 VARCHAR (128) --医师嘱托
	,
	BCK01B INT --执行科室ID, 关联字段：BCK1.BCK01
	,
	VAF25 VARCHAR (10) --空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果
	,
	VAF26 VARCHAR (20) --执行频次
	,
	VAF27 INT --频率次数
	,
	VAF28 TINYINT --频率间隔
	,
	VAF29 VARCHAR (4) --间隔单位
	,
	VAF30 VARCHAR (64) --执行时间方案
	,
	VAF31 TINYINT --计价特性  0=正常, 1=自费
	,
	VAF32 TINYINT --0：正常; 1＝给药途径
	,
	VAF33 TINYINT --0：标记未用;1：正常 2：自动停止
	,
	VAF34 TINYINT --可否分零
	,
	VAF35 TINYINT --0：正常 1：紧急
	,
	VAF36 DATETIME --开始执行时间
	,
	VAF37 DATETIME --执行终止时间
	,
	VAF38 DATETIME --上次执行时间
	,
	BCK01C INT --开嘱科室ID, 关联字段：BCK1.BCK01
	,
	BCE02A VARCHAR (20) --医师编码, 关联字段：BCE1.BCE02
	,
	BCE03A VARCHAR (20) --开嘱医师, 关联字段：BCE1.BCE03
	,
	VAF42 DATETIME --开嘱时间
	,
	BCE03B VARCHAR (20) --开嘱护士, 关联字段：BCE1.BCE03
	,
	BCE03C VARCHAR (20) --校对护士, 关联字段：BCE1.BCE03
	,
	VAF45 DATETIME --校对时间
	,
	BCE03D VARCHAR (20) --停嘱医生, 关联字段：BCE1.BCE03
	,
	VAF47 DATETIME --停嘱时间
	,
	BCE03E VARCHAR (20) --停嘱护士, 关联字段：BCE1.BCE03
	,
	BCE03F VARCHAR (20) --停嘱校对护士, 关联字段：BCE1.BCE03
	,
	VAF50 DATETIME --执行停嘱时间
	,
	VAF51 INT --申请ID
	,
	VAF52 TINYINT --0：新开；1：上传
	,
	VAF53 INT --审查结果，用于药品合理用药审核。(描述性医嘱：执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它)
	,
	VAF54 TINYINT --0：否，1：忽略
	,
	VAF55 VARCHAR (1024) --摘要，医嘱备注
	,
	CBM01 INT --医嘱单id, 关联字段：CBM1.CBM01
	,
	BCK01D INT --给药科室, 关联字段：BCK1.BCK01
	,
	VAF58 TINYINT --0：正常， 1：自备药，2：离院带药
	,
	VAF59 INT --组号
	,
	VAF60 VARCHAR (10) --滴速
	,
	VAF61 NUMERIC (8, 2) --首日执行次数
	,
	VAF62 NUMERIC (8, 2) --末日执行次数
	,
	BCE01A INT --开嘱医师ID, 关联字段：BCE1.BCE01
	,
	BCE01B INT --开嘱护士ID, 关联字段：BCE1.BCE01
	,
	BCE01C INT --校对护士ID, 关联字段：BCE1.BCE01
	,
	BCE01D INT --停嘱医师ID, 关联字段：BCE1.BCE01
	,
	BCE01E INT --停嘱护士ID, 关联字段：BCE1.BCE01
	,
	BCE01F INT --停嘱校对护士ID, 关联字段：BCE1.BCE01
	,
	BCE01G INT --操作员ID, 关联字段：BCE1.BCE01
	,
	BCE03G VARCHAR (20) --操作员, 关联字段：BCE1.BCE03
	,
	VAF71 DATETIME --审核时间
	,
	DSK01 INT --药品批次id DSK_ID
	,
	VAF01C INT --原医嘱id  (-1 = 重整医嘱)
	,
	VAF74 DATETIME --重整医嘱时间
	,
	VAF75 TINYINT --药品用药标识
	,
	BCE01H INT --授权医师id, 关联字段：BCE1.BCE01
	,
	BCE03H VARCHAR (20) --授权医师, 关联字段：BCE1.BCE03
	,
	BIW02 VARCHAR (64) --用药目的, 关联字段：BIW1.BIW02
)