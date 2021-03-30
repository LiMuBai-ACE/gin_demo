package weather

import (
	"encoding/json"
	"flag"
	"fmt"
	"gin_demo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	LOG_FILE = "./logs/weather.log"
)

var (
	iServices = flag.Bool("s", false, "To running as a services")
	port      = flag.Int("port", 3244, "The TCP port that the server listens on")
	address   = flag.String("address", "", "The net address that the server listens")
	crt       = flag.String("crt", "", "Specify the server credential file")
	key       = flag.String("key", "", "Specify the server key file")
	handle    *Weather
	once      sync.Once
	sigs      = make(chan os.Signal, 1)
	exit      = make(chan bool, 1)
)

//天气日志
func init() {
	flag.CommandLine.Usage = help
	if logout, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		log.SetOutput(logout)
		log.SetPrefix("[Info] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println(time.Now().Format(time.RFC3339), strings.Title(runtime.GOARCH), strings.Title(runtime.GOOS))
	} else {
		fmt.Println(time.Now().Format(time.RFC3339), strings.Title(runtime.GOARCH), strings.Title(runtime.GOOS))
		fmt.Printf("Not found the [ logs ] directory, the log will be displayed on the terminal\n")
	}
}

func getWeatherHandle() (weatherhandle *Weather) {
	once.Do(func() {
		handle = New(DEFAULT_LIMIT_SIZE)
		if err := handle.InitRegionTree(); err != nil {
			log.Println(err)
		}
	})
	return handle
}

//当前城市日历
func ShowWeather(c *gin.Context) {
	city := c.Query("city")
	params := strings.Split(city, ",")
	paramsLen := len(params)
	weatherHandle := getWeatherHandle()
	var err error
	var Resp *WeatherInfo
	fmt.Println(weatherHandle)
	if nil == weatherHandle {
		log.Printf("weatherHandle is nil, please check")
		errResp(c.Writer, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	switch paramsLen {
	case 3:
		Resp, err = weatherHandle.ShowCityWeather(params[0], params[1], params[2])
	case 2:
		Resp, err = weatherHandle.ShowCityWeather(params[0], params[1], params[1])
	case 1:
		Resp, err = weatherHandle.ShowCityWeather(params[0], params[0], params[0])
	default:
		errResp(c.Writer, http.StatusBadRequest, "parameter error")
		return
	}
	if nil != err {
		errResp(c.Writer, http.StatusBadRequest, err.Error())
		return
	}
	if nil == Resp {
		errResp(c.Writer, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	//jsonStr, err := json.Marshal(Resp)  // 序列化转字符串
	if nil != err {
		errResp(c.Writer, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   Resp,
		"msg":    errmsg.GetErrmsg(200),
	})
}

//城市地址
func ShowCityList(c *gin.Context) {
	weatherHandle := getWeatherHandle()
	city := c.Query("city")
	params := strings.Split(city, ",")
	if city != "" {
		resp, _ := weatherHandle.ShowCityList(params[0])
		var m interface{}
		err := json.Unmarshal(resp, &m)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, m)
	} else {
		resp, _ := weatherHandle.ShowCityList("")
		var m interface{}
		err := json.Unmarshal(resp, &m)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, m)
	}
}

func ShowStatus(c *gin.Context) {
	weatherHandle := getWeatherHandle()
	status := weatherHandle.Stats()
	//strStatus, _ := json.Marshal(status)  // 序列化 转字符串
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   status,
		"msg":    errmsg.GetErrmsg(200),
	})
	//c.Writer.Write(status)
}

func errResp(w gin.ResponseWriter, rCode int, rMsg string) {
	var Jmap = make(map[string]interface{})
	Jmap[RESP_RCODE_FIELD] = rCode
	Jmap[RESP_RMSG_FIELD] = rMsg
	strResp, _ := json.Marshal(Jmap)
	w.Write(strResp)
}

func help() {
	fmt.Printf("Provide weather access interface based on laboratory environment.\n")
	fmt.Printf("Usage: %s [OPTION]...\n", filepath.Base(os.Args[0]))
	fmt.Println("     -s\t\tSet process running as a services, using [false] by default")
	fmt.Println("     -address\tSet the listener address, using [0.0.0.0] by default")
	fmt.Println("     -port\tSet the listener port, using port [3244] by default")
	fmt.Println("     -crt\tSpecify the server credential file")
	fmt.Println("     -key\tSpecify the server key file")
	fmt.Println("     -help\tdisplay help info and exit")
}
