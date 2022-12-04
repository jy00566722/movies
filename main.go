package main

import (
	"github.com/jy00566722/movies/global"
)

func main() {
	global.GlobalInit() //初始化
	MyCron()
	// go movice.MoviceCtronGetDate() //获取数据信息
	// go movice.Movice1080CtronGetDate() //获取数据信息1080
	// go movice.SaveImageFormDbToBz() //搬运图片到BZ
	// go movice.SaveImageFormDbToBz1080() //搬运图片到BZ1080
	// go DelBzVersion()
	// go FindTest()
	// go movice.MergeMovies()
	r := SetupRouter() //原框架默认路由
	r.Run(":8080")

}
