package mzjtime

import (
	"fmt"
	"time"
)

//TimeFormat 时间格式
type TimeFormat int

const (
	timeTeplate1 TimeFormat = iota
	timeTeplate2
	timeTeplate3
	timeTeplate4
	timeTeplate5
)

func (t TimeFormat) String() string {
	switch t {
	case 0:
		return "2006-01-02 15:04:05" //yyyy-MM-dd HH:mm:ss
	case 1:
		return "2006/01/02 15:04:05" //yyyy/MM/dd HH:mm:ss
	case 2:
		return "2006-01-02" //yyyy-MM-dd
	case 3:
		return "2006/01/02" //yyyy/MM/dd
	case 4:
		return "15:04:05" //HH:mm:ss
	default:
		return "时间格式未定义"
	}
}

//Format 时间转字符串
func Format(times time.Time, f TimeFormat) string {
	return times.Format(f.String())
}

//UnixFormat 时间戳转字符串
func UnixFormat(sec int64, f TimeFormat) string {
	return time.Unix(sec, 0).Format(f.String())
}

//ParseInlocation 字符串转时间
func ParseInlocation(str string, f TimeFormat) (time.Time, error) {
	return time.ParseInLocation(f.String(), str, time.Local)
}

//Now 获取当前时间
func Now() time.Time {
	return time.Now()
}

//Unix 时间转成时间戳（秒）
func Unix(times time.Time) int64 {
	return times.Unix()
}

//UnixNano 时间转成时间戳（毫秒）
func UnixNano(times time.Time) int64 {
	return times.UnixNano() / 1e6
}

//UnixToTime 时间戳转化成时间
func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

//StringToUnix 字符串转时间戳（秒）
func StringToUnix(str string, f TimeFormat) int64 {
	ts, _ := ParseInlocation(str, f)
	return Unix(ts)
}

//StringToUnixNano 字符串转时间戳(毫秒)
func StringToUnixNano(str string, f TimeFormat) int64 {
	ts, _ := ParseInlocation(str, f)
	return UnixNano(ts)
}

//Add 添加日期
func Add(ts time.Time, d time.Duration) time.Time {
	return ts.Add(d)
}

//AddDate 添加日期
func AddDate(ts time.Time, years, months, days int) time.Time {
	return ts.AddDate(years, months, days)
}

func main() {
	x := Now()
	fmt.Println(x)

}
