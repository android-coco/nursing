//  Created by JP

package model

import (
	"fit"
)

// 首次护理记录单
type NRL2 struct {
	ID     int64    `xorm:"pk autoincr comment(文书id)"`
	PatientId  int64  `xorm:"comment(patientid病人id)"`
	BCK01  int    `xorm:"comment(classid科室id)"`
	NRL38  string `xorm:"comment(recordDate记录时间)"`
	NurseId string `xorm:"comment(NursingId责任护士ID)"`
	NurseName string `xorm:"comment(NursingName责任护士签名)"`

	NRL02  string `xorm:"comment(education教育程度)"`
	NRL03  string `xorm:"comment(datasource资料来源)"`
	NRL04  string `xorm:"comment(caregiver日常照顾者)"`
	NRL05  string `xorm:"comment(admission入院方式)"`
	NRL06  string `xorm:"comment(allergy过敏史)"`
	NRL06A string `xorm:"comment(allergy过敏史,json字符串)"`
	NRL07  string `xorm:"comment(payment医疗费用支付方式)"`
	NRL07A string `xorm:"comment(payment医疗费用支付方式的其它方式)"`
	NRL08  string `xorm:"comment(respond意识状态-呼之-能否回应)"`
	NRL09  string `xorm:"comment(infocus意识状态-对答-是否切题)"`

	NRL10   string `xorm:"comment(diet饮食)"`
	NRL11   string `xorm:"comment(masticate咀嚼困难)"`
	NRL12   string `xorm:"comment(oralmucosa口腔黏膜)"`
	NRL12A  string `xorm:"comment(oralmucosa口腔黏膜的其它状况)"`
	NRL13   string `xorm:"comment(deglutition吞咽困难)"`
	NRL14   string `xorm:"comment(sleep睡眠)"`
	NRL14A  string `xorm:"comment(sleep睡眠辅助药物)"`
	NRL15   string `xorm:"comment(fatigue醒后疲劳感)"`
	NRL16   string `xorm:"comment(micturition排尿)"`
	NRL16A  string `xorm:"comment(micturition排尿的其它方式)"`
	NRL17   string `xorm:"comment(cacation排便)"`
	NRL17A  string `xorm:"comment(cacation排便的其它方式)"`
	NRL18   string `xorm:"comment(time次数,1=代表1次/*天,2=代表*次/天)"`
	NRL19   string `xorm:"comment(limbExercise肢体活动)"`
	LimbsId string `xorm:"comment(关联肢体表Limbs.LimbsId)"`

	NRL20  string `xorm:"comment(self-careAbility自理能力)"`
	NRL20A string `xorm:"comment(self-careAbility自理能力部分自理内容)"`
	NRL21  string `xorm:"comment(skin皮肤状况)"`
	NRL22  string `xorm:"comment(skinDesc皮肤状况详细描述)"`
	NRL23  string `xorm:"comment(language常用语言)"`
	NRL24  string `xorm:"comment(abilityOfExpression语言表达能力)"`
	NRL25  string `xorm:"comment(smoke生活习惯-吸烟)"`
	NRL26  string `xorm:"comment(smoke生活习惯-吸烟的次数,*支/天)"`
	NRL27  string `xorm:"comment(alcoholomania生活习惯-嗜酒)"`
	NRL28  string `xorm:"comment(alcoholomania生活习惯-嗜酒的次数,*两/天)"`
	NRL29  string `xorm:"comment(signDesc其他症状或者体征描述)"`

	NRL30  string `xorm:"comment(inHospitalNotification住院告知)"`
	NRL30A string `xorm:"comment(inHospitalNotification住院告知的其它方式)"`
	NRL31  string `xorm:"comment(basicNursing基础护理)"`
	NRL32  string `xorm:"comment(specificNursing专科护理)"`
	NRL33  string `xorm:"comment(patientSafety患者安全)"`
	NRL34  string `xorm:"comment(other其他护理)"`
	NRL35  string `xorm:"comment(shiftChange交接班重点)"`
	NRL36  string `xorm:"comment(focus提醒医生给予关注)"`
	NRL37  string `xorm:"comment(care提醒医生给予关爱)"`
	NRL39A string `xorm:"comment(审核时间)"`
	NRL39B string `xorm:"comment(审核护士id)"`
	NRL39C string `xorm:"comment(审核护士签名)"`
}

/*插入 首次护理记录*/
func (nrl2 *NRL2) InsertToDatabase() (int64, error) {
	_, err := fit.MySqlEngine().InsertOne(nrl2)
	var nrl_id int64
	if err == nil {
		slice := make([]NRL2, 0)
		err = fit.MySqlEngine().SQL("select id from NRL2 where NRL38 = ? and PatientId = ?", nrl2.NRL38, nrl2.PatientId).Find(&slice)
		nrl_id = slice[0].ID
	}
	return nrl_id, err
}

/*查询是否存在某个病人的首次护理单*/
func IsExistNRL2(withPatientId string) (bool, error) {
	has, err := fit.MySqlEngine().SQL("select id from NRL2 where PatientId = ?", withPatientId).Exist()
	if err == nil && has {
		return true, nil
	} else {
		return false, err
	}

}

func (nrl2 *NRL2) UpdateRecordDatas(withNrlId int64) error {
	//Cols("NRL02", "NRL03", "NRL04", "NRL05", "NRL06", "NRL06A", "NRL07", "NRL07A", "NRL08", "NRL09", "NRL10", "NRL11", "NRL12", "NRL12A", "NRL13", "NRL14", "NRL14A", "NRL15", "NRL16", "NRL16A", "NRL17", "NRL17A", "NRL18", "NRL19", "NRL20", "NRL20A", "NRL21", "NRL22", "NRL23", "NRL24", "NRL25", "NRL26", "NRL27", "NRL28", "NRL29", "NRL30", "NRL30A", "NRL31", "NRL32", "NRL33", "NRL34", "NRL35", "NRL36", "NRL37", "NRL38", "NRL39A", "NRL39B", "NRL39C")
	_, err := fit.MySqlEngine().Table("NRL2").Where("id = ?", withNrlId).AllCols().Omit("ID").Update(nrl2)
	return err
}

func QueryNRL2(withNRLId string) (NRL2, error) {
	var nrl2 NRL2
	_, err := fit.MySqlEngine().SQL("select * from NRL2 where id = ?", withNRLId).Get(&nrl2)
	return nrl2, err
}

func QueryNRL2WithPid(pid string) (NRL2, error) {
	var nrl2 NRL2
	_, err := fit.MySqlEngine().SQL("select * from NRL2 where PatientId = ?", pid).Get(&nrl2)
	return nrl2, err
}
