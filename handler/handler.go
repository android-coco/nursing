//  Created by JP

package handler

import (
	"fit"
	"nursing/model"
	"errors"
	"fmt"
	"strconv"

	"time"
)

type PCHandlerInterface interface {
	// 获取存储在本地的用户信息，并且会校验登录状态，如果没登录会跳转到登录页
	GetLocalUserinfo(w *fit.Response, r *fit.Request) (userinfo model.UserInfoDup, err error)
	// 在加载View之前校验登录状态，如果没登录会跳转到登录页
	LoadViewSafely(w *fit.Response, r *fit.Request, tplname string) (success bool)
}

/*
func (c Controller) Get(w *fit.Response, r *fit.Request, p fit.Params) {

}

func (c Controller) Post(w *fit.Response, r *fit.Request, p fit.Params) {

}


func (c *Controller) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make(map[string]interface{}, 0))
}

func (c *Controller) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}

	timeNow := time.Now()
	defer func() {
		fmt.Println("***JK: 耗时：",fmt.Sprintf("%.1f",time.Now().Sub(timeNow).Seconds()))
	}()

*/

type PCController struct {
	fit.Controller
}

//type PDAController struct {
//	fit.Controller
//}

/*在加载View之前校验登录状态，如果没登录会跳转到登录页*/
func (c *PCController) LoadViewSafely(w *fit.Response, r *fit.Request, tplname ...string) (success bool) {
	session, err_s := fit.GlobalManager().SessionStart(w, r)
	if err_s != nil || session == nil {
		// Session失效 重新登录
		fit.Logger().LogError("**JP**", "Session失效 重新登录"+err_s.Error())
		c.Redirect(w, r, "/pc/login", 302)
		return false
	} else {
		userinfo := session.Get("UserInfo")
		if userinfo == nil {
			// 未登录
			fit.Logger().LogDebug("**JP**", "未登录")
			c.Redirect(w, r, "/pc/login", 302)
			return false
		} else {
			// 已登录
			fit.Logger().LogDebug("**JP**", "已登录")
			c.LoadView(w, tplname...)
			return true
		}
	}
}

/*获取存储在本地的用户信息，并且会校验登录状态，如果没登录会跳转到登录页*/
func (c *PCController) GetLocalUserinfo(w *fit.Response, r *fit.Request) (userinfo model.UserInfoDup, err error) {
	ses, err_s := fit.GlobalManager().SessionStart(w, r)
	if err_s != nil || ses == nil {
		// Session失效 重新登录
		fit.Logger().LogError("**JP**", "Session失效 重新登录"+err_s.Error())
		c.Redirect(w, r, "/pc/login", 302)
		return userinfo, err_s
	} else {
		temp := ses.Get("UserInfo")
		if temp == nil {
			// 未登录
			fit.Logger().LogDebug("**JP**", "未登录")
			c.Redirect(w, r, "/pc/login", 302)
			return userinfo, errors.New("未登录，无法获取用户信息")
		} else {
			// 已登录
			fit.Logger().LogDebug("**JP**", "已登录")
			userinfo = temp.(model.UserInfoDup)
			return userinfo, nil
		}
	}
}

/*
护理单 PC用
*/
// 护士信息 床位表   病人id  病人信息
func (c *PCController) GetBedsAndUserinfo(w *fit.Response, r *fit.Request, nrlType string) (userinfo model.UserInfoDup, beds []model.PCBedDup, pid string, pInfo model.PCBedDup, isHas bool) {
	// 护士信息
	isHas = false
	var err error
	userinfo, err = c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		fit.Logger().LogError("Error", "参数错误！  user info error", err)
		return
	}

	beds, err = model.QueryDepartmentBeds(userinfo.DepartmentID, false)
	if err != nil {
		fmt.Fprintln(w, "query beds err:", err)
		fit.Logger().LogError("query beds err:", err)
		return
	}

	pid = r.FormValue("pid")
	if pid == "" {
		if len(beds) == 0 {
			fit.Logger().LogError("beds is empty")
			fmt.Fprintln(w, "beds is empty")
			return
		}

		pidnum := beds[0].VAA01
		pid = strconv.FormatInt(pidnum, 10)
		url := "/pc/record/nrl" + nrlType + "?pid=" + pid
		if nrlType == "9" {
			url = "/pc/templist" + "?pid=" + pid
		}
		c.Redirect(w, r, url, 302)
		return userinfo, beds, pid, pInfo, false
	}

	// 病人信息
	for _, val := range beds {
		if strconv.FormatInt(val.VAA01, 10) == pid {
			pInfo = val
			break
		}
	}

	if pInfo.VAA01 == 0 {
		fit.Logger().LogError("pc nrl pInfo is empty")
		fmt.Fprintln(w, "pc nrl pInfo is empty")
		return userinfo, beds, pid, pInfo, false
	}

	return userinfo, beds, pid, pInfo, true
}

// pc 文书 翻页处理时间的页码的
func (c *PCController) GetPageInfo(w *fit.Response, r *fit.Request, nrlType, pid string) (datestr1, datestr2 string, pageindex, pagenum int, err error) {
	// 时间
	date1, errs := strconv.ParseInt(r.FormValue("sdate"), 10, 64)
	date2, erre := strconv.ParseInt(r.FormValue("edate"), 10, 64)
	if errs != nil || erre != nil {
		datestr1 = ""
		datestr2 = ""
	} else {
		datestr1 = time.Unix(date1/1000-60*60*8, 0).Format("2006-01-02 15:04:05")
		datestr2 = time.Unix(date2/1000+60*60*16, 0).Format("2006-01-02 15:04:05")
	}

	// 总条数
	count, errCount := model.PCQUeryNRLPageCount(nrlType, pid, datestr1, datestr2)
	if errCount != nil {
		fmt.Fprintln(w, "nrl list err :", errCount)
		fit.Logger().LogError("nrl page info :", errCount)
		err = errCount
		return
	}

	var peerPage int64 = 9
	switch nrlType {
	case "1":
		peerPage = 9
	case "3":
		peerPage = 9
	case "4":
		peerPage = 9
	case "5":
		peerPage = 5
	case "6":
		peerPage = 4
	case "7":
		peerPage = 8
	case "8":
		peerPage = 9
	default:
		peerPage = 9
	}


	//总页数
	pagenum = int((count-1)/peerPage) + 1
	//当前页数
	index := r.FormValue("num")
	pageindex, errnum := strconv.Atoi(index)
	if errnum != nil {
		pageindex = int(pagenum)
	}
	if pageindex < 1 {
		pageindex = 1
	} else if pageindex > pagenum {
		pageindex = pagenum
	}
	fmt.Println("count:", count, "pageNum:", pagenum, "pageindex:", pageindex)

	return
}
