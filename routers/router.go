package nursing

import (
	"nursing/handler"
	"fit"
)

func init() {
	//路由配置
	fit.Router().AddRouter("/", &handler.MainController{})
	fit.Router().AddRouter("/login", new(handler.LoginController))
	fit.Router().AddRouter("/signsiput", new(handler.SignsiputController))
	fit.Router().AddRouter("/signsout", new(handler.SignsoutController))
	fit.Router().AddRouter("/warnadd", new(handler.WarnController))
	fit.Router().AddRouter("/warnlist", new(handler.WarnListController))
	fit.Router().AddRouter("/accessadd", new(handler.AccessController))
	fit.Router().AddRouter("/accesslist", new(handler.AccessListController))
	fit.Router().AddRouter("/iov/collect", new(handler.IntakeOutputCollectController))
	fit.Router().AddRouter("/iov/query", new(handler.IntakeOutputQueryController))
	//医嘱查询
	fit.Router().AddRouter("/medical/advice",new(handler.MedicalAdvice))


	//数据库配置[]interface{}{new(model.Warn),new(model.Warn),new(model.Warn),new(model.Warn)}
	//fit.App().InitModels([]interface{}{new(model.Warn), new(model.Access),new(model.IntakeOutput),
	//new(model.Temperature),new(model.Pulse),new(model.Breathe),new(model.Pressure),new(model.Heartrate),
	//new(model.Spo2h),new(model.Glucose),new(model.Weight),new(model.Height)})
}
