package config

import (
	"fmt"
	"trpc.group/trpc-go/trpc-go/plugin"
)

const (
	pluginName = "democonfig"
	pluginType = "demoapp"
)

func init() {
	plugin.Register(pluginName, DefaultAppConfig)
}

var DefaultAppConfig = &AppConfig{}

// AppConfig 自定义插件配置
type AppConfig struct {
	config CustomConfig
}

type CustomConfig struct {
	Test    string `yaml:"test"`
	TestObj struct {
		Key1 string `yaml:"key1"`
		Key2 bool   `yaml:"key2"`
		Key3 int32  `yaml:"key3"`
	} `yaml:"test_obj"`
}

// Type 插件类型
func (custom *AppConfig) Type() string {
	return pluginType
}

// Setup 设置 setup
func (custom *AppConfig) Setup(name string, decoder plugin.Decoder) error {

	if err := decoder.Decode(&DefaultAppConfig.config); err != nil {
		return err
	}

	return nil
}

// Record record print string
// 1)可以通过查询公共函数, 2)可以通过查询全局变量
func Record() {
	fmt.Printf("key1 : %s, key2 : %t, key3 : %d \n", DefaultAppConfig.config.TestObj.Key1, DefaultAppConfig.config.TestObj.Key2, DefaultAppConfig.config.TestObj.Key3)
}
