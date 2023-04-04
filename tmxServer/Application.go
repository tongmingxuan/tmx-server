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

// ApplicationStart
// @Description: 基础框架启动
// @param []BaseFrame 核心服务切片
// @param []DefaultConfig 配置切片
func ApplicationStart(frameList []BaseFrame, defaultConfig []DefaultConfig) *Container {
	frameContainer = CreateContainer()

	config := &Config{}

	frameContainer.Set(config.Key(), config.Handle())

	for _, item := range defaultConfig {
		key := item.GetConfigKey()

		val := item.GetConfigValue()

		config.SetBaseConfigMap(key, val)
	}

	for _, frame := range frameList {
		obj := frame

		frameContainer.Set(obj.Key(), obj.Handle())
	}

	return frameContainer
}
