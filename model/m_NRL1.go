package model

import (
	"fit"
	"strconv"
	"nursing/utils"
	"sort"
	"time"
	"fmt"
)

//护理记录单 表头
type NRL1Title struct {
	ID    int64 `xorm:"pk autoincr comment(文书id)"`
	PatientId int64 `xorm:"comment(patientid病人id)"`
	BCK01 int   `xorm:"comment(classid科室id)"`

	NRT01  string `xorm:"comment(表头1，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT01V string `xorm:"comment(表头1，自定义内容)"`
	NRT02  string `xorm:"comment(表头2，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT02V string `xorm:"comment(表头2，自定义内容)"`
	NRT03  string `xorm:"comment(表头3，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT03V string `xorm:"comment(表头3，1=意识，自定义内容)"`
	NRT04  string `xorm:"comment(表头4，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT04V string `xorm:"comment(表头4，1=意识，自定义内容)"`
	NRT05  string `xorm:"comment(表头5，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT05V string `xorm:"comment(表头5，1=意识，自定义内容)"`
	NRT06  string `xorm:"comment(表头6，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT06V string `xorm:"comment(表头6，1=意识，自定义内容)"`
	NRT07  string `xorm:"comment(表头7，1=意识，2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分，7=门冬胰岛素，8=瞳孔左mm，9=瞳孔右mm，10=吸氧L/min，11=自定义)"`
	NRT07V string `xorm:"comment(表头7，1=意识，自定义内容)"`
}

//护理记录单
// NurseChat 结构
type NRLData struct {
	ID          int64        `json:"id" xorm:"pk autoincr comment(ID)"`
	HeadType    string       `json:"headType" xorm:"notnull comment(头部id,对应头部类型)"`
	TestTime    fit.JsonTime `json:"testTime" xorm:"notnull comment(测试时间)"`
	DateTimeStr string       `xorm:"-"`
	DateStr     string       `xorm:"-"`
	TimeStr     string       `xorm:"-"`
	SubType     int          `json:"subType,string" xorm:"notnull comment(类型,)"`
	Other       int          `json:"other" xorm:"notnull comment(其他可能选项,)"`
	ValueTitle  string       `xorm:"-"`
	Value       string       `json:"value" xorm:"notnull comment(值)"`
	OtherStr    string       `json:"otherStr" xorm:"comment"`
	PatientId   int64        `json:"patientid" xorm:"notnull comment(病人id)"`
	NurseId     int          `json:"nurseid" xorm:"notnull comment(护士id)"`
	NurseName   string       `json:"nursename" xorm:"notnull comment(护士姓名)"`
}

/*
头部类型:
1体温2脉搏3呼吸4血压5心率6血氧7血糖8体重9身高10皮试12事件13大便次数14其他15入量16出量17病情记录18=出入量统计，
31=自定义1，32=自定义2，33=自定义3，34=自定义4，35=自定义5，36=自定义6，37=自定义7，
*/
type NRLModel struct {
	ID          int64        `xorm:"pk autoincr comment(文书id)"`
	PatientId       int64        `xorm:"comment(patientid病人id)"`
	//BCK01       int          `xorm:"comment(classid科室id)"`
	NurseId      string       `xorm:"comment(NursingId责任护士ID)"`
	NurseName      string       `xorm:"comment(NursingName责任护士签名)"`
	DateTime    fit.JsonTime `json:"dateTime"`
	DateTimeStr string       `xorm:"-"`
	DateStr     string       `xorm:"-"`
	TimeStr     string       `xorm:"-"`
	Mod1        NRLData // 1体温
	Mod2        NRLData // 2脉搏
	Mod3        NRLData // 3呼吸
	Mod4        NRLData // 4血压
	Mod6        NRLData // 6血氧
	Mod7        NRLData // 7血糖
	Mod15       NRLData // 15入量
	Mod16       NRLData // 16出量
	Mod17       NRLData // 17病情记录
	Mod18       NRLData // 18出入量统计

	Mod31 NRLData // 自定义1
	Mod32 NRLData // 自定义2
	Mod33 NRLData // 自定义3
	Mod34 NRLData // 自定义4
	Mod35 NRLData // 自定义5
	Mod36 NRLData // 自定义6
	Mod37 NRLData // 自定义7
}

