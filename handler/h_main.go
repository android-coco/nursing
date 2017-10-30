package handler

import (
	"fmt"
	"fit"
)

type MainController struct {
	fit.Controller
}

type Sdata struct {
	Name string `json:"name"`
}

type X struct {
	Name     string `json:"name"`
	Datatime string `json:"datatime"`
}

type User1 struct {
	Name     string
	Datatime string
	Name1    string
	VAA01    int
	VAA05    string
}

type User2 struct {
	Name     string       `json:"name"`
	Datatime fit.JsonTime `json:"datatime"`
	VAA01    int          `json:"vaa_01"`
	VAA05    string       `json:"vaa_05"`
}

type VAA struct {
	VAA01 int
	VAA02 string
	VAA03 string
	VAA04 string
	VAA05 string
}

func (c MainController) GetFunc(w *fit.Response, r *fit.Request, p fit.Params) {
	//fmt.Fprint(w.Writer(), "OK")
	//fit.MySqlEngine().ShowSQL(true)
	//m := model.Warn{ClassId: "123", WarnTime: "2017-10-09 08:01,2017-10-09 08:00:00"}
	//id, err := fit.MySqlEngine().SQL("DELETE FROM Warn WHERE WarnTime IN (" + m.WarnTime + ") AND classid = " + m.ClassId).Delete(&m)
	//fmt.Fprintln(w, id, err)

	//x11 := model.NursingRecords{Updated: "2006-01-02 15:04:05", NursType: 1, NursingId: "1", NursingName: "1", ClassId: "1", PatientId: 1, RecordId: 1}
	//z,err8 := model.InsertNRecords(x11)
	//z,err8 := model.UpadteNRecords(1)
	//z,err8 := model.QueryNRecords(1)
	//z,err8 := model.QueryNRecordsByTypeAndTime(0,"2017-10-11","2017-10-11")
	//fmt.Println(z,err8)
	//导出excel表格
	//titles := []string{"a", "b", "c", "d"}
	//datas := [][]string{{"a1", "b1", "c1", "d1"}, {"a2", "b2", "c2", "d2"}, {"a3", "b3", "c3", "d3"}, {"a4", "b4", "c4", "d4"}}
	//error := fit.ExportExcel(w, titles, datas, "My")
	//if error != nil {
	//	fmt.Fprint(w.Writer(), "Fiale"+error.Error())
	//} else {
	//	fmt.Fprint(w.Writer(), "OK")
	//}
	c.Redirect(w,r,"pc/home",302)
	//c.LoadView(w, "v_index.html")
}

//func (c MainController)PostFunc(w *fit.Response, r *fit.Request, p fit.Params)  {
//	fmt.Fprintln(w,"PostFunc")
//}
func (c MainController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)

	//t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	//fit.SQLServerEngine().Query("insert into [dbo].[user](name,age,datatime,name1) VALUES(?,?,?,?)","游浩",29,"2019-06-15 08:37:18","捡垃圾")

	/*=======================================Oracle===========================*/
	//results := make([]Sdata, 0)
	//err := fit.OracleEngine().SQL("select * from SDATA").Find(&results)
	//fit.Logger().LogError("Mian", err)

	/**    SQL Server ----start ============================================= */
	//myUser1 := User1{Name:"胡杨",Datatime:"2020-10-10 12:12:12",Name1:"胡杨",VAA01:12,VAA05:"胡杨"}
	//inserId,err := fit.SQLServerEngine().Insert(myUser1)
	//fit.Logger().LogInfo("MainController:",inserId, err)

	//results, err := fit.SQLServerEngine().Query("select * from user1")
	//fit.Logger().LogInfo("MainController:",results, err)

	//results := make([]VAA,0)
	//err1 := fit.SQLServerEngine().SQL("select * from VAA1").Find(&results)
	//fit.Logger().LogInfo("MainController:",results[0].VAA05, err1)

	//c.JsonData.Result = 0
	//c.JsonData.ErrorMsg = "OK"
	//c.JsonData.Datas = results
	/**    SQL Server ---- end  ===================================================== */

	/**    MySql ----start ============================================= */
	//x := X{Name: "胡杨", Datatime: "2020-10-10 12:12:12"}
	//xid, err := fit.MySqlEngine().Insert(x)
	//fit.Logger().LogInfo("MainController:", xid, err)
	//
	//results1 := make([]X, 0)
	//err2 := fit.MySqlEngine().SQL("select * from x").Find(&results1)
	//fit.Logger().LogInfo("MainController:", results1, err2)
	//c.JsonData.Result = 0
	//c.JsonData.ErrorMsg = "OK"
	//c.JsonData.Datas = results1
	/**    MySql ----end ============================================= */

	//c.Data = fit.Data{"Title":"标题"}
	//c.LoadView(w,"tmpl.html")
	//fmt.Fprint(w.Writer(), "WelcomeControllerGet!\n"+r.FormValue("name"))

}
func (c MainController) Post(w *fit.Response, r *fit.Request, p fit.Params) {
	fmt.Fprint(w.Writer(), "WelcomeControllerPost!\n"+r.FormValue("name"))
}
