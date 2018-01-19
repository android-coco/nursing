package model

import (
	"fit"
	"errors"
	"fmt"
	"strconv"
)

/*3-8 公用查询方法*/
/*pda 端， 编辑页 查看某一个文书*/
func QueryNRLWithRid(nrlType, rid string) (nrl interface{}, pid, uid string, err error) {
	tableName := "NRL" + nrlType
	switch nrlType {
	case "3":
		var nr3 NRL3
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr3)
		if err != nil {
			return
		} else {
			nr3.DateStr = nr3.DateTime.ParseDate()
			return nr3, strconv.FormatInt(nr3.PatientId, 10), nr3.NurseId, nil
		}
	case "4":
		var nr4 NRL4
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr4)
		if err != nil {
			return
		} else {
			nr4.DateStr = nr4.DateTime.ParseDate()
			return nr4, strconv.FormatInt(nr4.PatientId, 10), nr4.NurseId, nil
		}
	case "5":
		var nr5 NRL5
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr5)
		if err != nil {
			return
		} else {
			nr5.DateStr = nr5.DateTime.ParseDate()
			return nr5, strconv.FormatInt(nr5.PatientId, 10), nr5.NurseId, nil
		}
	case "6":
		var nr6 NRL6
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr6)
		if err != nil {
			return
		} else {
			nr6.DateStr = nr6.DateTime.ParseDate()
			return nr6, strconv.FormatInt(nr6.PatientId, 10), nr6.NurseId, nil
		}
	case "7":
		var nr7 NRL7
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr7)
		if err != nil {
			return
		} else {
			nr7.DateStr = nr7.DateTime.ParseDate()
			nr7.TimeStr = nr7.DateTime.ParseTime()
			return nr7, strconv.FormatInt(nr7.PatientId, 10), nr7.NurseId, nil
		}
	case "8":
		var nr8 NRL8
		_, err = fit.MySqlEngine().Table(tableName).Where("id = ?", rid).Get(&nr8)
		if err != nil {
			return
		} else {
			nr8.DateStr = nr8.DateTime.ParseDate()
			nr8.TimeStr = nr8.DateTime.ParseTime()
			return nr8, strconv.FormatInt(nr8.PatientId, 10), nr8.NurseId, nil
		}
	default:
		err = errors.New("QueryNRLWithRid : invalid nrltype type")
		return
	}

}
//根据rid查找对应文书
func NRLQueryNRLModel(nrlMod interface{}) (pid int64, uid string, err error) {
	//var rid int64 = 0
	var tableName = ""
	switch obj := nrlMod.(type) {
	case *NRL3:
		//rid = obj.ID
		tableName = "NRL3"
	case *NRL4:
		//rid = obj.ID
		tableName = "NRL4"
	case *NRL5:
		//rid = obj.ID
		tableName = "NRL5"
	case *NRL6:
		//rid = obj.ID
		tableName = "NRL6"
	case *NRL7:
		//rid = obj.ID
		tableName = "NRL7"
	case *NRL8:
		//rid = obj.ID
		tableName = "NRL8"
	default:
		fit.Logger().LogError("hy:", "invalid nrl mod:", obj, nrlMod)
		return 0, "", errors.New("NRLQueryNRLModel : invalid nrlMod")
	}

	_, err = fit.MySqlEngine().Table(tableName).Get(nrlMod)
	// .Where("id = ?", rid)
	if err != nil {
		return 0, "", err
	} else {
		//fmt.Println("nrlmodel:", nrlMod)
		switch obj := nrlMod.(type) {
		case *NRL3:
			pid = obj.PatientId
			uid = obj.NurseId
			obj.DateStr = obj.DateTime.ParseDate()
		case *NRL4:
			pid = obj.PatientId
			uid = obj.NurseId
			obj.DateStr = obj.DateTime.ParseDate()
		case *NRL5:
			pid = obj.PatientId
			uid = obj.NurseId
			obj.DateStr = obj.DateTime.ParseDate()
		case *NRL6:
			pid = obj.PatientId
			uid = obj.NurseId
			obj.DateStr = obj.DateTime.ParseDate()
		case *NRL7:
			pid = obj.PatientId
			uid = obj.NurseId
			obj.DateStr = obj.DateTime.ParseDate()
			obj.TimeStr = obj.DateTime.ParseTime()
		case *NRL8:
			pid = obj.PatientId
			uid = obj.NurseId
			obj.DateStr = obj.DateTime.ParseDate()
			obj.TimeStr = obj.DateTime.ParseTime()
		default:
			err = errors.New("QueryNRLWithRid : invalid nrltype type")
			return 0, "", err
		}

		return pid, uid, nil
	}
}

