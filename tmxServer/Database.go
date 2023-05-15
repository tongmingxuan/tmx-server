// Package tmxServer /*
package tmxServer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

type DatabasePool struct {
	DbConfig map[string]map[string]string
	Pool     sync.Map
}

func (databasePool *DatabasePool) Key() string {
	return "database"
}

func (databasePool *DatabasePool) Handle() *BaseFrame {

	frameConfig := (*Config)(unsafe.Pointer(frameContainer.Get(new(Config).Key())))

	databaseConfig, ok := frameConfig.BaseConfigMap["database"]

	if ok != true {
		panic("数据库配置异常")
	}

	databasePool.DbConfig = databaseConfig

	for connectionName, connectionConfig := range databaseConfig {
		databasePool.CreateConnection(connectionName, connectionConfig)
	}

	return (*BaseFrame)(unsafe.Pointer(databasePool))
}

func (databasePool *DatabasePool) CreateConnection(connectionName string, connectionConfig map[string]string) {
	userName := connectionConfig["db_username"]
	passWord := connectionConfig["db_password"]
	host := connectionConfig["db_host"]
	port := connectionConfig["db_port"]
	dbname := connectionConfig["db_database"]

	maxConnection, _ := strconv.Atoi(connectionConfig["max_connection"])
	maxOpenConnection, _ := strconv.Atoi(connectionConfig["max_open_connection"])

	dsn := userName + ":" + passWord + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("CreateConnection:" + connectionName + ":异常:" + err.Error())
	}

	sqlDB, sqlErr := db.DB()

	if sqlErr != nil {
		panic("CreateConnection:" + connectionName + "获取链接异常:" + sqlErr.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxOpenConnection)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(maxConnection)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(60 * time.Second)

	databasePool.Pool.Store(connectionName, db)
}
