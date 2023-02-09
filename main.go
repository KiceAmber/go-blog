package main

import (
	"fmt"
	"go-blog/config"
	"go-blog/dao/mysql"
	"go-blog/log"
	"go-blog/pkg/snowflake"
	"go-blog/router"
	"os"

	"go.uber.org/zap"
)

// Go Web 开发通用脚手架

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Need config File arguments...")
		return
	}
	// 1. 加载配置文件
	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
		return
	}

	// 初始化日志
	if err := log.Init(config.Conf.LogConfig); err != nil {
		fmt.Printf("init log failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("log init success...")

	// 初始化 MySQL 连接
	if err := mysql.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化 Redis 连接
	// if err := redis.Init(config.Conf.RedisConfig); err != nil {
	// 	fmt.Printf("init redis failed, err:%v\n", err)
	// 	return
	// }
	// defer redis.Close()

	// 初始化雪花算法
	if err := snowflake.Init(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 注册路由 router
	r := router.Init()

	// 启动服务(优雅关机)
	router.Setup(r)
}
