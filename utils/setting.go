package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type DataStruct struct {
	Server struct {
		Mode   string
		Port   string
		Jwtkey string
	}
	Mysql struct {
		Db       string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Qiniu struct {
		Accesskey string
		Secretkey string
		Bucket    string
		Sever     string
	}
}

var Data = new(DataStruct)

func init() {

	if initConfig() != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", initConfig())
	}
	fmt.Printf("%+v", Data)
}

func initConfig() error {
	env := os.Getenv("GIN_ENV")
	fmt.Println(env, "---------------")
	var err error
	var data []byte
	if env == "" {
		data, err = ioutil.ReadFile(`config/` + "local" + ".yaml")
	} else {
		data, err = ioutil.ReadFile(`config/` + env + ".yaml")
	}

	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, Data)
	return nil
}
