package global

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	GLM_BZCRONSTATUS     bool = false //搬运图片上BZ的任务状态
	GLM_BZCRONSTATUS1080 bool = false //搬运图片上BZ的任务状态
	CLM_mongodbClient    *mongo.Client
	GLM_mongodbColl      *mongo.Collection
	GLM_s3Client         *s3.S3
	GLM_bucket           *string

	QmgoClient         *qmgo.Client
	QmgoDatabase       *qmgo.Database
	QmgoCollMovice     *qmgo.Collection
	QmgoCollMovice1080 *qmgo.Collection
)

const uri = "mongodb://t.deey.top:57890/?maxPoolSize=20&w=majority"

func GlobalInit() {
	CLM_mongodbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	GLM_mongodbColl = CLM_mongodbClient.Database("movicego").Collection("movice")

	GLM_bucket = aws.String("oeoli-movice")
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("00483978fec24030000000001", "K004HuS8EXcW4IoJX1bU/meK18nrI1s", ""),
		Endpoint:         aws.String("https://s3.us-west-004.backblazeb2.com"),
		Region:           aws.String("us-west-004"),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		fmt.Println("Err:", err)
	}
	GLM_s3Client = s3.New(newSession)

	//qmgo初始化
	ctx := context.Background()
	QmgoClient, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://t.deey.top:57890"})
	if err != nil {
		panic(err)
	}
	QmgoDatabase := QmgoClient.Database("movicego")
	QmgoCollMovice = QmgoDatabase.Collection("movice")
	QmgoCollMovice1080 = QmgoDatabase.Collection("movice1080")
}
