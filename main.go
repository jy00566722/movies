package main

// var moviceService movice.MoviceService

func main() {
	MyCron()
	r := SetupRouter() //原框架默认路由
	r.Run(":8080")

}
