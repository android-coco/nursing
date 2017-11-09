//  Created by JP

package handler

import (
	"fit"
	"nursing/model"
	"errors"
	"fmt"
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
		fit.Logger().LogError("**JP**", "Session失效 重新登录" + err_s.Error())
		c.Redirect(w,r,"/pc/login",302)
		return false
	} else {
		userinfo := session.Get("UserInfo")
		if userinfo == nil {
			// 未登录
			fit.Logger().LogDebug("**JP**", "未登录")
			c.Redirect(w,r,"/pc/login",302)
			return false
		} else {
			// 已登录
			fit.Logger().LogDebug("**JP**", "已登录")
			c.LoadView(w,tplname...)
			return true
		}
	}
}



/*获取存储在本地的用户信息，并且会校验登录状态，如果没登录会跳转到登录页*/
func (c *PCController)GetLocalUserinfo(w *fit.Response, r *fit.Request) (userinfo model.UserInfoDup, err error) {
	ses, err_s := fit.GlobalManager().SessionStart(w, r)
	if err_s != nil || ses == nil {
		// Session失效 重新登录
		fit.Logger().LogError("**JP**", "Session失效 重新登录" + err_s.Error())
		c.Redirect(w,r,"/pc/login",302)
		return userinfo , err_s
	} else {
		temp := ses.Get("UserInfo")
		if temp == nil {
			// 未登录
			fit.Logger().LogDebug("**JP**", "未登录")
			c.Redirect(w,r,"/pc/login",302)
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
func (c PCController) NRLRecord(w *fit.Response, r *fit.Request, p fit.Params) (userinfo model.UserInfoDup, beds []model.PCBedDup, err error) {
	// 护士信息
	userinfo, err = c.GetLocalUserinfo(w, r)
	if err != nil {
		fmt.Fprintln(w, "参数错误！  user info error", err)
		return userinfo, nil,err
	}

	beds, err = model.QueryDepartmentBeds(userinfo.DepartmentID, false)
	if err != nil {
		fit.Logger().LogError("pc nrl2", err)
		return userinfo, nil, err
	}

	return userinfo, beds, nil
}