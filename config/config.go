/*
 * @Author: GG
 * @Date: 2022-10-26 15:55:13
 * @LastEditTime: 2022-10-26 16:07:40
 * @LastEditors: GG
 * @Description: config
 * @FilePath: \shop-api\config\config.go
 *
 */
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfgReader *configReader

type Configuration struct {
	DatabaseSettings
	JwtSettings
}

// 数据库配置
type DatabaseSettings struct {
	DatabaseURI  string
	DatabaseName string
	Username     string
	Password     string
}

// jwt配置
type JwtSettings struct {
	SecretKey string
}

// 配置读取
type configReader struct {
	configFile string
	v          *viper.Viper
}

// 实例化configReader
func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}

// 获取所有配置
func GettAllConfigValues(configFile string) (configuation *Configuration, err error) {
	newConfigReader(configFile)

	// 读取配置
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("配置读取失败：%s", err)
		return nil, err
	}

	if err = cfgReader.v.Unmarshal(&configuation); err != nil {
		fmt.Printf("解析配置文件到结构体失败：%s", err)
		return nil, err
	}

	return configuation, err
}
