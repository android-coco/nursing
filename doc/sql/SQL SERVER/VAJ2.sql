CREATE TABLE VAJ2 --住院病人费用明细
(
	VAJ01 INT --ID
	,
	VAA01 INT --病人ID, 关联字段：VAA1.VAA01
	,
	VAA07 INT --就诊ID, 主页ID, 关联字段：VAA1.VAA07
	,
	VAJ04 INT --类型, 1=门诊挂号, 2=门诊划价, 3=门诊记帐, 4=门诊收费, 门诊结帐, 5=医技记帐, 6=住院记帐
	,
	VAJ05 INT --记录状态：1：收费划价/记账划价；2：已收费/已记账；3：已退费/已销账  ; 4：退费记录  是根据类型变化 ; 9：作废
	,
	ROWNR INT --次序
	,
	VAJ01A INT --上级ID, 从属父号, 关联字段：VAJ1.VAJ01A
	,
	VAJ01B INT --关联ID, 关联字段：VAJ1.VAJ01B
	,
	VAJ09 INT --冲销ID
	,
	VAJ10 INT --合并标志, 0=否, 1=是, merge tag, 多帐单合并, 多病人单
	,
	VAI01 INT --单据ID, 记帐单ID, 关联字段：VAI1.VAI01
	,
	VAF01 INT --医嘱ID, OrderID, 关联字段：VAF1.VAF01
	,
	VAK01 INT --结帐ID, 关联字段：VAK1.VAK01
	,
	ACF01 INT --医疗服务, 不能取0,3值, 关联字段：ACF1.ACF01
	,
	VAJ15 TINYINT --记帐标志
	,
	BCK01A INT --病区ID, 关联字段：BCK1.BCK01
	,
	BCK01B INT --科室ID, 关联字段：BCK1.BCK01
	,
	BDN01 VARCHAR (2) --类型, 编码, 关联字段：BDN1.BDN01
	,
	BBY01 INT --收费项目ID, 关联字段：BBY1.BBY01
	,
	BCJ02 VARCHAR (32) --发药窗口, 关联字段：BCJ1.BCJ02
	,
	VAJ21 TINYINT --加班标志, 暂时不用
	,
	VAJ22 TINYINT --特殊标志, 附加标志, 不同位置用途不同, 挂号时：存储项目特性(1=挂号, 6=诊金, 7=病历本, 8=就诊卡)
	,
	VAJ23 TINYINT --剂数, 中药剂数
	,
	VAJ24 NUMERIC (18, 4) --单量
	,
	VAJ25 NUMERIC (18, 4) --数量, 数次, 总数量
	,
	VAJ26 TINYINT --急诊标志, emergency tag
	,
	VAJ27 INT --婴儿费,对应VAP1表中VAP01
	,
	VAJ28 NUMERIC (9, 4) --税率, 暂时不用
	,
	VAJ29 NUMERIC (18, 4) --税费, 暂时不用
	,
	VAJ30 NUMERIC (9, 4) --折扣率分子, discount rate numerator
	,
	VAJ31 NUMERIC (9, 4) --折扣率分母, discount rate denominator
	,
	VAJ32 NUMERIC (18, 6) --全价
	,
	VAJ33 NUMERIC (18, 6) --单价, 标准单价
	,
	VAJ34 NUMERIC (18, 4) --包装
	,
	VAJ35 VARCHAR (20) --单位, 计算单位
	,
	VAJ36 NUMERIC (18, 2) --全额,原始价格计算得金额
	,
	VAJ37 NUMERIC (18, 2) --应收金额, 未临时打折前的金额(可能经过费别打折)
	,
	VAJ38 NUMERIC (18, 2) --结帐金额(结账时应付金额)，发票打印以此金额为准
	,
	VAJ39 TINYINT --费用标志, 0=正常, 1=自费, 2=免费
	,
	VAJ40 NUMERIC (18, 4) --自负金额
	,
	VAJ41 NUMERIC (18, 4) --保险金额, 统筹金额
	,
	BCE03A VARCHAR (20) --划价人, 关联字段：BCE1.BCE03
	,
	BCK01C INT --开单科室ID, OrderDeptID, 关联字段：BCK1.BCK01
	,
	BCE02B VARCHAR (20) --开单人号, 关联字段：BCE1.BCE02
	,
	BCE03B VARCHAR (20) --开单人, Physician, 一般为医师, 关联字段：BCE1.BCE03
	,
	VAJ46 DATETIME --记帐时间, 手工时间
	,
	VAJ47 DATETIME --交易时间, 机器时间
	,
	VAJ48 INT --执行ID
	,
	BCK01D INT --执行科室ID, 关联字段：BCK1.BCK01
	,
	BCE03C VARCHAR (20) --执行者, 关联字段：BCE1.BCE03
	,
	VAJ51 DATETIME --执行时间
	,
	VAJ52 DATETIME --执行交易时间
	,
	VAJ53 TINYINT --执行情况：0：未执行; 1：执行完成; 2：拒绝执行; 3：正在执行;4：过期挂起
	,
	VAJ54 VARCHAR (255) --备注
	,
	BCE02D VARCHAR (20) --操作员#, 关联字段：BCE1.BCE02
	,
	BCE03D VARCHAR (20) --操作员, 关联字段：BCE1.BCE03
	,
	VAJ57 VARCHAR (1024) --摘要;  收费项目为主从项目时  摘要=主项目名称
	,
	FAB03 VARCHAR (20) --销售单位, 药品门诊或住院单位, 原先为发票号, 关联字段：FAB1.FAB03
	,
	VAJ59 NUMERIC (18, 6) --成本价
	,
	BCE02C VARCHAR (20) --执行者号
	,
	VAJ61 NUMERIC (18, 2) --核算金额，财务核算时用到
	,
	VAJ62 DATETIME --业务时间、默认记账时间，销账时取被销账那条明细的记账时间
	,
	BCK01E INT --给药科室ID, 关联字段：BCK1.BCK01
	,
	VAJ64 DATETIME --发生时间、用于住院长嘱发送时记跨天的费用
	,
	VAJ65 TINYINT --住院中途结帐时，为1参与本次结帐，否则不参与
	,
	DSK01 INT --药品批次id DSK_ID
	,
	VAJ67 NUMERIC (18, 6) --原价
	,
	BCE01E INT --住院医师id, 关联字段：BCE1.BCE01
	,
	BCE03E VARCHAR (20) --住院医师, 关联字段：BCE1.BCE03
	,
	BCK01F INT --病人床位对应病区, 关联字段：BCK1.BCK01
	,
	BCQ04 VARCHAR (20) --病人床号, 关联字段：BCQ1.BCQ04
)