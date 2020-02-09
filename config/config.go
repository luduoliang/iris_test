package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var (
	Conf = NewConfig()
)

//返回配置，单例
func NewConfig() *toml.Tree {
	config, err := toml.LoadFile("./config/config.toml")
	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}
	return config
}


//获取字符串配置
func GetString(key string, default_value string) string {
	if(Conf.Has(key)){
		return Conf.Get(key).(string)
	}else{
		return default_value
	}
}


//获取bool类型配置
func GetBool(key string, default_value bool) bool {
	if(Conf.Has(key)){
		return Conf.Get(key).(bool)
	}else{
		return default_value
	}
}


//获取int配置
func GetInt(key string, default_value int) int {
	if(Conf.Has(key)){
		return Conf.Get(key).(int)
	}else{
		return default_value
	}
}


//获取int64配置
func GetInt64(key string, default_value int64) int64 {
	if(Conf.Has(key)){
		return Conf.Get(key).(int64)
	}else{
		return default_value
	}
}
