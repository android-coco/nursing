CREATE TABLE BCK1 --部门表
(
  BCK01 INT --ID, DEPTU=业务科室, DEPEU=辅科室(如执行科室, 执行药房等), DEPWU=病区
  ,
  BCK02 VARCHAR (10) --编码, Code
  ,
  BCK03 VARCHAR (64) --名称, NameA
  ,
  BCK04 VARCHAR (128) --英文名称
  ,
  ABBRP VARCHAR (10) --拼音, PYImc
  ,
  ABBRW VARCHAR (10) --五笔, WBImc
  ,
  BCK01A INT --上级ID, ParentID
  ,
  LVLNR INT --隶属科室id：对应于科室下面分组的情况时，发票上打印科室名称时取此字段对应的科室
  ,
  BCK09 VARCHAR (20) --电话, Telephone
  ,
  BCK10 VARCHAR (128) --位置, Location, Site
  ,
  BCK11 VARCHAR (2) --业务性质 0： 无 1： 管理 2：医疗 3：药事
  ,
  ACA01 VARCHAR (8) --诊疗科目编码, 参见表TMCPCL
  ,
  BCK13 VARCHAR (255) --说明, Description
  ,
  ABY01 VARCHAR (2) --洁净等级, 参见表TMECLV(净洁等级表), 1=Ⅰ级, 2=Ⅱ级, 3=Ⅲ, 4=Ⅳ, 分为4级
  ,
  BCK15 DATETIME --创建时间, createDate
  ,
  BCK16 DATETIME --撤销时间, Expirydate
  ,
  ADR01 INT --分支机构id
  ,
  BCK18 TINYINT --用于药房药品价格管理 0：零价销售 1＝进价销售
  ,
  BCK19 NUMERIC (18, 4) --药物配额(%)
  ,
  BLL01 INT --区域ID
  ,
  ABW01 VARCHAR (1) --''0''=不限制  ''1''=男  ''2''=女
)