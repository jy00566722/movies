package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jy00566722/movies/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DelBzVersion() {
	maxKeys := int64(1000)
	// versionIdMarker := "99"
	keyMarker := "20221124/1669291699032200psttun.jpeg"
	o, err := global.GLM_s3Client.ListObjectVersions(&s3.ListObjectVersionsInput{
		Bucket:    global.GLM_bucket,
		MaxKeys:   &maxKeys,
		KeyMarker: &keyMarker,
		// VersionIdMarker: &versionIdMarker,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("NextVersionIdMarker: %v\nNextKeyMarke%v\n", *o.NextVersionIdMarker, *o.NextKeyMarker)
		for _, object := range o.DeleteMarkers {
			fileName := object.Key
			versionId := object.VersionId
			err := global.GLM_mongodbColl.FindOne(context.TODO(), bson.D{{Key: "bz_pic", Value: fileName}}).Err()
			if err != nil {
				if err == mongo.ErrNoDocuments {
					//文件名不在数据库中,需要把它从BZ中删除
					fmt.Printf("准备删除:%v\n", *fileName)
					delOut, err := global.GLM_s3Client.DeleteObject(&s3.DeleteObjectInput{
						Bucket:    global.GLM_bucket,
						Key:       fileName,
						VersionId: versionId,
					})
					if err != nil {
						fmt.Printf("删除失败: %v\n", err)
						// fmt.Println("==================")
					} else {
						fmt.Printf("删除成功: %v\n", *delOut.VersionId)
						// fmt.Println("==================")
					}
					time.Sleep(300 * time.Millisecond)
				} else {
					fmt.Printf("\"在这里哟\": %v\n", "在这里哟")
				}
			} else {
				fmt.Printf("不能删除有引用的文件: %v\n", *fileName)
				fmt.Println("==================")
			}
		}
		for _, object := range o.Versions {
			fileName := object.Key
			versionId := object.VersionId
			err := global.GLM_mongodbColl.FindOne(context.TODO(), bson.D{{Key: "bz_pic", Value: fileName}}).Err()
			if err != nil {
				if err == mongo.ErrNoDocuments {
					//文件名不在数据库中,需要把它从BZ中删除
					fmt.Printf("准备删除:%v\n", *fileName)
					delOut, err := global.GLM_s3Client.DeleteObject(&s3.DeleteObjectInput{
						Bucket:    global.GLM_bucket,
						Key:       fileName,
						VersionId: versionId,
					})
					if err != nil {
						fmt.Printf("删除失败: %v\n", err)
						// fmt.Println("==================")
					} else {
						fmt.Printf("删除成功: %v\n", *delOut.VersionId)
						// fmt.Println("==================")
					}
					time.Sleep(300 * time.Millisecond)
				} else {
					fmt.Printf("\"在这里哟\": %v\n", "在这里哟")
				}
			} else {
				fmt.Printf("不能删除有引用的文件: %v\n", *fileName)
				fmt.Println("==================")
			}

		}
		fmt.Println("此轮结束,")
		fmt.Printf("NextVersionIdMarker: %v\nNextKeyMarke: %v\n", *o.NextVersionIdMarker, *o.NextKeyMarker)
	}
}

func GTmain() {
	fileName := "girl/1.docx"
	versionId := "4_zb853d947d89f9e0c82440013_f11717d3f8d747ffd_d20221124_m160043_c004_v0402006_t0000_u01669305643008"
	delOut, err := global.GLM_s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket:    global.GLM_bucket,
		Key:       &fileName,
		VersionId: &versionId,
	})
	if err != nil {
		fmt.Printf("删除失败: %v\n", err)
		fmt.Println("==================")
	} else {
		fmt.Printf("删除成功: %+v\n", delOut)
		fmt.Println("==================")
	}
}

func GetObjectTest() {
	// fileName := "girl/1.docx"
	maxKeys := int64(30)
	o, err := global.GLM_s3Client.ListObjectVersions(&s3.ListObjectVersionsInput{
		Bucket:  global.GLM_bucket,
		MaxKeys: &maxKeys,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("o: %v-%v\n", len(o.DeleteMarkers), len(o.Versions))
	}

}
