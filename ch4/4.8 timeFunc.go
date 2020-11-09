package main

import (
	"fmt"
	"time"
)

func main() {

	// 1.获取当前时间
	t := time.Now() // 生成当前时间对象
	fmt.Println(t)  // 2020-11-09 11:44:21.514546 +0800 CST m=+0.000079631


	// 2.当前时间的年月日时分秒
	fmt.Println("年", t.Year())		// 2020
	fmt.Println("月", t.Month())	// November
	fmt.Println("日", t.Day())		// 9
	fmt.Println("时", t.Hour())		// 11
	fmt.Println("分", t.Minute())	// 45
	fmt.Println("秒", t.Second())	// 50


	// 3.格式化时间
	// 第一种
	timeFormat1 := 	fmt.Sprintf("当前时间：%d-%d-%d %d:%d:%d\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	fmt.Println(timeFormat1) // 当前时间：2020-11-9 11:53:16

	// 第二种(2006-01-02 15:04:05 是Go诞生的日子, 以此代替 YmdHis)
	timeFormat2 := t.Format("2006-01-02 15:04:05")
	fmt.Println(timeFormat2) // 2020-11-09 11:53:16
	timeFormat3 := t.Format("2006-01-02")
	fmt.Println(timeFormat3) // 2020-11-09
	timeFormat4 := t.Format("15:04:05")
	fmt.Println(timeFormat4) // 11:53:16


	// 4.时间戳
	// Unix: 返回当前时间距1970年1月1日有多少秒
	// UnixNano: 返回当前时间距1970年1月1日有多少纳秒
	fmt.Println(time.Now().Unix()) // 1604894880
	fmt.Println(time.Now().UnixNano()) // 1604894880637230000


	// 5.时间常量
	// 休眠 1秒
	time.Sleep(time.Second)
	// 休眠 100 毫秒(0.1秒)
	time.Sleep(time.Millisecond * 100)
	// 休眠1分钟
	time.Sleep(time.Minute)
}