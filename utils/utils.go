package utils

import (
	"nursing/model"
	"time"
	"fmt"
	"crypto/sha1"
	"io"
)

//数据库模型转JSON 数据
func Transfer(slice []model.I_BaseMoel) []interface{} {
	var ifSlice = make([]model.I_BaseMoel, len(slice))
	for idx, v := range slice {
		ifSlice[idx] = v
	}
	var intfaceSlice = make([]interface{}, len(ifSlice))
	for idx, v := range ifSlice {
		intfaceSlice[idx] = v
	}
	return intfaceSlice
}

//  函数执行时间
func Trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() {
		fmt.Printf("exit %s (%s)\n", msg, time.Since(start))
	}
}

/*sha1加密*/
func Sha1Encryption(material string) string {
	if material == "" {
		return material
	}
	handler := sha1.New()
	io.WriteString(handler, material)
	pwBt := handler.Sum(nil)
	return fmt.Sprintf("%x", pwBt)
}
