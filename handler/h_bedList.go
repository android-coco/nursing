//  Created by JP

package handler

import (
	"fit"
	"strconv"
	"nursing/model"
)

/*床位列表页面*/
type BedListController struct {
	fit.Controller
}

/*PDA 床位列表API*/
func (c BedListController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	defer c.ResponseToJson(w)
	r.ParseForm()
	department_id := r.FormValue("department_id")

	if department_id == "" {
		c.RenderingJsonAutomatically(1, "参数不完整")
		return
	}

	depid_i, err_dep := strconv.Atoi(department_id)
	if err_dep != nil || depid_i < 0 {
		c.RenderingJsonAutomatically(2, "参数错误： department_id")
		return
	}

	response, err := model.QueryDepartmentBedList(depid_i)

	if err != nil {
		c.RenderingJsonAutomatically(3, "Database "+err.Error())
	} else {
		c.RenderingJson(0, "成功", response)
	}
}

func (c *BedListController) RenderingJsonAutomatically(result int, errMsg string) {
	c.RenderingJson(result, errMsg, make([]interface{}, 0))
}
func (c *BedListController) RenderingJson(result int, errMsg string, datas interface{}) {
	c.JsonData.Datas = datas
	c.JsonData.ErrorMsg = errMsg
	c.JsonData.Result = result
}

