package model

import (
	"bytes"
	"context"
	"fmt"
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var accessKey = utils.Data.Qiniu.Accesskey
var secretKey = utils.Data.Qiniu.Secretkey
var bucket = utils.Data.Qiniu.Bucket
var imgUrl = utils.Data.Qiniu.Server

type File struct {
	FileName string `json:"file_name"`
	File     string `json:"file"`
	Type     string `json:"type"`
}

func UpLoadFile(file string, fileName string, fileType string) (string, int) {
	//鉴权
	mac := qbox.NewMac(accessKey, secretKey)
	//上传策略
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//获取上传token
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
	//上传Config对象
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 上传区域
		UseCdnDomains: false,                // 是否使用https域名加速
		UseHTTPS:      false,                // 是否使用CDN上传加速
	}

	putExtra := storage.PutExtra{}
	//构建上传对象
	base64Uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	var num uint = 0
	if len(fileType) == 9 {
		num = 22
	} else if len(fileType) == 10 {
		num = 23
	}
	file = string([]byte(file)[num:])
	// 图片base64格式的数据 注意 需要去掉 前面类似data:image/png;base64,的数据
	data := []byte(file)
	dataLen := int64(len(data))
	err := base64Uploader.Put(context.Background(), &ret, upToken, fileName, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	fmt.Println(ret)
	url := imgUrl + ret.Key
	return url, errmsg.SUCCSE
}
