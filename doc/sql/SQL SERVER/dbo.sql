/*
Navicat SQL Server Data Transfer

Source Server         : SQLserver
Source Server Version : 110000
Source Host           : 192.168.0.129:1433
Source Database       : his
Source Schema         : dbo

Target Server Type    : SQL Server
Target Server Version : 110000
File Encoding         : 65001

Date: 2017-10-30 10:15:37
*/


-- ----------------------------
-- Table structure for AAG1
-- ----------------------------
DROP TABLE [dbo].[AAG1]
GO
CREATE TABLE [dbo].[AAG1] (
[AAG01] int NULL ,
[AAG02] varchar(32) NULL ,
[AAG03] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'AAG1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'护理级别'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'护理级别'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'AAG1', 
'COLUMN', N'AAG01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'编码，实际只用1位'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
, @level2type = 'COLUMN', @level2name = N'AAG01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'编码，实际只用1位'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
, @level2type = 'COLUMN', @level2name = N'AAG01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'AAG1', 
'COLUMN', N'AAG02')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
, @level2type = 'COLUMN', @level2name = N'AAG02'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
, @level2type = 'COLUMN', @level2name = N'AAG02'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'AAG1', 
'COLUMN', N'AAG03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'颜色，0=无， > 0指具体颜色'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
, @level2type = 'COLUMN', @level2name = N'AAG03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'颜色，0=无， > 0指具体颜色'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'AAG1'
, @level2type = 'COLUMN', @level2name = N'AAG03'
GO

-- ----------------------------
-- Table structure for BAK1
-- ----------------------------
DROP TABLE [dbo].[BAK1]
GO
CREATE TABLE [dbo].[BAK1] (
[BAK01] int NULL ,
[BAK02] varchar(32) NULL ,
[BAK03] tinyint NULL ,
[BAK04] varchar(32) NULL ,
[BAK05] varchar(128) NULL ,
[BAK06] varchar(128) NULL ,
[BAK07] varchar(255) NULL ,
[ABBRP] varchar(32) NULL ,
[ABBRW] varchar(32) NULL ,
[ABW01] varchar(1) NULL ,
[ABX01] varchar(1) NULL ,
[BAK12] tinyint NULL ,
[ACH01] varchar(1) NULL ,
[BAK14] varchar(10) NULL ,
[BAF01] varchar(1) NULL ,
[BAH01] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病目录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病目录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK02')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'编码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK02'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'编码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK02'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'优先级'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'优先级'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK03'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK04')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'附码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK04'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'附码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK04'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK05')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK05'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK05'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK06')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'英文名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK06'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'英文名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK06'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'说明'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'说明'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'ABBRP')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'拼音'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABBRP'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'拼音'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABBRP'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'ABBRW')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'中文五笔缩写'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABBRW'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'中文五笔缩写'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABBRW'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'ABW01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'性别限制0=不限 1=男  2=女'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABW01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'性别限制0=不限 1=男  2=女'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABW01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'ABX01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疗效限制, TDOTC, CurativeEffect'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABX01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疗效限制, TDOTC, CurativeEffect'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ABX01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK12')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'分娩, 0=否, 1=是, Labour'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK12'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'分娩, 0=否, 1=是, Labour'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK12'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'ACH01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'手术类型'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ACH01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'手术类型'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'ACH01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAK14')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'统计码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK14'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'统计码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAK14'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAF01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病类型'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAF01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病类型'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAF01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAK1', 
'COLUMN', N'BAH01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病分类ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAH01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病分类ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAK1'
, @level2type = 'COLUMN', @level2name = N'BAH01'
GO

