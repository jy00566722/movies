package movice

// Import resty into your code and refer it as `resty`.
import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

//从url获取图片
func (moviceService *MoviceService) GetImageFromUrlf(imgUrl string) (*resty.Response, error) {
	client := resty.New()
	resp, err := client.R().Get(imgUrl)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//上传图片
func (moviceService *MoviceService) UploadFile(fileName *string, contentType *string, fileBody io.ReadSeeker) (*s3.PutObjectOutput, error) {
	bucket := aws.String("oeoli-movice")
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("00483978fec24030000000001", "K004HuS8EXcW4IoJX1bU/meK18nrI1s", ""),
		Endpoint:         aws.String("https://s3.us-west-004.backblazeb2.com"),
		Region:           aws.String("us-west-004"),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	s3Client := s3.New(newSession)

	//Upload a new object "testfile.txt" with the string "S3 Compatible API" io.ReadSeeker
	file, err := s3Client.PutObject(&s3.PutObjectInput{
		Body:        fileBody,
		Bucket:      bucket,
		Key:         fileName,
		ContentType: contentType,
	})
	if err != nil {
		fmt.Printf("Failed to upload object %s/%s, %s\n", *bucket, *fileName, err.Error())
		return nil, err
	}

	return file, nil
}
