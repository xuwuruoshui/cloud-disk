package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
)

// 文件分片
const chunkSize = 1 * 1024 * 1024 //100MB
func TestGenerateChunkFile(t *testing.T) {
	// 1.获取文件信息
	fileInfo, err := os.Stat("bg2.jpg")
	if err != nil {
		t.Fatal(err)
	}

	// 2.计算分片个数
	chunkNum := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))

	// 3.获取文件
	myFile, err := os.OpenFile("bg2.jpg", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < chunkNum; i++ {
		// 指定读取文件的起始位置
		myFile.Seek(int64(i*chunkNum), 0)

		// 如果分片大于剩余的数据, 总量-i*一次分片=剩余需要直接读取的大小
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)
		// 写入
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 文件分片的合并
func TestMergeChunkFile(t *testing.T) {
	myfile, err := os.OpenFile("bg2.jpg", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 6; i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myfile.Write(b)
		f.Close()
	}
	myfile.Close()
}

// 文件一致性校验
func TestCheck(t *testing.T) {
	// 获取第一个文件的信息
	file, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	// 获取第二个文件的信息
	file2, err := os.OpenFile("fff.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}

	data2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}

	fh1 := fmt.Sprintf("%x", md5.Sum(data))
	fh2 := fmt.Sprintf("%x", md5.Sum(data2))

	fmt.Println(fh1)
	fmt.Println(fh2)
	fmt.Println(fh1 == fh2)
}

// 腾讯云存储
// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://1-1258379676.cos.ap-chengdu.myqcloud.com/1-1258379676")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})
	key := "cloud-disk/qwq.jpg"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID //1663035726701b93cb3354952ba066e2263318c1a205e3822fe2f40295df73328173d6c9a2
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("https://1-1258379676.cos.ap-chengdu.myqcloud.com/1-1258379676")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})
	key := "cloud-disk/qwq.jpg"
	uploadId := "1663035726701b93cb3354952ba066e2263318c1a205e3822fe2f40295df73328173d6c9a2"

	for i := 0; i < 4; i++ {
		f, err := os.ReadFile(fmt.Sprintf("%d.chunk", i)) // .chunk的md5值: 2d44f02749e59809dbd6ff69152f92df
		if err != nil {
			t.Fatal(err)
		}
		// opt可选
		resp, err := client.Object.UploadPart(
			context.Background(), key, uploadId, i+1, bytes.NewReader(f), nil,
		)
		if err != nil {
			panic(err)
		}
		PartETag := resp.Header.Get("ETag")
		fmt.Println(PartETag)
	}

}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://1-1258379676.cos.ap-chengdu.myqcloud.com/1-1258379676")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})
	key := "cloud-disk/qwq.jpg"
	uploadId := "1663035726701b93cb3354952ba066e2263318c1a205e3822fe2f40295df73328173d6c9a2"
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts,
		cos.Object{PartNumber: 1, ETag: "a08b370829123bc3fbfd3c520d6b40a1"},
		cos.Object{PartNumber: 2, ETag: "67f49428570f8e45dad8fc9768072156"},
		cos.Object{PartNumber: 3, ETag: "933275d31dc9e19cb53718ef6a208668"},
		cos.Object{PartNumber: 4, ETag: "97470ee7c1435d98ee22245335fcee3b"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
