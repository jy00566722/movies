package movice

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func LoadMoviceRouter(e *gin.Engine) {

}

//定时任务
func MoviceCron() {
	fmt.Println("启动定时获取电影数据任务任务:")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0 0,13 * * *", func() {
		fmt.Printf("运行定时任务:%v\n", time.Now())
		MoviceCtronGetDate()
	})
	c.Start()
	fmt.Println("定时任务启动成功!....")
}

var moviceService MoviceService

func MoviceCtronGetDate() {
	var url = "https://api.apibdzy.com/api.php/provide/vod/"
	var moviceReq = make(map[string]string)
	moviceReq["ac"] = "list"
	moviceReq["pg"] = "1"
	moviceReq["h"] = "72"
	result := &MoviceResp{}
	_, err := client.R().SetQueryParams(moviceReq).SetResult(result).ForceContentType("application/json").Get(url)
	if err != nil {
		fmt.Printf("\"请求出现错误\": %v\n", "请求出现错误")
		fmt.Printf("err: %v\n", err)
	} else {
		// fmt.Printf("请求回来的电影数据如下: %v\n", result)
		fmt.Printf("result.Total: %v\n", result.Total)
		fmt.Printf("result.Pagecount: %v\n", result.Pagecount)
		for i := 1; i <= result.Pagecount; i++ {
			fmt.Printf("\"当前请求第几页\": %v\n", i)
			moviceService.GetMovice(strconv.Itoa(i), "")
		}

	}
}
