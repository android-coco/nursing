CREATE TABLE FHB1 --病人信息扩展表
(
	FHB01 VARCHAR (36) --ID
	,
	VAA01 INT --病人ID
	,
	ACF01 INT --业务类型：0=挂号；1=门诊；2=住院
	,
	VAA07 INT --就诊ID
	,
	VAK01 INT --结帐ID
	,
	VAI01 INT --单据ID
	,
	VAJ01 INT --明细ID
	,
	BBX01 INT --诊疗ID
	,
	BBY01 INT --收费项目ID
	,
	BCE03 VARCHAR (20) --操作员
	,
	FHB11 DATETIME --操作时间
	,
	FHB12 TINYINT --性质：1=需要隐藏的项目
	,
	FHB13 INT --状态：奇数=隐藏；偶数=还原
)