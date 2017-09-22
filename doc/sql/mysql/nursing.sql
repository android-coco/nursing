/*
Navicat MySQL Data Transfer

Source Server         : 京东mysql
Source Server Version : 50717
Source Host           : 116.196.82.249:3306
Source Database       : nursing

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2017-09-19 09:59:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `access_time` datetime NOT NULL COMMENT '出入时间',
  `access_type` int(11) NOT NULL,
  `access_reason` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for base_model
-- ----------------------------
DROP TABLE IF EXISTS `base_model`;
CREATE TABLE `base_model` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for breathe
-- ----------------------------
DROP TABLE IF EXISTS `breathe`;
CREATE TABLE `breathe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '值',
  `whethertbm` tinyint(1) NOT NULL COMMENT '是否上呼吸机',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for glucose
-- ----------------------------
DROP TABLE IF EXISTS `glucose`;
CREATE TABLE `glucose` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` float NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for heartrate
-- ----------------------------
DROP TABLE IF EXISTS `heartrate`;
CREATE TABLE `heartrate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for height
-- ----------------------------
DROP TABLE IF EXISTS `height`;
CREATE TABLE `height` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` double NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for intake_output
-- ----------------------------
DROP TABLE IF EXISTS `intake_output`;
CREATE TABLE `intake_output` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `type` int(11) NOT NULL COMMENT '出入量类型，1：入量，2：出量',
  `subtype` int(11) NOT NULL COMMENT '出入量的子类型，1：其他入量/ml，2：输液入量/ml，3：饮食入量/ml，101：其他出量/ml，102：排尿出量/ml，103：大便出量/次',
  `record_time` varchar(255) NOT NULL COMMENT '采集时间',
  `value` int(11) NOT NULL COMMENT '采集值',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `nurse_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pressure
-- ----------------------------
DROP TABLE IF EXISTS `pressure`;
CREATE TABLE `pressure` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `diavalue` int(11) NOT NULL COMMENT '低压值',
  `sysvalue` int(11) NOT NULL COMMENT '高压值',
  `pulsevalue` int(11) NOT NULL COMMENT '脉率值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pulse
-- ----------------------------
DROP TABLE IF EXISTS `pulse`;
CREATE TABLE `pulse` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '值',
  `whetherbriefness` tinyint(1) NOT NULL COMMENT '是否短促',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for spo2h
-- ----------------------------
DROP TABLE IF EXISTS `spo2h`;
CREATE TABLE `spo2h` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` int(11) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for temperature
-- ----------------------------
DROP TABLE IF EXISTS `temperature`;
CREATE TABLE `temperature` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `ttemptype` int(11) NOT NULL COMMENT '体温的类型',
  `coolingtype` int(11) NOT NULL COMMENT '降温的类型',
  `value` float NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for VAF2
-- ----------------------------
DROP TABLE IF EXISTS `VAF2`;
CREATE TABLE `VAF2` (
  `VAF01` int(11) DEFAULT NULL COMMENT 'ID',
  `VAF01A` int(11) DEFAULT NULL COMMENT '相关ID, 关联字段：VAF1.VAF01',
  `VAF01B` int(11) DEFAULT NULL COMMENT '前提ID, 关联字段：VAF1.VAF01',
  `VAF04` tinyint(4) DEFAULT NULL COMMENT '1：门诊;2：住院',
  `VAA01` int(11) DEFAULT NULL COMMENT '病人ID, 关联字段：VAA1.VAA01',
  `VAF06` int(11) DEFAULT NULL COMMENT '就诊ID, 主页ID',
  `VAF07` int(11) DEFAULT NULL COMMENT '婴儿ID',
  `BCK01A` int(11) DEFAULT NULL COMMENT '病人科室ID, 关联字段：BCK1.BCK01',
  `ROWNR` int(11) DEFAULT NULL COMMENT '次序',
  `VAF10` tinyint(4) DEFAULT NULL COMMENT '1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果',
  `VAF11` tinyint(4) DEFAULT NULL COMMENT '医嘱类型, 1=长期医嘱, 2=临时医嘱',
  `BDA01` varchar(2) DEFAULT NULL COMMENT '诊疗类型, 关联字段：BDA1.BDA01',
  `BBX01` int(11) DEFAULT NULL COMMENT '诊疗项目ID, 关联字段：BBX1.BBX01',
  `VAF14` varchar(60) DEFAULT NULL COMMENT '标本部位',
  `VAF15` varchar(30) DEFAULT NULL COMMENT '检查方法',
  `BBY01` int(11) DEFAULT NULL COMMENT '收费项目ID, 关联字段：BBY1.BBY01',
  `VAF17` int(11) DEFAULT NULL COMMENT '天数, day number',
  `VAF18` decimal(18,4) DEFAULT NULL COMMENT '剂量, 单次用量',
  `VAF19` varchar(10) DEFAULT NULL COMMENT '用量',
  `VAF20` decimal(18,4) DEFAULT NULL COMMENT '单量',
  `VAF21` decimal(18,4) DEFAULT NULL COMMENT '数量',
  `VAF22` varchar(1024) DEFAULT NULL COMMENT '医嘱',
  `VAF23` varchar(128) DEFAULT NULL COMMENT '医师嘱托',
  `BCK01B` int(11) DEFAULT NULL COMMENT '执行科室ID, 关联字段：BCK1.BCK01',
  `VAF25` varchar(10) DEFAULT NULL COMMENT '空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果',
  `VAF26` varchar(20) DEFAULT NULL COMMENT '执行频次',
  `VAF27` int(11) DEFAULT NULL COMMENT '频率次数',
  `VAF28` tinyint(4) DEFAULT NULL COMMENT '频率间隔',
  `VAF29` varchar(4) DEFAULT NULL COMMENT '间隔单位',
  `VAF30` varchar(64) DEFAULT NULL COMMENT '执行时间方案',
  `VAF31` tinyint(4) DEFAULT NULL COMMENT '计价特性  0=正常, 1=自费',
  `VAF32` tinyint(4) DEFAULT NULL COMMENT '0：正常; 1＝给药途径',
  `VAF33` tinyint(4) DEFAULT NULL COMMENT '0：标记未用;1：正常 2：自动停止',
  `VAF34` tinyint(4) DEFAULT NULL COMMENT '可否分零',
  `VAF35` tinyint(4) DEFAULT NULL COMMENT '0：正常 1：紧急',
  `VAF36` datetime DEFAULT NULL COMMENT '开始执行时间',
  `VAF37` datetime DEFAULT NULL COMMENT '执行终止时间',
  `VAF38` datetime DEFAULT NULL COMMENT '上次执行时间',
  `BCK01C` int(11) DEFAULT NULL COMMENT '开嘱科室ID, 关联字段：BCK1.BCK01',
  `BCE02A` varchar(20) DEFAULT NULL COMMENT '医师编码, 关联字段：BCE1.BCE02',
  `BCE03A` varchar(20) DEFAULT NULL COMMENT '开嘱医师, 关联字段：BCE1.BCE03',
  `VAF42` datetime DEFAULT NULL COMMENT '开嘱时间',
  `BCE03B` varchar(20) DEFAULT NULL COMMENT '开嘱护士, 关联字段：BCE1.BCE03',
  `BCE03C` varchar(20) DEFAULT NULL COMMENT '校对护士, 关联字段：BCE1.BCE03',
  `VAF45` datetime DEFAULT NULL COMMENT '校对时间',
  `BCE03D` varchar(20) DEFAULT NULL COMMENT '停嘱医生, 关联字段：BCE1.BCE03',
  `VAF47` datetime DEFAULT NULL COMMENT '停嘱时间',
  `BCE03E` varchar(20) DEFAULT NULL COMMENT '停嘱护士, 关联字段：BCE1.BCE03',
  `BCE03F` varchar(20) DEFAULT NULL COMMENT '停嘱校对护士, 关联字段：BCE1.BCE03',
  `VAF50` datetime DEFAULT NULL COMMENT '执行停嘱时间',
  `VAF51` int(11) DEFAULT NULL COMMENT '申请ID',
  `VAF52` tinyint(4) DEFAULT NULL COMMENT '0：新开；1：上传',
  `VAF53` int(11) DEFAULT NULL COMMENT '审查结果，用于药品合理用药审核。(描述性医嘱：执行分类  -1：普通，0：口服单，1：注射单，2：输液单，3：治疗单，4：皮试单，5：输血单，6：护理单，9：其它)',
  `VAF54` tinyint(4) DEFAULT NULL COMMENT '0：否，1：忽略',
  `VAF55` varchar(1024) DEFAULT NULL COMMENT '摘要，医嘱备注',
  `CBM01` int(11) DEFAULT NULL COMMENT '医嘱单id, 关联字段：CBM1.CBM01',
  `BCK01D` int(11) DEFAULT NULL COMMENT '给药科室, 关联字段：BCK1.BCK01',
  `VAF58` tinyint(4) DEFAULT NULL COMMENT '0：正常， 1：自备药，2：离院带药',
  `VAF59` int(11) DEFAULT NULL COMMENT '组号',
  `VAF60` varchar(10) DEFAULT NULL COMMENT '滴速',
  `VAF61` decimal(8,2) DEFAULT NULL COMMENT '首日执行次数',
  `VAF62` decimal(8,2) DEFAULT NULL COMMENT '末日执行次数',
  `BCE01A` int(11) DEFAULT NULL COMMENT '开嘱医师ID, 关联字段：BCE1.BCE01',
  `BCE01B` int(11) DEFAULT NULL COMMENT '开嘱护士ID, 关联字段：BCE1.BCE01',
  `BCE01C` int(11) DEFAULT NULL COMMENT '校对护士ID, 关联字段：BCE1.BCE01',
  `BCE01D` int(11) DEFAULT NULL COMMENT '停嘱医师ID, 关联字段：BCE1.BCE01',
  `BCE01E` int(11) DEFAULT NULL COMMENT '停嘱护士ID, 关联字段：BCE1.BCE01',
  `BCE01F` int(11) DEFAULT NULL COMMENT '停嘱校对护士ID, 关联字段：BCE1.BCE01',
  `BCE01G` int(11) DEFAULT NULL COMMENT '操作员ID, 关联字段：BCE1.BCE01',
  `BCE03G` varchar(20) DEFAULT NULL COMMENT '操作员, 关联字段：BCE1.BCE03',
  `VAF71` datetime DEFAULT NULL COMMENT '审核时间',
  `DSK01` int(11) DEFAULT NULL COMMENT '药品批次id DSK_ID',
  `VAF01C` int(11) DEFAULT NULL COMMENT '原医嘱id  (-1 = 重整医嘱)',
  `VAF74` datetime DEFAULT NULL COMMENT '重整医嘱时间',
  `VAF75` tinyint(4) DEFAULT NULL COMMENT '药品用药标识',
  `BCE01H` int(11) DEFAULT NULL COMMENT '授权医师id, 关联字段：BCE1.BCE01',
  `BCE03H` varchar(20) DEFAULT NULL COMMENT '授权医师, 关联字段：BCE1.BCE03',
  `BIW02` varchar(64) DEFAULT NULL COMMENT '用药目的, 关联字段：BIW1.BIW02'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for warn
-- ----------------------------
DROP TABLE IF EXISTS `warn`;
CREATE TABLE `warn` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL COMMENT '提醒名称',
  `warn_time` datetime NOT NULL COMMENT '提醒时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for weight
-- ----------------------------
DROP TABLE IF EXISTS `weight`;
CREATE TABLE `weight` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurse_id` varchar(255) NOT NULL COMMENT '护士id',
  `patient_id` varchar(255) NOT NULL COMMENT '病人id',
  `date_time` datetime DEFAULT NULL,
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` int(11) NOT NULL COMMENT '测试的场景',
  `value` float NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

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
