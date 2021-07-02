package model

import (
	"context"
	"fmt"
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
)

var accessKey = utils.Data.Qiniu.Accesskey
var secretKey = utils.Data.Qiniu.Secretkey
var bucket = utils.Data.Qiniu.Bucket
var imgUrl = utils.Data.Qiniu.Server

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: true,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.Put(context.Background(), &ret, upToken, "111.png", file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	fmt.Println(ret)
	url := imgUrl + ret.Key
	return url, errmsg.SUCCSE
}
