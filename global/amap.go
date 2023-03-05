package global

import (
	"fmt"
	"github.com/spf13/viper"
)

var AMapKey *string

func init() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	aMapConfig := viper.Sub("amap")

	key := aMapConfig.GetString("key")

	AMapKey = &key
}
