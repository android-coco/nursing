package handler

import (
	"fit"
	"fmt"
	"nursing/model"
)

type PCNRL5Controller struct {
	PCController
}

type APNModel struct {
	ModelA model.NRL5
	ModelP model.NRL5
	ModelN model.NRL5
}

func (c PCNRL5Controller) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) {
	// 护士信息 床位表 病人id  病人信息
	userinfo, beds, pid, pInfo, has := c.GetBedsAndUserinfo(w, r, "5")
	if !has {
		return
	}
	// 起止时间  页码
	datestr1, datestr2, pageindex, pagenum, err := c.GetPageInfo(w, r, "5", pid)
	if err != nil {
		return
	}


	// 护理单
	mods, err13 := model.PCQueryNRL5(pid, datestr1, datestr2, pageindex)

	var list []APNModel
	arrlen := len(mods)
	length := (arrlen - 1) / 3 + 1

	for ii := 0; ii < length ; ii++ {
		var apnmol APNModel
		if ii == length - 1 {
			fmt.Println("-------------")
			if arrlen % 3 == 1 {
				modP := model.NRL5{
					NRL01:"2",
				}
				modN := model.NRL5{
					NRL01:"3",
				}
				apnmol.ModelA = mods[ii * 3]
				apnmol.ModelP = modP
				apnmol.ModelN = modN
			} else if arrlen % 3 == 2 {
				modN := model.NRL5{
					NRL01:"3",
				}
				apnmol.ModelA = mods[ii * 3]
				apnmol.ModelP = mods[ii * 3 + 1]
				apnmol.ModelN = modN
			}

		} else {
			fmt.Println("############")
			apnmol.ModelA = mods[ii * 3]
			apnmol.ModelP = mods[ii * 3 + 1]
			apnmol.ModelN = mods[ii * 3 + 2]
		}
		list = append(list, apnmol)
	}


	if err13 != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return
	}
	fmt.Printf("mods %+v\n %d\n\n", list, len(list))


	c.Data = fit.Data{
		"Userinfo":  userinfo, // 护士信息
		"PInfo":     pInfo,    // 病人信息
		"Beds":      beds,     // 床位list
		"NRLList":   list,
		"PageNum":   pagenum,
		"PageIndex": pageindex,
		"Menuindex": "7-5",
	}


	c.LoadViewSafely(w, r, "pc/v_nrl5.html", "pc/header_side.html", "pc/header_top.html")
}

