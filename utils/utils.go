package utils

import "nursing/model"

//数据库模型转JSON 数据
func Transfer(slice []model.I_BaseMoel) []interface{} {
	var ifSlice []model.I_BaseMoel = make([]model.I_BaseMoel, len(slice))
	for idx, v := range slice {
		ifSlice[idx] = v
	}
	var intfaceSlice []interface{} = make([]interface{}, len(ifSlice))
	for idx,v:=range ifSlice{
		intfaceSlice[idx] = v
	}
	return intfaceSlice
}