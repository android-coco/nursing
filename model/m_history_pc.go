//  Created by JP

package model

import (
	"fit"
	"time"
)

type HistoryPatient struct {
	VAA01 int64     `json:"patientId"`
	VAE02 int64     `json:"rgtNum"`
	VAA05 string    `json:"name"`
	VAA04 string    `json:"hospNum"`
	VAE26 time.Time `json:"-"`
	Date  string    `json:"dcgDate"`
}

/*根据出院时间区间查找历史病人*/
func SearchHistoryParientsWithTimeInterval(startTime, endTime string, did int) ([]HistoryPatient, error) {
	response := make([]HistoryPatient, 0)
	err := fit.SQLServerEngine().SQL("select VAE1.VAA01, VAE1.VAE02, VAE1.VAE26, VAA1.VAA04, VAA1.VAA05 from VAE1, VAA1 where VAE1.BCK01D = ? and (VAE1.VAE26 between ? and ?) and VAA1.VAA01 = VAE1.VAA01", did, startTime, endTime).Find(&response)
	for i, v := range response {
		response[i].Date = v.VAE26.Format("2006-01-02 15:04")
	}
	return response, err
}

/*根据登记号查找*/
func SearchHistoryParientsWithRegisterNum(rgtNum string, did int) ([]HistoryPatient, error) {
	response := make([]HistoryPatient, 0)
	err := fit.SQLServerEngine().SQL("select VAE1.VAA01, VAE1.VAE02, VAE1.VAE26, VAA1.VAA04, VAA1.VAA05 from VAE1, VAA1 where VAE1.BCK01D = ? and VAE1.VAE02 = ? and VAA1.VAA01 = VAE1.VAA01 and VAE1.VAE26 is not null", did, rgtNum).Find(&response)
	for i, v := range response {
		response[i].Date = v.VAE26.Format("2006-01-02 15:04")
	}
	return response, err
}

/*根据住院号查找*/
func SearchHistoryParientsWithHospitalizationNum(hospNum string, did int) ([]HistoryPatient, error) {
	response := make([]HistoryPatient, 0)
	err := fit.SQLServerEngine().SQL("select VAA1.VAA01, VAA1.VAA04, VAA1.VAA05, VAE1.VAE02, VAE1.VAE26 from VAA1, VAE1 where VAA1.BCK01B = ? and VAA1.VAA04 = ? and VAA1.VAA01 = VAE1.VAA01 and VAE1.VAE26 is not null", did, hospNum).Find(&response)
	for i, v := range response {
		response[i].Date = v.VAE26.Format("2006-01-02 15:04")
	}
	return response, err
}

/*根据姓名查找*/
func SearchHistoryParientsWithName(name string, did int) ([]HistoryPatient, error) {
	response := make([]HistoryPatient, 0)
	err := fit.SQLServerEngine().SQL("select VAA1.VAA01, VAA1.VAA04, VAA1.VAA05, VAE1.VAE02, VAE1.VAE26 from VAA1, VAE1 where VAA1.BCK01B = ? and (VAA1.VAA05 = ? or VAA1.ABBRP = ?) and VAA1.VAA01 = VAE1.VAA01 and VAE1.VAE26 is not null", did, name, name).Find(&response)
	for i, v := range response {
		response[i].Date = v.VAE26.Format("2006-01-02 15:04")
	}
	return response, err
}

