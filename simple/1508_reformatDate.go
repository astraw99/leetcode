package main

import (
	"fmt"
	"strings"
)

func main() {
	date := "1st Oct 2052"
	fmt.Println(reformatDate(date))
}

// reformatDate
// 给你一个字符串 date ，它的格式为 Day Month Year ，其中：
// Day 是集合 {"1st", "2nd", "3rd", "4th", ..., "30th", "31st"} 中的一个元素。
// Month 是集合 {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"} 中的一个元素。
// Year 的范围在 ​[1900, 2100] 之间。
func reformatDate(date string) string {
	monthMap := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	dateArr := strings.Split(date, " ")
	year := dateArr[2]
	month := ""
	if v, ok := monthMap[dateArr[1]]; ok {
		month = v
	}

	dayStr := dateArr[0]
	day := dayStr[:len(dayStr)-2]
	if len(day) < 2 {
		day = "0" + day
	}

	return year + "-" + month + "-" + day
}
