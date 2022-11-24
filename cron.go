package main

import (
	"fmt"
	"time"

	"github.com/jy00566722/movies/global"
	"github.com/jy00566722/movies/movice"
	"github.com/robfig/cron/v3"
)

//定时任务
func MyCron() {
	fmt.Println("启动定时任务:....")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 40 * * * *", func() {
		fmt.Printf("运行定时任务:%v\n", time.Now())
		go movice.MoviceCtronGetDate() //定时获取电影信息
	})
	c.AddFunc("50 1 * * * *", func() {
		fmt.Printf("运行定时任务:%v\n", time.Now())
		if !global.GLM_BZCRONSTATUS {
			global.GLM_BZCRONSTATUS = true
			go movice.SaveImageFormDbToBz() //定时搬运影视图片到BZ
		} else {
			fmt.Println("BZ搬运图片任务已经在运行")
		}

	})
	c.Start()
	fmt.Println("定时任务启动成功!....")
}
