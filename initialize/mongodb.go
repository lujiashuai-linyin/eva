package initialize

import (
	"context"
	"eva/global"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func Mongodb() {
	// 设置连接选项
	//global.EVA_CONFIG.Mongodb
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		global.EVA_LOG.Error("mongodb connect failed, err:", zap.Error(err))
	}

	// 检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		global.EVA_LOG.Error("mongodb connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("Connected to MongoDB!")
		global.EVA_MongoDB = client
	}
}
