//  Created by JP

package handler

import (
	"fit"
	"nursing/model"
	"errors"
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

/*
非管理员加载通道
在加载View之前校验登录状态，如果没登录会跳转到登录页
*/
func (c *PCController) LoadViewSafely(w *fit.Response, r *fit.Request, tplname ...string) (success bool) {
	return c.LoadViewForAdministrator(false, w, r, tplname...)
	//session, err_s := fit.GlobalManager().SessionStart(w, r)
	//if err_s != nil || session == nil {
	//	// Session失效 重新登录
	//	fit.Logger().LogError("**JP**", "Session失效 重新登录"+err_s.Error())
	//	c.Redirect(w, r, "/pc/login", 302)
	//	return false
	//} else {
	//	userinfo := session.Get("UserInfo")
	//	if userinfo == nil {
	//		// 未登录
	//		fit.Logger().LogDebug("**JP**", "未登录")
	//		c.Redirect(w, r, "/pc/login", 302)
	//		return false
	//	} else {
	//		account := userinfo.(model.UserInfoDup)
	//		if account.Authority >= 2 {
	//			fit.Logger().LogDebug("**JP**", "禁止访问")
	//			c.Redirect(w, r, "/pc/account/manage", 302)
	//			return false
	//		} else {
	//			// 已登录
	//			c.LoadView(w, tplname...)
	//			return true
	//		}
	//	}
	//}
}

/*
可授权管理员账号加载
在加载View之前校验登录状态，如果没登录会跳转到登录页，管理员账号只能进专属页面
*/
func (c *PCController) LoadViewForAdministrator(administrator bool,w *fit.Response, r *fit.Request, tplname ...string) (success bool) {
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
			account := userinfo.(model.UserInfoDup)
			if administrator == true {
				if account.Authority >= 2 {
					// 已登录
					c.LoadView(w, tplname...)
					return true
				} else {
					fit.Logger().LogDebug("**JP**", "禁止访问")
					c.Redirect(w, r, "/pc/home", 302)
					return false
				}
			} else {
				if account.Authority < 2 {
					// 已登录
					c.LoadView(w, tplname...)
					return true
				} else {
					fit.Logger().LogDebug("**JP**", "禁止访问")
					c.Redirect(w, r, "/pc/account/manage", 302)
					return false
				}
			}
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
			userinfo = temp.(model.UserInfoDup)
			return userinfo, nil
		}
	}
}



func (c *PCController) CheckingTimeFormat(param string) (res time.Time,err error) {
	res, err = time.Parse("2006-01-02 15:04:05", param)
	return
}

