package main

import "github.com/jy00566722/movies/movice"

// var moviceService movice.MoviceService

func main() {
	MyCron()
	// go movice.GetImageFromUrl("https://bdzyimg.com/upload/vod/20220306-1/4858006906045f8b56bca30ed1399601.jpg")
	go movice.SaveImageFormDbToBz()
	r := SetupRouter() //原框架默认路由
	r.Run(":8080")

}
