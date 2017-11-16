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

	// 护理记录单 pda端
	nrlist()

	// pc 端
	pclist()


	// pc 护理记录
	pcnrlist()
}

func pdalist() {
	fit.Router().AddRouter("/", &handler.MainController{}, "get,post:GetFunc")
	fit.Router().AddRouter("/login", new(handler.LoginController))

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

	fit.Router().AddRouter("/inputnursechat", new(handler.NurseChatInputController))  //护理体温单提交
	fit.Router().AddRouter("/outnursechat", new(handler.NurseChatOutputController))  //获取护理体温单数据
}

func nrlist() {
	fit.Router().AddRouter("/record/query", new(handler.QueryRecordController))
	// 护理记录单
	fit.Router().AddRouter("/record/nr1", new(handler.NRL1Controller))
	fit.Router().AddRouter("/record/nr1/edit", new(handler.NRL1Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr1/add", new(handler.NRL1Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr1/update", new(handler.NRL1Controller), "get,post:UpdateRecord")
	// 首次护理记录单
	fit.Router().AddRouter("/record/nr2", new(handler.NRL2Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr2/edit", new(handler.NRL2Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr2/update", new(handler.NRL2Controller), "get,post:AddRecord")
	//fit.Router().AddRouter("/record/nr2/update", new(handler.NRL2Controller), "get,post:UpdateRecord")
	//fit.Router().AddRouter("/record/nr2/add", new(handler.FirstNursingRecordController))
	//fit.Router().AddRouter("/record/nr2/edit", new(handler.FirstNursingRecordController))
	//fit.Router().AddRouter("/record/nr2", new(handler.QueryFirstNursingRecordController))
	//fit.Router().AddRouter("/record/nr2/exist", new(handler.QueryFirstNursingRecordController), "get:Exist")

	//基本生活活动能力BADL
	fit.Router().AddRouter("/record/nr3", new(handler.NRL3Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr3/edit", new(handler.NRL3Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr3/add", new(handler.NRL3Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr3/update", new(handler.NRL3Controller), "get,post:UpdateRecord")

	//深静脉血栓形成风险评估表
	fit.Router().AddRouter("/record/nr4", new(handler.NRL4Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr4/edit", new(handler.NRL4Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr4/add", new(handler.NRL4Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr4/update", new(handler.NRL4Controller), "get,post:UpdateRecord")

	//深静脉血栓观察表
	fit.Router().AddRouter("/record/nr5", new(handler.NRL5Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr5/edit", new(handler.NRL5Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr5/add", new(handler.NRL5Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr5/update", new(handler.NRL5Controller), "get,post:UpdateRecord")
	//查询是否存在某个班次的深静脉血栓护理观察单
	fit.Router().AddRouter("/record/nr5/exist", new(handler.NRL5Controller), "get,post:Exist")

	//压疮风险因素评估表
	fit.Router().AddRouter("/record/nr6", new(handler.NRL6Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr6/update", new(handler.NRL6Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr6/edit", new(handler.NRL6Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr6/add", new(handler.NRL6Controller), "get,post:AddRecord")

	//患者跌倒风险评估护理单
	fit.Router().AddRouter("/record/nr7", new(handler.NRL7Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr7/edit", new(handler.NRL7Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr7/add", new(handler.NRL7Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr7/update", new(handler.NRL7Controller), "get,post:UpdateRecord")
	//fit.Router().AddRouter("/record/nr7/addtitle", new(handler.NRL7Controller), "get,post:InsertTitle")
	fit.Router().AddRouter("/record/nr7/updatetitle", new(handler.NRL7Controller), "get,post:UpdateTitle")

	//疼痛强度评分量表
	fit.Router().AddRouter("/record/nr8", new(handler.NRL8Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr8/edit", new(handler.NRL8Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr8/add", new(handler.NRL8Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr8/update", new(handler.NRL8Controller), "get,post:UpdateRecord")

}

func pclist() {
	// PC主页
	fit.Router().AddRouter("/pc/home", new(handler.PCHomeController))
	fit.Router().AddRouter("/pc/home/api", new(handler.PCBedController))
	fit.Router().AddRouter("/pc/login", new(handler.PCLoginController))
	fit.Router().AddRouter("/pc/logout", new(handler.PCLoginController), "get:Logout")

	// 账号管理
	fit.Router().AddRouter("/pc/account/changepwd", new(handler.ChangePasswordController))
	fit.Router().AddRouter("/pc/account/manage", new(handler.AccountManageController), "get:Manage")
	fit.Router().AddRouter("/pc/account/manage/created", new(handler.AccountManageController), "get:List")
	fit.Router().AddRouter("/pc/account/create", new(handler.AccountManageController), "post:Create")
	fit.Router().AddRouter("/pc/account/update", new(handler.AccountManageController), "post:Update")
	fit.Router().AddRouter("/pc/account/uncreated", new(handler.AccountManageController), "get:Uncreated")
	fit.Router().AddRouter("/pc/account/created", new(handler.AccountManageController), "get:Created")

	// 设备管理
	fit.Router().AddRouter("/pc/device/manage", new(handler.DeviceManageController))
	fit.Router().AddRouter("/pc/host/config", new(handler.HostConfigController))

	// 交接班
	fit.Router().AddRouter("/pc/succession", new(handler.PCSuccessController))

	// 出入管理
	fit.Router().AddRouter("/pc/access/manage", new(handler.PCAccessController))

	// 医嘱信息
	fit.Router().AddRouter("/pc/medicaladvice/message", new(handler.MedicalAdviceMessage))
	fit.Router().AddRouter("/pc/medicaladvice/Detail", new(handler.MedicalAdviceDetail))

	// 医嘱拆分
	fit.Router().AddRouter("/pc/medicaladvice/split", new(handler.MedicalAdviceSplit))

	// 体征录入
	fit.Router().AddRouter("/pc/batvhinput", new(handler.PCBatvhinputController))
	fit.Router().AddRouter("/pc/patvhhistory", new(handler.PCBatvhHistoryController), "get,post:TZHistory")

	//各种打印页预览
	//腕带
	fit.Router().AddRouter("/pc/wdprint", new(handler.PCWristStrapController))
	//屏贴
	fit.Router().AddRouter("/pc/ptprint", new(handler.PCBottleStrapController))
	// 体温单
	fit.Router().AddRouter("/pc/templist", new(handler.TempChartController), "get,post:LoadTable")

	//提醒管理
	fit.Router().AddRouter("/pc/warn", new(handler.PCWarnController))
	fit.Router().AddRouter("/pc/warn/del/:id", new(handler.PCWarnController), "get,post:DelWarn")
	fit.Router().AddRouter("/pc/warn/modify", new(handler.PCWarnController), "get,post:ModifyWarn")

	//历史记录
	fit.Router().AddRouter("/pc/history", new(handler.PCHistoryController))
	fit.Router().AddRouter("/pc/history/search",new(handler.PCHistoryController),"get:SearchPatients")
	fit.Router().AddRouter("/pc/history/signs", new(handler.PCHistoryController),"get,post:Signs")
	fit.Router().AddRouter("/pc/history/temperature", new(handler.PCHistoryController),"get,post:Temperature")
	fit.Router().AddRouter("/pc/history/advice", new(handler.PCHistoryController),"get,post:Advice")
}

func pcnrlist()  {
	// pc 端
	fit.Router().AddRouter("/pc/record/nrl1", new(handler.PCNRL1Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl2", new(handler.PCNRL2Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl3", new(handler.PCNRL3Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl4", new(handler.PCNRL4Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl5", new(handler.PCNRL5Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl6", new(handler.PCNRL6Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl7", new(handler.PCNRL7Controller), "get,post:NRLRecord")
	fit.Router().AddRouter("/pc/record/nrl8", new(handler.PCNRL8Controller), "get,post:NRLRecord")
}