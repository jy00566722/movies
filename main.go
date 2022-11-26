package main

import (
	"github.com/jy00566722/movies/global"
	"github.com/jy00566722/movies/movice"
)

func main() {
	global.GlobalInit() //初始化
	// MyCron()
	// go movice.MoviceCtronGetDate()  //获取数据信息
	go movice.SaveImageFormDbToBz() //搬运图片到BZ
	// go DelBzVersion()
	r := SetupRouter() //原框架默认路由
	r.Run(":8080")

}
