package nursing

import (
	"nursing/handler"
	"fit"
	//"nursing/model"
)

func init() {
	//路由配置
	//fit.Router().AddRouter("/", &handler.MainController{})
	// PDA接口
	pdalist()

	// 护理记录

	nrlist()

	// pc 端
	pclist()

	//数据库配置[]interface{}{new(model.Warn),new(model.Warn),new(model.Warn),new(model.Warn)}
	//fit.App().InitModels([]interface{}{new(model.Warn), new(model.Access),new(model.IntakeOutput),
	//new(model.Temperature),new(model.Pulse),new(model.Breathe),new(model.Pressure),new(model.Heartrate),
	//new(model.Spo2h),new(model.Glucose),new(model.Weight),new(model.Height)})
	//fit.App().InitModels([]interface{}{new(model.Warn), new(model.Access)})
}

func pdalist() {
	fit.Router().AddRouter("/", &handler.MainController{}, "get,post:GetFunc")
	fit.Router().AddRouter("/login", new(handler.LoginController))
	fit.Router().AddRouter("/signsiput", new(handler.SignsiputController))
	fit.Router().AddRouter("/signsout", new(handler.SignsoutController))

	fit.Router().AddRouter("/warn/add", new(handler.WarnController))
	fit.Router().AddRouter("/warn/del", new(handler.WarnController), "get,post:DelWarn")
	fit.Router().AddRouter("/warn/list", new(handler.WarnListController))
	fit.Router().AddRouter("/access/add", new(handler.AccessController))
	fit.Router().AddRouter("/access/list", new(handler.AccessListController))
	fit.Router().AddRouter("/access/search", new(handler.AccessSearchController))

	fit.Router().AddRouter("/iov/collect", new(handler.IntakeOutputCollectController))
	fit.Router().AddRouter("/iov/query", new(handler.IntakeOutputQueryController))
	fit.Router().AddRouter("/bedlist", new(handler.BedListController))
	fit.Router().AddRouter("/departments", new(handler.DepartmentController))
	fit.Router().AddRouter("/patient/info", new(handler.PatientInfoController))
	//医嘱查询
	fit.Router().AddRouter("/advice/query", new(handler.MedicalAdviceQuery))
	fit.Router().AddRouter("/advice/execute", new(handler.MedicalAdviceExecute))
	fit.Router().AddRouter("/advice/statequery", new(handler.MedicalAdviceStateQuery))
	fit.Router().AddRouter("/advice/newpause", new(handler.MedicalAdviceNewPause))

	fit.Router().AddRouter("/deviceiput", new(handler.DeviceiputController))
	fit.Router().AddRouter("/deviceout", new(handler.DeviceoutController))
}

func nrlist() {
	fit.Router().AddRouter("/record/query", new(handler.QueryRecordController))
	// 护理记录单
	fit.Router().AddRouter("/record/nr1", new(handler.RecordController))
	fit.Router().AddRouter("/record/nr1/edit", new(handler.RecordController), "get,post:Edit")
	fit.Router().AddRouter("/record/nr1/add", new(handler.NRLController), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr1/update", new(handler.NRLController), "get,post:UpdateRecord")
	// 首次护理记录单
	fit.Router().AddRouter("/record/nr2/add", new(handler.FirstNursingRecordController))
	fit.Router().AddRouter("/record/nr2/edit", new(handler.FirstNursingRecordController))
	fit.Router().AddRouter("/record/nr2", new(handler.QueryFirstNursingRecordController))
	fit.Router().AddRouter("/record/nr2/exist", new(handler.QueryFirstNursingRecordController), "get:Exist")

	// pc 端
	fit.Router().AddRouter("/pc/record/nr1", new(handler.PNRLController), "get,post:NRL1Record")
	fit.Router().AddRouter("/pc/tempchart", new(handler.TempChartController), "get,post:TempChart")
}

func pclist() {
	// PC主页
	fit.Router().AddRouter("/pc/home", new(handler.PCHomeController))
	fit.Router().AddRouter("/pc/home/beds", new(handler.PCBedController))
	fit.Router().AddRouter("/pc/login", new(handler.PCLoginController))
	fit.Router().AddRouter("/pc/logout", new(handler.PCLoginController), "get:Logout")

	// 交接班
	fit.Router().AddRouter("/pc/succession", new(handler.PCSuccessController))

	//各种打印页预览
	//腕带
	fit.Router().AddRouter("/pc/wdprint", new(handler.PCWristStrapController))
	//屏贴
	fit.Router().AddRouter("/pc/ptprint", new(handler.PCBottleStrapController))

	fit.Router().AddRouter("/pc/templist", new(handler.TempChartController), "get,post:LoadTable")
}
