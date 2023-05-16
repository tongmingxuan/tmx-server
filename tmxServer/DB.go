// Package tmxServer /*
package tmxServer

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
	"unsafe"
)

type InterfaceModel interface {
	//TableName 获取当前表名称
	TableName() string
	//GetPollName 返回默认mysql连接池Name
	GetPollName() string
	//GetConnection 获取链接
	GetConnection() *gorm.DB
}

// Connection
// @Description: 获取链接
// @param connection
// @return *gorm.DB
func Connection(connection string) *gorm.DB {
	dbPoll := (*DatabasePool)(unsafe.Pointer(frameContainer.Get(new(DatabasePool).Key())))

	conn, ok := dbPoll.Pool.Load(connection)

	if ok != true {
		panic("dbPoll.Pool.Load(connection):error")
	}

	return conn.(*gorm.DB)
}

// BaseModel
// @Description: 基础模型
type BaseModel struct {
	CreatedAt MyTime  `json:"created_at"`
	UpdatedAt MyTime  `json:"updated_at"`
	DeletedAt *MyTime `json:"deleted_at"`
}

// Model
// @Description: 获取模型实例
// @receiver m
// @param model
// @return *gorm.DB
func (m BaseModel) Model(model InterfaceModel) *gorm.DB {
	connection := model.GetPollName()

	if connection == "" {
		panic("GetModel:model.GetConnection():返回空")
	}

	conn := Connection(connection)

	return conn.Model(model)
}

// MyTime 自定义时间
type MyTime struct {
	Time time.Time
}

func Now() MyTime {
	return MyTime{Time: time.Now()}
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	// 前端接收的时间字符串
	str := string(data)
	// 去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")

	// 空字符串处理
	if timeStr == "" {
		*t = MyTime{Time: time.Time{}}
		return err
	}

	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime{Time: t1}
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	// 0值置空
	var tt string
	if t.Time.IsZero() {
		tt = ""
	} else {
		tt = t.Time.Format("2006-01-02 15:04:05")
	}
	formatted := fmt.Sprintf("\"%v\"", tt)
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	// tTime := time.Time(t.Time)
	return t.Time.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime{Time: vt}
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format("2006-01-02 15:04:05")
}
