package handler

import (
	"fmt"
	"fit"
)

type MainController struct {
	fit.Controller
}

type X struct{
	Name    string `json:"name"`
	Datatime string `json:"datatime"`
}

type User1 struct {
	Name string
	Datatime string
	Name1 string
}

type User2 struct {
	Name string `json:"name"`
	Datatime fit.JsonTime `json:"datatime"`
}

func (c MainController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	//t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	//fit.SQLServerEngine().Query("insert into [dbo].[user](name,age,datatime,name1) VALUES(?,?,?,?)","游浩",29,"2019-06-15 08:37:18","捡垃圾")

	/**    SQL Server ----start ============================================= */
	//myUser1 := User1{Name:"胡杨",Datatime:"2020-10-10 12:12:12",Name1:"胡杨"}
	//inserId,err := fit.SQLServerEngine().Insert(myUser1)
	//fit.Logger().LogInfo("MainController:",inserId, err)

	//results, err := fit.SQLServerEngine().Query("select * from [dbo].[user1]")
	//fit.Logger().LogInfo("MainController:",string(results[0]["name1"]), err)

	//results := make([]User2,0)
	//err1 := fit.SQLServerEngine().SQL("select * from [dbo].[user1]").Find(&results)
	//fit.Logger().LogInfo("MainController:",results[0].Name, err1)

	//c.JsonData.Result = 0
	//c.JsonData.ErrorMsg = "OK"
	//c.JsonData.Datas = results
	/**    SQL Server ---- end  ===================================================== */


	/**    MySql ----start ============================================= */

	x := X{Name:"胡杨",Datatime:"2020-10-10 12:12:12"}
	xid,err := fit.Engine().Insert(x)
	fit.Logger().LogInfo("MainController:",xid, err)

	results1 := make([]X,0)
	err2 := fit.Engine().SQL("select * from x").Find(&results1)
	fit.Logger().LogInfo("MainController:",results1, err2)
	c.JsonData.Result = 0
	c.JsonData.ErrorMsg = "OK"
	c.JsonData.Datas = results1
	/**    MySql ----end ============================================= */
	//c.Data = fit.Data{"Title":"标题"}
	//c.LoadView(w,"tmpl.html")
	//fmt.Fprint(w.Writer(), "WelcomeControllerGet!\n"+r.FormValue("name"))
}
func (c MainController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	fmt.Fprint(w.Writer(), "WelcomeControllerPost!\n"+r.FormValue("name"))
}



