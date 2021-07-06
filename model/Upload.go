package model

import (
	"context"
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.Data.Qiniu.Accesskey
var SecretKey = utils.Data.Qiniu.Secretkey
var Bucket = utils.Data.Qiniu.Bucket
var ImgUrl = utils.Data.Qiniu.Server

func UpLoadFile(file multipart.File, fileSize int64, fileName string) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: true,
		UseHTTPS:      true,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{
		Key: fileName,
	}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCSE
}
