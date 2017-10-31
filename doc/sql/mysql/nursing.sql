/*
Navicat MySQL Data Transfer

Source Server         : youhao
Source Server Version : 50710
Source Host           : 127.0.0.1:3306
Source Database       : nursing

Target Server Type    : MYSQL
Target Server Version : 50710
File Encoding         : 65001

Date: 2017-10-30 10:05:48
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `NurseId` varchar(255) NOT NULL DEFAULT '' COMMENT '护士id',
  `NurseName` varchar(255) DEFAULT NULL COMMENT '护士姓名',
  `PatientId` varchar(255) NOT NULL DEFAULT '' COMMENT '病人id',
  `PatientName` varchar(255) DEFAULT NULL,
  `ClassId` varchar(255) DEFAULT NULL COMMENT '科室id',
  `BedId` varchar(255) DEFAULT NULL COMMENT '床号',
  `date_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上传时当前时间',
  `AccessType` int(11) NOT NULL COMMENT '出入类型： 1=回室，2=外出, 4=全部',
  `AccessReason` int(11) NOT NULL COMMENT '外出原因：1=检查，2=手术，4=其他',
  `AccessTime` varchar(255) NOT NULL DEFAULT '' COMMENT '提醒时间',
  `DateTime` varchar(255) NOT NULL DEFAULT '' COMMENT '测量时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8 COMMENT='出入管理';

-- ----------------------------
-- Table structure for ache
-- ----------------------------
DROP TABLE IF EXISTS `ache`;
CREATE TABLE `ache` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(250) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `value` varchar(255) NOT NULL COMMENT '值',
  `recordscene` varchar(255) NOT NULL COMMENT '测试场景',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for advicestate
-- ----------------------------
DROP TABLE IF EXISTS `advicestate`;
CREATE TABLE `advicestate` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `patientid` int(20) NOT NULL COMMENT '病人id',
  `advicestateId` int(20) NOT NULL COMMENT '医嘱id',
  `state` varchar(255) NOT NULL COMMENT '医嘱状态',
  `time` datetime NOT NULL COMMENT '打点时间',
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `period` varchar(255) NOT NULL COMMENT '周期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4;

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
-- Table structure for bce1
-- ----------------------------
DROP TABLE IF EXISTS `bce1`;
CREATE TABLE `bce1` (
  `BCE01` int(11) DEFAULT NULL,
  `BCE02` varchar(20) DEFAULT NULL COMMENT '编码, 工号, Code',
  `BCE03` varchar(20) DEFAULT NULL,
  `BCE04` varchar(20) DEFAULT NULL,
  `BCE05` varchar(10) DEFAULT NULL,
  `BCE06` varchar(20) DEFAULT NULL,
  `ABBRP` varchar(10) DEFAULT NULL,
  `ABBRW` varchar(10) DEFAULT NULL,
  `ABW01` varchar(1) DEFAULT NULL,
  `ACK01` varchar(2) DEFAULT NULL,
  `BCE11` varchar(20) DEFAULT NULL,
  `ABQ01` varchar(4) DEFAULT NULL,
  `BCE13` datetime DEFAULT NULL,
  `BCK01` int(11) DEFAULT NULL,
  `ACP01` varchar(2) DEFAULT NULL,
  `AAY01` varchar(4) DEFAULT NULL,
  `ACT01` varchar(4) DEFAULT NULL,
  `BCE18` varchar(64) DEFAULT NULL,
  `BCE19` datetime DEFAULT NULL,
  `ABU02` varchar(64) DEFAULT NULL,
  `AAD01` varchar(4) DEFAULT NULL,
  `ABG01` varchar(4) DEFAULT NULL,
  `BCE23` varchar(32) DEFAULT NULL,
  `BCE24` varchar(32) DEFAULT NULL,
  `ABS01` varchar(8) DEFAULT NULL,
  `ABI01` varchar(4) DEFAULT NULL,
  `AAH01` varchar(2) DEFAULT NULL,
  `BCE28` datetime DEFAULT NULL,
  `AAQ01` varchar(4) DEFAULT NULL,
  `ABE01` varchar(10) DEFAULT NULL,
  `BCE31` datetime DEFAULT NULL,
  `BCE32` datetime DEFAULT NULL,
  `BCE33` varchar(32) DEFAULT NULL,
  `BCE34` varchar(128) DEFAULT NULL,
  `BCE35` varchar(20) DEFAULT NULL,
  `BCE36` varchar(20) DEFAULT NULL,
  `BCE37` varchar(20) DEFAULT NULL,
  `BCE38` varchar(64) DEFAULT NULL,
  `BCE39` varchar(255) DEFAULT NULL,
  `BCE41` smallint(6) DEFAULT NULL COMMENT '状态，0=试用，1=在职，2=离职，3=退休'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='CREATE TABLE BCE1—员工表, 人员表\r\n(\r\n       BCE01 INT  --ID\r\n      ,BCE02 VARCHAR(20)  --编码, 工号, Code\r\n      ,BCE03 VARCHAR(20)  --姓名, 在不同场合的名称： EMPNN=护士, EMPPN=医师, EMPDN=药剂师, EMPWN=仓管员, EMPTN=医技师,\r\n      ,BCE04 VARCHAR(20)  --英文名, NameB\r\n      ,BCE05 VARCHAR(10)  --称谓, Title\r\n      ,BCE06 VARCHAR(20)  --曾用名, PreviousName\r\n      ,ABBRP VARCHAR(10)  --拼音, PYImc\r\n      ,ABBRW VARCHAR(10)  --五笔, WBImc\r\n      ,ABW01 VARCHAR(1)  --性别, Sex\r\n      ,ACK01 VARCHAR(2)  --婚姻, Marriage, Marriaged\r\n      ,BCE11 VARCHAR(20)  --身份证号, IDNumber\r\n      ,ABQ01 VARCHAR(4)  --民族, Nation\r\n      ,BCE13 DATETIME  --出生日期, BirthDate\r\n      ,BCK01 INT  --部门ID\r\n      ,ACP01 VARCHAR(2)  --政治面貌, TPLAC.PLACC, PoliticalAffiliationID\r\n      ,AAY01 VARCHAR(4)  --学历, 最高学历, 参见表, TEDUL\r\n      ,ACT01 VARCHAR(4)  --学位, TDEGR\r\n      ,BCE18 VARCHAR(64)  --毕业学校, graduated school\r\n      ,BCE19 DATETIME  --毕业时间, graduated date\r\n      ,ABU02 VARCHAR(64)  --ADIVN\r\n      ,AAD01 VARCHAR(4)  --所学专业\r\n      ,ABG01 VARCHAR(4)  --从事专业, TPPST\r\n      ,BCE23 VARCHAR(32)  --执业证号\r\n      ,BCE24 VARCHAR(32)  --户口所在地, RegisteredResidence\r\n      ,ABS01 VARCHAR(8)  --行政职务\r\n      ,ABI01 VARCHAR(4)  --技术职务, 职称\r\n      ,AAH01 VARCHAR(2)  --聘任职务\r\n      ,BCE28 DATETIME  --工作日期, joindate\r\n      ,AAQ01 VARCHAR(4)  --执业类别\r\n      ,ABE01 VARCHAR(10)  --执业范围, 需要单独的表\r\n      ,BCE31 DATETIME  --进院日期, HireDate\r\n      ,BCE32 DATETIME  --离职日期, DimissionDate\r\n      ,BCE33 VARCHAR(32)  --离职说明, DimissionDescription\r\n      ,BCE34 VARCHAR(128)  --住址, Adress\r\n      ,BCE35 VARCHAR(20)  --办公电话, officephone\r\n      ,BCE36 VARCHAR(20)  --联系电话, phone\r\n      ,BCE37 VARCHAR(20)  --移动电话, mobilephone\r\n      ,BCE38 VARCHAR(64)  --电子邮箱, email\r\n      ,BCE39 VARCHAR(255)  --备注, comment\r\n      ,BCE41 SMALLINT(2)  --状态，0=试用，1=在职，2=离职，3=退休\r\n)';

-- ----------------------------
-- Table structure for breathe
-- ----------------------------
DROP TABLE IF EXISTS `breathe`;
CREATE TABLE `breathe` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景  1辅助呼吸,2停辅助呼吸',
  `value` varchar(255) NOT NULL COMMENT '呼吸值',
  `whethertbm` varchar(255) NOT NULL COMMENT '是否上呼吸机  0否1是',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8 COMMENT='呼吸数据表';

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `devicesclass` int(16) unsigned NOT NULL COMMENT '套餐的科室信息',
  `devicesname` varchar(255) NOT NULL COMMENT '套餐的名字',
  `Devicelist` varchar(255) NOT NULL COMMENT '套餐的设备',
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 COMMENT='套餐数据表';

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
-- Table structure for glucose
-- ----------------------------
DROP TABLE IF EXISTS `glucose`;
CREATE TABLE `glucose` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景 1外出,2检查,3请假,4拒试,5无法侧,6不在',
  `teststate` varchar(255) NOT NULL COMMENT '测试的状态    1空腹,2早餐后1h,3早餐后2h,4中餐前,5中餐后1h,6中餐后2h,7晚餐前,8晚餐后1h,9晚餐后2h,10睡前',
  `value` varchar(255) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='血糖数据表';

-- ----------------------------
-- Table structure for heartrate
-- ----------------------------
DROP TABLE IF EXISTS `heartrate`;
CREATE TABLE `heartrate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景',
  `value` varchar(255) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='心率数据表';

-- ----------------------------
-- Table structure for height
-- ----------------------------
DROP TABLE IF EXISTS `height`;
CREATE TABLE `height` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(250) NOT NULL COMMENT '测试的场景  1卧床,2轮椅,3平车',
  `value` varchar(250) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='身高数据表';

-- ----------------------------
-- Table structure for incident
-- ----------------------------
DROP TABLE IF EXISTS `incident`;
CREATE TABLE `incident` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(250) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `recordscene` varchar(255) NOT NULL COMMENT '测试场景  1入院,2出院,3手术,4分娩,5出生,6转入,7转科,8转院,9死亡,10外出',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `value` varchar(255) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for intakeoutput
-- ----------------------------
DROP TABLE IF EXISTS `intakeoutput`;
CREATE TABLE `intakeoutput` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '这条数据创建的时间',
  `type` int(11) NOT NULL COMMENT '出入量类型，1：入量，2：出量',
  `subtype` int(11) NOT NULL COMMENT '出入量的子类型，1：其他入量或其他出量/ml，2：输液入量或排尿出量/ml，3：饮食入量/ml或大便出量/次',
  `testtime` varchar(255) NOT NULL COMMENT '采集时间',
  `value` int(11) NOT NULL COMMENT '采集值',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `nursename` varchar(255) NOT NULL COMMENT '护士姓名',
  `operationtype` int(11) NOT NULL COMMENT '操作类型，0：无，1：尿管排尿或者排便失禁，2：灌肠排便，3：清洁灌肠，4：肠镜排便，5：人工肛门排便',
  `nurseid` int(4) DEFAULT NULL COMMENT '护士ID',
  `otherdesc` varchar(64) DEFAULT NULL COMMENT '其它出量的补充描述',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8 COMMENT='出入量管理';

-- ----------------------------
-- Table structure for limbs
-- ----------------------------
DROP TABLE IF EXISTS `limbs`;
CREATE TABLE `limbs` (
  `limbsid` smallint(1) DEFAULT NULL COMMENT '肢体ID，1=左上肢，2=左下肢，3=右上肢，4=右下肢',
  `limbs` varchar(64) DEFAULT NULL COMMENT '肢体名称',
  `id` int(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='肢体结构';

-- ----------------------------
-- Table structure for nrl1
-- ----------------------------
DROP TABLE IF EXISTS `nrl1`;
CREATE TABLE `nrl1` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ClassId` varchar(30) DEFAULT NULL COMMENT '科室id',
  `PatientId` varchar(30) DEFAULT '' COMMENT '病人id',
  `NurseId` varchar(30) DEFAULT NULL COMMENT '责任护士id',
  `NurseName` varchar(30) DEFAULT NULL COMMENT '责任护士姓名',
  `DateTime` datetime DEFAULT NULL COMMENT '记录时间',
  `Temperature` varchar(11) DEFAULT NULL COMMENT '体温',
  `Pulse` varchar(11) DEFAULT NULL COMMENT '脉搏/心率',
  `Heartrate` varchar(11) DEFAULT NULL COMMENT '心率',
  `Breathe` varchar(11) DEFAULT NULL COMMENT '呼吸',
  `PressureDIA` varchar(11) DEFAULT NULL COMMENT '血压，低压',
  `PressureSYS` varchar(11) DEFAULT NULL COMMENT '血压，高压',
  `NRA5` varchar(11) DEFAULT NULL COMMENT '意识，1=清醒，2=嗜睡，3=昏睡，4=浅昏迷，5=深昏迷，6=意识浑浊，7=擅妄状态',
  `NRA6A` varchar(100) DEFAULT NULL COMMENT '入量：内容',
  `NRA6B` varchar(11) DEFAULT NULL COMMENT '入量：值，单位ml',
  `NRA7A` varchar(100) DEFAULT NULL COMMENT '出量：内容',
  `NRA7B` varchar(11) DEFAULT NULL COMMENT '出量：值，单位ml',
  `NRA8` text COMMENT '特殊情况',
  `NRA9A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA9B` text COMMENT '自定义项：内容',
  `NRA10A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA10B` text COMMENT '自定义项：内容',
  `NRA11A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA11B` text COMMENT '自定义项：内容',
  `NRA12A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA12B` text COMMENT '自定义项：内容',
  `NRA13A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA13B` text COMMENT '自定义项：内容',
  `NRA14A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA14B` text COMMENT '自定义项：内容',
  `NRA15A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA15B` text COMMENT '自定义项：内容',
  `NRA16A` varchar(60) DEFAULT NULL COMMENT '自定义项：标题',
  `NRA16B` text COMMENT '自定义项：内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for nrl2
-- ----------------------------
DROP TABLE IF EXISTS `nrl2`;
CREATE TABLE `nrl2` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文书ID',
  `VAA01` int(11) NOT NULL COMMENT 'classid科室id',
  `BCK01` int(11) NOT NULL COMMENT 'patientid病人id',
  `NRL02` int(4) DEFAULT NULL COMMENT 'education教育程度',
  `NRL03` int(4) DEFAULT NULL COMMENT 'datasource资料来源',
  `NRL04` int(4) DEFAULT NULL COMMENT 'caregiver日常照顾者',
  `NRL05` int(4) DEFAULT NULL COMMENT 'admission入院方式',
  `NRL06` int(4) DEFAULT NULL COMMENT 'allergy过敏史,0:无,1=有,2=未知',
  `NRL06A` varchar(128) DEFAULT NULL COMMENT 'allergy过敏史详情,Json字符串,1:食物,2:药物,3:其他',
  `NRL07` int(4) DEFAULT NULL COMMENT 'payment医疗费用支付方式',
  `NRL08` int(4) DEFAULT NULL COMMENT 'respond意识状态：呼之（能否回应）',
  `NRL09` int(4) DEFAULT NULL COMMENT 'infocus意识状态：对答（是否切题）',
  `NRL10` int(4) DEFAULT NULL COMMENT 'diet饮食',
  `NRL11` int(4) DEFAULT NULL COMMENT 'masticate咀嚼困难',
  `NRL12` int(4) DEFAULT NULL COMMENT 'oralmucosa口腔黏膜',
  `NRL13` int(4) DEFAULT NULL COMMENT 'deglutition吞咽困难',
  `NRL14` int(4) DEFAULT NULL COMMENT 'sleep睡眠',
  `NRL14A` varchar(128) DEFAULT NULL COMMENT 'sleep睡眠辅助药物',
  `NRL15` int(4) DEFAULT NULL COMMENT 'fatigue醒后疲劳感',
  `NRL16` int(4) DEFAULT NULL COMMENT 'micturition排尿',
  `NRL16A` varchar(128) DEFAULT NULL COMMENT 'micturition排尿的其它方式',
  `NRL17` varchar(128) DEFAULT NULL COMMENT 'cacation排便',
  `NRL17A` varchar(128) DEFAULT NULL COMMENT 'cacation排便的其它方式',
  `NRL18` varchar(64) DEFAULT NULL COMMENT 'time次数,1=代表1次/*天,2=代表*次/天',
  `NRL19` int(4) DEFAULT NULL COMMENT 'limb exercise肢体活动 ',
  `LimbsId` smallint(1) DEFAULT NULL COMMENT '关联肢体表Limbs.LimbsId ',
  `NRL20` int(4) DEFAULT NULL COMMENT 'self-careAbility自理能力',
  `NRL20A` varchar(64) DEFAULT NULL COMMENT 'self-careAbility自理能力部分自理内容',
  `NRL21` int(4) DEFAULT NULL COMMENT 'skin皮肤状况',
  `NRL22` varchar(255) DEFAULT NULL COMMENT 'skinDesc皮肤状况详细描述',
  `NRL23` varchar(64) DEFAULT NULL COMMENT 'language常用语言',
  `NRL24` int(4) DEFAULT NULL COMMENT 'abilityOfExpression语言表达能力',
  `NRL25` int(1) DEFAULT NULL COMMENT 'smoke生活习惯-吸烟',
  `NRL26` int(4) DEFAULT NULL COMMENT 'smoke生活习惯-吸烟的次数.*支/天',
  `NRL27` int(1) DEFAULT NULL COMMENT 'alcoholomania生活习惯-嗜酒',
  `NRL28` int(4) DEFAULT NULL COMMENT 'alcoholomania生活习惯-嗜酒的次数.*两/天',
  `NRL29` varchar(255) DEFAULT NULL COMMENT 'signDesc其他症状或者体征描述',
  `NRL30` int(4) DEFAULT NULL COMMENT 'inHospitalNotification住院告知',
  `NRL30A` varchar(64) DEFAULT NULL COMMENT 'inHospitalNotification住院告知的其它内容',
  `NRL31` varchar(255) DEFAULT NULL COMMENT 'basicNursing基础护理',
  `NRL32` varchar(255) DEFAULT NULL COMMENT 'specificNursing专科护理',
  `NRL33` varchar(255) DEFAULT NULL COMMENT 'patientSafety患者安全',
  `NRL34` varchar(255) DEFAULT NULL COMMENT 'other其它护理',
  `NRL35` varchar(255) DEFAULT NULL COMMENT 'shiftChange交接班重点',
  `NRL36` varchar(255) DEFAULT NULL COMMENT 'focus提醒医生给予关注',
  `NRL37` varchar(255) DEFAULT NULL COMMENT 'care提醒医生给予关爱',
  `NRL38` datetime DEFAULT NULL COMMENT 'recordDate记录时间',
  `BCE01A` int(11) NOT NULL COMMENT 'NursingId责任护士ID',
  `BCE03A` varchar(64) NOT NULL COMMENT 'NursingName责任护士签名',
  `NRL07A` varchar(64) DEFAULT NULL COMMENT 'payment医疗费用支付方式的其它支付方式',
  `NRL12A` varchar(64) DEFAULT NULL COMMENT 'oralmucosa口腔黏膜其它状况',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8 COMMENT='首次护理记录单';

-- ----------------------------
-- Table structure for nursingrecords
-- ----------------------------
DROP TABLE IF EXISTS `nursingrecords`;
CREATE TABLE `nursingrecords` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Updated` datetime DEFAULT NULL COMMENT '更新时间',
  `NursType` int(20) DEFAULT NULL COMMENT '文书类型，1=护理记录单，2=首次护理记录单',
  `NursingId` varchar(30) DEFAULT NULL COMMENT '责任护士id',
  `NursingName` varchar(30) DEFAULT NULL COMMENT '责任护士姓名',
  `ClassId` varchar(30) DEFAULT NULL COMMENT '科室id',
  `PatientId` varchar(30) DEFAULT NULL COMMENT '病人ID',
  `RecordId` int(11) DEFAULT NULL COMMENT '文书id',
  `Comment` varchar(255) NOT NULL COMMENT '备注（动作）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=236 DEFAULT CHARSET=utf8mb4 COMMENT='文书查询记录表';

-- ----------------------------
-- Table structure for pressure
-- ----------------------------
DROP TABLE IF EXISTS `pressure`;
CREATE TABLE `pressure` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) DEFAULT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(11) NOT NULL COMMENT '测试的场景  1外出,2检查,3请假,4拒试,5无法侧,6未测',
  `diavalue` varchar(11) NOT NULL COMMENT '低压值',
  `sysvalue` varchar(11) NOT NULL COMMENT '高压值',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='血压数据表';

-- ----------------------------
-- Table structure for pulse
-- ----------------------------
DROP TABLE IF EXISTS `pulse`;
CREATE TABLE `pulse` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景',
  `value` varchar(255) NOT NULL COMMENT '脉搏值',
  `whetherbriefness` varchar(255) NOT NULL COMMENT '是否短促  0否1是',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 COMMENT='脉搏测量数据表';

-- ----------------------------
-- Table structure for skin
-- ----------------------------
DROP TABLE IF EXISTS `skin`;
CREATE TABLE `skin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(250) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `value` varchar(255) NOT NULL COMMENT '值',
  `recordscene` varchar(255) NOT NULL COMMENT '测试场景',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for spo2h
-- ----------------------------
DROP TABLE IF EXISTS `spo2h`;
CREATE TABLE `spo2h` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(250) NOT NULL COMMENT '测试的场景',
  `value` varchar(250) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='血氧数据表';

-- ----------------------------
-- Table structure for succession
-- ----------------------------
DROP TABLE IF EXISTS `succession`;
CREATE TABLE `succession` (
  `ID` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `DataTime` datetime DEFAULT NULL COMMENT '日期',
  `Type` tinyint(2) NOT NULL COMMENT '类型1，白班 2，晚班 3，夜班',
  `NursingName` varchar(32) DEFAULT NULL COMMENT '护士名称',
  `ClassId` varchar(32) DEFAULT NULL COMMENT '科室ID',
  `NoldPatient` varchar(32) DEFAULT NULL COMMENT '原病人数',
  `NnowPatient` varchar(32) DEFAULT NULL COMMENT '现病人数',
  `NintHospital` varchar(32) DEFAULT NULL COMMENT '入院病人数',
  `NoutHospital` varchar(32) DEFAULT NULL COMMENT '出院',
  `Ninto` varchar(32) DEFAULT NULL COMMENT '转入数',
  `Nout` varchar(32) DEFAULT NULL COMMENT '转出',
  `Nsurgery` varchar(32) DEFAULT NULL COMMENT '手术',
  `Nchildbirth` varchar(32) DEFAULT NULL COMMENT '分娩',
  `Ncritically` varchar(32) DEFAULT NULL COMMENT '病危',
  `Ndeath` varchar(32) DEFAULT NULL COMMENT '死亡',
  `NintensiveCare` varchar(32) DEFAULT NULL COMMENT '特护',
  `NprimaryCare` varchar(32) NOT NULL COMMENT '一级护理',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='交接班';

-- ----------------------------
-- Table structure for successiondetails
-- ----------------------------
DROP TABLE IF EXISTS `successiondetails`;
CREATE TABLE `successiondetails` (
  `ID` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `BedID` varchar(32) DEFAULT NULL COMMENT '床位号',
  `PatientName` varchar(32) DEFAULT NULL COMMENT '病人名称',
  `Piagnosis` varchar(32) DEFAULT NULL COMMENT '诊断',
  `Typte` tinyint(2) DEFAULT NULL COMMENT '1,入院，2，出院，3，转入，4，转出，5，手术，6，分娩，7，病危，8，死亡，9，特护，10，一级护理',
  `Comment1` varchar(255) DEFAULT NULL COMMENT '备注',
  `DataTime` datetime DEFAULT NULL COMMENT '日期',
  `ClassId` varchar(32) DEFAULT NULL COMMENT '科室ID',
  `Comment2` varchar(255) DEFAULT NULL COMMENT '备注',
  `Comment3` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='交接班详情';

-- ----------------------------
-- Table structure for temperature
-- ----------------------------
DROP TABLE IF EXISTS `temperature`;
CREATE TABLE `temperature` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景 1物理降温,2药物降温,3冰毯降温,4停冰毯降温,5药物+物理降温,6无降温,7不升,8外出,9检查,10请假,11拒试,12无法侧,13未测',
  `ttemptype` varchar(255) NOT NULL DEFAULT '' COMMENT '体温的类型,1腋温,2耳温,3口温,4肛温,5额温',
  `coolingvalue` varchar(255) NOT NULL COMMENT '降温的体温',
  `value` varchar(255) NOT NULL COMMENT '体温值',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=181 DEFAULT CHARSET=utf8 COMMENT='体温测量数据表';

-- ----------------------------
-- Table structure for temperature_copy
-- ----------------------------
DROP TABLE IF EXISTS `temperature_copy`;
CREATE TABLE `temperature_copy` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景',
  `ttemptype` varchar(255) NOT NULL COMMENT '体温的类型',
  `coolingvalue` varchar(255) NOT NULL COMMENT '降温的体温',
  `value` varchar(255) NOT NULL COMMENT '体温值',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8 COMMENT='体温测量数据表';

-- ----------------------------
-- Table structure for temperature_copy2
-- ----------------------------
DROP TABLE IF EXISTS `temperature_copy2`;
CREATE TABLE `temperature_copy2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景',
  `ttemptype` varchar(255) NOT NULL COMMENT '体温的类型',
  `coolingvalue` varchar(255) NOT NULL COMMENT '降温的体温',
  `value` varchar(255) NOT NULL COMMENT '体温值',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8 COMMENT='体温测量数据表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `code` varchar(32) NOT NULL COMMENT '员工工号',
  `password` varchar(128) NOT NULL COMMENT '密码，MD5加密，管理员密码fitcome.com.nursing',
  `employeeid` int(11) NOT NULL COMMENT '员工ID，关联BCE1.BCE01',
  `authority` tinyint(2) DEFAULT NULL COMMENT '权限登记，0：正常，1：护士长，2：管理员',
  `createdate` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  `departmentid` int(4) DEFAULT NULL COMMENT '科室ID',
  `key` varchar(128) DEFAULT NULL COMMENT '明文密码，用于找回明码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Table structure for vaa1
-- ----------------------------
DROP TABLE IF EXISTS `vaa1`;
CREATE TABLE `vaa1` (
  `VAA01` int(11) NOT NULL COMMENT '病人ID, PATFU, PATID',
  `VAA02` varchar(64) DEFAULT NULL COMMENT '会员卡号',
  `VAA03` varchar(20) DEFAULT NULL COMMENT '门诊号 Clinic Patient No, Out patient No, OPTNO',
  `VAA04` varchar(20) DEFAULT NULL COMMENT '住院号, In Patient No, IPTNO',
  `VAA05` varchar(64) DEFAULT NULL COMMENT '姓名, PATNM, FirstName + Middlename',
  `VAA06` varchar(64) DEFAULT NULL COMMENT '监护人1',
  `ABBRP` varchar(10) DEFAULT NULL COMMENT '拼音',
  `ABBRW` varchar(10) DEFAULT NULL COMMENT '五笔',
  `ABW01` varchar(1) DEFAULT NULL COMMENT '性别',
  `VAA10` int(11) DEFAULT NULL COMMENT '年龄',
  `AAU01` varchar(1) DEFAULT NULL COMMENT '年龄单位',
  `VAA12` datetime DEFAULT NULL COMMENT '出生日期, birth date',
  `ACK01` varchar(2) DEFAULT NULL COMMENT '婚姻, marital status',
  `VAA14` varchar(2) DEFAULT NULL COMMENT '身份证件, 指公安机关签发的有效证件',
  `VAA15` varchar(20) DEFAULT NULL COMMENT '身份证号, IDNo, IDNumber',
  `VAA16` varchar(20) DEFAULT NULL COMMENT '其他证件',
  `ABJ01` varchar(2) DEFAULT NULL COMMENT '付款方式, 医疗付款方式',
  `BDP02` varchar(50) DEFAULT NULL COMMENT '病人类别',
  `ABC02` varchar(20) DEFAULT NULL COMMENT '病人费别',
  `VAA20` varchar(64) DEFAULT NULL COMMENT '出生地点, 出生地, BirthPlace',
  `ACM02` varchar(20) DEFAULT NULL COMMENT '从业状况, 身份',
  `ACC02` varchar(32) DEFAULT NULL COMMENT '国籍, Nationality',
  `ABQ02` varchar(32) DEFAULT NULL COMMENT '民族, Nation',
  `VAA25` varchar(64) DEFAULT NULL COMMENT '籍贯, 省.市, NativePlace',
  `VAA26` varchar(20) DEFAULT NULL COMMENT '宗教, 佛教, 伊斯兰教, 基督教, 天主教, 犹太教, 耶稣教, 其他, Religion, 暂时不用',
  `VAA27` varchar(20) DEFAULT NULL COMMENT '种族, 人种, Ethnic, 暂时不用',
  `VAA28` varchar(128) DEFAULT NULL COMMENT '户口地址, 户籍地址, RegisteredAddress',
  `VAA29` varchar(6) DEFAULT NULL COMMENT '户籍邮编, RegisteredPostCode',
  `VAA30` varchar(20) DEFAULT NULL COMMENT '户籍电话, RegisteredPhone',
  `VAA31` varchar(20) DEFAULT NULL COMMENT '省市',
  `VAA32` varchar(20) DEFAULT NULL COMMENT '县市, 城市, 县区市',
  `VAA33` varchar(64) DEFAULT NULL COMMENT '地址, 常住地址, resident address',
  `VAA34` varchar(20) DEFAULT NULL COMMENT '电话',
  `VAA35` varchar(13) DEFAULT NULL COMMENT '移动电话',
  `VAA36` varchar(128) DEFAULT NULL COMMENT '电子邮箱',
  `VAA37` varchar(64) DEFAULT NULL COMMENT '其他联系方式, 如QQ号, MSN号, Skype号等, OtherCommunications',
  `VAA38` varchar(48) DEFAULT NULL COMMENT '学历',
  `VAA39` varchar(64) DEFAULT NULL COMMENT '监护人2, Guardian',
  `VAA40` varchar(20) DEFAULT NULL COMMENT '联系人姓名, contact person name',
  `VAA41` varchar(32) DEFAULT NULL COMMENT '与病人关系, RelationShip',
  `VAA42` varchar(64) DEFAULT NULL COMMENT '联系人地址, Contact person address',
  `VAA43` varchar(20) DEFAULT NULL COMMENT '联系人电话, contact person telephone',
  `VAA44` varchar(16) DEFAULT NULL COMMENT '联系人移动电话, Contact person Mobile Phone',
  `BAQ01` int(11) DEFAULT NULL COMMENT '合同单位ID',
  `BAQ02` varchar(10) DEFAULT NULL COMMENT '单位编码, CoCode',
  `VAA47` varchar(64) DEFAULT NULL COMMENT '工作单位, CoName',
  `VAA48` varchar(20) DEFAULT NULL COMMENT '单位电话, CoTelephone',
  `VAA49` varchar(64) DEFAULT NULL COMMENT '单位地址, CoAddress',
  `VAA50` varchar(6) DEFAULT NULL COMMENT '单位邮编, CoPostCode',
  `VAA51` varchar(64) DEFAULT NULL COMMENT '单位开户行',
  `VAA52` varchar(20) DEFAULT NULL COMMENT '单位银行帐号',
  `VAA53` varchar(20) DEFAULT NULL COMMENT '担保人, Guarantor',
  `VAA54` decimal(18,2) DEFAULT NULL COMMENT '信用额度, 担保额度, CreditLimit',
  `VAA55` tinyint(4) DEFAULT NULL COMMENT '担保性质, CreditType',
  `VAA56` int(11) DEFAULT NULL COMMENT '住院次数, HospitalizationNumber',
  `VAA57` datetime DEFAULT NULL COMMENT '就诊时间, LastVisitDate',
  `BCK01A` int(11) DEFAULT NULL COMMENT '就诊科室, LastVisitDeptID',
  `VAA61` tinyint(4) DEFAULT NULL COMMENT '就诊状态, 0=无, 1=门诊, 2=住院, 3=出院, 4=转院, 5=死亡, 9=其他` VisitState',
  `VAA62` varchar(255) DEFAULT NULL COMMENT '过敏史',
  `BDX02` varchar(64) DEFAULT NULL COMMENT '了解途径, 病人了解医院的方式(如电视广告, 介绍, 户外广告等)',
  `VAA64` varchar(255) DEFAULT NULL COMMENT '备注',
  `VBU01` int(11) DEFAULT NULL COMMENT '帐号ID',
  `VAA66` varchar(64) DEFAULT NULL COMMENT '病案号',
  `VAA67` varchar(64) DEFAULT NULL COMMENT '查询密码',
  `IAK05` varchar(50) DEFAULT NULL COMMENT '社会保障号',
  `IAA01` int(11) DEFAULT NULL COMMENT '保险机构',
  `BCK01B` int(11) DEFAULT NULL COMMENT '科室ID',
  `BCK01C` int(11) DEFAULT NULL COMMENT '病区ID',
  `BCQ04` varchar(20) DEFAULT NULL COMMENT '床号',
  `VAA73` datetime DEFAULT NULL COMMENT '入院时间',
  `VAA74` datetime DEFAULT NULL COMMENT '出院时间',
  `VAA75` datetime DEFAULT NULL COMMENT '建档时间',
  `BEP05` decimal(18,4) DEFAULT NULL COMMENT '住院报警值',
  `BEP06` decimal(18,4) DEFAULT NULL COMMENT '住院信用额度',
  `ABL01A` varchar(2) DEFAULT NULL COMMENT '正定型',
  `ABL01B` varchar(2) DEFAULT NULL COMMENT '反定型',
  `VAA76` datetime DEFAULT NULL COMMENT '有效时间',
  `ABL01` varchar(2) DEFAULT NULL COMMENT '血型',
  `BEP06B` decimal(18,4) DEFAULT NULL COMMENT '门诊信用额度',
  `VAA78` varchar(1) DEFAULT NULL COMMENT 'Rh血型',
  `VAA82` varchar(64) DEFAULT NULL COMMENT '健康卡号',
  `VAA01A` int(11) DEFAULT NULL COMMENT '相关ID',
  `VAA84` varchar(20) DEFAULT NULL COMMENT '体检登记号'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='病人资料';

-- ----------------------------
-- Table structure for vaf2
-- ----------------------------
DROP TABLE IF EXISTS `vaf2`;
CREATE TABLE `vaf2` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='CREATE TABLE VAF2--住院病人医嘱记录\r\n(\r\n        VAF01  INT  --ID\r\n      ,VAF01A INT  --相关ID, 关联字段：VAF1.VAF01\r\n      ,VAF01B INT  --前提ID, 关联字段：VAF1.VAF01\r\n      ,VAF04  TINYINT  --1：门诊;2：住院\r\n      ,VAA01  INT  --病人ID, 关联字段：VAA1.VAA01\r\n      ,VAF06  INT  --就诊ID, 主页ID\r\n      ,VAF07  INT  --婴儿ID\r\n      ,BCK01A INT  --病人科室ID, 关联字段：BCK1.BCK01\r\n      ,ROWNR  INT  --次序\r\n      ,VAF10  TINYINT  --1：新嘱;2：疑问;3：校对;4：作废;5：删除 ;6：暂停;7：启用;8：已发送或停止;9：确认停止;10：皮试结果\r\n      ,VAF11  TINYINT  --医嘱类型, 1=长期医嘱, 2=临时医嘱\r\n      ,BDA01  VARCHAR(2)  --诊疗类型, 关联字段：BDA1.BDA01\r\n      ,BBX01  INT  --诊疗项目ID, 关联字段：BBX1.BBX01\r\n      ,VAF14  VARCHAR(60)  --标本部位\r\n      ,VAF15  VARCHAR(30)  --检查方法\r\n      ,BBY01  INT  --收费项目ID, 关联字段：BBY1.BBY01\r\n      ,VAF17  INT  --天数, day number\r\n      ,VAF18  NUMERIC(18, 4)  --剂量, 单次用量\r\n      ,VAF19  VARCHAR(10)  --用量\r\n      ,VAF20  NUMERIC(18, 4)  --单量\r\n      ,VAF21  NUMERIC(18, 4)  --数量\r\n      ,VAF22  VARCHAR(1024)  --医嘱\r\n      ,VAF23  VARCHAR(128)  --医师嘱托\r\n      ,BCK01B INT  --执行科室ID, 关联字段：BCK1.BCK01\r\n      ,VAF25  VARCHAR(10)  --空值=不需要做皮试 、 +或-表示阳性或阴性  、？表示需要做皮试但还没填皮试结果\r\n      ,VAF26  VARCHAR(20)  --执行频次\r\n      ,VAF27  INT  --频率次数\r\n      ,VAF28  TINYINT  --频率间隔\r\n      ,VAF29  VARCHAR(4)  --间隔单位\r\n      ,VAF30  VARCHAR(64)  --执行时间方案\r\n      ,VAF31  TINYINT  --计价特性  0=正常, 1=自费\r\n      ,VAF32  TINYINT  --0：正常; 1＝给药途径\r\n      ,VAF33  TINYINT  --0：标记未用;1：正常 2：自动停止\r\n      ,VAF34  TINYINT  --可否分零\r\n      ,VAF35  TINYINT  --0：正常 1：紧急\r\n      ,VAF36  DATETIME  --开始执行时间\r\n      ,VAF37  DATETIME  --执行终止时间\r\n      ,VAF38  DATETIME  --上次执行时间\r\n      ,BCK01C INT  --开嘱科室ID, 关联字段：BCK1.BCK01\r\n      ,BCE02A VARCHAR(20)  --医师编码, 关联字段：BCE1.BCE02\r\n      ,BCE03A VARCHAR(20)  --开嘱医师, 关联字段：BCE1.BCE03\r\n      ,VAF42  DATETIME  --开嘱时间\r\n      ,BCE03B VARCHAR(20)  --开嘱护士, 关联字段：BCE1.BCE03\r\n      ,BCE03C VARCHAR(20)  --校对护士, 关联字段：BCE1.BCE03\r\n      ,VAF45  DATETIME  --校对时间\r\n      ,BCE03D VARCHAR(20)  --停嘱医生, 关联字段：BCE1.BCE03\r\n      ,VAF47  DATETIME  --停嘱时间\r\n      ,BCE03E VARCHAR(20)  --停嘱护士, 关联字段：BCE1.BCE03\r\n      ,BCE03F VARCHAR(20)  --停嘱校对护士, 关联字段：BCE1.BCE03\r\n    ';

-- ----------------------------
-- Table structure for warn
-- ----------------------------
DROP TABLE IF EXISTS `warn`;
CREATE TABLE `warn` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `NurseId` varchar(255) NOT NULL DEFAULT '' COMMENT '护士id',
  `NurseName` varchar(255) NOT NULL COMMENT '护士姓名',
  `PatientId` varchar(255) DEFAULT NULL COMMENT '病人id',
  `ClassId` varchar(255) NOT NULL DEFAULT '' COMMENT '科室id',
  `BedId` varchar(255) DEFAULT NULL COMMENT '床号',
  `WarnTime` varchar(255) NOT NULL COMMENT '提醒时间',
  `WarnType` int(11) DEFAULT NULL COMMENT '提醒类型，1=响铃，2=震动，4=响铃+震动',
  `Name` varchar(255) DEFAULT NULL COMMENT '提醒名称',
  `Desc` varchar(255) DEFAULT NULL COMMENT '提醒描述',
  `date_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上传时当前时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8 COMMENT='提醒表';

-- ----------------------------
-- Table structure for weight
-- ----------------------------
DROP TABLE IF EXISTS `weight`;
CREATE TABLE `weight` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nurseid` varchar(255) NOT NULL COMMENT '护士id',
  `nursename` varchar(255) NOT NULL COMMENT '护士名字',
  `patientid` varchar(255) NOT NULL COMMENT '病人id',
  `testtime` datetime NOT NULL COMMENT '测试时间',
  `recordscene` varchar(255) NOT NULL COMMENT '测试的场景  1卧床,2轮椅,3平车',
  `value` varchar(255) NOT NULL COMMENT '值',
  PRIMARY KEY (`id`),
  KEY `patientid` (`patientid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='体重数据表';

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
