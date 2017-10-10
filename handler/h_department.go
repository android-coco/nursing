package handler

import (
	"fit"
	"nursing/model"
)

/*科室列表*/
type DepartmentController struct {
	fit.Controller
}

func (c DepartmentController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	departments, err := model.QueryDepartmentList()
	if err != nil {
		c.RenderingJsonAutomatically(3, "Database "+err.Error())
	} else {
		c.RenderingJson(0, "成功", departments)
	}
}

func (c *DepartmentController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}

func (c *DepartmentController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}
