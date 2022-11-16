package movice

// Import resty into your code and refer it as `resty`.
import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/qiniu/qmgo"
)

//获取电影数据
type MoviceService struct {
}

var client = resty.New()
var ctx = context.Background()
var cli *qmgo.QmgoClient

func init() {
	cli, _ = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://t.deey.top:57890", Database: "movicego", Coll: "movice"})
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
}

//https://api.apibdzy.com/api.php/provide/vod/?ac=detail&pg=2&h=40
var url = "https://api.apibdzy.com/api.php/provide/vod/"

func (moviceService *MoviceService) GetMovice(pg string, h string) {
	var moviceReq = make(map[string]string)
	moviceReq["ac"] = "detail"
	moviceReq["pg"] = pg
	moviceReq["h"] = h
	result := &MoviceResp{}
	_, err := client.R().SetQueryParams(moviceReq).SetResult(result).ForceContentType("application/json").Get(url)
	if err != nil {
		fmt.Printf("\"请求出现错误\": %v\n", "请求出现错误")
		fmt.Printf("err: %v\n", err)
	} else {
		b := cli.Bulk()
		for _, v := range result.List {
			b.Upsert(qmgo.M{"vod_id": v.VodId}, v)
		}
		// fmt.Printf("b: %v\n", b)
		r, err := b.Run(ctx)
		if err != nil {
			fmt.Printf("插入mongodb出错: %v\n", err)
		} else {
			fmt.Printf("\"捶入mongodb成功\": %+v\n", r)
		}
	}

}
