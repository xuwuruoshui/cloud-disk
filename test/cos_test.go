package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByFilePath(t *testing.T){
	u, _ := url.Parse("https://1-1258379676.cos.ap-chengdu.myqcloud.com/1-1258379676")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	key := "cloud-disk/test.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/bg10.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}


func TestFileUploadByReader(t *testing.T){
	u, _ := url.Parse("https://1-1258379676.cos.ap-chengdu.myqcloud.com/1-1258379676")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	key := "cloud-disk/2.jpg"

	f, err := os.ReadFile("./img/bg10.jpg")
	if err != nil {
		t.Fatal("read file err: ",err)
	}
	_, err = client.Object.Put(context.Background(), key,bytes.NewReader(f) , nil)
	if err != nil {
		t.Fatal("put file err: ",err)
	}
}