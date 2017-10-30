/*
Navicat MySQL Data Transfer

Source Server         : 京东mysql
Source Server Version : 50717
Source Host           : 116.196.82.249:3306
Source Database       : nursing

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2017-10-10 14:21:49
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for Access
-- ----------------------------
DROP TABLE IF EXISTS `Access`;
CREATE TABLE `Access` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `NurseId` varchar(255) NOT NULL DEFAULT '' COMMENT '护士id',
  `NurseName` varchar(255) DEFAULT NULL,
  `PatientId` varchar(255) NOT NULL DEFAULT '' COMMENT '病人id',
  `PatientName` varchar(255) DEFAULT NULL,
  `ClassId` varchar(255) DEFAULT NULL,
  `BedId` varchar(255) DEFAULT NULL,
  `date_time` datetime DEFAULT NULL COMMENT '测量时间',
  `AccessType` int(11) NOT NULL COMMENT '出入类型： 1=回室，2=外出, 4=全部',
  `AccessReason` int(11) NOT NULL COMMENT '外出原因：1=检查，2=手术，4=其他',
  `AccessTime` varchar(255) NOT NULL DEFAULT '' COMMENT '提醒时间',
  `DateTime` varchar(255) NOT NULL DEFAULT '' COMMENT '测量时间',
  `PatientIdCopy` varchar(255) NOT NULL DEFAULT '' COMMENT '病人id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='出入管理';

-- ----------------------------
-- Records of Access
-- ----------------------------
INSERT INTO `Access` VALUES ('1', '123', null, '56789', null, '123', null, '2017-09-14 03:33:13', '3', '1', '', '', '');
INSERT INTO `Access` VALUES ('2', '123456', null, '56789', null, '123', null, '2017-09-14 03:33:23', '1', '1', '', '', '');
INSERT INTO `Access` VALUES ('3', '123', null, '56789', null, '123', null, '2017-09-14 03:39:50', '1', '4', '', '', '');
INSERT INTO `Access` VALUES ('4', '123', null, '123', null, '123', null, null, '1', '1', '2017-09-13 18:00:00', '', '');
INSERT INTO `Access` VALUES ('5', '123', '王小二', '123', '张三', '123', 'T2', null, '1', '4', '2017-10-09 18:00:00', '', '');

-- ----------------------------
-- Table structure for advice_log
-- ----------------------------
DROP TABLE IF EXISTS `advice_log`;
CREATE TABLE `advice_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `datetime` date DEFAULT NULL COMMENT '时间',
  `state` varchar(22) DEFAULT NULL COMMENT '状态',
  `advice_id` int(11) DEFAULT NULL COMMENT '医嘱ID',
  `nursing_name` varchar(22) DEFAULT NULL COMMENT '护士名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='医嘱执行状态记录表';

-- ----------------------------
-- Records of advice_log
-- ----------------------------

-- ----------------------------
-- Table structure for Breathe
-- ----------------------------
DROP TABLE IF EXISTS `Breathe`;
CREATE TABLE `Breathe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '呼吸值',
  `whethertbm` tinyint(1) NOT NULL COMMENT '是否上呼吸机',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='呼吸数据表';

-- ----------------------------
-- Records of Breathe
-- ----------------------------

-- ----------------------------
-- Table structure for CBM2
-- ----------------------------
DROP TABLE IF EXISTS `CBM2`;
CREATE TABLE `CBM2` (
  `01` int(4) DEFAULT NULL COMMENT 'ID',
  `02` int(4) DEFAULT NULL COMMENT '描述:  病人ID, 关联字段:VAA1.VAA01',
  `03` int(4) DEFAULT NULL COMMENT '描述:  就诊ID, 关联字段:VAA1.VAA07',
  `04` int(4) DEFAULT NULL COMMENT '描述:  婴儿ID, 关联字段:VAP1.VAP01',
  `05` int(4) DEFAULT NULL COMMENT '1:门诊、2:住院, 关联字段:ACF1.ACF01',
  `06` int(4) DEFAULT NULL COMMENT '性质: 0:不区分中、西药处方  1=西、成药, 2=中药, 3=手术记录, 4 = 成药,5=卫材, 9=其他',
  `07` int(4) DEFAULT NULL COMMENT '描述:  1=普通, 2=急诊, 3=儿童, 4=麻醉, 5=第一类精神药品, 6=第二类精神药品, 7=放射药品,8=毒性药品,9=检查,10=检验,11=手术,12=治疗,99=其它',
  `08` varchar(20) DEFAULT NULL COMMENT '描述:  单据号',
  `09` varchar(1024) DEFAULT NULL COMMENT '描述:  摘要',
  `10` int(4) DEFAULT NULL COMMENT '描述:  病区ID, 关联字段:BCK1.BCK01',
  `11` int(4) DEFAULT NULL COMMENT '描述:  病人科室ID, 关联字段:BCK1.BCK01',
  `12` int(4) DEFAULT NULL COMMENT '描述:  开单科室ID, 关联字段:BCK1.BCK01',
  `13` int(4) DEFAULT NULL COMMENT '描述:  开单人ID, 关联字段:BCE1.BCE01',
  `14` varchar(20) DEFAULT NULL COMMENT '描述:  操作员, 关联字段:BCE1.BCE03',
  `15` datetime(6) DEFAULT NULL COMMENT '描述:  开单时间',
  `16` datetime(6) DEFAULT NULL COMMENT '描述:  交易时间',
  `17` datetime(6) DEFAULT NULL COMMENT '描述:  撤销时间',
  `18` varchar(255) DEFAULT NULL COMMENT '描述:  备注',
  `19` tinyint(1) DEFAULT NULL COMMENT '描述:  状态',
  `20` int(4) DEFAULT NULL COMMENT '描述:  收费小票打印次数',
  `21` int(4) DEFAULT NULL COMMENT '描述:  处方打印次数',
  `22` varchar(22) DEFAULT NULL COMMENT '描述:  处方分类, 关联字段:BJW1.BJW02',
  `23` int(4) DEFAULT NULL COMMENT '描述:  复诊预约id, 关联字段:VCY1.VCY01',
  `24` int(4) DEFAULT NULL COMMENT '描述:  诊疗申请单id, 关联字段:VBQ1.VBQ01',
  `25` varchar(10) DEFAULT NULL COMMENT '描述:  医疗类别编码, 关联字段:IAI1.IAI03'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of CBM2
-- ----------------------------

-- ----------------------------
-- Table structure for Devices
-- ----------------------------
DROP TABLE IF EXISTS `Devices`;
CREATE TABLE `Devices` (
  `devicesclass` int(16) unsigned NOT NULL COMMENT '套餐的科室信息',
  `devicesname` varchar(255) NOT NULL COMMENT '套餐的名字',
  `Devicelist` varchar(255) NOT NULL COMMENT '套餐的设备',
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='套餐数据表';

-- ----------------------------
-- Records of Devices
-- ----------------------------
INSERT INTO `Devices` VALUES ('102', 'uiokjji', '[{\"name\":\"SpO2013240\",\"address\":\"8C:DE:52:B2:4B:9A\",\"devicename\":\"uiokjji\"},{\"name\":\"FitThm_X03_1wu\",\"address\":\"DC:0D:30:00:0C:79\",\"devicename\":\"uiokjji\"}]', '12');

-- ----------------------------
-- Table structure for doc_nursing
-- ----------------------------
DROP TABLE IF EXISTS `doc_nursing`;
CREATE TABLE `doc_nursing` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `patientid` int(11) DEFAULT NULL COMMENT '病人ID',
  `temperature` float DEFAULT NULL COMMENT '体温 摄氏度',
  `pulse` int(11) DEFAULT '0' COMMENT '脉搏  次/分',
  `breathing` int(11) DEFAULT '0' COMMENT '呼吸 次/分',
  `spo2h` int(11) DEFAULT NULL COMMENT '血氧 %',
  `pressure` varchar(64) DEFAULT NULL COMMENT ' 血压 x/y mmhg',
  `consciousness` varchar(32) DEFAULT NULL COMMENT '意识',
  `contraction` varchar(32) DEFAULT NULL COMMENT '宫缩',
  `vulnus` varchar(32) DEFAULT NULL COMMENT '伤口',
  `lyma` varchar(32) DEFAULT NULL COMMENT '恶露',
  `milk` varchar(32) DEFAULT NULL COMMENT '乳汁',
  `fundus_h` varchar(32) DEFAULT NULL COMMENT '宫底高度',
  `input` varchar(32) DEFAULT NULL COMMENT '入量 ml',
  `output` varchar(32) DEFAULT NULL COMMENT '出量 ml',
  `desc` text COMMENT '描述',
  `who` varchar(20) DEFAULT NULL COMMENT '责任人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='护理记录表';

-- ----------------------------
-- Records of doc_nursing
-- ----------------------------

-- ----------------------------
-- Table structure for Glucose
-- ----------------------------
DROP TABLE IF EXISTS `Glucose`;
CREATE TABLE `Glucose` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` float NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='血糖数据表';

-- ----------------------------
-- Records of Glucose
-- ----------------------------

-- ----------------------------
-- Table structure for Heartrate
-- ----------------------------
DROP TABLE IF EXISTS `Heartrate`;
CREATE TABLE `Heartrate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='心率数据表';

-- ----------------------------
-- Records of Heartrate
-- ----------------------------

-- ----------------------------
-- Table structure for Height
-- ----------------------------
DROP TABLE IF EXISTS `Height`;
CREATE TABLE `Height` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` double NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='身高数据表';

-- ----------------------------
-- Records of Height
-- ----------------------------

-- ----------------------------
-- Table structure for IntakeOutput
-- ----------------------------
DROP TABLE IF EXISTS `IntakeOutput`;
CREATE TABLE `IntakeOutput` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '这条数据创建的时间',
  `type` int(11) NOT NULL COMMENT '出入量类型，1：入量，2：出量',
  `subtype` int(11) NOT NULL COMMENT '出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次',
  `recordtime` varchar(255) NOT NULL COMMENT '采集时间',
  `value` int(11) NOT NULL COMMENT '采集值',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `nursename` varchar(255) NOT NULL COMMENT '护士姓名',
  `operationtype` int(11) NOT NULL COMMENT '操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便',
  `nurseid` int(4) DEFAULT NULL COMMENT '护士ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8 COMMENT='出入量管理';

-- ----------------------------
-- Records of IntakeOutput
-- ----------------------------
INSERT INTO `IntakeOutput` VALUES ('30', '121', '2017-10-10 10:04:19', '2', '3', '2017-11-11 10:02:00', '2', '12', '见康云专用账号', '1', '10010');
INSERT INTO `IntakeOutput` VALUES ('31', '121', '2017-10-10 10:14:06', '1', '2', '2017-09-10 10:13:00', '236', 'Hugh', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('32', '122', '2017-10-10 10:30:17', '2', '3', '2017-10-10 10:42:00', '2', 'hdhdhd', '见康云专用账号', '2', '10010');
INSERT INTO `IntakeOutput` VALUES ('33', '122', '2017-10-10 10:30:37', '1', '2', '2017-10-10 10:23:00', '555', 'hhh', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('34', '121', '2017-10-10 10:34:10', '2', '1', '2017-09-22 15:45:00', '222', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('35', '121', '2017-10-10 10:34:16', '2', '1', '2017-09-22 15:35:00', '222', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('36', '121', '2017-10-10 10:34:20', '2', '1', '2017-10-22 15:35:00', '222', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('37', '121', '2017-10-10 10:34:34', '2', '1', '2017-10-09 15:35:00', '222', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('38', '121', '2017-10-10 10:34:38', '2', '1', '2017-10-19 15:35:00', '222', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('39', '121', '2017-10-10 10:35:57', '2', '1', '2017-10-10 05:45:00', '1999', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('40', '121', '2017-10-10 10:36:05', '2', '1', '2017-10-11 05:45:00', '1999', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('41', '121', '2017-10-10 10:36:10', '2', '1', '2017-10-13 05:45:00', '1999', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('42', '121', '2017-10-10 10:36:17', '2', '1', '2017-09-13 05:45:00', '1999', '描述', '见康云专用账号', '0', '10010');
INSERT INTO `IntakeOutput` VALUES ('43', '121', '2017-10-10 10:36:28', '2', '1', '2017-10-10 15:25:00', '1999', '描述', '见康云专用账号', '0', '10010');

-- ----------------------------
-- Table structure for Pressure
-- ----------------------------
DROP TABLE IF EXISTS `Pressure`;
CREATE TABLE `Pressure` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `diavalue` int(11) NOT NULL COMMENT '低压值',
  `sysvalue` int(11) NOT NULL COMMENT '高压值',
  `pulsevalue` int(11) NOT NULL COMMENT '脉率值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='血压数据表';

-- ----------------------------
-- Records of Pressure
-- ----------------------------

-- ----------------------------
-- Table structure for Pulse
-- ----------------------------
DROP TABLE IF EXISTS `Pulse`;
CREATE TABLE `Pulse` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '脉搏值',
  `whetherbriefness` tinyint(1) NOT NULL COMMENT '是否短促',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='脉搏测量数据表';

-- ----------------------------
-- Records of Pulse
-- ----------------------------

-- ----------------------------
-- Table structure for Spo2h
-- ----------------------------
DROP TABLE IF EXISTS `Spo2h`;
CREATE TABLE `Spo2h` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='血氧数据表';

-- ----------------------------
-- Records of Spo2h
-- ----------------------------

-- ----------------------------
-- Table structure for Temperature
-- ----------------------------
DROP TABLE IF EXISTS `Temperature`;
CREATE TABLE `Temperature` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `ttemptype` int(11) NOT NULL COMMENT '体温的类型',
  `coolingtype` int(11) NOT NULL COMMENT '降温的类型',
  `value` float NOT NULL COMMENT '体温值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='体温测量数据表';

-- ----------------------------
-- Records of Temperature
-- ----------------------------
INSERT INTO `Temperature` VALUES ('1', '45', '45', '2006-01-02 15:04:05', '1', '1', '0', '45');
INSERT INTO `Temperature` VALUES ('2', '45', '45', '2006-01-02 15:04:05', '1', '1', '0', '45');
INSERT INTO `Temperature` VALUES ('3', '45', '45', '2006-01-02 15:04:05', '1', '1', '0', '45');

-- ----------------------------
-- Table structure for VAF1
-- ----------------------------
DROP TABLE IF EXISTS `VAF1`;
CREATE TABLE `VAF1` (
  `VAF01` int(11) DEFAULT NULL,
  `VAF01A` int(11) DEFAULT NULL,
  `VAF01B` int(11) DEFAULT NULL,
  `VAF04` tinyint(4) DEFAULT NULL,
  `VAA01` int(11) DEFAULT NULL,
  `VAF06` int(11) DEFAULT NULL,
  `VAF07` int(11) DEFAULT NULL,
  `BCK01A` int(11) DEFAULT NULL,
  `ROWNR` int(11) DEFAULT NULL,
  `VAF10` int(11) DEFAULT NULL,
  `VAF11` tinyint(4) DEFAULT NULL,
  `BDA01` varchar(2) DEFAULT NULL,
  `BBX01` int(11) DEFAULT NULL,
  `VAF14` varchar(60) DEFAULT NULL,
  `VAF15` varchar(30) DEFAULT NULL,
  `BBY01` int(11) DEFAULT NULL,
  `VAF17` int(11) DEFAULT NULL,
  `VAF18` decimal(18,4) DEFAULT NULL,
  `VAF19` varchar(10) DEFAULT NULL,
  `VAF20` decimal(18,4) DEFAULT NULL,
  `VAF21` decimal(18,4) DEFAULT NULL,
  `VAF22` text,
  `VAF23` varchar(128) DEFAULT NULL,
  `BCK01B` int(11) DEFAULT NULL,
  `VAF25` varchar(10) DEFAULT NULL,
  `VAF26` varchar(20) DEFAULT NULL,
  `VAF27` int(11) DEFAULT NULL,
  `VAF28` tinyint(4) DEFAULT NULL,
  `VAF29` varchar(4) DEFAULT NULL,
  `VAF30` varchar(64) DEFAULT NULL,
  `VAF31` tinyint(4) DEFAULT NULL,
  `VAF32` tinyint(4) DEFAULT NULL,
  `VAF33` tinyint(4) DEFAULT NULL,
  `VAF34` tinyint(4) DEFAULT NULL,
  `VAF35` tinyint(4) DEFAULT NULL,
  `VAF36` datetime DEFAULT NULL,
  `VAF37` datetime DEFAULT NULL,
  `VAF38` datetime DEFAULT NULL,
  `BCK01C` int(11) DEFAULT NULL,
  `BCE02A` varchar(20) DEFAULT NULL,
  `BCE03A` varchar(20) DEFAULT NULL,
  `VAF42` datetime DEFAULT NULL,
  `BCE03B` varchar(20) DEFAULT NULL,
  `BCE03C` varchar(20) DEFAULT NULL,
  `VAF45` datetime DEFAULT NULL,
  `BCE03D` varchar(20) DEFAULT NULL,
  `VAF47` datetime DEFAULT NULL,
  `BCE03E` varchar(20) DEFAULT NULL,
  `BCE03F` varchar(20) DEFAULT NULL,
  `VAF50` datetime DEFAULT NULL,
  `VAF51` int(11) DEFAULT NULL,
  `VAF52` tinyint(4) DEFAULT NULL,
  `VAF53` int(11) DEFAULT NULL,
  `VAF54` tinyint(4) DEFAULT NULL,
  `VAF55` text,
  `CBM01` int(11) DEFAULT NULL,
  `BCK01D` int(11) DEFAULT NULL,
  `VAF58` tinyint(4) DEFAULT NULL,
  `VAF59` int(11) DEFAULT NULL,
  `VAF60` varchar(10) DEFAULT NULL,
  `VAF61` decimal(8,2) DEFAULT NULL,
  `VAF62` decimal(8,2) DEFAULT NULL,
  `BCE01A` int(11) DEFAULT NULL,
  `BCE01B` int(11) DEFAULT NULL,
  `BCE01C` int(11) DEFAULT NULL,
  `BCE01D` int(11) DEFAULT NULL,
  `BCE01E` int(11) DEFAULT NULL,
  `BCE01F` int(11) DEFAULT NULL,
  `BCE01G` int(11) DEFAULT NULL,
  `BCE03G` varchar(20) DEFAULT NULL,
  `VAF71` datetime DEFAULT NULL,
  `DSK01` int(11) DEFAULT NULL,
  `VAF01C` int(11) DEFAULT NULL,
  `VAF74` datetime DEFAULT NULL,
  `VAF75` tinyint(4) DEFAULT NULL,
  `BCE01H` int(11) DEFAULT NULL,
  `BCE03H` varchar(20) DEFAULT NULL,
  `BIW02` varchar(64) DEFAULT NULL,
  `Crypt` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='CREATE TABLE VAF1--住院病人医嘱记录\r\n(\r\n        VAF01  INT  --ID\r\n      ,VAF01A INT  --相关ID, 关联字段：VAF1.VAF01\r\n      ,VAF01B INT  --前提ID, 关联字段：VAF1.VAF01\r\n      ,VAF04  TINYINT  --1：门诊;2：住院\r\n      ,VAA01  INT  --病人ID, 关联字段：VAA1.VAA01\r\n      ,VAF06  INT  --就诊ID, 主页ID\r\n      ,VAF07  INT  --婴儿ID\r\n      ,BCK01A INT  --病人科室ID, 关联字段：BCK1.BCK01\r\n      ,ROWNR  INT  --次序\r\n      ,VAF10  TINYINT  --1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果\r\n      ,VAF11  TINYINT  --医嘱类型, 1=长期医嘱, 2=临时医嘱\r\n      ,BDA01  VARCHAR(2)  --诊疗类型, 关联字段：BDA1.BDA01\r\n      ,BBX01  INT  --诊疗项目ID, 关联字段：BBX1.BBX01\r\n      ,VAF14  VARCHAR(60)  --标本部位\r\n      ,VAF15  VARCHAR(30)  --检查方法\r\n      ,BBY01  INT  --收费项目ID, 关联字段：BBY1.BBY01\r\n      ,VAF17  INT  --天数, day number\r\n      ,VAF18  NUMERIC(18, 4)  --剂量, 单次用量\r\n      ,VAF19  VARCHAR(10)  --用量\r\n      ,VAF20  NUMERIC(18, 4)  --单量\r\n      ,VAF21  NUMERIC(18, 4)  --数量\r\n      ,VAF22  VARCHAR(1024)  --医嘱\r\n      ,VAF23  VARCHAR(128)  --医师嘱托\r\n      ,BCK01B INT  --执行科室ID, 关联字段：BCK1.BCK01\r\n      ,VAF25  VARCHAR(10)  --空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果\r\n      ,VAF26  VARCHAR(20)  --执行频次\r\n      ,VAF27  INT  --频率次数\r\n      ,VAF28  TINYINT  --频率间隔\r\n      ,VAF29  VARCHAR(4)  --间隔单位\r\n      ,VAF30  VARCHAR(64)  --执行时间方案\r\n      ,VAF31  TINYINT  --计价特性  0=正常, 1=自费\r\n      ,VAF32  TINYINT  --0：正常; 1＝给药途径\r\n      ,VAF33  TINYINT  --0：标记未用;1：正常 2：自动停止\r\n      ,VAF34  TINYINT  --可否分零\r\n      ,VAF35  TINYINT  --0：正常 1：紧急\r\n      ,VAF36  DATETIME  --开始执行时间\r\n      ,VAF37  DATETIME  --执行终止时间\r\n      ,VAF38  DATETIME  --上次执行时间\r\n      ,BCK01C INT  --开嘱科室ID, 关联字段：BCK1.BCK01\r\n      ,BCE02A VARCHAR(20)  --医师编码, 关联字段：BCE1.BCE02\r\n      ,BCE03A VARCHAR(20)  --开嘱医师, 关联字段：BCE1.BCE03\r\n      ,VAF42  DATETIME  --开嘱时间\r\n      ,BCE03B VARCHAR(20)  --开嘱护士, 关联字段：BCE1.BCE03\r\n      ,BCE03C VARCHAR(20)  --校对护士, 关联字段：BCE1.BCE03\r\n      ,VAF45  DATETIME  --校对时间\r\n      ,BCE03D VARCHAR(20)  --停嘱医生, 关联字段：BCE1.BCE03\r\n      ,VAF47  DATETIME  --停嘱时间\r\n      ,BCE03E VARCHAR(20)  --停嘱护士, 关联字段：BCE1.BCE03\r\n      ,BCE03F VARCHAR(20)  --停嘱校对护士, 关联字段：BCE1.BCE03\r\n    ';

-- ----------------------------
-- Records of VAF1
-- ----------------------------
INSERT INTO `VAF1` VALUES ('1711129', '1711128', '10161', '1', '97027', '183103', '0', '26', '1', '1', '2', '1', '6576', '0.9%氯化钠注射液', '', '8877', '1', '250.0000', '250 ml', '250.0000', '1.0000', '0.9%氯化钠注射液 250ml', '', '28', ' ', '每日1次', '1', '1', 'D', '', '0', '0', '0', '0', '0', '2012-03-25 04:28:00', '2012-03-25 04:29:00', '1900-01-01 00:00:00', '26', '9700', '李小里', '2012-03-25 04:29:00', '', ' ', '1900-01-01 00:00:00', '', '1900-01-01 00:00:00', ' ', ' ', '1900-01-01 00:00:00', '0', '0', '-1', '0', '', '582829', '111', '0', '1', '0', '0.00', '0.00', '356', '0', '0', '0', '0', '0', '356', '李小里', '2012-03-25 04:29:32', '0', '0', '2012-03-25 04:28:16', '0', '0', ' ', ' ', '0');
INSERT INTO `VAF1` VALUES ('1711130', '1711128', '10161', '1', '97027', '183103', '0', '26', '2', '1', '2', '1', '6464', '盐酸克林霉素注射液', '', '6464', '1', '0.9000', '0.9 g', '0.9000', '3.0000', '盐酸克林霉素注射液 0.3g:2ml', '', '28', ' ', '每日1次', '1', '1', 'D', '', '0', '0', '0', '0', '0', '2012-03-25 04:28:00', '2012-03-25 04:29:00', '1900-01-01 00:00:00', '26', '9700', '李小里', '2012-03-25 04:29:00', '', ' ', '1900-01-01 00:00:00', '', '1900-01-01 00:00:00', ' ', ' ', '1900-01-01 00:00:00', '0', '0', '-1', '0', '', '582829', '111', '0', '1', '0', '0.00', '0.00', '356', '0', '0', '0', '0', '0', '356', '李小里', '2012-03-25 04:29:32', '0', '0', '2012-03-25 04:28:16', '0', '0', ' ', ' ', '0');
INSERT INTO `VAF1` VALUES ('1711131', '0', '11123', '1', '97027', '183103', '0', '26', '0', '1', '2', 'T', '7544', '', '', '0', '1', '0.0000', '', '0.0000', '1.0000', '续滴(不加药)', '', '111', ' ', '每日1次', '1', '1', 'D', '', '0', '1', '0', '0', '0', '2012-03-25 04:28:00', '2012-03-25 04:29:00', '1900-01-01 00:00:00', '26', '9700', '李小里', '2012-03-25 04:29:00', '', ' ', '1900-01-01 00:00:00', '', '1900-01-01 00:00:00', ' ', ' ', '1900-01-01 00:00:00', '0', '0', '-1', '0', '', '582829', '28', '0', '2', '0', '0.00', '0.00', '356', '0', '0', '0', '0', '0', '356', '李小里', '2012-03-25 04:29:32', '0', '0', '2012-03-25 04:28:16', '0', '0', ' ', ' ', '0');
INSERT INTO `VAF1` VALUES ('1711132', '1711131', '11123', '1', '97027', '183103', '0', '26', '1', '1', '2', '1', '6475', '甲硝唑氯化钠注射液', '', '6475', '1', '1.0000', '1.0 g', '1.0000', '2.0000', '甲硝唑氯化钠注射液 0.5g:100ml', '', '28', ' ', '每日1次', '1', '1', 'D', '', '0', '0', '0', '0', '0', '2012-03-25 04:28:00', '2012-03-25 04:29:00', '1900-01-01 00:00:00', '26', '9700', '李小里', '2012-03-25 04:29:00', '', ' ', '1900-01-01 00:00:00', '', '1900-01-01 00:00:00', ' ', ' ', '1900-01-01 00:00:00', '0', '0', '-1', '0', '', '582829', '111', '0', '2', '0', '0.00', '0.00', '356', '0', '0', '0', '0', '0', '356', '李小里', '2012-03-25 04:29:32', '0', '0', '2012-03-25 04:28:16', '0', '0', ' ', ' ', '0');
INSERT INTO `VAF1` VALUES ('1711133', '0', '10693', '1', '97027', '183103', '0', '26', '0', '1', '2', 'T', '7535', '', '', '0', '3', '0.0000', '', '0.0000', '9.0000', '口服', '', '26', ' ', '每日3次', '3', '1', 'D', '', '0', '1', '0', '0', '0', '2012-03-25 04:28:00', '2012-03-25 04:29:00', '1900-01-01 00:00:00', '26', '9700', '李小里', '2012-03-25 04:29:00', '', ' ', '1900-01-01 00:00:00', '', '1900-01-01 00:00:00', ' ', ' ', '1900-01-01 00:00:00', '0', '0', '-1', '0', '', '582829', '28', '0', '3', '0', '0.00', '0.00', '356', '0', '0', '0', '0', '0', '356', '李小里', '2012-03-25 04:29:32', '0', '0', '2012-03-25 04:28:16', '0', '0', ' ', ' ', '0');

-- ----------------------------
-- Table structure for Warn
-- ----------------------------
DROP TABLE IF EXISTS `Warn`;
CREATE TABLE `Warn` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `NurseId` varchar(255) NOT NULL DEFAULT '' COMMENT '护士id',
  `PatientId` varchar(255) DEFAULT NULL COMMENT '病人id',
  `ClassId` varchar(255) NOT NULL DEFAULT '' COMMENT '科室id',
  `date_time` datetime DEFAULT NULL,
  `Name` varchar(255) DEFAULT NULL COMMENT '提醒名称',
  `WarnTime` varchar(255) NOT NULL DEFAULT '' COMMENT '提醒时间',
  `BedId` varchar(255) DEFAULT NULL,
  `Desc` varchar(255) DEFAULT NULL,
  `WarnType` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='提醒表';

-- ----------------------------
-- Records of Warn
-- ----------------------------
INSERT INTO `Warn` VALUES ('1', '123456', '123', '5678', '2017-09-14 07:28:43', '浩浩', '2017-09-12 22:00:00', null, null, null);
INSERT INTO `Warn` VALUES ('5', '', '', '123', null, '量体温', '2017-10-09 08:00:00', null, '整个科室', '1');
INSERT INTO `Warn` VALUES ('6', '', '', '123', null, '量体温', '2017-10-09 08:00', null, '整个科室', '1');
INSERT INTO `Warn` VALUES ('7', '', '', '123', null, '量体温', '2017-10-09 08:00', null, '整个科室', '1');
INSERT INTO `Warn` VALUES ('8', '', '', '123', null, '量体温', '2017-10-09 08:01', null, '整个科室', '1');
INSERT INTO `Warn` VALUES ('9', '123', '123', '123', null, '量体温', '2017-09-13 22:00:00', null, null, null);
INSERT INTO `Warn` VALUES ('11', '', '', '123', null, '量体温', '2017-10-09 08:00', null, '整个科室', '1');
INSERT INTO `Warn` VALUES ('12', '', '', '123', null, '量体温', '2017-10-09 08:00', null, '整个科室', '1');
INSERT INTO `Warn` VALUES ('13', '123', '123', '123', null, '量体温', '2017-09-13 22:00:00', null, null, null);

-- ----------------------------
-- Table structure for Weight
-- ----------------------------
DROP TABLE IF EXISTS `Weight`;
CREATE TABLE `Weight` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` float NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='体重数据表';

-- ----------------------------
-- Records of Weight
-- ----------------------------
INSERT INTO `Weight` VALUES ('1', '111', '101', '2006-01-02 23:15:15', '1', '20');
INSERT INTO `Weight` VALUES ('2', '111', '101', '2005-01-05 23:04:05', '1', '20');
INSERT INTO `Weight` VALUES ('3', '111', '101', '2004-01-05 23:04:05', '1', '20');
INSERT INTO `Weight` VALUES ('4', '111', '101', '2004-01-05 15:04:05', '1', '20');
INSERT INTO `Weight` VALUES ('5', '111', '101', '2004-01-05 15:04:05', '1', '20');
INSERT INTO `Weight` VALUES ('6', '111', '101', '2004-01-05 15:04:05', '1', '20');
INSERT INTO `Weight` VALUES ('7', '111', '101', '2004-01-05 15:04:05', '1', '20');

-- ----------------------------
-- Table structure for x
-- ----------------------------
DROP TABLE IF EXISTS `x`;
CREATE TABLE `x` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `datatime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of x
-- ----------------------------
INSERT INTO `x` VALUES ('1', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('2', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('3', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('4', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('5', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('6', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('7', '胡杨', '2019-06-15 08:37:18');
INSERT INTO `x` VALUES ('8', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('9', '胡杨', '2014-06-15 08:37:18');
INSERT INTO `x` VALUES ('10', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('11', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('12', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('13', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('14', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('15', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('16', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('17', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('18', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('19', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('20', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('21', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('22', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('23', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('24', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('25', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('26', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('27', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('28', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('29', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('30', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('31', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('32', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('33', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('34', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('35', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('36', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('37', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('38', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('39', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('40', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('41', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('42', '胡杨', '2020-10-10 12:12:12');
INSERT INTO `x` VALUES ('43', '胡杨', '2020-10-10 12:12:12');
