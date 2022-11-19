package main

import (
	"fmt"
	"time"

	"github.com/jy00566722/movies/movice"
	"github.com/robfig/cron/v3"
)

//定时任务
func MyCron() {
	fmt.Println("启动定时任务:....")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 30 * * * *", func() {
		fmt.Printf("运行定时任务:%v\n", time.Now())
		go movice.MoviceCtronGetDate()
	})
	c.Start()
	fmt.Println("定时任务启动成功!....")
}
