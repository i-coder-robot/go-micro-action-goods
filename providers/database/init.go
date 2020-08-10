package database

import (
	"fmt"
	"github.com/i-coder-robot/go-micro-action-core/env"
	tgorm "github.com/i-coder-robot/go-micro-action-core/gorm"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

var (
	DB *gorm.DB
)

type DBConf struct {
	Driver string
	// Host 主机地址
	Host string
	// Port 主机端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// DbName 数据库名称
	DbName string
	// Charset 数据库编码
	Charset string
}


func init() {
	fmt.Println("数据库 init")
	var err error
	conf := &tgorm.Config{
		Driver: env.Getenv("DB_DRIVER", "mysql"),
		// Host 主机地址
		Host: env.Getenv("DB_HOST", "127.0.0.1"),
		// Port 主机端口
		Port: env.Getenv("DB_PORT", "3306"),
		// User 用户名
		User: env.Getenv("DB_USER", "root"),
		// Password 密码
		Password: env.Getenv("DB_PASSWORD", "smartwell"),
		// DbName 数据库名称
		DbName: env.Getenv("DB_NAME", "user"),
		// Charset 数据库编码
		Charset: env.Getenv("DB_CHARSET", "utf8"),
	}
	DB, err = tgorm.Connection(conf)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	fmt.Println("数据库 init 结束...")
}
