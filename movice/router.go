package movice

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/duke-git/lancet/random"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoadMoviceRouter(e *gin.Engine) {

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
		fmt.Println("====所有请求完成!====")

	}
}

//从数据库中提取图片未保存到BZ的记录，把影视图片搬到BZ长久保存
func SaveImageFormDbToBz() {
	var moviesInfo []Movice
	//$exists
	err := cli.Find(ctx, bson.M{"bz_pic": nil}).Sort("_id").Limit(100).All(&moviesInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for _, v := range moviesInfo {
			time.Sleep(time.Second)
			if v.VodPic != "" {
				loadFileinfo, err := GetImageFromUrl(v.VodPic)
				if err != nil {
					fmt.Printf("err: %v\n", err)
					err := cli.UpdateOne(ctx, bson.M{"vod_id": v.VodId}, bson.M{"$set": bson.M{"bz_pic": "BAD"}})
					if err != nil {
						fmt.Printf("插入数据库出错1: %v\n", err)
					} else {
						fmt.Printf("插入搬运错误信息BAD进入mongdb: %v\n", v.VodName)
					}
				} else {
					// fmt.Printf("上传成功后的信息: %+v\n", loadFileinfo.FileName)
					//https://f004.backblazeb2.com/file/oeoli-movice/20221120/1668878893235997JfcZCw.jpeg
					err := cli.UpdateOne(ctx, bson.M{"vod_id": v.VodId}, bson.M{"$set": bson.M{"bz_pic": loadFileinfo.FileName}})
					if err != nil {
						fmt.Printf("插入数据库出错: %v\n", err)
					} else {
						fmt.Printf("\"插入mongodb成功\": %v\n", v.VodName)
					}

				}
			}
		}

	}
}

//从url获取图片并上传到bz,返回文件名与唯一ID
func GetImageFromUrl(url string) (loadFileinfo LoadFileinfo, err error) {
	// fmt.Println("开始搬图片到BZ:")
	resp, err := moviceService.GetImageFromUrlf(url)
	if err != nil {
		fmt.Printf("从url获取图片出错: %v\n", err)
		return loadFileinfo, err
	}
	fmt.Println("从URL获取图像成功")
	if resp.Size() > 1024*1024*3 {
		fmt.Printf("图片大小超过3M")
		return loadFileinfo, err
	}
	d := resp.Header().Get("content-type")
	// fmt.Printf("d:图片类型为: %v\n", d)
	d_s := strings.Split(d, "/")
	extstring := "." + d_s[1]
	d_s_1 := strings.ToUpper(d_s[1])
	testArray := []string{"JPG", "JPEG", "PNG", "GIF"}
	testFlag := false
	for _, v := range testArray {
		if v == d_s_1 {
			testFlag = !testFlag
			break
		}
	}
	if !testFlag {
		return loadFileinfo, errors.New("获取图片的类型不正确:" + d_s[1])
	}
	now := time.Now()
	fileName := now.Format("20060102") + "/" + strconv.FormatInt(now.UnixMicro(), 10) + random.RandString(6) + extstring
	contentType := "image/" + extstring[1:]
	imgBodyReader := bytes.NewReader(resp.Body())
	fileInfo, err := moviceService.UploadFile(&fileName, &contentType, imgBodyReader)
	// fileInfo, err := UploadFile(&fileName, &contentType, imgBodyReader)
	if err != nil {
		fmt.Printf("上传失败了啊,%v\n", err)
		return loadFileinfo, err
	}
	loadFileinfo.FileName = fileName
	loadFileinfo.B2OutPut = fileInfo
	loadFileinfo.Size = resp.Size()
	// fmt.Printf("上传成功后的信息: %+v\n", loadFileinfo)
	return loadFileinfo, nil
}