func NRLInsertData(nrlMod interface{}) (rid int64, err error) {
	var tableName = ""
	var pid int64 = 0
	switch obj := nrlMod.(type) {
	case *NRL3:
		pid = obj.PatientId
		tableName = "NRL3"
	case *NRL4:
		pid = obj.PatientId
		tableName = "NRL4"
	case *NRL5:
		pid = obj.PatientId
		tableName = "NRL5"
	case *NRL6:
		pid = obj.PatientId
		tableName = "NRL6"
	case *NRL7:
		pid = obj.PatientId
		tableName = "NRL7"
	case *NRL8:
		pid = obj.PatientId
		tableName = "NRL8"
	default:
		err = errors.New("NRL InsertData : invalid nrltype type")
		return
	}

	_, err = fit.MySqlEngine().Table(tableName).Insert(nrlMod)
	if err == nil {
		switch obj := nrlMod.(type) {
		case *NRL3:
			rid = obj.ID
		case *NRL4:
			rid = obj.ID
		case *NRL5:
			rid = obj.ID
		case *NRL6:
			rid = obj.ID
		case *NRL7:
			rid = obj.ID
		case *NRL8:
			rid = obj.ID
		default:
			err = errors.New("QueryNRLWithRid : invalid nrltype type")
			return
		}
		fmt.Println("-------- ", pid, nrlMod)
		//time.Sleep(100 * time.Millisecond)
		//fmt.Println("-------- after", pid, object)
		return rid, err
	}
	return 0, err
}

func NRLUpdateData11(nrlMod interface{}) (int64, error) {
	var rid int64 = 0
	var tableName = ""
	switch obj := nrlMod.(type) {
	case *NRL3:
		rid = obj.ID
		tableName = "NRL3"
		fmt.Println("model ------- :", obj)
	case *NRL4:
		rid = obj.ID
		tableName = "NRL4"
	case *NRL5:
		rid = obj.ID
		tableName = "NRL5"
	case *NRL6:
		rid = obj.ID
		tableName = "NRL6"
	case *NRL7:
		rid = obj.ID
		tableName = "NRL7"
	case *NRL8:
		rid = obj.ID
		tableName = "NRL8"
	default:
		fmt.Println("-----obj:", obj, nrlMod)
		return 0, errors.New("NRLQueryNRLModel : invalid nrlMod")
	}
	return fit.MySqlEngine().Table(tableName).ID(rid).AllCols().Omit("ID").Update(nrlMod)
}

func NRLUpdateData(id int64, mod interface{}) (int64, error) {
	return fit.MySqlEngine().ID(id).AllCols().Omit("ID").Update(mod)
}

func NRLDeleteData(id int64, mod interface{}) (int64, error) {
	DeleteNRecords(id)
	return fit.MySqlEngine().ID(id).Delete(mod)
}

func fetchTableName(nrlMod interface{}) (tableName string)  {
	switch obj := nrlMod.(type) {
	case *NRL3:
		tableName = "NRL3"
	case *NRL4:
		tableName = "NRL4"
	case *NRL5:
		tableName = "NRL5"
	case *NRL6:
		tableName = "NRL6"
	case *NRL7:
		tableName = "NRL7"
	case *NRL8:
		tableName = "NRL8"
	default:
		fmt.Println("invalid nrlMod-----obj:", obj, nrlMod)
		return ""
	}
	return
}

func PCNRLQueryNRL3(pid, datestr1, datestr2 string, pagenum int, mods []interface{}) (err error) {
	if pagenum == -1 {
		err = fit.MySqlEngine().Table("NRL3").Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Asc("DateTime").Find(&mods)
	} else {
		if datestr2 == "" || datestr1 == "" {
			err = fit.MySqlEngine().Table("NRL3").Where("PatientId = ?", pid).Limit(9, (pagenum-1)*9).Asc("DateTime").Find(&mods)
		} else {
			err = fit.MySqlEngine().Table("NRL3").Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Limit(9, (pagenum-1)*9).Asc("DateTime").Find(&mods)
		}
	}
	if err != nil {
		return err
	}
	//for key, _ := range mods {
	//	mods[key].DateStr = mods[key].DateTime.ParseDate()
	//}
	return  nil
}


/*PC端*/
//查询页数
func PCQueryNRLPageCount(nrlType, pid, datestr1, datestr2 string) (counts int64, err error) {
	tablename := "NRL" + nrlType
	if nrlType == "1" {
		return -1, nil
		//return 0, errors.New("PCQUeryNRLPageCount: invalid type")
	} else if nrlType == "5" {
		return -1, nil
		//return 0, errors.New("PCQUeryNRLPageCount: invalid type")
	} else {
		if datestr2 == "" || datestr1 == "" {
			counts, err = fit.MySqlEngine().Table(tablename).Where("PatientId = ?", pid).Count()
		} else {
			counts, err = fit.MySqlEngine().Table(tablename).Where("PatientId = ? AND DateTime >= ? AND DateTime < ?", pid, datestr1, datestr2).Count()
		}
	}
	return counts, err
}

func UpdateNRLChcker(nrlType, id, datestr, nurseName string) (err error) {

	sql := ""
	switch nrlType {
	case "2":
		sql = "update NRL2 set NRL39A = ?, NRL39C = ? where ID = ?"
		_, err = fit.MySqlEngine().Exec(sql, datestr, nurseName, id)
	case "6":
		sql = "update NRL6 set NRL15B = ? where ID = ?"
		_, err = fit.MySqlEngine().Exec(sql, nurseName, id)
	case "7":
		sql = "update NRL7 set NRL09B = ? where ID = ?"
		_, err = fit.MySqlEngine().Exec(sql, nurseName, id)
	case "8":
		sql = "update NRL8 set NRL06B = ? where ID = ?"
		_, err = fit.MySqlEngine().Exec(sql, nurseName, id)
	default:
		err = errors.New("API --- update nurse name : invalid type")
		return
	}
	return err
}