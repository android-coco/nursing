package nursing

import (
	"nursing/handler"
	"fit"

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
	//PC 打印
	pcprint()
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

	//医嘱
	fit.Router().AddRouter("/madvice/search", new(handler.MedicalAdviceController),"post:Search")
	fit.Router().AddRouter("/madvice/status/search", new(handler.MedicalAdviceController),"post:StatusSearch")
	fit.Router().AddRouter("/madvice/exec/search", new(handler.MedicalAdviceController),"post:ExecSearch")
	fit.Router().AddRouter("/madvice/exec/detail", new(handler.MedicalAdviceController),"post:ExecDetail")
	fit.Router().AddRouter("/madvice/exec", new(handler.MedicalAdviceController),"post:Execute")

	fit.Router().AddRouter("/deviceiput", new(handler.DeviceiputController))
	fit.Router().AddRouter("/deviceout", new(handler.DeviceoutController))

	fit.Router().AddRouter("/inputnursechat", new(handler.NurseChatController),"post:NurseChatInput")  //护理体温单提交
	fit.Router().AddRouter("/outnursechat", new(handler.NurseChatController),"post:NurseChatOutput")  //获取护理体温单数据
}

func nrlist() {
	fit.Router().AddRouter("/record/query", new(handler.QueryRecordController))
	// 护理记录单
	fit.Router().AddRouter("/record/nr1/update", new(handler.NRL1Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr1/updatetitle", new(handler.NRL1Controller), "get,post:UpdateTitle")
	//fit.Router().AddRouter("/record/nr1", new(handler.NRL1Controller))
	//fit.Router().AddRouter("/record/nr1/edit", new(handler.NRL1Controller), "get,post:Edit")
	//fit.Router().AddRouter("/record/nr1/add", new(handler.NRL1Controller), "get,post:AddRecord")
	//fit.Router().AddRouter("/record/nr1/update", new(handler.NRL1Controller), "get,post:UpdateRecord")
	// 首次护理记录单
	fit.Router().AddRouter("/record/nr2", new(handler.NRL2Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr2/edit", new(handler.NRL2Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr2/update", new(handler.NRL2Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr2/exist", new(handler.NRL2Controller), "get:Exist")

	//基本生活活动能力BADL
	fit.Router().AddRouter("/record/nr3", new(handler.NRL3Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr3/edit", new(handler.NRL3Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr3/add", new(handler.NRL3Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr3/update", new(handler.NRL3Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr3/delete", new(handler.NRL3Controller), "get,post:DeleteRecord")

	//深静脉血栓形成风险评估表
	fit.Router().AddRouter("/record/nr4", new(handler.NRL4Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr4/edit", new(handler.NRL4Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr4/add", new(handler.NRL4Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr4/update", new(handler.NRL4Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr4/delete", new(handler.NRL4Controller), "get,post:DeleteRecord")

	//深静脉血栓观察表
	fit.Router().AddRouter("/record/nr5", new(handler.NRL5Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr5/edit", new(handler.NRL5Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr5/add", new(handler.NRL5Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr5/update", new(handler.NRL5Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr5/delete", new(handler.NRL5Controller), "get,post:DeleteRecord")
	//查询是否存在某个班次的深静脉血栓护理观察单
	fit.Router().AddRouter("/record/nr5/exist", new(handler.NRL5Controller), "get,post:Exist")

	//压疮风险因素评估表
	fit.Router().AddRouter("/record/nr6", new(handler.NRL6Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr6/update", new(handler.NRL6Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr6/edit", new(handler.NRL6Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr6/add", new(handler.NRL6Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr6/delete", new(handler.NRL6Controller), "get,post:DeleteRecord")

	//患者跌倒风险评估护理单
	fit.Router().AddRouter("/record/nr7", new(handler.NRL7Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr7/edit", new(handler.NRL7Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr7/add", new(handler.NRL7Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr7/update", new(handler.NRL7Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr7/delete", new(handler.NRL7Controller), "get,post:DeleteRecord")
	fit.Router().AddRouter("/record/nr7/updatetitle", new(handler.NRL7Controller), "get,post:UpdateTitle")

	//疼痛强度评分量表
	fit.Router().AddRouter("/record/nr8", new(handler.NRL8Controller), "get,post:Check")
	fit.Router().AddRouter("/record/nr8/edit", new(handler.NRL8Controller), "get,post:Edit")
	fit.Router().AddRouter("/record/nr8/add", new(handler.NRL8Controller), "get,post:AddRecord")
	fit.Router().AddRouter("/record/nr8/update", new(handler.NRL8Controller), "get,post:UpdateRecord")
	fit.Router().AddRouter("/record/nr8/delete", new(handler.NRL8Controller), "get,post:DeleteRecord")

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

	// 医嘱
	fit.Router().AddRouter("/pc/madvice/search",new(handler.PCMedicalAdviceController),"get:PCSearch")
	fit.Router().AddRouter("/pc/madvice/exec/state",new(handler.PCMedicalAdviceController),"get:PCExecState")
	fit.Router().AddRouter("/pc/madvice/split",new(handler.PCMedicalAdviceController),"get:PCSplit")
	fit.Router().AddRouter("/pc/madvice/exec/search", new(handler.PCMedicalAdviceController),"post:PCExecSearch")
	fit.Router().AddRouter("/pc/madvice/split/search",new(handler.PCMedicalAdviceController),"post:SpiltSearch")
	fit.Router().AddRouter("/pc/madvice/exec/detail",new(handler.PCMedicalAdviceController),"get:PCExecDetail")
	fit.Router().AddRouter("/pc/madvice/search/api",new(handler.PCMedicalAdviceController),"post:Search")


	// 体征批量录入
	fit.Router().AddRouter("/pc/batvhinput", new(handler.PCBatvhinputController))
	//体征历史
	fit.Router().AddRouter("/pc/patvhhistory", new(handler.PCBatvhHistoryController), "get,post:TZHistory")
	fit.Router().AddRouter("/pc/patvhdelete", new(handler.PCBatvhHistoryController),"get,post:TZDelete")
	fit.Router().AddRouter("/pc/patvhupdate", new(handler.PCBatvhHistoryController),"get,post:TZUpdate")

	// 体温单
	fit.Router().AddRouter("/pc/templist", new(handler.TempChartController), "get,post:LoadTable")
	fit.Router().AddRouter("/pc/templist/print", new(handler.TempChartController), "get,post:PrintTempChart")

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

	// 出入量统计接口
	fit.Router().AddRouter("/pc/record/nrl/io", new(handler.PCNRL1Controller), "get,post:NRLIOStatistcs")
	fit.Router().AddRouter("/pc/record/nrl/insertio", new(handler.PCNRL1Controller), "get,post:NRLIOTypeIn")
}

func pcprint(){
	//打印腕带
	fit.Router().AddRouter("/pc/wdprint", new(handler.PCWristStrapController))
	//打印屏贴
	fit.Router().AddRouter("/pc/ptprint", new(handler.PCBottleStrapController))
	//打印深圳万丰医院护理记录单
	fit.Router().AddRouter("/pc/nrl1print", new(handler.PCNrl1Controller))
	//打印首次护理记录单
	fit.Router().AddRouter("/pc/nrl2print", new(handler.PCNrl2Controller))
	//打印深圳万丰医院基本生活活动能力(BADL)
	fit.Router().AddRouter("/pc/nrl3print", new(handler.PCNrl3Controller))
	//深圳万丰医院深静脉血栓(DVT)形成风险评估表
	fit.Router().AddRouter("/pc/nrl4print", new(handler.PCNrl4Controller))
	//打印深静脉血栓护理观察表
	fit.Router().AddRouter("/pc/nrl5print", new(handler.PCNrl5Controller))
	//打印深圳万丰医院压疮风险因素评估表（Braden评分）
	fit.Router().AddRouter("/pc/nrl6print", new(handler.PCNrl6Controller))
	//打印深圳万丰医院患者跌到风险评估护理单
	fit.Router().AddRouter("/pc/nrl7print", new(handler.PCNrl7Controller))
	//打印深圳万丰医院深圳万丰医院疼痛强度评分量表
	fit.Router().AddRouter("/pc/nrl8print", new(handler.PCNrl8Controller))

	//交接班打印
	fit.Router().AddRouter("/pc/successprint", new(handler.PCSuccessionController))

}