//  Created by JP

package handler

import (
	"fit"
	"nursing/model"
)

/*科室列表*/
type DepartmentController struct {
	fit.Controller
}

/*API 或者科室列表*/
func (c DepartmentController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	departments, err := model.QueryDepartmentList(true)
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
