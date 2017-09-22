CREATE TABLE BCE1 --工表, 人员表
(
	BCE01 INT --ID
	,
	BCE02 VARCHAR (20) --编码, 工号, Code
	,
	BCE03 VARCHAR (20) --姓名, 在不同场合的名称： EMPNN=护士, EMPPN=医师, EMPDN=药剂师, EMPWN=仓管员, EMPTN=医技师,
	,
	BCE04 VARCHAR (20) --英文名, NameB
	,
	BCE05 VARCHAR (10) --称谓, Title
	,
	BCE06 VARCHAR (20) --曾用名, PreviousName
	,
	ABBRP VARCHAR (10) --拼音, PYImc
	,
	ABBRW VARCHAR (10) --五笔, WBImc
	,
	ABW01 VARCHAR (1) --性别, Sex
	,
	ACK01 VARCHAR (2) --婚姻, Marriage, Marriaged
	,
	BCE11 VARCHAR (20) --身份证号, IDNumber
	,
	ABQ01 VARCHAR (4) --民族, Nation
	,
	BCE13 DATETIME --出生日期, BirthDate
	,
	BCK01 INT --部门ID
	,
	ACP01 VARCHAR (2) --政治面貌, TPLAC.PLACC, PoliticalAffiliationID
	,
	AAY01 VARCHAR (4) --学历, 最高学历, 参见表, TEDUL
	,
	ACT01 VARCHAR (4) --学位, TDEGR
	,
	BCE18 VARCHAR (64) --毕业学校, graduated school
	,
	BCE19 DATETIME --毕业时间, graduated date
	,
	ABU02 VARCHAR (64) --ADIVN
	,
	AAD01 VARCHAR (4) --所学专业
	,
	ABG01 VARCHAR (4) --从事专业, TPPST
	,
	BCE23 VARCHAR (32) --执业证号
	,
	BCE24 VARCHAR (32) --户口所在地, RegisteredResidence
	,
	ABS01 VARCHAR (8) --行政职务
	,
	ABI01 VARCHAR (4) --技术职务, 职称
	,
	AAH01 VARCHAR (2) --聘任职务
	,
	BCE28 DATETIME --工作日期, joindate
	,
	AAQ01 VARCHAR (4) --执业类别
	,
	ABE01 VARCHAR (10) --执业范围, 需要单独的表
	,
	BCE31 DATETIME --进院日期, HireDate
	,
	BCE32 DATETIME --离职日期, DimissionDate
	,
	BCE33 VARCHAR (32) --离职说明, DimissionDescription
	,
	BCE34 VARCHAR (128) --住址, Adress
	,
	BCE35 VARCHAR (20) --办公电话, officephone
	,
	BCE36 VARCHAR (20) --联系电话, phone
	,
	BCE37 VARCHAR (20) --移动电话, mobilephone
	,
	BCE38 VARCHAR (64) --电子邮箱, email
	,
	BCE39 VARCHAR (255) --备注, comment
)