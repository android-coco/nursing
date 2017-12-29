package utils

import (

	"time"
	"fmt"
	"crypto/sha1"
	"io"
)

//数据库模型转JSON 数据
/*func Transfer(slice []model.I_BaseMoel) []interface{} {
	var ifSlice = make([]model.I_BaseMoel, len(slice))
	for idx, v := range slice {
		ifSlice[idx] = v
	}
	var intfaceSlice = make([]interface{}, len(ifSlice))
	for idx, v := range ifSlice {
		intfaceSlice[idx] = v
	}
	return intfaceSlice
}*/

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

//2个时间比较大小  time1 < time2 reutn true
func CompareTime(time1, time2 string) bool {
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		return true
	}
	return false
}

//2个时间比较大小  time1 < time2 reutn true
func CompareTimeNow(time1 string) bool {
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", time1, time.Local)
	fmt.Println(t1)
	t2 := time.Now()
	fmt.Println(t2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		return true
	}
	return false
}

//数字时间转汉字时间
var timeChina =  []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
//var dateChina =  []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
func sinicizingTime(intval int) (datestr string) {
	if intval == 0 {
		datestr = "零"
	} else if intval < 10 {
		datestr = timeChina[intval]
	} else if intval == 10 {
		datestr = "十"
	} else {
		datestr = timeChina[intval / 10] + "十" + timeChina[intval % 10]
	}

	//if intval >= 0 && intval <= 60 {
	//
	//} else {
	//	datestr = ""
	//}
	return
}

func STime(t time.Time) (datestr string) {
	return sinicizingTime(t.Hour()) + "时:" + sinicizingTime(t.Minute()) + "分"
}