// Package tmxServer  /*
package tmxServer

// BaseFrame
// @Description: 基础加载接口
type BaseFrame interface {
	Key() string
	Handle() *BaseFrame
}

// DefaultConfig
// tmx-project-frame-tmxCore
// @Description: 基础加载配置文件接口
type DefaultConfig interface {
	GetConfigKey() string
	GetConfigValue() map[string]map[string]string
}
