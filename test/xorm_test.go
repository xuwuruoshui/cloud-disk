package test

import (
	"bytes"
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T){
	var err error
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(192.168.0.132:3306)/cloud-disk?charset=utf8mb4&parseTime=True")
	if err!=nil{
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err!=nil{
		t.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err!=nil{
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst,b,"","\t")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(dst.String())

}

func TestCrypt(t *testing.T){
	pwd := "123456"
	encrypt, _ := helper.Encrypt(pwd)
	t.Log(encrypt)
	success := helper.Decrypt(pwd, encrypt)
	t.Log(success)
}