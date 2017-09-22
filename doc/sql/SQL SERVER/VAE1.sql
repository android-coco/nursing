CREATE TABLE VAE1 --病人登记记录
(
	VAE01 INT --ID, 登记ID, CaseID
	,
	VAE02 VARCHAR (20) --登记号, CaseNo
	,
	VAA01 INT --病人ID, PatID, PatientID
	,
	VAE04 TINYINT --病人性质, 0=住院, 1=门诊留观, 2=住院留观
	,
	ABJ01 VARCHAR (2) --医疗付款方式
	,
	BDP02 VARCHAR (50) --病人类别
	,
	ABC02 VARCHAR (20) --折扣费别
	,
	VAE08 TINYINT --再入院
	,
	BCK01A INT --入院病区ID
	,
	BCK01B INT --入院科室ID
	,
	VAE11 DATETIME --入院日期
	,
	ABO01 VARCHAR (2) --入院病情
	,
	ABR01 VARCHAR (1) --入院方式
	,
	ABT02 VARCHAR (20) --就诊方式
	,
	VAE15 VARCHAR (64) --转院名称
	,
	ABZ02 VARCHAR (10) --入院待遇
	,
	ABK02 VARCHAR (20) --住院目的
	,
	BCQ04A VARCHAR (20) --入院床位
	,
	VAE19 INT --住院次数
	,
	VAE20 TINYINT --陪伴, accompany personnel
	,
	AAG01 INT --对应收费项目中为护理类型的项目id(BBY01)
	,
	VAE22 VARCHAR (2) --住院病情
	,
	BCK01C INT --住院病区ID
	,
	BCK01D INT --住院科室ID
	,
	BCQ04B VARCHAR (20) --住院床位
	,
	VAE26 DATETIME --出院日期
	,
	VAE27 NUMERIC (12, 2) --住院天数
	,
	ABV01 VARCHAR (2) --出院类型, 出院方式 0：正常；1：转院；2：死亡
	,
	VAE29 TINYINT --确诊, Accurate Diagnosis  1：确诊  0：疑诊
	,
	VAE30 DATETIME --确诊日期
	,
	VAE31 TINYINT --新发肿瘤, newly diagnosed tumor
	,
	VAE33 INT --抢救次数
	,
	VAE34 INT --成功次数
	,
	VAE35 TINYINT --随诊标志
	,
	VAE36 INT --随诊期限
	,
	VAE37 VARCHAR (2) --随诊单位 Y表示年, M表示月, D表示天, W表示周
	,
	VAE38 TINYINT --尸检标志 1=是 2＝否
	,
	BCE03A VARCHAR (20) --门诊医师
	,
	BCE03B VARCHAR (20) --责任护士/管床护士
	,
	BCE03C VARCHAR (20) --住院医师
	,
	VAE42 VARCHAR (64) --疾病病种
	,
	VAE44 TINYINT -- 0=预约, 1=入院, 2=在院, 3=转科, 4=预出院, 5=结账,6=离院(未结账)9=取消
	,
	VAE45 NUMERIC (9, 0) --金额
	,
	VAE46 INT --年龄
	,
	AAU01 VARCHAR (1) --年龄单位
	,
	ACK01 VARCHAR (2) --婚姻状况
	,
	AAT02 VARCHAR (128) --职业, Occupation
	,
	ACC02 VARCHAR (32) --国籍, Nationality
	,
	AAY02 VARCHAR (48) --学历
	,
	BAQ01 INT --合约单位ID
	,
	BAQ02 VARCHAR (10) --单位编码, CoCode
	,
	BAQ03 VARCHAR (64) --工作单位, CoName
	,
	VAE55 VARCHAR (32) --单位电话, CoTelephone
	,
	VAE56 VARCHAR (64) --单位地址, CoAddress
	,
	VAE57 VARCHAR (128) --户口地址, 户籍地址, RegisteredAddress
	,
	VAE58 VARCHAR (6) --户籍邮编, RegisteredPostCode
	,
	VAE59 VARCHAR (20) --户籍电话, RegisteredPhone
	,
	VAE60 VARCHAR (20) --(现住址)省(区、市)
	,
	VAE61 VARCHAR (20) --(现住址)市
	,
	VAE62 VARCHAR (64) --(现住址)地址, 常住地址, 街道, 乡镇村, resident address
	,
	VAE63 VARCHAR (20) --电话
	,
	VAE64 VARCHAR (13) --移动电话
	,
	VAE65 VARCHAR (128) --电子邮箱
	,
	VAE66 VARCHAR (64) --其他联系方式, 如QQ号, MSN号, Skype号等, OtherCommunications
	,
	VAE67 VARCHAR (64) --监护人, Guardian
	,
	VAE68 VARCHAR (20) --联系人姓名, contact person name
	,
	AAZ02 VARCHAR (32) --与病人关系, RelationShip
	,
	VAE70 VARCHAR (64) --联系人地址, Contact person address
	,
	VAE71 VARCHAR (20) --联系人电话, contact person telephone
	,
	VAE72 VARCHAR (16) --联系人移动电话, Contact person Mobile Phone
	,
	VAE73 VARCHAR (4) --中医治疗类别
	,
	IAA01 INT --保险机构
	,
	UAA01 INT --社区
	,
	VAE76 TINYINT --审核
	,
	BCE03D VARCHAR (20) --审核人
	,
	VAE78 DATETIME --审核日期
	,
	VAE79 TINYINT --传染病上传标志
	,
	VAE80 TINYINT --转出标志
	,
	BCE03E VARCHAR (20) --登记人
	,
	VAE82 DATETIME --制表时间
	,
	VAE83 VARCHAR (255) --备注
	,
	VAE84 INT --档案状态
	,
	BCE02C VARCHAR (20) --住院医师编码
	,
	VAE85 DATETIME --存档时间
	,
	VAE86 VARCHAR (32) --介绍人
	,
	VAE88 DATETIME --说明：超过此时间，禁止对此病人的一切操作，包括发药、记账，下达医嘱，医技操作，病人结帐等。
	,
	VAE87 VARCHAR (20) --存储格式：000Y00M00D00H00N00W
	,
	SCF01 INT --CRM预约ID
	,
	VAE89 NUMERIC (9, 3) --入院时体重(g)
	,
	VAE90 NUMERIC (9, 3) --新生儿出生体重(g)
	,
	VAE91 NUMERIC (9, 2) --身高(cm)
	,
	VAE92 VARCHAR (64) --转院医疗机构
	,
	VAA08 VARCHAR (20) --病案号'
)