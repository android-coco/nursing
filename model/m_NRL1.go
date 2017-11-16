package model

import (
	"time"
	"fit"
)

//护理记录单 表头
type NRL1Title struct {
	ID    int64  `xorm:"pk autoincr comment(文书id)"`
	BCK01 int64  `xorm:"comment(classid科室id)"`
	VAA01 int64  `xorm:"comment(patientid病人id)"`

	NRT01 string `xorm:"comment(表头1，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT01V string `xorm:"comment(表头1，自定义内容)"`
	NRT02 string `xorm:"comment(表头2，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT02V string `xorm:"comment(表头2，自定义内容)"`
	NRT03 string `xorm:"comment(表头3，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT03V string `xorm:"comment(表头3，1=意识，自定义内容)"`
	NRT04 string `xorm:"comment(表头4，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT04V string `xorm:"comment(表头4，1=意识，自定义内容)"`
	NRT05 string `xorm:"comment(表头5，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT05V string `xorm:"comment(表头5，1=意识，自定义内容)"`
	NRT06 string `xorm:"comment(表头6，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT06V string `xorm:"comment(表头6，1=意识，自定义内容)"`
	NRT07 string `xorm:"comment(表头7，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT07V string `xorm:"comment(表头7，1=意识，自定义内容)"`
}

//护理记录单
type NRL1 struct {
	Id          int64
	PatientId   string    // 病人id
	ClassId     string    // 科室id
	Temperature string    // 体温
	Pulse       string    // 脉搏
	Heartrate   string    // /心率
	Breathe     string    // 呼吸
	PressureDIA string    // 血压，低压
	PressureSYS string    // 血压，高压
	NRA5        string    // 意识，1=清醒，2=嗜睡，3=昏睡，4=浅昏迷，5=深昏迷，6=意识浑浊，7=擅妄状态
	NRA6A       string    // 入量：内容
	NRA6B       string    // 入量：值， 单位ml
	NRA7A       string    // 出量：内容
	NRA7B       string    // 出量：值，单位ml
	NRA8        string    // 特殊情况说明
	DateTime    time.Time // 记录时间
	NurseId     string    // 责任护士id
	NurseName   string    // 责任护士姓名
	NRA9A       string    // 自定义项：标题
	NRA9B       string    // 自定义项：内容
	NRA10A      string    // 自定义项：标题
	NRA10B      string    // 自定义项：内容
	NRA11A      string    // 自定义项：标题
	NRA11B      string    // 自定义项：内容
	NRA12A      string    // 自定义项：标题
	NRA12B      string    // 自定义项：内容
	NRA13A      string    // 自定义项：标题
	NRA13B      string    // 自定义项：内容
	NRA14A      string    // 自定义项：标题
	NRA14B      string    // 自定义项：内容
	NRA15A      string    // 自定义项：标题
	NRA15B      string    // 自定义项：内容
	//NRA16A      string    // 自定义项：标题
	NRA16B string // 瞳孔 1=灵敏，2=迟钝，3=消失
}

/*type NRL1 struct {
	id          int64
	PatientId   string     // 病人id
	Temperature float32   // 体温
	Pulse       int       // 脉搏/心率
	Breathe     int       // 呼吸
	PressureDIA int       // 血压，低压
	PressureSYS int       // 血压，高压
	NRA5        int       // 意识，1=清醒，2=嗜睡，3=昏睡，4=浅昏迷，5=深昏迷，6=意识浑浊，7=擅妄状态
	NRA6A       string    // 入量：内容
	NRA6B       int       // 入量：值， 单位ml
	NRA7A       string    // 出量：内容
	NRA7B       int       // 出量：值，单位ml
	NRA8        string    // 特殊情况说明
	DateTime    time.Time // 记录时间
	NurseId     string     // 责任护士id
	NurseName   string    // 责任护士姓名
	NRA9A       string    // 自定义项：标题
	NRA9B       string    // 自定义项：内容
	NRA10A      string    // 自定义项：标题
	NRA10B      string    // 自定义项：内容
	NRA11A      string    // 自定义项：标题
	NRA11B      string    // 自定义项：内容
	NRA12A      string    // 自定义项：标题
	NRA12B      string    // 自定义项：内容
	NRA13A      string    // 自定义项：标题
	NRA13B      string    // 自定义项：内容
	NRA14A      string    // 自定义项：标题
	NRA14B      string    // 自定义项：内容
	NRA15A      string    // 自定义项：标题
	NRA15B      string    // 自定义项：内容
	NRA16A      string    // 自定义项：标题
	NRA16B      string    // 自定义项：内容
}*/

func (m *NRL1) InsertData() (int64, error) {
	_, err := fit.MySqlEngine().Table("NRL1").Insert(m)
	var rid int64 = 0
	if err == nil {
		_, err1 := fit.MySqlEngine().Table("NRL1").Where("patientid = ?", m.PatientId).Cols("id").Desc("id").Get(&rid)
		if err1 == nil {
			m.Id = rid
			return rid, err
		}
	} else {
		return 0, err
	}
	return 0, err
}

func (m NRL1) UpdateData(id int64) (int64, error) {
	return fit.MySqlEngine().Id(id).Update(&m)
}

// 查询某一条护理记录单
func QueryNRL1(rid string) (NRL1, error) {
	var nr1 NRL1
	_, err := fit.MySqlEngine().Table("NRL1").Where("id = ?", rid).Get(&nr1)
	if err != nil {
		return NRL1{}, err
	} else {
		return nr1, nil
	}
}

// 根据病人id，时间范围 查询他所有的护理记录单
func (m NRL1) QueryNRL1ByDate(startDate, endDate string, page int) ([]NRL1, error) {
	var nrl1List []NRL1
	var err error
	if startDate == "" || endDate == "" {
		err = fit.MySqlEngine().Table("NRL1").Where("patientid = ?", m.PatientId).Limit(20, page*20).Find(&nrl1List)
	} else {
		err = fit.MySqlEngine().Table("NRL1").Where("patientid = ?", m.PatientId).And("datetime > ?", startDate).And("datetime < ?", endDate).Limit(20, page*20).Find(&nrl1List)
	}
	return nrl1List, err
}

/*
// 护理记录单录入
1=体温
2=呼吸
3=脉搏
4=心率
5=事件
6A=输入量
6B=排出量
7=血压
8=体重
9=身高
10=皮试
*/
func GetNRL1Data(ty, pid string, day time.Time) ([]map[string]string, error) {
	loc, _ := time.LoadLocation("Local")
	var tablename string
	//var results []map[string]string
	switch ty {
	case "1":
		tablename = "`Temperature` "
	case "2":
		tablename = "`Breathe` "
	case "3":
		tablename = "`Pulse` "
	case "4":
		tablename = "`Heartrate` "
	case "5":
		tablename = "`Incident` "
	case "6A":
		tablename = "`IntakeOutput` "
	case "6B":
		tablename = "`IntakeOutput` "
	case "7":
		tablename = "`Pressure` "
	case "8":
		tablename = "`Weight` "
	case "9":
		tablename = "`Height` "
	case "10":
		tablename = "`Skin` "
	default:
		tablename = ""
	}

	sqlstr := "SELECT * FROM " + tablename + "WHERE `patientid` = ? AND `testtime` >= ? AND `testtime` < ?"

	//if "5" == ty {
	//	sqlstr = "SELECT * FROM " + tablename + "WHERE `type` = 1 AND `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1"
	//} else if "6" == ty {
	//	sqlstr = "SELECT * FROM " + tablename + "WHERE `type` = 2 AND `patientid` = ? AND `testtime` >= ? AND `testtime` < ? ORDER BY `testtime` DESC LIMIT 1"
	//}

	time1 := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
	time2 := time.Date(day.Year(), day.Month(), day.Day()+1, 0, 0, 0, 0, loc)
	resultmap, err := fit.MySqlEngine().QueryString(sqlstr, pid, time1.Format("2006-01-02 15:04:05"), time2.Format("2006-01-02 15:04:05"))
	if err != nil {
		fit.Logger().LogError("temp chart ", err)
		return nil, err
	}
	//if len(resultmap) > 0 {
	//	results = append(results, resultmap[0])
	//}

	return resultmap, nil
}

// pc端接口
func PCQueryNRL1(pid, datestr1, datestr2 string, pagenum int) ([]NRL1, error) {
	var mods []NRL1
	var err error
	if datestr2 == "" || datestr1 == "" {
		err = fit.MySqlEngine().Table("NRL5").Where("VAA01 = ?", pid).Limit(9, (pagenum-1)*9).Asc("datetime", "NRL01").Find(&mods)
	} else {
		err = fit.MySqlEngine().Table("NRL5").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum-1)*9).Find(&mods)
	}
	if err != nil {
		return nil, err
	}
	//for key,_ := range mods {
	//	val := mods[key]
	//	mods[key].DateStr = val.DateTime.Format("2006-01-02")
	//}
	return mods, nil
}

func PCQueryNRL1PageCount(pid, datestr1, datestr2 string) (counts int64, err error) {
	if datestr2 == "" || datestr1 == "" {
		counts, err = fit.MySqlEngine().Table("NRL1").Where("Patientid = ?", pid).Count()
	} else {
		counts, err = fit.MySqlEngine().Table("NRL1").Where("VAA01 = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
	}
	return counts, err
}
