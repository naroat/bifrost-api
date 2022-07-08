package utils

import (
	"bifrost/server/global"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// 连接库
func Connection() *mongo.Client {
	//连接mongodb
	config := global.GVA_CONFIG.MongoDB
	opts := options.Client().ApplyURI("mongodb://" + config.Username + ":" + config.Password + "@" + config.Path + ":" + config.Port + "/" + config.Dbname)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}
	//database := client.Database(config.Dbname)
	return client
}

// 设置表
func Table(database *mongo.Database, table string) *mongo.Collection {
	return database.Collection(table)
}

func FindOne(connection mongo.Client, filter bson.M) {
	//定义存储查询结果的值
	/*var result Student
	//查询
	err = connection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}*/
}
