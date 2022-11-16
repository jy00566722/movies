package main

import "github.com/jy00566722/movies/movice"

// var moviceService movice.MoviceService

func main() {
	go movice.MoviceCtronGetDate()
	r := SetupRouter() //原框架默认路由
	r.Run(":8080")

}
