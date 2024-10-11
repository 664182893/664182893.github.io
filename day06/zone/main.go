package main

import (
	"fmt"
	"time"
)

// func f1() {
// }

func f2() {
	now := time.Now() //当前时间
	fmt.Println(now)

	//根据东八区的时区和格式去解析一个字符串格式的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		if err != nil {
			fmt.Printf("locate time failed,err:%v\n", err)
			return
		}
	}
	//按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2024-10-11 19:59:00", loc)
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	f2()
}
