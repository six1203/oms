package global

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"order/global/logger"
)

var RedisClient *redis.Client

// loadRedisConfig 从配置文件中读取Redis配置参数
func loadRedisConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v.Sub("redis"), nil
}

// newRedisClient 初始化Redis客户端
func newRedisClient() *redis.Client {
	redisConfig, err := loadRedisConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	host := redisConfig.GetString("host")
	port := redisConfig.GetInt("port")
	pwd := redisConfig.GetString("password")
	db := redisConfig.GetInt("db")

	addr := fmt.Sprintf("%s:%d", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	logger.Info("redis连接成功")
	return client
}

func init() {
	RedisClient = newRedisClient()
}
