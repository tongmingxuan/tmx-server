// Package tmxServer /*
package tmxServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func AntToString(value interface{}) string {
	var str string
	if value == nil {
		return str
	}
	// vt := value.(type)
	switch value.(type) {
	case float64:
		ft := value.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		str = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = value.(string)
	case []byte:
		str = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		str = string(newValue)
	}

	return str
}

func Dump(v interface{}) {
	jsonFormat, errs := json.MarshalIndent(v, "", "  ")
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(string(jsonFormat))
}

func GetEnv(key string, defaultValue string) string {
	result, isSet := os.LookupEnv(key)

	if isSet {
		return result
	}

	return defaultValue
}

func SetEnv(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic("设置env异常,key:" + key + ":value:" + value)
	}
}

func GetDebugTraceBySlice() []string {
	var res []string

	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		res = append(res, "当前文件名称:"+file+":调用行数:"+strconv.Itoa(line))
	}

	return res
}

func StringToInt(string, errMessage string) int {
	res, err := strconv.Atoi(string)

	if err != nil {
		panic(errMessage + ":" + err.Error())
	}

	return res
}

func GetGoId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
