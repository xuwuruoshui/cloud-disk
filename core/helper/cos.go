package helper

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// 文件上传到腾讯云
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	file, header, err := r.FormFile("file")
	key := "cloud-disk/" + UUID() + path.Ext(header.Filename)

	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}

// 分片上传初始化
func CosInitPartUpload(ext string) (string, string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})
	// 文件名
	key := "cloud-disk/" + UUID() + ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}

	return key, v.UploadID, nil
}

// 分片上传
func CosPartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	key := r.PostForm.Get("key")
	uploadId := r.PostForm.Get("uploadId")
	partNumStr := r.PostForm.Get("partNum")
	partNum, err := strconv.Atoi(partNumStr)
	if err!=nil{
		return "", err
	}

	f, _, err := r.FormFile("file")
	if err!=nil{
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadId, partNum, buf, nil,
	)
	if err != nil {
		return "",err
	}

	return strings.Trim(resp.Header.Get("ETag"),"\""),nil
}

// 分片上传完成
func CosPartUploadComplete(key,uploadId string,cosObject []cos.Object) (error){
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts,cosObject...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	if err != nil {
		return err
	}
	return nil
}