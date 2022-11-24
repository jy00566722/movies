package main

import "github.com/jy00566722/movies/movice"

func main() {
	// MyCron()
	// go movice.MoviceCtronGetDate() //获取数据信息
	go movice.SaveImageFormDbToBz() //搬运图片到BZ
	r := SetupRouter()              //原框架默认路由
	r.Run(":8080")

}
