package mtime

import (
	"fmt"
	"strings"
	"time"
)

var days = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// 当前时间转换成字符串
func Date(format string, nowTime time.Time) string {
	if format == "" {
		return ""
	}
	//判断对象nowTime是否是Time类型
	_, ok := interface{}(nowTime).(time.Time)
	if !ok {
		panic("two param type is error.")
		return ""
	}

	var str string
	for i := 0; i < len(format); i++ {
		str += dateFormat(format[i], nowTime)
		if format[i] == '/' {
			str += "/"
		}
		if format[i] == '-' {
			str += "-"
		}
		if format[i] == ' ' {
			str += " "
		}
		if format[i] == ':' {
			str += ":"
		}
	}

	return str
}

func dateFormat(b byte, nowTime time.Time) string {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	nowTime.In(cstSh)
	var date string
	//nowTime.Format("2006-01-02 15:04:05")
	switch {
	case b == 'Y':
		// 四位数的年份
		date = nowTime.Format("2006")
	case b == 'm':
		// 数字表示的月份，有前导零
		date = nowTime.Format("01")
	case b == 'M':
		// 英文 月份,完整的文本格式 January 到 December
		date = nowTime.Month().String()
	case b == 'd':
		// day
		date = nowTime.Format("02")
	case b == 'H':
		//hours
		date = nowTime.Format("15")
	case b == 'i':
		//minute
		date = nowTime.Format("04")
	case b == 's':
		//second
		date = nowTime.Format("05")
	case b == 'W':
		//本年份中的第几周
		_, week := nowTime.ISOWeek()
		date = fmt.Sprintf("%d", week)
	case b == 'w':
		// 星期中的第几天
		w := nowTime.Weekday().String()
		for i := 0; i < len(days); i++ {
			if days[i] == w {
				date = string(i)
				break
			}
		}
	case b == 'l':
		//星期几， 完整的格式，Sunday 到 Saturday
		date = nowTime.Weekday().String()
	case b == 'z':
		//年份中第几天
		date = fmt.Sprintf("%d", nowTime.YearDay())
	}
	return date
}

// 获取当前的时间 - Unix时间戳
func CurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func CurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func CurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

// unix时间戳 转换字符串
func FromUnixTime(ti int64, format string) string {
	t := time.Unix(ti, 0)
	var str string
	for i := 0; i < len(format); i++ {
		str += dateFormat(format[i], t)
		if format[i] == '/' {
			str += "/"
		}
		if format[i] == '-' {
			str += "-"
		}
		if format[i] == ' ' {
			str += " "
		}
		if format[i] == ':' {
			str += ":"
		}
	}
	return str
}

var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Parse Date use PHP time format.
func DateParse(dateString, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, dateString, time.Local)
}

func DateToUnix(dateString, format string) int64 {
	tm, _ := DateParse(dateString, format)
	return tm.Unix()
}