-- ----------------------------
-- Table structure for BAO2
-- ----------------------------
DROP TABLE [dbo].[BAO2]
GO
CREATE TABLE [dbo].[BAO2] (
[VAA01] int NULL ,
[BAK01A] int NULL ,
[VAO19] datetime NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BAO2', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'住院病人诊断记录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAO2'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'住院病人诊断记录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BAO2'
GO

-- ----------------------------
-- Table structure for BBY1
-- ----------------------------
DROP TABLE [dbo].[BBY1]
GO
CREATE TABLE [dbo].[BBY1] (
[BBY01] int NULL ,
[BDN01] varchar(2) NULL ,
[BCA01] int NULL ,
[BBY04] varchar(20) NULL ,
[BBY05] varchar(128) NULL ,
[BBY06] varchar(48) NULL ,
[BBE02] varchar(64) NULL ,
[BBY08] varchar(20) NULL ,
[BCF01] tinyint NULL ,
[BBY10] varchar(50) NULL ,
[BBY11] varchar(50) NULL ,
[BAX01] int NULL ,
[AAS01] varchar(4) NULL ,
[ABF01] varchar(8) NULL ,
[ABA01] varchar(4) NULL ,
[BCG01] varchar(4) NULL ,
[BCH01] varchar(4) NULL ,
[ACF01] int NULL ,
[BBY18] tinyint NULL ,
[BBY19] tinyint NULL ,
[BBY20] tinyint NULL ,
[BBY21] tinyint NULL ,
[BBY22] tinyint NULL ,
[BBY23] numeric(12,4) NULL ,
[BCK01] int NULL ,
[BBY25] numeric(18,6) NULL ,
[BBY26] numeric(18,6) NULL ,
[BBY27] numeric(18,6) NULL ,
[BBY28] numeric(18,6) NULL ,
[BBY29] varchar(512) NULL ,
[BBY30] datetime NULL ,
[BBY31] datetime NULL ,
[BBY32] tinyint NULL ,
[BBY34] tinyint NULL ,
[BBY35] numeric(18,4) NULL ,
[BGF01] varchar(8) NULL ,
[BHH01] varchar(8) NULL ,
[BBY38] tinyint NULL ,
[BBY39] tinyint NULL ,
[BBY40] tinyint NULL ,
[BBY41] varchar(128) NULL ,
[BCE03A] varchar(20) NULL ,
[BBY43] datetime NULL ,
[BBY44] varchar(64) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BBY1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BBY1--收费项目目录
(
			 BBY01  INT  --ID
      ,BDN01  VARCHAR(2)  --类型ID
      ,BCA01  INT  --分类ID
      ,BBY04  VARCHAR(20)  --编码
      ,BBY05  VARCHAR(128)  --名称, 药品的此处名称为商品名称
      ,BBY06  VARCHAR(48)  --规格, specification, STRENGTH
      ,BBE02  VARCHAR(64)  --产地, 生产商
      ,BBY08  VARCHAR(20)  --单位, 药品此处为基本单位
      ,BCF01  TINYINT  --性质, 项目特性, 1=挂号, 2=急诊, 3=特别护理, 4=常规护理, 5=尾数处理, 6=诊金, 7=病历本, 8=就诊卡
      ,BBY10  VARCHAR(50)  --标准编码, StandardCode
      ,BBY11  VARCHAR(50)  --备用编码, ReserveCode
      ,BAX01  INT  --收入项目
      ,AAS01  VARCHAR(4)  --偿付类别, 费用类型
      ,ABF01  VARCHAR(8)  --收据费别
      ,ABA01  VARCHAR(4)  --病案费别
      ,BCG01  VARCHAR(4)  --其他费别
      ,BCH01  VARCHAR(4)  --折扣费别
      ,ACF01  INT  --业务类别, 服务对象
      ,BBY18  TINYINT  --忽略折扣, 屏蔽费别, 指是否打折, no discount
      ,BBY19  TINYINT  --非药品：允许变价, 价格可变, 是否变价 0=否 1=允许变价, Price Variable   药品：对应药价属性，0：定价 1=时价 2＝ 指导价 
      ,BBY20  TINYINT  --加班加价, OvertimeMarkupType, 0=不加价, 1=比例加价, 2=定额加价, 3=指定价格, 暂时只用0, 1
      ,BBY21  TINYINT  --补充说明, 在记帐时是否要补充说明
      ,BBY22  TINYINT  --收费确认, 费用确认, charge confirm 0=否,1=是
      ,BBY23  NUMERIC(12, 4)  --限制用量, LimitQuantity
      ,BCK01  INT  --执行科室, execute dept 0=不指定, 1=病人科室, 2=指定科室, 3=病人病区, 4=操作科室, 5=院外执行, 6=开单科室
      ,BBY25  NUMERIC(18, 6)  --单价, 对应当前售价
      ,BBY26  NUMERIC(18, 6)  --首部位加价, First part markup
      ,BBY27  NUMERIC(18, 6)  --最低价, MinPrice
      ,BBY28  NUMERIC(18, 6)  --最高价, MaxPrice, AdministeredPrice
      ,BBY29  VARCHAR(512)  --说明
      ,BBY30  DATETIME  --创建时间
      ,BBY31  DATETIME  --有效时间
      ,BBY32  TINYINT  --计算方式 0= 人工；1=自动
      ,BBY34  TINYINT  --组合方式(0：无 1：主从属 2：套餐 3：组合)
      ,BBY35  NUMERIC(18, 4)  --药库单位价格
      ,BGF01  VARCHAR(8)  --财务费别
      ,BHH01  VARCHAR(8)  --业务费别
      ,BBY38  TINYINT  --日志级别 ：0 ：无需记录日志 1：记录日志
      ,BBY39  TINYINT  --分单标识 0 = 否;1=分单
      ,BBY40  TINYINT  --忽略零库存
      ,BBY41  VARCHAR(128)  --生产商
      ,BCE03A VARCHAR(20)  --最后修改人
      ,BBY43  DATETIME  --最后修改时间
      ,BBY44  VARCHAR(64)  --名称缩写

)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BBY1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BBY1--收费项目目录
(
			 BBY01  INT  --ID
      ,BDN01  VARCHAR(2)  --类型ID
      ,BCA01  INT  --分类ID
      ,BBY04  VARCHAR(20)  --编码
      ,BBY05  VARCHAR(128)  --名称, 药品的此处名称为商品名称
      ,BBY06  VARCHAR(48)  --规格, specification, STRENGTH
      ,BBE02  VARCHAR(64)  --产地, 生产商
      ,BBY08  VARCHAR(20)  --单位, 药品此处为基本单位
      ,BCF01  TINYINT  --性质, 项目特性, 1=挂号, 2=急诊, 3=特别护理, 4=常规护理, 5=尾数处理, 6=诊金, 7=病历本, 8=就诊卡
      ,BBY10  VARCHAR(50)  --标准编码, StandardCode
      ,BBY11  VARCHAR(50)  --备用编码, ReserveCode
      ,BAX01  INT  --收入项目
      ,AAS01  VARCHAR(4)  --偿付类别, 费用类型
      ,ABF01  VARCHAR(8)  --收据费别
      ,ABA01  VARCHAR(4)  --病案费别
      ,BCG01  VARCHAR(4)  --其他费别
      ,BCH01  VARCHAR(4)  --折扣费别
      ,ACF01  INT  --业务类别, 服务对象
      ,BBY18  TINYINT  --忽略折扣, 屏蔽费别, 指是否打折, no discount
      ,BBY19  TINYINT  --非药品：允许变价, 价格可变, 是否变价 0=否 1=允许变价, Price Variable   药品：对应药价属性，0：定价 1=时价 2＝ 指导价 
      ,BBY20  TINYINT  --加班加价, OvertimeMarkupType, 0=不加价, 1=比例加价, 2=定额加价, 3=指定价格, 暂时只用0, 1
      ,BBY21  TINYINT  --补充说明, 在记帐时是否要补充说明
      ,BBY22  TINYINT  --收费确认, 费用确认, charge confirm 0=否,1=是
      ,BBY23  NUMERIC(12, 4)  --限制用量, LimitQuantity
      ,BCK01  INT  --执行科室, execute dept 0=不指定, 1=病人科室, 2=指定科室, 3=病人病区, 4=操作科室, 5=院外执行, 6=开单科室
      ,BBY25  NUMERIC(18, 6)  --单价, 对应当前售价
      ,BBY26  NUMERIC(18, 6)  --首部位加价, First part markup
      ,BBY27  NUMERIC(18, 6)  --最低价, MinPrice
      ,BBY28  NUMERIC(18, 6)  --最高价, MaxPrice, AdministeredPrice
      ,BBY29  VARCHAR(512)  --说明
      ,BBY30  DATETIME  --创建时间
      ,BBY31  DATETIME  --有效时间
      ,BBY32  TINYINT  --计算方式 0= 人工；1=自动
      ,BBY34  TINYINT  --组合方式(0：无 1：主从属 2：套餐 3：组合)
      ,BBY35  NUMERIC(18, 4)  --药库单位价格
      ,BGF01  VARCHAR(8)  --财务费别
      ,BHH01  VARCHAR(8)  --业务费别
      ,BBY38  TINYINT  --日志级别 ：0 ：无需记录日志 1：记录日志
      ,BBY39  TINYINT  --分单标识 0 = 否;1=分单
      ,BBY40  TINYINT  --忽略零库存
      ,BBY41  VARCHAR(128)  --生产商
      ,BCE03A VARCHAR(20)  --最后修改人
      ,BBY43  DATETIME  --最后修改时间
      ,BBY44  VARCHAR(64)  --名称缩写

)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BBY1'
GO

-- ----------------------------
-- Table structure for BCE1
-- ----------------------------
DROP TABLE [dbo].[BCE1]
GO
CREATE TABLE [dbo].[BCE1] (
[BCE01] int NULL ,
[BCE02] varchar(20) NULL ,
[BCE03] varchar(20) NULL ,
[BCE04] varchar(20) NULL ,
[BCE05] varchar(10) NULL ,
[BCE06] varchar(20) NULL ,
[ABBRP] varchar(10) NULL ,
[ABBRW] varchar(10) NULL ,
[ABW01] varchar(1) NULL ,
[ACK01] varchar(2) NULL ,
[BCE11] varchar(20) NULL ,
[ABQ01] varchar(4) NULL ,
[BCE13] datetime NULL ,
[BCK01] int NULL ,
[ACP01] varchar(2) NULL ,
[AAY01] varchar(4) NULL ,
[ACT01] varchar(4) NULL ,
[BCE18] varchar(64) NULL ,
[BCE19] datetime NULL ,
[ABU02] varchar(64) NULL ,
[AAD01] varchar(4) NULL ,
[ABG01] varchar(4) NULL ,
[BCE23] varchar(32) NULL ,
[BCE24] varchar(32) NULL ,
[ABS01] varchar(8) NULL ,
[ABI01] varchar(4) NULL ,
[AAH01] varchar(2) NULL ,
[BCE28] datetime NULL ,
[AAQ01] varchar(4) NULL ,
[ABE01] varchar(10) NULL ,
[BCE31] datetime NULL ,
[BCE32] datetime NULL ,
[BCE33] varchar(32) NULL ,
[BCE34] varchar(128) NULL ,
[BCE35] varchar(20) NULL ,
[BCE36] varchar(20) NULL ,
[BCE37] varchar(20) NULL ,
[BCE38] varchar(64) NULL ,
[BCE39] varchar(255) NULL ,
[BCE41] smallint NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BCE1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BCE1--工表, 人员表
(
       BCE01 INT  --ID
      ,BCE02 VARCHAR(20)  --编码, 工号, Code
      ,BCE03 VARCHAR(20)  --姓名, 在不同场合的名称： EMPNN=护士, EMPPN=医师, EMPDN=药剂师, EMPWN=仓管员, EMPTN=医技师,
      ,BCE04 VARCHAR(20)  --英文名, NameB
      ,BCE05 VARCHAR(10)  --称谓, Title
      ,BCE06 VARCHAR(20)  --曾用名, PreviousName
      ,ABBRP VARCHAR(10)  --拼音, PYImc
      ,ABBRW VARCHAR(10)  --五笔, WBImc
      ,ABW01 VARCHAR(1)  --性别, Sex
      ,ACK01 VARCHAR(2)  --婚姻, Marriage, Marriaged
      ,BCE11 VARCHAR(20)  --身份证号, IDNumber
      ,ABQ01 VARCHAR(4)  --民族, Nation
      ,BCE13 DATETIME  --出生日期, BirthDate
      ,BCK01 INT  --部门ID
      ,ACP01 VARCHAR(2)  --政治面貌, TPLAC.PLACC, PoliticalAffiliationID
      ,AAY01 VARCHAR(4)  --学历, 最高学历, 参见表, TEDUL
      ,ACT01 VARCHAR(4)  --学位, TDEGR
      ,BCE18 VARCHAR(64)  --毕业学校, graduated school
      ,BCE19 DATETIME  --毕业时间, graduated date
      ,ABU02 VARCHAR(64)  --ADIVN
      ,AAD01 VARCHAR(4)  --所学专业
      ,ABG01 VARCHAR(4)  --从事专业, TPPST
      ,BCE23 VARCHAR(32)  --执业证号
      ,BCE24 VARCHAR(32)  --户口所在地, RegisteredResidence
      ,ABS01 VARCHAR(8)  --行政职务
      ,ABI01 VARCHAR(4)  --技术职务, 职称
      ,AAH01 VARCHAR(2)  --聘任职务
      ,BCE28 DATETIME  --工作日期, joindate
      ,AAQ01 VARCHAR(4)  --执业类别
      ,ABE01 VARCHAR(10)  --执业范围, 需要单独的表
      ,BCE31 DATETIME  --进院日期, HireDate
      ,BCE32 DATETIME  --离职日期, DimissionDate
      ,BCE33 VARCHAR(32)  --离职说明, DimissionDescription
      ,BCE34 VARCHAR(128)  --住址, Adress
      ,BCE35 VARCHAR(20)  --办公电话, officephone
      ,BCE36 VARCHAR(20)  --联系电话, phone
      ,BCE37 VARCHAR(20)  --移动电话, mobilephone
      ,BCE38 VARCHAR(64)  --电子邮箱, email
      ,BCE39 VARCHAR(255)  --备注, comment
      ,BCE41 SMALLINT(2)  --状态，0=试用，1=在职，2=离职，3=退休
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCE1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BCE1--工表, 人员表
(
       BCE01 INT  --ID
      ,BCE02 VARCHAR(20)  --编码, 工号, Code
      ,BCE03 VARCHAR(20)  --姓名, 在不同场合的名称： EMPNN=护士, EMPPN=医师, EMPDN=药剂师, EMPWN=仓管员, EMPTN=医技师,
      ,BCE04 VARCHAR(20)  --英文名, NameB
      ,BCE05 VARCHAR(10)  --称谓, Title
      ,BCE06 VARCHAR(20)  --曾用名, PreviousName
      ,ABBRP VARCHAR(10)  --拼音, PYImc
      ,ABBRW VARCHAR(10)  --五笔, WBImc
      ,ABW01 VARCHAR(1)  --性别, Sex
      ,ACK01 VARCHAR(2)  --婚姻, Marriage, Marriaged
      ,BCE11 VARCHAR(20)  --身份证号, IDNumber
      ,ABQ01 VARCHAR(4)  --民族, Nation
      ,BCE13 DATETIME  --出生日期, BirthDate
      ,BCK01 INT  --部门ID
      ,ACP01 VARCHAR(2)  --政治面貌, TPLAC.PLACC, PoliticalAffiliationID
      ,AAY01 VARCHAR(4)  --学历, 最高学历, 参见表, TEDUL
      ,ACT01 VARCHAR(4)  --学位, TDEGR
      ,BCE18 VARCHAR(64)  --毕业学校, graduated school
      ,BCE19 DATETIME  --毕业时间, graduated date
      ,ABU02 VARCHAR(64)  --ADIVN
      ,AAD01 VARCHAR(4)  --所学专业
      ,ABG01 VARCHAR(4)  --从事专业, TPPST
      ,BCE23 VARCHAR(32)  --执业证号
      ,BCE24 VARCHAR(32)  --户口所在地, RegisteredResidence
      ,ABS01 VARCHAR(8)  --行政职务
      ,ABI01 VARCHAR(4)  --技术职务, 职称
      ,AAH01 VARCHAR(2)  --聘任职务
      ,BCE28 DATETIME  --工作日期, joindate
      ,AAQ01 VARCHAR(4)  --执业类别
      ,ABE01 VARCHAR(10)  --执业范围, 需要单独的表
      ,BCE31 DATETIME  --进院日期, HireDate
      ,BCE32 DATETIME  --离职日期, DimissionDate
      ,BCE33 VARCHAR(32)  --离职说明, DimissionDescription
      ,BCE34 VARCHAR(128)  --住址, Adress
      ,BCE35 VARCHAR(20)  --办公电话, officephone
      ,BCE36 VARCHAR(20)  --联系电话, phone
      ,BCE37 VARCHAR(20)  --移动电话, mobilephone
      ,BCE38 VARCHAR(64)  --电子邮箱, email
      ,BCE39 VARCHAR(255)  --备注, comment
      ,BCE41 SMALLINT(2)  --状态，0=试用，1=在职，2=离职，3=退休
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCE1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BCE1', 
'COLUMN', N'BCE02')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'编码, 工号, Code'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCE1'
, @level2type = 'COLUMN', @level2name = N'BCE02'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'编码, 工号, Code'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCE1'
, @level2type = 'COLUMN', @level2name = N'BCE02'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BCE1', 
'COLUMN', N'BCE41')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'状态，0=试用，1=在职，2=离职，3=退休'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCE1'
, @level2type = 'COLUMN', @level2name = N'BCE41'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'状态，0=试用，1=在职，2=离职，3=退休'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCE1'
, @level2type = 'COLUMN', @level2name = N'BCE41'
GO

-- ----------------------------
-- Table structure for BCK1
-- ----------------------------
DROP TABLE [dbo].[BCK1]
GO
CREATE TABLE [dbo].[BCK1] (
[BCK01] int NULL ,
[BCK02] varchar(10) NULL ,
[BCK03] varchar(64) NULL ,
[BCK04] varchar(128) NULL ,
[ABBRP] varchar(10) NULL ,
[ABBRW] varchar(10) NULL ,
[BCK01A] int NULL ,
[LVLNR] int NULL ,
[BCK09] varchar(20) NULL ,
[BCK10] varchar(128) NULL ,
[BCK11] varchar(2) NULL ,
[ACA01] varchar(8) NULL ,
[BCK13] varchar(255) NULL ,
[ABY01] varchar(2) NULL ,
[BCK15] datetime NULL ,
[BCK16] datetime NULL ,
[ADR01] int NULL ,
[BCK18] tinyint NULL ,
[BCK19] numeric(18,4) NULL ,
[BLL01] int NULL ,
[ABW01] varchar(1) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BCK1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BCK1(部门表)
(
       BCK01  INT  --ID, DEPTU=业务科室, DEPEU=辅科室(如执行科室, 执行药房等), DEPWU=病区
      ,BCK02  VARCHAR(10)  --编码, Code
      ,BCK03  VARCHAR(64)  --名称, NameA
      ,BCK04  VARCHAR(128)  --英文名称
      ,ABBRP  VARCHAR(10)  --拼音, PYImc
      ,ABBRW  VARCHAR(10)  --五笔, WBImc
      ,BCK01A INT  --上级ID, ParentID
      ,LVLNR  INT  --隶属科室id：对应于科室下面分组的情况时，发票上打印科室名称时取此字段对应的科室
      ,BCK09  VARCHAR(20)  --电话, Telephone
      ,BCK10  VARCHAR(128)  --位置, Location, Site
      ,BCK11  VARCHAR(2)  --业务性质 0： 无 1： 管理 2：医疗 3：药事
      ,ACA01  VARCHAR(8)  --诊疗科目编码, 参见表TMCPCL
      ,BCK13  VARCHAR(255)  --说明, Description
      ,ABY01  VARCHAR(2)  --洁净等级, 参见表TMECLV(净洁等级表), 1=Ⅰ级, 2=Ⅱ级, 3=Ⅲ, 4=Ⅳ, 分为4级
      ,BCK15  DATETIME  --创建时间, createDate
      ,BCK16  DATETIME  --撤销时间, Expirydate
      ,ADR01  INT  --分支机构id
      ,BCK18  TINYINT  --用于药房药品价格管理 0：零价销售 1＝进价销售
      ,BCK19  NUMERIC(18, 4)  --药物配额(%)
      ,BLL01  INT  --区域ID
      ,ABW01  VARCHAR(1)  --''0''=不限制  ''1''=男  ''2''=女

)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCK1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BCK1(部门表)
(
       BCK01  INT  --ID, DEPTU=业务科室, DEPEU=辅科室(如执行科室, 执行药房等), DEPWU=病区
      ,BCK02  VARCHAR(10)  --编码, Code
      ,BCK03  VARCHAR(64)  --名称, NameA
      ,BCK04  VARCHAR(128)  --英文名称
      ,ABBRP  VARCHAR(10)  --拼音, PYImc
      ,ABBRW  VARCHAR(10)  --五笔, WBImc
      ,BCK01A INT  --上级ID, ParentID
      ,LVLNR  INT  --隶属科室id：对应于科室下面分组的情况时，发票上打印科室名称时取此字段对应的科室
      ,BCK09  VARCHAR(20)  --电话, Telephone
      ,BCK10  VARCHAR(128)  --位置, Location, Site
      ,BCK11  VARCHAR(2)  --业务性质 0： 无 1： 管理 2：医疗 3：药事
      ,ACA01  VARCHAR(8)  --诊疗科目编码, 参见表TMCPCL
      ,BCK13  VARCHAR(255)  --说明, Description
      ,ABY01  VARCHAR(2)  --洁净等级, 参见表TMECLV(净洁等级表), 1=Ⅰ级, 2=Ⅱ级, 3=Ⅲ, 4=Ⅳ, 分为4级
      ,BCK15  DATETIME  --创建时间, createDate
      ,BCK16  DATETIME  --撤销时间, Expirydate
      ,ADR01  INT  --分支机构id
      ,BCK18  TINYINT  --用于药房药品价格管理 0：零价销售 1＝进价销售
      ,BCK19  NUMERIC(18, 4)  --药物配额(%)
      ,BLL01  INT  --区域ID
      ,ABW01  VARCHAR(1)  --''0''=不限制  ''1''=男  ''2''=女

)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCK1'
GO

-- ----------------------------
-- Table structure for BCQ1
-- ----------------------------
DROP TABLE [dbo].[BCQ1]
GO
CREATE TABLE [dbo].[BCQ1] (
[BCQ01] int NULL ,
[BCK01A] int NULL ,
[BCQ03] varchar(10) NULL ,
[BCQ04] varchar(20) NULL ,
[BCK01B] int NULL ,
[ABW01] varchar(1) NULL ,
[ACG02] varchar(20) NULL ,
[BBY01] int NULL ,
[BCQ09] tinyint NULL ,
[ROWNR] int NULL ,
[BCQ11] varchar(128) NULL ,
[VAA01] int NULL ,
[BCQ13] tinyint NULL ,
[BBY01B] int NULL ,
[BCQ15] tinyint NULL ,
[BBY01C] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'BCQ1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BCQ1--病床编制表
(
        BCQ01  INT  --ID
      ,BCK01A INT  --病区ID
      ,BCQ03  VARCHAR(10)  --病房, 病室, 房间
      ,BCQ04  VARCHAR(20)  --床位, 床号, 病床
      ,BCK01B INT  --科室ID
      ,ABW01  VARCHAR(1)  --性别, 性别限制
      ,ACG02  VARCHAR(20)  --编制性质
      ,BBY01  INT  --收费项目ID, 等级
      ,BCQ09  TINYINT  --共用 1：共用
      ,ROWNR  INT  --次序
      ,BCQ11  VARCHAR(128)  --说明
      ,VAA01  INT  --病人ID, 这个不用显示维护
      ,BCQ13  TINYINT  --状态, 0=空床, 1=占用, 2=陪床(已占用),3=包房, 4=维修,5=撤编 CWBST
      ,BBY01B INT  --包床项目
      ,BCQ15  TINYINT  --0=普通, 1=CCU, 2=挂床
      ,BBY01C INT  --标准收费项目id
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCQ1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE BCQ1--病床编制表
(
        BCQ01  INT  --ID
      ,BCK01A INT  --病区ID
      ,BCQ03  VARCHAR(10)  --病房, 病室, 房间
      ,BCQ04  VARCHAR(20)  --床位, 床号, 病床
      ,BCK01B INT  --科室ID
      ,ABW01  VARCHAR(1)  --性别, 性别限制
      ,ACG02  VARCHAR(20)  --编制性质
      ,BBY01  INT  --收费项目ID, 等级
      ,BCQ09  TINYINT  --共用 1：共用
      ,ROWNR  INT  --次序
      ,BCQ11  VARCHAR(128)  --说明
      ,VAA01  INT  --病人ID, 这个不用显示维护
      ,BCQ13  TINYINT  --状态, 0=空床, 1=占用, 2=陪床(已占用),3=包房, 4=维修,5=撤编 CWBST
      ,BBY01B INT  --包床项目
      ,BCQ15  TINYINT  --0=普通, 1=CCU, 2=挂床
      ,BBY01C INT  --标准收费项目id
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'BCQ1'
GO

-- ----------------------------
-- Table structure for CAM1
-- ----------------------------
DROP TABLE [dbo].[CAM1]
GO
CREATE TABLE [dbo].[CAM1] (
[CAM01] int NULL ,
[CAM02] tinyint NULL ,
[CAM03] varchar(10) NULL ,
[CAM04] varchar(64) NULL ,
[CAM05] varchar(1024) NULL ,
[CAM06] varchar(64) NULL ,
[CAM07] int NULL ,
[CAM08] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病参考目录（已弃用：JP 2017.10.18）'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病参考目录（已弃用：JP 2017.10.18）'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM02')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'类型：1=中医， 2=西医'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM02'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'类型：1=中医， 2=西医'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM02'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'编码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'编码'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM03'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM04')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM04'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM04'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM05')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'说明'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM05'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'说明'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM05'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM06')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'编者'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM06'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'编者'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM06'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疑似--Suspected, similarity measure(相似度)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疑似--Suspected, similarity measure(相似度)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CAM1', 
'COLUMN', N'CAM08')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'临床'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM08'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'临床'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CAM1'
, @level2type = 'COLUMN', @level2name = N'CAM08'
GO

-- ----------------------------
-- Table structure for CBM1
-- ----------------------------
DROP TABLE [dbo].[CBM1]
GO
CREATE TABLE [dbo].[CBM1] (
[CBM01] int NULL ,
[VAA01] int NULL ,
[VAA07] int NULL ,
[VAP01] int NULL ,
[ACF01] int NULL ,
[CBM06] int NULL ,
[CBM07] int NULL ,
[CBM08] nvarchar(20) NULL ,
[CBM09] nvarchar(1024) NULL ,
[BCK01A] int NULL ,
[BCK01B] int NULL ,
[BCK01C] int NULL ,
[BCE01] int NULL ,
[BCE03] nvarchar(20) NULL ,
[CBM15] datetime NULL ,
[CBM16] datetime NULL ,
[CBM17] datetime NULL ,
[CBM18] nvarchar(255) NULL ,
[CBM19] tinyint NULL ,
[PRNCP] int NULL ,
[CBM21] int NULL ,
[BJW02] nvarchar(22) NULL ,
[VCY01] int NULL ,
[VBQ01] int NULL ,
[IAI03] nvarchar(10) NULL ,
[CBM27] varchar(1) NULL ,
[CBM28] varchar(1) NULL ,
[CBM29] varchar(1) NULL ,
[CBM30] varchar(1) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'住院病人医嘱单'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'住院病人医嘱单'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'VAA01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  病人ID, 关联字段:VAA1.VAA01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VAA01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  病人ID, 关联字段:VAA1.VAA01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VAA01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'VAA07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  就诊ID, 关联字段:VAA1.VAA07'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VAA07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  就诊ID, 关联字段:VAA1.VAA07'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VAA07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'VAP01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  婴儿ID, 关联字段:VAP1.VAP01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VAP01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  婴儿ID, 关联字段:VAP1.VAP01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VAP01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'ACF01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'1:门诊、2:住院, 关联字段:ACF1.ACF01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'ACF01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'1:门诊、2:住院, 关联字段:ACF1.ACF01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'ACF01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM06')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'性质: 0:不区分中、西药处方  1=西、成药, 2=中药, 3=手术记录, 4 = 成药,5=卫材, 9=其他'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM06'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'性质: 0:不区分中、西药处方  1=西、成药, 2=中药, 3=手术记录, 4 = 成药,5=卫材, 9=其他'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM06'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  1=普通, 2=急诊, 3=儿童, 4=麻醉, 5=第一类精神药品, 6=第二类精神药品, 7=放射药品,8=毒性药品,9=检查,10=检验,11=手术,12=治疗,99=其它'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  1=普通, 2=急诊, 3=儿童, 4=麻醉, 5=第一类精神药品, 6=第二类精神药品, 7=放射药品,8=毒性药品,9=检查,10=检验,11=手术,12=治疗,99=其它'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM08')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  单据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM08'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  单据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM08'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM09')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  摘要'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM09'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  摘要'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM09'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'BCK01A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  病区ID, 关联字段:BCK1.BCK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCK01A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  病区ID, 关联字段:BCK1.BCK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCK01A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'BCK01B')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  病人科室ID, 关联字段:BCK1.BCK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCK01B'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  病人科室ID, 关联字段:BCK1.BCK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCK01B'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'BCK01C')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  开单科室ID, 关联字段:BCK1.BCK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCK01C'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  开单科室ID, 关联字段:BCK1.BCK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCK01C'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'BCE01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  开单人ID, 关联字段:BCE1.BCE01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCE01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  开单人ID, 关联字段:BCE1.BCE01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCE01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'BCE03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  操作员, 关联字段:BCE1.BCE03'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCE03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  操作员, 关联字段:BCE1.BCE03'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BCE03'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM15')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  开单时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM15'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  开单时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM15'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM16')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  交易时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM16'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  交易时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM16'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM17')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  撤销时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM17'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  撤销时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM17'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM18')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  备注'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM18'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  备注'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM18'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM19')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  状态'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM19'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  状态'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM19'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'PRNCP')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  收费小票打印次数'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'PRNCP'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  收费小票打印次数'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'PRNCP'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'CBM21')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  处方打印次数'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM21'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  处方打印次数'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'CBM21'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'BJW02')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  处方分类, 关联字段:BJW1.BJW02'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BJW02'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  处方分类, 关联字段:BJW1.BJW02'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'BJW02'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'VCY01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  复诊预约id, 关联字段:VCY1.VCY01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VCY01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  复诊预约id, 关联字段:VCY1.VCY01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VCY01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'VBQ01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  诊疗申请单id, 关联字段:VBQ1.VBQ01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VBQ01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  诊疗申请单id, 关联字段:VBQ1.VBQ01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'VBQ01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBM1', 
'COLUMN', N'IAI03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'描述:  医疗类别编码, 关联字段:IAI1.IAI03'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'IAI03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'描述:  医疗类别编码, 关联字段:IAI1.IAI03'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBM1'
, @level2type = 'COLUMN', @level2name = N'IAI03'
GO

