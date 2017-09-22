CREATE TABLE BCQ1 --病床编制表
(
	BCQ01 INT --ID
	,
	BCK01A INT --病区ID
	,
	BCQ03 VARCHAR (10) --病房, 病室, 房间
	,
	BCQ04 VARCHAR (20) --床位, 床号, 病床
	,
	BCK01B INT --科室ID
	,
	ABW01 VARCHAR (1) --性别, 性别限制
	,
	ACG02 VARCHAR (20) --编制性质
	,
	BBY01 INT --收费项目ID, 等级
	,
	BCQ09 TINYINT --共用 1：共用
	,
	ROWNR INT --次序
	,
	BCQ11 VARCHAR (128) --说明
	,
	VAA01 INT --病人ID, 这个不用显示维护
	,
	BCQ13 TINYINT --状态, 0=空床, 1=占用, 2=陪床(已占用),3=包房, 4=维修,5=撤编 CWBST
	,
	BBY01B INT --包床项目
	,
	BCQ15 TINYINT --0=普通, 1=CCU, 2=挂床
	,
	BBY01C INT --标准收费项目id
)