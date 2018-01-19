package model

import (
	"fit"
	"github.com/go-xorm/xorm"
	"time"
	"errors"
	"strconv"
)

const (
	Day_Shift       string = "Day_Shift"       //白班
	Night_Shift     string = "Night_Shift"     //晚班
	Graveyard_Shift string = "Graveyard_Shift" //夜班
	Comment_Shift   string = "Comment_Shift"   //内容
)

//交接班
type Succession struct {
	ID          int          `json:"id" xorm:"notnull comment(数据id)"`
	DataTime    FitTime `json:"datatime" xorm:"notnull comment(日期)"`
	Type        int          `json:"type" xorm:"notnull comment('类型1，白班 2，晚班 3，夜班',)"`
	NursingName string       `json:"nursingname" xorm:"notnull comment(护士名称)"`
	NursingID   int          `json:"nursingid" xorm:"notnull comment(护士id)"`
	ClassId     string       `json:"classid" xorm:"notnull comment(科室ID)"`

	NoldPatient    string `json:"noldpatient" xorm:"notnull comment(原病人数)"`
	NnowPatient    string `json:"nnowpatient" xorm:"notnull comment(现病人数)"`
	NintHospital   string `json:"ninthospital" xorm:"notnull comment(入院病人数)"`
	NoutHospital   string `json:"nouthospital" xorm:"notnull comment(出院)"`
	Ninto          string `json:"ninto" xorm:"notnull comment(转入数)"`
	Nout           string `json:"nout" xorm:"notnull comment(转出)"`
	Nsurgery       string `json:"nsurgery" xorm:"notnull comment(手术)"`
	Nchildbirth    string `json:"nchildbirth" xorm:"notnull comment(分娩)"`
	Ncritically    string `json:"ncritically" xorm:"notnull comment(病危)"`
	Ndeath         string `json:"ndeath" xorm:"notnull comment(死亡)"`
	NintensiveCare string `json:"nintensivecare" xorm:"notnull comment(特护)"`
	NprimaryCare   string `json:"nprimarycare" xorm:"notnull comment(一级护理)"`
}

func IputSuccession(session *xorm.Session, strData map[string]string) error {
	var item = &Succession{}

	if v, ok := strData["datatime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
		if err != nil {
			return errors.New("没有datatime")
		} else {
			item.DataTime = FitTime(texttime)
		}
	} else {
		return errors.New("没有datatime")
	}

	if v, ok := strData["type"]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("没有type")
		} else {
			item.Type = i
		}
	} else {
		return errors.New("没有type")
	}

	if v, ok := strData["nursingid"]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("没有type")
		} else {
			item.NursingID = i
		}
	} else {
		return errors.New("nursingid")
	}

	if v, ok := strData["nursingname"]; ok {
		item.NursingName = v
	} else {
		return errors.New("没有nursingname")
	}

	if v, ok := strData["classid"]; ok {
		item.ClassId = v
	} else {
		return errors.New("没有classid")
	}

	if v, ok := strData["noldpatient"]; ok {
		item.NoldPatient = v
	} else {
		return errors.New("没有noldpatient")
	}

	if v, ok := strData["nnowpatient"]; ok {
		item.NnowPatient = v
	} else {
		return errors.New("没有nnowpatient")
	}

	if v, ok := strData["ninthospital"]; ok {
		item.NintHospital = v
	} else {
		return errors.New("没有ninthospital")
	}

	if v, ok := strData["nouthospital"]; ok {
		item.NoutHospital = v
	} else {
		return errors.New("没有nouthospital")
	}

	if v, ok := strData["ninto"]; ok {
		item.Ninto = v
	} else {
		return errors.New("没有ninto")
	}

	if v, ok := strData["nout"]; ok {
		item.Nout = v
	} else {
		return errors.New("没有nout")
	}

	if v, ok := strData["nsurgery"]; ok {
		item.Nsurgery = v
	} else {
		return errors.New("没有nsurgery")
	}

	if v, ok := strData["nchildbirth"]; ok {
		item.Nchildbirth = v
	} else {
		return errors.New("没有nchildbirth")
	}

	if v, ok := strData["ncritically"]; ok {
		item.Ncritically = v
	} else {
		return errors.New("没有ncritically")
	}

	if v, ok := strData["ndeath"]; ok {
		item.Ndeath = v
	} else {
		return errors.New("没有ndeath")
	}

	if v, ok := strData["nintensivecare"]; ok {
		item.NintensiveCare = v
	} else {
		return errors.New("没有nintensivecare")
	}

	if v, ok := strData["nprimarycare"]; ok {
		item.NprimaryCare = v
	} else {
		return errors.New("没有nprimarycare")
	}

	var item1 = &Succession{}

	//var err error

	fit.Logger().LogError("gk dd", "插入验证")

	has, err := session.Table("Succession").Where("datatime = ? and classid = ? and nursingname = ? and Type = ?", item.DataTime.String(), item.ClassId, item.NursingName, item.Type).Get(item1)
	if err != nil {
		fit.Logger().LogError("gk dd", err)
		return err
	}

	if has {
		_, err = session.Where("id = ?", item1.ID).AllCols().Omit("id").Update(item);
		fit.Logger().LogError("gk dd", "更新正确")
	} else {
		_, err = session.Insert(item);
		fit.Logger().LogError("gk dd", "插入正确")
	}

	return err
}