-- ----------------------------
-- Table structure for CBP1
-- ----------------------------
DROP TABLE [dbo].[CBP1]
GO
CREATE TABLE [dbo].[CBP1] (
[VAA01] int NULL ,
[BCK01] int NULL ,
[AAG01] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'CBP1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病人护理记录新版'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBP1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病人护理记录新版'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'CBP1'
GO

-- ----------------------------
-- Table structure for FHB1
-- ----------------------------
DROP TABLE [dbo].[FHB1]
GO
CREATE TABLE [dbo].[FHB1] (
[FHB01] varchar(36) NULL ,
[VAA01] int NULL ,
[ACF01] int NULL ,
[VAA07] int NULL ,
[VAK01] int NULL ,
[VAI01] int NULL ,
[VAJ01] int NULL ,
[BBX01] int NULL ,
[BBY01] int NULL ,
[BCE03] varchar(20) NULL ,
[FHB11] datetime NULL ,
[FHB12] tinyint NULL ,
[FHB13] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'FHB1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE FHB1(病人信息扩展表)
(
        FHB01 VARCHAR(36)  --ID
      ,VAA01 INT  --病人ID
      ,ACF01 INT  --业务类型：0=挂号；1=门诊；2=住院
      ,VAA07 INT  --就诊ID
      ,VAK01 INT  --结帐ID
      ,VAI01 INT  --单据ID
      ,VAJ01 INT  --明细ID
      ,BBX01 INT  --诊疗ID
      ,BBY01 INT  --收费项目ID
      ,BCE03 VARCHAR(20)  --操作员
      ,FHB11 DATETIME  --操作时间
      ,FHB12 TINYINT  --性质：1=需要隐藏的项目
      ,FHB13 INT  --状态：奇数=隐藏；偶数=还原
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'FHB1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE FHB1(病人信息扩展表)
(
        FHB01 VARCHAR(36)  --ID
      ,VAA01 INT  --病人ID
      ,ACF01 INT  --业务类型：0=挂号；1=门诊；2=住院
      ,VAA07 INT  --就诊ID
      ,VAK01 INT  --结帐ID
      ,VAI01 INT  --单据ID
      ,VAJ01 INT  --明细ID
      ,BBX01 INT  --诊疗ID
      ,BBY01 INT  --收费项目ID
      ,BCE03 VARCHAR(20)  --操作员
      ,FHB11 DATETIME  --操作时间
      ,FHB12 TINYINT  --性质：1=需要隐藏的项目
      ,FHB13 INT  --状态：奇数=隐藏；偶数=还原
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'FHB1'
GO

-- ----------------------------
-- Table structure for IAN1
-- ----------------------------
DROP TABLE [dbo].[IAN1]
GO
CREATE TABLE [dbo].[IAN1] (
[IAN01] int NULL ,
[BCE01] int NULL ,
[IAA01] int NULL ,
[IAN04] varchar(20) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'IAN1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE IAN1—人员对照表
(
        IAN01  INT  --ID
      , BCE01 INT  --员工ID
      , IAA01 INT  --保险机构
      , IAN04  VARCHAR(20)  --医保系统中的人员编码
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'IAN1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE IAN1—人员对照表
(
        IAN01  INT  --ID
      , BCE01 INT  --员工ID
      , IAA01 INT  --保险机构
      , IAN04  VARCHAR(20)  --医保系统中的人员编码
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'IAN1'
GO

-- ----------------------------
-- Table structure for SYS_Users
-- ----------------------------
DROP TABLE [dbo].[SYS_Users]
GO
CREATE TABLE [dbo].[SYS_Users] (
[ID] int NULL ,
[Code] varchar(32) NULL ,
[EmployeeID] int NULL ,
[Name] varchar(20) NULL ,
[FullName] varchar(64) NULL ,
[Password] varchar(128) NULL ,
[Description] varchar(255) NULL ,
[Privilege] tinyint NULL ,
[Authorized] tinyint NULL ,
[CreateDate] datetime NULL ,
[ExpiryDate] datetime NULL ,
[LoginHost] varchar(256) NULL ,
[LoginState] tinyint NULL ,
[WorkPass] varchar(128) NULL ,
[CheckState] tinyint NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'SYS_Users', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE SYS_Users(用户表)
(
        ID          INT  --ID
      ,Code        VARCHAR(32)  --用户, 这里可以自定义用户的代码, 也可以取自员工表的员工工号
      ,EmployeeID  INT  --对照员工表的ID, 如果为0表示不连到员工表
      ,Name        VARCHAR(20)  --如果是员工表的用户，姓名取自员工表
      ,FullName    VARCHAR(64)  --全名
      ,Password    VARCHAR(128)  --密码, 密码以用户名作密钥加密，这样相同明文，密文也不同, 再作Hash操作，变成不可逆
      ,Description VARCHAR(255)  --说明
      ,Privilege   TINYINT  --特权用户, 此类用户是系统内置帐号，不允许外部删除及修改用户名, 此字段不用于编辑
      ,Authorized  TINYINT  --授权： 0=正常, 1=禁用
      ,CreateDate  DATETIME  --创建日期
      ,ExpiryDate  DATETIME  --有效日期, 帐户有效期
      ,LoginHost   VARCHAR(256)  --当前操作员登录的机器名
      ,LoginState  TINYINT  --登录状态  0：无 1：已登录  2：登出  4：签到   8：签出
      ,WorkPass    VARCHAR(128)  --业务密码

)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'SYS_Users'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE SYS_Users(用户表)
(
        ID          INT  --ID
      ,Code        VARCHAR(32)  --用户, 这里可以自定义用户的代码, 也可以取自员工表的员工工号
      ,EmployeeID  INT  --对照员工表的ID, 如果为0表示不连到员工表
      ,Name        VARCHAR(20)  --如果是员工表的用户，姓名取自员工表
      ,FullName    VARCHAR(64)  --全名
      ,Password    VARCHAR(128)  --密码, 密码以用户名作密钥加密，这样相同明文，密文也不同, 再作Hash操作，变成不可逆
      ,Description VARCHAR(255)  --说明
      ,Privilege   TINYINT  --特权用户, 此类用户是系统内置帐号，不允许外部删除及修改用户名, 此字段不用于编辑
      ,Authorized  TINYINT  --授权： 0=正常, 1=禁用
      ,CreateDate  DATETIME  --创建日期
      ,ExpiryDate  DATETIME  --有效日期, 帐户有效期
      ,LoginHost   VARCHAR(256)  --当前操作员登录的机器名
      ,LoginState  TINYINT  --登录状态  0：无 1：已登录  2：登出  4：签到   8：签出
      ,WorkPass    VARCHAR(128)  --业务密码

)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'SYS_Users'
GO

-- ----------------------------
-- Table structure for sysdiagrams
-- ----------------------------
DROP TABLE [dbo].[sysdiagrams]
GO
CREATE TABLE [dbo].[sysdiagrams] (
[name] sysname NOT NULL ,
[principal_id] int NOT NULL ,
[diagram_id] int NOT NULL IDENTITY(1,1) ,
[version] int NULL ,
[definition] varbinary(MAX) NULL 
)


GO

-- ----------------------------
-- Table structure for VAA1
-- ----------------------------
DROP TABLE [dbo].[VAA1]
GO
CREATE TABLE [dbo].[VAA1] (
[VAA01] int NULL ,
[VAA02] varchar(64) NULL ,
[VAA03] varchar(20) NULL ,
[VAA04] varchar(20) NULL ,
[VAA05] varchar(64) NULL ,
[VAA06] varchar(64) NULL ,
[ABBRP] varchar(10) NULL ,
[ABBRW] varchar(10) NULL ,
[ABW01] varchar(1) NULL ,
[VAA10] int NULL ,
[AAU01] varchar(1) NULL ,
[VAA12] datetime NULL ,
[ACK01] varchar(2) NULL ,
[VAA14] varchar(2) NULL ,
[VAA15] varchar(20) NULL ,
[VAA16] varchar(20) NULL ,
[ABJ01] varchar(2) NULL ,
[BDP02] varchar(50) NULL ,
[ABC02] varchar(20) NULL ,
[VAA20] varchar(64) NULL ,
[ACM02] varchar(20) NULL ,
[ACC02] varchar(32) NULL ,
[ABQ02] varchar(32) NULL ,
[VAA25] varchar(64) NULL ,
[VAA26] varchar(20) NULL ,
[VAA27] varchar(20) NULL ,
[VAA28] varchar(128) NULL ,
[VAA29] varchar(6) NULL ,
[VAA30] varchar(20) NULL ,
[VAA31] varchar(20) NULL ,
[VAA32] varchar(20) NULL ,
[VAA33] varchar(64) NULL ,
[VAA34] varchar(20) NULL ,
[VAA35] varchar(13) NULL ,
[VAA36] varchar(128) NULL ,
[VAA37] varchar(64) NULL ,
[VAA38] varchar(48) NULL ,
[VAA39] varchar(64) NULL ,
[VAA40] varchar(20) NULL ,
[VAA41] varchar(32) NULL ,
[VAA42] varchar(64) NULL ,
[VAA43] varchar(20) NULL ,
[VAA44] varchar(16) NULL ,
[BAQ01] int NULL ,
[BAQ02] varchar(10) NULL ,
[VAA47] varchar(64) NULL ,
[VAA48] varchar(20) NULL ,
[VAA49] varchar(64) NULL ,
[VAA50] varchar(6) NULL ,
[VAA51] varchar(64) NULL ,
[VAA52] varchar(20) NULL ,
[VAA53] varchar(20) NULL ,
[VAA54] numeric(18,2) NULL ,
[VAA55] tinyint NULL ,
[VAA56] int NULL ,
[VAA57] datetime NULL ,
[BCK01A] int NULL ,
[VAA61] tinyint NULL ,
[VAA62] varchar(255) NULL ,
[BDX02] varchar(64) NULL ,
[VAA64] varchar(255) NULL ,
[VBU01] int NULL ,
[VAA66] varchar(64) NULL ,
[VAA67] varchar(64) NULL ,
[IAK05] varchar(50) NULL ,
[IAA01] int NULL ,
[BCK01B] int NULL ,
[BCK01C] int NULL ,
[BCQ04] varchar(20) NULL ,
[VAA73] datetime NULL ,
[VAA74] datetime NULL ,
[VAA75] datetime NULL ,
[BEP05] numeric(18,4) NULL ,
[BEP06] numeric(18,4) NULL ,
[ABL01A] varchar(2) NULL ,
[ABL01B] varchar(2) NULL ,
[VAA76] datetime NULL ,
[ABL01] varchar(2) NULL ,
[BEP06B] numeric(18,4) NULL ,
[VAA78] varchar(1) NULL ,
[VAA82] varchar(64) NULL ,
[VAA01A] int NULL ,
[VAA84] varchar(20) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAA1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAA1(病人资料)
(
       VAA01  INT  --病人ID, PATFU, PATID
      ,VAA02  VARCHAR(64)  --会员卡号
      ,VAA03  VARCHAR(20)  --门诊号 Clinic Patient No, Out patient No, OPTNO
      ,VAA04  VARCHAR(20)  --住院号, In Patient No, IPTNO
      ,VAA05  VARCHAR(64)  --姓名, PATNM, FirstName + Middlename
      ,VAA06  VARCHAR(64)  --监护人1
      ,ABBRP  VARCHAR(10)  --拼音
      ,ABBRW  VARCHAR(10)  --五笔
      ,ABW01  VARCHAR(1)  --性别
      ,VAA10  INT  --年龄
      ,AAU01  VARCHAR(1)  --年龄单位
      ,VAA12  DATETIME  --出生日期, birth date
      ,ACK01  VARCHAR(2)  --婚姻, marital status
      ,VAA14  VARCHAR(2)  --身份证件, 指公安机关签发的有效证件
      ,VAA15  VARCHAR(20)  --身份证号, IDNo, IDNumber
      ,VAA16  VARCHAR(20)  --其他证件
      ,ABJ01  VARCHAR(2)  --付款方式, 医疗付款方式
      ,BDP02  VARCHAR(50)  --病人类别
      ,ABC02  VARCHAR(20)  --病人费别
      ,VAA20  VARCHAR(64)  --出生地点, 出生地, BirthPlace
      ,ACM02  VARCHAR(20)  --从业状况, 身份
      ,ACC02  VARCHAR(32)  --国籍, Nationality
      ,ABQ02  VARCHAR(32)  --民族, Nation
      ,VAA25  VARCHAR(64)  --籍贯, 省.市, NativePlace
      ,VAA26  VARCHAR(20)  --宗教, 佛教, 伊斯兰教, 基督教, 天主教, 犹太教, 耶稣教, 其他, Religion, 暂时不用
      ,VAA27  VARCHAR(20)  --种族, 人种, Ethnic, 暂时不用
      ,VAA28  VARCHAR(128)  --户口地址, 户籍地址, RegisteredAddress
      ,VAA29  VARCHAR(6)  --户籍邮编, RegisteredPostCode
      ,VAA30  VARCHAR(20)  --户籍电话, RegisteredPhone
      ,VAA31  VARCHAR(20)  --省市
      ,VAA32  VARCHAR(20)  --县市, 城市, 县区市
      ,VAA33  VARCHAR(64)  --地址, 常住地址, resident address
      ,VAA34  VARCHAR(20)  --电话
      ,VAA35  VARCHAR(13)  --移动电话
      ,VAA36  VARCHAR(128)  --电子邮箱
      ,VAA37  VARCHAR(64)  --其他联系方式, 如QQ号, MSN号, Skype号等, OtherCommunications
      ,VAA38  VARCHAR(48)  --学历
      ,VAA39  VARCHAR(64)  --监护人2, Guardian
      ,VAA40  VARCHAR(20)  --联系人姓名, contact person name
      ,VAA41  VARCHAR(32)  --与病人关系, RelationShip
      ,VAA42  VARCHAR(64)  --联系人地址, Contact person address
      ,VAA43  VARCHAR(20)  --联系人电话, contact person telephone
      ,VAA44  VARCHAR(16)  --联系人移动电话, Contact person Mobile Phone
      ,BAQ01  INT  --合同单位ID
      ,BAQ02  VARCHAR(10)  --单位编码, CoCode
      ,VAA47  VARCHAR(64)  --工作单位, CoName
      ,VAA48  VARCHAR(20)  --单位电话, CoTelephone
      ,VAA49  VARCHAR(64)  --单位地址, CoAddress
      ,VAA50  VARCHAR(6)  --单位邮编, CoPostCode
      ,VAA51  VARCHAR(64)  --单位开户行
      ,VAA52  VARCHAR(20)  --单位银行帐号
      ,VAA53  VARCHAR(20)  --担保人, Guarantor
      ,VAA54  NUMERIC(18, 2)  --信用额度, 担保额度, CreditLimit
      ,VAA55  TINYINT  --担保性质, CreditType
      ,VAA56  INT  --住院次数, HospitalizationNumber
      ,VAA57  DATETIME  --就诊时间, LastVisitDate
      ,BCK01A INT  --就诊科室, LastVisitDeptID
      ,VAA61  TINYINT  --就诊状态, 0=无, 1=门诊, 2=住院, 3=出院, 4=转院, 5=死亡, 9=其他 VisitState
      ,VAA62  VARCHAR(255)  --过敏史
      ,BDX02  VARCHAR(64)  --了解途径, 病人了解医院的方式(如电视广告, 介绍, 户外广告等)
      ,VAA64  VARCHAR(255)  --备注
      ,VBU01  INT  --帐号ID
      ,VAA66  VARCHAR(64)  --病案号
      ,VAA67  VARCHAR(64)  --查询密码
      ,IAK05  VARCHAR(50)  --社会保障号
      ,IAA01  INT  --保险机构
      ,BCK01B INT  --科室ID
      ,BCK01C INT  --病区ID
      ,BCQ04  VARCHAR(20)  --床号
      ,VAA73  DATETIME  --入院时间
      ,VAA74  DATETIME  --出院时间
      ,VAA75  DATETIME  --建档时间
      ,BEP05  NUMERIC(18, 4)  --住院报警值
      ,BEP06  NUMERIC(18, 4)  --住院信用额度
      ,ABL01A VARCHAR(2)  --正定型
      ,ABL01B VARCHAR(2)  --反定型
      ,VAA76  DATETIME  --有效时间
      ,ABL01  VARCHAR(2)  --血型
      ,BEP06B NUMERIC(18, 4)  --门诊信用额度
      ,VAA78  VARCHAR(1)  --Rh血型
      ,VAA82  VARCHAR(64)  --健康卡号
      ,VAA01A INT  --相关ID
      ,VAA84  VARCHAR(20)  --体检登记号
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAA1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAA1(病人资料)
(
       VAA01  INT  --病人ID, PATFU, PATID
      ,VAA02  VARCHAR(64)  --会员卡号
      ,VAA03  VARCHAR(20)  --门诊号 Clinic Patient No, Out patient No, OPTNO
      ,VAA04  VARCHAR(20)  --住院号, In Patient No, IPTNO
      ,VAA05  VARCHAR(64)  --姓名, PATNM, FirstName + Middlename
      ,VAA06  VARCHAR(64)  --监护人1
      ,ABBRP  VARCHAR(10)  --拼音
      ,ABBRW  VARCHAR(10)  --五笔
      ,ABW01  VARCHAR(1)  --性别
      ,VAA10  INT  --年龄
      ,AAU01  VARCHAR(1)  --年龄单位
      ,VAA12  DATETIME  --出生日期, birth date
      ,ACK01  VARCHAR(2)  --婚姻, marital status
      ,VAA14  VARCHAR(2)  --身份证件, 指公安机关签发的有效证件
      ,VAA15  VARCHAR(20)  --身份证号, IDNo, IDNumber
      ,VAA16  VARCHAR(20)  --其他证件
      ,ABJ01  VARCHAR(2)  --付款方式, 医疗付款方式
      ,BDP02  VARCHAR(50)  --病人类别
      ,ABC02  VARCHAR(20)  --病人费别
      ,VAA20  VARCHAR(64)  --出生地点, 出生地, BirthPlace
      ,ACM02  VARCHAR(20)  --从业状况, 身份
      ,ACC02  VARCHAR(32)  --国籍, Nationality
      ,ABQ02  VARCHAR(32)  --民族, Nation
      ,VAA25  VARCHAR(64)  --籍贯, 省.市, NativePlace
      ,VAA26  VARCHAR(20)  --宗教, 佛教, 伊斯兰教, 基督教, 天主教, 犹太教, 耶稣教, 其他, Religion, 暂时不用
      ,VAA27  VARCHAR(20)  --种族, 人种, Ethnic, 暂时不用
      ,VAA28  VARCHAR(128)  --户口地址, 户籍地址, RegisteredAddress
      ,VAA29  VARCHAR(6)  --户籍邮编, RegisteredPostCode
      ,VAA30  VARCHAR(20)  --户籍电话, RegisteredPhone
      ,VAA31  VARCHAR(20)  --省市
      ,VAA32  VARCHAR(20)  --县市, 城市, 县区市
      ,VAA33  VARCHAR(64)  --地址, 常住地址, resident address
      ,VAA34  VARCHAR(20)  --电话
      ,VAA35  VARCHAR(13)  --移动电话
      ,VAA36  VARCHAR(128)  --电子邮箱
      ,VAA37  VARCHAR(64)  --其他联系方式, 如QQ号, MSN号, Skype号等, OtherCommunications
      ,VAA38  VARCHAR(48)  --学历
      ,VAA39  VARCHAR(64)  --监护人2, Guardian
      ,VAA40  VARCHAR(20)  --联系人姓名, contact person name
      ,VAA41  VARCHAR(32)  --与病人关系, RelationShip
      ,VAA42  VARCHAR(64)  --联系人地址, Contact person address
      ,VAA43  VARCHAR(20)  --联系人电话, contact person telephone
      ,VAA44  VARCHAR(16)  --联系人移动电话, Contact person Mobile Phone
      ,BAQ01  INT  --合同单位ID
      ,BAQ02  VARCHAR(10)  --单位编码, CoCode
      ,VAA47  VARCHAR(64)  --工作单位, CoName
      ,VAA48  VARCHAR(20)  --单位电话, CoTelephone
      ,VAA49  VARCHAR(64)  --单位地址, CoAddress
      ,VAA50  VARCHAR(6)  --单位邮编, CoPostCode
      ,VAA51  VARCHAR(64)  --单位开户行
      ,VAA52  VARCHAR(20)  --单位银行帐号
      ,VAA53  VARCHAR(20)  --担保人, Guarantor
      ,VAA54  NUMERIC(18, 2)  --信用额度, 担保额度, CreditLimit
      ,VAA55  TINYINT  --担保性质, CreditType
      ,VAA56  INT  --住院次数, HospitalizationNumber
      ,VAA57  DATETIME  --就诊时间, LastVisitDate
      ,BCK01A INT  --就诊科室, LastVisitDeptID
      ,VAA61  TINYINT  --就诊状态, 0=无, 1=门诊, 2=住院, 3=出院, 4=转院, 5=死亡, 9=其他 VisitState
      ,VAA62  VARCHAR(255)  --过敏史
      ,BDX02  VARCHAR(64)  --了解途径, 病人了解医院的方式(如电视广告, 介绍, 户外广告等)
      ,VAA64  VARCHAR(255)  --备注
      ,VBU01  INT  --帐号ID
      ,VAA66  VARCHAR(64)  --病案号
      ,VAA67  VARCHAR(64)  --查询密码
      ,IAK05  VARCHAR(50)  --社会保障号
      ,IAA01  INT  --保险机构
      ,BCK01B INT  --科室ID
      ,BCK01C INT  --病区ID
      ,BCQ04  VARCHAR(20)  --床号
      ,VAA73  DATETIME  --入院时间
      ,VAA74  DATETIME  --出院时间
      ,VAA75  DATETIME  --建档时间
      ,BEP05  NUMERIC(18, 4)  --住院报警值
      ,BEP06  NUMERIC(18, 4)  --住院信用额度
      ,ABL01A VARCHAR(2)  --正定型
      ,ABL01B VARCHAR(2)  --反定型
      ,VAA76  DATETIME  --有效时间
      ,ABL01  VARCHAR(2)  --血型
      ,BEP06B NUMERIC(18, 4)  --门诊信用额度
      ,VAA78  VARCHAR(1)  --Rh血型
      ,VAA82  VARCHAR(64)  --健康卡号
      ,VAA01A INT  --相关ID
      ,VAA84  VARCHAR(20)  --体检登记号
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAA1'
GO

-- ----------------------------
-- Table structure for VAE1
-- ----------------------------
DROP TABLE [dbo].[VAE1]
GO
CREATE TABLE [dbo].[VAE1] (
[VAE01] int NULL ,
[VAE02] varchar(20) NULL ,
[VAA01] int NULL ,
[VAE04] tinyint NULL ,
[ABJ01] varchar(2) NULL ,
[BDP02] varchar(50) NULL ,
[ABC02] varchar(20) NULL ,
[VAE08] tinyint NULL ,
[BCK01A] int NULL ,
[BCK01B] int NULL ,
[VAE11] datetime NULL ,
[ABO01] varchar(2) NULL ,
[ABR01] varchar(1) NULL ,
[ABT02] varchar(20) NULL ,
[VAE15] varchar(64) NULL ,
[ABZ02] varchar(10) NULL ,
[ABK02] varchar(20) NULL ,
[BCQ04A] varchar(20) NULL ,
[VAE19] int NULL ,
[VAE20] tinyint NULL ,
[AAG01] int NULL ,
[VAE22] varchar(2) NULL ,
[BCK01C] int NULL ,
[BCK01D] int NULL ,
[BCQ04B] varchar(20) NULL ,
[VAE26] datetime NULL ,
[VAE27] numeric(12,2) NULL ,
[ABV01] varchar(2) NULL ,
[VAE29] tinyint NULL ,
[VAE30] datetime NULL ,
[VAE31] tinyint NULL ,
[VAE33] int NULL ,
[VAE34] int NULL ,
[VAE35] tinyint NULL ,
[VAE36] int NULL ,
[VAE37] varchar(2) NULL ,
[VAE38] tinyint NULL ,
[BCE03A] varchar(20) NULL ,
[BCE03B] varchar(20) NULL ,
[BCE03C] varchar(20) NULL ,
[VAE42] varchar(64) NULL ,
[VAE44] tinyint NULL ,
[VAE45] numeric(9) NULL ,
[VAE46] int NULL ,
[AAU01] varchar(1) NULL ,
[ACK01] varchar(2) NULL ,
[AAT02] varchar(128) NULL ,
[ACC02] varchar(32) NULL ,
[AAY02] varchar(48) NULL ,
[BAQ01] int NULL ,
[BAQ02] varchar(10) NULL ,
[BAQ03] varchar(64) NULL ,
[VAE55] varchar(32) NULL ,
[VAE56] varchar(64) NULL ,
[VAE57] varchar(128) NULL ,
[VAE58] varchar(6) NULL ,
[VAE59] varchar(20) NULL ,
[VAE60] varchar(20) NULL ,
[VAE61] varchar(20) NULL ,
[VAE62] varchar(64) NULL ,
[VAE63] varchar(20) NULL ,
[VAE64] varchar(13) NULL ,
[VAE65] varchar(128) NULL ,
[VAE66] varchar(64) NULL ,
[VAE67] varchar(64) NULL ,
[VAE68] varchar(20) NULL ,
[AAZ02] varchar(32) NULL ,
[VAE70] varchar(64) NULL ,
[VAE71] varchar(20) NULL ,
[VAE72] varchar(16) NULL ,
[VAE73] varchar(4) NULL ,
[IAA01] int NULL ,
[UAA01] int NULL ,
[VAE76] tinyint NULL ,
[BCE03D] varchar(20) NULL ,
[VAE78] datetime NULL ,
[VAE79] tinyint NULL ,
[VAE80] tinyint NULL ,
[BCE03E] varchar(20) NULL ,
[VAE82] datetime NULL ,
[VAE83] varchar(255) NULL ,
[VAE84] int NULL ,
[BCE02C] varchar(20) NULL ,
[VAE85] datetime NULL ,
[VAE86] varchar(32) NULL ,
[VAE88] datetime NULL ,
[VAE87] varchar(20) NULL ,
[SCF01] int NULL ,
[VAE89] numeric(9,3) NULL ,
[VAE90] numeric(9,3) NULL ,
[VAE91] numeric(9,2) NULL ,
[VAE92] varchar(64) NULL ,
[VAA08] varchar(20) NULL ,
[VAE94] varchar(20) NULL ,
[VAE95] varchar(255) NULL ,
[VAE96] int NULL ,
[VAE97] varchar(255) NULL ,
[VAE98] varchar(255) NULL ,
[BCE02A] varchar(255) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAE1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAE1 --病人登记记录
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
	VAA08 VARCHAR (20) --病案号''
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAE1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAE1 --病人登记记录
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
	VAA08 VARCHAR (20) --病案号''
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAE1'
GO

-- ----------------------------
-- Table structure for VAF2
-- ----------------------------
DROP TABLE [dbo].[VAF2]
GO
CREATE TABLE [dbo].[VAF2] (
[VAF01] int NULL ,
[VAF01A] int NULL ,
[VAF01B] int NULL ,
[VAF04] tinyint NULL ,
[VAA01] int NULL ,
[VAF06] int NULL ,
[VAF07] int NULL ,
[BCK01A] int NULL ,
[ROWNR] int NULL ,
[VAF10] int NULL ,
[VAF11] tinyint NULL ,
[BDA01] varchar(2) NULL ,
[BBX01] int NULL ,
[VAF14] varchar(60) NULL ,
[VAF15] varchar(30) NULL ,
[BBY01] int NULL ,
[VAF17] int NULL ,
[VAF18] numeric(18,4) NULL ,
[VAF19] varchar(10) NULL ,
[VAF20] numeric(18,4) NULL ,
[VAF21] numeric(18,4) NULL ,
[VAF22] varchar(1024) NULL ,
[VAF23] varchar(128) NULL ,
[BCK01B] int NULL ,
[VAF25] varchar(10) NULL ,
[VAF26] varchar(20) NULL ,
[VAF27] int NULL ,
[VAF28] tinyint NULL ,
[VAF29] varchar(4) NULL ,
[VAF30] varchar(64) NULL ,
[VAF31] tinyint NULL ,
[VAF32] tinyint NULL ,
[VAF33] tinyint NULL ,
[VAF34] tinyint NULL ,
[VAF35] tinyint NULL ,
[VAF36] datetime NULL ,
[VAF37] datetime NULL ,
[VAF38] datetime NULL ,
[BCK01C] int NULL ,
[BCE02A] varchar(20) NULL ,
[BCE03A] varchar(20) NULL ,
[VAF42] datetime NULL ,
[BCE03B] varchar(20) NULL ,
[BCE03C] varchar(20) NULL ,
[VAF45] datetime NULL ,
[BCE03D] varchar(20) NULL ,
[VAF47] datetime NULL ,
[BCE03E] varchar(20) NULL ,
[BCE03F] varchar(20) NULL ,
[VAF50] datetime NULL ,
[VAF51] int NULL ,
[VAF52] tinyint NULL ,
[VAF53] int NULL ,
[VAF54] tinyint NULL ,
[VAF55] varchar(1024) NULL ,
[CBM01] int NULL ,
[BCK01D] int NULL ,
[VAF58] tinyint NULL ,
[VAF59] int NULL ,
[VAF60] varchar(10) NULL ,
[VAF61] numeric(8,2) NULL ,
[VAF62] numeric(8,2) NULL ,
[BCE01A] int NULL ,
[BCE01B] int NULL ,
[BCE01C] int NULL ,
[BCE01D] int NULL ,
[BCE01E] int NULL ,
[BCE01F] int NULL ,
[BCE01G] int NULL ,
[BCE03G] varchar(20) NULL ,
[VAF71] datetime NULL ,
[DSK01] int NULL ,
[VAF01C] int NULL ,
[VAF74] datetime NULL ,
[VAF75] tinyint NULL ,
[BCE01H] int NULL ,
[BCE03H] varchar(20) NULL ,
[BIW02] varchar(64) NULL ,
[Crypt] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAF2', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAF2--住院病人医嘱记录
(
        VAF01  INT  --ID
      ,VAF01A INT  --相关ID, 关联字段：VAF1.VAF01
      ,VAF01B INT  --前提ID, 关联字段：VAF1.VAF01
      ,VAF04  TINYINT  --1：门诊;2：住院
      ,VAA01  INT  --病人ID, 关联字段：VAA1.VAA01
      ,VAF06  INT  --就诊ID, 主页ID
      ,VAF07  INT  --婴儿ID
      ,BCK01A INT  --病人科室ID, 关联字段：BCK1.BCK01
      ,ROWNR  INT  --次序
      ,VAF10  TINYINT  --1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果
      ,VAF11  TINYINT  --医嘱类型, 1=长期医嘱, 2=临时医嘱
      ,BDA01  VARCHAR(2)  --诊疗类型, 关联字段：BDA1.BDA01
      ,BBX01  INT  --诊疗项目ID, 关联字段：BBX1.BBX01
      ,VAF14  VARCHAR(60)  --标本部位
      ,VAF15  VARCHAR(30)  --检查方法
      ,BBY01  INT  --收费项目ID, 关联字段：BBY1.BBY01
      ,VAF17  INT  --天数, day number
      ,VAF18  NUMERIC(18, 4)  --剂量, 单次用量
      ,VAF19  VARCHAR(10)  --用量
      ,VAF20  NUMERIC(18, 4)  --单量
      ,VAF21  NUMERIC(18, 4)  --数量
      ,VAF22  VARCHAR(1024)  --医嘱
      ,VAF23  VARCHAR(128)  --医师嘱托
      ,BCK01B INT  --执行科室ID, 关联字段：BCK1.BCK01
      ,VAF25  VARCHAR(10)  --空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果
      ,VAF26  VARCHAR(20)  --执行频次
      ,VAF27  INT  --频率次数
      ,VAF28  TINYINT  --频率间隔
      ,VAF29  VARCHAR(4)  --间隔单位
      ,VAF30  VARCHAR(64)  --执行时间方案
      ,VAF31  TINYINT  --计价特性  0=正常, 1=自费
      ,VAF32  TINYINT  --0：正常; 1＝给药途径
      ,VAF33  TINYINT  --0：标记未用;1：正常 2：自动停止
      ,VAF34  TINYINT  --可否分零
      ,VAF35  TINYINT  --0：正常 1：紧急
      ,VAF36  DATETIME  --开始执行时间
      ,VAF37  DATETIME  --执行终止时间
      ,VAF38  DATETIME  --上次执行时间
      ,BCK01C INT  --开嘱科室ID, 关联字段：BCK1.BCK01
      ,BCE02A VARCHAR(20)  --医师编码, 关联字段：BCE1.BCE02
      ,BCE03A VARCHAR(20)  --开嘱医师, 关联字段：BCE1.BCE03
      ,VAF42  DATETIME  --开嘱时间
      ,BCE03B VARCHAR(20)  --开嘱护士, 关联字段：BCE1.BCE03
      ,BCE03C VARCHAR(20)  --校对护士, 关联字段：BCE1.BCE03
      ,VAF45  DATETIME  --校对时间
      ,BCE03D VARCHAR(20)  --停嘱医生, 关联字段：BCE1.BCE03
      ,VAF47  DATETIME  --停嘱时间
      ,BCE03E VARCHAR(20)  --停嘱护士, 关联字段：BCE1.BCE03
      ,BCE03F VARCHAR(20)  --停嘱校对护士, 关联字段：BCE1.BCE03
      ,VAF50  DATETIME  --执行停嘱时间
      ,VAF51  INT  --申请ID
      ,VAF52  TINYINT  --0：新开；1：上传
      ,VAF53  INT  --审查结果，用于药品合理用药审核。(描述性医嘱：执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它)
      ,VAF54  TINYINT  --0：否，1：忽略
      ,VAF55  VARCHAR(1024)  --摘要，医嘱备注
      ,CBM01  INT  --医嘱单id, 关联字段：CBM1.CBM01
      ,BCK01D INT  --给药科室, 关联字段：BCK1.BCK01
      ,VAF58  TINYINT  --0：正常， 1：自备药，2：离院带药
      ,VAF59  INT  --组号
      ,VAF60  VARCHAR(10)  --滴速
      ,VAF61  NUMERIC(8, 2)  --首日执行次数
      ,VAF62  NUMERIC(8, 2)  --末日执行次数
      ,BCE01A INT  --开嘱医师ID, 关联字段：BCE1.BCE01
      ,BCE01B INT  --开嘱护士ID, 关联字段：BCE1.BCE01
      ,BCE01C INT  --校对护士ID, 关联字段：BCE1.BCE01
      ,BCE01D INT  --停嘱医师ID, 关联字段：BCE1.BCE01
      ,BCE01E INT  --停嘱护士ID, 关联字段：BCE1.BCE01
      ,BCE01F INT  --停嘱校对护士ID, 关联字段：BCE1.BCE01
      ,BCE01G INT  --操作员ID, 关联字段：BCE1.BCE01
      ,BCE03G VARCHAR(20)  --操作员, 关联字段：BCE1.BCE03
      ,VAF71  DATETIME  --审核时间
      ,DSK01  INT  --药品批次id DSK_ID
      ,VAF01C INT  --原医嘱id  (-1 = 重整医嘱)
      ,VAF74  DATETIME  --重整医嘱时间
      ,VAF75  TINYINT  --药品用药标识
      ,BCE01H INT  --授权医师id, 关联字段：BCE1.BCE01
      ,BCE03H VARCHAR(20)  --授权医师, 关联字段：BCE1.BCE03
      ,BIW02  VARCHAR(64)  --用药目的, 关联字段：BIW1.BIW02
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAF2'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAF2--住院病人医嘱记录
(
        VAF01  INT  --ID
      ,VAF01A INT  --相关ID, 关联字段：VAF1.VAF01
      ,VAF01B INT  --前提ID, 关联字段：VAF1.VAF01
      ,VAF04  TINYINT  --1：门诊;2：住院
      ,VAA01  INT  --病人ID, 关联字段：VAA1.VAA01
      ,VAF06  INT  --就诊ID, 主页ID
      ,VAF07  INT  --婴儿ID
      ,BCK01A INT  --病人科室ID, 关联字段：BCK1.BCK01
      ,ROWNR  INT  --次序
      ,VAF10  TINYINT  --1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果
      ,VAF11  TINYINT  --医嘱类型, 1=长期医嘱, 2=临时医嘱
      ,BDA01  VARCHAR(2)  --诊疗类型, 关联字段：BDA1.BDA01
      ,BBX01  INT  --诊疗项目ID, 关联字段：BBX1.BBX01
      ,VAF14  VARCHAR(60)  --标本部位
      ,VAF15  VARCHAR(30)  --检查方法
      ,BBY01  INT  --收费项目ID, 关联字段：BBY1.BBY01
      ,VAF17  INT  --天数, day number
      ,VAF18  NUMERIC(18, 4)  --剂量, 单次用量
      ,VAF19  VARCHAR(10)  --用量
      ,VAF20  NUMERIC(18, 4)  --单量
      ,VAF21  NUMERIC(18, 4)  --数量
      ,VAF22  VARCHAR(1024)  --医嘱
      ,VAF23  VARCHAR(128)  --医师嘱托
      ,BCK01B INT  --执行科室ID, 关联字段：BCK1.BCK01
      ,VAF25  VARCHAR(10)  --空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果
      ,VAF26  VARCHAR(20)  --执行频次
      ,VAF27  INT  --频率次数
      ,VAF28  TINYINT  --频率间隔
      ,VAF29  VARCHAR(4)  --间隔单位
      ,VAF30  VARCHAR(64)  --执行时间方案
      ,VAF31  TINYINT  --计价特性  0=正常, 1=自费
      ,VAF32  TINYINT  --0：正常; 1＝给药途径
      ,VAF33  TINYINT  --0：标记未用;1：正常 2：自动停止
      ,VAF34  TINYINT  --可否分零
      ,VAF35  TINYINT  --0：正常 1：紧急
      ,VAF36  DATETIME  --开始执行时间
      ,VAF37  DATETIME  --执行终止时间
      ,VAF38  DATETIME  --上次执行时间
      ,BCK01C INT  --开嘱科室ID, 关联字段：BCK1.BCK01
      ,BCE02A VARCHAR(20)  --医师编码, 关联字段：BCE1.BCE02
      ,BCE03A VARCHAR(20)  --开嘱医师, 关联字段：BCE1.BCE03
      ,VAF42  DATETIME  --开嘱时间
      ,BCE03B VARCHAR(20)  --开嘱护士, 关联字段：BCE1.BCE03
      ,BCE03C VARCHAR(20)  --校对护士, 关联字段：BCE1.BCE03
      ,VAF45  DATETIME  --校对时间
      ,BCE03D VARCHAR(20)  --停嘱医生, 关联字段：BCE1.BCE03
      ,VAF47  DATETIME  --停嘱时间
      ,BCE03E VARCHAR(20)  --停嘱护士, 关联字段：BCE1.BCE03
      ,BCE03F VARCHAR(20)  --停嘱校对护士, 关联字段：BCE1.BCE03
      ,VAF50  DATETIME  --执行停嘱时间
      ,VAF51  INT  --申请ID
      ,VAF52  TINYINT  --0：新开；1：上传
      ,VAF53  INT  --审查结果，用于药品合理用药审核。(描述性医嘱：执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它)
      ,VAF54  TINYINT  --0：否，1：忽略
      ,VAF55  VARCHAR(1024)  --摘要，医嘱备注
      ,CBM01  INT  --医嘱单id, 关联字段：CBM1.CBM01
      ,BCK01D INT  --给药科室, 关联字段：BCK1.BCK01
      ,VAF58  TINYINT  --0：正常， 1：自备药，2：离院带药
      ,VAF59  INT  --组号
      ,VAF60  VARCHAR(10)  --滴速
      ,VAF61  NUMERIC(8, 2)  --首日执行次数
      ,VAF62  NUMERIC(8, 2)  --末日执行次数
      ,BCE01A INT  --开嘱医师ID, 关联字段：BCE1.BCE01
      ,BCE01B INT  --开嘱护士ID, 关联字段：BCE1.BCE01
      ,BCE01C INT  --校对护士ID, 关联字段：BCE1.BCE01
      ,BCE01D INT  --停嘱医师ID, 关联字段：BCE1.BCE01
      ,BCE01E INT  --停嘱护士ID, 关联字段：BCE1.BCE01
      ,BCE01F INT  --停嘱校对护士ID, 关联字段：BCE1.BCE01
      ,BCE01G INT  --操作员ID, 关联字段：BCE1.BCE01
      ,BCE03G VARCHAR(20)  --操作员, 关联字段：BCE1.BCE03
      ,VAF71  DATETIME  --审核时间
      ,DSK01  INT  --药品批次id DSK_ID
      ,VAF01C INT  --原医嘱id  (-1 = 重整医嘱)
      ,VAF74  DATETIME  --重整医嘱时间
      ,VAF75  TINYINT  --药品用药标识
      ,BCE01H INT  --授权医师id, 关联字段：BCE1.BCE01
      ,BCE03H VARCHAR(20)  --授权医师, 关联字段：BCE1.BCE03
      ,BIW02  VARCHAR(64)  --用药目的, 关联字段：BIW1.BIW02
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAF2'
GO

-- ----------------------------
-- Table structure for VAJ2
-- ----------------------------
DROP TABLE [dbo].[VAJ2]
GO
CREATE TABLE [dbo].[VAJ2] (
[VAJ01] int NULL ,
[VAA01] int NULL ,
[VAA07] int NULL ,
[VAJ04] int NULL ,
[VAJ05] int NULL ,
[ROWNR] int NULL ,
[VAJ01A] int NULL ,
[VAJ01B] int NULL ,
[VAJ09] int NULL ,
[VAJ10] int NULL ,
[VAI01] int NULL ,
[VAF01] int NULL ,
[VAK01] int NULL ,
[ACF01] int NULL ,
[VAJ15] tinyint NULL ,
[BCK01A] int NULL ,
[BCK01B] int NULL ,
[BDN01] varchar(2) NULL ,
[BBY01] int NULL ,
[BCJ02] varchar(32) NULL ,
[VAJ21] tinyint NULL ,
[VAJ22] tinyint NULL ,
[VAJ23] tinyint NULL ,
[VAJ24] numeric(18,4) NULL ,
[VAJ25] numeric(18,4) NULL ,
[VAJ26] tinyint NULL ,
[VAJ27] int NULL ,
[VAJ28] numeric(9,4) NULL ,
[VAJ29] numeric(18,4) NULL ,
[VAJ30] numeric(9,4) NULL ,
[VAJ31] numeric(9,4) NULL ,
[VAJ32] numeric(18,6) NULL ,
[VAJ33] numeric(18,6) NULL ,
[VAJ34] numeric(18,4) NULL ,
[VAJ35] varchar(20) NULL ,
[VAJ36] numeric(18,2) NULL ,
[VAJ37] numeric(18,2) NULL ,
[VAJ38] numeric(18,2) NULL ,
[VAJ39] tinyint NULL ,
[VAJ40] numeric(18,4) NULL ,
[VAJ41] numeric(18,4) NULL ,
[BCE03A] varchar(20) NULL ,
[BCK01C] int NULL ,
[BCE02B] varchar(20) NULL ,
[BCE03B] varchar(20) NULL ,
[VAJ46] datetime NULL ,
[VAJ47] datetime NULL ,
[VAJ48] int NULL ,
[BCK01D] int NULL ,
[BCE03C] varchar(20) NULL ,
[VAJ51] datetime NULL ,
[VAJ52] datetime NULL ,
[VAJ53] tinyint NULL ,
[VAJ54] varchar(255) NULL ,
[BCE02D] varchar(20) NULL ,
[BCE03D] varchar(20) NULL ,
[VAJ57] varchar(1024) NULL ,
[FAB03] varchar(20) NULL ,
[VAJ59] numeric(18,6) NULL ,
[BCE02C] varchar(20) NULL ,
[VAJ61] numeric(18,2) NULL ,
[VAJ62] datetime NULL ,
[BCK01E] int NULL ,
[VAJ64] datetime NULL ,
[VAJ65] tinyint NULL ,
[DSK01] int NULL ,
[VAJ67] numeric(18,6) NULL ,
[BCE01E] int NULL ,
[BCE03E] varchar(20) NULL ,
[BCK01F] int NULL ,
[BCQ04] varchar(20) NULL ,
[VAJ72] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAJ2', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAJ2(住院病人费用明细)
(
     VAJ01  INT  --ID
      ,VAA01  INT  --病人ID, 关联字段：VAA1.VAA01
      ,VAA07  INT  --就诊ID, 主页ID, 关联字段：VAA1.VAA07
      ,VAJ04  INT  --类型, 1=门诊挂号, 2=门诊划价, 3=门诊记帐, 4=门诊收费, 门诊结帐, 5=医技记帐, 6=住院记帐
      ,VAJ05  INT  --记录状态：1：收费划价/记账划价；2：已收费/已记账；3：已退费/已销账  ; 4：退费记录  是根据类型变化 ; 9：作废
      ,ROWNR  INT  --次序
      ,VAJ01A INT  --上级ID, 从属父号, 关联字段：VAJ1.VAJ01A
      ,VAJ01B INT  --关联ID, 关联字段：VAJ1.VAJ01B
      ,VAJ09  INT  --冲销ID
      ,VAJ10  INT  --合并标志, 0=否, 1=是, merge tag, 多帐单合并, 多病人单
      ,VAI01  INT  --单据ID, 记帐单ID, 关联字段：VAI1.VAI01
      ,VAF01  INT  --医嘱ID, OrderID, 关联字段：VAF1.VAF01
      ,VAK01  INT  --结帐ID, 关联字段：VAK1.VAK01
      ,ACF01  INT  --医疗服务, 不能取0,3值, 关联字段：ACF1.ACF01
      ,VAJ15  TINYINT  --记帐标志
      ,BCK01A INT  --病区ID, 关联字段：BCK1.BCK01
      ,BCK01B INT  --科室ID, 关联字段：BCK1.BCK01
      ,BDN01  VARCHAR(2)  --类型, 编码, 关联字段：BDN1.BDN01
      ,BBY01  INT  --收费项目ID, 关联字段：BBY1.BBY01
      ,BCJ02  VARCHAR(32)  --发药窗口, 关联字段：BCJ1.BCJ02
      ,VAJ21  TINYINT  --加班标志, 暂时不用
      ,VAJ22  TINYINT  --特殊标志, 附加标志, 不同位置用途不同, 挂号时：存储项目特性(1=挂号, 6=诊金, 7=病历本, 8=就诊卡)
      ,VAJ23  TINYINT  --剂数, 中药剂数
      ,VAJ24  NUMERIC(18, 4)  --单量
      ,VAJ25  NUMERIC(18, 4)  --数量, 数次, 总数量
      ,VAJ26  TINYINT  --急诊标志, emergency tag
      ,VAJ27  INT  --婴儿费,对应VAP1表中VAP01
      ,VAJ28  NUMERIC(9, 4)  --税率, 暂时不用
      ,VAJ29  NUMERIC(18, 4)  --税费, 暂时不用
      ,VAJ30  NUMERIC(9, 4)  --折扣率分子, discount rate numerator
      ,VAJ31  NUMERIC(9, 4)  --折扣率分母, discount rate denominator
      ,VAJ32  NUMERIC(18, 6)  --全价
      ,VAJ33  NUMERIC(18, 6)  --单价, 标准单价
      ,VAJ34  NUMERIC(18, 4)  --包装
      ,VAJ35  VARCHAR(20)  --单位, 计算单位
      ,VAJ36  NUMERIC(18, 2)  --全额,原始价格计算得金额
      ,VAJ37  NUMERIC(18, 2)  --应收金额, 未临时打折前的金额(可能经过费别打折)
      ,VAJ38  NUMERIC(18, 2)  --结帐金额(结账时应付金额)，发票打印以此金额为准
      ,VAJ39  TINYINT  --费用标志, 0=正常, 1=自费, 2=免费
      ,VAJ40  NUMERIC(18, 4)  --自负金额
      ,VAJ41  NUMERIC(18, 4)  --保险金额, 统筹金额
      ,BCE03A VARCHAR(20)  --划价人, 关联字段：BCE1.BCE03
      ,BCK01C INT  --开单科室ID, OrderDeptID, 关联字段：BCK1.BCK01
      ,BCE02B VARCHAR(20)  --开单人号, 关联字段：BCE1.BCE02
      ,BCE03B VARCHAR(20)  --开单人, Physician, 一般为医师, 关联字段：BCE1.BCE03
      ,VAJ46  DATETIME  --记帐时间, 手工时间
      ,VAJ47  DATETIME  --交易时间, 机器时间
      ,VAJ48  INT  --执行ID
      ,BCK01D INT  --执行科室ID, 关联字段：BCK1.BCK01
      ,BCE03C VARCHAR(20)  --执行者, 关联字段：BCE1.BCE03
      ,VAJ51  DATETIME  --执行时间
      ,VAJ52  DATETIME  --执行交易时间
      ,VAJ53  TINYINT  --执行情况：0：未执行; 1：执行完成; 2：拒绝执行; 3：正在执行;4：过期挂起
      ,VAJ54  VARCHAR(255)  --备注
      ,BCE02D VARCHAR(20)  --操作员#, 关联字段：BCE1.BCE02
      ,BCE03D VARCHAR(20)  --操作员, 关联字段：BCE1.BCE03
      ,VAJ57  VARCHAR(1024)  --摘要;  收费项目为主从项目时  摘要=主项目名称
      ,FAB03  VARCHAR(20)  --销售单位, 药品门诊或住院单位, 原先为发票号, 关联字段：FAB1.FAB03
      ,VAJ59  NUMERIC(18, 6)  --成本价
      ,BCE02C VARCHAR(20)  --执行者号
      ,VAJ61  NUMERIC(18, 2)  --核算金额，财务核算时用到
      ,VAJ62  DATETIME  --业务时间、默认记账时间，销账时取被销账那条明细的记账时间
      ,BCK01E INT  --给药科室ID, 关联字段：BCK1.BCK01
      ,VAJ64  DATETIME  --发生时间、用于住院长嘱发送时记跨天的费用
      ,VAJ65  TINYINT  --住院中途结帐时，为1参与本次结帐，否则不参与
      ,DSK01  INT  --药品批次id DSK_ID
      ,VAJ67  NUMERIC(18, 6)  --原价
      ,BCE01E INT  --住院医师id, 关联字段：BCE1.BCE01
      ,BCE03E VARCHAR(20)  --住院医师, 关联字段：BCE1.BCE03
      ,BCK01F INT  --病人床位对应病区, 关联字段：BCK1.BCK01
      ,BCQ04  VARCHAR(20)  --病人床号, 关联字段：BCQ1.BCQ04
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAJ2'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAJ2(住院病人费用明细)
(
     VAJ01  INT  --ID
      ,VAA01  INT  --病人ID, 关联字段：VAA1.VAA01
      ,VAA07  INT  --就诊ID, 主页ID, 关联字段：VAA1.VAA07
      ,VAJ04  INT  --类型, 1=门诊挂号, 2=门诊划价, 3=门诊记帐, 4=门诊收费, 门诊结帐, 5=医技记帐, 6=住院记帐
      ,VAJ05  INT  --记录状态：1：收费划价/记账划价；2：已收费/已记账；3：已退费/已销账  ; 4：退费记录  是根据类型变化 ; 9：作废
      ,ROWNR  INT  --次序
      ,VAJ01A INT  --上级ID, 从属父号, 关联字段：VAJ1.VAJ01A
      ,VAJ01B INT  --关联ID, 关联字段：VAJ1.VAJ01B
      ,VAJ09  INT  --冲销ID
      ,VAJ10  INT  --合并标志, 0=否, 1=是, merge tag, 多帐单合并, 多病人单
      ,VAI01  INT  --单据ID, 记帐单ID, 关联字段：VAI1.VAI01
      ,VAF01  INT  --医嘱ID, OrderID, 关联字段：VAF1.VAF01
      ,VAK01  INT  --结帐ID, 关联字段：VAK1.VAK01
      ,ACF01  INT  --医疗服务, 不能取0,3值, 关联字段：ACF1.ACF01
      ,VAJ15  TINYINT  --记帐标志
      ,BCK01A INT  --病区ID, 关联字段：BCK1.BCK01
      ,BCK01B INT  --科室ID, 关联字段：BCK1.BCK01
      ,BDN01  VARCHAR(2)  --类型, 编码, 关联字段：BDN1.BDN01
      ,BBY01  INT  --收费项目ID, 关联字段：BBY1.BBY01
      ,BCJ02  VARCHAR(32)  --发药窗口, 关联字段：BCJ1.BCJ02
      ,VAJ21  TINYINT  --加班标志, 暂时不用
      ,VAJ22  TINYINT  --特殊标志, 附加标志, 不同位置用途不同, 挂号时：存储项目特性(1=挂号, 6=诊金, 7=病历本, 8=就诊卡)
      ,VAJ23  TINYINT  --剂数, 中药剂数
      ,VAJ24  NUMERIC(18, 4)  --单量
      ,VAJ25  NUMERIC(18, 4)  --数量, 数次, 总数量
      ,VAJ26  TINYINT  --急诊标志, emergency tag
      ,VAJ27  INT  --婴儿费,对应VAP1表中VAP01
      ,VAJ28  NUMERIC(9, 4)  --税率, 暂时不用
      ,VAJ29  NUMERIC(18, 4)  --税费, 暂时不用
      ,VAJ30  NUMERIC(9, 4)  --折扣率分子, discount rate numerator
      ,VAJ31  NUMERIC(9, 4)  --折扣率分母, discount rate denominator
      ,VAJ32  NUMERIC(18, 6)  --全价
      ,VAJ33  NUMERIC(18, 6)  --单价, 标准单价
      ,VAJ34  NUMERIC(18, 4)  --包装
      ,VAJ35  VARCHAR(20)  --单位, 计算单位
      ,VAJ36  NUMERIC(18, 2)  --全额,原始价格计算得金额
      ,VAJ37  NUMERIC(18, 2)  --应收金额, 未临时打折前的金额(可能经过费别打折)
      ,VAJ38  NUMERIC(18, 2)  --结帐金额(结账时应付金额)，发票打印以此金额为准
      ,VAJ39  TINYINT  --费用标志, 0=正常, 1=自费, 2=免费
      ,VAJ40  NUMERIC(18, 4)  --自负金额
      ,VAJ41  NUMERIC(18, 4)  --保险金额, 统筹金额
      ,BCE03A VARCHAR(20)  --划价人, 关联字段：BCE1.BCE03
      ,BCK01C INT  --开单科室ID, OrderDeptID, 关联字段：BCK1.BCK01
      ,BCE02B VARCHAR(20)  --开单人号, 关联字段：BCE1.BCE02
      ,BCE03B VARCHAR(20)  --开单人, Physician, 一般为医师, 关联字段：BCE1.BCE03
      ,VAJ46  DATETIME  --记帐时间, 手工时间
      ,VAJ47  DATETIME  --交易时间, 机器时间
      ,VAJ48  INT  --执行ID
      ,BCK01D INT  --执行科室ID, 关联字段：BCK1.BCK01
      ,BCE03C VARCHAR(20)  --执行者, 关联字段：BCE1.BCE03
      ,VAJ51  DATETIME  --执行时间
      ,VAJ52  DATETIME  --执行交易时间
      ,VAJ53  TINYINT  --执行情况：0：未执行; 1：执行完成; 2：拒绝执行; 3：正在执行;4：过期挂起
      ,VAJ54  VARCHAR(255)  --备注
      ,BCE02D VARCHAR(20)  --操作员#, 关联字段：BCE1.BCE02
      ,BCE03D VARCHAR(20)  --操作员, 关联字段：BCE1.BCE03
      ,VAJ57  VARCHAR(1024)  --摘要;  收费项目为主从项目时  摘要=主项目名称
      ,FAB03  VARCHAR(20)  --销售单位, 药品门诊或住院单位, 原先为发票号, 关联字段：FAB1.FAB03
      ,VAJ59  NUMERIC(18, 6)  --成本价
      ,BCE02C VARCHAR(20)  --执行者号
      ,VAJ61  NUMERIC(18, 2)  --核算金额，财务核算时用到
      ,VAJ62  DATETIME  --业务时间、默认记账时间，销账时取被销账那条明细的记账时间
      ,BCK01E INT  --给药科室ID, 关联字段：BCK1.BCK01
      ,VAJ64  DATETIME  --发生时间、用于住院长嘱发送时记跨天的费用
      ,VAJ65  TINYINT  --住院中途结帐时，为1参与本次结帐，否则不参与
      ,DSK01  INT  --药品批次id DSK_ID
      ,VAJ67  NUMERIC(18, 6)  --原价
      ,BCE01E INT  --住院医师id, 关联字段：BCE1.BCE01
      ,BCE03E VARCHAR(20)  --住院医师, 关联字段：BCE1.BCE03
      ,BCK01F INT  --病人床位对应病区, 关联字段：BCK1.BCK01
      ,BCQ04  VARCHAR(20)  --病人床号, 关联字段：BCQ1.BCQ04
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAJ2'
GO

-- ----------------------------
-- Table structure for VAK1
-- ----------------------------
DROP TABLE [dbo].[VAK1]
GO
CREATE TABLE [dbo].[VAK1] (
[VAK01] int NULL ,
[VAA01] int NULL ,
[VAA07] int NULL ,
[VAI04] varchar(20) NULL ,
[FAB03] varchar(20) NULL ,
[VAK06] int NULL ,
[VAK07] numeric(18) NULL ,
[VAK08] numeric(18) NULL ,
[VAK09] numeric(18) NULL ,
[VAK10] numeric(18) NULL ,
[BCE02A] varchar(20) NULL ,
[BCE03A] varchar(20) NULL ,
[VAK13] datetime NULL ,
[VAK01A] int NULL ,
[VAK15] datetime NULL ,
[VAK16] datetime NULL ,
[VAK17] varchar(255) NULL ,
[VAK18] int NULL ,
[VAK19] tinyint NULL ,
[VAK20] numeric(18) NULL ,
[VAK21] numeric(18) NULL ,
[PNRCP] int NULL ,
[VAK23] numeric(18) NULL ,
[VAK24] numeric(18) NULL ,
[BCE01A] int NULL ,
[VAK26] int NULL ,
[VAK27] int NULL ,
[VAK28] int NULL ,
[FAB03A] varchar(20) NULL ,
[PNRCPA] int NULL ,
[BCE03B] varchar(290) NULL ,
[FAC01] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病人结账记录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病人结账记录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'结账ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'结账ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAA01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病人id'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAA01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病人id'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAA01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAA07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'就诊id'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAA07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'就诊id'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAA07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAI04')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'单据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAI04'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'单据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAI04'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'FAB03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'收据号,票据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'FAB03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'收据号,票据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'FAB03'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK06')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'1=门诊挂号,2=门诊收费,3=住院中结,4=住院a收费'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK06'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'1=门诊挂号,2=门诊收费,3=住院中结,4=住院a收费'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK06'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'位数处理的金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'位数处理的金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK08')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'应付,自付金额,发票打印金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK08'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'应付,自付金额,发票打印金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK08'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK09')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'实付'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK09'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'实付'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK09'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK10')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'应找'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK10'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'应找'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK10'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'BCE02A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'收费账号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE02A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'收费账号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE02A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'BCE03A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'收费员'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE03A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'收费员'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE03A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK13')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'交易时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK13'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'交易时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK13'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK01A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'冲销ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK01A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'冲销ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK01A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK15')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'开始时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK15'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'开始时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK15'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK16')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'截止时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK16'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'截止时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK16'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK18')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'缴款ID，如果是退费记录为0'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK18'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'缴款ID，如果是退费记录为0'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK18'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK19')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'1:收费2:退费'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK19'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'1:收费2:退费'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK19'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK20')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'冲预交款金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK20'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'冲预交款金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK20'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK21')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'费用总额(V36合计)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK21'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'费用总额(V36合计)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK21'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'PNRCP')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'冲预交款金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'PNRCP'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'冲预交款金额'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'PNRCP'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK23')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'未临时打折前的金额(VAJ38合计)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK23'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'未临时打折前的金额(VAJ38合计)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK23'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK24')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'财务核对时用到(VAJ61合计)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK24'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'财务核对时用到(VAJ61合计)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK24'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'BCE01A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'收费员ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE01A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'收费员ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE01A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK26')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'分单结账'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK26'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'分单结账'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK26'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK27')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'传输标志，财务接口对外传输时用'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK27'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'传输标志，财务接口对外传输时用'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK27'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'VAK28')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'批次ID，分单结账时一起结账的记录写入第一条结账ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK28'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'批次ID，分单结账时一起结账的记录写入第一条结账ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'VAK28'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'FAB03A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'打印收据时管理收据号参数有效时，此字段写入起始收据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'FAB03A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'打印收据时管理收据号参数有效时，此字段写入起始收据号'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'FAB03A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'PNRCPA')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'收据打印次数'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'PNRCPA'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'收据打印次数'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'PNRCPA'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'BCE03B')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'退费授权人(性能(工号))'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE03B'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'退费授权人(性能(工号))'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'BCE03B'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAK1', 
'COLUMN', N'FAC01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'票据打印内容'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'FAC01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'票据打印内容'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAK1'
, @level2type = 'COLUMN', @level2name = N'FAC01'
GO

-- ----------------------------
-- Table structure for VAM1
-- ----------------------------
DROP TABLE [dbo].[VAM1]
GO
CREATE TABLE [dbo].[VAM1] (
[VAA07] int NULL ,
[VAM02] varchar(10) NULL ,
[VAM03] tinyint NULL ,
[VAM05] varchar(1) NULL ,
[VAM06] tinyint NULL ,
[VAM07] varchar(2) NULL ,
[VAM08] tinyint NULL ,
[VAM09] tinyint NULL ,
[VAM10] tinyint NULL ,
[VAM11] tinyint NULL ,
[VAM12] tinyint NULL ,
[VAM13] tinyint NULL ,
[VAM14] varchar(1) NULL ,
[VAM15] varchar(1) NULL ,
[VAM16] varchar(1) NULL ,
[VAM17] varchar(1) NULL ,
[VAM18] tinyint NULL ,
[VAM19] tinyint NULL ,
[VAM20] tinyint NULL ,
[VAM21] tinyint NULL ,
[VAM22] tinyint NULL ,
[VAM23] tinyint NULL ,
[VAM91] tinyint NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病人登记附表'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病人登记附表'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAA07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'就诊ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAA07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'就诊ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAA07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM02')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病室'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM02'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病室'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM02'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM03')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'有无转科，0：否  1：是'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM03'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'有无转科，0：否  1：是'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM03'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM05')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病例分型, A, B, C, D'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM05'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病例分型, A, B, C, D'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM05'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM06')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'入院前经外院诊治,值域: 0=否, 1=是'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM06'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'入院前经外院诊治,值域: 0=否, 1=是'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM06'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'最高诊断依据'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'最高诊断依据'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM08')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'HBsAg 0=未做、1=阴性、2=阳性'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM08'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'HBsAg 0=未做、1=阴性、2=阳性'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM08'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM09')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'HCV-Ab 0=未做、1=阴性、2=阳性'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM09'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'HCV-Ab 0=未做、1=阴性、2=阳性'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM09'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM10')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'HIV-Ab 0=未做、1=阴性、2=阳性'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM10'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'HIV-Ab 0=未做、1=阴性、2=阳性'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM10'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM11')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疑难, 中医住院期间病情'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM11'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疑难, 中医住院期间病情'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM11'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM12')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'危重, 中医住院期间病情'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM12'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'危重, 中医住院期间病情'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM12'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM13')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'急症, 中医住院期间病情'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM13'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'急症, 中医住院期间病情'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM13'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM14')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'辨证，准确度'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM14'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'辨证，准确度'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM14'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM15')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'治法, 准确度'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM15'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'治法, 准确度'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM15'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM16')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'方药, 准确度'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM16'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'方药, 准确度'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM16'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM17')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'抢救方法, 值域: 1=中, 2=西, 3=中西'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM17'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'抢救方法, 值域: 1=中, 2=西, 3=中西'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM17'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM18')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'自制中药, Homemade, 值域: 0=无, 1=有'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM18'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'自制中药, Homemade, 值域: 0=无, 1=有'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM18'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM19')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'中医门诊与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM19'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'中医门诊与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM19'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM20')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'中医入院与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM20'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'中医入院与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM20'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM21')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'门诊与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM21'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'门诊与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM21'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM22')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'入院与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM22'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'入院与出院 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM22'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM23')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'术前与术后, Ante-Post 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM23'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'术前与术后, Ante-Post 1未做、2=符合、3=不符合、4=不确定'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM23'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAM1', 
'COLUMN', N'VAM91')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'0=未欠费、1=欠费'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM91'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'0=未欠费、1=欠费'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAM1'
, @level2type = 'COLUMN', @level2name = N'VAM91'
GO

