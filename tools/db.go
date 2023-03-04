package tools

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strconv"
	"time"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {

	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	mysqlConfig := viper.Sub("mysql")

	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	host := mysqlConfig.GetString("host")
	port := mysqlConfig.GetInt("port")
	dbName := mysqlConfig.GetString("dbname")
	username := mysqlConfig.GetString("username")
	charset := mysqlConfig.GetString("charset")
	parseTime := mysqlConfig.GetString("parseTime")
	loc := mysqlConfig.GetString("loc")
	timeout := strconv.Itoa(mysqlConfig.GetInt("timeout")) + "s"
	maxIdleConns := mysqlConfig.GetInt("max-idle-conns")
	maxOpenConns := mysqlConfig.GetInt("max-open-conns")
	connMaxLifeTime := time.Duration(viper.GetInt("conn-max-life-time"))

	dsn := fmt.Sprintf("%s@(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s&timeout=%s", username, host, port, dbName, charset, parseTime, loc, timeout)
	log.Printf("mysql数据库的连接url==>%s", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		// 表名前缀，`Product` 的表名应该是 `t_products`， 默认不起用
		//TablePrefix: "",
		// 使用单数表名，启用该选项，默认启用，此时，`User` 的表名应该是 `t_user`
		SingularTable: true}})
	if err != nil {
		fmt.Println(err)
	}
	mysqlConn, _ := db.DB()
	//延时关闭数据库连接
	// FIXME  这个地方加上defer会报错 sql: database is closed,先记一下
	//defer mysqlConn.Close()
	mysqlConn.SetMaxIdleConns(maxIdleConns)
	mysqlConn.SetMaxOpenConns(maxOpenConns)
	mysqlConn.SetConnMaxLifetime(connMaxLifeTime)
	data, _ := json.Marshal(mysqlConn.Stats()) //获得当前的SQL配置情况
	fmt.Println(string(data))

}

func GetDB() *gorm.DB {
	return db
}
