package model

import "context"

func Init(ctx context.Context) {
	// 初始化Redis客户端
	initRedis(ctx)
	InitMysql(ctx)
}