-- ----------------------------
-- Table structure for VAO2
-- ----------------------------
DROP TABLE [dbo].[VAO2]
GO
CREATE TABLE [dbo].[VAO2] (
[VAO01] int NULL ,
[VAA01] int NULL ,
[VAA07] int NULL ,
[VAF01] int NULL ,
[ACF01] int NULL ,
[VAO06] int NULL ,
[VAO07] int NULL ,
[VAO08] int NULL ,
[VAO09] int NULL ,
[VAO10] tinyint NULL ,
[VAO11] tinyint NULL ,
[BAK01A] int NULL ,
[CAM01] int NULL ,
[BAK01B] int NULL ,
[VAO15] varchar(150) NULL ,
[ABX01] varchar(1) NULL ,
[VAO17] tinyint NULL ,
[VAO18] tinyint NULL ,
[VAO19] datetime NULL ,
[BCE03A] varchar(20) NULL ,
[VAO21] datetime NULL ,
[BCE03B] varchar(20) NULL ,
[VAO22] tinyint NULL ,
[VAO24] tinyint NULL ,
[VAO25] tinyint NULL ,
[VAO26] varchar(128) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'住院病人诊断记录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'住院病人诊断记录'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAA01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病人ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAA01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病人ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAA01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAA07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'就诊ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAA07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'就诊ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAA07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAF01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'医嘱单诊断时对应的是CBM01;记账单诊断时对应的是VAI01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAF01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'医嘱单诊断时对应的是CBM01;记账单诊断时对应的是VAI01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAF01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'ACF01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'1=门诊, 2=住院, 3=病案, 关联字段:ACF1.ACF01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'ACF01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'1=门诊, 2=住院, 3=病案, 关联字段:ACF1.ACF01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'ACF01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO06')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'诊断次序'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO06'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'诊断次序'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO06'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO07')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'0=初步诊断, 1=正常诊断'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO07'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'0=初步诊断, 1=正常诊断'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO07'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO08')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病历ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO08'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病历ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO08'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO09')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'病例ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO09'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'病例ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO09'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO10')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'0=西医,1=中医'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO10'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'0=西医,1=中医'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO10'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO11')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'1=门诊诊断, 2=入院诊断, 3=出院诊断, 4=, 5=医院感染, 6=病理诊断, 7=损伤中毒, 8=术前诊断, 9=术后诊断, 10=并发症, 11=病原学诊断，12＝尸检主要诊断'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO11'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'1=门诊诊断, 2=入院诊断, 3=出院诊断, 4=, 5=医院感染, 6=病理诊断, 7=损伤中毒, 8=术前诊断, 9=术后诊断, 10=并发症, 11=病原学诊断，12＝尸检主要诊断'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO11'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'BAK01A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'疾病ID, 关联字段:BAK1.BAK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BAK01A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'疾病ID, 关联字段:BAK1.BAK01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BAK01A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'CAM01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'诊断ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'CAM01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'诊断ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'CAM01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'BAK01B')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'症候ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BAK01B'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'症候ID'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BAK01B'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO15')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'诊断途述'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO15'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'诊断途述'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO15'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'ABX01')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'出院情况, 关联字段:ABX1.ABX01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'ABX01'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'出院情况, 关联字段:ABX1.ABX01'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'ABX01'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO17')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N' 未治'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO17'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N' 未治'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO17'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO18')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'1=疑诊,  其它正常'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO18'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'1=疑诊,  其它正常'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO18'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO19')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'记录日期'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO19'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'记录日期'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO19'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'BCE03A')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'记录人'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BCE03A'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'记录人'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BCE03A'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO21')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'取消时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO21'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'取消时间'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO21'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'BCE03B')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'取消人'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BCE03B'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'取消人'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'BCE03B'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO22')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'0:其它地方诊断 1：医嘱单诊断 2:记账单诊断'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO22'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'0:其它地方诊断 1：医嘱单诊断 2:记账单诊断'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO22'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO24')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'入院病情: 1.有，2.临床未确定，3.情况不明，4.无'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO24'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'入院病情: 1.有，2.临床未确定，3.情况不明，4.无'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO24'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO25')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'诊断标识  1：主要诊断，2：次要诊断 (对应与VAO11字段类型中再分的)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO25'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'诊断标识  1：主要诊断，2：次要诊断 (对应与VAO11字段类型中再分的)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO25'
GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAO2', 
'COLUMN', N'VAO26')) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'中医症候名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO26'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'中医症候名称'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAO2'
, @level2type = 'COLUMN', @level2name = N'VAO26'
GO

