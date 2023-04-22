// Package HelperFunction /*
package HelperFunction

import (
	"encoding/json"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"net/http"
	"time"
)

// GetDateTime
// @Description: 返回年月日时分秒
// @return string
func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// JsonSuccess
// @Description: http成功
// @param message
// @param data
// @return string
func JsonSuccess(message string, data map[string]interface{}) string {
	if data == nil {
		data = map[string]interface{}{}
	}

	mapRes := map[string]interface{}{
		"code":    200,
		"message": message,
		"data":    data,
	}

	requestId, ok := tmxServer.GetContext().Get("requestId")

	if ok {
		mapRes["request_id"] = requestId
	}

	result, err := json.Marshal(mapRes)

	if err != nil {
		panic("JsonSuccess:error:" + err.Error())
	}

	return string(result)
}

func JsonError(message string, data map[string]interface{}) string {
	if data == nil {
		data = map[string]interface{}{}
	}

	mapRes := map[string]interface{}{
		"code":    500,
		"message": message,
		"data":    data,
	}

	requestId, ok := tmxServer.GetContext().Get("requestId")

	if ok {
		mapRes["request_id"] = requestId
	}

	result, err := json.Marshal(mapRes)

	if err != nil {
		panic("JsonSuccess:error:" + err.Error())
	}

	return string(result)
}

func GetQueryParams(c *gin.Context) map[string]interface{} {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]interface{}, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}

func GetPostFormParams(c *gin.Context) map[string]interface{} {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			panic("GetPostFormParams:ErrNotMultipart:" + err.Error())
		}
	}
	var postMap = make(map[string]interface{}, len(c.Request.PostForm))
	for k, v := range c.Request.PostForm {
		if len(v) > 1 {
			postMap[k] = v
		} else if len(v) == 1 {
			postMap[k] = v[0]
		}
	}

	return postMap
}

func GetSnowflakeIdByInt64() string {
	startTime := "2021-12-03"

	var machineID int64 = 1

	var st time.Time

	// 格式化 1月2号下午3时4分5秒  2006年
	st, err := time.Parse("2006-01-02", startTime)

	if err != nil {
		panic("InitExtendLoad:Init:error:" + err.Error())
	}

	snowflake.Epoch = st.UnixNano() / 1e6

	node, err := snowflake.NewNode(machineID)

	if err != nil {
		panic("InitExtendLoad:NewNode:error:" + err.Error())
	}

	return node.Generate().String()
}
