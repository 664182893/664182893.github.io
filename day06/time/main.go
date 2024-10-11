package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(time.Unix(1564803667, 0))

	//时间间隔
	fmt.Println(time.Second)
	//now +24小时
	fmt.Println(now.Add(24 * time.Millisecond * 1000 * 60 * 60))

	//格式化时间
	t, _ := time.Parse(time.Stamp, "Jan 11,2024 at 7.12am(PST)")
	fmt.Println(time.Stamp)
	fmt.Println(t)
	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2019-08-03")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	//sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02", "2024-10-11")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	d := now.Sub(nextYear)
	fmt.Println(d)
	fmt.Println("------------------")
	nextYear1 := nextYear.UTC()
	fmt.Println(nextYear1)
	fmt.Println("-----------")
	time.Sleep(time.Duration(100) * time.Second)
	fmt.Println("100s has gone")
}
