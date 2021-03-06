package utils

import "strconv"

//string是否为空
func WhetherStrNiu (str string)bool{
	if len(str) == 0{
		return true
	}
	return false
}

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func DelLastIndex(str string) string  {
	rs := []rune(str)
	length := len(rs)
	return string(rs[0:length - 1])
}

func Int64Value(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func FormatInt64(num int64) string  {
	return strconv.FormatInt(num, 10)
}

