package model

import (
	"fit"
)

type MonitorInfo struct {
	ID             int64           `xorm:"pk autoincr comment(数据ID自动增长)" json:"id"`
	V1             string          `xorm:"comment(病人总数)" json:"v1"`
	V2             string          `xorm:"comment(新入)" json:"v2"`
	V3             string          `xorm:"comment(出院)" json:"v3"`
	V4             string          `xorm:"comment(一级护理)" json:"v4"`
	V5             string          `xorm:"comment(病重)" json:"v5"`
	V6             string          `xorm:"comment(病危)" json:"v6"`
	V7             string          `xorm:"comment(心电监护)" json:"v7"`
	V8             string          `xorm:"comment(吸氧)" json:"v8"`
	V9             string          `xorm:"comment(神志瞳孔监测)" json:"v9"`
	V10            string          `xorm:"comment(雾化吸入)" json:"v10"`
	V11            string          `xorm:"comment(口腔护理)" json:"v11"`
	V12            string          `xorm:"comment(尿道口护理)" json:"v12"`
	V13            string          `xorm:"comment(留置导尿)" json:"v13"`
	V14            string          `xorm:"comment(计24小时尿量)" json:"v14"`
	V15            string          `xorm:"comment(计24小时出入量)" json:"v15"`
	V16            string          `xorm:"comment(血压监测)" json:"v16"`
	V17            string          `xorm:"comment(血糖监测)" json:"v17"`
	V18            string          `xorm:"comment(自定义1)" json:"v18"`
	V19            string          `xorm:"comment(自定义2)" json:"v19"`
	V20            string          `xorm:"comment(其他)" json:"v20"`
	NurseName      string          `xorm:"comment(值班护士)" json:"nursename"`
	DoctorName     string          `xorm:"comment(值班医生)" json:"doctorname"`
	Speed          string          `xorm:"comment(轮播速度等级1,2,3,4)" json:"speed"`
	Display        string          `xorm:"comment(是否显示床位)" json:"display"`
	UpdateTime     string          `xorm:"comment(最后更新时间)" json:"updatetime"`
	UserName       string          `xorm:"comment(最后编辑人)" json:"username"`
	ClassName      string          `xorm:"comment(科室名称)" json:"classname"`
	ClassId        string          `xorm:"comment(科室ID)" json:"classid"`
	MonitorNotifys []MonitorNotify `xorm:"extends" json:"notifys"` //通知内容 多个内容
}

type MonitorNotify struct {
	ID            int64  `xorm:"pk autoincr comment(数据ID自动增长)" json:"id"`
	MonitorInfoId int64  `xorm:"comment(monitorinfo表数据ID)" json:"mid"`
	NotifyInfo    string `xorm:"comment(通知内容)" json:"notifyinfo"`
}

// 查询数据
func TVQueryMonitor(classid string) (monitors []MonitorInfo, err error) {
	monitors = make([]MonitorInfo, 0)
	err = fit.MySqlEngine().SQL("SELECT * FROM monitorinfo WHERE classid = ?  ORDER BY updatetime DESC", classid).Find(&monitors)
	if err != nil {
		return
	}
	for i := 0; i < len(monitors); i++ {
		monitorNotify := make([]MonitorNotify, 0)
		err = fit.MySqlEngine().SQL("SELECT * FROM monitornotify WHERE monitorinfoid = ?", monitors[i].ID).Find(&monitorNotify)
		if err != nil {
			break
		}
		monitors[i].MonitorNotifys = monitorNotify
	}
	return
}

// 删除数据  通知
func TVDelMonitorNotify(notify MonitorNotify) (err error) {
	_, err = fit.MySqlEngine().SQL("DELETE FROM monitornotify WHERE id = ? ", notify.ID).Delete(&notify)
	if err != nil {
		return
	}
	return
}

// 更新数据
func TVUpdataMonitorInfo(info MonitorInfo) (err error) {
	if info.ID <= 0 {
		_, err = fit.MySqlEngine().Table("monitorinfo").Insert(&info)
		if len(info.MonitorNotifys) > 0 {
			for _, notify := range info.MonitorNotifys {
				if notify.ID <= 0 {
					notify.MonitorInfoId = info.ID
					_, err = fit.MySqlEngine().Table("monitornotify").Insert(&notify)
					if err != nil {
						break
					}
				}else{
					_, err = fit.MySqlEngine().Table("monitornotify").ID(notify.ID).AllCols().Update(&notify)
					if err != nil {
						break
					}
				}
			}
		}
	} else {
		_, err = fit.MySqlEngine().Table("monitorinfo").ID(info.ID).AllCols().Update(&info)
		if len(info.MonitorNotifys) > 0 {
			for _, notify := range info.MonitorNotifys {
				if notify.ID <= 0 {
					notify.MonitorInfoId = info.ID
					_, err = fit.MySqlEngine().Table("monitornotify").Insert(&notify)
					if err != nil {
						break
					}
				}else{
					_, err = fit.MySqlEngine().Table("monitornotify").ID(notify.ID).AllCols().Update(&notify)
					if err != nil {
						break
					}
				}
			}
		}
	}
	if err != nil {
		return
	}
	return
}