type IOStatistics struct {
	ID        int64     `xorm:"pk autoincr comment(id)"`
	PatientId     int64     `xorm:"comment(patientid病人id)"`
	//BCK01     int       `xorm:"comment(classid科室id)"`
	NurseId    string    `xorm:"comment(NursingId责任护士ID)"`
	NurseName    string    `xorm:"comment(NursingName责任护士签名)"`
	DateTime1 time.Time `xorm:"comment(记录时间)"`
	DateTime2 time.Time `xorm:"comment(记录时间)"`
	DataType  string    `xorm:"comment(1=同步到体温单，0=护理记录单)"`

	// 入量 1输液2饮食3饮水4其他
	IntakeA     string `xorm:"comment(入量输液)"`
	IntakeB     string `xorm:"comment(入量饮食)"`
	IntakeC     string `xorm:"comment(入量饮水)"`
	IntakeD     string `xorm:"comment(入量其他)"`
	IntakeDV    string `xorm:"comment(入量其他,描述)"`
	IntakeTotal string `xorm:"comment(总入量)"`

	// 出量 1尿量2大便3其他
	OutputA     string `xorm:"comment(出量尿量)"`
	OutputB     string `xorm:"comment(出量大便)"`
	OutputC     string `xorm:"comment(出量其他)"`
	OutputCV    string `xorm:"comment(出量其他,描述)"`
	OutputTotal string `xorm:"comment(总出量)"`
}

//NRLData 排序
type Datas []NRLData

func (mods Datas) Len() int {
	return len(mods)
}

func (mods Datas) Less(i, j int) bool {
	return mods[i].TestTime.Before(mods[j].TestTime)
}

func (mods Datas) Swap(i, j int) {
	mods[i], mods[j] = mods[j], mods[i]
}

func appendData(src ...NRLData) (mods Datas) {
	mods = append(mods, src...)
	return mods
}

//NRLModel 排序
type DataModels []NRLModel

func (mods DataModels) Len() int {
	return len(mods)
}

func (mods DataModels) Less(i, j int) bool {
	//date1 := mods[i].DateStr + " " + mods[i].TimeStr
	//date2 := mods[j].DateStr + " " + mods[j].TimeStr
	//return date1 < date2
	return mods[i].DateTime.Before(mods[j].DateTime)
}

func (mods DataModels) Swap(i, j int) {
	mods[i], mods[j] = mods[j], mods[i]
}

func appendDataModel(src ...NRLModel) (mods DataModels) {
	mods = append(mods, src...)
	return mods
}

/*
护理记录单
*/
func GetNRL1Data(pid, datestr1, datestr2 string) ([]NRLModel, error) {
	// 常规数据
	//fit.MySqlEngine().ShowSQL(true)
	var mods []NRLData
	var err error
	if datestr1 == "" || datestr2 == "" {
		err = fit.MySqlEngine().Table("NurseChat").Where("patientid = ?", pid).In("HeadType", "1", "2", "3", "4", "6", "7", "15", "16", "17", "31", "32", "33", "34", "35", "36", "37").Asc("TestTime", "NurseId").Find(&mods)
	} else {
		err = fit.MySqlEngine().Table("NurseChat").Where("patientid = ? and TestTime >= ? and TestTime < ?", pid, datestr1, datestr2).In("HeadType", "1", "2", "3", "4", "6", "7", "15", "16", "17", "31", "32", "33", "34", "35", "36", "37").Asc("TestTime", "NurseId").Find(&mods)
	}
	//fit.MySqlEngine().ShowSQL(false)
	if err != nil {
		fit.Logger().LogError("temp chart ", err)
		return nil, err
	}
	parseNRLData(mods)

	//变化表头部分数据
	VAA01, _ := utils.Int64Value(pid)
	nrl1Title := NRL1Title{PatientId: VAA01}

	// 自定义表头
	nrl1Title.PCQueryNRL1Title()
	data1, err1 := NRLRelatesToMarkSheet(nrl1Title.NRT01, "31", pid, datestr1, datestr2)
	data2, err2 := NRLRelatesToMarkSheet(nrl1Title.NRT02, "32", pid, datestr1, datestr2)
	data3, err3 := NRLRelatesToMarkSheet(nrl1Title.NRT03, "33", pid, datestr1, datestr2)
	data4, err4 := NRLRelatesToMarkSheet(nrl1Title.NRT04, "34", pid, datestr1, datestr2)
	data5, err5 := NRLRelatesToMarkSheet(nrl1Title.NRT05, "35", pid, datestr1, datestr2)
	data6, err6 := NRLRelatesToMarkSheet(nrl1Title.NRT06, "36", pid, datestr1, datestr2)
	data7, err7 := NRLRelatesToMarkSheet(nrl1Title.NRT07, "37", pid, datestr1, datestr2)
	checkerr("NRLRelatesToMarkSheet", err1, err2, err3, err4, err5, err6, err7)

	mods = append(mods, data1...)
	mods = append(mods, data2...)
	mods = append(mods, data3...)
	mods = append(mods, data4...)
	mods = append(mods, data5...)
	mods = append(mods, data6...)
	mods = append(mods, data7...)

	var data = appendData(mods...)
	sort.Sort(data)
	// 组合成DataModel
	result1 := formatNRLData(data)

	// 出入量统计部分数据
	ioData, errio := NRLRelatesToIO(pid, datestr1, datestr2)
	checkerr("err relates to io :", errio)

	result1 = append(result1, ioData...)

	var result = appendDataModel(result1...)
	sort.Sort(result)

	//for _, val := range result {
	//	fmt.Println("--- date :", val.DateStr, val.TimeStr)
	//}
	return result, nil
}