-- ----------------------------
-- Table structure for VAT1
-- ----------------------------
DROP TABLE [dbo].[VAT1]
GO
CREATE TABLE [dbo].[VAT1] (
[VAT01] int NULL ,
[VAA01] int NULL ,
[VAA07] int NULL ,
[VAT04] int NULL ,
[ACF01] int NULL ,
[VAF01] int NULL ,
[VAQ01] int NULL ,
[VAT08] datetime NULL ,
[VAT09] datetime NULL ,
[VAT10] datetime NULL ,
[ACH01] varchar(2) NULL ,
[ACZ01] varchar(2) NULL ,
[VAT13] int NULL ,
[VAT14] datetime NULL ,
[VAT15] datetime NULL ,
[VAT16] varchar(128) NULL ,
[BBX01] int NULL ,
[ACI02] varchar(30) NULL ,
[VAT19] varchar(10) NULL ,
[VAT20] numeric(18,2) NULL ,
[VAT21] datetime NULL ,
[VAT22] datetime NULL ,
[AAB01] varchar(2) NULL ,
[ABD01] varchar(2) NULL ,
[VAT25] tinyint NULL ,
[VAT26] varchar(128) NULL ,
[VAT27] varchar(10) NULL ,
[VAT28] varchar(10) NULL ,
[VAT29] varchar(10) NULL ,
[BEE03] varchar(20) NULL ,
[BEE01] int NULL ,
[BCK01] int NULL ,
[VAT33] varchar(255) NULL ,
[VAT34] datetime NULL ,
[BCE03A] varchar(20) NULL ,
[VAT36] datetime NULL ,
[BCE03B] varchar(20) NULL ,
[VAT38] int NULL ,
[VAT39] datetime NULL ,
[BJX01] varchar(2) NULL ,
[BJY01] varchar(4) NULL ,
[VCY01] int NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VAT1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAT1-- 病人手术记录
(
       VAT01  INT  --ID
      ,VAA01  INT  --病人ID
      ,VAA07  INT  --就诊ID
      ,VAT04  INT  --手术状态, 0=申请, 1=审核, 2=安排, 3=手术, 4=完成, 9=拒绝
      ,ACF01  INT  --医疗服务
      ,VAF01  INT  --医嘱ID
      ,VAQ01  INT  --病历ID
      ,VAT08  DATETIME  --手术日期
      ,VAT09  DATETIME  --手术开始时间
      ,VAT10  DATETIME  --手术结束时间
      ,ACH01  VARCHAR(2)  --手术级别
      ,ACZ01  VARCHAR(2)  --紧急程度
      ,VAT13  INT  --接台手术
      ,VAT14  DATETIME  --麻醉开始时间
      ,VAT15  DATETIME  --麻醉结束时间
      ,VAT16  VARCHAR(128)  --麻醉方式
      ,BBX01  INT  --麻醉方式ID, 诊疗项目.ID
      ,ACI02  VARCHAR(30)  --麻醉类型
      ,VAT19  VARCHAR(10)  --1=优,2=佳,3=劣,4=危(急)
      ,VAT20  NUMERIC(18, 2)  --输液总量
      ,VAT21  DATETIME  --输氧开始时间
      ,VAT22  DATETIME  --输氧结束时间
      ,AAB01  VARCHAR(2)  --切口等级
      ,ABD01  VARCHAR(2)  --愈合等级
      ,VAT25  TINYINT  --0=污染手术,1=感染手术
      ,VAT26  VARCHAR(128)  --污染内容(艾滋病(HIV), 乙肝, 丙肝, 梅毒), 可多选用;分开
      ,VAT27  VARCHAR(10)  --1=良好, 2=好, 3=坏
      ,VAT28  VARCHAR(10)  --1=良好, 2=好, 3=坏
      ,VAT29  VARCHAR(10)  --1=良好, 2=好, 3=坏
      ,BEE03  VARCHAR(20)  --手术间
      ,BEE01  INT  --手术间ID
      ,BCK01  INT  --手术室ID
      ,VAT33  VARCHAR(255)  --说明
      ,VAT34  DATETIME  --记录日期
      ,BCE03A VARCHAR(20)  --记录人
      ,VAT36  DATETIME  --取消时间
      ,BCE03B VARCHAR(20)  --取消人
      ,VAT38  INT  --术前准备时间 （单位：天）
      ,VAT39  DATETIME  --术前预防性抗菌药物给药时间
      ,BJX01  VARCHAR(2)  --ASA麻醉分级
      ,BJY01  VARCHAR(4)  --操作部位
      ,VCY01  INT  --复诊预约id
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAT1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VAT1-- 病人手术记录
(
       VAT01  INT  --ID
      ,VAA01  INT  --病人ID
      ,VAA07  INT  --就诊ID
      ,VAT04  INT  --手术状态, 0=申请, 1=审核, 2=安排, 3=手术, 4=完成, 9=拒绝
      ,ACF01  INT  --医疗服务
      ,VAF01  INT  --医嘱ID
      ,VAQ01  INT  --病历ID
      ,VAT08  DATETIME  --手术日期
      ,VAT09  DATETIME  --手术开始时间
      ,VAT10  DATETIME  --手术结束时间
      ,ACH01  VARCHAR(2)  --手术级别
      ,ACZ01  VARCHAR(2)  --紧急程度
      ,VAT13  INT  --接台手术
      ,VAT14  DATETIME  --麻醉开始时间
      ,VAT15  DATETIME  --麻醉结束时间
      ,VAT16  VARCHAR(128)  --麻醉方式
      ,BBX01  INT  --麻醉方式ID, 诊疗项目.ID
      ,ACI02  VARCHAR(30)  --麻醉类型
      ,VAT19  VARCHAR(10)  --1=优,2=佳,3=劣,4=危(急)
      ,VAT20  NUMERIC(18, 2)  --输液总量
      ,VAT21  DATETIME  --输氧开始时间
      ,VAT22  DATETIME  --输氧结束时间
      ,AAB01  VARCHAR(2)  --切口等级
      ,ABD01  VARCHAR(2)  --愈合等级
      ,VAT25  TINYINT  --0=污染手术,1=感染手术
      ,VAT26  VARCHAR(128)  --污染内容(艾滋病(HIV), 乙肝, 丙肝, 梅毒), 可多选用;分开
      ,VAT27  VARCHAR(10)  --1=良好, 2=好, 3=坏
      ,VAT28  VARCHAR(10)  --1=良好, 2=好, 3=坏
      ,VAT29  VARCHAR(10)  --1=良好, 2=好, 3=坏
      ,BEE03  VARCHAR(20)  --手术间
      ,BEE01  INT  --手术间ID
      ,BCK01  INT  --手术室ID
      ,VAT33  VARCHAR(255)  --说明
      ,VAT34  DATETIME  --记录日期
      ,BCE03A VARCHAR(20)  --记录人
      ,VAT36  DATETIME  --取消时间
      ,BCE03B VARCHAR(20)  --取消人
      ,VAT38  INT  --术前准备时间 （单位：天）
      ,VAT39  DATETIME  --术前预防性抗菌药物给药时间
      ,BJX01  VARCHAR(2)  --ASA麻醉分级
      ,BJY01  VARCHAR(4)  --操作部位
      ,VCY01  INT  --复诊预约id
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VAT1'
GO

-- ----------------------------
-- Table structure for VCF1
-- ----------------------------
DROP TABLE [dbo].[VCF1]
GO
CREATE TABLE [dbo].[VCF1] (
[VCF01] int NULL ,
[ACF01] int NULL ,
[VAA01] int NULL ,
[VAA07] int NULL ,
[VAP01] int NULL ,
[BCK01] int NULL ,
[AAG01] int NULL ,
[VCF08] tinyint NULL ,
[VCF09] datetime NULL ,
[VCF10] datetime NULL ,
[VCF11] varchar(10) NULL ,
[BCE01A] int NULL ,
[BCE03A] varchar(20) NULL ,
[VCF14] tinyint NULL ,
[BCE01B] int NULL ,
[BCE03B] varchar(20) NULL ,
[VCF17] datetime NULL ,
[VCF18] tinyint NULL ,
[BCE03C] varchar(20) NULL ,
[VCF20] datetime NULL ,
[VCF21] varchar(20) NULL ,
[VCF22] varchar(1) NULL ,
[VCF23] varchar(1) NULL ,
[VCF24] varchar(10) NULL ,
[VCF25] varchar(10) NULL ,
[VCF26] varchar(4) NULL ,
[VCF27] varchar(1) NULL ,
[VCF28] varchar(10) NULL ,
[VCF29] varchar(10) NULL ,
[VCF30] varchar(10) NULL ,
[VCF31] varchar(10) NULL ,
[VCF32] varchar(1) NULL ,
[VCF33] varchar(10) NULL ,
[VCF34] varchar(10) NULL ,
[VCF35] varchar(10) NULL ,
[VCF36] varchar(10) NULL ,
[VCF37] varchar(10) NULL ,
[VCF38] varchar(10) NULL ,
[VCF39] varchar(128) NULL ,
[VCF40] varchar(10) NULL ,
[VCF41] varchar(10) NULL ,
[VCF42] varchar(10) NULL ,
[VCF43] varchar(10) NULL ,
[VCF44] varchar(10) NULL ,
[VCF45] varchar(10) NULL ,
[VCF46] varchar(10) NULL ,
[VCF47] varchar(10) NULL ,
[VCF48] varchar(10) NULL ,
[VCF49] varchar(10) NULL ,
[VCF50] varchar(128) NULL ,
[VCF51] varchar(10) NULL ,
[VCF52] varchar(128) NULL ,
[VCF53] varchar(10) NULL ,
[VCF54] varchar(1) NULL ,
[VCF55] varchar(10) NULL ,
[VCF56] varchar(5) NULL ,
[VCF57] varchar(1) NULL ,
[VCF58] varchar(5) NULL ,
[VCF59] varchar(10) NULL ,
[VCF60] varchar(50) NULL ,
[VCF61] varchar(10) NULL ,
[VCF62] varchar(10) NULL ,
[VCF63] varchar(10) NULL ,
[VCF64] varchar(10) NULL ,
[VCF65] varchar(128) NULL ,
[VCF66] varchar(128) NULL ,
[VCF67] varchar(128) NULL ,
[VCF68] varchar(128) NULL ,
[VCF69] varchar(128) NULL ,
[VCF70] varchar(128) NULL ,
[VCF71] varchar(128) NULL ,
[VCF72] varchar(128) NULL ,
[VCF73] varchar(128) NULL ,
[VCF74] varchar(128) NULL ,
[VCF75] varchar(128) NULL ,
[VCF76] varchar(128) NULL ,
[VCF77] varchar(128) NULL ,
[VCF78] varchar(128) NULL ,
[VCF79] varchar(128) NULL ,
[VCF80] varchar(128) NULL ,
[VCF81] varchar(128) NULL ,
[VCF82] varchar(128) NULL ,
[VCF83] varchar(128) NULL ,
[VCF84] varchar(128) NULL ,
[VAQ01] int NULL ,
[VCF86] tinyint NULL ,
[VCF87] varchar(10) NULL 
)


GO
IF ((SELECT COUNT(*) from fn_listextendedproperty('MS_Description', 
'SCHEMA', N'dbo', 
'TABLE', N'VCF1', 
NULL, NULL)) > 0) 
EXEC sp_updateextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VCF1(病人护理记录)(
 VCF01  INT  --ID
      ,ACF01  INT  --医疗服务
      ,VAA01  INT  --病人ID
      ,VAA07  INT  --就诊ID
      ,VAP01  INT  --婴儿ID
      ,BCK01  INT  --科室ID
      ,AAG01  INT  --护理级别
      ,VCF08  TINYINT  --性质, 1=图标, 2=文档
      ,VCF09  DATETIME  --业务时间
      ,VCF10  DATETIME  --操作时间
      ,VCF11  VARCHAR(10)  --情况： 01-入院, 02=手术, 03=分娩, 04=转入, 05=回室, 06=出院, 07=死亡,  08=转出, 11=拒测, 12=请假, 13=外出
      ,BCE01A INT  --护士ID
      ,BCE03A VARCHAR(20)  --护士
      ,VCF14  TINYINT  --复试合格  0：不合格 1：合格
      ,BCE01B INT  --更新人ID
      ,BCE03B VARCHAR(20)  --更新人
      ,VCF17  DATETIME  --更新时间
      ,VCF18  TINYINT  --状态, 1=记录, 2=审核, 3=归档
      ,BCE03C VARCHAR(20)  --归档人
      ,VCF20  DATETIME  --归档时间
      ,VCF21  VARCHAR(20)  --班次
      ,VCF22  VARCHAR(1)  --体温部位, 1=腋温, 2=口温, 3=肛温
      ,VCF23  VARCHAR(1)  --0=自然体温, 1=物理降温
      ,VCF24  VARCHAR(10)  --体温
      ,VCF25  VARCHAR(10)  --体温单位
      ,VCF26  VARCHAR(4)  --体温符号
      ,VCF27  VARCHAR(1)  --起博器 0=自然心率, 1=起博器
      ,VCF28  VARCHAR(10)  --心率
      ,VCF29  VARCHAR(10)  --心率单位
      ,VCF30  VARCHAR(10)  --脉博
      ,VCF31  VARCHAR(10)  --脉博单位
      ,VCF32  VARCHAR(1)  --呼吸方式 0=自主呼吸, 1=呼吸机
      ,VCF33  VARCHAR(10)  --呼吸
      ,VCF34  VARCHAR(10)  --呼吸单位
      ,VCF35  VARCHAR(10)  --身高
      ,VCF36  VARCHAR(10)  --身高单位
      ,VCF37  VARCHAR(10)  --体重
      ,VCF38  VARCHAR(10)  --体重单位
      ,VCF39  VARCHAR(128)  --皮试药物
      ,VCF40  VARCHAR(10)  --皮试结果
      ,VCF41  VARCHAR(10)  --收缩压
      ,VCF42  VARCHAR(10)  --舒张压
      ,VCF43  VARCHAR(10)  --血压单位
      ,VCF44  VARCHAR(10)  --血氧饱和度 SO2
      ,VCF45  VARCHAR(10)  --意识, 神志
      ,VCF46  VARCHAR(10)  --瞳孔大小左
      ,VCF47  VARCHAR(10)  --瞳孔大小右
      ,VCF48  VARCHAR(10)  --瞳孔反射左
      ,VCF49  VARCHAR(10)  --瞳孔反射右
      ,VCF50  VARCHAR(128)  --摄入药物
      ,VCF51  VARCHAR(10)  --摄入药量
      ,VCF52  VARCHAR(128)  --摄入食物
      ,VCF53  VARCHAR(10)  --摄入食量
      ,VCF54  VARCHAR(1)  --排尿方式, 1=导尿  2=失禁, 4=预置导尿
      ,VCF55  VARCHAR(10)  --尿量
      ,VCF56  VARCHAR(5)  --尿次
      ,VCF57  VARCHAR(1)  --大便方式 0=正常 1=灌肠, 2=失禁, 4=人工肛门
      ,VCF58  VARCHAR(5)  --大便次数
      ,VCF59  VARCHAR(10)  --大便量
      ,VCF60  VARCHAR(50)  --排出物
      ,VCF61  VARCHAR(10)  --排出量
      ,VCF62  VARCHAR(10)  --翻身体位
      ,VCF63  VARCHAR(10)  --腹围
      ,VCF64  VARCHAR(10)  --腹围单位
      ,VCF65  VARCHAR(128)  --其他1
      ,VCF66  VARCHAR(128)  --其他2
      ,VCF67  VARCHAR(128)  --其他3
      ,VCF68  VARCHAR(128)  --其他4
      ,VCF69  VARCHAR(128)  --其他5
      ,VCF70  VARCHAR(128)  --其他6
      ,VCF71  VARCHAR(128)  --其他7
      ,VCF72  VARCHAR(128)  --其他8
      ,VCF73  VARCHAR(128)  --其他9
      ,VCF74  VARCHAR(128)  --其他10
      ,VCF75  VARCHAR(128)  --其他11
      ,VCF76  VARCHAR(128)  --其他12
      ,VCF77  VARCHAR(128)  --其他13
      ,VCF78  VARCHAR(128)  --其他14
      ,VCF79  VARCHAR(128)  --其他15
      ,VCF80  VARCHAR(128)  --其他16
      ,VCF81  VARCHAR(128)  --其他17
      ,VCF82  VARCHAR(128)  --其他18
      ,VCF83  VARCHAR(128)  --其他19
      ,VCF84  VARCHAR(128)  --其他20
      ,VAQ01  INT  --病人病历ID
      ,VCF86  TINYINT  --时区 1：3时   2：7时以此类推
      ,VCF87  VARCHAR(10)  --疼痛强度，0~10
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VCF1'
ELSE
EXEC sp_addextendedproperty @name = N'MS_Description', @value = N'CREATE TABLE VCF1(病人护理记录)(
 VCF01  INT  --ID
      ,ACF01  INT  --医疗服务
      ,VAA01  INT  --病人ID
      ,VAA07  INT  --就诊ID
      ,VAP01  INT  --婴儿ID
      ,BCK01  INT  --科室ID
      ,AAG01  INT  --护理级别
      ,VCF08  TINYINT  --性质, 1=图标, 2=文档
      ,VCF09  DATETIME  --业务时间
      ,VCF10  DATETIME  --操作时间
      ,VCF11  VARCHAR(10)  --情况： 01-入院, 02=手术, 03=分娩, 04=转入, 05=回室, 06=出院, 07=死亡,  08=转出, 11=拒测, 12=请假, 13=外出
      ,BCE01A INT  --护士ID
      ,BCE03A VARCHAR(20)  --护士
      ,VCF14  TINYINT  --复试合格  0：不合格 1：合格
      ,BCE01B INT  --更新人ID
      ,BCE03B VARCHAR(20)  --更新人
      ,VCF17  DATETIME  --更新时间
      ,VCF18  TINYINT  --状态, 1=记录, 2=审核, 3=归档
      ,BCE03C VARCHAR(20)  --归档人
      ,VCF20  DATETIME  --归档时间
      ,VCF21  VARCHAR(20)  --班次
      ,VCF22  VARCHAR(1)  --体温部位, 1=腋温, 2=口温, 3=肛温
      ,VCF23  VARCHAR(1)  --0=自然体温, 1=物理降温
      ,VCF24  VARCHAR(10)  --体温
      ,VCF25  VARCHAR(10)  --体温单位
      ,VCF26  VARCHAR(4)  --体温符号
      ,VCF27  VARCHAR(1)  --起博器 0=自然心率, 1=起博器
      ,VCF28  VARCHAR(10)  --心率
      ,VCF29  VARCHAR(10)  --心率单位
      ,VCF30  VARCHAR(10)  --脉博
      ,VCF31  VARCHAR(10)  --脉博单位
      ,VCF32  VARCHAR(1)  --呼吸方式 0=自主呼吸, 1=呼吸机
      ,VCF33  VARCHAR(10)  --呼吸
      ,VCF34  VARCHAR(10)  --呼吸单位
      ,VCF35  VARCHAR(10)  --身高
      ,VCF36  VARCHAR(10)  --身高单位
      ,VCF37  VARCHAR(10)  --体重
      ,VCF38  VARCHAR(10)  --体重单位
      ,VCF39  VARCHAR(128)  --皮试药物
      ,VCF40  VARCHAR(10)  --皮试结果
      ,VCF41  VARCHAR(10)  --收缩压
      ,VCF42  VARCHAR(10)  --舒张压
      ,VCF43  VARCHAR(10)  --血压单位
      ,VCF44  VARCHAR(10)  --血氧饱和度 SO2
      ,VCF45  VARCHAR(10)  --意识, 神志
      ,VCF46  VARCHAR(10)  --瞳孔大小左
      ,VCF47  VARCHAR(10)  --瞳孔大小右
      ,VCF48  VARCHAR(10)  --瞳孔反射左
      ,VCF49  VARCHAR(10)  --瞳孔反射右
      ,VCF50  VARCHAR(128)  --摄入药物
      ,VCF51  VARCHAR(10)  --摄入药量
      ,VCF52  VARCHAR(128)  --摄入食物
      ,VCF53  VARCHAR(10)  --摄入食量
      ,VCF54  VARCHAR(1)  --排尿方式, 1=导尿  2=失禁, 4=预置导尿
      ,VCF55  VARCHAR(10)  --尿量
      ,VCF56  VARCHAR(5)  --尿次
      ,VCF57  VARCHAR(1)  --大便方式 0=正常 1=灌肠, 2=失禁, 4=人工肛门
      ,VCF58  VARCHAR(5)  --大便次数
      ,VCF59  VARCHAR(10)  --大便量
      ,VCF60  VARCHAR(50)  --排出物
      ,VCF61  VARCHAR(10)  --排出量
      ,VCF62  VARCHAR(10)  --翻身体位
      ,VCF63  VARCHAR(10)  --腹围
      ,VCF64  VARCHAR(10)  --腹围单位
      ,VCF65  VARCHAR(128)  --其他1
      ,VCF66  VARCHAR(128)  --其他2
      ,VCF67  VARCHAR(128)  --其他3
      ,VCF68  VARCHAR(128)  --其他4
      ,VCF69  VARCHAR(128)  --其他5
      ,VCF70  VARCHAR(128)  --其他6
      ,VCF71  VARCHAR(128)  --其他7
      ,VCF72  VARCHAR(128)  --其他8
      ,VCF73  VARCHAR(128)  --其他9
      ,VCF74  VARCHAR(128)  --其他10
      ,VCF75  VARCHAR(128)  --其他11
      ,VCF76  VARCHAR(128)  --其他12
      ,VCF77  VARCHAR(128)  --其他13
      ,VCF78  VARCHAR(128)  --其他14
      ,VCF79  VARCHAR(128)  --其他15
      ,VCF80  VARCHAR(128)  --其他16
      ,VCF81  VARCHAR(128)  --其他17
      ,VCF82  VARCHAR(128)  --其他18
      ,VCF83  VARCHAR(128)  --其他19
      ,VCF84  VARCHAR(128)  --其他20
      ,VAQ01  INT  --病人病历ID
      ,VCF86  TINYINT  --时区 1：3时   2：7时以此类推
      ,VCF87  VARCHAR(10)  --疼痛强度，0~10
)'
, @level0type = 'SCHEMA', @level0name = N'dbo'
, @level1type = 'TABLE', @level1name = N'VCF1'
GO

-- ----------------------------
-- Indexes structure for table sysdiagrams
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table sysdiagrams
-- ----------------------------
ALTER TABLE [dbo].[sysdiagrams] ADD PRIMARY KEY ([diagram_id])
GO

-- ----------------------------
-- Uniques structure for table sysdiagrams
-- ----------------------------
ALTER TABLE [dbo].[sysdiagrams] ADD UNIQUE ([principal_id] ASC, [name] ASC)
GO