func OutSuccession(sql string, msg ...interface{}) ([]Succession, error) {
	items := make([]Succession, 0)
	err := fit.MySqlEngine().Table("Succession").Where(sql, msg...).Find(&items)
	return items, err
}

//交接班详情
type SuccessionDetails struct {
	ID          int          `json:"id" xorm:"notnull comment(数据id)"`
	DataTime    FitTime      `json:"datatime" xorm:"notnull comment(日期)"`
	ClassId     string       `json:"classid" xorm:"notnull comment(科室ID)"`
	BedID       string       `json:"bedid" xorm:"notnull comment(床位号)"`
	PatientName string       `json:"patientname" xorm:"notnull comment(病人名称)"`
	PatientId   string       `json:"patientid" xorm:"notnull comment(病人id)"`
	Piagnosis   string       `json:"piagnosis" xorm:"notnull comment(诊断)"`

	Typte    int    `json:"typte" xorm:"notnull comment(1入院，2，出院，3，转入，4，转出，5，手术，6，分娩，7，病危，8，死亡，9，特护，10，一级护理)"`
	Comment1 string `json:"comment1" xorm:"notnull comment(备注)"`
	Comment2 string `json:"comment2" xorm:"notnull comment(备注)"`
	Comment3 string `json:"comment3" xorm:"notnull comment(备注)"`
}

func IputSuccessionDetails(session *xorm.Session, strData map[string]string) error {
	var item = &SuccessionDetails{}

	if v, ok := strData["datatime"]; ok {
		texttime, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
		if err != nil {
			return errors.New("没有datatime")
		} else {
			item.DataTime = FitTime(texttime)
		}
	} else {
		return errors.New("没有datatime")
	}

	if v, ok := strData["classid"]; ok {
		item.ClassId = v
	} else {
		return errors.New("没有classid")
	}

	if v, ok := strData["bedid"]; ok {
		item.BedID = v
	} else {
		return errors.New("没有bedid")
	}

	if v, ok := strData["patientname"]; ok {
		item.PatientName = v
	} else {
		return errors.New("没有patientname")
	}

	if v, ok := strData["patientid"]; ok {
		item.PatientId = v
	} else {
		return errors.New("patientid")
	}

	if v, ok := strData["piagnosis"]; ok {
		item.Piagnosis = v
	} else {
		return errors.New("没有piagnosis")
	}

	if v, ok := strData["typte"]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("没有typte")
		} else {
			item.Typte = i
		}
	} else {
		return errors.New("没有typte")
	}

	if v, ok := strData["comment1"]; ok {
		item.Comment1 = v
	} else {
		return errors.New("没有comment1")
	}

	if v, ok := strData["comment2"]; ok {
		item.Comment2 = v
	} else {
		return errors.New("没有comment2")
	}

	if v, ok := strData["comment3"]; ok {
		item.Comment3 = v
	} else {
		return errors.New("没有comment3")
	}

	var item1 = &SuccessionDetails{}

	has, err := session.Table("SuccessionDetails").Where("datatime = ? and classid = ? and BedID =  ?", item.DataTime.String(), item.ClassId, item.BedID).Get(item1)
	if err != nil {
		return err
	}

	if has {
		_, err = session.Where("id = ?", item1.ID).AllCols().Omit("id").Update(item);
	} else {
		_, err = session.Insert(item);
	}

	fit.Logger().LogError("fffffff", "aaaaaaa")

	return err
}

func DeletsSuccessionDetails(session *xorm.Session, id string) error {
	var item = &SuccessionDetails{}
	_, err := session.Where("ID = ?", id).Delete(item);
	return err
}

func OutSuccessionDetails(sql string, msg ...interface{}) ([]SuccessionDetails, error) {
	items := make([]SuccessionDetails, 0)
	err := fit.MySqlEngine().Table("SuccessionDetails").Where(sql, msg...).Find(&items)
	return items, err
}
