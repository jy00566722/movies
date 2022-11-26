package movice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/duke-git/lancet/random"
	"github.com/gin-gonic/gin"
	"github.com/jy00566722/movies/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://t.deey.top:57890/?maxPoolSize=20&w=majority"

var clientM *mongo.Client
var cliM *mongo.Collection

func init() {
	clientM, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// if err != nil {
	// 	panic(err)
	// }
	cliM = clientM.Database("movicego").Collection("movice")
}

func GetDataTest() {
	fmt.Println("测试.")
	filter := bson.D{{Key: "vod_name", Value: bson.D{{Key: "$all", Value: bson.A{"大使的厨房"}}}}}
	cursor, err := cliM.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Movice
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

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
		FixImgUrl()

	}
}

//从数据库中提取图片未保存到BZ的记录，把影视图片搬到BZ长久保存
func SaveImageFormDbToBz() {
	var moviesInfo []Movice
	err := cli.Find(ctx, bson.M{"bz_pic": "BAD"}).Sort("-_id").Limit(1000).All(&moviesInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for i, v := range moviesInfo {
			time.Sleep(time.Microsecond * 500)
			if v.VodPic != "" && v.BzPic == "BAD" {
				fmt.Printf("要搬运的信息:VodPic: %v,BzPic:%v\n", v.VodPic, v.BzPic)
				loadFileinfo, err := GetImageFromUrl(v.VodPic)
				if err != nil {
					fmt.Printf("err: %v\n", err)
					err := cli.UpdateOne(ctx, bson.M{"vod_id": v.VodId}, bson.M{"$set": bson.M{"bz_pic": "BAD"}})
					if err != nil {
						fmt.Printf("插入数据库出错1: %v\n", err)
					} else {
						fmt.Printf("插入搬运错误信息BAD进入mongdb: %v,%v\n", v.VodName, v.VodPic)
					}
				} else {
					// fmt.Printf("上传成功后的信息: %+v\n", loadFileinfo.FileName)
					//https://f004.backblazeb2.com/file/oeoli-movice/20221120/1668878893235997JfcZCw.jpeg
					err := cli.UpdateOne(ctx, bson.M{"vod_id": v.VodId}, bson.M{"$set": bson.M{"bz_pic": loadFileinfo.FileName}})
					if err != nil {
						fmt.Printf("插入数据库出错: %v\n", err)
					} else {
						fmt.Printf("\"插入mongodb成功%v\": %v\n", i, v.VodName)
					}
				}
			} else {
				fmt.Printf("不用搬运的信息:%v: %v,BzPic:%v\n", v.VodName, v.VodPic, v.BzPic)
			}
		}
		fmt.Println("本轮搬运完成")
		global.GLM_BZCRONSTATUS = false
	}
}

//修整天vod_pic的地址问题,前面多的域名去掉
func FixImgUrl() {
	var moviesInfo []Movice
	err := cli.Find(ctx, bson.M{}).Sort("-_id").Limit(500).All(&moviesInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		reg1, err := regexp.Compile(`(https{0,1}:\/\/.{1,}?\/)(http.{1,})`)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		//http://bdzyimg.com/https://img.liangzipic.com/upload/vod/20221116-1/ec4c3cc14323210c0df2c853a6016955.jpg#err2022-11-20
		for _, v := range moviesInfo {
			flag := reg1.Match([]byte(v.VodPic))
			if flag {
				// fmt.Printf("  不   正确的VodPic: %v\n", v.VodPic)
				newPic := reg1.ReplaceAllString(v.VodPic, "$2")
				// fmt.Printf("newPic: %v\n", newPic)
				err := cli.UpdateOne(ctx, bson.M{"vod_id": v.VodId}, bson.M{"$set": bson.M{"vod_pic": newPic}})
				if err != nil {
					fmt.Printf("err: %v\n", err)
				} else {
					fmt.Printf("newPic修改成功: %v\n", newPic)
				}
			}
		}
		fmt.Printf("\"本轮FixPicUrl结束\": %v\n", "本轮结束")
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
	// fmt.Println("从URL获取图像成功")
	if resp.Size() > 1024*1024*3 {
		fmt.Printf("图片大小超过3M")
		return loadFileinfo, err
	}
	d := resp.Header().Get("content-type")
	d_s := strings.Split(d, "/")
	if len(d_s) < 2 {
		fmt.Printf("d:图片类型为: %v\n", d)
		err = errors.New("d_s长度小于2,不能判断出图片类型")
		return loadFileinfo, err
	}
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
