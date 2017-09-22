#医嘱表
CREATE TABLE VAF2 (
	VAF01 INT COMMENT 'ID',
	VAF01A INT COMMENT '相关ID, 关联字段：VAF1.VAF01',
	VAF01B INT COMMENT '前提ID, 关联字段：VAF1.VAF01',
	VAF04 TINYINT COMMENT '1：门诊;2：住院',
	VAA01 INT COMMENT '病人ID, 关联字段：VAA1.VAA01',
	VAF06 INT COMMENT '就诊ID, 主页ID',
	VAF07 INT COMMENT '婴儿ID',
	BCK01A INT COMMENT '病人科室ID, 关联字段：BCK1.sql.BCK01',
	ROWNR INT COMMENT '次序',
	VAF10 TINYINT COMMENT '1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果',
	VAF11 TINYINT COMMENT '医嘱类型, 1=长期医嘱, 2=临时医嘱',
	BDA01 VARCHAR (2) COMMENT '诊疗类型, 关联字段：BDA1.BDA01',
	BBX01 INT COMMENT '诊疗项目ID, 关联字段：BBX1.BBX01',
	VAF14 VARCHAR (60) COMMENT '标本部位',
	VAF15 VARCHAR (30) COMMENT '检查方法',
	BBY01 INT COMMENT '收费项目ID, 关联字段：BBY1.BBY01',
	VAF17 INT COMMENT '天数, day number',
	VAF18 NUMERIC (18, 4) COMMENT '剂量, 单次用量',
	VAF19 VARCHAR (10) COMMENT '用量',
	VAF20 NUMERIC (18, 4) COMMENT '单量',
	VAF21 NUMERIC (18, 4) COMMENT '数量',
	VAF22 VARCHAR (1024) COMMENT '医嘱',
	VAF23 VARCHAR (128) COMMENT '医师嘱托',
	BCK01B INT COMMENT '执行科室ID, 关联字段：BCK1.sql.BCK01',
	VAF25 VARCHAR (10) COMMENT '空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果',
	VAF26 VARCHAR (20) COMMENT '执行频次',
	VAF27 INT COMMENT '频率次数',
	VAF28 TINYINT COMMENT '频率间隔',
	VAF29 VARCHAR (4) COMMENT '间隔单位',
	VAF30 VARCHAR (64) COMMENT '执行时间方案',
	VAF31 TINYINT COMMENT '计价特性  0=正常, 1=自费',
	VAF32 TINYINT COMMENT '0：正常; 1＝给药途径',
	VAF33 TINYINT COMMENT '0：标记未用;1：正常 2：自动停止',
	VAF34 TINYINT COMMENT '可否分零',
	VAF35 TINYINT COMMENT '0：正常 1：紧急',
	VAF36 DATETIME COMMENT '开始执行时间',
	VAF37 DATETIME COMMENT '执行终止时间',
	VAF38 DATETIME COMMENT '上次执行时间',
	BCK01C INT COMMENT '开嘱科室ID, 关联字段：BCK1.sql.BCK01',
	BCE02A VARCHAR (20) COMMENT '医师编码, 关联字段：BCE1.BCE02',
	BCE03A VARCHAR (20) COMMENT '开嘱医师, 关联字段：BCE1.BCE03',
	VAF42 DATETIME COMMENT '开嘱时间',
	BCE03B VARCHAR (20) COMMENT '开嘱护士, 关联字段：BCE1.BCE03',
	BCE03C VARCHAR (20) COMMENT '校对护士, 关联字段：BCE1.BCE03',
	VAF45 DATETIME COMMENT '校对时间',
	BCE03D VARCHAR (20) COMMENT '停嘱医生, 关联字段：BCE1.BCE03',
	VAF47 DATETIME COMMENT '停嘱时间',
	BCE03E VARCHAR (20) COMMENT '停嘱护士, 关联字段：BCE1.BCE03',
	BCE03F VARCHAR (20) COMMENT '停嘱校对护士, 关联字段：BCE1.BCE03',
	VAF50 DATETIME COMMENT '执行停嘱时间',
	VAF51 INT COMMENT '申请ID',
	VAF52 TINYINT COMMENT '0：新开；1：上传',
	VAF53 INT COMMENT '审查结果，用于药品合理用药审核。(描述性医嘱：执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它)',
	VAF54 TINYINT COMMENT '0：否，1：忽略',
	VAF55 VARCHAR (1024) COMMENT '摘要，医嘱备注',
	CBM01 INT COMMENT '医嘱单id, 关联字段：CBM1.CBM01',
	BCK01D INT COMMENT '给药科室, 关联字段：BCK1.sql.BCK01',
	VAF58 TINYINT COMMENT '0：正常， 1：自备药，2：离院带药',
	VAF59 INT COMMENT '组号',
	VAF60 VARCHAR (10) COMMENT '滴速',
	VAF61 NUMERIC (8, 2) COMMENT '首日执行次数',
	VAF62 NUMERIC (8, 2) COMMENT '末日执行次数',
	BCE01A INT COMMENT '开嘱医师ID, 关联字段：BCE1.BCE01',
	BCE01B INT COMMENT '开嘱护士ID, 关联字段：BCE1.BCE01',
	BCE01C INT COMMENT '校对护士ID, 关联字段：BCE1.BCE01',
	BCE01D INT COMMENT '停嘱医师ID, 关联字段：BCE1.BCE01',
	BCE01E INT COMMENT '停嘱护士ID, 关联字段：BCE1.BCE01',
	BCE01F INT COMMENT '停嘱校对护士ID, 关联字段：BCE1.BCE01',
	BCE01G INT COMMENT '操作员ID, 关联字段：BCE1.BCE01',
	BCE03G VARCHAR (20) COMMENT '操作员, 关联字段：BCE1.BCE03',
	VAF71 DATETIME COMMENT '审核时间',
	DSK01 INT COMMENT '药品批次id DSK_ID',
	VAF01C INT COMMENT '原医嘱id  (-1 = 重整医嘱)',
	VAF74 DATETIME COMMENT '重整医嘱时间',
	VAF75 TINYINT COMMENT '药品用药标识',
	BCE01H INT COMMENT '授权医师id, 关联字段：BCE1.BCE01',
	BCE03H VARCHAR (20) COMMENT '授权医师, 关联字段：BCE1.BCE03',
	BIW02 VARCHAR (64) COMMENT '用药目的, 关联字段：BIW1.BIW02'
)