func parseNRLData(mods []NRLData) {
	for key, _ := range mods {
		val := mods[key]

		mods[key].DateTimeStr = val.TestTime.NormParse()
		mods[key].DateStr = val.TestTime.ParseDate()
		mods[key].TimeStr = val.TestTime.ParseTime()
		switch val.HeadType {
		case "15":
			switch val.SubType {
			case 1:
				mods[key].ValueTitle = "输液"
			case 2:
				mods[key].ValueTitle = "饮食"
			case 3:
				mods[key].ValueTitle = "饮水"
			case 4:
				//mods[key].ValueTitle = "其他"
				mods[key].ValueTitle = val.OtherStr
			default:
				mods[key].ValueTitle = ""
			}
		case "16":
			switch val.SubType {
			case 1:
				mods[key].ValueTitle = "尿量"
			case 2:
				mods[key].ValueTitle = "大便"
			case 3:
				//mods[key].ValueTitle = "其他"
				mods[key].ValueTitle = val.OtherStr
			default:
				mods[key].ValueTitle = ""
			}
		default:
		}
	}
}

func formatNRLData(mods []NRLData) []NRLModel {
	var results []NRLModel
	for index := 0; index < len(mods); index++ {
		mod := mods[index]
		//fmt.Println("--- time  :", mod.TestTime, mod.DateTimeStr)
		if index == 0 {
			var resultModel NRLModel
			resultModel.DateTime = mod.TestTime
			resultModel.DateTimeStr = mod.DateTimeStr
			resultModel.DateStr = mod.DateStr
			resultModel.TimeStr = mod.TimeStr
			resultModel.NurseId = strconv.Itoa(mod.NurseId)
			resultModel.NurseName = mod.NurseName
			mateNRLData(&mod, &resultModel)
			results = append(results, resultModel)

		} else {
			oldresultModel := &results[len(results)-1]
			//if mod.DateStr == oldresultModel.DateStr && mod.TimeStr == oldresultModel.TimeStr && strconv.Itoa(mod.NurseId) == oldresultModel.BCE01A {
			if mod.TestTime == oldresultModel.DateTime && strconv.Itoa(mod.NurseId) == oldresultModel.NurseId { // 测量时间相同，护士相同
				if oldresultModel.Mod15.HeadType == mod.HeadType || oldresultModel.Mod16.HeadType == mod.HeadType { // 同一测量时间下， 出入量多条
					var newresultModel NRLModel
					newresultModel.DateTime = mod.TestTime
					newresultModel.DateTimeStr = mod.DateTimeStr
					newresultModel.DateStr = mod.DateStr
					newresultModel.TimeStr = mod.TimeStr
					newresultModel.NurseId = strconv.Itoa(mod.NurseId)
					newresultModel.NurseName = mod.NurseName
					mateNRLData(&mod, &newresultModel)
					results = append(results, newresultModel)
				} else {
					mateNRLData(&mod, oldresultModel)
				}
			} else {
				var newresultModel NRLModel
				newresultModel.DateTime = mod.TestTime
				newresultModel.DateTimeStr = mod.DateTimeStr
				newresultModel.DateStr = mod.DateStr
				newresultModel.TimeStr = mod.TimeStr
				newresultModel.NurseId = strconv.Itoa(mod.NurseId)
				newresultModel.NurseName = mod.NurseName
				mateNRLData(&mod, &newresultModel)
				results = append(results, newresultModel)
			}

		}
	}
	return results
}

