package handler

import (
	"fit"
	"fmt"
	"nursing/model"
	"strconv"
	"encoding/json"
	"strings"
	"time"
)

type PNRL2Controller struct {
	PCController
}

func (c PNRL2Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息
	userinfo, err := c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}
	beds, err := model.QueryDepartmentBeds(userinfo.DepartmentID, false)

	//fmt.Printf("bed1 : %+v\n", beds)
	if err != nil {
		fit.Logger().LogError("pc nrl2", err)
	}

	// 病人id  病人信息
	var pInfo model.PCBedDup
	pid := r.FormValue("pid")
	if pid == "" {
		pidnum := beds[0].VAA01
		pid = strconv.Itoa(pidnum)

		url := "/pc/record/nrl2?pid=" + pid
		c.Redirect(w, r, url, 302)
		return
	}

	//fmt.Println("pid", pid)
	// 病人信息
	for _, val := range beds {
		if strconv.Itoa(val.VAA01) == pid {
			pInfo = val
			break
		}
	}
	if pInfo.VAA01 == 0 {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	// 护理单
	flag, errExist := model.IsExistNRL2(pid)
	if errExist != nil {
		fit.Logger().LogError("pc nr2 is exist nrl2?", err)
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}

	var nrl2 model.NRL2
	if !flag {
		nrl2 = model.NRL2{}
	} else {
		var errn2 error
		nrl2, errn2 = model.QueryNRL2WithPid(pid)
		if errn2 != nil {
			fit.Logger().LogError("pc nr2 query nrl2:", err)
			fmt.Fprintln(w, "参数错误！  user info error", err)
			return
		}
	}
	// 过敏史-过敏源-NRL06-Json字符串，key：对应index，value：对应内容
	NRL06A := ""
	NRL06B := ""
	if nrl2.NRL06 == "2" {
		by := []byte(nrl2.NRL06A)
		anyObj := make(map[string]string, 0)
		json.Unmarshal(by, &anyObj)
		for k, v := range anyObj {
			NRL06A = k
			NRL06B = v
		}
	}

	// 排便次数- "n,m"，n,m代表2个空的数值，即n次/天，1次/m天
	NRL18A := ""
	NRL18B := ""
	if nrl2.NRL18 != "" {
		slice := strings.Split(nrl2.NRL18, ",")
		length := len(slice)
		if length == 1 {
			NRL18A = slice[0]
		} else if length == 2 {
			NRL18A = slice[0]
			NRL18B = slice[1]
		}
	}
	// 拆分护理单录入时间
	tempTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nrl2.NRL38, time.Local)
	NRL38A := tempTime.Format("2006-01-02")
	NRL38B := tempTime.Format("15:04")

	tempTime, _ = time.ParseInLocation("2006-01-02 15:04:05", nrl2.NRL39A, time.Local)
	NRL39B := tempTime.Format("2006-01-02")
	NRL39C := tempTime.Format("15:04")

	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRL":       nrl2,
		"Menuindex": "7-2",
		"NRL06A":    NRL06A, // 过敏源的index
		"NRL06B":    NRL06B, // 过敏源的补充内容
		"NRL18A":    NRL18A, // n次/天
		"NRL18B":    NRL18B, // 1次/m天
		"NRL38A":    NRL38A, // 录入护理单的年月日
		"NRL38B":    NRL38B, // 录入护理单的时分
		"NRL39B":    NRL39B,
		"NRL39C":    NRL39C,
	}

	c.LoadViewSafely(w, r, "pc/v_nrl2.html", "pc/header_side.html", "pc/header_top.html")
}
