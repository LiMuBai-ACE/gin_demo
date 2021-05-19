package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

//完整方法
type MyTime struct {
	time.Time
}

//重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；程序中解析到JSON
func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

//JSON中解析到程序中
func (t *MyTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+"2006-01-02 15:04:05"+`"`, string(data), time.Local)
	*t = MyTime{Time: now}
	return
}

//写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
