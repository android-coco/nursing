CREATE TABLE VAA1 --病人资料
(
	VAA01 INT --病人ID, PATFU, PATID
	,
	VAA02 VARCHAR (64) --会员卡号
	,
	VAA03 VARCHAR (20) --门诊号 Clinic Patient No, Out patient No, OPTNO
	,
	VAA04 VARCHAR (20) --住院号, In Patient No, IPTNO
	,
	VAA05 VARCHAR (64) --姓名, PATNM, FirstName + Middlename
	,
	VAA06 VARCHAR (64) --监护人1
	,
	ABBRP VARCHAR (10) --拼音
	,
	ABBRW VARCHAR (10) --五笔
	,
	ABW01 VARCHAR (1) --性别
	,
	VAA10 INT --年龄
	,
	AAU01 VARCHAR (1) --年龄单位
	,
	VAA12 DATETIME --出生日期, birth date
	,
	ACK01 VARCHAR (2) --婚姻, marital status
	,
	VAA14 VARCHAR (2) --身份证件, 指公安机关签发的有效证件
	,
	VAA15 VARCHAR (20) --身份证号, IDNo, IDNumber
	,
	VAA16 VARCHAR (20) --其他证件
	,
	ABJ01 VARCHAR (2) --付款方式, 医疗付款方式
	,
	BDP02 VARCHAR (50) --病人类别
	,
	ABC02 VARCHAR (20) --病人费别
	,
	VAA20 VARCHAR (64) --出生地点, 出生地, BirthPlace
	,
	ACM02 VARCHAR (20) --从业状况, 身份
	,
	ACC02 VARCHAR (32) --国籍, Nationality
	,
	ABQ02 VARCHAR (32) --民族, Nation
	,
	VAA25 VARCHAR (64) --籍贯, 省.市, NativePlace
	,
	VAA26 VARCHAR (20) --宗教, 佛教, 伊斯兰教, 基督教, 天主教, 犹太教, 耶稣教, 其他, Religion, 暂时不用
	,
	VAA27 VARCHAR (20) --种族, 人种, Ethnic, 暂时不用
	,
	VAA28 VARCHAR (128) --户口地址, 户籍地址, RegisteredAddress
	,
	VAA29 VARCHAR (6) --户籍邮编, RegisteredPostCode
	,
	VAA30 VARCHAR (20) --户籍电话, RegisteredPhone
	,
	VAA31 VARCHAR (20) --省市
	,
	VAA32 VARCHAR (20) --县市, 城市, 县区市
	,
	VAA33 VARCHAR (64) --地址, 常住地址, resident address
	,
	VAA34 VARCHAR (20) --电话
	,
	VAA35 VARCHAR (13) --移动电话
	,
	VAA36 VARCHAR (128) --电子邮箱
	,
	VAA37 VARCHAR (64) --其他联系方式, 如QQ号, MSN号, Skype号等, OtherCommunications
	,
	VAA38 VARCHAR (48) --学历
	,
	VAA39 VARCHAR (64) --监护人2, Guardian
	,
	VAA40 VARCHAR (20) --联系人姓名, contact person name
	,
	VAA41 VARCHAR (32) --与病人关系, RelationShip
	,
	VAA42 VARCHAR (64) --联系人地址, Contact person address
	,
	VAA43 VARCHAR (20) --联系人电话, contact person telephone
	,
	VAA44 VARCHAR (16) --联系人移动电话, Contact person Mobile Phone
	,
	BAQ01 INT --合同单位ID
	,
	BAQ02 VARCHAR (10) --单位编码, CoCode
	,
	VAA47 VARCHAR (64) --工作单位, CoName
	,
	VAA48 VARCHAR (20) --单位电话, CoTelephone
	,
	VAA49 VARCHAR (64) --单位地址, CoAddress
	,
	VAA50 VARCHAR (6) --单位邮编, CoPostCode
	,
	VAA51 VARCHAR (64) --单位开户行
	,
	VAA52 VARCHAR (20) --单位银行帐号
	,
	VAA53 VARCHAR (20) --担保人, Guarantor
	,
	VAA54 NUMERIC (18, 2) --信用额度, 担保额度, CreditLimit
	,
	VAA55 TINYINT --担保性质, CreditType
	,
	VAA56 INT --住院次数, HospitalizationNumber
	,
	VAA57 DATETIME --就诊时间, LastVisitDate
	,
	BCK01A INT --就诊科室, LastVisitDeptID
	,
	VAA61 TINYINT --就诊状态, 0=无, 1=门诊, 2=住院, 3=出院, 4=转院, 5=死亡, 9=其他 VisitState
	,
	VAA62 VARCHAR (255) --过敏史
	,
	BDX02 VARCHAR (64) --了解途径, 病人了解医院的方式(如电视广告, 介绍, 户外广告等)
	,
	VAA64 VARCHAR (255) --备注
	,
	VBU01 INT --帐号ID
	,
	VAA66 VARCHAR (64) --病案号
	,
	VAA67 VARCHAR (64) --查询密码
	,
	IAK05 VARCHAR (50) --社会保障号
	,
	IAA01 INT --保险机构
	,
	BCK01B INT --科室ID
	,
	BCK01C INT --病区ID
	,
	BCQ04 VARCHAR (20) --床号
	,
	VAA73 DATETIME --入院日期
	,
	VAA74 DATETIME --出院时间
	,
	VAA75 DATETIME --建档时间
	,
	BEP05 NUMERIC (18, 4) --住院报警值
	,
	BEP06 NUMERIC (18, 4) --住院信用额度
	,
	ABL01A VARCHAR (2) --正定型
	,
	ABL01B VARCHAR (2) --反定型
	,
	VAA76 DATETIME --有效时间
	,
	ABL01 VARCHAR (2) --血型
	,
	BEP06B NUMERIC (18, 4) --门诊信用额度
	,
	VAA78 VARCHAR (1) --Rh血型
	,
	VAA82 VARCHAR (64) --健康卡号
	,
	VAA01A INT --相关ID
	,
	VAA84 VARCHAR (20) --体检登记号
)