func mateNRLData(mod *NRLData, resultModel *NRLModel) {
	switch mod.HeadType {
	case "1":
		resultModel.Mod1 = *mod
	case "2":
		resultModel.Mod2 = *mod
	case "3":
		resultModel.Mod3 = *mod
	case "4":
		resultModel.Mod4 = *mod
	case "6":
		resultModel.Mod6 = *mod
	case "7":
		resultModel.Mod7 = *mod
	case "15":
		resultModel.Mod15 = *mod
	case "16":
		resultModel.Mod16 = *mod
	case "17":
		resultModel.Mod17 = *mod
	case "18":
		resultModel.Mod18 = *mod
	case "31":
		resultModel.Mod31 = *mod
	case "32":
		resultModel.Mod32 = *mod
	case "33":
		resultModel.Mod33 = *mod
	case "34":
		resultModel.Mod34 = *mod
	case "35":
		resultModel.Mod35 = *mod
	case "36":
		resultModel.Mod36 = *mod
	case "37":
		resultModel.Mod37 = *mod
	default:
		fit.Logger().LogError("FormatNRLData err , invalid type:", mod.HeadType)
	}

}

// 护理记录单
// 表头
func (m *NRL1Title) PCQueryNRL1Title() (err error) {
	_, err = fit.MySqlEngine().Table("NRL1Title").Where("PatientId = ?", m.PatientId).Get(m)
	return
}

func (m *NRL1Title) PCUpdateNRT1Title() error {
	has, err := fit.MySqlEngine().Table("NRL1Title").Where("PatientId = ?", m.PatientId).Exist()
	if err != nil {
		return err
	}
	if has {
		var mod NRL1Title
		_, err = fit.MySqlEngine().Table("NRL1Title").Where("PatientId = ?", m.PatientId).Get(&mod)
		if err != nil {
			return err
		}

		if mod.NRT01 == "11" {
			m.NRT01V = ""
		}
		if mod.NRT02 == "11" {
			m.NRT02V = ""
		}
		if mod.NRT03 == "11" {
			m.NRT03V = ""
		}
		if mod.NRT04 == "11" {
			m.NRT04V = ""
		}
		if mod.NRT05 == "11" {
			m.NRT05V = ""
		}
		if mod.NRT06 == "11" {
			m.NRT06V = ""
		}
		if mod.NRT07 == "11" {
			m.NRT07V = ""
		}
		_, err = fit.MySqlEngine().Table("NRL1Title").ID(mod.ID).Omit("ID", "BCK01", "PatientId").Update(m)

	} else {
		_, err = fit.MySqlEngine().Table("NRL1Title").Insert(m)

	}
	if err != nil {
		return err
	}

	return nil
}

/*
护理记录单 表头关联评分
//2=BADL评分，3=DVT评分，4=跌倒评分，5=压疮评分，6=疼痛评分
*/
func NRLRelatesToMarkSheet(nrlType, headType, pid, datestr1, datestr2 string) (results []NRLData, err error) {
	var tablename = ""
	switch nrlType {
	case "2":
		tablename = " NRL3"
	case "3":
		tablename = " NRL4"
	case "4":
		tablename = " NRL7"
	case "5":
		tablename = " NRL6"
	case "6":
		tablename = " NRL8"
	default:
		return
	}

	var resultmap []map[string]string
	if datestr1 == "" || datestr2 == "" {
		sqlstr := "select ID,BCK01,PatientId,BCE01A,BCE03A,DateTime,Score from" + tablename + " where PatientId = ?"
		resultmap, err = fit.MySqlEngine().QueryString(sqlstr, pid)
	} else {
		sqlstr := "select ID,BCK01,PatientId,BCE01A,BCE03A,DateTime,Score from" + tablename + " where PatientId = ? and DateTime >= ? and DateTime < ?"
		resultmap, err = fit.MySqlEngine().QueryString(sqlstr, pid, datestr1, datestr2)

	}

	if err != nil {
		fit.Logger().LogError("NRL1", "relates to mark sheet error", err)
		return
	}

	for _, val := range resultmap {
		id, _ := utils.Int64Value(val["ID"])
		pid, _ := utils.Int64Value(val["PatientId"])
		uid, _ := strconv.Atoi(val["BCE01A"])
		datetime, _ := time.ParseInLocation("2006-01-02 15:04:05", val["DateTime"], time.Local)
		datestr := utils.Substr(val["DateTime"], 0, 10)
		timestr := utils.Substr(val["DateTime"], 11, 5)
		if timestr == "00:00" {
			timestr = ""
		}

		nrldata := NRLData{
			ID:          id,
			HeadType:    headType,
			NurseId:     uid,
			NurseName:   val["BCE03A"],
			PatientId:   pid,
			Value:       val["Score"],
			TestTime:    fit.JsonTime(datetime),
			DateTimeStr: datetime.Format("2006-01-02 15:04:05"),
			DateStr:     datestr,
			TimeStr:     timestr,
		}
		if nrldata.Value != "" {
			results = append(results, nrldata)
		}
	}
	return
}

