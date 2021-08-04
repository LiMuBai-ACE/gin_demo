package model

import (
	"encoding/base64"
	"fmt"
	"gin_demo/utils"
	"gin_demo/utils/errmsg"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

var AccessKey = utils.Data.Qiniu.Accesskey
var SecretKey = utils.Data.Qiniu.Secretkey
var Bucket = utils.Data.Qiniu.Bucket
var ImgUrl = utils.Data.Qiniu.Server

type File struct {
	File string `json:"file"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func UpLoadFile(file File) (string, int) {

	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, file.File)
	fmt.Println(b, file.File)
	if !b {
		return "", errmsg.ERROR
	}

	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	//allData := re.FindAllSubmatch([]byte(file.File), 2)
	//fileType := string(allData[0][1]) //png ，jpeg 后缀获取

	base64Str := re.ReplaceAllString(file.File, "")

	date := time.Now().Format("2006-01-02")
	if ok := utils.IsFileExist("./images/" + date); !ok {
		os.Mkdir("./images/"+date, 0666)
	}

	//curFileStr := strconv.FormatInt(time.Now().UnixNano(), 10)

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//n := r.Intn(99999)

	//var file1 string = "./images" + "/" + date + "/" + curFileStr + strconv.Itoa(n) + "." + fileType
	var file1 string = "./images" + "/" + date + "/" + file.Name
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(file1, byte, 0666)
	if err != nil {
		return "图片上传失败", errmsg.ERROR
	}

	//putPolicy := storage.PutPolicy{
	//	Scope: Bucket,
	//}
	//mac := qbox.NewMac(AccessKey, SecretKey)
	//upToken := putPolicy.UploadToken(mac)
	//
	//cfg := storage.Config{
	//	Zone:          &storage.ZoneHuadong,
	//	UseCdnDomains: true,
	//	UseHTTPS:      true,
	//}
	//
	//putExtra := storage.PutExtra{}
	//
	//formUploader := storage.NewFormUploader(&cfg)
	//ret := storage.PutRet{
	//	Key: fileName,
	//}
	//
	//err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	//if err != nil {
	//	return "", errmsg.ERROR
	//}
	//url := ImgUrl + ret.Key
	return "", errmsg.SUCCSE
}
