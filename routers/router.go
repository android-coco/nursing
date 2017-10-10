package nursing

import (
	"nursing/handler"
	"fit"
	//"nursing/model"
)

func init() {
	//路由配置
	//fit.Router().AddRouter("/", &handler.MainController{})
	fit.Router().AddRouter("/", &handler.MainController{},"get,post:GetFunc")
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
	//医嘱查询
	fit.Router().AddRouter("/medical/advice",new(handler.MedicalAdvice))
	fit.Router().AddRouter("/deviceiput", new(handler.DeviceiputController))
	fit.Router().AddRouter("/deviceout", new(handler.DeviceoutController))

	//数据库配置[]interface{}{new(model.Warn),new(model.Warn),new(model.Warn),new(model.Warn)}
	//fit.App().InitModels([]interface{}{new(model.Warn), new(model.Access),new(model.IntakeOutput),
	//new(model.Temperature),new(model.Pulse),new(model.Breathe),new(model.Pressure),new(model.Heartrate),
	//new(model.Spo2h),new(model.Glucose),new(model.Weight),new(model.Height)})
	//fit.App().InitModels([]interface{}{new(model.Warn), new(model.Access)})
}