/*
护理记录单 关联出入量统计
*/
func NRLRelatesToIO(pid, datestr1, datestr2 string) (results []NRLModel, err error) {

	var ioModels []IOStatistics
	if datestr1 == "" || datestr2 == "" {
		err = fit.MySqlEngine().Table("IOStatistics").Where("PatientId = ?", pid).Find(&ioModels)
	} else {
		err = fit.MySqlEngine().Table("IOStatistics").Where("PatientId = ? and DateTime1 >= ? and DateTime1 < ?", pid, datestr1, datestr2).Find(&ioModels)
	}
	if err != nil {
		fit.Logger().LogError("NRL1", "relates to IO sheet error", err)
		return
	}
	for _, val := range ioModels {
		time1 := val.DateTime1
		time2 := val.DateTime2
		minutes := time2.Sub(time1).Minutes()
		str := fmt.Sprintf("%s - %s  共%.0f小时%d分 小结", time1.Format("15:04"), time2.Format("15:04"), minutes/60, int(minutes)%60)
		nrldata := NRLModel{
			ID:       val.ID,
			PatientId:    val.PatientId,
			//BCK01:    val.BCK01,
			NurseId:   val.NurseId,
			NurseName:   val.NurseName,
			DateTime: fit.JsonTime(time1),
			DateStr:  val.DateTime1.Format("2006-01-02"),
			TimeStr:  str,

			Mod1: NRLData{
				HeadType: "18",
				TestTime: fit.JsonTime(time1),
				Value:    val.IntakeA,
			},
			Mod2: NRLData{
				TestTime: fit.JsonTime(time2),
				Value:    val.IntakeB,
			},
			Mod3: NRLData{
				Value: val.IntakeC,
			},
			Mod4: NRLData{
				Value: val.IntakeD,
			},
			Mod6: NRLData{
				Value: val.IntakeTotal,
			},
			Mod7: NRLData{
				Value: val.OutputA,
			},
			Mod15: NRLData{
				Value: val.OutputB,
			},
			Mod16: NRLData{
				Value: val.OutputC,
			},
			Mod17: NRLData{
				Value: val.OutputTotal,
			},
		}
		results = append(results, nrldata)
	}
	return
}

/*
出入量统计
*/
func PCQueryNRLIntakeOutputData(pid, ty, datestr1, datestr2 string) ([]NRLData, error) {
	var mods []NRLData
	err := fit.MySqlEngine().Table("NurseChat").Where("patientid = ? AND HeadType = ? AND TestTime >= ? AND TestTime < ?", pid, ty, datestr1, datestr2).Asc("TestTime").Find(&mods)
	if err != nil {
		fit.Logger().LogError("temp chart ", err)
		return nil, err
	}
	parseNRLData(mods)
	for _, mod := range mods {
		//fmt.Println(mod)
		mod.DateStr = mod.TestTime.IOParse()
	}

	return mods, nil
}

func (m *IOStatistics) InsertData() (int64, error) {
	return fit.MySqlEngine().Table("IOStatistics").Insert(m)
}

func (m *IOStatistics) DeleteData(id int64) (int64, error) {
	return fit.MySqlEngine().ID(id).Delete(m)
}

func checkerr(tag string, err ...error) bool {
	var flag = false
	for _, val := range err {
		if val != nil {
			fit.Logger().LogError(tag, val)
			flag = true
		}
	}
	return flag
}
