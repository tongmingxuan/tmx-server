// Package tmxServer /*
package tmxServer

import (
	"fmt"
	"github.com/spf13/viper"
	"unsafe"
)

type Config struct {
	// ENV配置
	ConfigMap map[string]interface{}
	//config目录下 配置文件配置
	BaseConfigMap map[string]map[string]map[string]string
}

func (config *Config) Key() string {
	return "config"
}

func (config *Config) Handle() *BaseFrame {
	viper.SetConfigType("yml")

	viper.SetConfigFile("./config.yml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
		panic("Viper:func:error:" + err.Error())
	}

	allSettings := viper.AllSettings()

	config.ConfigMap = make(map[string]interface{}, len(allSettings))

	for key, value := range allSettings {
		configKey := AntToString(key)

		configValue := AntToString(value)

		SetEnv(configKey, configValue)

		config.ConfigMap[configKey] = configValue
	}

	return (*BaseFrame)(unsafe.Pointer(config))
}

func (config *Config) SetBaseConfigMap(configFileKey string, configValue map[string]map[string]string) {
	if config.BaseConfigMap == nil {
		config.BaseConfigMap = make(map[string]map[string]map[string]string)
	}

	config.BaseConfigMap[configFileKey] = configValue
}

func (config *Config) GetBaseConfigMap(configFileKey, key string) interface{} {
	value, ok := config.BaseConfigMap[configFileKey]

	if ok != true {
		return nil
	}

	res, _ok := value[key]

	if _ok != true {
		return nil
	}

	return res
}

func (config *Config) GetConfigMap(key string) interface{} {
	value, ok := config.ConfigMap[key]

	if ok {
		return value
	}
	return nil
}

func (config *Config) SetConfig(key string, value interface{}) {
	config.ConfigMap[key] = value
}
