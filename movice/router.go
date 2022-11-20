package movice

import (
	"bytes"
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

	}
}

//从数据库中提取图片未保存到BZ的记录，把影视图片搬到BZ长久保存
func SaveImageFormDbToBz() {
	var moviesInfo []Movice
	err := cli.Find(ctx, bson.M{"vod_pic_slide": bson.M{"$eq": ""}}).Limit(3).All(&moviesInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for _, v := range moviesInfo {
			// fmt.Printf("视频名称: %v,原图片地址:%v,BZ图片地址:%v\n", v.VodName, v.VodPic, v.VodPicSlide)
			fmt.Printf("视频名称: %v\n", v.VodName)
			if v.VodPicSlide == "" && v.VodPic != "" {
				loadFileinfo, err := GetImageFromUrl(v.VodPic)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				} else {
					fmt.Printf("上传成功后的信息: %+v\n", loadFileinfo.FileName)
					err := cli.UpdateOne(ctx, bson.M{"vod_id": v.VodId}, bson.M{"$set": bson.M{"vod_pic_slide": loadFileinfo.FileName}})
					if err != nil {
						fmt.Printf("插入数据库出错: %v\n", err)
					} else {
						fmt.Printf("\"插入mongodb成功\": %v\n", "")
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

	// fmt.Printf("extstring: %v\n", extstring)
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
