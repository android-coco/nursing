CREATE TABLE BBY1 --收费项目目录
(
  BBY01 INT --ID
  ,
  BDN01 VARCHAR (2) --类型ID
  ,
  BCA01 INT --分类ID
  ,
  BBY04 VARCHAR (20) --编码
  ,
  BBY05 VARCHAR (128) --名称, 药品的此处名称为商品名称
  ,
  BBY06 VARCHAR (48) --规格, specification, STRENGTH
  ,
  BBE02 VARCHAR (64) --产地, 生产商
  ,
  BBY08 VARCHAR (20) --单位, 药品此处为基本单位
  ,
  BCF01 TINYINT --性质, 项目特性, 1=挂号, 2=急诊, 3=特别护理, 4=常规护理, 5=尾数处理, 6=诊金, 7=病历本, 8=就诊卡
  ,
  BBY10 VARCHAR (50) --标准编码, StandardCode
  ,
  BBY11 VARCHAR (50) --备用编码, ReserveCode
  ,
  BAX01 INT --收入项目
  ,
  AAS01 VARCHAR (4) --偿付类别, 费用类型
  ,
  ABF01 VARCHAR (8) --收据费别
  ,
  ABA01 VARCHAR (4) --病案费别
  ,
  BCG01 VARCHAR (4) --其他费别
  ,
  BCH01 VARCHAR (4) --折扣费别
  ,
  ACF01 INT --业务类别, 服务对象
  ,
  BBY18 TINYINT --忽略折扣, 屏蔽费别, 指是否打折, no discount
  ,
  BBY19 TINYINT --非药品：允许变价, 价格可变, 是否变价 0=否 1=允许变价, Price Variable   药品：对应药价属性，0：定价 1=时价 2＝ 指导价
  ,
  BBY20 TINYINT --加班加价, OvertimeMarkupType, 0=不加价, 1=比例加价, 2=定额加价, 3=指定价格, 暂时只用0, 1
  ,
  BBY21 TINYINT --补充说明, 在记帐时是否要补充说明
  ,
  BBY22 TINYINT --收费确认, 费用确认, charge confirm 0=否,1=是
  ,
  BBY23 NUMERIC (12, 4) --限制用量, LimitQuantity
  ,
  BCK01 INT --执行科室, execute dept 0=不指定, 1=病人科室, 2=指定科室, 3=病人病区, 4=操作科室, 5=院外执行, 6=开单科室
  ,
  BBY25 NUMERIC (18, 6) --单价, 对应当前售价
  ,
  BBY26 NUMERIC (18, 6) --首部位加价, First part markup
  ,
  BBY27 NUMERIC (18, 6) --最低价, MinPrice
  ,
  BBY28 NUMERIC (18, 6) --最高价, MaxPrice, AdministeredPrice
  ,
  BBY29 VARCHAR (512) --说明
  ,
  BBY30 DATETIME --创建时间
  ,
  BBY31 DATETIME --有效时间
  ,
  BBY32 TINYINT --计算方式 0= 人工；1=自动
  ,
  BBY34 TINYINT --组合方式(0：无 1：主从属 2：套餐 3：组合)
  ,
  BBY35 NUMERIC (18, 4) --药库单位价格
  ,
  BGF01 VARCHAR (8) --财务费别
  ,
  BHH01 VARCHAR (8) --业务费别
  ,
  BBY38 TINYINT --日志级别 ：0 ：无需记录日志 1：记录日志
  ,
  BBY39 TINYINT --分单标识 0 = 否;1=分单
  ,
  BBY40 TINYINT --忽略零库存
  ,
  BBY41 VARCHAR (128) --生产商
  ,
  BCE03A VARCHAR (20) --最后修改人
  ,
  BBY43 DATETIME --最后修改时间
  ,
  BBY44 VARCHAR (64) --名称缩写
)